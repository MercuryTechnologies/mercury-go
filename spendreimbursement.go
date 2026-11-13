// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	"github.com/MercuryTechnologies/mercury-go/internal/apiquery"
	shimjson "github.com/MercuryTechnologies/mercury-go/internal/encoding/json"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// Manage Spend Management reimbursements. Automated access is permitted, including
// from apps and AI agents. You're responsible for your API tokens and for the apps
// and agents that use them. See
// [Getting Started](https://docs.mercury.com/docs/getting-started) and sections 2
// (Our Proprietary Rights) and 3.2 (Third-Party AI and Connected Services) of the
// [Terms of Use](https://mercury.com/legal/terms).
//
// SpendReimbursementService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpendReimbursementService] method instead.
type SpendReimbursementService struct {
	options []option.RequestOption
}

// NewSpendReimbursementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSpendReimbursementService(opts ...option.RequestOption) (r SpendReimbursementService) {
	r = SpendReimbursementService{}
	r.options = opts
	return
}

// Submit a new reimbursement as the token's user. Only that user can edit or
// cancel it. Returns the new reimbursement id.
func (r *SpendReimbursementService) New(ctx context.Context, params SpendReimbursementNewParams, opts ...option.RequestOption) (res *APIV1CreateReimbursementResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "spend/reimbursements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve details of a specific reimbursement by its id, including its review and
// payment status. Attachments include an id, a file name, and a presigned download
// URL. The URL expires after 24 hours. Request the reimbursement again for a new
// URL.
func (r *SpendReimbursementService) Get(ctx context.Context, reimbursementID string, opts ...option.RequestOption) (res *SpendReimbursementGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if reimbursementID == "" {
		err = errors.New("missing required reimbursementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("spend/reimbursements/%s", reimbursementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replace the fields of an existing reimbursement. Only reimbursements still in
// review can be edited.
func (r *SpendReimbursementService) Update(ctx context.Context, reimbursementID string, params SpendReimbursementUpdateParams, opts ...option.RequestOption) (res *APIV1CreateReimbursementResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if reimbursementID == "" {
		err = errors.New("missing required reimbursementId parameter")
		return nil, err
	}
	path := fmt.Sprintf("spend/reimbursements/%s", reimbursementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Retrieve a paginated list of reimbursements for the organization ordered by
// creation time, ascending (oldest first) by default. Each reimbursement includes
// its attachments, with an id, a file name, and a presigned download URL. The URL
// expires after 24 hours. Request the reimbursement again for a new URL. Supports
// cursor-based pagination with limit, order, start_at, start_after, and end_before
// query parameters.
func (r *SpendReimbursementService) List(ctx context.Context, query SpendReimbursementListParams, opts ...option.RequestOption) (res *SpendReimbursementListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "spend/reimbursements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Approve a submitted reimbursement. The token's user needs the "Reimbursement
// approval and submission" or "Policy management and approval" permission, and
// access to the organization's reimbursement account. The token needs the
// `reimbursements:manage` scope or Read and Write access. Self-approval works only
// when the organization allows it or has only one admin.
//
// Approved reimbursements pay out in a batch each weekday at 1:49 PM Eastern. A
// payout waits for a later batch when the reimbursement account can't cover it or
// the submitter has no verified payout account.
func (r *SpendReimbursementService) Approve(ctx context.Context, reimbursementID string, body SpendReimbursementApproveParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", body.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if reimbursementID == "" {
		err = errors.New("missing required reimbursementId parameter")
		return err
	}
	path := fmt.Sprintf("spend/reimbursements/%s/approve", reimbursementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Cancel a reimbursement that is still in review.
func (r *SpendReimbursementService) Cancel(ctx context.Context, reimbursementID string, body SpendReimbursementCancelParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", body.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if reimbursementID == "" {
		err = errors.New("missing required reimbursementId parameter")
		return err
	}
	path := fmt.Sprintf("spend/reimbursements/%s/cancel", reimbursementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Deny a submitted reimbursement. The reason is recorded as a reviewer comment.
// The token's user needs the "Reimbursement approval and submission" or "Policy
// management and approval" permission, and access to the organization's
// reimbursement account. The token needs the `reimbursements:manage` scope or Read
// and Write access.
func (r *SpendReimbursementService) Deny(ctx context.Context, reimbursementID string, params SpendReimbursementDenyParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if reimbursementID == "" {
		err = errors.New("missing required reimbursementId parameter")
		return err
	}
	path := fmt.Sprintf("spend/reimbursements/%s/deny", reimbursementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Record that the submitter was paid outside Mercury, for example by check or
// payroll. Mercury won’t pay the reimbursement, and its payment status reads as
// paidOutsideMercury. Denied or canceled reimbursements, and ones with a Mercury
// payment in progress or completed, can’t be marked. The optional note is recorded
// as a reviewer comment. Approving or denying the reimbursement afterwards undoes
// the mark.
func (r *SpendReimbursementService) MarkPaid(ctx context.Context, reimbursementID string, params SpendReimbursementMarkPaidParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if reimbursementID == "" {
		err = errors.New("missing required reimbursementId parameter")
		return err
	}
	path := fmt.Sprintf("spend/reimbursements/%s/mark-paid", reimbursementID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Create a reimbursement (a regular, non-mileage expense).
//
// The properties Amount, Currency, DatePaidFor, MerchantName are required.
type APIV1CreateReimbursementRequestParam struct {
	// The positive amount spent.
	Amount float64 `json:"amount" api:"required"`
	// The currency the money was spent in (e.g. "USD").
	Currency string `json:"currency" api:"required"`
	// When the money was actually spent.
	DatePaidFor time.Time `json:"datePaidFor" api:"required" format:"date"`
	// Where the money was spent, e.g. "Starbucks".
	MerchantName string `json:"merchantName" api:"required"`
	// ID for the category
	CategoryID param.Opt[string] `json:"categoryId,omitzero" format:"uuid"`
	Notes      param.Opt[string] `json:"notes,omitzero"`
	paramObj
}

func (r APIV1CreateReimbursementRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow APIV1CreateReimbursementRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV1CreateReimbursementRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The stable id of a freshly created reimbursement; GET it for the full view.
type APIV1CreateReimbursementResponse struct {
	// ID for the expense
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIV1CreateReimbursementResponse) RawJSON() string { return r.JSON.raw }
func (r *APIV1CreateReimbursementResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A file linked to a reimbursement.
type APIV1ReimbursementAttachment struct {
	// ID for the expense attachment
	ID       string `json:"id" api:"required" format:"uuid"`
	FileName string `json:"fileName" api:"required"`
	// Presigned download URL. Expires after 24 hours. Request the reimbursement again
	// for a new URL.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FileName    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIV1ReimbursementAttachment) RawJSON() string { return r.JSON.raw }
func (r *APIV1ReimbursementAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payout state, separate from review.
//
//   - `pending`: no payout in progress. The reimbursement isn't approved yet, or
//     it's approved and waiting for a payout batch. A canceled payout, or one
//     Mercury retries automatically, also returns to `pending`.
//   - `initiated`: the payout is in progress.
//   - `paid`: the payout completed.
//   - `failed`: the payout failed. Mercury retries in a later batch once the
//     submitter fixes their payout account, unless they've left the organization.
//   - `paidOutsideMercury`: an approver recorded that the submitter was paid outside
//     Mercury, so Mercury won't pay it. Its review status is `approved`.
type ReimbursementPaymentStatus string

const (
	ReimbursementPaymentStatusPending            ReimbursementPaymentStatus = "pending"
	ReimbursementPaymentStatusInitiated          ReimbursementPaymentStatus = "initiated"
	ReimbursementPaymentStatusPaid               ReimbursementPaymentStatus = "paid"
	ReimbursementPaymentStatusFailed             ReimbursementPaymentStatus = "failed"
	ReimbursementPaymentStatusPaidOutsideMercury ReimbursementPaymentStatus = "paidOutsideMercury"
)

// Review state, separate from payment.
//
// - `submitted`: waiting for review.
// - `updatesRequested`: a reviewer asked the submitter for changes.
// - `approved`: a reviewer approved it. `paymentStatus` shows the payout.
// - `denied`: a reviewer denied it.
// - `canceled`: the submitter canceled it.
type ReimbursementReviewStatus string

const (
	ReimbursementReviewStatusSubmitted        ReimbursementReviewStatus = "submitted"
	ReimbursementReviewStatusUpdatesRequested ReimbursementReviewStatus = "updatesRequested"
	ReimbursementReviewStatusApproved         ReimbursementReviewStatus = "approved"
	ReimbursementReviewStatusDenied           ReimbursementReviewStatus = "denied"
	ReimbursementReviewStatusCanceled         ReimbursementReviewStatus = "canceled"
)

// Public view of a single reimbursement, keyed on the stable 'ExpenseId'.
type SpendReimbursementGetResponse struct {
	// ID for the expense
	ID string `json:"id" api:"required" format:"uuid"`
	// The positive amount spent in the submitted currency.
	Amount      float64                        `json:"amount" api:"required"`
	Attachments []APIV1ReimbursementAttachment `json:"attachments" api:"required"`
	Currency    string                         `json:"currency" api:"required"`
	DatePaidFor time.Time                      `json:"datePaidFor" api:"required" format:"date"`
	// Payout state, separate from review.
	//
	//   - `pending`: no payout in progress. The reimbursement isn't approved yet, or
	//     it's approved and waiting for a payout batch. A canceled payout, or one
	//     Mercury retries automatically, also returns to `pending`.
	//   - `initiated`: the payout is in progress.
	//   - `paid`: the payout completed.
	//   - `failed`: the payout failed. Mercury retries in a later batch once the
	//     submitter fixes their payout account, unless they've left the organization.
	//   - `paidOutsideMercury`: an approver recorded that the submitter was paid outside
	//     Mercury, so Mercury won't pay it. Its review status is `approved`.
	//
	// Any of "pending", "initiated", "paid", "failed", "paidOutsideMercury".
	PaymentStatus ReimbursementPaymentStatus `json:"paymentStatus" api:"required"`
	// Review state, separate from payment.
	//
	// - `submitted`: waiting for review.
	// - `updatesRequested`: a reviewer asked the submitter for changes.
	// - `approved`: a reviewer approved it. `paymentStatus` shows the payout.
	// - `denied`: a reviewer denied it.
	// - `canceled`: the submitter canceled it.
	//
	// Any of "submitted", "updatesRequested", "approved", "denied", "canceled".
	ReviewStatus ReimbursementReviewStatus `json:"reviewStatus" api:"required"`
	SubmittedAt  string                    `json:"submittedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// ID for the user
	SubmitterUserID string `json:"submitterUserId" api:"required" format:"uuid"`
	// ID for the category
	CategoryID   string `json:"categoryId" api:"nullable" format:"uuid"`
	MerchantName string `json:"merchantName" api:"nullable"`
	Notes        string `json:"notes" api:"nullable"`
	// ID for this transaction
	TransactionID string `json:"transactionId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Amount          respjson.Field
		Attachments     respjson.Field
		Currency        respjson.Field
		DatePaidFor     respjson.Field
		PaymentStatus   respjson.Field
		ReviewStatus    respjson.Field
		SubmittedAt     respjson.Field
		SubmitterUserID respjson.Field
		CategoryID      respjson.Field
		MerchantName    respjson.Field
		Notes           respjson.Field
		TransactionID   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendReimbursementGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SpendReimbursementGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A page of reimbursements plus the cursor for the next/previous page.
type SpendReimbursementListResponse struct {
	Page           SpendReimbursementListResponsePage            `json:"page" api:"required"`
	Reimbursements []SpendReimbursementListResponseReimbursement `json:"reimbursements" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page           respjson.Field
		Reimbursements respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendReimbursementListResponse) RawJSON() string { return r.JSON.raw }
func (r *SpendReimbursementListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendReimbursementListResponsePage struct {
	// ID for the expense
	NextPage string `json:"nextPage" format:"uuid"`
	// ID for the expense
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
func (r SpendReimbursementListResponsePage) RawJSON() string { return r.JSON.raw }
func (r *SpendReimbursementListResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A reimbursement as it appears in the list. Dates, merchant details, notes,
// category, and transaction id are available from the detail endpoint.
type SpendReimbursementListResponseReimbursement struct {
	// ID for the expense
	ID string `json:"id" api:"required" format:"uuid"`
	// The positive amount spent in the submitted currency.
	Amount      float64                        `json:"amount" api:"required"`
	Attachments []APIV1ReimbursementAttachment `json:"attachments" api:"required"`
	Currency    string                         `json:"currency" api:"required"`
	// Payout state, separate from review.
	//
	//   - `pending`: no payout in progress. The reimbursement isn't approved yet, or
	//     it's approved and waiting for a payout batch. A canceled payout, or one
	//     Mercury retries automatically, also returns to `pending`.
	//   - `initiated`: the payout is in progress.
	//   - `paid`: the payout completed.
	//   - `failed`: the payout failed. Mercury retries in a later batch once the
	//     submitter fixes their payout account, unless they've left the organization.
	//   - `paidOutsideMercury`: an approver recorded that the submitter was paid outside
	//     Mercury, so Mercury won't pay it. Its review status is `approved`.
	//
	// Any of "pending", "initiated", "paid", "failed", "paidOutsideMercury".
	PaymentStatus ReimbursementPaymentStatus `json:"paymentStatus" api:"required"`
	// Review state, separate from payment.
	//
	// - `submitted`: waiting for review.
	// - `updatesRequested`: a reviewer asked the submitter for changes.
	// - `approved`: a reviewer approved it. `paymentStatus` shows the payout.
	// - `denied`: a reviewer denied it.
	// - `canceled`: the submitter canceled it.
	//
	// Any of "submitted", "updatesRequested", "approved", "denied", "canceled".
	ReviewStatus ReimbursementReviewStatus `json:"reviewStatus" api:"required"`
	SubmittedAt  string                    `json:"submittedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// ID for the user
	SubmitterUserID string `json:"submitterUserId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Amount          respjson.Field
		Attachments     respjson.Field
		Currency        respjson.Field
		PaymentStatus   respjson.Field
		ReviewStatus    respjson.Field
		SubmittedAt     respjson.Field
		SubmitterUserID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpendReimbursementListResponseReimbursement) RawJSON() string { return r.JSON.raw }
func (r *SpendReimbursementListResponseReimbursement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendReimbursementNewParams struct {
	// Create a reimbursement (a regular, non-mileage expense).
	APIV1CreateReimbursementRequest APIV1CreateReimbursementRequestParam
	IdempotencyKey                  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r SpendReimbursementNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.APIV1CreateReimbursementRequest)
}
func (r *SpendReimbursementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendReimbursementUpdateParams struct {
	// Create a reimbursement (a regular, non-mileage expense).
	APIV1CreateReimbursementRequest APIV1CreateReimbursementRequestParam
	IdempotencyKey                  param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r SpendReimbursementUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.APIV1CreateReimbursementRequest)
}
func (r *SpendReimbursementUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendReimbursementListParams struct {
	// The ID of the reimbursement to end the page before (exclusive). When provided,
	// results will end just before this ID and work backwards. Use this for reverse
	// pagination or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the reimbursement to start the page after (exclusive). When provided,
	// results will begin with the reimbursement immediately following this ID. Use
	// this for standard forward pagination to get the next page of results. Cannot be
	// combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// The ID of the reimbursement to start the page at (inclusive). Cannot be combined
	// with start_after or end_before.
	StartAt param.Opt[string] `query:"start_at,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order SpendReimbursementListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpendReimbursementListParams]'s query parameters as
// `url.Values`.
func (r SpendReimbursementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type SpendReimbursementListParamsOrder string

const (
	SpendReimbursementListParamsOrderAsc  SpendReimbursementListParamsOrder = "asc"
	SpendReimbursementListParamsOrderDesc SpendReimbursementListParamsOrder = "desc"
)

type SpendReimbursementApproveParams struct {
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

type SpendReimbursementCancelParams struct {
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

type SpendReimbursementDenyParams struct {
	Reason         string            `json:"reason" api:"required"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r SpendReimbursementDenyParams) MarshalJSON() (data []byte, err error) {
	type shadow SpendReimbursementDenyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendReimbursementDenyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendReimbursementMarkPaidParams struct {
	// Recorded as a reviewer comment, e.g. how and when the submitter was paid.
	Note           param.Opt[string] `json:"note,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r SpendReimbursementMarkPaidParams) MarshalJSON() (data []byte, err error) {
	type shadow SpendReimbursementMarkPaidParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendReimbursementMarkPaidParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
