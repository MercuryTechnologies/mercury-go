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
	"github.com/MercuryTechnologies/mercury-go/packages/pagination"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// WebhookService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

// Register a new webhook endpoint to receive event notifications
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve details of a specific webhook endpoint by ID
func (r *WebhookService) Get(ctx context.Context, webhookEndpointID string, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if webhookEndpointID == "" {
		err = errors.New("missing required webhookEndpointId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the configuration of an existing webhook endpoint. A webhook that has
// been disabled due to consecutive delivery failures can be reactivated by setting
// its status to 'active'.
func (r *WebhookService) Update(ctx context.Context, webhookEndpointID string, body WebhookUpdateParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if webhookEndpointID == "" {
		err = errors.New("missing required webhookEndpointId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of all webhook endpoints for your organization.
// Supports filtering by status.
func (r *WebhookService) List(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) (res *pagination.CursorIDWebhooks[Webhook], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8"), option.WithResponseInto(&raw)}, opts...)
	path := "webhooks"
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

// Retrieve a paginated list of all webhook endpoints for your organization.
// Supports filtering by status.
func (r *WebhookService) ListAutoPaging(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) *pagination.CursorIDWebhooksAutoPager[Webhook] {
	return pagination.NewCursorIDWebhooksAutoPager(r.List(ctx, query, opts...))
}

// Delete a webhook endpoint
func (r *WebhookService) Delete(ctx context.Context, webhookEndpointID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookEndpointID == "" {
		err = errors.New("missing required webhookEndpointId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s", webhookEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Send a test event to verify the webhook endpoint is properly configured and
// reachable. The request body accepts an optional 'eventType' field to specify
// which event type to test (e.g., 'transaction.created', 'transaction.updated').
// If omitted from the request body, defaults to 'transaction.created'.
func (r *WebhookService) Verify(ctx context.Context, webhookEndpointID string, body WebhookVerifyParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if webhookEndpointID == "" {
		err = errors.New("missing required webhookEndpointId parameter")
		return
	}
	path := fmt.Sprintf("webhooks/%s/verify", webhookEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Webhook configuration details
type Webhook struct {
	// ID for the webhook
	ID string `json:"id" api:"required" format:"uuid"`
	// When the webhook was created
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// The status of the webhook endpoint. 'active': delivering events normally.
	// 'paused': paused by the user. 'disabled': automatically disabled by the system
	// due to consecutive delivery failures. A disabled webhook can be reactivated by
	// updating its status to 'active'.
	//
	// Any of "active", "paused", "disabled".
	Status WebhookStatus `json:"status" api:"required"`
	// When the webhook was last updated
	UpdatedAt string `json:"updatedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// The URL that will receive webhook POST requests
	URL string `json:"url" api:"required"`
	// Optional array of event types this webhook is subscribed to. Nothing means all
	// events.
	//
	// Any of "transaction.created", "transaction.updated",
	// "checkingAccount.balance.updated", "savingsAccount.balance.updated",
	// "treasuryAccount.balance.updated", "investmentAccount.balance.updated",
	// "creditAccount.balance.updated".
	EventTypes []string `json:"eventTypes" api:"nullable"`
	// Optional array of resource field paths to filter events by. Nothing means no
	// filtering.
	//
	// Any of "transaction.amount", "transaction.bankDescription",
	// "transaction.categoryData", "transaction.customCategory",
	// "transaction.customCategory.id", "transaction.customCategory.name",
	// "transaction.mercuryCategory", "transaction.estimatedDeliveryDate",
	// "transaction.externalMemo", "transaction.failedAt", "transaction.note",
	// "transaction.postedAt", "transaction.reasonForFailure", "transaction.status",
	// "checkingAccount.availableBalance", "checkingAccount.currentBalance",
	// "checkingAccount.inFlightBalance", "savingsAccount.availableBalance",
	// "savingsAccount.currentBalance", "savingsAccount.inFlightBalance",
	// "treasuryAccount.availableBalance", "treasuryAccount.currentBalance",
	// "treasuryAccount.inFlightBalance", "investmentAccount.availableBalance",
	// "investmentAccount.currentBalance", "investmentAccount.inFlightBalance",
	// "creditAccount.availableBalance", "creditAccount.currentBalance",
	// "creditAccount.inFlightBalance".
	FilterPaths []string `json:"filterPaths" api:"nullable"`
	// Webhook signing secret. Only returned on creation (POST), not on GET or UPDATE
	// operations.
	Secret string `json:"secret" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Status      respjson.Field
		UpdatedAt   respjson.Field
		URL         respjson.Field
		EventTypes  respjson.Field
		FilterPaths respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Webhook) RawJSON() string { return r.JSON.raw }
func (r *Webhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the webhook endpoint. 'active': delivering events normally.
// 'paused': paused by the user. 'disabled': automatically disabled by the system
// due to consecutive delivery failures. A disabled webhook can be reactivated by
// updating its status to 'active'.
type WebhookStatus string

const (
	WebhookStatusActive   WebhookStatus = "active"
	WebhookStatusPaused   WebhookStatus = "paused"
	WebhookStatusDisabled WebhookStatus = "disabled"
)

type WebhookNewParams struct {
	// The URL to which webhook events will be delivered
	URL string `json:"url" api:"required"`
	// Optional array of event types to subscribe to. Nothing means subscribe to all
	// event types.
	//
	// Any of "transaction.created", "transaction.updated",
	// "checkingAccount.balance.updated", "savingsAccount.balance.updated",
	// "treasuryAccount.balance.updated", "investmentAccount.balance.updated",
	// "creditAccount.balance.updated".
	EventTypes []string `json:"eventTypes,omitzero"`
	// Optional array of resource field paths to filter events by. When specified,
	// webhook events will only be sent when one of these fields changes. Nothing means
	// no filtering (all events are sent).
	//
	// Any of "transaction.amount", "transaction.bankDescription",
	// "transaction.categoryData", "transaction.customCategory",
	// "transaction.customCategory.id", "transaction.customCategory.name",
	// "transaction.mercuryCategory", "transaction.estimatedDeliveryDate",
	// "transaction.externalMemo", "transaction.failedAt", "transaction.note",
	// "transaction.postedAt", "transaction.reasonForFailure", "transaction.status",
	// "checkingAccount.availableBalance", "checkingAccount.currentBalance",
	// "checkingAccount.inFlightBalance", "savingsAccount.availableBalance",
	// "savingsAccount.currentBalance", "savingsAccount.inFlightBalance",
	// "treasuryAccount.availableBalance", "treasuryAccount.currentBalance",
	// "treasuryAccount.inFlightBalance", "investmentAccount.availableBalance",
	// "investmentAccount.currentBalance", "investmentAccount.inFlightBalance",
	// "creditAccount.availableBalance", "creditAccount.currentBalance",
	// "creditAccount.inFlightBalance".
	FilterPaths []string `json:"filterPaths,omitzero"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	// The URL to which webhook events will be delivered. Omit to leave unchanged.
	URL param.Opt[string] `json:"url,omitzero"`
	// Event types to subscribe to. Send null to subscribe to all event types. Send an
	// array to subscribe to specific types. Omit to leave unchanged.
	//
	// Any of "transaction.created", "transaction.updated",
	// "checkingAccount.balance.updated", "savingsAccount.balance.updated",
	// "treasuryAccount.balance.updated", "investmentAccount.balance.updated",
	// "creditAccount.balance.updated".
	EventTypes []string `json:"eventTypes,omitzero"`
	// Resource field paths to filter events by. When specified, webhook events will
	// only be sent when one of these fields changes. Send null for no filtering. Send
	// an array to filter by specific fields. Omit to leave unchanged.
	//
	// Any of "transaction.amount", "transaction.bankDescription",
	// "transaction.categoryData", "transaction.customCategory",
	// "transaction.customCategory.id", "transaction.customCategory.name",
	// "transaction.mercuryCategory", "transaction.estimatedDeliveryDate",
	// "transaction.externalMemo", "transaction.failedAt", "transaction.note",
	// "transaction.postedAt", "transaction.reasonForFailure", "transaction.status",
	// "checkingAccount.availableBalance", "checkingAccount.currentBalance",
	// "checkingAccount.inFlightBalance", "savingsAccount.availableBalance",
	// "savingsAccount.currentBalance", "savingsAccount.inFlightBalance",
	// "treasuryAccount.availableBalance", "treasuryAccount.currentBalance",
	// "treasuryAccount.inFlightBalance", "investmentAccount.availableBalance",
	// "investmentAccount.currentBalance", "investmentAccount.inFlightBalance",
	// "creditAccount.availableBalance", "creditAccount.currentBalance",
	// "creditAccount.inFlightBalance".
	FilterPaths []string `json:"filterPaths,omitzero"`
	// Webhook status. Only 'active' and 'paused' values are allowed. Omit to leave
	// unchanged.
	//
	// Any of "active", "paused".
	Status WebhookUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook status. Only 'active' and 'paused' values are allowed. Omit to leave
// unchanged.
type WebhookUpdateParamsStatus string

const (
	WebhookUpdateParamsStatusActive WebhookUpdateParamsStatus = "active"
	WebhookUpdateParamsStatusPaused WebhookUpdateParamsStatus = "paused"
)

type WebhookListParams struct {
	// The ID of the webhook to end the page before (exclusive). When provided, results
	// will end just before this ID and work backwards. Use this for reverse pagination
	// or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the webhook to start the page after (exclusive). When provided,
	// results will begin with the webhook immediately following this ID. Use this for
	// standard forward pagination to get the next page of results. Cannot be combined
	// with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order WebhookListParamsOrder `query:"order,omitzero" json:"-"`
	// Any of "active", "paused", "disabled", "deleted".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type WebhookListParamsOrder string

const (
	WebhookListParamsOrderAsc  WebhookListParamsOrder = "asc"
	WebhookListParamsOrderDesc WebhookListParamsOrder = "desc"
)

type WebhookVerifyParams struct {
	// Optional event type to test. If not specified, defaults to transaction.created.
	//
	// Any of "transaction.created", "transaction.updated",
	// "checkingAccount.balance.updated", "savingsAccount.balance.updated",
	// "treasuryAccount.balance.updated", "investmentAccount.balance.updated",
	// "creditAccount.balance.updated".
	EventType WebhookVerifyParamsEventType `json:"eventType,omitzero"`
	paramObj
}

func (r WebhookVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional event type to test. If not specified, defaults to transaction.created.
type WebhookVerifyParamsEventType string

const (
	WebhookVerifyParamsEventTypeTransactionCreated              WebhookVerifyParamsEventType = "transaction.created"
	WebhookVerifyParamsEventTypeTransactionUpdated              WebhookVerifyParamsEventType = "transaction.updated"
	WebhookVerifyParamsEventTypeCheckingAccountBalanceUpdated   WebhookVerifyParamsEventType = "checkingAccount.balance.updated"
	WebhookVerifyParamsEventTypeSavingsAccountBalanceUpdated    WebhookVerifyParamsEventType = "savingsAccount.balance.updated"
	WebhookVerifyParamsEventTypeTreasuryAccountBalanceUpdated   WebhookVerifyParamsEventType = "treasuryAccount.balance.updated"
	WebhookVerifyParamsEventTypeInvestmentAccountBalanceUpdated WebhookVerifyParamsEventType = "investmentAccount.balance.updated"
	WebhookVerifyParamsEventTypeCreditAccountBalanceUpdated     WebhookVerifyParamsEventType = "creditAccount.balance.updated"
)
