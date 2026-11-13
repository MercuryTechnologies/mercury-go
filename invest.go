// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"encoding/json"
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

// Retrieve Invest accounts and their activity
//
// InvestService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestService] method instead.
type InvestService struct {
	options []option.RequestOption
}

// NewInvestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInvestService(opts ...option.RequestOption) (r InvestService) {
	r = InvestService{}
	r.options = opts
	return
}

// Retrieve a paginated list of all Invest accounts associated with the
// organization.
func (r *InvestService) List(ctx context.Context, query InvestListParams, opts ...option.RequestOption) (res *pagination.CursorIDAccounts[InvestListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "invest"
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

// Retrieve a paginated list of all Invest accounts associated with the
// organization.
func (r *InvestService) ListAutoPaging(ctx context.Context, query InvestListParams, opts ...option.RequestOption) *pagination.CursorIDAccountsAutoPager[InvestListResponse] {
	return pagination.NewCursorIDAccountsAutoPager(r.List(ctx, query, opts...))
}

// Retrieve the activity history (no pagination) for a specific Invest account:
// transfers, trades, fees, dividends, interest payments, credits and portfolio
// strategy changes.
func (r *InvestService) GetActivity(ctx context.Context, investID string, opts ...option.RequestOption) (res *InvestGetActivityResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if investID == "" {
		err = errors.New("missing required investId parameter")
		return nil, err
	}
	path := fmt.Sprintf("invest/%s/activity", investID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type InvestActivityStatus string

const (
	InvestActivityStatusPending  InvestActivityStatus = "pending"
	InvestActivityStatusCanceled InvestActivityStatus = "canceled"
	InvestActivityStatusFailed   InvestActivityStatus = "failed"
	InvestActivityStatusComplete InvestActivityStatus = "complete"
)

// An Invest account
type InvestListResponse struct {
	// ID for a Mercury account.
	ID string `json:"id" api:"required" format:"uuid"`
	// Available balance of the account
	AvailableBalance float64 `json:"availableBalance" api:"required"`
	// When the account was created
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Current balance of the account
	CurrentBalance float64 `json:"currentBalance" api:"required"`
	// Status of the account
	//
	// Any of "active", "deleted", "pending", "archived".
	Status AccountStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AvailableBalance respjson.Field
		CreatedAt        respjson.Field
		CurrentBalance   respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestListResponse) RawJSON() string { return r.JSON.raw }
func (r *InvestListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Activity for a single Invest account.
type InvestGetActivityResponse struct {
	// Activity items.
	Activity []InvestGetActivityResponseActivity `json:"activity" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Activity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponse) RawJSON() string { return r.JSON.raw }
func (r *InvestGetActivityResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single item of Invest account activity.
type InvestGetActivityResponseActivity struct {
	// The day the activity occurred.
	Day time.Time `json:"day" api:"required" format:"date"`
	// The activity itself.
	Details InvestGetActivityResponseActivityDetailsUnion `json:"details" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day         respjson.Field
		Details     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivity) RawJSON() string { return r.JSON.raw }
func (r *InvestGetActivityResponseActivity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvestGetActivityResponseActivityDetailsUnion contains all possible properties
// and values from
// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransfer],
// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTrade],
// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadFee],
// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadDividend],
// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadInterestPayment],
// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadCredit],
// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChange].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InvestGetActivityResponseActivityDetailsUnion struct {
	// This field is a union of
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents],
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTradeContents],
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadFeeContents],
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadDividendContents],
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadInterestPaymentContents],
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadCreditContents],
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChangeContents]
	Contents InvestGetActivityResponseActivityDetailsUnionContents `json:"contents"`
	Kind     string                                                `json:"kind"`
	JSON     struct {
		Contents respjson.Field
		Kind     respjson.Field
		raw      string
	} `json:"-"`
}

