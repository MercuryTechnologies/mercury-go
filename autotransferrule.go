// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"encoding/json"
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

// Retrieve auto-transfer rules
//
// AutoTransferRuleService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutoTransferRuleService] method instead.
type AutoTransferRuleService struct {
	options []option.RequestOption
}

// NewAutoTransferRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAutoTransferRuleService(opts ...option.RequestOption) (r AutoTransferRuleService) {
	r = AutoTransferRuleService{}
	r.options = opts
	return
}

// Retrieve a paginated list of the organization's active auto-transfer rules:
// rules that distribute fixed amounts or percentages of an account's balance to
// other accounts, and rules that keep an account at a target balance. Pass
// accountId to return only the rules that manage that account's balance.
func (r *AutoTransferRuleService) GetAutoTransferRules(ctx context.Context, query AutoTransferRuleGetAutoTransferRulesParams, opts ...option.RequestOption) (res *AutoTransferRuleGetAutoTransferRulesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "auto-transfer-rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A page of the organization's active auto-transfer rules.
type AutoTransferRuleGetAutoTransferRulesResponse struct {
	// Cursors for navigating to the next and previous pages.
	Page AutoTransferRuleGetAutoTransferRulesResponsePage `json:"page" api:"required"`
	// Rules in the current page.
	Rules []AutoTransferRuleGetAutoTransferRulesResponseRule `json:"rules" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		Rules       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoTransferRuleGetAutoTransferRulesResponse) RawJSON() string { return r.JSON.raw }
func (r *AutoTransferRuleGetAutoTransferRulesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cursors for navigating to the next and previous pages.
type AutoTransferRuleGetAutoTransferRulesResponsePage struct {
	NextPage     string `json:"nextPage" format:"uuid"`
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
func (r AutoTransferRuleGetAutoTransferRulesResponsePage) RawJSON() string { return r.JSON.raw }
func (r *AutoTransferRuleGetAutoTransferRulesResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An active auto-transfer rule.
type AutoTransferRuleGetAutoTransferRulesResponseRule struct {
	// The rule's ID. Editing a rule replaces it with a new version that has a new ID.
	ID string `json:"id" api:"required" format:"uuid"`
	// ID for a Mercury account.
	ControlledAccountID string `json:"controlledAccountId" api:"required" format:"uuid"`
	// A sentence describing what the rule does.
	Description string `json:"description" api:"required"`
	// The rule's kind and its kind-specific details.
	Details AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion `json:"details" api:"required"`
	// How often the rule runs, e.g. "Every business day".
	Frequency string `json:"frequency" api:"required"`
	// Whether the rule is paused.
	IsPaused bool `json:"isPaused" api:"required"`
	// The rule's name, as shown in the Mercury dashboard.
	Name string `json:"name" api:"required"`
	// When the rule next runs, if scheduled.
	NextTransferAt string `json:"nextTransferAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// A nickname the user gave the rule, if any.
	Nickname string `json:"nickname" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		ControlledAccountID respjson.Field
		Description         respjson.Field
		Details             respjson.Field
		Frequency           respjson.Field
		IsPaused            respjson.Field
		Name                respjson.Field
		NextTransferAt      respjson.Field
		Nickname            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoTransferRuleGetAutoTransferRulesResponseRule) RawJSON() string { return r.JSON.raw }
func (r *AutoTransferRuleGetAutoTransferRulesResponseRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion contains all
// possible properties and values from
// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmount],
// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentage],
// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalance].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion struct {
	// This field is a union of
	// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContents],
	// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContents],
	// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalanceContents]
	Contents AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnionContents `json:"contents"`
	Kind     string                                                               `json:"kind"`
	JSON     struct {
		Contents respjson.Field
		Kind     respjson.Field
		raw      string
	} `json:"-"`
}

func (u AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion) AsAutoTransferRuleKindDistributeAmount() (v AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmount) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion) AsAutoTransferRuleKindDistributePercentage() (v AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion) AsAutoTransferRuleKindMaintainTargetBalance() (v AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnionContents is an
// implicit subunion of
// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion].
// AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnionContents provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion].
type AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnionContents struct {
	// This field is a union of
	// [[]AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContentsReservoirAccount],
	// [[]AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContentsReservoirAccount]
	ReservoirAccounts AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnionContentsReservoirAccounts `json:"reservoirAccounts"`
	// This field is from variant
	// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalanceContents].
	ReservoirAccountID string `json:"reservoirAccountId"`
	// This field is from variant
	// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalanceContents].
	TargetBalance float64 `json:"targetBalance"`
	// This field is from variant
	// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalanceContents].
	LowerBound float64 `json:"lowerBound"`
	// This field is from variant
	// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalanceContents].
	UpperBound float64 `json:"upperBound"`
	JSON       struct {
		ReservoirAccounts  respjson.Field
		ReservoirAccountID respjson.Field
		TargetBalance      respjson.Field
		LowerBound         respjson.Field
		UpperBound         respjson.Field
		raw                string
	} `json:"-"`
}

func (r *AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnionContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnionContentsReservoirAccounts
// is an implicit subunion of
// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion].
// AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnionContentsReservoirAccounts
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfAutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContentsReservoirAccounts
// OfAutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContentsReservoirAccounts]
type AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnionContentsReservoirAccounts struct {
	// This field will be present if the value is a
	// [[]AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContentsReservoirAccount]
	// instead of an object.
	OfAutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContentsReservoirAccounts []AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContentsReservoirAccount `json:",inline"`
	// This field will be present if the value is a
	// [[]AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContentsReservoirAccount]
	// instead of an object.
	OfAutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContentsReservoirAccounts []AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContentsReservoirAccount `json:",inline"`
	JSON                                                                                                                       struct {
		OfAutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContentsReservoirAccounts     respjson.Field
		OfAutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContentsReservoirAccounts respjson.Field
		raw                                                                                                                        string
	} `json:"-"`
}

func (r *AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsUnionContentsReservoirAccounts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmount struct {
	// Distributes fixed amounts from the controlled account to one or more reservoir
	// accounts.
	Contents AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContents `json:"contents" api:"required"`
	// Any of "distributeAmount".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmount) RawJSON() string {
	return r.JSON.raw
}
func (r *AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Distributes fixed amounts from the controlled account to one or more reservoir
// accounts.
type AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContents struct {
	// The reservoir accounts and the amount each receives.
	ReservoirAccounts []AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContentsReservoirAccount `json:"reservoirAccounts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReservoirAccounts respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContents) RawJSON() string {
	return r.JSON.raw
}
func (r *AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A reservoir account and the fixed amount a distribute rule moves to it.
type AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContentsReservoirAccount struct {
	// ID for a Mercury account.
	AccountID string `json:"accountId" api:"required" format:"uuid"`
	// A dollar amount
	Amount float64 `json:"amount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		Amount      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContentsReservoirAccount) RawJSON() string {
	return r.JSON.raw
}
func (r *AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributeAmountContentsReservoirAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentage struct {
	// Distributes a percentage of the controlled account's balance to one or more
	// reservoir accounts.
	Contents AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContents `json:"contents" api:"required"`
	// Any of "distributePercentage".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentage) RawJSON() string {
	return r.JSON.raw
}
func (r *AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Distributes a percentage of the controlled account's balance to one or more
// reservoir accounts.
type AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContents struct {
	// The reservoir accounts and the percentage each receives.
	ReservoirAccounts []AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContentsReservoirAccount `json:"reservoirAccounts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReservoirAccounts respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContents) RawJSON() string {
	return r.JSON.raw
}
func (r *AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A reservoir account and the share of the controlled account's balance a
// distribute rule moves to it.
type AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContentsReservoirAccount struct {
	// ID for a Mercury account.
	AccountID string `json:"accountId" api:"required" format:"uuid"`
	// The percentage, 0 through 100, of the controlled account's balance moved to this
	// account.
	Percentage float64 `json:"percentage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID   respjson.Field
		Percentage  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContentsReservoirAccount) RawJSON() string {
	return r.JSON.raw
}
func (r *AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindDistributePercentageContentsReservoirAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalance struct {
	// Keeps the controlled account's balance near a target by moving money between it
	// and a reservoir account. When the balance moves outside the bounds, the rule
	// moves exactly enough money to bring it back to targetBalance. A rule with both
	// bounds equal to targetBalance moves money whenever the balance differs from the
	// target.
	Contents AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalanceContents `json:"contents" api:"required"`
	// Any of "maintainTargetBalance".
	Kind string `json:"kind" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalance) RawJSON() string {
	return r.JSON.raw
}
func (r *AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Keeps the controlled account's balance near a target by moving money between it
// and a reservoir account. When the balance moves outside the bounds, the rule
// moves exactly enough money to bring it back to targetBalance. A rule with both
// bounds equal to targetBalance moves money whenever the balance differs from the
// target.
type AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalanceContents struct {
	// ID for a Mercury account.
	ReservoirAccountID string `json:"reservoirAccountId" api:"required" format:"uuid"`
	// A dollar amount
	TargetBalance float64 `json:"targetBalance" api:"required"`
	// A dollar amount
	LowerBound float64 `json:"lowerBound" api:"nullable"`
	// A dollar amount
	UpperBound float64 `json:"upperBound" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReservoirAccountID respjson.Field
		TargetBalance      respjson.Field
		LowerBound         respjson.Field
		UpperBound         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalanceContents) RawJSON() string {
	return r.JSON.raw
}
func (r *AutoTransferRuleGetAutoTransferRulesResponseRuleDetailsAutoTransferRuleKindMaintainTargetBalanceContents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutoTransferRuleGetAutoTransferRulesParams struct {
	// ID for a Mercury account.
	AccountID param.Opt[string] `query:"accountId,omitzero" format:"uuid" json:"-"`
	// The ID of the auto-transfer rule to end the page before (exclusive). When
	// provided, results will end just before this ID and work backwards. Use this for
	// reverse pagination or to retrieve previous pages. Cannot be combined with
	// start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the auto-transfer rule to start the page after (exclusive). When
	// provided, results will begin with the auto-transfer rule immediately following
	// this ID. Use this for standard forward pagination to get the next page of
	// results. Cannot be combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order AutoTransferRuleGetAutoTransferRulesParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AutoTransferRuleGetAutoTransferRulesParams]'s query
// parameters as `url.Values`.
func (r AutoTransferRuleGetAutoTransferRulesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type AutoTransferRuleGetAutoTransferRulesParamsOrder string

const (
	AutoTransferRuleGetAutoTransferRulesParamsOrderAsc  AutoTransferRuleGetAutoTransferRulesParamsOrder = "asc"
	AutoTransferRuleGetAutoTransferRulesParamsOrderDesc AutoTransferRuleGetAutoTransferRulesParamsOrder = "desc"
)
