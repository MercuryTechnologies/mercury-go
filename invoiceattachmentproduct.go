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
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// Manage invoicing products
//
// InvoiceAttachmentProductService contains methods and other services that help
// with interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvoiceAttachmentProductService] method instead.
type InvoiceAttachmentProductService struct {
	options []option.RequestOption
}

// NewInvoiceAttachmentProductService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewInvoiceAttachmentProductService(opts ...option.RequestOption) (r InvoiceAttachmentProductService) {
	r = InvoiceAttachmentProductService{}
	r.options = opts
	return
}

// Create a new invoicing product (product catalog item). The combination of name
// and unit price must be unique within the organization.
func (r *InvoiceAttachmentProductService) New(ctx context.Context, body InvoiceAttachmentProductNewParams, opts ...option.RequestOption) (res *APIV1ArProductResponseData, err error) {
	opts = slices.Concat(r.options, opts)
	path := "ar/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of invoicing products (product catalog items).
// Supports cursor-based pagination with limit, order, start_after, and end_before
// query parameters.
func (r *InvoiceAttachmentProductService) List(ctx context.Context, query InvoiceAttachmentProductListParams, opts ...option.RequestOption) (res *InvoiceAttachmentProductListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "ar/products"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete an invoicing product (product catalog item). The product is unlinked from
// line items on open, pending, and recurring invoices.
func (r *InvoiceAttachmentProductService) Delete(ctx context.Context, productID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if productID == "" {
		err = errors.New("missing required productId parameter")
		return err
	}
	path := fmt.Sprintf("ar/products/%s", productID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A product catalog item for Accounts Receivable API endpoints
type APIV1ArProductResponseData struct {
	// The ID of an invoicing product. Use the /api/v1/ar/products endpoint to list
	// your products.
	ID string `json:"id" api:"required" format:"uuid"`
	// Name of the product.
	Name string `json:"name" api:"required"`
	// A dollar amount
	UnitPrice float64 `json:"unitPrice" api:"required"`
	// The time the product's current details were last set.
	UpdatedAt string `json:"updatedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Description of the product, if any.
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		UnitPrice   respjson.Field
		UpdatedAt   respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIV1ArProductResponseData) RawJSON() string { return r.JSON.raw }
func (r *APIV1ArProductResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated response data for Accounts Receivable products API endpoint
type InvoiceAttachmentProductListResponse struct {
	// Pagination cursors for navigating to next/previous pages
	Page InvoiceAttachmentProductListResponsePage `json:"page" api:"required"`
	// The list of products for this page
	Products []APIV1ArProductResponseData `json:"products" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		Products    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceAttachmentProductListResponse) RawJSON() string { return r.JSON.raw }
func (r *InvoiceAttachmentProductListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination cursors for navigating to next/previous pages
type InvoiceAttachmentProductListResponsePage struct {
	// The ID of an invoicing product. Use the /api/v1/ar/products endpoint to list
	// your products.
	NextPage string `json:"nextPage" format:"uuid"`
	// The ID of an invoicing product. Use the /api/v1/ar/products endpoint to list
	// your products.
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
func (r InvoiceAttachmentProductListResponsePage) RawJSON() string { return r.JSON.raw }
func (r *InvoiceAttachmentProductListResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceAttachmentProductNewParams struct {
	// Name of the product. Must be unique per organization in combination with the
	// unit price.
	Name string `json:"name" api:"required"`
	// A dollar amount
	UnitPrice float64 `json:"unitPrice" api:"required"`
	// Optional description of the product.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r InvoiceAttachmentProductNewParams) MarshalJSON() (data []byte, err error) {
	type shadow InvoiceAttachmentProductNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InvoiceAttachmentProductNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceAttachmentProductListParams struct {
	// The ID of the product to end the page before (exclusive). When provided, results
	// will end just before this ID and work backwards. Use this for reverse pagination
	// or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the product to start the page after (exclusive). When provided,
	// results will begin with the product immediately following this ID. Use this for
	// standard forward pagination to get the next page of results. Cannot be combined
	// with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order InvoiceAttachmentProductListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [InvoiceAttachmentProductListParams]'s query parameters as
// `url.Values`.
func (r InvoiceAttachmentProductListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type InvoiceAttachmentProductListParamsOrder string

const (
	InvoiceAttachmentProductListParamsOrderAsc  InvoiceAttachmentProductListParamsOrder = "asc"
	InvoiceAttachmentProductListParamsOrderDesc InvoiceAttachmentProductListParamsOrder = "desc"
)