func (u InvestGetActivityResponseActivityDetailsUnion) AsInvestActivityPayloadTransfer() (v InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransfer) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvestGetActivityResponseActivityDetailsUnion) AsInvestActivityPayloadTrade() (v InvestGetActivityResponseActivityDetailsInvestActivityPayloadTrade) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvestGetActivityResponseActivityDetailsUnion) AsInvestActivityPayloadFee() (v InvestGetActivityResponseActivityDetailsInvestActivityPayloadFee) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvestGetActivityResponseActivityDetailsUnion) AsInvestActivityPayloadDividend() (v InvestGetActivityResponseActivityDetailsInvestActivityPayloadDividend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvestGetActivityResponseActivityDetailsUnion) AsInvestActivityPayloadInterestPayment() (v InvestGetActivityResponseActivityDetailsInvestActivityPayloadInterestPayment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvestGetActivityResponseActivityDetailsUnion) AsInvestActivityPayloadCredit() (v InvestGetActivityResponseActivityDetailsInvestActivityPayloadCredit) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InvestGetActivityResponseActivityDetailsUnion) AsInvestActivityPayloadPortfolioChange() (v InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChange) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InvestGetActivityResponseActivityDetailsUnion) RawJSON() string { return u.JSON.raw }

func (r *InvestGetActivityResponseActivityDetailsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InvestGetActivityResponseActivityDetailsUnionContents is an implicit subunion of
// [InvestGetActivityResponseActivityDetailsUnion].
// InvestGetActivityResponseActivityDetailsUnionContents provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [InvestGetActivityResponseActivityDetailsUnion].
type InvestGetActivityResponseActivityDetailsUnionContents struct {
	Amount float64 `json:"amount"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents].
	CounterpartyID string `json:"counterpartyId"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents].
	CounterpartyName string `json:"counterpartyName"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents].
	IsAutomated bool `json:"isAutomated"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents].
	Kind string `json:"kind"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents].
	Status InvestActivityStatus `json:"status"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents].
	TransactionID string `json:"transactionId"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents].
	PostedAt string `json:"postedAt"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTradeContents].
	Action string `json:"action"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTradeContents].
	Price float64 `json:"price"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTradeContents].
	Quantity string `json:"quantity"`
	Security string `json:"security"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadTradeContents].
	Total float64 `json:"total"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChangeContents].
	From string `json:"from"`
	// This field is from variant
	// [InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChangeContents].
	To   string `json:"to"`
	JSON struct {
		Amount           respjson.Field
		CounterpartyID   respjson.Field
		CounterpartyName respjson.Field
		IsAutomated      respjson.Field
		Kind             respjson.Field
		Status           respjson.Field
		TransactionID    respjson.Field
		PostedAt         respjson.Field
		Action           respjson.Field
		Price            respjson.Field
		Quantity         respjson.Field
		Security         respjson.Field
		Total            respjson.Field
		From             respjson.Field
		To               respjson.Field
		raw              string
	} `json:"-"`
}

