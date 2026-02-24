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

// UserService contains methods and other services that help with interacting with
// the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// Get user by ID
func (r *UserService) Get(ctx context.Context, userID string, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return
	}
	path := fmt.Sprintf("users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all users
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Details of a user within an organization.
type User struct {
	// User's email address
	Email string `json:"email,required"`
	// User's first name
	FirstName string `json:"firstName,required"`
	// User's last name
	LastName string `json:"lastName,required"`
	// User's role within the organization
	//
	// Any of "administrator", "bookkeeper", "customUser", "cardOnlyUser", "employee".
	OrganizationRole UserOrganizationRole `json:"organizationRole,required"`
	// ID for the user
	UserID string `json:"userId,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email            respjson.Field
		FirstName        respjson.Field
		LastName         respjson.Field
		OrganizationRole respjson.Field
		UserID           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User's role within the organization
type UserOrganizationRole string

const (
	UserOrganizationRoleAdministrator UserOrganizationRole = "administrator"
	UserOrganizationRoleBookkeeper    UserOrganizationRole = "bookkeeper"
	UserOrganizationRoleCustomUser    UserOrganizationRole = "customUser"
	UserOrganizationRoleCardOnlyUser  UserOrganizationRole = "cardOnlyUser"
	UserOrganizationRoleEmployee      UserOrganizationRole = "employee"
)

// Paginated response containing a list of organization users.
type UserListResponse struct {
	// Pagination information including cursors for navigating to next/previous pages
	Page UserListResponsePage `json:"page,required"`
	// List of users in the current page
	Users []User `json:"users,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination information including cursors for navigating to next/previous pages
type UserListResponsePage struct {
	// ID for the user
	NextPage string `json:"nextPage" format:"uuid"`
	// ID for the user
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
func (r UserListResponsePage) RawJSON() string { return r.JSON.raw }
func (r *UserListResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListParams struct {
	// The ID of the user to end the page before (exclusive). When provided, results
	// will end just before this ID and work backwards. Use this for reverse pagination
	// or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the user to start the page after (exclusive). When provided, results
	// will begin with the user immediately following this ID. Use this for standard
	// forward pagination to get the next page of results. Cannot be combined with
	// end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order UserListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type UserListParamsOrder string

const (
	UserListParamsOrderAsc  UserListParamsOrder = "asc"
	UserListParamsOrderDesc UserListParamsOrder = "desc"
)
