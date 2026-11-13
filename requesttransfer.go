// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	"github.com/MercuryTechnologies/mercury-go/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// Manage transfer approval requests
//
// RequestTransferService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRequestTransferService] method instead.
type RequestTransferService struct {
	options []option.RequestOption
}

// NewRequestTransferService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRequestTransferService(opts ...option.RequestOption) (r RequestTransferService) {
	r = RequestTransferService{}
	r.options = opts
	return
}

// Retrieve a specific transfer approval request by ID.
func (r *RequestTransferService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *TransferMoneyApprovalRequestResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if requestID == "" {
		err = errors.New("missing required requestId parameter")
		return nil, err
	}
	path := fmt.Sprintf("request-transfer/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a transfer between two accounts in your organization that will require
// approval based on your organization's approval policies.
func (r *RequestTransferService) RequestTransfer(ctx context.Context, body RequestTransferRequestTransferParams, opts ...option.RequestOption) (res *TransferMoneyApprovalRequestResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "request-transfer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of transfer approval requests for the authenticated
// organization. Supports filtering by source account, destination account, and
// status.
func (r *RequestTransferService) GetRequestTransfer(ctx context.Context, query RequestTransferGetRequestTransferParams, opts ...option.RequestOption) (res *RequestTransferGetRequestTransferResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "request-transfer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// An approval request for a transfer between accounts.
type TransferMoneyApprovalRequestResponse struct {
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount" api:"required"`
	// When the request was created.
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// ID for a Mercury account.
	DestinationAccountID string `json:"destinationAccountId" api:"required" format:"uuid"`
	// ID for the user
	RequestedByUserID string `json:"requestedByUserId" api:"required" format:"uuid"`
	RequestID         string `json:"requestId" api:"required"`
	// Approval decisions recorded on this request, oldest first.
	Reviews []PaymentApprovalReview `json:"reviews" api:"required"`
	// ID for a Mercury account.
	SourceAccountID string `json:"sourceAccountId" api:"required" format:"uuid"`
	// Any of "pendingApproval", "approved", "rejected", "cancelled".
	Status ReviewRequestStatus `json:"status" api:"required"`
	// Memo that appears on the resulting transaction.
	Memo string `json:"memo" api:"nullable"`
	// Internal note recorded on the request. Not shown to any counterparty.
	Note string `json:"note" api:"nullable"`
	// Number of approvals required before the transfer is sent.
	NumberOfApproversRequired int64 `json:"numberOfApproversRequired" api:"nullable"`
	// Whether the requester may approve this request. Omitted whenever
	// numberOfApproversRequired is omitted.
	RequesterMayApprove bool `json:"requesterMayApprove" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount                    respjson.Field
		CreatedAt                 respjson.Field
		DestinationAccountID      respjson.Field
		RequestedByUserID         respjson.Field
		RequestID                 respjson.Field
		Reviews                   respjson.Field
		SourceAccountID           respjson.Field
		Status                    respjson.Field
		Memo                      respjson.Field
		Note                      respjson.Field
		NumberOfApproversRequired respjson.Field
		RequesterMayApprove       respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferMoneyApprovalRequestResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferMoneyApprovalRequestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page of transfer approval requests.
type RequestTransferGetRequestTransferResponse struct {
	Page     RequestTransferGetRequestTransferResponsePage `json:"page" api:"required"`
	Requests []TransferMoneyApprovalRequestResponse        `json:"requests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		Requests    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequestTransferGetRequestTransferResponse) RawJSON() string { return r.JSON.raw }
func (r *RequestTransferGetRequestTransferResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequestTransferGetRequestTransferResponsePage struct {
	// ID for the transfer money approval request
	NextPage string `json:"nextPage" format:"uuid"`
	// ID for the transfer money approval request
	PreviousPage string `json:"previousPage" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPage     respjson.Field
		PreviousPage respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequestTransferGetRequestTransferResponsePage) RawJSON() string { return r.JSON.raw }
func (r *RequestTransferGetRequestTransferResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequestTransferRequestTransferParams struct {
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount" api:"required"`
	// ID for a Mercury account.
	DestinationAccountID string `json:"destinationAccountId" api:"required" format:"uuid"`
	IdempotencyKey       string `json:"idempotencyKey" api:"required"`
	// ID for a Mercury account.
	SourceAccountID string            `json:"sourceAccountId" api:"required" format:"uuid"`
	Note            param.Opt[string] `json:"note,omitzero"`
	paramObj
}

func (r RequestTransferRequestTransferParams) MarshalJSON() (data []byte, err error) {
	type shadow RequestTransferRequestTransferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RequestTransferRequestTransferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RequestTransferGetRequestTransferParams struct {
	// ID for a Mercury account.
	DestinationAccountID param.Opt[string] `query:"destinationAccountId,omitzero" format:"uuid" json:"-"`
	// The ID of the transfer approval request to end the page before (exclusive). When
	// provided, results will end just before this ID and work backwards. Use this for
	// reverse pagination or to retrieve previous pages. Cannot be combined with
	// start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// ID for a Mercury account.
	SourceAccountID param.Opt[string] `query:"sourceAccountId,omitzero" format:"uuid" json:"-"`
	// The ID of the transfer approval request to start the page after (exclusive).
	// When provided, results will begin with the transfer approval request immediately
	// following this ID. Use this for standard forward pagination to get the next page
	// of results. Cannot be combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// The ID of the transfer approval request to start the page at (inclusive). Cannot
	// be combined with start_after or end_before.
	StartAt param.Opt[string] `query:"start_at,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order RequestTransferGetRequestTransferParamsOrder `query:"order,omitzero" json:"-"`
	// Any of "pendingApproval", "approved", "rejected", "cancelled".
	Status RequestTransferGetRequestTransferParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RequestTransferGetRequestTransferParams]'s query parameters
// as `url.Values`.
func (r RequestTransferGetRequestTransferParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type RequestTransferGetRequestTransferParamsOrder string

const (
	RequestTransferGetRequestTransferParamsOrderAsc  RequestTransferGetRequestTransferParamsOrder = "asc"
	RequestTransferGetRequestTransferParamsOrderDesc RequestTransferGetRequestTransferParamsOrder = "desc"
)

type RequestTransferGetRequestTransferParamsStatus string

const (
	RequestTransferGetRequestTransferParamsStatusPendingApproval RequestTransferGetRequestTransferParamsStatus = "pendingApproval"
	RequestTransferGetRequestTransferParamsStatusApproved        RequestTransferGetRequestTransferParamsStatus = "approved"
	RequestTransferGetRequestTransferParamsStatusRejected        RequestTransferGetRequestTransferParamsStatus = "rejected"
	RequestTransferGetRequestTransferParamsStatusCancelled       RequestTransferGetRequestTransferParamsStatus = "cancelled"
)
