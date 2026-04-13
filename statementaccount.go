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

// Manage bank accounts
//
// StatementAccountService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatementAccountService] method instead.
type StatementAccountService struct {
	Options []option.RequestOption
}

// NewStatementAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStatementAccountService(opts ...option.RequestOption) (r StatementAccountService) {
	r = StatementAccountService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of monthly statements for a specific account. Supports
// cursor-based pagination with limit, order, start_after, and end_before query
// parameters, as well as date range filtering with start and end parameters.
func (r *StatementAccountService) List(ctx context.Context, accountID string, query StatementAccountListParams, opts ...option.RequestOption) (res *pagination.CursorIDAccountStatements[StatementAccountListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/%s/statements", accountID)
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

// Retrieve a paginated list of monthly statements for a specific account. Supports
// cursor-based pagination with limit, order, start_after, and end_before query
// parameters, as well as date range filtering with start and end parameters.
func (r *StatementAccountService) ListAutoPaging(ctx context.Context, accountID string, query StatementAccountListParams, opts ...option.RequestOption) *pagination.CursorIDAccountStatementsAutoPager[StatementAccountListResponse] {
	return pagination.NewCursorIDAccountStatementsAutoPager(r.List(ctx, accountID, query, opts...))
}

type StatementAccountListResponse struct {
	// ID for the account statement
	ID                  string  `json:"id" api:"required" format:"uuid"`
	AccountNumber       string  `json:"accountNumber" api:"required"`
	CompanyLegalAddress Address `json:"companyLegalAddress" api:"required"`
	CompanyLegalName    string  `json:"companyLegalName" api:"required"`
	DownloadURL         string  `json:"downloadUrl" api:"required"`
	EndDate             string  `json:"endDate" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// A dollar amount
	EndingBalance float64                                   `json:"endingBalance" api:"required"`
	RoutingNumber string                                    `json:"routingNumber" api:"required"`
	StartDate     string                                    `json:"startDate" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	Transactions  []StatementAccountListResponseTransaction `json:"transactions" api:"required"`
	Ein           string                                    `json:"ein" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AccountNumber       respjson.Field
		CompanyLegalAddress respjson.Field
		CompanyLegalName    respjson.Field
		DownloadURL         respjson.Field
		EndDate             respjson.Field
		EndingBalance       respjson.Field
		RoutingNumber       respjson.Field
		StartDate           respjson.Field
		Transactions        respjson.Field
		Ein                 respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatementAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *StatementAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatementAccountListResponseTransaction struct {
	// ID for this transaction
	ID        string `json:"id" api:"required" format:"uuid"`
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	PostedAt  string `json:"postedAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		PostedAt    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatementAccountListResponseTransaction) RawJSON() string { return r.JSON.raw }
func (r *StatementAccountListResponseTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatementAccountListParams struct {
	// Filter statements where the period start date is on or before this date. If the
	// date is in the future, defaults to the current date. Format: YYYY-MM-DD
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// The ID of the statement to end the page before (exclusive). When provided,
	// results will end just before this ID and work backwards. Use this for reverse
	// pagination or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter statements where the period start date is on or after this date. Format:
	// YYYY-MM-DD
	Start param.Opt[string] `query:"start,omitzero" json:"-"`
	// The ID of the statement to start the page after (exclusive). When provided,
	// results will begin with the statement immediately following this ID. Use this
	// for standard forward pagination to get the next page of results. Cannot be
	// combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'
	//
	// Any of "asc", "desc".
	Order StatementAccountListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StatementAccountListParams]'s query parameters as
// `url.Values`.
func (r StatementAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'
type StatementAccountListParamsOrder string

const (
	StatementAccountListParamsOrderAsc  StatementAccountListParamsOrder = "asc"
	StatementAccountListParamsOrderDesc StatementAccountListParamsOrder = "desc"
)
