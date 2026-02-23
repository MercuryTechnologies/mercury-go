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
	"github.com/stainless-sdks/mercury-go/packages/param"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
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
func (r *AccountTransactionService) List(ctx context.Context, accountID string, query AccountTransactionListParams, opts ...option.RequestOption) (res *AccountTransactionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("account/%s/transactions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
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

type AccountTransactionListResponse struct {
	Total        int64         `json:"total,required"`
	Transactions []Transaction `json:"transactions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Total        respjson.Field
		Transactions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountTransactionListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountTransactionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// Payment method to use, currently only supports "ach".
	//
	// Any of "ach".
	PaymentMethod AccountTransactionSendParamsPaymentMethod `json:"paymentMethod,omitzero,required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId,required" format:"uuid"`
	// Optional external memo
	ExternalMemo param.Opt[string] `json:"externalMemo,omitzero"`
	// Optional note
	Note param.Opt[string] `json:"note,omitzero"`
	paramObj
}

func (r AccountTransactionSendParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountTransactionSendParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountTransactionSendParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payment method to use, currently only supports "ach".
type AccountTransactionSendParamsPaymentMethod string

const (
	AccountTransactionSendParamsPaymentMethodACH AccountTransactionSendParamsPaymentMethod = "ach"
)
