// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/mercury-go/internal/apijson"
	"github.com/stainless-sdks/mercury-go/internal/apiquery"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/pagination"
	"github.com/stainless-sdks/mercury-go/packages/param"
)

// AccountTransactionService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountTransactionService] method instead.
type AccountTransactionService struct {
	Options []option.RequestOption
}

// NewAccountTransactionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountTransactionService(opts ...option.RequestOption) (r AccountTransactionService) {
	r = AccountTransactionService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of transactions for a specific account. Supports
// filtering by date range, status, and search terms.
func (r *AccountTransactionService) List(ctx context.Context, accountID string, query AccountTransactionListParams, opts ...option.RequestOption) (res *pagination.OffsetAccountTransactions[Transaction], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8"), option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("account/%s/transactions", accountID)
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

// Retrieve a paginated list of transactions for a specific account. Supports
// filtering by date range, status, and search terms.
func (r *AccountTransactionService) ListAutoPaging(ctx context.Context, accountID string, query AccountTransactionListParams, opts ...option.RequestOption) *pagination.OffsetAccountTransactionsAutoPager[Transaction] {
	return pagination.NewOffsetAccountTransactionsAutoPager(r.List(ctx, accountID, query, opts...))
}

// Send money from an account to a recipient. Creates a transaction that will be
// processed immediately or may require approval.
func (r *AccountTransactionService) Send(ctx context.Context, accountID string, body AccountTransactionSendParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("account/%s/transactions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountTransactionListParams struct {
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
	// Search term to filter transactions by description or counterparty name
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Earliest date to filter transactions. If not provided, defaults to 30 days
	// before the current date. Format: YYYY-MM-DD or ISO 8601 string
	Start param.Opt[string] `query:"start,omitzero" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'
	//
	// Any of "asc", "desc".
	Order AccountTransactionListParamsOrder `query:"order,omitzero" json:"-"`
	// Any of "pending", "sent", "cancelled", "failed", "reversed", "blocked".
	Status AccountTransactionListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountTransactionListParams]'s query parameters as
// `url.Values`.
func (r AccountTransactionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'
type AccountTransactionListParamsOrder string

const (
	AccountTransactionListParamsOrderAsc  AccountTransactionListParamsOrder = "asc"
	AccountTransactionListParamsOrderDesc AccountTransactionListParamsOrder = "desc"
)

type AccountTransactionListParamsStatus string

const (
	AccountTransactionListParamsStatusPending   AccountTransactionListParamsStatus = "pending"
	AccountTransactionListParamsStatusSent      AccountTransactionListParamsStatus = "sent"
	AccountTransactionListParamsStatusCancelled AccountTransactionListParamsStatus = "cancelled"
	AccountTransactionListParamsStatusFailed    AccountTransactionListParamsStatus = "failed"
	AccountTransactionListParamsStatusReversed  AccountTransactionListParamsStatus = "reversed"
	AccountTransactionListParamsStatusBlocked   AccountTransactionListParamsStatus = "blocked"
)

type AccountTransactionSendParams struct {
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount,required"`
	// Unique string identifying the transaction
	IdempotencyKey string `json:"idempotencyKey,required"`
	// If domesticWire is used, then the purpose field is required.
	//
	// Any of "ach", "domesticWire".
	PaymentMethod AccountTransactionSendParamsPaymentMethod `json:"paymentMethod,omitzero,required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId,required" format:"uuid"`
	// Optional external memo
	ExternalMemo param.Opt[string] `json:"externalMemo,omitzero"`
	// Optional note
	Note param.Opt[string] `json:"note,omitzero"`
	// External API representation of SendMoneyPurpose. Only exposes the 'simple' field
	// to decouple internal implementation from external API.
	Purpose AccountTransactionSendParamsPurpose `json:"purpose,omitzero"`
	paramObj
}

func (r AccountTransactionSendParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountTransactionSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountTransactionSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// If domesticWire is used, then the purpose field is required.
type AccountTransactionSendParamsPaymentMethod string

const (
	AccountTransactionSendParamsPaymentMethodACH          AccountTransactionSendParamsPaymentMethod = "ach"
	AccountTransactionSendParamsPaymentMethodDomesticWire AccountTransactionSendParamsPaymentMethod = "domesticWire"
)

// External API representation of SendMoneyPurpose. Only exposes the 'simple' field
// to decouple internal implementation from external API.
type AccountTransactionSendParamsPurpose struct {
	Simple AccountTransactionSendParamsPurposeSimple `json:"simple,omitzero"`
	paramObj
}

func (r AccountTransactionSendParamsPurpose) MarshalJSON() (data []byte, err error) {
	type shadow AccountTransactionSendParamsPurpose
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountTransactionSendParamsPurpose) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Category is required.
type AccountTransactionSendParamsPurposeSimple struct {
	// Payment category.
	//
	// Any of "Employee", "Landlord", "Vendor", "Contractor", "Subsidiary",
	// "TransferToMyExternalAccount", "FamilyMemberOrFriend", "ForGoodsOrServices",
	// "AngelInvestment", "SavingsOrInvestments", "Expenses", "Travel", "Other".
	Category string `json:"category,omitzero,required"`
	// Additional information. Required for: Vendor (vendor name), Contractor
	// (contractor name), Other (payment description). Optional for Subsidiary
	// (subsidiary name). Not accepted for any other categories.
	AdditionalInfo param.Opt[string] `json:"additionalInfo,omitzero"`
	paramObj
}

func (r AccountTransactionSendParamsPurposeSimple) MarshalJSON() (data []byte, err error) {
	type shadow AccountTransactionSendParamsPurposeSimple
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountTransactionSendParamsPurposeSimple) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AccountTransactionSendParamsPurposeSimple](
		"category", "Employee", "Landlord", "Vendor", "Contractor", "Subsidiary", "TransferToMyExternalAccount", "FamilyMemberOrFriend", "ForGoodsOrServices", "AngelInvestment", "SavingsOrInvestments", "Expenses", "Travel", "Other",
	)
}
