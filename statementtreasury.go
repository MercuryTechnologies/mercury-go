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

// Manage treasury accounts and transactions
//
// StatementTreasuryService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatementTreasuryService] method instead.
type StatementTreasuryService struct {
	Options []option.RequestOption
}

// NewStatementTreasuryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStatementTreasuryService(opts ...option.RequestOption) (r StatementTreasuryService) {
	r = StatementTreasuryService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of statements for a specific treasury account.
// Supports cursor-based pagination and filtering by document type.
func (r *StatementTreasuryService) List(ctx context.Context, treasuryID string, query StatementTreasuryListParams, opts ...option.RequestOption) (res *pagination.CursorIDAccountStatements[StatementTreasuryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if treasuryID == "" {
		err = errors.New("missing required treasuryId parameter")
		return nil, err
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
func (r *StatementTreasuryService) ListAutoPaging(ctx context.Context, treasuryID string, query StatementTreasuryListParams, opts ...option.RequestOption) *pagination.CursorIDAccountStatementsAutoPager[StatementTreasuryListResponse] {
	return pagination.NewCursorIDAccountStatementsAutoPager(r.List(ctx, treasuryID, query, opts...))
}

// Individual treasury statement in the response
type StatementTreasuryListResponse struct {
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
	DocumentType StatementTreasuryListResponseDocumentType `json:"documentType" api:"required"`
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
func (r StatementTreasuryListResponse) RawJSON() string { return r.JSON.raw }
func (r *StatementTreasuryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of document (e.g. monthly statement, trade confirmation, tax form)
type StatementTreasuryListResponseDocumentType string

const (
	StatementTreasuryListResponseDocumentTypeMonthlyStatement  StatementTreasuryListResponseDocumentType = "MonthlyStatement"
	StatementTreasuryListResponseDocumentTypeTradeConfirmation StatementTreasuryListResponseDocumentType = "TradeConfirmation"
	StatementTreasuryListResponseDocumentType1099              StatementTreasuryListResponseDocumentType = "1099"
	StatementTreasuryListResponseDocumentType1099R             StatementTreasuryListResponseDocumentType = "1099R"
	StatementTreasuryListResponseDocumentType1042S             StatementTreasuryListResponseDocumentType = "1042S"
	StatementTreasuryListResponseDocumentType5498              StatementTreasuryListResponseDocumentType = "5498"
	StatementTreasuryListResponseDocumentType5498Esa           StatementTreasuryListResponseDocumentType = "5498ESA"
	StatementTreasuryListResponseDocumentType1099Q             StatementTreasuryListResponseDocumentType = "1099Q"
	StatementTreasuryListResponseDocumentTypeFmv               StatementTreasuryListResponseDocumentType = "FMV"
	StatementTreasuryListResponseDocumentTypeSdira             StatementTreasuryListResponseDocumentType = "SDIRA"
)

type StatementTreasuryListParams struct {
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
	DocumentType StatementTreasuryListParamsDocumentType `query:"documentType,omitzero" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order StatementTreasuryListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatementTreasuryListParams]'s query parameters as
// `url.Values`.
func (r StatementTreasuryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter statements by document type.
type StatementTreasuryListParamsDocumentType string

const (
	StatementTreasuryListParamsDocumentTypeMonthlyStatement  StatementTreasuryListParamsDocumentType = "MonthlyStatement"
	StatementTreasuryListParamsDocumentTypeTradeConfirmation StatementTreasuryListParamsDocumentType = "TradeConfirmation"
	StatementTreasuryListParamsDocumentType1099              StatementTreasuryListParamsDocumentType = "1099"
	StatementTreasuryListParamsDocumentType1099R             StatementTreasuryListParamsDocumentType = "1099R"
	StatementTreasuryListParamsDocumentType1042S             StatementTreasuryListParamsDocumentType = "1042S"
	StatementTreasuryListParamsDocumentType5498              StatementTreasuryListParamsDocumentType = "5498"
	StatementTreasuryListParamsDocumentType5498Esa           StatementTreasuryListParamsDocumentType = "5498ESA"
	StatementTreasuryListParamsDocumentType1099Q             StatementTreasuryListParamsDocumentType = "1099Q"
	StatementTreasuryListParamsDocumentTypeFmv               StatementTreasuryListParamsDocumentType = "FMV"
	StatementTreasuryListParamsDocumentTypeSdira             StatementTreasuryListParamsDocumentType = "SDIRA"
)

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type StatementTreasuryListParamsOrder string

const (
	StatementTreasuryListParamsOrderAsc  StatementTreasuryListParamsOrder = "asc"
	StatementTreasuryListParamsOrderDesc StatementTreasuryListParamsOrder = "desc"
)
