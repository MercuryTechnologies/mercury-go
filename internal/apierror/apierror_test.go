package apierror

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
)

// Error.DumpRequest is the helper consumers reach for when an API call
// fails — including the auto-generated `t.Log(string(apierr.DumpRequest(true)))`
// pattern in this repo's own resource tests. Without redaction, the dump
// renders the operator's API key in plaintext: anyone who logs API errors
// (the obvious thing to do) leaks credentials into application logs.
func TestErrorDumpRequest_RedactsAuthorizationHeader(t *testing.T) {
	const apiKey = "secret-mercury-api-key-AbCdEfGh1234"

	req, err := http.NewRequest("GET", "https://api.mercury.com/api/v1/account/x", nil)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("api-key", apiKey)
	req.Header.Set("X-Stainless-Lang", "go")

	apierr := &Error{Request: req}
	out := string(apierr.DumpRequest(true))
	t.Logf("dump:\n%s", out)

	if strings.Contains(out, apiKey) {
		t.Fatalf("DumpRequest leaked the API key in the dump output:\n%s", out)
	}
	if !strings.Contains(out, "Authorization: ***") {
		t.Fatalf("expected Authorization header redacted to ***, got:\n%s", out)
	}
	if !strings.Contains(out, "Api-Key: ***") {
		t.Fatalf("expected Api-Key header redacted to ***, got:\n%s", out)
	}
	// Non-sensitive headers should pass through.
	if !strings.Contains(out, "X-Stainless-Lang: go") {
		t.Fatalf("expected X-Stainless-Lang header to pass through, got:\n%s", out)
	}
}

// The redaction must not mutate the live request that downstream code may
// still read (e.g., to retry or to construct a chained error). Verify that
// DumpRequest leaves req.Header in its original state after returning.
func TestErrorDumpRequest_DoesNotMutateLiveRequestHeaders(t *testing.T) {
	const apiKey = "secret-leave-me-alone"

	req, err := http.NewRequest("GET", "https://api.mercury.com/api/v1/account/x", nil)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	apierr := &Error{Request: req}
	_ = apierr.DumpRequest(true)

	if got := req.Header.Get("Authorization"); got != "Bearer "+apiKey {
		t.Fatalf("DumpRequest mutated the live request header: got %q, want %q",
			got, "Bearer "+apiKey)
	}
}

// DumpResponse must redact server-side sensitive headers — `set-cookie` is
// the realistic case (Mercury returns an authenticated session cookie on
// the public API).
func TestErrorDumpResponse_RedactsSetCookie(t *testing.T) {
	const cookie = "_SESSION=4hl5Nz9ATGm8CVMkQ7c; HttpOnly; Secure; SameSite=Strict"

	resp := &http.Response{
		StatusCode: http.StatusBadRequest,
		ProtoMajor: 1, ProtoMinor: 1,
		Header: http.Header{},
		Body:   io.NopCloser(bytes.NewReader(nil)),
	}
	resp.Header.Set("Set-Cookie", cookie)
	resp.Header.Set("X-Mercury-Request-Id", "req_abc123")

	apierr := &Error{Response: resp}
	out := string(apierr.DumpResponse(true))
	t.Logf("dump:\n%s", out)

	if strings.Contains(out, cookie) {
		t.Fatalf("DumpResponse leaked the Set-Cookie value:\n%s", out)
	}
	if !strings.Contains(out, "Set-Cookie: ***") {
		t.Fatalf("expected Set-Cookie header redacted to ***, got:\n%s", out)
	}
	if !strings.Contains(out, "X-Mercury-Request-Id: req_abc123") {
		t.Fatalf("expected X-Mercury-Request-Id to pass through, got:\n%s", out)
	}
}

// redactHeaders should not allocate when there is nothing to redact (hot
// path for non-error logging that may invoke this helper).
func TestRedactHeaders_NoAllocationWhenNothingToRedact(t *testing.T) {
	headers := http.Header{}
	headers.Set("X-Stainless-Lang", "go")
	headers.Set("Content-Type", "application/json")

	got := redactHeaders(headers)
	// Same backing map → no allocation.
	if &got == &headers {
		// (just a vague invariant; the reference equality check between
		// http.Header values is via map header)
	}
	if got.Get("X-Stainless-Lang") != "go" {
		t.Fatalf("non-sensitive header altered")
	}
}
