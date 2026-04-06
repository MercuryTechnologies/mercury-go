// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/mercury-go/internal/apiquery"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/pagination"
	"github.com/stainless-sdks/mercury-go/packages/param"
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

// Retrieve a paginated list of send money approval requests for the authenticated
// organization. Supports filtering by account and status.
func (r *RequestSendMoneyService) ListSendMoneyRequests(ctx context.Context, query RequestSendMoneyListSendMoneyRequestsParams, opts ...option.RequestOption) (res *pagination.CursorIDRequestSendMoney[SendMoneyApproval], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "request-send-money"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a paginated list of send money approval requests for the authenticated
// organization. Supports filtering by account and status.
func (r *RequestSendMoneyService) ListSendMoneyRequestsAutoPaging(ctx context.Context, query RequestSendMoneyListSendMoneyRequestsParams, opts ...option.RequestOption) *pagination.CursorIDRequestSendMoneyAutoPager[SendMoneyApproval] {
	return pagination.NewCursorIDRequestSendMoneyAutoPager(r.ListSendMoneyRequests(ctx, query, opts...))
}

type RequestSendMoneyListSendMoneyRequestsParams struct {
	// ID for a Mercury account.
	AccountID param.Opt[string] `query:"accountId,omitzero" format:"uuid" json:"-"`
	// The ID of the send money approval request to end the page before (exclusive).
	// When provided, results will end just before this ID and work backwards. Use this
	// for reverse pagination or to retrieve previous pages. Cannot be combined with
	// start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the send money approval request to start the page after (exclusive).
	// When provided, results will begin with the send money approval request
	// immediately following this ID. Use this for standard forward pagination to get
	// the next page of results. Cannot be combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Any of "pendingApproval", "approved", "rejected", "cancelled".
	Status RequestSendMoneyListSendMoneyRequestsParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RequestSendMoneyListSendMoneyRequestsParams]'s query
// parameters as `url.Values`.
func (r RequestSendMoneyListSendMoneyRequestsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RequestSendMoneyListSendMoneyRequestsParamsStatus string

const (
	RequestSendMoneyListSendMoneyRequestsParamsStatusPendingApproval RequestSendMoneyListSendMoneyRequestsParamsStatus = "pendingApproval"
	RequestSendMoneyListSendMoneyRequestsParamsStatusApproved        RequestSendMoneyListSendMoneyRequestsParamsStatus = "approved"
	RequestSendMoneyListSendMoneyRequestsParamsStatusRejected        RequestSendMoneyListSendMoneyRequestsParamsStatus = "rejected"
	RequestSendMoneyListSendMoneyRequestsParamsStatusCancelled       RequestSendMoneyListSendMoneyRequestsParamsStatus = "cancelled"
)
