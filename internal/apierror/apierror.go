// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package apierror

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// sensitiveLogHeaders are redacted before the request or response is dumped
// for diagnostic display. Keep this in sync with option.sensitiveLogHeaders.
var sensitiveLogHeaders = []string{"authorization", "api-key", "x-api-key", "cookie", "set-cookie"}

// redactHeaders returns headers with sensitiveLogHeaders replaced by "***".
// Returns the input unmodified when no redaction is needed (no allocation in
// the common case).
func redactHeaders(headers http.Header) http.Header {
	var redacted http.Header
	for _, name := range sensitiveLogHeaders {
		values := headers.Values(name)
		if len(values) == 0 {
			continue
		}
		if redacted == nil {
			redacted = headers.Clone()
		}
		redacted.Del(name)
		for range values {
			redacted.Add(name, "***")
		}
	}
	if redacted == nil {
		return headers
	}
	return redacted
}

// Error represents an error that originates from the API, i.e. when a request is
// made and the API returns a response with a HTTP status code. Other errors are
// not wrapped by this SDK.
type Error struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

// Returns the unmodified JSON received from the API
func (r Error) RawJSON() string { return r.JSON.raw }
func (r *Error) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r *Error) Error() string {
	// Attempt to re-populate the response body
	return fmt.Sprintf("%s %q: %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), r.JSON.raw)
}

// DumpRequest returns the dump of the HTTP request that produced this error,
// with sensitive headers (Authorization, api-key, cookie) replaced by "***".
// Without redaction the dump renders the operator's API key in plaintext —
// the auto-generated tests in this repo demonstrate the typical consumer
// pattern of logging this output, so the redaction matters in real code.
func (r *Error) DumpRequest(body bool) []byte {
	if r.Request.GetBody != nil {
		r.Request.Body, _ = r.Request.GetBody()
	}
	origHeaders := r.Request.Header
	r.Request.Header = redactHeaders(origHeaders)
	defer func() { r.Request.Header = origHeaders }()
	out, _ := httputil.DumpRequestOut(r.Request, body)
	return out
}

// DumpResponse returns the dump of the HTTP response that produced this
// error, with sensitive headers (set-cookie etc.) replaced by "***".
func (r *Error) DumpResponse(body bool) []byte {
	origHeaders := r.Response.Header
	r.Response.Header = redactHeaders(origHeaders)
	defer func() { r.Response.Header = origHeaders }()
	out, _ := httputil.DumpResponse(r.Response, body)
	return out
}
