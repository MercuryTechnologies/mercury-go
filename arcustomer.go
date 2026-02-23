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

// ArCustomerService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArCustomerService] method instead.
type ArCustomerService struct {
	Options []option.RequestOption
}

// NewArCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewArCustomerService(opts ...option.RequestOption) (r ArCustomerService) {
	r = ArCustomerService{}
	r.Options = opts
	return
}

// Create a new customer for the organization
func (r *ArCustomerService) New(ctx context.Context, body ArCustomerNewParams, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "ar/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details of a specific customer by their ID
func (r *ArCustomerService) Get(ctx context.Context, customerID string, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("ar/customers/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing customer
func (r *ArCustomerService) Update(ctx context.Context, customerID string, body ArCustomerUpdateParams, opts ...option.RequestOption) (res *CustomerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("ar/customers/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of customers. Supports cursor-based pagination with
// limit, order, start_after, and end_before query parameters.
func (r *ArCustomerService) List(ctx context.Context, query ArCustomerListParams, opts ...option.RequestOption) (res *ArCustomerListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "ar/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a customer. This action cannot be undone.
func (r *ArCustomerService) Delete(ctx context.Context, customerID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("ar/customers/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Address input for creating or updating customers
//
// The properties Address1, City, Country, Name, PostalCode, Region are required.
type AddressInputParam struct {
	// Primary street address line.
	Address1 string `json:"address1,required"`
	// City name.
	City string `json:"city,required"`
	// Two-letter country code (ISO 3166-1 alpha-2).
	Country string `json:"country,required"`
	// The mailing name of the address.
	Name string `json:"name,required"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode,required"`
	// Either a two-letter US state code i.e. "CA" for California or a free-form
	// identification of a particular region worldwide.
	Region string `json:"region,required"`
	// Secondary street address line (optional).
	Address2 param.Opt[string] `json:"address2,omitzero"`
	paramObj
}

func (r AddressInputParam) MarshalJSON() (data []byte, err error) {
	type shadow AddressInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response data for Accounts Receivable customer API endpoints
type CustomerResponse struct {
	// The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint
	// to list your customers and find the corresponding id, or create a new customer
	// first.
	ID string `json:"id,required" format:"uuid"`
	// Email of customer.
	Email string `json:"email,required"`
	// Name of customer.
	Name string `json:"name,required"`
	// Customer address information for Accounts Receivable API
	Address CustomerResponseAddress `json:"address,nullable"`
	// The time the customer was deleted, if it was deleted.
	DeletedAt string `json:"deletedAt,nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		Address     respjson.Field
		DeletedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer address information for Accounts Receivable API
type CustomerResponseAddress struct {
	// Primary street address line.
	Address1 string `json:"address1,required"`
	// City name.
	City string `json:"city,required"`
	// Two-letter country code (ISO 3166-1 alpha-2).
	Country string `json:"country,required"`
	// Postal or ZIP code
	PostalCode string `json:"postalCode,required"`
	// State, province, or region.
	Region string `json:"region,required"`
	// Secondary street address line (optional).
	Address2 string `json:"address2,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address1    respjson.Field
		City        respjson.Field
		Country     respjson.Field
		PostalCode  respjson.Field
		Region      respjson.Field
		Address2    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerResponseAddress) RawJSON() string { return r.JSON.raw }
func (r *CustomerResponseAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated response data for Accounts Receivable customers API endpoint
type ArCustomerListResponse struct {
	// The list of customers for this page
	Customers []CustomerResponse `json:"customers,required"`
	// Pagination cursors for navigating to next/previous pages
	Page ArCustomerListResponsePage `json:"page,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Customers   respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ArCustomerListResponse) RawJSON() string { return r.JSON.raw }
func (r *ArCustomerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination cursors for navigating to next/previous pages
type ArCustomerListResponsePage struct {
	// The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint
	// to list your customers and find the corresponding id, or create a new customer
	// first.
	NextPage string `json:"nextPage" format:"uuid"`
	// The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint
	// to list your customers and find the corresponding id, or create a new customer
	// first.
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
func (r ArCustomerListResponsePage) RawJSON() string { return r.JSON.raw }
func (r *ArCustomerListResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArCustomerNewParams struct {
	// The email address for the customer.
	Email string `json:"email,required"`
	// The name of the customer.
	Name string `json:"name,required"`
	// Address input for creating or updating customers
	Address AddressInputParam `json:"address,omitzero"`
	paramObj
}

func (r ArCustomerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ArCustomerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ArCustomerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArCustomerUpdateParams struct {
	// The email address for the customer.
	Email string `json:"email,required"`
	// The name of the customer.
	Name string `json:"name,required"`
	// Open invoices for the customer will be resent with updated data when this is
	// true.
	ResendOpenInvoices bool `json:"resendOpenInvoices,required"`
	// Address input for creating or updating customers
	Address AddressInputParam `json:"address,omitzero"`
	paramObj
}

func (r ArCustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ArCustomerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ArCustomerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArCustomerListParams struct {
	// The ID of the customer to end the page before (exclusive). When provided,
	// results will end just before this ID and work backwards. Use this for reverse
	// pagination or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the customer to start the page after (exclusive). When provided,
	// results will begin with the customer immediately following this ID. Use this for
	// standard forward pagination to get the next page of results. Cannot be combined
	// with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order ArCustomerListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ArCustomerListParams]'s query parameters as `url.Values`.
func (r ArCustomerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type ArCustomerListParamsOrder string

const (
	ArCustomerListParamsOrderAsc  ArCustomerListParamsOrder = "asc"
	ArCustomerListParamsOrderDesc ArCustomerListParamsOrder = "desc"
)
