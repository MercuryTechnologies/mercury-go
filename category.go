// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
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

// CategoryService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCategoryService] method instead.
type CategoryService struct {
	Options []option.RequestOption
}

// NewCategoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCategoryService(opts ...option.RequestOption) (r CategoryService) {
	r = CategoryService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of all available custom expense categories for the
// organization. Supports cursor-based pagination with limit, order, start_after,
// and end_before query parameters.
func (r *CategoryService) List(ctx context.Context, query CategoryListParams, opts ...option.RequestOption) (res *CategoryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Paginated response containing a list of categories. | Use the page cursor
// information to fetch additional pages of categories.
type CategoryListResponse struct {
	// List of categories in the current page
	Categories []CategoryData `json:"categories" api:"required"`
	// Pagination information including cursors for navigating to next/previous pages
	Page CategoryListResponsePage `json:"page" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories  respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CategoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *CategoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination information including cursors for navigating to next/previous pages
type CategoryListResponsePage struct {
	// ID for the category
	NextPage string `json:"nextPage" format:"uuid"`
	// ID for the category
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
func (r CategoryListResponsePage) RawJSON() string { return r.JSON.raw }
func (r *CategoryListResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CategoryListParams struct {
	// The ID of the category to end the page before (exclusive). When provided,
	// results will end just before this ID and work backwards. Use this for reverse
	// pagination or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the category to start the page after (exclusive). When provided,
	// results will begin with the category immediately following this ID. Use this for
	// standard forward pagination to get the next page of results. Cannot be combined
	// with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order CategoryListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CategoryListParams]'s query parameters as `url.Values`.
func (r CategoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type CategoryListParamsOrder string

const (
	CategoryListParamsOrderAsc  CategoryListParamsOrder = "asc"
	CategoryListParamsOrderDesc CategoryListParamsOrder = "desc"
)