func (r *InvestGetActivityResponseActivityDetailsUnionContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransfer struct {
	// A movement of funds between a depository account and the Invest account.
	Contents InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents `json:"contents" api:"required"`
	// Any of "transfer".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransfer) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransfer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A movement of funds between a depository account and the Invest account.
type InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents struct {
	// A dollar amount
	Amount float64 `json:"amount" api:"required"`
	// ID for a Mercury account.
	CounterpartyID string `json:"counterpartyId" api:"required" format:"uuid"`
	// Name of the account on the other side of the transfer.
	CounterpartyName string `json:"counterpartyName" api:"required"`
	// Whether the transfer resulted from an auto-transfer rule.
	IsAutomated bool `json:"isAutomated" api:"required"`
	// Whether funds moved into or out of the Invest account.
	//
	// Any of "investment", "liquidation".
	Kind string `json:"kind" api:"required"`
	// Status of the transfer.
	//
	// Any of "pending", "canceled", "failed", "complete".
	Status InvestActivityStatus `json:"status" api:"required"`
	// ID for this transaction
	TransactionID string `json:"transactionId" api:"required" format:"uuid"`
	// When the transfer posted, if it has posted.
	PostedAt string `json:"postedAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount           respjson.Field
		CounterpartyID   respjson.Field
		CounterpartyName respjson.Field
		IsAutomated      respjson.Field
		Kind             respjson.Field
		Status           respjson.Field
		TransactionID    respjson.Field
		PostedAt         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadTransferContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestGetActivityResponseActivityDetailsInvestActivityPayloadTrade struct {
	// A purchase or sale of a security.
	Contents InvestGetActivityResponseActivityDetailsInvestActivityPayloadTradeContents `json:"contents" api:"required"`
	// Any of "trade".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadTrade) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadTrade) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A purchase or sale of a security.
type InvestGetActivityResponseActivityDetailsInvestActivityPayloadTradeContents struct {
	// Whether the trade was a buy or a sell.
	//
	// Any of "buy", "sell".
	Action string `json:"action" api:"required"`
	// A dollar amount
	Price float64 `json:"price" api:"required"`
	// Quantity of the security exchanged, to five decimal places.
	Quantity string `json:"quantity" api:"required"`
	// Ticker symbol of the security traded.
	Security string `json:"security" api:"required"`
	// Status of the trade.
	//
	// Any of "pending", "canceled", "failed", "complete".
	Status InvestActivityStatus `json:"status" api:"required"`
	// A dollar amount
	Total float64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Price       respjson.Field
		Quantity    respjson.Field
		Security    respjson.Field
		Status      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadTradeContents) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadTradeContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestGetActivityResponseActivityDetailsInvestActivityPayloadFee struct {
	// A fee charged against the Invest account.
	Contents InvestGetActivityResponseActivityDetailsInvestActivityPayloadFeeContents `json:"contents" api:"required"`
	// Any of "fee".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadFee) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fee charged against the Invest account.
type InvestGetActivityResponseActivityDetailsInvestActivityPayloadFeeContents struct {
	// A dollar amount
	Amount float64 `json:"amount" api:"required"`
	// Status of the fee.
	//
	// Any of "pending", "canceled", "failed", "complete".
	Status InvestActivityStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadFeeContents) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadFeeContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestGetActivityResponseActivityDetailsInvestActivityPayloadDividend struct {
	// A dividend paid by a held security.
	Contents InvestGetActivityResponseActivityDetailsInvestActivityPayloadDividendContents `json:"contents" api:"required"`
	// Any of "dividend".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadDividend) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadDividend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dividend paid by a held security.
type InvestGetActivityResponseActivityDetailsInvestActivityPayloadDividendContents struct {
	// A dollar amount
	Amount float64 `json:"amount" api:"required"`
	// Ticker symbol of the security that paid the dividend.
	Security string `json:"security" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Security    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadDividendContents) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadDividendContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestGetActivityResponseActivityDetailsInvestActivityPayloadInterestPayment struct {
	// An interest payment credited to the Invest account.
	Contents InvestGetActivityResponseActivityDetailsInvestActivityPayloadInterestPaymentContents `json:"contents" api:"required"`
	// Any of "interestPayment".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadInterestPayment) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadInterestPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An interest payment credited to the Invest account.
type InvestGetActivityResponseActivityDetailsInvestActivityPayloadInterestPaymentContents struct {
	// A dollar amount
	Amount float64 `json:"amount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadInterestPaymentContents) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadInterestPaymentContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestGetActivityResponseActivityDetailsInvestActivityPayloadCredit struct {
	// A credit applied to the Invest account.
	Contents InvestGetActivityResponseActivityDetailsInvestActivityPayloadCreditContents `json:"contents" api:"required"`
	// Any of "credit".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadCredit) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadCredit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A credit applied to the Invest account.
type InvestGetActivityResponseActivityDetailsInvestActivityPayloadCreditContents struct {
	// A dollar amount
	Amount float64 `json:"amount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadCreditContents) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadCreditContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChange struct {
	// A change of portfolio strategy on the Invest account.
	Contents InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChangeContents `json:"contents" api:"required"`
	// Any of "portfolioChange".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChange) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A change of portfolio strategy on the Invest account.
type InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChangeContents struct {
	// Name of the previous portfolio strategy.
	From string `json:"from" api:"required"`
	// Status of the portfolio change.
	//
	// Any of "pending", "canceled", "failed", "complete".
	Status InvestActivityStatus `json:"status" api:"required"`
	// Name of the new portfolio strategy.
	To string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Status      respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChangeContents) RawJSON() string {
	return r.JSON.raw
}
func (r *InvestGetActivityResponseActivityDetailsInvestActivityPayloadPortfolioChangeContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvestListParams struct {
	// The ID of the account to end the page before (exclusive). When provided, results
	// will end just before this ID and work backwards. Use this for reverse pagination
	// or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the account to start the page after (exclusive). When provided,
	// results will begin with the account immediately following this ID. Use this for
	// standard forward pagination to get the next page of results. Cannot be combined
	// with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order InvestListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvestListParams]'s query parameters as `url.Values`.
func (r InvestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type InvestListParamsOrder string

const (
	InvestListParamsOrderAsc  InvestListParamsOrder = "asc"
	InvestListParamsOrderDesc InvestListParamsOrder = "desc"
)
