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
	"github.com/stainless-sdks/mercury-go/packages/pagination"
	"github.com/stainless-sdks/mercury-go/packages/param"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// AccountsRecievableCustomerService contains methods and other services that help
// with interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountsRecievableCustomerService] method instead.
type AccountsRecievableCustomerService struct {
	Options []option.RequestOption
}

// NewAccountsRecievableCustomerService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountsRecievableCustomerService(opts ...option.RequestOption) (r AccountsRecievableCustomerService) {
	r = AccountsRecievableCustomerService{}
	r.Options = opts
	return
}

// Create a new customer for the organization
func (r *AccountsRecievableCustomerService) New(ctx context.Context, body AccountsRecievableCustomerNewParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ar/customers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details of a specific customer by their ID
func (r *AccountsRecievableCustomerService) Get(ctx context.Context, customerID string, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.Options, opts)
	if customerID == "" {
		err = errors.New("missing required customerId parameter")
		return
	}
	path := fmt.Sprintf("ar/customers/%s", customerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing customer
func (r *AccountsRecievableCustomerService) Update(ctx context.Context, customerID string, body AccountsRecievableCustomerUpdateParams, opts ...option.RequestOption) (res *Customer, err error) {
	opts = slices.Concat(r.Options, opts)
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
func (r *AccountsRecievableCustomerService) List(ctx context.Context, query AccountsRecievableCustomerListParams, opts ...option.RequestOption) (res *pagination.CursorIDArCustomers[Customer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ar/customers"
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

// Retrieve a paginated list of customers. Supports cursor-based pagination with
// limit, order, start_after, and end_before query parameters.
func (r *AccountsRecievableCustomerService) ListAutoPaging(ctx context.Context, query AccountsRecievableCustomerListParams, opts ...option.RequestOption) *pagination.CursorIDArCustomersAutoPager[Customer] {
	return pagination.NewCursorIDArCustomersAutoPager(r.List(ctx, query, opts...))
}

// Delete a customer. This action cannot be undone.
func (r *AccountsRecievableCustomerService) Delete(ctx context.Context, customerID string, opts ...option.RequestOption) (err error) {
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
	Address1 string `json:"address1" api:"required"`
	// City name.
	City string `json:"city" api:"required"`
	// Two-letter country code (ISO 3166-1 alpha-2).
	Country string `json:"country" api:"required"`
	// The mailing name of the address.
	Name string `json:"name" api:"required"`
	// Postal or ZIP code.
	PostalCode string `json:"postalCode" api:"required"`
	// Either a two-letter US state code i.e. "CA" for California or a free-form
	// identification of a particular region worldwide.
	Region string `json:"region" api:"required"`
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
type Customer struct {
	// The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint
	// to list your customers and find the corresponding id, or create a new customer
	// first.
	ID string `json:"id" api:"required" format:"uuid"`
	// Email of customer.
	Email string `json:"email" api:"required"`
	// Name of customer.
	Name string `json:"name" api:"required"`
	// Customer address information for Accounts Receivable API
	Address CustomerAddress `json:"address" api:"nullable"`
	// The time the customer was deleted, if it was deleted.
	DeletedAt string `json:"deletedAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
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
func (r Customer) RawJSON() string { return r.JSON.raw }
func (r *Customer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Customer address information for Accounts Receivable API
type CustomerAddress struct {
	// Primary street address line.
	Address1 string `json:"address1" api:"required"`
	// City name.
	City string `json:"city" api:"required"`
	// Two-letter country code (ISO 3166-1 alpha-2).
	Country string `json:"country" api:"required"`
	// Postal or ZIP code
	PostalCode string `json:"postalCode" api:"required"`
	// State, province, or region.
	Region string `json:"region" api:"required"`
	// Secondary street address line (optional).
	Address2 string `json:"address2" api:"nullable"`
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
func (r CustomerAddress) RawJSON() string { return r.JSON.raw }
func (r *CustomerAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountsRecievableCustomerNewParams struct {
	// The email address for the customer.
	Email string `json:"email" api:"required"`
	// The name of the customer.
	Name string `json:"name" api:"required"`
	// Address input for creating or updating customers
	Address AddressInputParam `json:"address,omitzero"`
	paramObj
}

func (r AccountsRecievableCustomerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountsRecievableCustomerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountsRecievableCustomerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountsRecievableCustomerUpdateParams struct {
	// The email address for the customer.
	Email string `json:"email" api:"required"`
	// The name of the customer.
	Name string `json:"name" api:"required"`
	// Open invoices for the customer will be resent with updated data when this is
	// true.
	ResendOpenInvoices bool `json:"resendOpenInvoices" api:"required"`
	// Address input for creating or updating customers
	Address AddressInputParam `json:"address,omitzero"`
	paramObj
}

func (r AccountsRecievableCustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountsRecievableCustomerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountsRecievableCustomerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountsRecievableCustomerListParams struct {
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
	Order AccountsRecievableCustomerListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountsRecievableCustomerListParams]'s query parameters as
// `url.Values`.
func (r AccountsRecievableCustomerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type AccountsRecievableCustomerListParamsOrder string

const (
	AccountsRecievableCustomerListParamsOrderAsc  AccountsRecievableCustomerListParamsOrder = "asc"
	AccountsRecievableCustomerListParamsOrderDesc AccountsRecievableCustomerListParamsOrder = "desc"
)
