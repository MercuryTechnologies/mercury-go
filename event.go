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

// EventService contains methods and other services that help with interacting with
// the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventService] method instead.
type EventService struct {
	Options []option.RequestOption
}

// NewEventService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEventService(opts ...option.RequestOption) (r EventService) {
	r = EventService{}
	r.Options = opts
	return
}

// Get event by ID
func (r *EventService) Get(ctx context.Context, eventID string, opts ...option.RequestOption) (res *Event, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if eventID == "" {
		err = errors.New("missing required eventId parameter")
		return
	}
	path := fmt.Sprintf("events/%s", eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all events
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *EventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Represents a single event in the Mercury API event stream. | Events track
// changes to resources over time, providing an audit trail | of all modifications
// with before/after values and metadata about what changed.
type Event struct {
	// ID for the API event
	ID string `json:"id" api:"required" format:"uuid"`
	// List of JSON paths that were modified in this event
	ChangedPaths []string `json:"changedPaths" api:"required"`
	// JSON object containing the fields that were changed and their new values
	MergePatch any `json:"mergePatch" api:"required"`
	// Timestamp when the event occurred
	OccurredAt string `json:"occurredAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// The type of operation performed (e.g., create, update, delete)
	//
	// Any of "create", "update", "delete".
	OperationType EventOperationType `json:"operationType" api:"required"`
	// The ID of the resource that was affected
	ResourceID string `json:"resourceId" api:"required" format:"uuid"`
	// The type of resource that was affected (e.g., transaction, account)
	//
	// Any of "transaction", "checkingAccount", "savingsAccount", "treasuryAccount",
	// "investmentAccount", "creditAccount".
	ResourceType EventResourceType `json:"resourceType" api:"required"`
	// Version number of the resource after this change
	ResourceVersion int64 `json:"resourceVersion" api:"required"`
	// JSON object containing the fields that were changed and their previous values
	// before the update
	PreviousValues any `json:"previousValues" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ChangedPaths    respjson.Field
		MergePatch      respjson.Field
		OccurredAt      respjson.Field
		OperationType   respjson.Field
		ResourceID      respjson.Field
		ResourceType    respjson.Field
		ResourceVersion respjson.Field
		PreviousValues  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Event) RawJSON() string { return r.JSON.raw }
func (r *Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of operation performed (e.g., create, update, delete)
type EventOperationType string

const (
	EventOperationTypeCreate EventOperationType = "create"
	EventOperationTypeUpdate EventOperationType = "update"
	EventOperationTypeDelete EventOperationType = "delete"
)

// The type of resource that was affected (e.g., transaction, account)
type EventResourceType string

const (
	EventResourceTypeTransaction       EventResourceType = "transaction"
	EventResourceTypeCheckingAccount   EventResourceType = "checkingAccount"
	EventResourceTypeSavingsAccount    EventResourceType = "savingsAccount"
	EventResourceTypeTreasuryAccount   EventResourceType = "treasuryAccount"
	EventResourceTypeInvestmentAccount EventResourceType = "investmentAccount"
	EventResourceTypeCreditAccount     EventResourceType = "creditAccount"
)

// Paginated response containing a list of API events. | Use the page cursor
// information to fetch additional pages of events.
type EventListResponse struct {
	// List of events in the current page
	Events []Event `json:"events" api:"required"`
	// Pagination information including cursors for navigating to next/previous pages
	Page EventListResponsePage `json:"page" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events      respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EventListResponse) RawJSON() string { return r.JSON.raw }
func (r *EventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination information including cursors for navigating to next/previous pages
type EventListResponsePage struct {
	// ID for the API event
	NextPage string `json:"nextPage" format:"uuid"`
	// ID for the API event
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
func (r EventListResponsePage) RawJSON() string { return r.JSON.raw }
func (r *EventListResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventListParams struct {
	// The ID of the event to end the page before (exclusive). When provided, results
	// will end just before this ID and work backwards. Use this for reverse pagination
	// or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit      param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	ResourceID param.Opt[string] `query:"resourceId,omitzero" format:"uuid" json:"-"`
	// The ID of the event to start the page after (exclusive). When provided, results
	// will begin with the event immediately following this ID. Use this for standard
	// forward pagination to get the next page of results. Cannot be combined with
	// end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order EventListParamsOrder `query:"order,omitzero" json:"-"`
	// Any of "transaction", "checkingAccount", "savingsAccount", "treasuryAccount",
	// "investmentAccount", "creditAccount".
	ResourceType EventListParamsResourceType `query:"resourceType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EventListParams]'s query parameters as `url.Values`.
func (r EventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type EventListParamsOrder string

const (
	EventListParamsOrderAsc  EventListParamsOrder = "asc"
	EventListParamsOrderDesc EventListParamsOrder = "desc"
)

type EventListParamsResourceType string

const (
	EventListParamsResourceTypeTransaction       EventListParamsResourceType = "transaction"
	EventListParamsResourceTypeCheckingAccount   EventListParamsResourceType = "checkingAccount"
	EventListParamsResourceTypeSavingsAccount    EventListParamsResourceType = "savingsAccount"
	EventListParamsResourceTypeTreasuryAccount   EventListParamsResourceType = "treasuryAccount"
	EventListParamsResourceTypeInvestmentAccount EventListParamsResourceType = "investmentAccount"
	EventListParamsResourceTypeCreditAccount     EventListParamsResourceType = "creditAccount"
)
