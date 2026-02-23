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
func (r *TreasuryService) List(ctx context.Context, query TreasuryListParams, opts ...option.RequestOption) (res *pagination.CursorPage[TreasuryListResponse], err error) {
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
func (r *TreasuryService) ListAutoPaging(ctx context.Context, query TreasuryListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[TreasuryListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a paginated list of statements for a specific treasury account.
// Supports cursor-based pagination and filtering by document type.
func (r *TreasuryService) GetStatements(ctx context.Context, treasuryID string, query TreasuryGetStatementsParams, opts ...option.RequestOption) (res *pagination.CursorPage[TreasuryGetStatementsResponse], err error) {
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
func (r *TreasuryService) GetStatementsAutoPaging(ctx context.Context, treasuryID string, query TreasuryGetStatementsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[TreasuryGetStatementsResponse] {
	return pagination.NewCursorPageAutoPager(r.GetStatements(ctx, treasuryID, query, opts...))
}

// Retrieve paginated treasury transactions for a specific treasury account.
func (r *TreasuryService) GetTransactions(ctx context.Context, treasuryID string, query TreasuryGetTransactionsParams, opts ...option.RequestOption) (res *pagination.TreasuryCursorPage[TreasuryGetTransactionsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8"), option.WithResponseInto(&raw)}, opts...)
	if treasuryID == "" {
		err = errors.New("missing required treasuryId parameter")
		return
	}
	path := fmt.Sprintf("treasury/%s/transactions", treasuryID)
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

// Retrieve paginated treasury transactions for a specific treasury account.
func (r *TreasuryService) GetTransactionsAutoPaging(ctx context.Context, treasuryID string, query TreasuryGetTransactionsParams, opts ...option.RequestOption) *pagination.TreasuryCursorPageAutoPager[TreasuryGetTransactionsResponse] {
	return pagination.NewTreasuryCursorPageAutoPager(r.GetTransactions(ctx, treasuryID, query, opts...))
}

// Paginated response type for treasury accounts API endpoint
type TreasuryListResponse struct {
	// List of treasury accounts in the current page
	Accounts []TreasuryListResponseAccount `json:"accounts,required"`
	// Pagination information including cursors for navigating to next/previous pages
	Page TreasuryListResponsePage `json:"page,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TreasuryListResponse) RawJSON() string { return r.JSON.raw }
func (r *TreasuryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TreasuryListResponseAccount struct {
	// ID for a Mercury account.
	ID               string  `json:"id,required" format:"uuid"`
	AvailableBalance float64 `json:"availableBalance,required"`
	CreatedAt        string  `json:"createdAt,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	CurrentBalance   float64 `json:"currentBalance,required"`
	// Monthly net return breakdown with dividend and fee details
	NetReturns []TreasuryListResponseAccountNetReturn `json:"netReturns,required"`
	// Any of "active", "deleted", "pending", "archived".
	Status AccountStatus `json:"status,required"`
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
func (r TreasuryListResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *TreasuryListResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Monthly net return breakdown for a treasury account
type TreasuryListResponseAccountNetReturn struct {
	// List of dividends received by security
	Dividends []TreasuryListResponseAccountNetReturnDividend `json:"dividends,required"`
	// First day of the month for this net return
	Month time.Time `json:"month,required" format:"date"`
	// Net return amount (dividends minus fees)
	NetAmount float64 `json:"netAmount,required"`
	// Status of this net return calculation
	//
	// Any of "processing", "pending", "charged", "error".
	Status string `json:"status,required"`
	// Treasury fee charged for this period (positive value)
	TreasuryFee float64 `json:"treasuryFee,required"`
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
func (r TreasuryListResponseAccountNetReturn) RawJSON() string { return r.JSON.raw }
func (r *TreasuryListResponseAccountNetReturn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Dividend information for a specific treasury security
type TreasuryListResponseAccountNetReturnDividend struct {
	// Security identifier (e.g., "617455696")
	ID string `json:"id,required"`
	// Dividend amount for this security
	Amount float64 `json:"amount,required"`
	// Human-readable security name (e.g., "Morgan Stanley Ultra-Short Income Portfolio
	// Class IR")
	SecurityName string `json:"securityName,required"`
	// Security identifier type
	//
	// Any of "cusip".
	Type string `json:"type,required"`
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
func (r TreasuryListResponseAccountNetReturnDividend) RawJSON() string { return r.JSON.raw }
func (r *TreasuryListResponseAccountNetReturnDividend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination information including cursors for navigating to next/previous pages
type TreasuryListResponsePage struct {
	// ID for a Mercury account.
	NextPage string `json:"nextPage" format:"uuid"`
	// ID for a Mercury account.
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
func (r TreasuryListResponsePage) RawJSON() string { return r.JSON.raw }
func (r *TreasuryListResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated response for treasury account statements
type TreasuryGetStatementsResponse struct {
	Page       TreasuryGetStatementsResponsePage        `json:"page,required"`
	Statements []TreasuryGetStatementsResponseStatement `json:"statements,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		Statements  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TreasuryGetStatementsResponse) RawJSON() string { return r.JSON.raw }
func (r *TreasuryGetStatementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TreasuryGetStatementsResponsePage struct {
	// ID for the account statement
	NextPage string `json:"nextPage" format:"uuid"`
	// ID for the account statement
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
func (r TreasuryGetStatementsResponsePage) RawJSON() string { return r.JSON.raw }
func (r *TreasuryGetStatementsResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Individual treasury statement in the response
type TreasuryGetStatementsResponseStatement struct {
	// ID for the account statement
	ID string `json:"id,required" format:"uuid"`
	// ID for a Mercury account.
	AccountID string `json:"accountId,required" format:"uuid"`
	// Timestamp when the record was created
	CreatedAt string `json:"createdAt,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Date the statement was created by the custodian
	CreationDate string `json:"creationDate,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Human-readable description of the statement
	Description string `json:"description,required"`
	// Type of document (e.g. monthly statement, trade confirmation, tax form)
	//
	// Any of "MonthlyStatement", "TradeConfirmation", "1099", "1099R", "1042S",
	// "5498", "5498ESA", "1099Q", "FMV", "SDIRA".
	DocumentType string `json:"documentType,required"`
	// URL to download the statement PDF
	DownloadURL string `json:"downloadUrl,required"`
	// End of the period covered by the statement
	PeriodEnd time.Time `json:"periodEnd,required" format:"date"`
	// Start of the period covered by the statement
	PeriodStart time.Time `json:"periodStart,required" format:"date"`
	// Timestamp when the record was last updated
	UpdatedAt string `json:"updatedAt,required" format:"yyyy-mm-ddThh:MM:ssZ"`
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
func (r TreasuryGetStatementsResponseStatement) RawJSON() string { return r.JSON.raw }
func (r *TreasuryGetStatementsResponseStatement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response type for treasury transactions API endpoint
type TreasuryGetTransactionsResponse struct {
	// List of treasury transactions in the response
	Transactions []TreasuryGetTransactionsResponseTransaction `json:"transactions,required"`
	// Pagination cursor for retrieving next batch of transactions
	Cursor int64 `json:"cursor,nullable"`
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
	ID string `json:"id,required" format:"uuid"`
	// ID for a Mercury account.
	AccountID    string    `json:"accountId,required" format:"uuid"`
	Amount       float64   `json:"amount,required"`
	Balance      float64   `json:"balance,required"`
	CanonicalDay time.Time `json:"canonicalDay,required" format:"date"`
	Description  string    `json:"description,required"`
	// Any of "depositCanceled", "depositComplete", "depositFailed", "depositReturned",
	// "mercuryFeePosted", "mercuryFeeFailed", "mercuryFeeRefunded",
	// "mercuryFeeCanceled", "withdrawalPosted", "withdrawalFailed",
	// "withdrawalCanceled", "withdrawalReturned", "revertTxn", "interestPosted",
	// "interestCanceled", "manualAmendmentPosted", "mercuryCreditPosted",
	// "mercuryCreditFailed", "dividendPosted", "dividendCanceled",
	// "dividendReinvestmentPosted", "mutualFundTradeFailed", "mutualFundTradePosted",
	// "sweepInPosted", "sweepOutPosted", "sweepReconcilePosted",
	// "valuationChangePosted".
	Type              string                                            `json:"type,required"`
	AdditionalDetails string                                            `json:"additionalDetails,nullable"`
	Details           TreasuryGetTransactionsResponseTransactionDetails `json:"details,nullable"`
	Security          string                                            `json:"security,nullable"`
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
	CreditDescription string `json:"creditDescription,nullable"`
	// ID for a Mercury account.
	DepositCounterpartyID      string `json:"depositCounterpartyId,nullable" format:"uuid"`
	FeeDescription             string `json:"feeDescription,nullable"`
	ManualAmendmentDescription string `json:"manualAmendmentDescription,nullable"`
	Security                   string `json:"security,nullable"`
	SweepDirection             string `json:"sweepDirection,nullable"`
	TradeAction                string `json:"tradeAction,nullable"`
	// ID for a Mercury account.
	WithdrawalCounterpartyID string `json:"withdrawalCounterpartyId,nullable" format:"uuid"`
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
