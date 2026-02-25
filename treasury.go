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

	"github.com/stainless-sdks/mercury-go/internal/apijson"
	"github.com/stainless-sdks/mercury-go/internal/apiquery"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/pagination"
	"github.com/stainless-sdks/mercury-go/packages/param"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// TreasuryService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTreasuryService] method instead.
type TreasuryService struct {
	Options []option.RequestOption
}

// NewTreasuryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTreasuryService(opts ...option.RequestOption) (r TreasuryService) {
	r = TreasuryService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of all treasury accounts associated with the
// authenticated organization. Use cursor parameters (start_after, end_before) for
// pagination.
func (r *TreasuryService) List(ctx context.Context, query TreasuryListParams, opts ...option.RequestOption) (res *pagination.CursorIDAccounts[TreasuryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8"), option.WithResponseInto(&raw)}, opts...)
	path := "treasury"
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

// Retrieve a paginated list of all treasury accounts associated with the
// authenticated organization. Use cursor parameters (start_after, end_before) for
// pagination.
func (r *TreasuryService) ListAutoPaging(ctx context.Context, query TreasuryListParams, opts ...option.RequestOption) *pagination.CursorIDAccountsAutoPager[TreasuryListResponse] {
	return pagination.NewCursorIDAccountsAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a paginated list of statements for a specific treasury account.
// Supports cursor-based pagination and filtering by document type.
func (r *TreasuryService) GetStatements(ctx context.Context, treasuryID string, query TreasuryGetStatementsParams, opts ...option.RequestOption) (res *pagination.CursorIDAccountStatements[TreasuryGetStatementsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8"), option.WithResponseInto(&raw)}, opts...)
	if treasuryID == "" {
		err = errors.New("missing required treasuryId parameter")
		return
	}
	path := fmt.Sprintf("treasury/%s/statements", treasuryID)
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

// Retrieve a paginated list of statements for a specific treasury account.
// Supports cursor-based pagination and filtering by document type.
func (r *TreasuryService) GetStatementsAutoPaging(ctx context.Context, treasuryID string, query TreasuryGetStatementsParams, opts ...option.RequestOption) *pagination.CursorIDAccountStatementsAutoPager[TreasuryGetStatementsResponse] {
	return pagination.NewCursorIDAccountStatementsAutoPager(r.GetStatements(ctx, treasuryID, query, opts...))
}

// Retrieve paginated treasury transactions for a specific treasury account.
func (r *TreasuryService) GetTransactions(ctx context.Context, treasuryID string, query TreasuryGetTransactionsParams, opts ...option.RequestOption) (res *TreasuryGetTransactionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if treasuryID == "" {
		err = errors.New("missing required treasuryId parameter")
		return
	}
	path := fmt.Sprintf("treasury/%s/transactions", treasuryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type TreasuryListResponse struct {
	// ID for a Mercury account.
	ID               string  `json:"id" api:"required" format:"uuid"`
	AvailableBalance float64 `json:"availableBalance" api:"required"`
	CreatedAt        string  `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	CurrentBalance   float64 `json:"currentBalance" api:"required"`
	// Monthly net return breakdown with dividend and fee details
	NetReturns []TreasuryListResponseNetReturn `json:"netReturns" api:"required"`
	// Any of "active", "deleted", "pending", "archived".
	Status AccountStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AvailableBalance respjson.Field
		CreatedAt        respjson.Field
		CurrentBalance   respjson.Field
		NetReturns       respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TreasuryListResponse) RawJSON() string { return r.JSON.raw }
func (r *TreasuryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly net return breakdown for a treasury account
type TreasuryListResponseNetReturn struct {
	// List of dividends received by security
	Dividends []TreasuryListResponseNetReturnDividend `json:"dividends" api:"required"`
	// First day of the month for this net return
	Month time.Time `json:"month" api:"required" format:"date"`
	// Net return amount (dividends minus fees)
	NetAmount float64 `json:"netAmount" api:"required"`
	// Status of this net return calculation
	//
	// Any of "processing", "pending", "charged", "error".
	Status string `json:"status" api:"required"`
	// Treasury fee charged for this period (positive value)
	TreasuryFee float64 `json:"treasuryFee" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dividends   respjson.Field
		Month       respjson.Field
		NetAmount   respjson.Field
		Status      respjson.Field
		TreasuryFee respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TreasuryListResponseNetReturn) RawJSON() string { return r.JSON.raw }
func (r *TreasuryListResponseNetReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dividend information for a specific treasury security
type TreasuryListResponseNetReturnDividend struct {
	// Security identifier (e.g., "617455696")
	ID string `json:"id" api:"required"`
	// Dividend amount for this security
	Amount float64 `json:"amount" api:"required"`
	// Human-readable security name (e.g., "Morgan Stanley Ultra-Short Income Portfolio
	// Class IR")
	SecurityName string `json:"securityName" api:"required"`
	// Security identifier type
	//
	// Any of "cusip".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Amount       respjson.Field
		SecurityName respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TreasuryListResponseNetReturnDividend) RawJSON() string { return r.JSON.raw }
func (r *TreasuryListResponseNetReturnDividend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Individual treasury statement in the response
type TreasuryGetStatementsResponse struct {
	// ID for the account statement
	ID string `json:"id" api:"required" format:"uuid"`
	// ID for a Mercury account.
	AccountID string `json:"accountId" api:"required" format:"uuid"`
	// Timestamp when the record was created
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Date the statement was created by the custodian
	CreationDate string `json:"creationDate" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Human-readable description of the statement
	Description string `json:"description" api:"required"`
	// Type of document (e.g. monthly statement, trade confirmation, tax form)
	//
	// Any of "MonthlyStatement", "TradeConfirmation", "1099", "1099R", "1042S",
	// "5498", "5498ESA", "1099Q", "FMV", "SDIRA".
	DocumentType TreasuryGetStatementsResponseDocumentType `json:"documentType" api:"required"`
	// URL to download the statement PDF
	DownloadURL string `json:"downloadUrl" api:"required"`
	// End of the period covered by the statement
	PeriodEnd time.Time `json:"periodEnd" api:"required" format:"date"`
	// Start of the period covered by the statement
	PeriodStart time.Time `json:"periodStart" api:"required" format:"date"`
	// Timestamp when the record was last updated
	UpdatedAt string `json:"updatedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		CreatedAt    respjson.Field
		CreationDate respjson.Field
		Description  respjson.Field
		DocumentType respjson.Field
		DownloadURL  respjson.Field
		PeriodEnd    respjson.Field
		PeriodStart  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TreasuryGetStatementsResponse) RawJSON() string { return r.JSON.raw }
func (r *TreasuryGetStatementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of document (e.g. monthly statement, trade confirmation, tax form)
type TreasuryGetStatementsResponseDocumentType string

const (
	TreasuryGetStatementsResponseDocumentTypeMonthlyStatement  TreasuryGetStatementsResponseDocumentType = "MonthlyStatement"
	TreasuryGetStatementsResponseDocumentTypeTradeConfirmation TreasuryGetStatementsResponseDocumentType = "TradeConfirmation"
	TreasuryGetStatementsResponseDocumentType1099              TreasuryGetStatementsResponseDocumentType = "1099"
	TreasuryGetStatementsResponseDocumentType1099R             TreasuryGetStatementsResponseDocumentType = "1099R"
	TreasuryGetStatementsResponseDocumentType1042S             TreasuryGetStatementsResponseDocumentType = "1042S"
	TreasuryGetStatementsResponseDocumentType5498              TreasuryGetStatementsResponseDocumentType = "5498"
	TreasuryGetStatementsResponseDocumentType5498Esa           TreasuryGetStatementsResponseDocumentType = "5498ESA"
	TreasuryGetStatementsResponseDocumentType1099Q             TreasuryGetStatementsResponseDocumentType = "1099Q"
	TreasuryGetStatementsResponseDocumentTypeFmv               TreasuryGetStatementsResponseDocumentType = "FMV"
	TreasuryGetStatementsResponseDocumentTypeSdira             TreasuryGetStatementsResponseDocumentType = "SDIRA"
)

// Response type for treasury transactions API endpoint
type TreasuryGetTransactionsResponse struct {
	// List of treasury transactions in the response
	Transactions []TreasuryGetTransactionsResponseTransaction `json:"transactions" api:"required"`
	// Pagination cursor for retrieving next batch of transactions
	Cursor int64 `json:"cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transactions respjson.Field
		Cursor       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TreasuryGetTransactionsResponse) RawJSON() string { return r.JSON.raw }
func (r *TreasuryGetTransactionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Treasury transaction data for external API consumption
type TreasuryGetTransactionsResponseTransaction struct {
	// ID for this treasury transaction
	ID string `json:"id" api:"required" format:"uuid"`
	// ID for a Mercury account.
	AccountID    string    `json:"accountId" api:"required" format:"uuid"`
	Amount       float64   `json:"amount" api:"required"`
	Balance      float64   `json:"balance" api:"required"`
	CanonicalDay time.Time `json:"canonicalDay" api:"required" format:"date"`
	Description  string    `json:"description" api:"required"`
	// Any of "depositCanceled", "depositComplete", "depositFailed", "depositReturned",
	// "mercuryFeePosted", "mercuryFeeFailed", "mercuryFeeRefunded",
	// "mercuryFeeCanceled", "withdrawalPosted", "withdrawalFailed",
	// "withdrawalCanceled", "withdrawalReturned", "revertTxn", "interestPosted",
	// "interestCanceled", "manualAmendmentPosted", "mercuryCreditPosted",
	// "mercuryCreditFailed", "dividendPosted", "dividendCanceled",
	// "dividendReinvestmentPosted", "mutualFundTradeFailed", "mutualFundTradePosted",
	// "sweepInPosted", "sweepOutPosted", "sweepReconcilePosted",
	// "valuationChangePosted".
	Type              string                                            `json:"type" api:"required"`
	AdditionalDetails string                                            `json:"additionalDetails" api:"nullable"`
	Details           TreasuryGetTransactionsResponseTransactionDetails `json:"details" api:"nullable"`
	Security          string                                            `json:"security" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AccountID         respjson.Field
		Amount            respjson.Field
		Balance           respjson.Field
		CanonicalDay      respjson.Field
		Description       respjson.Field
		Type              respjson.Field
		AdditionalDetails respjson.Field
		Details           respjson.Field
		Security          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TreasuryGetTransactionsResponseTransaction) RawJSON() string { return r.JSON.raw }
func (r *TreasuryGetTransactionsResponseTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TreasuryGetTransactionsResponseTransactionDetails struct {
	CreditDescription string `json:"creditDescription" api:"nullable"`
	// ID for a Mercury account.
	DepositCounterpartyID      string `json:"depositCounterpartyId" api:"nullable" format:"uuid"`
	FeeDescription             string `json:"feeDescription" api:"nullable"`
	ManualAmendmentDescription string `json:"manualAmendmentDescription" api:"nullable"`
	Security                   string `json:"security" api:"nullable"`
	SweepDirection             string `json:"sweepDirection" api:"nullable"`
	TradeAction                string `json:"tradeAction" api:"nullable"`
	// ID for a Mercury account.
	WithdrawalCounterpartyID string `json:"withdrawalCounterpartyId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditDescription          respjson.Field
		DepositCounterpartyID      respjson.Field
		FeeDescription             respjson.Field
		ManualAmendmentDescription respjson.Field
		Security                   respjson.Field
		SweepDirection             respjson.Field
		TradeAction                respjson.Field
		WithdrawalCounterpartyID   respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TreasuryGetTransactionsResponseTransactionDetails) RawJSON() string { return r.JSON.raw }
func (r *TreasuryGetTransactionsResponseTransactionDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TreasuryListParams struct {
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
	Order TreasuryListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TreasuryListParams]'s query parameters as `url.Values`.
func (r TreasuryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type TreasuryListParamsOrder string

const (
	TreasuryListParamsOrderAsc  TreasuryListParamsOrder = "asc"
	TreasuryListParamsOrderDesc TreasuryListParamsOrder = "desc"
)

type TreasuryGetStatementsParams struct {
	// The ID of the statement to end the page before (exclusive). When provided,
	// results will end just before this ID and work backwards. Use this for reverse
	// pagination or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the statement to start the page after (exclusive). When provided,
	// results will begin with the statement immediately following this ID. Use this
	// for standard forward pagination to get the next page of results. Cannot be
	// combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Filter statements by document type.
	//
	// Any of "MonthlyStatement", "TradeConfirmation", "1099", "1099R", "1042S",
	// "5498", "5498ESA", "1099Q", "FMV", "SDIRA".
	DocumentType TreasuryGetStatementsParamsDocumentType `query:"documentType,omitzero" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order TreasuryGetStatementsParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TreasuryGetStatementsParams]'s query parameters as
// `url.Values`.
func (r TreasuryGetStatementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter statements by document type.
type TreasuryGetStatementsParamsDocumentType string

const (
	TreasuryGetStatementsParamsDocumentTypeMonthlyStatement  TreasuryGetStatementsParamsDocumentType = "MonthlyStatement"
	TreasuryGetStatementsParamsDocumentTypeTradeConfirmation TreasuryGetStatementsParamsDocumentType = "TradeConfirmation"
	TreasuryGetStatementsParamsDocumentType1099              TreasuryGetStatementsParamsDocumentType = "1099"
	TreasuryGetStatementsParamsDocumentType1099R             TreasuryGetStatementsParamsDocumentType = "1099R"
	TreasuryGetStatementsParamsDocumentType1042S             TreasuryGetStatementsParamsDocumentType = "1042S"
	TreasuryGetStatementsParamsDocumentType5498              TreasuryGetStatementsParamsDocumentType = "5498"
	TreasuryGetStatementsParamsDocumentType5498Esa           TreasuryGetStatementsParamsDocumentType = "5498ESA"
	TreasuryGetStatementsParamsDocumentType1099Q             TreasuryGetStatementsParamsDocumentType = "1099Q"
	TreasuryGetStatementsParamsDocumentTypeFmv               TreasuryGetStatementsParamsDocumentType = "FMV"
	TreasuryGetStatementsParamsDocumentTypeSdira             TreasuryGetStatementsParamsDocumentType = "SDIRA"
)

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type TreasuryGetStatementsParamsOrder string

const (
	TreasuryGetStatementsParamsOrderAsc  TreasuryGetStatementsParamsOrder = "asc"
	TreasuryGetStatementsParamsOrderDesc TreasuryGetStatementsParamsOrder = "desc"
)

type TreasuryGetTransactionsParams struct {
	// Pagination cursor for retrieving next batch of transactions. Must be an
	// integer >= 0
	Cursor param.Opt[int64] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return. Defaults to 100
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order for transactions. Can be 'asc' or 'desc'. Defaults to 'desc'
	//
	// Any of "asc", "desc".
	Order TreasuryGetTransactionsParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TreasuryGetTransactionsParams]'s query parameters as
// `url.Values`.
func (r TreasuryGetTransactionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order for transactions. Can be 'asc' or 'desc'. Defaults to 'desc'
type TreasuryGetTransactionsParamsOrder string

const (
	TreasuryGetTransactionsParamsOrderAsc  TreasuryGetTransactionsParamsOrder = "asc"
	TreasuryGetTransactionsParamsOrderDesc TreasuryGetTransactionsParamsOrder = "desc"
)
