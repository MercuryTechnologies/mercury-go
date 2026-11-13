// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	"github.com/MercuryTechnologies/mercury-go/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// List merchants available for spend controls
//
// MerchantService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMerchantService] method instead.
type MerchantService struct {
	options []option.RequestOption
}

// NewMerchantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMerchantService(opts ...option.RequestOption) (r MerchantService) {
	r = MerchantService{}
	r.options = opts
	return
}

// Retrieve a paginated list of priority merchants that can be used for spend
// controls like merchant locking. Supports cursor-based pagination with limit,
// order, start_after, and end_before query parameters, plus an optional
// case-insensitive search by merchant name.
func (r *MerchantService) List(ctx context.Context, query MerchantListParams, opts ...option.RequestOption) (res *MerchantListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "merchants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Paginated response for merchants endpoint
type MerchantListResponse struct {
	Data []MerchantInfo           `json:"data" api:"required"`
	Page MerchantListResponsePage `json:"page" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MerchantListResponse) RawJSON() string { return r.JSON.raw }
func (r *MerchantListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MerchantListResponsePage struct {
	NextPage     string `json:"nextPage" format:"uuid"`
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
func (r MerchantListResponsePage) RawJSON() string { return r.JSON.raw }
func (r *MerchantListResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MerchantListParams struct {
	// The ID of the merchant to end the page before (exclusive). When provided,
	// results will end just before this ID and work backwards. Use this for reverse
	// pagination or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Case-insensitive search term to filter merchants by name.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// The ID of the merchant to start the page after (exclusive). When provided,
	// results will begin with the merchant immediately following this ID. Use this for
	// standard forward pagination to get the next page of results. Cannot be combined
	// with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order MerchantListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MerchantListParams]'s query parameters as `url.Values`.
func (r MerchantListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type MerchantListParamsOrder string

const (
	MerchantListParamsOrderAsc  MerchantListParamsOrder = "asc"
	MerchantListParamsOrderDesc MerchantListParamsOrder = "desc"
)
