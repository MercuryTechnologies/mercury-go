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

// Manage send money approval requests
//
// RequestSendMoneyService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestSendMoneyService] method instead.
type RequestSendMoneyService struct {
	Options []option.RequestOption
}

// NewRequestSendMoneyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRequestSendMoneyService(opts ...option.RequestOption) (r RequestSendMoneyService) {
	r = RequestSendMoneyService{}
	r.Options = opts
	return
}

// Get send money approval request by ID
func (r *RequestSendMoneyService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *SendMoneyApproval, err error) {
	opts = slices.Concat(r.Options, opts)
	if requestID == "" {
		err = errors.New("missing required requestId parameter")
		return nil, err
	}
	path := fmt.Sprintf("request-send-money/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
