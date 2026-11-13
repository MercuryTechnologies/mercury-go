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
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/pagination"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// PaymentService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	options []option.RequestOption
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r PaymentService) {
	r = PaymentService{}
	r.options = opts
	return
}

// Send money from an account to a recipient. Creates a transaction that will be
// processed immediately or may require approval.
func (r *PaymentService) New(ctx context.Context, accountID string, body PaymentNewParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/%s/transactions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of send money approval requests for the authenticated
// organization. Supports filtering by account and status.
func (r *PaymentService) List(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) (res *pagination.CursorIDRequestSendMoney[SendMoneyApproval], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
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
func (r *PaymentService) ListAutoPaging(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) *pagination.CursorIDRequestSendMoneyAutoPager[SendMoneyApproval] {
	return pagination.NewCursorIDRequestSendMoneyAutoPager(r.List(ctx, query, opts...))
}

// Get send money approval request by ID
func (r *PaymentService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *SendMoneyApproval, err error) {
	opts = slices.Concat(r.options, opts)
	if requestID == "" {
		err = errors.New("missing required requestId parameter")
		return nil, err
	}
	path := fmt.Sprintf("request-send-money/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create a "request to send money" that will require approval based on your
// organization's approval policies.
func (r *PaymentService) Request(ctx context.Context, accountID string, body PaymentRequestParams, opts ...option.RequestOption) (res *SendMoneyApproval, err error) {
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/%s/request-send-money", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of transactions for a specific account. Supports
// filtering by date range, status, and search terms.
func (r *PaymentService) GetTransactions(ctx context.Context, accountID string, query PaymentGetTransactionsParams, opts ...option.RequestOption) (res *PaymentGetTransactionsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/%s/transactions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Transfer funds between two accounts within the same organization. Supports
// transfers between depository accounts (checking/savings), from a depository
// account to a treasury/investment account, and from a treasury/investment account
// to a depository account. Creates paired debit and credit transactions.
func (r *PaymentService) Transfer(ctx context.Context, body PaymentTransferParams, opts ...option.RequestOption) (res *PaymentTransferResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "transfer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PaymentApprovalReview struct {
	ReviewedAt string `json:"reviewedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// ID for the user
	ReviewerUserID string `json:"reviewerUserId" api:"required" format:"uuid"`
	// Any of "approved", "rejected".
	Status PaymentApprovalReviewStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReviewedAt     respjson.Field
		ReviewerUserID respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentApprovalReview) RawJSON() string { return r.JSON.raw }
func (r *PaymentApprovalReview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentApprovalReviewStatus string

const (
	PaymentApprovalReviewStatusApproved PaymentApprovalReviewStatus = "approved"
	PaymentApprovalReviewStatusRejected PaymentApprovalReviewStatus = "rejected"
)

// External API representation of SendMoneyPurpose. Only exposes the 'simple' field
// to decouple internal implementation from external API.
type PostTransactionSendMoneyPurposeParam struct {
	Simple PostTransactionSendMoneyPurposeSimpleParam `json:"simple,omitzero"`
	paramObj
}

func (r PostTransactionSendMoneyPurposeParam) MarshalJSON() (data []byte, err error) {
	type shadow PostTransactionSendMoneyPurposeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTransactionSendMoneyPurposeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Category is required.
type PostTransactionSendMoneyPurposeSimpleParam struct {
	// Payment category.
	//
	// Any of "employee", "landlord", "vendor", "contractor", "subsidiary",
	// "transferToMyExternalAccount", "familyMemberOrFriend", "forGoodsOrServices",
	// "angelInvestment", "savingsOrInvestments", "expenses", "travel", "other".
	Category string `json:"category,omitzero" api:"required"`
	// Additional information. Required for: Vendor (vendor name), Contractor
	// (contractor name), Other (payment description). Optional for Subsidiary
	// (subsidiary name). Not accepted for any other categories.
	AdditionalInfo param.Opt[string] `json:"additionalInfo,omitzero"`
	paramObj
}

func (r PostTransactionSendMoneyPurposeSimpleParam) MarshalJSON() (data []byte, err error) {
	type shadow PostTransactionSendMoneyPurposeSimpleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PostTransactionSendMoneyPurposeSimpleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PostTransactionSendMoneyPurposeSimpleParam](
		"category", "employee", "landlord", "vendor", "contractor", "subsidiary", "transferToMyExternalAccount", "familyMemberOrFriend", "forGoodsOrServices", "angelInvestment", "savingsOrInvestments", "expenses", "travel", "other",
	)
}

type ReviewRequestStatus string

const (
	ReviewRequestStatusPendingApproval ReviewRequestStatus = "pendingApproval"
	ReviewRequestStatusApproved        ReviewRequestStatus = "approved"
	ReviewRequestStatusRejected        ReviewRequestStatus = "rejected"
	ReviewRequestStatusCancelled       ReviewRequestStatus = "cancelled"
)

// A pending or completed approval request for a Mercury payment.
type SendMoneyApproval struct {
	// ID for a Mercury account.
	AccountID string `json:"accountId" api:"required" format:"uuid"`
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount" api:"required"`
	// Time at which the payment request was created.
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Any of "ach", "check", "domesticWire", "internationalWire", "realTimePayment".
	PaymentMethod SendMoneyPaymentMethod `json:"paymentMethod" api:"required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId" api:"required" format:"uuid"`
	// ID for the user
	RequestedByUserID string `json:"requestedByUserId" api:"required" format:"uuid"`
	RequestID         string `json:"requestId" api:"required"`
	// Approval decisions recorded against this request, ordered from oldest to most
	// recent.
	Reviews []PaymentApprovalReview `json:"reviews" api:"required"`
	// Any of "pendingApproval", "approved", "rejected", "cancelled".
	Status ReviewRequestStatus `json:"status" api:"required"`
	Memo   string              `json:"memo" api:"nullable"`
	// Total number of approvals required for this payment to be sent. Omitted for
	// older requests where the requirement is not available.
	NumberOfApproversRequired int64 `json:"numberOfApproversRequired" api:"nullable"`
	// True when the approval requirement recorded for this payment names the requester
	// as its approver. Reflects the recorded requirement, not the requester's current
	// standing, so submitting the approval can still be refused. The value persists
	// after the request settles and says only whether the requirement names the
	// requester; whether an approval is still outstanding is carried by
	// numberOfApproversRequired and status. Omitted for older requests where the
	// requirement is not available.
	RequesterMayApprove bool `json:"requesterMayApprove" api:"nullable"`
	// Date on which the payment is scheduled to be sent once fully approved. Null when
	// the payment will be sent as soon as approvals are complete.
	ScheduledSendDate time.Time `json:"scheduledSendDate" api:"nullable" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID                 respjson.Field
		Amount                    respjson.Field
		CreatedAt                 respjson.Field
		PaymentMethod             respjson.Field
		RecipientID               respjson.Field
		RequestedByUserID         respjson.Field
		RequestID                 respjson.Field
		Reviews                   respjson.Field
		Status                    respjson.Field
		Memo                      respjson.Field
		NumberOfApproversRequired respjson.Field
		RequesterMayApprove       respjson.Field
		ScheduledSendDate         respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SendMoneyApproval) RawJSON() string { return r.JSON.raw }
func (r *SendMoneyApproval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMoneyPaymentMethod string

const (
	SendMoneyPaymentMethodACH               SendMoneyPaymentMethod = "ach"
	SendMoneyPaymentMethodCheck             SendMoneyPaymentMethod = "check"
	SendMoneyPaymentMethodDomesticWire      SendMoneyPaymentMethod = "domesticWire"
	SendMoneyPaymentMethodInternationalWire SendMoneyPaymentMethod = "internationalWire"
	SendMoneyPaymentMethodRealTimePayment   SendMoneyPaymentMethod = "realTimePayment"
)

type PaymentGetTransactionsResponse struct {
	Total        int64         `json:"total" api:"required"`
	Transactions []Transaction `json:"transactions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Total        respjson.Field
		Transactions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentGetTransactionsResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentGetTransactionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for POST /api/v1/transfer endpoint. Returns both the credit and debit
// transactions for the transfer (depository, treasury, or investment).
type PaymentTransferResponse struct {
	CreditTransaction Transaction `json:"creditTransaction" api:"required"`
	DebitTransaction  Transaction `json:"debitTransaction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTransaction respjson.Field
		DebitTransaction  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PaymentTransferResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentTransferResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentNewParams struct {
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount" api:"required"`
	// Unique string identifying the transaction
	IdempotencyKey string `json:"idempotencyKey" api:"required"`
	// If domesticWire is used, then the purpose field is required.
	//
	// Any of "ach", "check", "domesticWire", "realTimePayment".
	PaymentMethod PaymentNewParamsPaymentMethod `json:"paymentMethod,omitzero" api:"required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId" api:"required" format:"uuid"`
	// Set to true to skip the duplicate payment check. By default, payments are
	// blocked if the same account sent the same amount to the same recipient with the
	// same payment method in the last 24 hours, even when the idempotency key is
	// different. Use this field only when you intentionally send a repeat payment.
	DuplicateCheckOverride param.Opt[bool] `json:"duplicateCheckOverride,omitzero"`
	// Optional external memo
	ExternalMemo param.Opt[string] `json:"externalMemo,omitzero"`
	// Optional note
	Note param.Opt[string] `json:"note,omitzero"`
	// External API representation of SendMoneyPurpose. Only exposes the 'simple' field
	// to decouple internal implementation from external API.
	Purpose PostTransactionSendMoneyPurposeParam `json:"purpose,omitzero"`
	paramObj
}

func (r PaymentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// If domesticWire is used, then the purpose field is required.
type PaymentNewParamsPaymentMethod string

const (
	PaymentNewParamsPaymentMethodACH             PaymentNewParamsPaymentMethod = "ach"
	PaymentNewParamsPaymentMethodCheck           PaymentNewParamsPaymentMethod = "check"
	PaymentNewParamsPaymentMethodDomesticWire    PaymentNewParamsPaymentMethod = "domesticWire"
	PaymentNewParamsPaymentMethodRealTimePayment PaymentNewParamsPaymentMethod = "realTimePayment"
)

type PaymentListParams struct {
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
	Status PaymentListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaymentListParams]'s query parameters as `url.Values`.
func (r PaymentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PaymentListParamsStatus string

const (
	PaymentListParamsStatusPendingApproval PaymentListParamsStatus = "pendingApproval"
	PaymentListParamsStatusApproved        PaymentListParamsStatus = "approved"
	PaymentListParamsStatusRejected        PaymentListParamsStatus = "rejected"
	PaymentListParamsStatusCancelled       PaymentListParamsStatus = "cancelled"
)

type PaymentRequestParams struct {
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount" api:"required"`
	// Unique string identifying the transaction
	IdempotencyKey string `json:"idempotencyKey" api:"required"`
	// Payment method to use. If domesticWire or internationalWire is used, then the
	// purpose field is required.
	//
	// Any of "ach", "check", "domesticWire", "internationalWire", "realTimePayment".
	PaymentMethod SendMoneyPaymentMethod `json:"paymentMethod,omitzero" api:"required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId" api:"required" format:"uuid"`
	// Set to true to skip the duplicate payment check. By default, payments are
	// blocked if the same account sent the same amount to the same recipient with the
	// same payment method in the last 24 hours, even when the idempotency key is
	// different. Use this field only when you intentionally send a repeat payment. The
	// check occurs when you create the request. When an approver approves the request,
	// the check does not occur again.
	DuplicateCheckOverride param.Opt[bool] `json:"duplicateCheckOverride,omitzero"`
	// Optional external memo
	ExternalMemo param.Opt[string] `json:"externalMemo,omitzero"`
	// Optional note
	Note param.Opt[string] `json:"note,omitzero"`
	// Who pays intermediary bank fees on a USD international wire. Pass 'ours' so the
	// recipient receives the full amount (Mercury charges a $15 fee), or 'shared' to
	// split those fees. If omitted on a USD international wire, this defaults to
	// 'ours'. Do not include this field for other payment methods or accounts that are
	// not charged international-wire fees.
	//
	// Any of "ours", "shared".
	ChargeType PaymentRequestParamsChargeType `json:"chargeType,omitzero"`
	// External API representation of SendMoneyPurpose. Only exposes the 'simple' field
	// to decouple internal implementation from external API.
	Purpose PostTransactionSendMoneyPurposeParam `json:"purpose,omitzero"`
	paramObj
}

func (r PaymentRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Who pays intermediary bank fees on a USD international wire. Pass 'ours' so the
// recipient receives the full amount (Mercury charges a $15 fee), or 'shared' to
// split those fees. If omitted on a USD international wire, this defaults to
// 'ours'. Do not include this field for other payment methods or accounts that are
// not charged international-wire fees.
type PaymentRequestParamsChargeType string

const (
	PaymentRequestParamsChargeTypeOurs   PaymentRequestParamsChargeType = "ours"
	PaymentRequestParamsChargeTypeShared PaymentRequestParamsChargeType = "shared"
)

type PaymentGetTransactionsParams struct {
	// UUID of a custom category. Can be returned from /categories endpoint.
	CategoryID param.Opt[string] `query:"categoryId,omitzero" format:"uuid" json:"-"`
	// Latest date to filter transactions. If not provided, defaults to the current
	// date. Format: YYYY-MM-DD or ISO 8601 string
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Name of mercuryCategory you want to filter on. Merchant Type in the UI.
	MercuryCategory param.Opt[string] `query:"mercuryCategory,omitzero" json:"-"`
	// Number of results to skip for pagination
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// ID returned from /account/:id/request-send-money
	RequestID param.Opt[string] `query:"requestId,omitzero" format:"uuid" json:"-"`
	// Search term to filter transactions by description, note, recipient memo, or
	// counterparty name
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Earliest date to filter transactions. If not provided, defaults to 30 days
	// before the current date. Format: YYYY-MM-DD or ISO 8601 string
	Start param.Opt[string] `query:"start,omitzero" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'
	//
	// Any of "asc", "desc".
	Order PaymentGetTransactionsParamsOrder `query:"order,omitzero" json:"-"`
	// Any of "pending", "sent", "cancelled", "failed", "reversed", "blocked".
	Status PaymentGetTransactionsParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PaymentGetTransactionsParams]'s query parameters as
// `url.Values`.
func (r PaymentGetTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'
type PaymentGetTransactionsParamsOrder string

const (
	PaymentGetTransactionsParamsOrderAsc  PaymentGetTransactionsParamsOrder = "asc"
	PaymentGetTransactionsParamsOrderDesc PaymentGetTransactionsParamsOrder = "desc"
)

type PaymentGetTransactionsParamsStatus string

const (
	PaymentGetTransactionsParamsStatusPending   PaymentGetTransactionsParamsStatus = "pending"
	PaymentGetTransactionsParamsStatusSent      PaymentGetTransactionsParamsStatus = "sent"
	PaymentGetTransactionsParamsStatusCancelled PaymentGetTransactionsParamsStatus = "cancelled"
	PaymentGetTransactionsParamsStatusFailed    PaymentGetTransactionsParamsStatus = "failed"
	PaymentGetTransactionsParamsStatusReversed  PaymentGetTransactionsParamsStatus = "reversed"
	PaymentGetTransactionsParamsStatusBlocked   PaymentGetTransactionsParamsStatus = "blocked"
)

type PaymentTransferParams struct {
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

func (r PaymentTransferParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentTransferParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentTransferParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
