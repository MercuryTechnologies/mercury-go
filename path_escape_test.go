package mercury_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mercury "github.com/MercuryTechnologies/mercury-go"
	"github.com/MercuryTechnologies/mercury-go/option"
)

// captureURL spins up a local HTTP server that records the wire-form path
// of one inbound request. Returns the EscapedPath (encoded as it traveled
// the wire) and the URL.Path (after percent-decoding by Go's parser).
func captureURL(t *testing.T) (*httptest.Server, *capturedURL) {
	t.Helper()
	cap := &capturedURL{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cap.escapedPath = r.URL.EscapedPath()
		cap.path = r.URL.Path
		cap.method = r.Method
		w.WriteHeader(http.StatusNotFound)
	}))
	t.Cleanup(srv.Close)
	return srv, cap
}

type capturedURL struct {
	method      string
	escapedPath string
	path        string
}

// A path-param value containing "/" used to interpolate raw into the
// outbound URL. Go's net/url parser then ran RFC 3986 dot-segment collapse
// on the merged path — so a value of "x/../../webhooks/wh_attacker_id"
// passed to Accounts.Get rewrote the wire path from /account/X to a
// completely different endpoint (/webhooks/wh_attacker_id) under the
// operator's API key. After url.PathEscape, "/" is encoded as "%2F" and
// the dot-segments arrive at the API as literal bytes inside a single
// path segment that the Mercury API rejects as an invalid identifier.
func TestPathParamEscape_AccountsGet(t *testing.T) {
	srv, cap := captureURL(t)
	client := mercury.NewClient(
		option.WithBaseURL(srv.URL),
		option.WithAPIKey("operator-api-key"),
	)

	const malicious = "x/../../webhooks/wh_attacker_id"
	_, _ = client.Accounts.Get(context.Background(), malicious)

	if !strings.HasPrefix(cap.escapedPath, "/account/") {
		t.Fatalf("path traversal escaped to a sibling endpoint: %q", cap.escapedPath)
	}
	if !strings.Contains(cap.escapedPath, "%2F") {
		t.Fatalf("expected slash to be percent-encoded on the wire, got %q", cap.escapedPath)
	}
}

// Webhook is the highest-impact endpoint because Webhooks.Update can
// retarget delivery URLs, and Webhooks.Verify changes the verification
// secret. A traversal there would let the attacker rewire the operator's
// outbound webhook payloads.
func TestPathParamEscape_WebhooksUpdateAndVerify(t *testing.T) {
	for _, name := range []string{"Update", "Verify", "Get", "Delete"} {
		name := name
		t.Run(name, func(t *testing.T) {
			srv, cap := captureURL(t)
			client := mercury.NewClient(
				option.WithBaseURL(srv.URL),
				option.WithAPIKey("operator-api-key"),
			)

			const malicious = "x/../../account/operator_acct"
			ctx := context.Background()
			switch name {
			case "Get":
				_, _ = client.Webhooks.Get(ctx, malicious)
			case "Update":
				_, _ = client.Webhooks.Update(ctx, malicious, mercury.WebhookUpdateParams{})
			case "Delete":
				_ = client.Webhooks.Delete(ctx, malicious)
			case "Verify":
				_ = client.Webhooks.Verify(ctx, malicious, mercury.WebhookVerifyParams{})
			}

			if !strings.HasPrefix(cap.escapedPath, "/webhooks/") {
				t.Fatalf("%s: traversal escaped to %q (expected to stay under /webhooks/)",
					name, cap.escapedPath)
			}
		})
	}
}

// Sample of every path-param-bearing endpoint family. We assert each
// stays under its declared prefix on the wire when given a traversal-
// shaped id. The wire form must keep the dot-segments inert (percent-
// encoded slashes) so the API server cannot interpret them as separators.
func TestPathParamEscape_EndpointFamilies(t *testing.T) {
	const malicious = "x/../../intruder"
	cases := []struct {
		name   string
		prefix string
		invoke func(context.Context, mercury.Client)
	}{
		{"Customers.Get", "/ar/customers/", func(ctx context.Context, c mercury.Client) { _, _ = c.Customers.Get(ctx, malicious) }},
		{"Customers.Update", "/ar/customers/", func(ctx context.Context, c mercury.Client) { _, _ = c.Customers.Update(ctx, malicious, mercury.CustomerUpdateParams{}) }},
		{"Customers.Delete", "/ar/customers/", func(ctx context.Context, c mercury.Client) { _ = c.Customers.Delete(ctx, malicious) }},
		{"Invoices.Get", "/ar/invoices/", func(ctx context.Context, c mercury.Client) { _, _ = c.Invoices.Get(ctx, malicious) }},
		{"Invoices.Cancel", "/ar/invoices/", func(ctx context.Context, c mercury.Client) { _ = c.Invoices.Cancel(ctx, malicious) }},
		{"Invoices.Download", "/ar/invoices/", func(ctx context.Context, c mercury.Client) { _, _ = c.Invoices.Download(ctx, malicious) }},
		{"Events.Get", "/events/", func(ctx context.Context, c mercury.Client) { _, _ = c.Events.Get(ctx, malicious) }},
		{"Users.Get", "/users/", func(ctx context.Context, c mercury.Client) { _, _ = c.Users.Get(ctx, malicious) }},
		{"Recipients.Get", "/recipient/", func(ctx context.Context, c mercury.Client) { _, _ = c.Recipients.Get(ctx, malicious) }},
		{"Transactions.Get", "/transaction/", func(ctx context.Context, c mercury.Client) { _, _ = c.Transactions.Get(ctx, malicious) }},
		{"Safes.Get", "/safes/", func(ctx context.Context, c mercury.Client) { _, _ = c.Safes.Get(ctx, malicious) }},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			srv, cap := captureURL(t)
			client := mercury.NewClient(
				option.WithBaseURL(srv.URL),
				option.WithAPIKey("operator-api-key"),
			)
			tc.invoke(context.Background(), client)

			if !strings.HasPrefix(cap.escapedPath, tc.prefix) {
				t.Fatalf("%s: traversal escaped to %q (expected prefix %q)",
					tc.name, cap.escapedPath, tc.prefix)
			}
			if !strings.Contains(cap.escapedPath, "%2F") {
				t.Fatalf("%s: expected percent-encoded slashes on the wire, got %q",
					tc.name, cap.escapedPath)
			}
		})
	}
}

// Legitimate Mercury identifier shapes (alphanumeric + hyphens, UUID,
// prefixed) must continue to flow through unchanged. url.PathEscape is a
// no-op for chars in the unreserved set; this test pins that legitimate
// IDs aren't accidentally rewritten.
func TestPathParamEscape_LegitimateIDsUnchanged(t *testing.T) {
	cases := []string{
		"acct_abc123",
		"wh_zZ-09aA",
		"550e8400-e29b-41d4-a716-446655440000",
	}
	for _, id := range cases {
		id := id
		t.Run(id, func(t *testing.T) {
			srv, cap := captureURL(t)
			client := mercury.NewClient(
				option.WithBaseURL(srv.URL),
				option.WithAPIKey("operator-api-key"),
			)
			_, _ = client.Accounts.Get(context.Background(), id)

			want := "/account/" + id
			if cap.escapedPath != want {
				t.Fatalf("legitimate id rewritten: got %q, want %q", cap.escapedPath, want)
			}
		})
	}
}
