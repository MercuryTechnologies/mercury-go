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
func (r *AccountService) ListStatements(ctx context.Context, accountID string, query AccountListStatementsParams, opts ...option.RequestOption) (res *pagination.CursorIDAccountStatements[AccountListStatementsResponse], err error) {
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
func (r *AccountService) ListStatementsAutoPaging(ctx context.Context, accountID string, query AccountListStatementsParams, opts ...option.RequestOption) *pagination.CursorIDAccountStatementsAutoPager[AccountListStatementsResponse] {
	return pagination.NewCursorIDAccountStatementsAutoPager(r.ListStatements(ctx, accountID, query, opts...))
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
	ID                string  `json:"id" api:"required" format:"uuid"`
	AccountNumber     string  `json:"accountNumber" api:"required"`
	AvailableBalance  float64 `json:"availableBalance" api:"required"`
	CreatedAt         string  `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	CurrentBalance    float64 `json:"currentBalance" api:"required"`
	DashboardLink     string  `json:"dashboardLink" api:"required"`
	Kind              string  `json:"kind" api:"required"`
	LegalBusinessName string  `json:"legalBusinessName" api:"required"`
	Name              string  `json:"name" api:"required"`
	RoutingNumber     string  `json:"routingNumber" api:"required"`
	// Any of "active", "deleted", "pending", "archived".
	Status AccountStatus `json:"status" api:"required"`
	// Any of "mercury", "external", "recipient".
	Type                   AccountType `json:"type" api:"required"`
	CanReceiveTransactions bool        `json:"canReceiveTransactions" api:"nullable"`
	Nickname               string      `json:"nickname" api:"nullable"`
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
	ID string `json:"id" api:"required" format:"uuid"`
	// The name of the category
	Name string `json:"name" api:"required"`
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
	AccountID string `json:"accountId" api:"required" format:"uuid"`
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount" api:"required"`
	// Any of "ach", "check", "domesticWire", "internationalWire".
	PaymentMethod SendMoneyPaymentMethod `json:"paymentMethod" api:"required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId" api:"required" format:"uuid"`
	RequestID   string `json:"requestId" api:"required"`
	// Any of "pendingApproval", "approved", "rejected", "cancelled".
	Status SendMoneyApprovalStatus `json:"status" api:"required"`
	Memo   string                  `json:"memo" api:"nullable"`
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
	ID string `json:"id" api:"required" format:"uuid"`
	// ID for a Mercury account.
	AccountID                  string                  `json:"accountId" api:"required" format:"uuid"`
	Amount                     float64                 `json:"amount" api:"required"`
	Attachments                []TransactionAttachment `json:"attachments" api:"required"`
	CompliantWithReceiptPolicy bool                    `json:"compliantWithReceiptPolicy" api:"required"`
	// ID for a Mercury account.
	CounterpartyID        string `json:"counterpartyId" api:"required" format:"uuid"`
	CounterpartyName      string `json:"counterpartyName" api:"required"`
	CreatedAt             string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	DashboardLink         string `json:"dashboardLink" api:"required"`
	EstimatedDeliveryDate string `json:"estimatedDeliveryDate" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	HasGeneratedReceipt   bool   `json:"hasGeneratedReceipt" api:"required"`
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
	Kind                TransactionKind                 `json:"kind" api:"required"`
	RelatedTransactions []TransactionRelatedTransaction `json:"relatedTransactions" api:"required"`
	// Any of "pending", "sent", "cancelled", "failed", "reversed", "blocked".
	Status          TransactionStatus `json:"status" api:"required"`
	BankDescription string            `json:"bankDescription" api:"nullable"`
	// Represents an expense category for transaction classification.
	CategoryData CategoryData `json:"categoryData" api:"nullable"`
	// Present for check deposits and mailed checks; Nothing otherwise.
	CheckNumber          string `json:"checkNumber" api:"nullable"`
	CounterpartyNickname string `json:"counterpartyNickname" api:"nullable"`
	// ID for the credit statement period
	CreditAccountPeriodID string                          `json:"creditAccountPeriodId" api:"nullable" format:"uuid"`
	CurrencyExchangeInfo  TransactionCurrencyExchangeInfo `json:"currencyExchangeInfo" api:"nullable"`
	Details               TransactionDetails              `json:"details" api:"nullable"`
	ExternalMemo          string                          `json:"externalMemo" api:"nullable"`
	FailedAt              string                          `json:"failedAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// ID for this transaction
	FeeID string `json:"feeId" api:"nullable" format:"uuid"`
	// The name of the General Ledger (GL) code assigned to this transaction for
	// accounting categorization. GL codes act as "bins" that organize transactions
	// into accounting categories. This field is present when the transaction has been
	// categorized, either manually by a user, via an accounting integration sync, or
	// through auto-categorization rules. Nothing if the transaction has not been
	// assigned a GL code.
	GeneralLedgerCodeName string `json:"generalLedgerCodeName" api:"nullable"`
	// Merchant information for card transactions
	Merchant TransactionMerchant `json:"merchant" api:"nullable"`
	// Any of "Other", "Advertising", "Airlines", "AlcoholAndBars",
	// "BooksAndNewspaper", "CarRental", "Charity", "Clothing", "Conferences",
	// "Education", "Electronics", "Entertainment", "FacilitiesExpenses", "Fees",
	// "FoodDelivery", "FuelAndGas", "Gambling", "GovernmentServices", "Grocery",
	// "GroundTransportation", "Insurance", "InternetAndTelephone", "Legal", "Lodging",
	// "Medical", "Memberships", "OfficeSupplies", "OtherTravel", "Parking",
	// "Political", "ProfessionalServices", "Restaurants", "Retail",
	// "RideshareAndTaxis", "Shipping", "Software", "Taxes", "Utilities",
	// "VehicleExpenses".
	MercuryCategory  MercuryCategory `json:"mercuryCategory" api:"nullable"`
	Note             string          `json:"note" api:"nullable"`
	PostedAt         string          `json:"postedAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	ReasonForFailure string          `json:"reasonForFailure" api:"nullable"`
	RequestID        string          `json:"requestId" api:"nullable"`
	// Present for transactions that have tracking numbers (e.g., RTP, ACH, wires);
	// Nothing otherwise.
	TrackingNumber string `json:"trackingNumber" api:"nullable"`
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
	AttachmentType string `json:"attachmentType" api:"required"`
	FileName       string `json:"fileName" api:"required"`
	URL            string `json:"url" api:"required"`
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
	ID string `json:"id" api:"required" format:"uuid"`
	// ID for a Mercury account.
	AccountID string  `json:"accountId" api:"required" format:"uuid"`
	Amount    float64 `json:"amount" api:"required"`
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
	RelationKind string `json:"relationKind" api:"required"`
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
	ConvertedFromAmount   float64 `json:"convertedFromAmount" api:"required"`
	ConvertedFromCurrency string  `json:"convertedFromCurrency" api:"required"`
	ConvertedToAmount     float64 `json:"convertedToAmount" api:"required"`
	ConvertedToCurrency   string  `json:"convertedToCurrency" api:"required"`
	// Exchange rate goes from "from currency" to "to currency" (ie from currency \*
	// exchange rate = to currency)
	ExchangeRate  float64 `json:"exchangeRate" api:"required"`
	FeeAmount     float64 `json:"feeAmount" api:"required"`
	FeePercentage float64 `json:"feePercentage" api:"required"`
	// ID for this transaction
	FeeTransactionID string `json:"feeTransactionId" api:"nullable" format:"uuid"`
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
	Address                      AddressData                      `json:"address" api:"nullable"`
	CreditCardInfo               TransactionDetailsCreditCardInfo `json:"creditCardInfo" api:"nullable"`
	DebitCardInfo                TransactionDetailsDebitCardInfo  `json:"debitCardInfo" api:"nullable"`
	DomesticWireRoutingInfo      DomesticWireRoutingInfo          `json:"domesticWireRoutingInfo" api:"nullable"`
	ElectronicRoutingInfo        ElectronicRoutingInfo            `json:"electronicRoutingInfo" api:"nullable"`
	InternationalWireRoutingInfo InternationalWireRoutingInfo     `json:"internationalWireRoutingInfo" api:"nullable"`
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
	ID            string `json:"id" api:"required" format:"uuid"`
	PaymentMethod string `json:"paymentMethod" api:"required"`
	Email         string `json:"email" api:"nullable"`
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
	ID string `json:"id" api:"required" format:"uuid"`
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
	ID string `json:"id" api:"nullable"`
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
	Category MercuryCategory `json:"category" api:"nullable"`
	// 4-digit merchant category code (MCC) for card transactions
	CategoryCode string `json:"categoryCode" api:"nullable"`
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
	Cards []AccountListCardsResponseCard `json:"cards" api:"required"`
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
	CardID         string `json:"cardId" api:"required"`
	CreatedAt      string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	LastFourDigits string `json:"lastFourDigits" api:"required"`
	NameOnCard     string `json:"nameOnCard" api:"required"`
	// Any of "visa", "mastercard".
	Network string `json:"network" api:"required"`
	// Any of "active", "frozen", "cancelled", "inactive", "expired", "suspended".
	Status string `json:"status" api:"required"`
	// Any of "inactive", "active", "paused".
	PhysicalCardStatus string `json:"physicalCardStatus" api:"nullable"`
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
	ID                  string  `json:"id" api:"required" format:"uuid"`
	AccountNumber       string  `json:"accountNumber" api:"required"`
	CompanyLegalAddress Address `json:"companyLegalAddress" api:"required"`
	CompanyLegalName    string  `json:"companyLegalName" api:"required"`
	DownloadURL         string  `json:"downloadUrl" api:"required"`
	EndDate             string  `json:"endDate" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// A dollar amount
	EndingBalance float64                                    `json:"endingBalance" api:"required"`
	RoutingNumber string                                     `json:"routingNumber" api:"required"`
	StartDate     string                                     `json:"startDate" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	Transactions  []AccountListStatementsResponseTransaction `json:"transactions" api:"required"`
	Ein           string                                     `json:"ein" api:"nullable"`
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
	ID        string `json:"id" api:"required" format:"uuid"`
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	PostedAt  string `json:"postedAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
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
	Amount float64 `json:"amount" api:"required"`
	// Unique string identifying the transaction
	IdempotencyKey string `json:"idempotencyKey" api:"required"`
	// Payment method to use.
	//
	// Any of "ach", "check", "domesticWire", "internationalWire".
	PaymentMethod SendMoneyPaymentMethod `json:"paymentMethod,omitzero" api:"required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId" api:"required" format:"uuid"`
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
	AccountID string `path:"accountId" api:"required" format:"uuid" json:"-"`
	paramObj
}
