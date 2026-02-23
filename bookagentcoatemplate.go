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

// BookAgentCoaTemplateService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBookAgentCoaTemplateService] method instead.
type BookAgentCoaTemplateService struct {
	Options []option.RequestOption
}

// NewBookAgentCoaTemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBookAgentCoaTemplateService(opts ...option.RequestOption) (r BookAgentCoaTemplateService) {
	r = BookAgentCoaTemplateService{}
	r.Options = opts
	return
}

// Create a new agent-owned Chart of Accounts template. These templates can be used
// when creating new Books instances for clients.
func (r *BookAgentCoaTemplateService) New(ctx context.Context, params BookAgentCoaTemplateNewParams, opts ...option.RequestOption) (res *ChartOfAccountsTemplateForTeal, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "books/agent-coa-templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve details of a specific Chart of Accounts template.
func (r *BookAgentCoaTemplateService) Get(ctx context.Context, coaTemplateID string, query BookAgentCoaTemplateGetParams, opts ...option.RequestOption) (res *ChartOfAccountsTemplateForTeal, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if coaTemplateID == "" {
		err = errors.New("missing required coaTemplateId parameter")
		return
	}
	path := fmt.Sprintf("books/agent-coa-template/%s", coaTemplateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve a paginated list of all default and agent-owned Chart of Accounts
// templates. These templates can be used when creating new Books instances for
// clients.
func (r *BookAgentCoaTemplateService) List(ctx context.Context, query BookAgentCoaTemplateListParams, opts ...option.RequestOption) (res *BookAgentCoaTemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "books/agent-coa-templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a specific Chart of Accounts template.
func (r *BookAgentCoaTemplateService) Delete(ctx context.Context, coaTemplateID string, opts ...option.RequestOption) (res *[]BookAgentCoaTemplateDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if coaTemplateID == "" {
		err = errors.New("missing required coaTemplateId parameter")
		return
	}
	path := fmt.Sprintf("books/agent-coa-template/%s", coaTemplateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// A chart of accounts template
type ChartOfAccountsTemplateForBooks struct {
	// The ID of the template
	ID string `json:"id,required" format:"base58 (max length 22)"`
	// The accounting basis of the template
	//
	// Any of "cash", "accrual".
	AccountingBasis ChartOfAccountsTemplateForBooksAccountingBasis `json:"accountingBasis,required"`
	// The name of the template
	Name string `json:"name,required"`
	// Whether this template belongs to an agent or a platform
	//
	// Any of "agent", "platform".
	Type ChartOfAccountsTemplateForBooksType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountingBasis respjson.Field
		Name            respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChartOfAccountsTemplateForBooks) RawJSON() string { return r.JSON.raw }
func (r *ChartOfAccountsTemplateForBooks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The accounting basis of the template
type ChartOfAccountsTemplateForBooksAccountingBasis string

const (
	ChartOfAccountsTemplateForBooksAccountingBasisCash    ChartOfAccountsTemplateForBooksAccountingBasis = "cash"
	ChartOfAccountsTemplateForBooksAccountingBasisAccrual ChartOfAccountsTemplateForBooksAccountingBasis = "accrual"
)

// Whether this template belongs to an agent or a platform
type ChartOfAccountsTemplateForBooksType string

const (
	ChartOfAccountsTemplateForBooksTypeAgent    ChartOfAccountsTemplateForBooksType = "agent"
	ChartOfAccountsTemplateForBooksTypePlatform ChartOfAccountsTemplateForBooksType = "platform"
)

type ChartOfAccountsTemplateForTeal struct {
	// The unique ID of the object.
	ID string `json:"id,required" format:"base58 (max length 22)"`
	// The accounting basis of the Chart of Accounts template
	//
	// Any of "cash", "accrual".
	AccountingBasis ChartOfAccountsTemplateForTealAccountingBasis `json:"accounting_basis,required"`
	// List of IDs of the Ledger Template objects associated with the Chart of Accounts
	// template
	LedgerIDs []string `json:"ledger_ids,required" format:"base58 (max length 22)"`
	// Chart of Accounts template name
	Name string `json:"name,required"`
	// Whether this template belongs to an agent or a platform
	//
	// Any of "agent", "platform".
	Type ChartOfAccountsTemplateForTealType `json:"type,required"`
	// [Expandable] List of the Ledger Templates associated with the Chart of Accounts
	// template
	Ledgers []LedgerTemplate `json:"ledgers,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccountingBasis respjson.Field
		LedgerIDs       respjson.Field
		Name            respjson.Field
		Type            respjson.Field
		Ledgers         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChartOfAccountsTemplateForTeal) RawJSON() string { return r.JSON.raw }
func (r *ChartOfAccountsTemplateForTeal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The accounting basis of the Chart of Accounts template
type ChartOfAccountsTemplateForTealAccountingBasis string

const (
	ChartOfAccountsTemplateForTealAccountingBasisCash    ChartOfAccountsTemplateForTealAccountingBasis = "cash"
	ChartOfAccountsTemplateForTealAccountingBasisAccrual ChartOfAccountsTemplateForTealAccountingBasis = "accrual"
)

// Whether this template belongs to an agent or a platform
type ChartOfAccountsTemplateForTealType string

const (
	ChartOfAccountsTemplateForTealTypeAgent    ChartOfAccountsTemplateForTealType = "agent"
	ChartOfAccountsTemplateForTealTypePlatform ChartOfAccountsTemplateForTealType = "platform"
)

// Response containing a paginated list of Chart of Accounts templates
type BookAgentCoaTemplateListResponse struct {
	// Agent-created COA templates
	AgentTemplates []ChartOfAccountsTemplateForBooks `json:"agentTemplates,required"`
	// Platform default COA templates available to this agent
	DefaultTemplates []ChartOfAccountsTemplateForBooks `json:"defaultTemplates,required"`
	// Token for fetching the next page of results
	NextPageToken string `json:"nextPageToken,nullable"`
	// Token for fetching the previous page of results
	PrevPageToken string `json:"prevPageToken,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentTemplates   respjson.Field
		DefaultTemplates respjson.Field
		NextPageToken    respjson.Field
		PrevPageToken    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BookAgentCoaTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *BookAgentCoaTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BookAgentCoaTemplateDeleteResponse = any

type BookAgentCoaTemplateNewParams struct {
	// The accounting basis for this Chart of Accounts template.
	//
	// Any of "cash", "accrual".
	AccountingBasis BookAgentCoaTemplateNewParamsAccountingBasis `json:"accounting_basis,omitzero,required"`
	// An arbitrary string on the object, useful for identifying the Chart of Accounts
	// template.
	Name string `json:"name,required"`
	// Comma-separated list of expandable paths.
	//
	// Any of "ledgers".
	Expand BookAgentCoaTemplateNewParamsExpand `query:"expand,omitzero" json:"-"`
	paramObj
}

func (r BookAgentCoaTemplateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BookAgentCoaTemplateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BookAgentCoaTemplateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [BookAgentCoaTemplateNewParams]'s query parameters as
// `url.Values`.
func (r BookAgentCoaTemplateNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The accounting basis for this Chart of Accounts template.
type BookAgentCoaTemplateNewParamsAccountingBasis string

const (
	BookAgentCoaTemplateNewParamsAccountingBasisCash    BookAgentCoaTemplateNewParamsAccountingBasis = "cash"
	BookAgentCoaTemplateNewParamsAccountingBasisAccrual BookAgentCoaTemplateNewParamsAccountingBasis = "accrual"
)

// Comma-separated list of expandable paths.
type BookAgentCoaTemplateNewParamsExpand string

const (
	BookAgentCoaTemplateNewParamsExpandLedgers BookAgentCoaTemplateNewParamsExpand = "ledgers"
)

type BookAgentCoaTemplateGetParams struct {
	// Comma-separated list of expandable paths.
	//
	// Any of "ledgers".
	Expand BookAgentCoaTemplateGetParamsExpand `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BookAgentCoaTemplateGetParams]'s query parameters as
// `url.Values`.
func (r BookAgentCoaTemplateGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Comma-separated list of expandable paths.
type BookAgentCoaTemplateGetParamsExpand string

const (
	BookAgentCoaTemplateGetParamsExpandLedgers BookAgentCoaTemplateGetParamsExpand = "ledgers"
)

type BookAgentCoaTemplateListParams struct {
	// The number of records to return. Max limit is 100.
	Limit     param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	// Comma-separated list of expandable paths.
	//
	// Any of "ledgers".
	Expand BookAgentCoaTemplateListParamsExpand `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BookAgentCoaTemplateListParams]'s query parameters as
// `url.Values`.
func (r BookAgentCoaTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Comma-separated list of expandable paths.
type BookAgentCoaTemplateListParamsExpand string

const (
	BookAgentCoaTemplateListParamsExpandLedgers BookAgentCoaTemplateListParamsExpand = "ledgers"
)
