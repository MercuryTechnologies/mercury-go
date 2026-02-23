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
func (r *EventService) Get(ctx context.Context, eventID string, opts ...option.RequestOption) (res *APIEvent, err error) {
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
func (r *EventService) List(ctx context.Context, query EventListParams, opts ...option.RequestOption) (res *pagination.CursorIDEvents[APIEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8"), option.WithResponseInto(&raw)}, opts...)
	path := "events"
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

// Get all events
func (r *EventService) ListAutoPaging(ctx context.Context, query EventListParams, opts ...option.RequestOption) *pagination.CursorIDEventsAutoPager[APIEvent] {
	return pagination.NewCursorIDEventsAutoPager(r.List(ctx, query, opts...))
}

// Represents a single event in the Mercury API event stream. | Events track
// changes to resources over time, providing an audit trail | of all modifications
// with before/after values and metadata about what changed.
type APIEvent struct {
	// ID for the API event
	ID string `json:"id,required" format:"uuid"`
	// List of JSON paths that were modified in this event
	ChangedPaths []string `json:"changedPaths,required"`
	// JSON object containing the fields that were changed and their new values
	MergePatch any `json:"mergePatch,required"`
	// Timestamp when the event occurred
	OccurredAt string `json:"occurredAt,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// The type of operation performed (e.g., create, update, delete)
	//
	// Any of "create", "update", "delete".
	OperationType APIEventOperationType `json:"operationType,required"`
	// The ID of the resource that was affected
	ResourceID string `json:"resourceId,required" format:"uuid"`
	// The type of resource that was affected (e.g., transaction, account)
	//
	// Any of "transaction", "checkingAccount", "savingsAccount", "treasuryAccount",
	// "investmentAccount", "creditAccount".
	ResourceType APIEventResourceType `json:"resourceType,required"`
	// Version number of the resource after this change
	ResourceVersion int64 `json:"resourceVersion,required"`
	// JSON object containing the fields that were changed and their previous values
	// before the update
	PreviousValues any `json:"previousValues,nullable"`
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
func (r APIEvent) RawJSON() string { return r.JSON.raw }
func (r *APIEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of operation performed (e.g., create, update, delete)
type APIEventOperationType string

const (
	APIEventOperationTypeCreate APIEventOperationType = "create"
	APIEventOperationTypeUpdate APIEventOperationType = "update"
	APIEventOperationTypeDelete APIEventOperationType = "delete"
)

// The type of resource that was affected (e.g., transaction, account)
type APIEventResourceType string

const (
	APIEventResourceTypeTransaction       APIEventResourceType = "transaction"
	APIEventResourceTypeCheckingAccount   APIEventResourceType = "checkingAccount"
	APIEventResourceTypeSavingsAccount    APIEventResourceType = "savingsAccount"
	APIEventResourceTypeTreasuryAccount   APIEventResourceType = "treasuryAccount"
	APIEventResourceTypeInvestmentAccount APIEventResourceType = "investmentAccount"
	APIEventResourceTypeCreditAccount     APIEventResourceType = "creditAccount"
)

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
