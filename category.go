// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/mercury-go/internal/apiquery"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/pagination"
	"github.com/stainless-sdks/mercury-go/packages/param"
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
func (r *CategoryService) List(ctx context.Context, query CategoryListParams, opts ...option.RequestOption) (res *pagination.CursorIDCategories[CategoryData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8"), option.WithResponseInto(&raw)}, opts...)
	path := "categories"
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

// Retrieve a paginated list of all available custom expense categories for the
// organization. Supports cursor-based pagination with limit, order, start_after,
// and end_before query parameters.
func (r *CategoryService) ListAutoPaging(ctx context.Context, query CategoryListParams, opts ...option.RequestOption) *pagination.CursorIDCategoriesAutoPager[CategoryData] {
	return pagination.NewCursorIDCategoriesAutoPager(r.List(ctx, query, opts...))
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
