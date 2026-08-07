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

// Manage recipient invites
//
// RecipientInviteService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecipientInviteService] method instead.
type RecipientInviteService struct {
	options []option.RequestOption
}

// NewRecipientInviteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecipientInviteService(opts ...option.RequestOption) (r RecipientInviteService) {
	r = RecipientInviteService{}
	r.options = opts
	return
}

// Create an invite for a recipient to submit their payment details. Supply a
// recipientId to invite an existing recipient; omit it to invite someone new, in
// which case the recipient is created when the invitee completes onboarding.
func (r *RecipientInviteService) New(ctx context.Context, body RecipientInviteNewParams, opts ...option.RequestOption) (res *RecipientInviteAPIResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "recipients/invites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of all recipient invites for your organization.
// Supports filtering by status.
func (r *RecipientInviteService) List(ctx context.Context, query RecipientInviteListParams, opts ...option.RequestOption) (res *RecipientInviteListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "recipients/invites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete an active recipient invite.
func (r *RecipientInviteService) Delete(ctx context.Context, inviteID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if inviteID == "" {
		err = errors.New("missing required inviteId parameter")
		return err
	}
	path := fmt.Sprintf("recipients/invites/%s", url.PathEscape(inviteID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type RecipientInviteAPIResponse struct {
	// ID for the invite
	ID string `json:"id" api:"required"`
	// Recipient contact email the invite was created for.
	ContactEmail string `json:"contactEmail" api:"required"`
	// When the invite was created.
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Recipient name the invite was created for.
	Name string `json:"name" api:"required"`
	// URL where the recipient submits their payment details.
	OnboardingURL string `json:"onboardingUrl" api:"required"`
	// Payment methods the recipient may submit details for.
	PaymentMethods []PaymentMethod `json:"paymentMethods" api:"required"`
	// Whether the recipient must upload a tax document.
	RequireTaxDocument bool `json:"requireTaxDocument" api:"required"`
	// Status of the invite.
	//
	// Any of "created", "completed", "expired".
	Status RecipientInviteAPIResponseStatus `json:"status" api:"required"`
	// When the invite expires, if it has an expiry.
	ExpiresAt string `json:"expiresAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Notes shown to the recipient, if any.
	Notes string `json:"notes" api:"nullable"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		ContactEmail       respjson.Field
		CreatedAt          respjson.Field
		Name               respjson.Field
		OnboardingURL      respjson.Field
		PaymentMethods     respjson.Field
		RequireTaxDocument respjson.Field
		Status             respjson.Field
		ExpiresAt          respjson.Field
		Notes              respjson.Field
		RecipientID        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientInviteAPIResponse) RawJSON() string { return r.JSON.raw }
func (r *RecipientInviteAPIResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the invite.
type RecipientInviteAPIResponseStatus string

const (
	RecipientInviteAPIResponseStatusCreated   RecipientInviteAPIResponseStatus = "created"
	RecipientInviteAPIResponseStatusCompleted RecipientInviteAPIResponseStatus = "completed"
	RecipientInviteAPIResponseStatusExpired   RecipientInviteAPIResponseStatus = "expired"
)

// API response for listing recipient invites with pagination.
type RecipientInviteListResponse struct {
	// List of invites in the current page.
	Invites []RecipientInviteAPIResponse `json:"invites" api:"required"`
	// Pagination cursors (inviteId) for navigating to next/previous pages.
	Page RecipientInviteListResponsePage `json:"page" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invites     respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientInviteListResponse) RawJSON() string { return r.JSON.raw }
func (r *RecipientInviteListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination cursors (inviteId) for navigating to next/previous pages.
type RecipientInviteListResponsePage struct {
	// ID for the invite
	NextPage string `json:"nextPage"`
	// ID for the invite
	PreviousPage string `json:"previousPage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPage     respjson.Field
		PreviousPage respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientInviteListResponsePage) RawJSON() string { return r.JSON.raw }
func (r *RecipientInviteListResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientInviteNewParams struct {
	// Contact email the invite is sent to. When 'recipientId' is present, updates the
	// recipient's contact email to this value.
	ContactEmail string `json:"contactEmail" api:"required"`
	// Payment methods the recipient may submit details for.
	PaymentMethods []PaymentMethod `json:"paymentMethods,omitzero" api:"required"`
	// Whether the recipient must upload a tax document.
	RequireTaxDocument bool `json:"requireTaxDocument" api:"required"`
	// When true, sends an Email to the invitee. When false, does not send an email to
	// the invitee.
	SendEmail bool `json:"sendEmail" api:"required"`
	// Name the invite is created for. This field is required when 'recipientId' is
	// absent. When 'recipientId' is present, this field is optional and updates the
	// recipient's name to this value.
	Name param.Opt[string] `json:"name,omitzero"`
	// Optional notes shown to the recipient.
	Notes param.Opt[string] `json:"notes,omitzero"`
	// Optional organization name to display on the request.
	OrganizationNameOnRequest param.Opt[string] `json:"organizationNameOnRequest,omitzero"`
	// ID for a Mercury account.
	RecipientID param.Opt[string] `json:"recipientId,omitzero" format:"uuid"`
	paramObj
}

func (r RecipientInviteNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RecipientInviteNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientInviteNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientInviteListParams struct {
	// The ID of the recipient invite to end the page before (exclusive). When
	// provided, results will end just before this ID and work backwards. Use this for
	// reverse pagination or to retrieve previous pages. Cannot be combined with
	// start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the recipient invite to start the page after (exclusive). When
	// provided, results will begin with the recipient invite immediately following
	// this ID. Use this for standard forward pagination to get the next page of
	// results. Cannot be combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order RecipientInviteListParamsOrder `query:"order,omitzero" json:"-"`
	// Any of "created", "completed", "expired".
	Status RecipientInviteListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecipientInviteListParams]'s query parameters as
// `url.Values`.
func (r RecipientInviteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type RecipientInviteListParamsOrder string

const (
	RecipientInviteListParamsOrderAsc  RecipientInviteListParamsOrder = "asc"
	RecipientInviteListParamsOrderDesc RecipientInviteListParamsOrder = "desc"
)

type RecipientInviteListParamsStatus string

const (
	RecipientInviteListParamsStatusCreated   RecipientInviteListParamsStatus = "created"
	RecipientInviteListParamsStatusCompleted RecipientInviteListParamsStatus = "completed"
	RecipientInviteListParamsStatusExpired   RecipientInviteListParamsStatus = "expired"
)
