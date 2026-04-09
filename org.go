// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"net/http"
	"slices"

	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// Organization information
//
// OrgService contains methods and other services that help with interacting with
// the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgService] method instead.
type OrgService struct {
	Options []option.RequestOption
}

// NewOrgService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgService(opts ...option.RequestOption) (r OrgService) {
	r = OrgService{}
	r.Options = opts
	return
}

// Retrieve information about your organization including EIN, legal business name,
// and DBAs.
func (r *OrgService) Get(ctx context.Context, opts ...option.RequestOption) (res *OrgGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "organization"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Response containing organization details.
type OrgGetResponse struct {
	// Organization information
	Organization OrgGetResponseOrganization `json:"organization" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Organization respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrgGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OrgGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organization information
type OrgGetResponseOrganization struct {
	// Unique identifier for the organization
	ID string `json:"id" api:"required" format:"uuid"`
	// List of DBAs (Doing Business As names) for this organization
	Dbas []OrgGetResponseOrganizationDba `json:"dbas" api:"required"`
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
func (r OrgGetResponseOrganization) RawJSON() string { return r.JSON.raw }
func (r *OrgGetResponseOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DBA (Doing Business As) information
type OrgGetResponseOrganizationDba struct {
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
func (r OrgGetResponseOrganizationDba) RawJSON() string { return r.JSON.raw }
func (r *OrgGetResponseOrganizationDba) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
