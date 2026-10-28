// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	"github.com/MercuryTechnologies/mercury-go/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/pagination"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// Manage Spend Management budgets
//
// SpendBudgetService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpendBudgetService] method instead.
type SpendBudgetService struct {
	options []option.RequestOption
}

// NewSpendBudgetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpendBudgetService(opts ...option.RequestOption) (r SpendBudgetService) {
	r = SpendBudgetService{}
	r.options = opts
	return
}

// Create a budget as a draft. Add members and publish it from the Mercury
// dashboard.
func (r *SpendBudgetService) New(ctx context.Context, params SpendBudgetNewParams, opts ...option.RequestOption) (res *BudgetIDResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "spend/budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a budget's fields. A draft is edited in place; a published budget gets a
// new version. Omitted or null fields keep their current value. Frequency cannot
// change after publication. Effective changes to amountPerUserPerPeriod,
// frequency, startDate, or endDate return 409 when the budget has an active
// Agentic Card link; name and description updates remain allowed.
func (r *SpendBudgetService) Update(ctx context.Context, budgetID string, params SpendBudgetUpdateParams, opts ...option.RequestOption) (res *BudgetIDResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	if budgetID == "" {
		err = errors.New("missing required budgetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("spend/budgets/%s", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve the organization's non-deleted budgets ordered by creation time,
// ascending (oldest first) by default. Supports cursor-based pagination with
// limit, order, start_at, start_after, and end_before query parameters.
func (r *SpendBudgetService) List(ctx context.Context, query SpendBudgetListParams, opts ...option.RequestOption) (res *pagination.CursorIDBudgets[Budget], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "spend/budgets"
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

// Retrieve the organization's non-deleted budgets ordered by creation time,
// ascending (oldest first) by default. Supports cursor-based pagination with
// limit, order, start_at, start_after, and end_before query parameters.
func (r *SpendBudgetService) ListAutoPaging(ctx context.Context, query SpendBudgetListParams, opts ...option.RequestOption) *pagination.CursorIDBudgetsAutoPager[Budget] {
	return pagination.NewCursorIDBudgetsAutoPager(r.List(ctx, query, opts...))
}

// Archive a budget (a published budget becomes archived; a draft becomes deleted).
// Returns 409 when the budget has an active Agentic Card link.
func (r *SpendBudgetService) Archive(ctx context.Context, budgetID string, body SpendBudgetArchiveParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", body.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if budgetID == "" {
		err = errors.New("missing required budgetId parameter")
		return err
	}
	path := fmt.Sprintf("spend/budgets/%s/archive", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Assign a transaction to a budget so its spend counts against the cycle
// containing the transaction's date. If the assignment would overspend the budget,
// set allowOverspend to true; omitted and false values enforce the spend limit.
func (r *SpendBudgetService) AssignTransaction(ctx context.Context, budgetID string, params SpendBudgetAssignTransactionParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if budgetID == "" {
		err = errors.New("missing required budgetId parameter")
		return err
	}
	path := fmt.Sprintf("spend/budgets/%s/assign-transaction", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Retrieve a budget by id, including its lifecycle status, per-user limit,
// cadence, and effective window.
func (r *SpendBudgetService) Get(ctx context.Context, budgetID string, opts ...option.RequestOption) (res *Budget, err error) {
	opts = slices.Concat(r.options, opts)
	if budgetID == "" {
		err = errors.New("missing required budgetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("spend/budgets/%s", budgetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List the non-deleted expenditures assigned to a budget, ordered by expenditure
// creation time, ascending (oldest first) by default. Supports cursor-based
// pagination with limit, order, start_at, start_after, and end_before query
// parameters.
func (r *SpendBudgetService) ListExpenditures(ctx context.Context, budgetID string, query SpendBudgetListExpendituresParams, opts ...option.RequestOption) (res *pagination.CursorIDBudgetExpenditures[BudgetExpenditure], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if budgetID == "" {
		err = errors.New("missing required budgetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("spend/budgets/%s/expenditures", budgetID)
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

// List the non-deleted expenditures assigned to a budget, ordered by expenditure
// creation time, ascending (oldest first) by default. Supports cursor-based
// pagination with limit, order, start_at, start_after, and end_before query
// parameters.
func (r *SpendBudgetService) ListExpendituresAutoPaging(ctx context.Context, budgetID string, query SpendBudgetListExpendituresParams, opts ...option.RequestOption) *pagination.CursorIDBudgetExpendituresAutoPager[BudgetExpenditure] {
	return pagination.NewCursorIDBudgetExpendituresAutoPager(r.ListExpenditures(ctx, budgetID, query, opts...))
}

// List the members of a budget and their current or final cycle spend, ordered by
// the creation time of each budget member's assignment record, ascending (oldest
// first) by default. Supports cursor-based pagination with limit, order, start_at,
// start_after, and end_before query parameters.
func (r *SpendBudgetService) ListMembers(ctx context.Context, budgetID string, query SpendBudgetListMembersParams, opts ...option.RequestOption) (res *pagination.CursorIDBudgetMembers[BudgetMember], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if budgetID == "" {
		err = errors.New("missing required budgetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("spend/budgets/%s/members", budgetID)
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

// List the members of a budget and their current or final cycle spend, ordered by
// the creation time of each budget member's assignment record, ascending (oldest
// first) by default. Supports cursor-based pagination with limit, order, start_at,
// start_after, and end_before query parameters.
func (r *SpendBudgetService) ListMembersAutoPaging(ctx context.Context, budgetID string, query SpendBudgetListMembersParams, opts ...option.RequestOption) *pagination.CursorIDBudgetMembersAutoPager[BudgetMember] {
	return pagination.NewCursorIDBudgetMembersAutoPager(r.ListMembers(ctx, budgetID, query, opts...))
}

type Budget struct {
	// ID for the budget
	ID string `json:"id" api:"required" format:"uuid"`
	// Any of "oneTime", "daily", "weekly", "monthly", "yearly".
	Frequency BudgetFrequency `json:"frequency" api:"required"`
	// Any of "draft", "published", "archived".
	Status BudgetPublishStatus `json:"status" api:"required"`
	// A non-negative dollar amount
	AmountPerUserPerPeriod float64   `json:"amountPerUserPerPeriod" api:"nullable"`
	Description            string    `json:"description" api:"nullable"`
	EndDate                time.Time `json:"endDate" api:"nullable" format:"date"`
	Name                   string    `json:"name" api:"nullable"`
	StartDate              time.Time `json:"startDate" api:"nullable" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		Frequency              respjson.Field
		Status                 respjson.Field
		AmountPerUserPerPeriod respjson.Field
		Description            respjson.Field
		EndDate                respjson.Field
		Name                   respjson.Field
		StartDate              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Budget) RawJSON() string { return r.JSON.raw }
func (r *Budget) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetExpenditure struct {
	// ID for the budget expenditure
	ID string `json:"id" api:"required" format:"uuid"`
	// A dollar amount
	AssignedAmount float64 `json:"assignedAmount" api:"nullable"`
	// ID for the expense
	ExpenseID string `json:"expenseId" api:"nullable" format:"uuid"`
	// ID for this transaction
	TransactionID string `json:"transactionId" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AssignedAmount respjson.Field
		ExpenseID      respjson.Field
		TransactionID  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BudgetExpenditure) RawJSON() string { return r.JSON.raw }
func (r *BudgetExpenditure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetFrequency string

const (
	BudgetFrequencyOneTime BudgetFrequency = "oneTime"
	BudgetFrequencyDaily   BudgetFrequency = "daily"
	BudgetFrequencyWeekly  BudgetFrequency = "weekly"
	BudgetFrequencyMonthly BudgetFrequency = "monthly"
	BudgetFrequencyYearly  BudgetFrequency = "yearly"
)

type BudgetIDResponse struct {
	// ID for the budget
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BudgetIDResponse) RawJSON() string { return r.JSON.raw }
func (r *BudgetIDResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetMember struct {
	// A non-negative dollar amount
	CycleRemaining float64 `json:"cycleRemaining" api:"required"`
	// A non-negative dollar amount
	CycleSpent float64 `json:"cycleSpent" api:"required"`
	Email      string  `json:"email" api:"required"`
	FirstName  string  `json:"firstName" api:"required"`
	LastName   string  `json:"lastName" api:"required"`
	// ID for the organization user
	OrgUserID string `json:"orgUserId" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CycleRemaining respjson.Field
		CycleSpent     respjson.Field
		Email          respjson.Field
		FirstName      respjson.Field
		LastName       respjson.Field
		OrgUserID      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BudgetMember) RawJSON() string { return r.JSON.raw }
func (r *BudgetMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BudgetPublishStatus string

const (
	BudgetPublishStatusDraft     BudgetPublishStatus = "draft"
	BudgetPublishStatusPublished BudgetPublishStatus = "published"
	BudgetPublishStatusArchived  BudgetPublishStatus = "archived"
)

type SpendBudgetNewParams struct {
	// Any of "oneTime", "daily", "weekly", "monthly", "yearly".
	Frequency BudgetFrequency `json:"frequency,omitzero" api:"required"`
	// A non-negative dollar amount
	AmountPerUserPerPeriod param.Opt[float64]   `json:"amountPerUserPerPeriod,omitzero"`
	Description            param.Opt[string]    `json:"description,omitzero"`
	EndDate                param.Opt[time.Time] `json:"endDate,omitzero" format:"date"`
	Name                   param.Opt[string]    `json:"name,omitzero"`
	StartDate              param.Opt[time.Time] `json:"startDate,omitzero" format:"date"`
	IdempotencyKey         param.Opt[string]    `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r SpendBudgetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SpendBudgetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendBudgetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendBudgetUpdateParams struct {
	// A non-negative dollar amount
	AmountPerUserPerPeriod param.Opt[float64]   `json:"amountPerUserPerPeriod,omitzero"`
	Description            param.Opt[string]    `json:"description,omitzero"`
	EndDate                param.Opt[time.Time] `json:"endDate,omitzero" format:"date"`
	Name                   param.Opt[string]    `json:"name,omitzero"`
	StartDate              param.Opt[time.Time] `json:"startDate,omitzero" format:"date"`
	IdempotencyKey         param.Opt[string]    `header:"Idempotency-Key,omitzero" json:"-"`
	// Any of "oneTime", "daily", "weekly", "monthly", "yearly".
	Frequency BudgetFrequency `json:"frequency,omitzero"`
	paramObj
}

func (r SpendBudgetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SpendBudgetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendBudgetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendBudgetListParams struct {
	// The ID of the budget to end the page before (exclusive). When provided, results
	// will end just before this ID and work backwards. Use this for reverse pagination
	// or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the budget to start the page after (exclusive). When provided, results
	// will begin with the budget immediately following this ID. Use this for standard
	// forward pagination to get the next page of results. Cannot be combined with
	// end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// The ID of the budget to start the page at (inclusive). Cannot be combined with
	// start_after or end_before.
	StartAt param.Opt[string] `query:"start_at,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order SpendBudgetListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpendBudgetListParams]'s query parameters as `url.Values`.
func (r SpendBudgetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type SpendBudgetListParamsOrder string

const (
	SpendBudgetListParamsOrderAsc  SpendBudgetListParamsOrder = "asc"
	SpendBudgetListParamsOrderDesc SpendBudgetListParamsOrder = "desc"
)

type SpendBudgetArchiveParams struct {
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

type SpendBudgetAssignTransactionParams struct {
	// ID for this transaction
	TransactionID string `json:"transactionId" api:"required" format:"uuid"`
	// Whether to allow this assignment to exceed the budget. Defaults to false.
	AllowOverspend param.Opt[bool]   `json:"allowOverspend,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r SpendBudgetAssignTransactionParams) MarshalJSON() (data []byte, err error) {
	type shadow SpendBudgetAssignTransactionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendBudgetAssignTransactionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendBudgetListExpendituresParams struct {
	// The ID of the budget expenditure to end the page before (exclusive). When
	// provided, results will end just before this ID and work backwards. Use this for
	// reverse pagination or to retrieve previous pages. Cannot be combined with
	// start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the budget expenditure to start the page after (exclusive). When
	// provided, results will begin with the budget expenditure immediately following
	// this ID. Use this for standard forward pagination to get the next page of
	// results. Cannot be combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// The ID of the budget expenditure to start the page at (inclusive). Cannot be
	// combined with start_after or end_before.
	StartAt param.Opt[string] `query:"start_at,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order SpendBudgetListExpendituresParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpendBudgetListExpendituresParams]'s query parameters as
// `url.Values`.
func (r SpendBudgetListExpendituresParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type SpendBudgetListExpendituresParamsOrder string

const (
	SpendBudgetListExpendituresParamsOrderAsc  SpendBudgetListExpendituresParamsOrder = "asc"
	SpendBudgetListExpendituresParamsOrderDesc SpendBudgetListExpendituresParamsOrder = "desc"
)

type SpendBudgetListMembersParams struct {
	// The ID of the organization user to end the page before (exclusive). When
	// provided, results will end just before this ID and work backwards. Use this for
	// reverse pagination or to retrieve previous pages. Cannot be combined with
	// start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the organization user to start the page after (exclusive). When
	// provided, results will begin with the organization user immediately following
	// this ID. Use this for standard forward pagination to get the next page of
	// results. Cannot be combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// The ID of the organization user to start the page at (inclusive). Cannot be
	// combined with start_after or end_before.
	StartAt param.Opt[string] `query:"start_at,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order SpendBudgetListMembersParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpendBudgetListMembersParams]'s query parameters as
// `url.Values`.
func (r SpendBudgetListMembersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type SpendBudgetListMembersParamsOrder string

const (
	SpendBudgetListMembersParamsOrderAsc  SpendBudgetListMembersParamsOrder = "asc"
	SpendBudgetListMembersParamsOrderDesc SpendBudgetListMembersParamsOrder = "desc"
)
