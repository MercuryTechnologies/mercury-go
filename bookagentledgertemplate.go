// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/mercury-go/internal/apijson"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/param"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// BookAgentLedgerTemplateService contains methods and other services that help
// with interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBookAgentLedgerTemplateService] method instead.
type BookAgentLedgerTemplateService struct {
	Options []option.RequestOption
}

// NewBookAgentLedgerTemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBookAgentLedgerTemplateService(opts ...option.RequestOption) (r BookAgentLedgerTemplateService) {
	r = BookAgentLedgerTemplateService{}
	r.Options = opts
	return
}

// Update an existing ledger within an agent-owned Chart of Accounts template.
func (r *BookAgentLedgerTemplateService) Update(ctx context.Context, ledgerID string, body BookAgentLedgerTemplateUpdateParams, opts ...option.RequestOption) (res *LedgerTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if ledgerID == "" {
		err = errors.New("missing required ledgerId parameter")
		return
	}
	path := fmt.Sprintf("books/agent-ledger-template/%s", ledgerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete an existing ledger within an agent-owned Chart of Accounts template.
func (r *BookAgentLedgerTemplateService) Delete(ctx context.Context, ledgerID string, opts ...option.RequestOption) (res *[]BookAgentLedgerTemplateDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if ledgerID == "" {
		err = errors.New("missing required ledgerId parameter")
		return
	}
	path := fmt.Sprintf("books/agent-ledger-template/%s", ledgerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type LedgerTemplate struct {
	// The unique ID of the ledger template
	ID string `json:"id,required" format:"base58 (max length 22)"`
	// Whether this ledger is a debit or credit account
	//
	// Any of "debit", "credit".
	DebitCredit LedgerTemplateDebitCredit `json:"debit_credit,required"`
	// Whether the ledger can be edited
	Editable bool `json:"editable,required"`
	// Whether the ledger is required
	IsRequired bool `json:"is_required,required"`
	// Ledger name
	Name string `json:"name,required"`
	// Whether to report in cash flow
	ReportCashFlow bool `json:"report_cash_flow,required"`
	// Sort code for ordering
	SortCode string `json:"sort_code,required"`
	// Ledger sub-type
	//
	// Any of "current_assets", "non-current_assets", "transfers_between_accounts",
	// "uncategorized_assets", "current_liabilities", "non-current_liabilities",
	// "equity", "operating_revenues", "other_income", "cost_of_goods_sold",
	// "operating_expenses", "other_expenses", "retained_earnings".
	SubType LedgerTemplateSubType `json:"sub_type,required"`
	// Ledger type (asset, liability, etc.)
	//
	// Any of "asset", "liability", "equity", "revenue", "expense".
	Type LedgerTemplateType `json:"type,required"`
	// Child ledger templates
	Children []LedgerTemplate `json:"children,nullable"`
	// The financial account type
	//
	// Any of "bank_account", "credit_card", "payments", "payroll", "loan",
	// "prepaid_card", "accounts_receivable", "accounts_payable", "investment",
	// "treasury".
	FinancialAccountType LedgerTemplateFinancialAccountType `json:"financial_account_type,nullable"`
	// Parent ledger ID
	ParentID string `json:"parent_id,nullable" format:"base58 (max length 22)"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		DebitCredit          respjson.Field
		Editable             respjson.Field
		IsRequired           respjson.Field
		Name                 respjson.Field
		ReportCashFlow       respjson.Field
		SortCode             respjson.Field
		SubType              respjson.Field
		Type                 respjson.Field
		Children             respjson.Field
		FinancialAccountType respjson.Field
		ParentID             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LedgerTemplate) RawJSON() string { return r.JSON.raw }
func (r *LedgerTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether this ledger is a debit or credit account
type LedgerTemplateDebitCredit string

const (
	LedgerTemplateDebitCreditDebit  LedgerTemplateDebitCredit = "debit"
	LedgerTemplateDebitCreditCredit LedgerTemplateDebitCredit = "credit"
)

// Ledger sub-type
type LedgerTemplateSubType string

const (
	LedgerTemplateSubTypeCurrentAssets            LedgerTemplateSubType = "current_assets"
	LedgerTemplateSubTypeNonCurrentAssets         LedgerTemplateSubType = "non-current_assets"
	LedgerTemplateSubTypeTransfersBetweenAccounts LedgerTemplateSubType = "transfers_between_accounts"
	LedgerTemplateSubTypeUncategorizedAssets      LedgerTemplateSubType = "uncategorized_assets"
	LedgerTemplateSubTypeCurrentLiabilities       LedgerTemplateSubType = "current_liabilities"
	LedgerTemplateSubTypeNonCurrentLiabilities    LedgerTemplateSubType = "non-current_liabilities"
	LedgerTemplateSubTypeEquity                   LedgerTemplateSubType = "equity"
	LedgerTemplateSubTypeOperatingRevenues        LedgerTemplateSubType = "operating_revenues"
	LedgerTemplateSubTypeOtherIncome              LedgerTemplateSubType = "other_income"
	LedgerTemplateSubTypeCostOfGoodsSold          LedgerTemplateSubType = "cost_of_goods_sold"
	LedgerTemplateSubTypeOperatingExpenses        LedgerTemplateSubType = "operating_expenses"
	LedgerTemplateSubTypeOtherExpenses            LedgerTemplateSubType = "other_expenses"
	LedgerTemplateSubTypeRetainedEarnings         LedgerTemplateSubType = "retained_earnings"
)

// Ledger type (asset, liability, etc.)
type LedgerTemplateType string

const (
	LedgerTemplateTypeAsset     LedgerTemplateType = "asset"
	LedgerTemplateTypeLiability LedgerTemplateType = "liability"
	LedgerTemplateTypeEquity    LedgerTemplateType = "equity"
	LedgerTemplateTypeRevenue   LedgerTemplateType = "revenue"
	LedgerTemplateTypeExpense   LedgerTemplateType = "expense"
)

// The financial account type
type LedgerTemplateFinancialAccountType string

const (
	LedgerTemplateFinancialAccountTypeBankAccount        LedgerTemplateFinancialAccountType = "bank_account"
	LedgerTemplateFinancialAccountTypeCreditCard         LedgerTemplateFinancialAccountType = "credit_card"
	LedgerTemplateFinancialAccountTypePayments           LedgerTemplateFinancialAccountType = "payments"
	LedgerTemplateFinancialAccountTypePayroll            LedgerTemplateFinancialAccountType = "payroll"
	LedgerTemplateFinancialAccountTypeLoan               LedgerTemplateFinancialAccountType = "loan"
	LedgerTemplateFinancialAccountTypePrepaidCard        LedgerTemplateFinancialAccountType = "prepaid_card"
	LedgerTemplateFinancialAccountTypeAccountsReceivable LedgerTemplateFinancialAccountType = "accounts_receivable"
	LedgerTemplateFinancialAccountTypeAccountsPayable    LedgerTemplateFinancialAccountType = "accounts_payable"
	LedgerTemplateFinancialAccountTypeInvestment         LedgerTemplateFinancialAccountType = "investment"
	LedgerTemplateFinancialAccountTypeTreasury           LedgerTemplateFinancialAccountType = "treasury"
)

type BookAgentLedgerTemplateDeleteResponse = any

type BookAgentLedgerTemplateUpdateParams struct {
	// The name of the ledger.
	Name param.Opt[string] `json:"name,omitzero"`
	// Determines the display order in reports, ordered digit by digit,
	SortCode param.Opt[string] `json:"sort_code,omitzero"`
	paramObj
}

func (r BookAgentLedgerTemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BookAgentLedgerTemplateUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BookAgentLedgerTemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
