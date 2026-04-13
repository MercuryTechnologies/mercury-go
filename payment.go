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
	Options []option.RequestOption
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r PaymentService) {
	r = PaymentService{}
	r.Options = opts
	return
}

// Send money from an account to a recipient. Creates a transaction that will be
// processed immediately or may require approval.
func (r *PaymentService) New(ctx context.Context, accountID string, body PaymentNewParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *PaymentService) ListAutoPaging(ctx context.Context, query PaymentListParams, opts ...option.RequestOption) *pagination.CursorIDRequestSendMoneyAutoPager[SendMoneyApproval] {
	return pagination.NewCursorIDRequestSendMoneyAutoPager(r.List(ctx, query, opts...))
}

// Get send money approval request by ID
func (r *PaymentService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *SendMoneyApproval, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/%s/request-send-money", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transfer funds between two accounts within the same organization. Supports
// transfers between depository accounts (checking/savings), from a depository
// account to a treasury/investment account, and from a treasury/investment account
// to a depository account. Creates paired debit and credit transactions.
func (r *PaymentService) Transfer(ctx context.Context, body PaymentTransferParams, opts ...option.RequestOption) (res *PaymentTransferResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "transfer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Extremely close to the internal type, but strips out potentially unwanted fields
type SendMoneyApproval struct {
	// ID for a Mercury account.
	AccountID string `json:"accountId" api:"required" format:"uuid"`
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount" api:"required"`
	// Any of "ach", "check", "domesticWire", "internationalWire".
	PaymentMethod SendMoneyPaymentMethod `json:"paymentMethod" api:"required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId" api:"required" format:"uuid"`
	RequestID   string `json:"requestId" api:"required"`
	// Any of "pendingApproval", "approved", "rejected", "cancelled".
	Status SendMoneyApprovalStatus `json:"status" api:"required"`
	Memo   string                  `json:"memo" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID     respjson.Field
		Amount        respjson.Field
		PaymentMethod respjson.Field
		RecipientID   respjson.Field
		RequestID     respjson.Field
		Status        respjson.Field
		Memo          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SendMoneyApproval) RawJSON() string { return r.JSON.raw }
func (r *SendMoneyApproval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMoneyApprovalStatus string

const (
	SendMoneyApprovalStatusPendingApproval SendMoneyApprovalStatus = "pendingApproval"
	SendMoneyApprovalStatusApproved        SendMoneyApprovalStatus = "approved"
	SendMoneyApprovalStatusRejected        SendMoneyApprovalStatus = "rejected"
	SendMoneyApprovalStatusCancelled       SendMoneyApprovalStatus = "cancelled"
)

type SendMoneyPaymentMethod string

const (
	SendMoneyPaymentMethodACH               SendMoneyPaymentMethod = "ach"
	SendMoneyPaymentMethodCheck             SendMoneyPaymentMethod = "check"
	SendMoneyPaymentMethodDomesticWire      SendMoneyPaymentMethod = "domesticWire"
	SendMoneyPaymentMethodInternationalWire SendMoneyPaymentMethod = "internationalWire"
)

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
	// Any of "ach", "check", "domesticWire".
	PaymentMethod PaymentNewParamsPaymentMethod `json:"paymentMethod,omitzero" api:"required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId" api:"required" format:"uuid"`
	// Optional external memo
	ExternalMemo param.Opt[string] `json:"externalMemo,omitzero"`
	// Optional note
	Note param.Opt[string] `json:"note,omitzero"`
	// External API representation of SendMoneyPurpose. Only exposes the 'simple' field
	// to decouple internal implementation from external API.
	Purpose PaymentNewParamsPurpose `json:"purpose,omitzero"`
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
	PaymentNewParamsPaymentMethodACH          PaymentNewParamsPaymentMethod = "ach"
	PaymentNewParamsPaymentMethodCheck        PaymentNewParamsPaymentMethod = "check"
	PaymentNewParamsPaymentMethodDomesticWire PaymentNewParamsPaymentMethod = "domesticWire"
)

// External API representation of SendMoneyPurpose. Only exposes the 'simple' field
// to decouple internal implementation from external API.
type PaymentNewParamsPurpose struct {
	Simple PaymentNewParamsPurposeSimple `json:"simple,omitzero"`
	paramObj
}

func (r PaymentNewParamsPurpose) MarshalJSON() (data []byte, err error) {
	type shadow PaymentNewParamsPurpose
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentNewParamsPurpose) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Category is required.
type PaymentNewParamsPurposeSimple struct {
	// Payment category.
	//
	// Any of "Employee", "Landlord", "Vendor", "Contractor", "Subsidiary",
	// "TransferToMyExternalAccount", "FamilyMemberOrFriend", "ForGoodsOrServices",
	// "AngelInvestment", "SavingsOrInvestments", "Expenses", "Travel", "Other".
	Category string `json:"category,omitzero" api:"required"`
	// Additional information. Required for: Vendor (vendor name), Contractor
	// (contractor name), Other (payment description). Optional for Subsidiary
	// (subsidiary name). Not accepted for any other categories.
	AdditionalInfo param.Opt[string] `json:"additionalInfo,omitzero"`
	paramObj
}

func (r PaymentNewParamsPurposeSimple) MarshalJSON() (data []byte, err error) {
	type shadow PaymentNewParamsPurposeSimple
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentNewParamsPurposeSimple) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PaymentNewParamsPurposeSimple](
		"category", "Employee", "Landlord", "Vendor", "Contractor", "Subsidiary", "TransferToMyExternalAccount", "FamilyMemberOrFriend", "ForGoodsOrServices", "AngelInvestment", "SavingsOrInvestments", "Expenses", "Travel", "Other",
	)
}

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
	// Payment method to use.
	//
	// Any of "ach", "check", "domesticWire", "internationalWire".
	PaymentMethod SendMoneyPaymentMethod `json:"paymentMethod,omitzero" api:"required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId" api:"required" format:"uuid"`
	// Optional external memo
	ExternalMemo param.Opt[string] `json:"externalMemo,omitzero"`
	// Optional note
	Note param.Opt[string] `json:"note,omitzero"`
	paramObj
}

func (r PaymentRequestParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentRequestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentRequestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
