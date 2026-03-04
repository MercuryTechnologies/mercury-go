// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/mercury-go/internal/apijson"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// Organization information
//
// OrganizationService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r OrganizationService) {
	r = OrganizationService{}
	r.Options = opts
	return
}

// Retrieve information about your organization including EIN, legal business name,
// and DBAs.
func (r *OrganizationService) Get(ctx context.Context, opts ...option.RequestOption) (res *OrganizationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "organization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Response containing organization details.
type OrganizationGetResponse struct {
	// Organization information
	Organization OrganizationGetResponseOrganization `json:"organization" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization information
type OrganizationGetResponseOrganization struct {
	// Unique identifier for the organization
	ID string `json:"id" api:"required" format:"uuid"`
	// List of DBAs (Doing Business As names) for this organization
	Dbas []OrganizationGetResponseOrganizationDba `json:"dbas" api:"required"`
	// Whether this is a personal or business organization
	//
	// Any of "personal", "business".
	Kind string `json:"kind" api:"required"`
	// Legal business name as registered
	LegalBusinessName string `json:"legalBusinessName" api:"required"`
	// Employer Identification Number (EIN), if available
	Ein string `json:"ein" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Dbas              respjson.Field
		Kind              respjson.Field
		LegalBusinessName respjson.Field
		Ein               respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationGetResponseOrganization) RawJSON() string { return r.JSON.raw }
func (r *OrganizationGetResponseOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DBA (Doing Business As) information
type OrganizationGetResponseOrganizationDba struct {
	// Whether this DBA is set as the default for payments
	DbaIsDefault bool `json:"dbaIsDefault" api:"required"`
	// The DBA name
	DbaName string `json:"dbaName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DbaIsDefault respjson.Field
		DbaName      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationGetResponseOrganizationDba) RawJSON() string { return r.JSON.raw }
func (r *OrganizationGetResponseOrganizationDba) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
