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

	"github.com/stainless-sdks/mercury-go/internal/apijson"
	"github.com/stainless-sdks/mercury-go/internal/apiquery"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/param"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// BookService contains methods and other services that help with interacting with
// the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBookService] method instead.
type BookService struct {
	Options             []option.RequestOption
	AgentLedgerTemplate BookAgentLedgerTemplateService
	JournalEntries      BookJournalEntryService
	AgentCoaTemplates   BookAgentCoaTemplateService
}

// NewBookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBookService(opts ...option.RequestOption) (r BookService) {
	r = BookService{}
	r.Options = opts
	r.AgentLedgerTemplate = NewBookAgentLedgerTemplateService(opts...)
	r.JournalEntries = NewBookJournalEntryService(opts...)
	r.AgentCoaTemplates = NewBookAgentCoaTemplateService(opts...)
	return
}

// Create a new ledger within an agent-owned Chart of Accounts template.
func (r *BookService) NewAgentLedgerTemplate(ctx context.Context, body BookNewAgentLedgerTemplateParams, opts ...option.RequestOption) (res *LedgerTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "books/agent-ledger-templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a Journal Entry
func (r *BookService) GetJournalEntry(ctx context.Context, tealJournalEntryID string, params BookGetJournalEntryParams, opts ...option.RequestOption) (res *JournalEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if params.BooksID == "" {
		err = errors.New("missing required booksId parameter")
		return
	}
	if tealJournalEntryID == "" {
		err = errors.New("missing required tealJournalEntryId parameter")
		return
	}
	path := fmt.Sprintf("books/journal-entry/%s/%s", params.BooksID, tealJournalEntryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type JournalEntry struct {
	// The unique ID of the object.
	ID string `json:"id,required" format:"base58 (max length 22)"`
	// The timestamp when this journal entry record was created in the database.
	CreationDate string `json:"creation_date,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// The source of the journal entry creation: 'manual' if created manually by the
	// user, 'transaction' if auto-generated from transaction categorization (e.g.,
	// Plaid sync), or 'legacy' if auto-generated from the old AR/AP system.
	//
	// Any of "manual", "transaction", "accrual", "payment_application", "legacy".
	CreationSource JournalEntryCreationSource `json:"creation_source,required"`
	// The datetime the Journal Entry was created in UTC time
	Datetime string `json:"datetime,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// An arbitrary string on the object, useful for displaying to the user.
	Description string `json:"description,required"`
	// Whether the Journal Entry represents an opening balance.
	IsOpeningBalance bool `json:"is_opening_balance,required"`
	// List of Line Entries. Included in expanded responses.
	LineEntries []JournalEntryLineEntry `json:"line_entries,nullable"`
	// The ids of the Line Entries detailing the Journal Entry's movement of value.
	LineEntryIDs []string `json:"line_entry_ids,nullable" format:"base58 (max length 22)"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreationDate     respjson.Field
		CreationSource   respjson.Field
		Datetime         respjson.Field
		Description      respjson.Field
		IsOpeningBalance respjson.Field
		LineEntries      respjson.Field
		LineEntryIDs     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JournalEntry) RawJSON() string { return r.JSON.raw }
func (r *JournalEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The source of the journal entry creation: 'manual' if created manually by the
// user, 'transaction' if auto-generated from transaction categorization (e.g.,
// Plaid sync), or 'legacy' if auto-generated from the old AR/AP system.
type JournalEntryCreationSource string

const (
	JournalEntryCreationSourceManual             JournalEntryCreationSource = "manual"
	JournalEntryCreationSourceTransaction        JournalEntryCreationSource = "transaction"
	JournalEntryCreationSourceAccrual            JournalEntryCreationSource = "accrual"
	JournalEntryCreationSourcePaymentApplication JournalEntryCreationSource = "payment_application"
	JournalEntryCreationSourceLegacy             JournalEntryCreationSource = "legacy"
)

type JournalEntryLineEntry struct {
	// The unique ID of the object.
	ID string `json:"id,required" format:"base58 (max length 22)"`
	// A dollar amount
	Amount float64 `json:"amount,required"`
	// The datetime the Line Entry was created in UTC time.
	Datetime string `json:"datetime,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Indicates if the amount is a credit or debit.
	//
	// Any of "debit", "credit".
	DebitCredit string `json:"debit_credit,required"`
	// Indicates if the Line Entry can be manually added or removed from a Ledger, as
	// dictated by the `editable` property on the Ledger.
	Editable bool `json:"editable,required"`
	// An arbitrary string on the object, useful for displaying to the user.
	JournalEntryDescription string `json:"journal_entry_description,required"`
	// The ID of the associated Journal Entry.
	JournalEntryID string `json:"journal_entry_id,required" format:"base58 (max length 22)"`
	// The ID of the Line Entry's Ledger.
	LedgerID string `json:"ledger_id,required" format:"base58 (max length 22)"`
	// The name of the Line Entry's Ledger.
	LedgerName string `json:"ledger_name,required"`
	// An arbitrary string on the object, useful for displaying to the user or for
	// categorization.
	Description string `json:"description,nullable"`
	// Expandable. IDs of the ledgers on the other side(s) of the line entry with
	// opposite debit/credit.
	OpposingLedgerIDs   []string                                  `json:"opposing_ledger_ids,nullable" format:"base58 (max length 22)"`
	PaymentApplications []JournalEntryLineEntryPaymentApplication `json:"payment_applications,nullable"`
	// If the Line Entry is identified as a transfer between accounts, the ID of the
	// Line Entry in the opposing Ledger. Related guide:
	// [Auto-categorization](https://docs.teal.dev/guides/platform/categorization/pipeline#transfers-between-accounts)
	RelatedLineEntryID string `json:"related_line_entry_id,nullable" format:"base58 (max length 22)"`
	// The ID of the associated Transaction, if one exists.
	TransactionID string `json:"transaction_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Amount                  respjson.Field
		Datetime                respjson.Field
		DebitCredit             respjson.Field
		Editable                respjson.Field
		JournalEntryDescription respjson.Field
		JournalEntryID          respjson.Field
		LedgerID                respjson.Field
		LedgerName              respjson.Field
		Description             respjson.Field
		OpposingLedgerIDs       respjson.Field
		PaymentApplications     respjson.Field
		RelatedLineEntryID      respjson.Field
		TransactionID           respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JournalEntryLineEntry) RawJSON() string { return r.JSON.raw }
func (r *JournalEntryLineEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JournalEntryLineEntryPaymentApplication struct {
	Adjustments []JournalEntryLineEntryPaymentApplicationAdjustment `json:"_adjustments,required"`
	// A dollar amount
	MatchedAmount        float64 `json:"_matchedAmount,required"`
	PaymentApplicationID string  `json:"_paymentApplicationId,required"`
	// Any of "bank_transaction", "customer_credit", "manual".
	SourceType string `json:"_sourceType,required"`
	// Any of "proposed", "confirmed", "rejected".
	Status        string                                                `json:"_status,required"`
	AccrualSource JournalEntryLineEntryPaymentApplication_AccrualSource `json:"_accrualSource"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Adjustments          respjson.Field
		MatchedAmount        respjson.Field
		PaymentApplicationID respjson.Field
		SourceType           respjson.Field
		Status               respjson.Field
		AccrualSource        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JournalEntryLineEntryPaymentApplication) RawJSON() string { return r.JSON.raw }
func (r *JournalEntryLineEntryPaymentApplication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JournalEntryLineEntryPaymentApplicationAdjustment struct {
	AccrualAdjustmentID string `json:"_accrualAdjustmentId,required"`
	// A dollar amount
	Amount float64 `json:"_amount,required"`
	// The unique identifier for a Ledger
	CreditAccountID string `json:"_creditAccountId,required" format:"base58 (max length 22)"`
	// The unique identifier for a Ledger
	DebitAccountID string `json:"_debitAccountId,required" format:"base58 (max length 22)"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccrualAdjustmentID respjson.Field
		Amount              respjson.Field
		CreditAccountID     respjson.Field
		DebitAccountID      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JournalEntryLineEntryPaymentApplicationAdjustment) RawJSON() string { return r.JSON.raw }
func (r *JournalEntryLineEntryPaymentApplicationAdjustment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JournalEntryLineEntryPaymentApplication_AccrualSource struct {
	AccrualSourceID string    `json:"_accrualSourceId,required"`
	EffectiveDate   time.Time `json:"_effectiveDate,required" format:"date"`
	// A dollar amount
	GrossAmount     float64 `json:"_grossAmount,required"`
	ReferenceNumber string  `json:"_referenceNumber,required"`
	// Any of "invoice", "bill".
	Type       string `json:"_type,required"`
	ExternalID string `json:"_externalId"`
	// The unique identifier for an External Party
	ExternalPartyID string `json:"_externalPartyId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccrualSourceID respjson.Field
		EffectiveDate   respjson.Field
		GrossAmount     respjson.Field
		ReferenceNumber respjson.Field
		Type            respjson.Field
		ExternalID      respjson.Field
		ExternalPartyID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JournalEntryLineEntryPaymentApplication_AccrualSource) RawJSON() string { return r.JSON.raw }
func (r *JournalEntryLineEntryPaymentApplication_AccrualSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BookNewAgentLedgerTemplateParams struct {
	// The COA Template Id.
	CoaTemplateID string `json:"coa_template_id,required" format:"base58 (max length 22)"`
	// Indicates if the ledger is a credit or debit ledger.
	//
	// Any of "debit", "credit".
	DebitCredit BookNewAgentLedgerTemplateParamsDebitCredit `json:"debit_credit,omitzero,required"`
	// Whether line entries can be manually added or removed from the ledger.
	Editable bool `json:"editable,required"`
	// Indicates that the ledger represents a real-world financial account.
	//
	// Any of "bank_account", "credit_card", "payments", "payroll", "loan",
	// "prepaid_card", "accounts_receivable", "accounts_payable", "investment",
	// "treasury".
	FinancialAccountType BookNewAgentLedgerTemplateParamsFinancialAccountType `json:"financial_account_type,omitzero,required"`
	// Indicates similar characteristics and accounting treatment for a group of
	// ledgers within a `type`.
	//
	// Any of "current_assets", "non-current_assets", "transfers_between_accounts",
	// "uncategorized_assets", "current_liabilities", "non-current_liabilities",
	// "equity", "operating_revenues", "other_income", "cost_of_goods_sold",
	// "operating_expenses", "other_expenses", "retained_earnings".
	LedgerSubType BookNewAgentLedgerTemplateParamsLedgerSubType `json:"ledger_sub_type,omitzero,required"`
	// Indicates the purpose and location of the funds and value recorded in the ledger
	// and which report it is included in: `asset`, `liability` and `equity` ledgers
	// are displayed on the balance sheet; `revenue` and `expenses` are displayed on
	// the income statement.
	//
	// Any of "asset", "liability", "equity", "revenue", "expense".
	LedgerType BookNewAgentLedgerTemplateParamsLedgerType `json:"ledger_type,omitzero,required"`
	// The name of the ledger.
	Name string `json:"name,required"`
	// Whether Teal includes this ledger in the cash flow report.
	ReportCashFlow bool `json:"report_cash_flow,required"`
	// Determines the display order in reports, ordered digit by digit,
	SortCode string `json:"sort_code,required"`
	// If the ledger is a child Ledger, the ID of the parent ledger object.
	ParentID param.Opt[string] `json:"parent_id,omitzero" format:"base58 (max length 22)"`
	paramObj
}

func (r BookNewAgentLedgerTemplateParams) MarshalJSON() (data []byte, err error) {
	type shadow BookNewAgentLedgerTemplateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BookNewAgentLedgerTemplateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates if the ledger is a credit or debit ledger.
type BookNewAgentLedgerTemplateParamsDebitCredit string

const (
	BookNewAgentLedgerTemplateParamsDebitCreditDebit  BookNewAgentLedgerTemplateParamsDebitCredit = "debit"
	BookNewAgentLedgerTemplateParamsDebitCreditCredit BookNewAgentLedgerTemplateParamsDebitCredit = "credit"
)

// Indicates that the ledger represents a real-world financial account.
type BookNewAgentLedgerTemplateParamsFinancialAccountType string

const (
	BookNewAgentLedgerTemplateParamsFinancialAccountTypeBankAccount        BookNewAgentLedgerTemplateParamsFinancialAccountType = "bank_account"
	BookNewAgentLedgerTemplateParamsFinancialAccountTypeCreditCard         BookNewAgentLedgerTemplateParamsFinancialAccountType = "credit_card"
	BookNewAgentLedgerTemplateParamsFinancialAccountTypePayments           BookNewAgentLedgerTemplateParamsFinancialAccountType = "payments"
	BookNewAgentLedgerTemplateParamsFinancialAccountTypePayroll            BookNewAgentLedgerTemplateParamsFinancialAccountType = "payroll"
	BookNewAgentLedgerTemplateParamsFinancialAccountTypeLoan               BookNewAgentLedgerTemplateParamsFinancialAccountType = "loan"
	BookNewAgentLedgerTemplateParamsFinancialAccountTypePrepaidCard        BookNewAgentLedgerTemplateParamsFinancialAccountType = "prepaid_card"
	BookNewAgentLedgerTemplateParamsFinancialAccountTypeAccountsReceivable BookNewAgentLedgerTemplateParamsFinancialAccountType = "accounts_receivable"
	BookNewAgentLedgerTemplateParamsFinancialAccountTypeAccountsPayable    BookNewAgentLedgerTemplateParamsFinancialAccountType = "accounts_payable"
	BookNewAgentLedgerTemplateParamsFinancialAccountTypeInvestment         BookNewAgentLedgerTemplateParamsFinancialAccountType = "investment"
	BookNewAgentLedgerTemplateParamsFinancialAccountTypeTreasury           BookNewAgentLedgerTemplateParamsFinancialAccountType = "treasury"
)

// Indicates similar characteristics and accounting treatment for a group of
// ledgers within a `type`.
type BookNewAgentLedgerTemplateParamsLedgerSubType string

const (
	BookNewAgentLedgerTemplateParamsLedgerSubTypeCurrentAssets            BookNewAgentLedgerTemplateParamsLedgerSubType = "current_assets"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeNonCurrentAssets         BookNewAgentLedgerTemplateParamsLedgerSubType = "non-current_assets"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeTransfersBetweenAccounts BookNewAgentLedgerTemplateParamsLedgerSubType = "transfers_between_accounts"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeUncategorizedAssets      BookNewAgentLedgerTemplateParamsLedgerSubType = "uncategorized_assets"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeCurrentLiabilities       BookNewAgentLedgerTemplateParamsLedgerSubType = "current_liabilities"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeNonCurrentLiabilities    BookNewAgentLedgerTemplateParamsLedgerSubType = "non-current_liabilities"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeEquity                   BookNewAgentLedgerTemplateParamsLedgerSubType = "equity"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeOperatingRevenues        BookNewAgentLedgerTemplateParamsLedgerSubType = "operating_revenues"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeOtherIncome              BookNewAgentLedgerTemplateParamsLedgerSubType = "other_income"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeCostOfGoodsSold          BookNewAgentLedgerTemplateParamsLedgerSubType = "cost_of_goods_sold"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeOperatingExpenses        BookNewAgentLedgerTemplateParamsLedgerSubType = "operating_expenses"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeOtherExpenses            BookNewAgentLedgerTemplateParamsLedgerSubType = "other_expenses"
	BookNewAgentLedgerTemplateParamsLedgerSubTypeRetainedEarnings         BookNewAgentLedgerTemplateParamsLedgerSubType = "retained_earnings"
)

// Indicates the purpose and location of the funds and value recorded in the ledger
// and which report it is included in: `asset`, `liability` and `equity` ledgers
// are displayed on the balance sheet; `revenue` and `expenses` are displayed on
// the income statement.
type BookNewAgentLedgerTemplateParamsLedgerType string

const (
	BookNewAgentLedgerTemplateParamsLedgerTypeAsset     BookNewAgentLedgerTemplateParamsLedgerType = "asset"
	BookNewAgentLedgerTemplateParamsLedgerTypeLiability BookNewAgentLedgerTemplateParamsLedgerType = "liability"
	BookNewAgentLedgerTemplateParamsLedgerTypeEquity    BookNewAgentLedgerTemplateParamsLedgerType = "equity"
	BookNewAgentLedgerTemplateParamsLedgerTypeRevenue   BookNewAgentLedgerTemplateParamsLedgerType = "revenue"
	BookNewAgentLedgerTemplateParamsLedgerTypeExpense   BookNewAgentLedgerTemplateParamsLedgerType = "expense"
)

type BookGetJournalEntryParams struct {
	BooksID string `path:"booksId,required" format:"uuid" json:"-"`
	// An optional field to include when retrieving a Journal Entry
	//
	// Any of "line_entries", "line_entries.transaction", "line_entries.ledger",
	// "line_entries.opposing_ledger_ids", "line_entries.payment_applications".
	Expand BookGetJournalEntryParamsExpand `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BookGetJournalEntryParams]'s query parameters as
// `url.Values`.
func (r BookGetJournalEntryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// An optional field to include when retrieving a Journal Entry
type BookGetJournalEntryParamsExpand string

const (
	BookGetJournalEntryParamsExpandLineEntries                    BookGetJournalEntryParamsExpand = "line_entries"
	BookGetJournalEntryParamsExpandLineEntriesTransaction         BookGetJournalEntryParamsExpand = "line_entries.transaction"
	BookGetJournalEntryParamsExpandLineEntriesLedger              BookGetJournalEntryParamsExpand = "line_entries.ledger"
	BookGetJournalEntryParamsExpandLineEntriesOpposingLedgerIDs   BookGetJournalEntryParamsExpand = "line_entries.opposing_ledger_ids"
	BookGetJournalEntryParamsExpandLineEntriesPaymentApplications BookGetJournalEntryParamsExpand = "line_entries.payment_applications"
)
