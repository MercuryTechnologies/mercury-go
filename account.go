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

// AccountService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options      []option.RequestOption
	Transactions AccountTransactionService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.Options = opts
	r.Transactions = NewAccountTransactionService(opts...)
	return
}

// Get account by ID
func (r *AccountService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("account/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a paginated list of accounts. Supports cursor-based pagination with
// limit, order, start_after, and end_before query parameters.
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *pagination.CursorIDAccounts[Account], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8"), option.WithResponseInto(&raw)}, opts...)
	path := "accounts"
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

// Retrieve a paginated list of accounts. Supports cursor-based pagination with
// limit, order, start_after, and end_before query parameters.
func (r *AccountService) ListAutoPaging(ctx context.Context, query AccountListParams, opts ...option.RequestOption) *pagination.CursorIDAccountsAutoPager[Account] {
	return pagination.NewCursorIDAccountsAutoPager(r.List(ctx, query, opts...))
}

// Retrieve all debit and credit cards associated with a specific account.
func (r *AccountService) ListCards(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountListCardsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("account/%s/cards", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a paginated list of monthly statements for a specific account. Supports
// cursor-based pagination with limit, order, start_after, and end_before query
// parameters, as well as date range filtering with start and end parameters.
func (r *AccountService) ListStatements(ctx context.Context, accountID string, query AccountListStatementsParams, opts ...option.RequestOption) (res *pagination.CursorIDStatements[AccountListStatementsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8"), option.WithResponseInto(&raw)}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("account/%s/statements", accountID)
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

// Retrieve a paginated list of monthly statements for a specific account. Supports
// cursor-based pagination with limit, order, start_after, and end_before query
// parameters, as well as date range filtering with start and end parameters.
func (r *AccountService) ListStatementsAutoPaging(ctx context.Context, accountID string, query AccountListStatementsParams, opts ...option.RequestOption) *pagination.CursorIDStatementsAutoPager[AccountListStatementsResponse] {
	return pagination.NewCursorIDStatementsAutoPager(r.ListStatements(ctx, accountID, query, opts...))
}

// Create a "request to send money" that will require approval based on your
// organization's approval policies.
func (r *AccountService) RequestSendMoney(ctx context.Context, accountID string, body AccountRequestSendMoneyParams, opts ...option.RequestOption) (res *SendMoneyApproval, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("account/%s/request-send-money", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get transaction by ID
func (r *AccountService) GetTransaction(ctx context.Context, transactionID string, query AccountGetTransactionParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if query.AccountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("account/%s/transaction/%s", query.AccountID, transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Account struct {
	// ID for a Mercury account.
	ID                string  `json:"id,required" format:"uuid"`
	AccountNumber     string  `json:"accountNumber,required"`
	AvailableBalance  float64 `json:"availableBalance,required"`
	CreatedAt         string  `json:"createdAt,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	CurrentBalance    float64 `json:"currentBalance,required"`
	DashboardLink     string  `json:"dashboardLink,required"`
	Kind              string  `json:"kind,required"`
	LegalBusinessName string  `json:"legalBusinessName,required"`
	Name              string  `json:"name,required"`
	RoutingNumber     string  `json:"routingNumber,required"`
	// Any of "active", "deleted", "pending", "archived".
	Status AccountStatus `json:"status,required"`
	// Any of "mercury", "external", "recipient".
	Type                   AccountType `json:"type,required"`
	CanReceiveTransactions bool        `json:"canReceiveTransactions,nullable"`
	Nickname               string      `json:"nickname,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccountNumber          respjson.Field
		AvailableBalance       respjson.Field
		CreatedAt              respjson.Field
		CurrentBalance         respjson.Field
		DashboardLink          respjson.Field
		Kind                   respjson.Field
		LegalBusinessName      respjson.Field
		Name                   respjson.Field
		RoutingNumber          respjson.Field
		Status                 respjson.Field
		Type                   respjson.Field
		CanReceiveTransactions respjson.Field
		Nickname               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountType string

const (
	AccountTypeMercury   AccountType = "mercury"
	AccountTypeExternal  AccountType = "external"
	AccountTypeRecipient AccountType = "recipient"
)

type AccountStatus string

const (
	AccountStatusActive   AccountStatus = "active"
	AccountStatusDeleted  AccountStatus = "deleted"
	AccountStatusPending  AccountStatus = "pending"
	AccountStatusArchived AccountStatus = "archived"
)

// Represents an expense category for transaction classification.
type CategoryData struct {
	// ID for the category
	ID string `json:"id,required" format:"uuid"`
	// The name of the category
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CategoryData) RawJSON() string { return r.JSON.raw }
func (r *CategoryData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MercuryCategory string

const (
	MercuryCategoryOther                MercuryCategory = "Other"
	MercuryCategoryAdvertising          MercuryCategory = "Advertising"
	MercuryCategoryAirlines             MercuryCategory = "Airlines"
	MercuryCategoryAlcoholAndBars       MercuryCategory = "AlcoholAndBars"
	MercuryCategoryBooksAndNewspaper    MercuryCategory = "BooksAndNewspaper"
	MercuryCategoryCarRental            MercuryCategory = "CarRental"
	MercuryCategoryCharity              MercuryCategory = "Charity"
	MercuryCategoryClothing             MercuryCategory = "Clothing"
	MercuryCategoryConferences          MercuryCategory = "Conferences"
	MercuryCategoryEducation            MercuryCategory = "Education"
	MercuryCategoryElectronics          MercuryCategory = "Electronics"
	MercuryCategoryEntertainment        MercuryCategory = "Entertainment"
	MercuryCategoryFacilitiesExpenses   MercuryCategory = "FacilitiesExpenses"
	MercuryCategoryFees                 MercuryCategory = "Fees"
	MercuryCategoryFoodDelivery         MercuryCategory = "FoodDelivery"
	MercuryCategoryFuelAndGas           MercuryCategory = "FuelAndGas"
	MercuryCategoryGambling             MercuryCategory = "Gambling"
	MercuryCategoryGovernmentServices   MercuryCategory = "GovernmentServices"
	MercuryCategoryGrocery              MercuryCategory = "Grocery"
	MercuryCategoryGroundTransportation MercuryCategory = "GroundTransportation"
	MercuryCategoryInsurance            MercuryCategory = "Insurance"
	MercuryCategoryInternetAndTelephone MercuryCategory = "InternetAndTelephone"
	MercuryCategoryLegal                MercuryCategory = "Legal"
	MercuryCategoryLodging              MercuryCategory = "Lodging"
	MercuryCategoryMedical              MercuryCategory = "Medical"
	MercuryCategoryMemberships          MercuryCategory = "Memberships"
	MercuryCategoryOfficeSupplies       MercuryCategory = "OfficeSupplies"
	MercuryCategoryOtherTravel          MercuryCategory = "OtherTravel"
	MercuryCategoryParking              MercuryCategory = "Parking"
	MercuryCategoryPolitical            MercuryCategory = "Political"
	MercuryCategoryProfessionalServices MercuryCategory = "ProfessionalServices"
	MercuryCategoryRestaurants          MercuryCategory = "Restaurants"
	MercuryCategoryRetail               MercuryCategory = "Retail"
	MercuryCategoryRideshareAndTaxis    MercuryCategory = "RideshareAndTaxis"
	MercuryCategoryShipping             MercuryCategory = "Shipping"
	MercuryCategorySoftware             MercuryCategory = "Software"
	MercuryCategoryTaxes                MercuryCategory = "Taxes"
	MercuryCategoryUtilities            MercuryCategory = "Utilities"
	MercuryCategoryVehicleExpenses      MercuryCategory = "VehicleExpenses"
)

// Extremely close to the internal type, but strips out potentially unwanted fields
type SendMoneyApproval struct {
	// ID for a Mercury account.
	AccountID string `json:"accountId,required" format:"uuid"`
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount,required"`
	// Any of "ach", "check", "domesticWire", "internationalWire".
	PaymentMethod SendMoneyPaymentMethod `json:"paymentMethod,required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId,required" format:"uuid"`
	RequestID   string `json:"requestId,required"`
	// Any of "pendingApproval", "approved", "rejected", "cancelled".
	Status SendMoneyApprovalStatus `json:"status,required"`
	Memo   string                  `json:"memo,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID     respjson.Field
		Amount        respjson.Field
		PaymentMethod respjson.Field
		RecipientID   respjson.Field
		RequestID     respjson.Field
		Status        respjson.Field
		Memo          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SendMoneyApproval) RawJSON() string { return r.JSON.raw }
func (r *SendMoneyApproval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SendMoneyApprovalStatus string

const (
	SendMoneyApprovalStatusPendingApproval SendMoneyApprovalStatus = "pendingApproval"
	SendMoneyApprovalStatusApproved        SendMoneyApprovalStatus = "approved"
	SendMoneyApprovalStatusRejected        SendMoneyApprovalStatus = "rejected"
	SendMoneyApprovalStatusCancelled       SendMoneyApprovalStatus = "cancelled"
)

type SendMoneyPaymentMethod string

const (
	SendMoneyPaymentMethodACH               SendMoneyPaymentMethod = "ach"
	SendMoneyPaymentMethodCheck             SendMoneyPaymentMethod = "check"
	SendMoneyPaymentMethodDomesticWire      SendMoneyPaymentMethod = "domesticWire"
	SendMoneyPaymentMethodInternationalWire SendMoneyPaymentMethod = "internationalWire"
)

type Transaction struct {
	// ID for this transaction
	ID string `json:"id,required" format:"uuid"`
	// ID for a Mercury account.
	AccountID                  string                  `json:"accountId,required" format:"uuid"`
	Amount                     float64                 `json:"amount,required"`
	Attachments                []TransactionAttachment `json:"attachments,required"`
	CompliantWithReceiptPolicy bool                    `json:"compliantWithReceiptPolicy,required"`
	// ID for a Mercury account.
	CounterpartyID        string `json:"counterpartyId,required" format:"uuid"`
	CounterpartyName      string `json:"counterpartyName,required"`
	CreatedAt             string `json:"createdAt,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	DashboardLink         string `json:"dashboardLink,required"`
	EstimatedDeliveryDate string `json:"estimatedDeliveryDate,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	HasGeneratedReceipt   bool   `json:"hasGeneratedReceipt,required"`
	// Any of "externalTransfer", "internalTransfer", "outgoingPayment",
	// "creditCardCredit", "creditCardTransaction", "debitCardCredit",
	// "debitCardTransaction", "cardInternationalTransactionFee",
	// "cardInternationalTransactionFeeRebate",
	// "cardInternationalTransactionFeeReversal",
	// "cardInternationalTransactionFeeRebateReversal", "incomingDomesticWire",
	// "checkDeposit", "incomingInternationalWire", "treasuryTransfer",
	// "currencyCloudReturn", "wireFee", "personalBankingSubscriptionFee",
	// "billingEngineSubscriptionFee", "expenseReimbursement", "exogenousWireDrawdown",
	// "other".
	Kind                TransactionKind                 `json:"kind,required"`
	RelatedTransactions []TransactionRelatedTransaction `json:"relatedTransactions,required"`
	// Any of "pending", "sent", "cancelled", "failed", "reversed", "blocked".
	Status          TransactionStatus `json:"status,required"`
	BankDescription string            `json:"bankDescription,nullable"`
	// Represents an expense category for transaction classification.
	CategoryData CategoryData `json:"categoryData,nullable"`
	// Present for check deposits and mailed checks; Nothing otherwise.
	CheckNumber          string `json:"checkNumber,nullable"`
	CounterpartyNickname string `json:"counterpartyNickname,nullable"`
	// ID for the credit statement period
	CreditAccountPeriodID string                          `json:"creditAccountPeriodId,nullable" format:"uuid"`
	CurrencyExchangeInfo  TransactionCurrencyExchangeInfo `json:"currencyExchangeInfo,nullable"`
	Details               TransactionDetails              `json:"details,nullable"`
	ExternalMemo          string                          `json:"externalMemo,nullable"`
	FailedAt              string                          `json:"failedAt,nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// ID for this transaction
	FeeID string `json:"feeId,nullable" format:"uuid"`
	// The name of the General Ledger (GL) code assigned to this transaction for
	// accounting categorization. GL codes act as "bins" that organize transactions
	// into accounting categories. This field is present when the transaction has been
	// categorized, either manually by a user, via an accounting integration sync, or
	// through auto-categorization rules. Nothing if the transaction has not been
	// assigned a GL code.
	GeneralLedgerCodeName string `json:"generalLedgerCodeName,nullable"`
	// Merchant information for card transactions
	Merchant TransactionMerchant `json:"merchant,nullable"`
	// Any of "Other", "Advertising", "Airlines", "AlcoholAndBars",
	// "BooksAndNewspaper", "CarRental", "Charity", "Clothing", "Conferences",
	// "Education", "Electronics", "Entertainment", "FacilitiesExpenses", "Fees",
	// "FoodDelivery", "FuelAndGas", "Gambling", "GovernmentServices", "Grocery",
	// "GroundTransportation", "Insurance", "InternetAndTelephone", "Legal", "Lodging",
	// "Medical", "Memberships", "OfficeSupplies", "OtherTravel", "Parking",
	// "Political", "ProfessionalServices", "Restaurants", "Retail",
	// "RideshareAndTaxis", "Shipping", "Software", "Taxes", "Utilities",
	// "VehicleExpenses".
	MercuryCategory  MercuryCategory `json:"mercuryCategory,nullable"`
	Note             string          `json:"note,nullable"`
	PostedAt         string          `json:"postedAt,nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	ReasonForFailure string          `json:"reasonForFailure,nullable"`
	RequestID        string          `json:"requestId,nullable"`
	// Present for transactions that have tracking numbers (e.g., RTP, ACH, wires);
	// Nothing otherwise.
	TrackingNumber string `json:"trackingNumber,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		AccountID                  respjson.Field
		Amount                     respjson.Field
		Attachments                respjson.Field
		CompliantWithReceiptPolicy respjson.Field
		CounterpartyID             respjson.Field
		CounterpartyName           respjson.Field
		CreatedAt                  respjson.Field
		DashboardLink              respjson.Field
		EstimatedDeliveryDate      respjson.Field
		HasGeneratedReceipt        respjson.Field
		Kind                       respjson.Field
		RelatedTransactions        respjson.Field
		Status                     respjson.Field
		BankDescription            respjson.Field
		CategoryData               respjson.Field
		CheckNumber                respjson.Field
		CounterpartyNickname       respjson.Field
		CreditAccountPeriodID      respjson.Field
		CurrencyExchangeInfo       respjson.Field
		Details                    respjson.Field
		ExternalMemo               respjson.Field
		FailedAt                   respjson.Field
		FeeID                      respjson.Field
		GeneralLedgerCodeName      respjson.Field
		Merchant                   respjson.Field
		MercuryCategory            respjson.Field
		Note                       respjson.Field
		PostedAt                   respjson.Field
		ReasonForFailure           respjson.Field
		RequestID                  respjson.Field
		TrackingNumber             respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Transaction) RawJSON() string { return r.JSON.raw }
func (r *Transaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionAttachment struct {
	// Any of "checkImage", "receipt", "other".
	AttachmentType string `json:"attachmentType,required"`
	FileName       string `json:"fileName,required"`
	URL            string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachmentType respjson.Field
		FileName       respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionAttachment) RawJSON() string { return r.JSON.raw }
func (r *TransactionAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionKind string

const (
	TransactionKindExternalTransfer                              TransactionKind = "externalTransfer"
	TransactionKindInternalTransfer                              TransactionKind = "internalTransfer"
	TransactionKindOutgoingPayment                               TransactionKind = "outgoingPayment"
	TransactionKindCreditCardCredit                              TransactionKind = "creditCardCredit"
	TransactionKindCreditCardTransaction                         TransactionKind = "creditCardTransaction"
	TransactionKindDebitCardCredit                               TransactionKind = "debitCardCredit"
	TransactionKindDebitCardTransaction                          TransactionKind = "debitCardTransaction"
	TransactionKindCardInternationalTransactionFee               TransactionKind = "cardInternationalTransactionFee"
	TransactionKindCardInternationalTransactionFeeRebate         TransactionKind = "cardInternationalTransactionFeeRebate"
	TransactionKindCardInternationalTransactionFeeReversal       TransactionKind = "cardInternationalTransactionFeeReversal"
	TransactionKindCardInternationalTransactionFeeRebateReversal TransactionKind = "cardInternationalTransactionFeeRebateReversal"
	TransactionKindIncomingDomesticWire                          TransactionKind = "incomingDomesticWire"
	TransactionKindCheckDeposit                                  TransactionKind = "checkDeposit"
	TransactionKindIncomingInternationalWire                     TransactionKind = "incomingInternationalWire"
	TransactionKindTreasuryTransfer                              TransactionKind = "treasuryTransfer"
	TransactionKindCurrencyCloudReturn                           TransactionKind = "currencyCloudReturn"
	TransactionKindWireFee                                       TransactionKind = "wireFee"
	TransactionKindPersonalBankingSubscriptionFee                TransactionKind = "personalBankingSubscriptionFee"
	TransactionKindBillingEngineSubscriptionFee                  TransactionKind = "billingEngineSubscriptionFee"
	TransactionKindExpenseReimbursement                          TransactionKind = "expenseReimbursement"
	TransactionKindExogenousWireDrawdown                         TransactionKind = "exogenousWireDrawdown"
	TransactionKindOther                                         TransactionKind = "other"
)

// A Public API version of RelatedTransactionData.
type TransactionRelatedTransaction struct {
	// ID for this transaction
	ID string `json:"id,required" format:"uuid"`
	// ID for a Mercury account.
	AccountID string  `json:"accountId,required" format:"uuid"`
	Amount    float64 `json:"amount,required"`
	// Any of "ProvisionalCreditReversalToMerchantRefund",
	// "MerchantRefundToProvisionalCreditReversal", "MerchantRefundToFraudulentCharge",
	// "FraudulentChargeToMerchantRefund", "PaymentRefundToFailedPayment",
	// "FailedPaymentToPaymentRefund", "GiftCompensationToOriginalTransaction",
	// "FeePaymentToOriginalTransaction", "OriginalTransactionToFeePayment",
	// "FeePaymentToFeeRebate", "FeeRebateToFeePayment", "FeePaymentToFeeReversal",
	// "FeeReversalToFeePayment", "FeeRebateToFeeRebateReversal",
	// "FeeRebateReversalToFeeRebate", "TreasurySplitLiquidation",
	// "ProvisionalCreditToOriginalCharge", "OriginalChargeToProvisionalCredit",
	// "FeeAtmReimbursementToAtmTransaction", "AtmTransactionToFeeAtmReimbursement",
	// "AtmTransactionToAtmReimbursementReversal",
	// "AtmReimbursementReversalToAtmTransaction", "ReturnToOriginalTransaction",
	// "OriginalTransactionToReturn", "ProvisionalCreditToReversal",
	// "ReversalToProvisionalCredit".
	RelationKind string `json:"relationKind,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		Amount       respjson.Field
		RelationKind respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionRelatedTransaction) RawJSON() string { return r.JSON.raw }
func (r *TransactionRelatedTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusSent      TransactionStatus = "sent"
	TransactionStatusCancelled TransactionStatus = "cancelled"
	TransactionStatusFailed    TransactionStatus = "failed"
	TransactionStatusReversed  TransactionStatus = "reversed"
	TransactionStatusBlocked   TransactionStatus = "blocked"
)

type TransactionCurrencyExchangeInfo struct {
	ConvertedFromAmount   float64 `json:"convertedFromAmount,required"`
	ConvertedFromCurrency string  `json:"convertedFromCurrency,required"`
	ConvertedToAmount     float64 `json:"convertedToAmount,required"`
	ConvertedToCurrency   string  `json:"convertedToCurrency,required"`
	// Exchange rate goes from "from currency" to "to currency" (ie from currency \*
	// exchange rate = to currency)
	ExchangeRate  float64 `json:"exchangeRate,required"`
	FeeAmount     float64 `json:"feeAmount,required"`
	FeePercentage float64 `json:"feePercentage,required"`
	// ID for this transaction
	FeeTransactionID string `json:"feeTransactionId,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConvertedFromAmount   respjson.Field
		ConvertedFromCurrency respjson.Field
		ConvertedToAmount     respjson.Field
		ConvertedToCurrency   respjson.Field
		ExchangeRate          respjson.Field
		FeeAmount             respjson.Field
		FeePercentage         respjson.Field
		FeeTransactionID      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionCurrencyExchangeInfo) RawJSON() string { return r.JSON.raw }
func (r *TransactionCurrencyExchangeInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionDetails struct {
	Address                      AddressData                      `json:"address,nullable"`
	CreditCardInfo               TransactionDetailsCreditCardInfo `json:"creditCardInfo,nullable"`
	DebitCardInfo                TransactionDetailsDebitCardInfo  `json:"debitCardInfo,nullable"`
	DomesticWireRoutingInfo      DomesticWireRoutingInfo          `json:"domesticWireRoutingInfo,nullable"`
	ElectronicRoutingInfo        ElectronicRoutingInfo            `json:"electronicRoutingInfo,nullable"`
	InternationalWireRoutingInfo InternationalWireRoutingInfo     `json:"internationalWireRoutingInfo,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address                      respjson.Field
		CreditCardInfo               respjson.Field
		DebitCardInfo                respjson.Field
		DomesticWireRoutingInfo      respjson.Field
		ElectronicRoutingInfo        respjson.Field
		InternationalWireRoutingInfo respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionDetails) RawJSON() string { return r.JSON.raw }
func (r *TransactionDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionDetailsCreditCardInfo struct {
	ID            string `json:"id,required" format:"uuid"`
	PaymentMethod string `json:"paymentMethod,required"`
	Email         string `json:"email,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		PaymentMethod respjson.Field
		Email         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionDetailsCreditCardInfo) RawJSON() string { return r.JSON.raw }
func (r *TransactionDetailsCreditCardInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionDetailsDebitCardInfo struct {
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionDetailsDebitCardInfo) RawJSON() string { return r.JSON.raw }
func (r *TransactionDetailsDebitCardInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Merchant information for card transactions
type TransactionMerchant struct {
	// Merchant ID for card transactions
	ID string `json:"id,nullable"`
	// Mercury category for the merchant (e.g., "Restaurants", "Software")
	//
	// Any of "Other", "Advertising", "Airlines", "AlcoholAndBars",
	// "BooksAndNewspaper", "CarRental", "Charity", "Clothing", "Conferences",
	// "Education", "Electronics", "Entertainment", "FacilitiesExpenses", "Fees",
	// "FoodDelivery", "FuelAndGas", "Gambling", "GovernmentServices", "Grocery",
	// "GroundTransportation", "Insurance", "InternetAndTelephone", "Legal", "Lodging",
	// "Medical", "Memberships", "OfficeSupplies", "OtherTravel", "Parking",
	// "Political", "ProfessionalServices", "Restaurants", "Retail",
	// "RideshareAndTaxis", "Shipping", "Software", "Taxes", "Utilities",
	// "VehicleExpenses".
	Category MercuryCategory `json:"category,nullable"`
	// 4-digit merchant category code (MCC) for card transactions
	CategoryCode string `json:"categoryCode,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Category     respjson.Field
		CategoryCode respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionMerchant) RawJSON() string { return r.JSON.raw }
func (r *TransactionMerchant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListCardsResponse struct {
	Cards []AccountListCardsResponseCard `json:"cards,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cards       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountListCardsResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountListCardsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListCardsResponseCard struct {
	CardID         string `json:"cardId,required"`
	CreatedAt      string `json:"createdAt,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	LastFourDigits string `json:"lastFourDigits,required"`
	NameOnCard     string `json:"nameOnCard,required"`
	// Any of "visa", "mastercard".
	Network string `json:"network,required"`
	// Any of "active", "frozen", "cancelled", "inactive", "expired", "suspended".
	Status string `json:"status,required"`
	// Any of "inactive", "active", "paused".
	PhysicalCardStatus string `json:"physicalCardStatus,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CardID             respjson.Field
		CreatedAt          respjson.Field
		LastFourDigits     respjson.Field
		NameOnCard         respjson.Field
		Network            respjson.Field
		Status             respjson.Field
		PhysicalCardStatus respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountListCardsResponseCard) RawJSON() string { return r.JSON.raw }
func (r *AccountListCardsResponseCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListStatementsResponse struct {
	// ID for the account statement
	ID                  string  `json:"id,required" format:"uuid"`
	AccountNumber       string  `json:"accountNumber,required"`
	CompanyLegalAddress Address `json:"companyLegalAddress,required"`
	CompanyLegalName    string  `json:"companyLegalName,required"`
	DownloadURL         string  `json:"downloadUrl,required"`
	EndDate             string  `json:"endDate,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// A dollar amount
	EndingBalance float64                                    `json:"endingBalance,required"`
	RoutingNumber string                                     `json:"routingNumber,required"`
	StartDate     string                                     `json:"startDate,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	Transactions  []AccountListStatementsResponseTransaction `json:"transactions,required"`
	Ein           string                                     `json:"ein,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AccountNumber       respjson.Field
		CompanyLegalAddress respjson.Field
		CompanyLegalName    respjson.Field
		DownloadURL         respjson.Field
		EndDate             respjson.Field
		EndingBalance       respjson.Field
		RoutingNumber       respjson.Field
		StartDate           respjson.Field
		Transactions        respjson.Field
		Ein                 respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountListStatementsResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountListStatementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListStatementsResponseTransaction struct {
	// ID for this transaction
	ID        string `json:"id,required" format:"uuid"`
	CreatedAt string `json:"createdAt,required" format:"yyyy-mm-ddThh:MM:ssZ"`
	PostedAt  string `json:"postedAt,nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		PostedAt    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountListStatementsResponseTransaction) RawJSON() string { return r.JSON.raw }
func (r *AccountListStatementsResponseTransaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListParams struct {
	// The ID of the account to end the page before (exclusive). When provided, results
	// will end just before this ID and work backwards. Use this for reverse pagination
	// or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the account to start the page after (exclusive). When provided,
	// results will begin with the account immediately following this ID. Use this for
	// standard forward pagination to get the next page of results. Cannot be combined
	// with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order AccountListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type AccountListParamsOrder string

const (
	AccountListParamsOrderAsc  AccountListParamsOrder = "asc"
	AccountListParamsOrderDesc AccountListParamsOrder = "desc"
)

type AccountListStatementsParams struct {
	// Filter statements where the period start date is on or before this date. If the
	// date is in the future, defaults to the current date. Format: YYYY-MM-DD
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// The ID of the statement to end the page before (exclusive). When provided,
	// results will end just before this ID and work backwards. Use this for reverse
	// pagination or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter statements where the period start date is on or after this date. Format:
	// YYYY-MM-DD
	Start param.Opt[string] `query:"start,omitzero" json:"-"`
	// The ID of the statement to start the page after (exclusive). When provided,
	// results will begin with the statement immediately following this ID. Use this
	// for standard forward pagination to get the next page of results. Cannot be
	// combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'
	//
	// Any of "asc", "desc".
	Order AccountListStatementsParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountListStatementsParams]'s query parameters as
// `url.Values`.
func (r AccountListStatementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'
type AccountListStatementsParamsOrder string

const (
	AccountListStatementsParamsOrderAsc  AccountListStatementsParamsOrder = "asc"
	AccountListStatementsParamsOrderDesc AccountListStatementsParamsOrder = "desc"
)

type AccountRequestSendMoneyParams struct {
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount,required"`
	// Unique string identifying the transaction
	IdempotencyKey string `json:"idempotencyKey,required"`
	// Payment method to use.
	//
	// Any of "ach", "check", "domesticWire", "internationalWire".
	PaymentMethod SendMoneyPaymentMethod `json:"paymentMethod,omitzero,required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId,required" format:"uuid"`
	// Optional external memo
	ExternalMemo param.Opt[string] `json:"externalMemo,omitzero"`
	// Optional note
	Note param.Opt[string] `json:"note,omitzero"`
	paramObj
}

func (r AccountRequestSendMoneyParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountRequestSendMoneyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountRequestSendMoneyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetTransactionParams struct {
	// ID for a Mercury account.
	AccountID string `path:"accountId,required" format:"uuid" json:"-"`
	paramObj
}
