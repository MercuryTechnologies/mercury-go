// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/mercury-go/internal/apijson"
	"github.com/stainless-sdks/mercury-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/mercury-go/internal/encoding/json"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/param"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// BookJournalEntryService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBookJournalEntryService] method instead.
type BookJournalEntryService struct {
	Options []option.RequestOption
}

// NewBookJournalEntryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBookJournalEntryService(opts ...option.RequestOption) (r BookJournalEntryService) {
	r = BookJournalEntryService{}
	r.Options = opts
	return
}

// List all journal entries
func (r *BookJournalEntryService) List(ctx context.Context, booksID string, query BookJournalEntryListParams, opts ...option.RequestOption) (res *BookJournalEntryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if booksID == "" {
		err = errors.New("missing required booksId parameter")
		return
	}
	path := fmt.Sprintf("books/journal-entries/%s", booksID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Bulk delete journal entries
func (r *BookJournalEntryService) BulkDelete(ctx context.Context, booksID string, body BookJournalEntryBulkDeleteParams, opts ...option.RequestOption) (res *[]BookJournalEntryBulkDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if booksID == "" {
		err = errors.New("missing required booksId parameter")
		return
	}
	path := fmt.Sprintf("books/journal-entries/%s", booksID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Bulk update journal entries
func (r *BookJournalEntryService) BulkUpdate(ctx context.Context, booksID string, body BookJournalEntryBulkUpdateParams, opts ...option.RequestOption) (res *[]BookJournalEntryBulkUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if booksID == "" {
		err = errors.New("missing required booksId parameter")
		return
	}
	path := fmt.Sprintf("books/journal-entries/%s", booksID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Create multiple Journal Entries
func (r *BookJournalEntryService) NewMultiple(ctx context.Context, booksID string, params BookJournalEntryNewMultipleParams, opts ...option.RequestOption) (res *[]JournalEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if booksID == "" {
		err = errors.New("missing required booksId parameter")
		return
	}
	path := fmt.Sprintf("books/journal-entries/%s", booksID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type BookJournalEntryListResponse struct {
	// The token to use to request the next page
	NextPageToken string `json:"next_page_token,nullable"`
	// The token to use to request the previous page
	PrevPageToken string         `json:"prev_page_token,nullable"`
	Records       []JournalEntry `json:"records"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageToken respjson.Field
		PrevPageToken respjson.Field
		Records       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BookJournalEntryListResponse) RawJSON() string { return r.JSON.raw }
func (r *BookJournalEntryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BookJournalEntryBulkDeleteResponse = any

type BookJournalEntryBulkUpdateResponse = any

type BookJournalEntryListParams struct {
	// Only return journal entries with the specified monetary amount. If specified,
	// `min_amount` and `max_amount` parameters are ignored. If no decimal place
	// provided, rounded values will be matched. E.g. an input of `100` will match
	// 100.13 and 100.45, but an input of `100.13` will only match that value.
	Amount param.Opt[float64] `query:"amount,omitzero" json:"-"`
	// Only return Journal Entries before the end date (inclusive).
	EndDate param.Opt[time.Time] `query:"end_date,omitzero" format:"date" json:"-"`
	// Only return Journal Entries where the description matches the specified
	// keyword(s).
	Keywords param.Opt[string] `query:"keywords,omitzero" json:"-"`
	// Only return Journal Entries where at least one Line Entry belongs to the
	// specified Ledger.
	LedgerID param.Opt[string] `query:"ledger_id,omitzero" format:"base58 (max length 22)" json:"-"`
	// The maximum number of journal entries to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only return journal entries with a monetary amount less than specified amount.
	// Ignored if `amount` is set.
	MaxAmount param.Opt[float64] `query:"max_amount,omitzero" json:"-"`
	// Only return journal entries with a monetary amount of at least the specified
	// amount. Ignored if `amount` is set.
	MinAmount param.Opt[float64] `query:"min_amount,omitzero" json:"-"`
	PageToken param.Opt[string]  `query:"page_token,omitzero" json:"-"`
	// Only return Journal Entries after the start date (inclusive).
	StartDate param.Opt[time.Time] `query:"start_date,omitzero" format:"date" json:"-"`
	TimeZone  param.Opt[string]    `query:"time_zone,omitzero" json:"-"`
	// Only return Journal Entries with the specified creation source.
	//
	// Any of "manual", "transaction", "accrual", "payment_application", "legacy".
	CreationSource BookJournalEntryListParamsCreationSource `query:"creation_source,omitzero" json:"-"`
	// Comma-separated list of expandable paths.
	//
	// Any of "line_entries", "line_entries.transaction", "line_entries.ledger",
	// "line_entries.opposing_ledger_ids", "line_entries.payment_applications".
	Expand BookJournalEntryListParamsExpand `query:"expand,omitzero" json:"-"`
	// One or more comma-separated lists of ledger IDs. Only return Journal Entries
	// where at least one Line Entry belongs to one of the specified ledgers in a
	// group. Specifying multiple groups will return an intersection of the matches.
	LedgerIDGroups []string `query:"ledger_id_groups,omitzero" format:"base58 (max length 22)" json:"-"`
	// A comma separated list of sort orders for the results.
	//
	// Any of "description", "-description", "datetime", "-datetime".
	Sort BookJournalEntryListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BookJournalEntryListParams]'s query parameters as
// `url.Values`.
func (r BookJournalEntryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only return Journal Entries with the specified creation source.
type BookJournalEntryListParamsCreationSource string

const (
	BookJournalEntryListParamsCreationSourceManual             BookJournalEntryListParamsCreationSource = "manual"
	BookJournalEntryListParamsCreationSourceTransaction        BookJournalEntryListParamsCreationSource = "transaction"
	BookJournalEntryListParamsCreationSourceAccrual            BookJournalEntryListParamsCreationSource = "accrual"
	BookJournalEntryListParamsCreationSourcePaymentApplication BookJournalEntryListParamsCreationSource = "payment_application"
	BookJournalEntryListParamsCreationSourceLegacy             BookJournalEntryListParamsCreationSource = "legacy"
)

// Comma-separated list of expandable paths.
type BookJournalEntryListParamsExpand string

const (
	BookJournalEntryListParamsExpandLineEntries                    BookJournalEntryListParamsExpand = "line_entries"
	BookJournalEntryListParamsExpandLineEntriesTransaction         BookJournalEntryListParamsExpand = "line_entries.transaction"
	BookJournalEntryListParamsExpandLineEntriesLedger              BookJournalEntryListParamsExpand = "line_entries.ledger"
	BookJournalEntryListParamsExpandLineEntriesOpposingLedgerIDs   BookJournalEntryListParamsExpand = "line_entries.opposing_ledger_ids"
	BookJournalEntryListParamsExpandLineEntriesPaymentApplications BookJournalEntryListParamsExpand = "line_entries.payment_applications"
)

// A comma separated list of sort orders for the results.
type BookJournalEntryListParamsSort string

const (
	BookJournalEntryListParamsSortDescription      BookJournalEntryListParamsSort = "description"
	BookJournalEntryListParamsSortMinusDescription BookJournalEntryListParamsSort = "-description"
	BookJournalEntryListParamsSortDatetime         BookJournalEntryListParamsSort = "datetime"
	BookJournalEntryListParamsSortMinusDatetime    BookJournalEntryListParamsSort = "-datetime"
)

type BookJournalEntryBulkDeleteParams struct {
	// The journal entries to delete
	JournalEntryIDs []string `json:"journal_entry_ids,omitzero,required" format:"base58 (max length 22)"`
	paramObj
}

func (r BookJournalEntryBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow BookJournalEntryBulkDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BookJournalEntryBulkDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BookJournalEntryBulkUpdateParams struct {
	Body []BookJournalEntryBulkUpdateParamsBody
	paramObj
}

func (r BookJournalEntryBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *BookJournalEntryBulkUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The property ID is required.
type BookJournalEntryBulkUpdateParamsBody struct {
	// The unique ID of the object.
	ID string `json:"id,required" format:"base58 (max length 22)"`
	// The datetime the Journal Entry was created in UTC time.
	Datetime param.Opt[string] `json:"datetime,omitzero" format:"yyyy-mm-ddThh:MM:ssZ"`
	// An arbitrary string on the object, useful for displaying to the user.
	Description param.Opt[string] `json:"description,omitzero"`
	// An object with optional `create`, `update`, and `delete` parameters to modify
	// the Line Entries associated with the Journal Entry. The `create` and `update`
	// parameters accept lists of Line Entry objects, while the `delete` parameter
	// accepts a list of existing Line Entry IDs.
	LineEntryChanges BookJournalEntryBulkUpdateParamsBodyLineEntryChanges `json:"line_entry_changes,omitzero"`
	paramObj
}

func (r BookJournalEntryBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow BookJournalEntryBulkUpdateParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BookJournalEntryBulkUpdateParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with optional `create`, `update`, and `delete` parameters to modify
// the Line Entries associated with the Journal Entry. The `create` and `update`
// parameters accept lists of Line Entry objects, while the `delete` parameter
// accepts a list of existing Line Entry IDs.
type BookJournalEntryBulkUpdateParamsBodyLineEntryChanges struct {
	// A list of new Line Entry objects to be created on the Journal Entry.
	Create []BookJournalEntryBulkUpdateParamsBodyLineEntryChangesCreate `json:"create,omitzero"`
	// A list of existing Line Entry IDs to delete.
	Delete []string `json:"delete,omitzero" format:"base58 (max length 22)"`
	// A list of existing Line Entry objects to be updated on the Journal Entry.
	Update []BookJournalEntryBulkUpdateParamsBodyLineEntryChangesUpdate `json:"update,omitzero"`
	paramObj
}

func (r BookJournalEntryBulkUpdateParamsBodyLineEntryChanges) MarshalJSON() (data []byte, err error) {
	type shadow BookJournalEntryBulkUpdateParamsBodyLineEntryChanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BookJournalEntryBulkUpdateParamsBodyLineEntryChanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, DebitCredit, Description, LedgerID are required.
type BookJournalEntryBulkUpdateParamsBodyLineEntryChangesCreate struct {
	// A dollar amount
	Amount float64 `json:"amount,required"`
	// Indicates if the amount is a credit or debit.
	//
	// Any of "debit", "credit".
	DebitCredit string `json:"debit_credit,omitzero,required"`
	// An arbitrary string on the object, useful for displaying to the user or for
	// categorization.
	Description string `json:"description,required"`
	// The ID of the Line Entry's Ledger.
	LedgerID string `json:"ledger_id,required" format:"base58 (max length 22)"`
	// A list of IDs of the Tag objects associated with the Line Entry.
	TagIDs []string `json:"tag_ids,omitzero" format:"base58 (max length 22)"`
	paramObj
}

func (r BookJournalEntryBulkUpdateParamsBodyLineEntryChangesCreate) MarshalJSON() (data []byte, err error) {
	type shadow BookJournalEntryBulkUpdateParamsBodyLineEntryChangesCreate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BookJournalEntryBulkUpdateParamsBodyLineEntryChangesCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BookJournalEntryBulkUpdateParamsBodyLineEntryChangesCreate](
		"debit_credit", "debit", "credit",
	)
}

// The property ID is required.
type BookJournalEntryBulkUpdateParamsBodyLineEntryChangesUpdate struct {
	// The unique ID of the object.
	ID string `json:"id,required" format:"base58 (max length 22)"`
	// A dollar amount
	Amount param.Opt[float64] `json:"amount,omitzero"`
	// An arbitrary string on the object, useful for displaying to the user or for
	// categorization.
	Description param.Opt[string] `json:"description,omitzero"`
	// The ID of the Line Entry's Ledger.
	LedgerID param.Opt[string] `json:"ledger_id,omitzero" format:"base58 (max length 22)"`
	// Indicates if the amount is a credit or debit.
	//
	// Any of "debit", "credit".
	DebitCredit string `json:"debit_credit,omitzero"`
	// A list of IDs of the Tag objects associated with the Line Entry.
	TagIDs []string `json:"tag_ids,omitzero" format:"base58 (max length 22)"`
	paramObj
}

func (r BookJournalEntryBulkUpdateParamsBodyLineEntryChangesUpdate) MarshalJSON() (data []byte, err error) {
	type shadow BookJournalEntryBulkUpdateParamsBodyLineEntryChangesUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BookJournalEntryBulkUpdateParamsBodyLineEntryChangesUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BookJournalEntryBulkUpdateParamsBodyLineEntryChangesUpdate](
		"debit_credit", "debit", "credit",
	)
}

type BookJournalEntryNewMultipleParams struct {
	// An optional identifier for audit logging.
	ModifiedBy param.Opt[string] `query:"modified_by,omitzero" format:"uuid" json:"-"`
	// Comma-separated list of expandable paths.
	//
	// Any of "line_entries", "line_entries.transaction", "line_entries.ledger",
	// "line_entries.opposing_ledger_ids", "line_entries.payment_applications".
	Expand BookJournalEntryNewMultipleParamsExpand `query:"expand,omitzero" json:"-"`
	Body   []BookJournalEntryNewMultipleParamsBody
	paramObj
}

func (r BookJournalEntryNewMultipleParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *BookJournalEntryNewMultipleParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// URLQuery serializes [BookJournalEntryNewMultipleParams]'s query parameters as
// `url.Values`.
func (r BookJournalEntryNewMultipleParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Comma-separated list of expandable paths.
type BookJournalEntryNewMultipleParamsExpand string

const (
	BookJournalEntryNewMultipleParamsExpandLineEntries                    BookJournalEntryNewMultipleParamsExpand = "line_entries"
	BookJournalEntryNewMultipleParamsExpandLineEntriesTransaction         BookJournalEntryNewMultipleParamsExpand = "line_entries.transaction"
	BookJournalEntryNewMultipleParamsExpandLineEntriesLedger              BookJournalEntryNewMultipleParamsExpand = "line_entries.ledger"
	BookJournalEntryNewMultipleParamsExpandLineEntriesOpposingLedgerIDs   BookJournalEntryNewMultipleParamsExpand = "line_entries.opposing_ledger_ids"
	BookJournalEntryNewMultipleParamsExpandLineEntriesPaymentApplications BookJournalEntryNewMultipleParamsExpand = "line_entries.payment_applications"
)

// The properties Datetime, Description, LineEntries are required.
type BookJournalEntryNewMultipleParamsBody struct {
	// The datetime the Journal Entry was created in UTC time.
	Datetime string `json:"datetime,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// An arbitrary string on the object, useful for displaying to the user.
	Description string `json:"description,required"`
	// The Line Entry objects to be associated with the Journal Entry. Must include at
	// least two Line Entries, one debit and one credit, and the sum of the `amount`s
	// of all debit entries must equal the sum of the `amount`s of the credit entries.
	LineEntries []BookJournalEntryNewMultipleParamsBodyLineEntry `json:"line_entries,omitzero,required"`
	paramObj
}

func (r BookJournalEntryNewMultipleParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow BookJournalEntryNewMultipleParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BookJournalEntryNewMultipleParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Amount, DebitCredit, Description, LedgerID are required.
type BookJournalEntryNewMultipleParamsBodyLineEntry struct {
	// A dollar amount
	Amount float64 `json:"amount,required"`
	// Indicates if the amount is a credit or debit.
	//
	// Any of "debit", "credit".
	DebitCredit string `json:"debit_credit,omitzero,required"`
	// An arbitrary string on the object, useful for displaying to the user or for
	// categorization.
	Description string `json:"description,required"`
	// The ID of the Line Entry's Ledger.
	LedgerID string `json:"ledger_id,required" format:"base58 (max length 22)"`
	// A list of IDs of the Tag objects associated with the Line Entry.
	TagIDs []string `json:"tag_ids,omitzero" format:"base58 (max length 22)"`
	paramObj
}

func (r BookJournalEntryNewMultipleParamsBodyLineEntry) MarshalJSON() (data []byte, err error) {
	type shadow BookJournalEntryNewMultipleParamsBodyLineEntry
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BookJournalEntryNewMultipleParamsBodyLineEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BookJournalEntryNewMultipleParamsBodyLineEntry](
		"debit_credit", "debit", "credit",
	)
}
