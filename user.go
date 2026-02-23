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
func (r *UserService) Get(ctx context.Context, userID string, opts ...option.RequestOption) (res *UserDetails, err error) {
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
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *pagination.CursorIDUsers[UserDetails], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8"), option.WithResponseInto(&raw)}, opts...)
	path := "users"
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

// Get all users
func (r *UserService) ListAutoPaging(ctx context.Context, query UserListParams, opts ...option.RequestOption) *pagination.CursorIDUsersAutoPager[UserDetails] {
	return pagination.NewCursorIDUsersAutoPager(r.List(ctx, query, opts...))
}

// Details of a user within an organization.
type UserDetails struct {
	// User's email address
	Email string `json:"email,required"`
	// User's first name
	FirstName string `json:"firstName,required"`
	// User's last name
	LastName string `json:"lastName,required"`
	// User's role within the organization
	//
	// Any of "administrator", "bookkeeper", "customUser", "cardOnlyUser", "employee".
	OrganizationRole UserDetailsOrganizationRole `json:"organizationRole,required"`
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
func (r UserDetails) RawJSON() string { return r.JSON.raw }
func (r *UserDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User's role within the organization
type UserDetailsOrganizationRole string

const (
	UserDetailsOrganizationRoleAdministrator UserDetailsOrganizationRole = "administrator"
	UserDetailsOrganizationRoleBookkeeper    UserDetailsOrganizationRole = "bookkeeper"
	UserDetailsOrganizationRoleCustomUser    UserDetailsOrganizationRole = "customUser"
	UserDetailsOrganizationRoleCardOnlyUser  UserDetailsOrganizationRole = "cardOnlyUser"
	UserDetailsOrganizationRoleEmployee      UserDetailsOrganizationRole = "employee"
)

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
