// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
)

// Download account statements
//
// StatementService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatementService] method instead.
type StatementService struct {
	Options []option.RequestOption
}

// NewStatementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStatementService(opts ...option.RequestOption) (r StatementService) {
	r = StatementService{}
	r.Options = opts
	return
}

// Downloads a PDF file for the specified account statement. The response includes
// a Content-Disposition header for proper file download handling. Returns binary
// PDF data.
func (r *StatementService) Download(ctx context.Context, statementID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if statementID == "" {
		err = errors.New("missing required statementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("statements/%s/pdf", statementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
