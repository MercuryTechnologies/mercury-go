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

// Manage transactions
//
// TransactionService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	Options []option.RequestOption
	// Manage transactions
	Attachments TransactionAttachmentService
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r TransactionService) {
	r = TransactionService{}
	r.Options = opts
	r.Attachments = NewTransactionAttachmentService(opts...)
	return
}

// Update the note and/or category of an existing transaction. Use null values to
// clear existing data.
func (r *TransactionService) Update(ctx context.Context, transactionID string, body TransactionUpdateParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("transaction/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of all transactions across all accounts. Supports
// advanced filtering by date ranges, status, categories, and cursor-based
// pagination.
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *pagination.CursorIDTransactions[Transaction], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "transactions"
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

// Retrieve a paginated list of all transactions across all accounts. Supports
// advanced filtering by date ranges, status, categories, and cursor-based
// pagination.
func (r *TransactionService) ListAutoPaging(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) *pagination.CursorIDTransactionsAutoPager[Transaction] {
	return pagination.NewCursorIDTransactionsAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a single transaction by its ID. Returns full transaction details
// including attachments, check images, and related metadata.
func (r *TransactionService) Get(ctx context.Context, transactionID string, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("transaction/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CurrencyExchangeInfo struct {
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
func (r CurrencyExchangeInfo) RawJSON() string { return r.JSON.raw }
func (r *CurrencyExchangeInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A GL code allocation on a transaction — a GL code name paired with the amount
// allocated to it. When a transaction is fully categorized, the amounts across all
// allocations sum to the transaction total.
type GlAllocation struct {
	// The amount allocated to this GL code
	Amount float64 `json:"amount" api:"required"`
	// The name of the GL code from the connected accounting integration
	GlCodeName string `json:"glCodeName" api:"required"`
	// Optional user-provided description for this allocation
	Description string `json:"description" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		GlCodeName  respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GlAllocation) RawJSON() string { return r.JSON.raw }
func (r *GlAllocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Merchant information for card transactions
type MerchantData struct {
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
func (r MerchantData) RawJSON() string { return r.JSON.raw }
func (r *MerchantData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Public API version of RelatedTransactionData.
type RelatedTransactionData struct {
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
	RelationKind RelatedTransactionDataRelationKind `json:"relationKind" api:"required"`
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
func (r RelatedTransactionData) RawJSON() string { return r.JSON.raw }
func (r *RelatedTransactionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RelatedTransactionDataRelationKind string

const (
	RelatedTransactionDataRelationKindProvisionalCreditReversalToMerchantRefund RelatedTransactionDataRelationKind = "ProvisionalCreditReversalToMerchantRefund"
	RelatedTransactionDataRelationKindMerchantRefundToProvisionalCreditReversal RelatedTransactionDataRelationKind = "MerchantRefundToProvisionalCreditReversal"
	RelatedTransactionDataRelationKindMerchantRefundToFraudulentCharge          RelatedTransactionDataRelationKind = "MerchantRefundToFraudulentCharge"
	RelatedTransactionDataRelationKindFraudulentChargeToMerchantRefund          RelatedTransactionDataRelationKind = "FraudulentChargeToMerchantRefund"
	RelatedTransactionDataRelationKindPaymentRefundToFailedPayment              RelatedTransactionDataRelationKind = "PaymentRefundToFailedPayment"
	RelatedTransactionDataRelationKindFailedPaymentToPaymentRefund              RelatedTransactionDataRelationKind = "FailedPaymentToPaymentRefund"
	RelatedTransactionDataRelationKindGiftCompensationToOriginalTransaction     RelatedTransactionDataRelationKind = "GiftCompensationToOriginalTransaction"
	RelatedTransactionDataRelationKindFeePaymentToOriginalTransaction           RelatedTransactionDataRelationKind = "FeePaymentToOriginalTransaction"
	RelatedTransactionDataRelationKindOriginalTransactionToFeePayment           RelatedTransactionDataRelationKind = "OriginalTransactionToFeePayment"
	RelatedTransactionDataRelationKindFeePaymentToFeeRebate                     RelatedTransactionDataRelationKind = "FeePaymentToFeeRebate"
	RelatedTransactionDataRelationKindFeeRebateToFeePayment                     RelatedTransactionDataRelationKind = "FeeRebateToFeePayment"
	RelatedTransactionDataRelationKindFeePaymentToFeeReversal                   RelatedTransactionDataRelationKind = "FeePaymentToFeeReversal"
	RelatedTransactionDataRelationKindFeeReversalToFeePayment                   RelatedTransactionDataRelationKind = "FeeReversalToFeePayment"
	RelatedTransactionDataRelationKindFeeRebateToFeeRebateReversal              RelatedTransactionDataRelationKind = "FeeRebateToFeeRebateReversal"
	RelatedTransactionDataRelationKindFeeRebateReversalToFeeRebate              RelatedTransactionDataRelationKind = "FeeRebateReversalToFeeRebate"
	RelatedTransactionDataRelationKindTreasurySplitLiquidation                  RelatedTransactionDataRelationKind = "TreasurySplitLiquidation"
	RelatedTransactionDataRelationKindProvisionalCreditToOriginalCharge         RelatedTransactionDataRelationKind = "ProvisionalCreditToOriginalCharge"
	RelatedTransactionDataRelationKindOriginalChargeToProvisionalCredit         RelatedTransactionDataRelationKind = "OriginalChargeToProvisionalCredit"
	RelatedTransactionDataRelationKindFeeAtmReimbursementToAtmTransaction       RelatedTransactionDataRelationKind = "FeeAtmReimbursementToAtmTransaction"
	RelatedTransactionDataRelationKindAtmTransactionToFeeAtmReimbursement       RelatedTransactionDataRelationKind = "AtmTransactionToFeeAtmReimbursement"
	RelatedTransactionDataRelationKindAtmTransactionToAtmReimbursementReversal  RelatedTransactionDataRelationKind = "AtmTransactionToAtmReimbursementReversal"
	RelatedTransactionDataRelationKindAtmReimbursementReversalToAtmTransaction  RelatedTransactionDataRelationKind = "AtmReimbursementReversalToAtmTransaction"
	RelatedTransactionDataRelationKindReturnToOriginalTransaction               RelatedTransactionDataRelationKind = "ReturnToOriginalTransaction"
	RelatedTransactionDataRelationKindOriginalTransactionToReturn               RelatedTransactionDataRelationKind = "OriginalTransactionToReturn"
	RelatedTransactionDataRelationKindProvisionalCreditToReversal               RelatedTransactionDataRelationKind = "ProvisionalCreditToReversal"
	RelatedTransactionDataRelationKindReversalToProvisionalCredit               RelatedTransactionDataRelationKind = "ReversalToProvisionalCredit"
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
	// GL code allocations assigned to this transaction via a connected accounting
	// software integration (e.g. QuickBooks, Xero, NetSuite). Each allocation has a GL
	// code name and the amount allocated to it; amounts sum to the transaction total
	// when the transaction is fully categorized. Empty if no GL codes have been
	// assigned. Distinct from Mercury custom categories (see transactionCategoryData).
	GlAllocations       []GlAllocation `json:"glAllocations" api:"required"`
	HasGeneratedReceipt bool           `json:"hasGeneratedReceipt" api:"required"`
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
	Kind                TransactionKind          `json:"kind" api:"required"`
	RelatedTransactions []RelatedTransactionData `json:"relatedTransactions" api:"required"`
	// Any of "pending", "sent", "cancelled", "failed", "reversed", "blocked".
	Status          TransactionStatus `json:"status" api:"required"`
	BankDescription string            `json:"bankDescription" api:"nullable"`
	// Represents an expense category for transaction classification.
	CategoryData CategoryData `json:"categoryData" api:"nullable"`
	// Present for check deposits and mailed checks; Nothing otherwise.
	CheckNumber          string `json:"checkNumber" api:"nullable"`
	CounterpartyNickname string `json:"counterpartyNickname" api:"nullable"`
	// ID for the credit statement period
	CreditAccountPeriodID string                `json:"creditAccountPeriodId" api:"nullable" format:"uuid"`
	CurrencyExchangeInfo  CurrencyExchangeInfo  `json:"currencyExchangeInfo" api:"nullable"`
	Details               TransactionMethodData `json:"details" api:"nullable"`
	ExternalMemo          string                `json:"externalMemo" api:"nullable"`
	FailedAt              string                `json:"failedAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// ID for this transaction
	FeeID string `json:"feeId" api:"nullable" format:"uuid"`
	// Deprecated: use transactionGlAllocations instead. This field does not reflect GL
	// codes assigned via Mercury auto-categorization rules. Preserved for backwards
	// compatibility.
	GeneralLedgerCodeName string `json:"generalLedgerCodeName" api:"nullable"`
	// Merchant information for card transactions
	Merchant MerchantData `json:"merchant" api:"nullable"`
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
		GlAllocations              respjson.Field
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

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusSent      TransactionStatus = "sent"
	TransactionStatusCancelled TransactionStatus = "cancelled"
	TransactionStatusFailed    TransactionStatus = "failed"
	TransactionStatusReversed  TransactionStatus = "reversed"
	TransactionStatusBlocked   TransactionStatus = "blocked"
)

type TransactionMethodData struct {
	Address                      AddressData                         `json:"address" api:"nullable"`
	CreditCardInfo               TransactionMethodDataCreditCardInfo `json:"creditCardInfo" api:"nullable"`
	DebitCardInfo                TransactionMethodDataDebitCardInfo  `json:"debitCardInfo" api:"nullable"`
	DomesticWireRoutingInfo      DomesticWireRoutingInfo             `json:"domesticWireRoutingInfo" api:"nullable"`
	ElectronicRoutingInfo        ElectronicRoutingInfo               `json:"electronicRoutingInfo" api:"nullable"`
	InternationalWireRoutingInfo InternationalWireRoutingInfo        `json:"internationalWireRoutingInfo" api:"nullable"`
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
func (r TransactionMethodData) RawJSON() string { return r.JSON.raw }
func (r *TransactionMethodData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionMethodDataCreditCardInfo struct {
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
func (r TransactionMethodDataCreditCardInfo) RawJSON() string { return r.JSON.raw }
func (r *TransactionMethodDataCreditCardInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionMethodDataDebitCardInfo struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionMethodDataDebitCardInfo) RawJSON() string { return r.JSON.raw }
func (r *TransactionMethodDataDebitCardInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionUpdateParams struct {
	// Note update action. Omit field to keep current note, send null or empty string
	// to clear note, send text to set note.
	Note param.Opt[string] `json:"note,omitzero" api:"required"`
	// ID for the category
	CategoryID string `json:"categoryId" api:"required" format:"uuid"`
	paramObj
}

func (r TransactionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListParams struct {
	// UUID of a custom category. Can be returned from /categories endpoint.
	CategoryID param.Opt[string] `query:"categoryId,omitzero" json:"-"`
	// Latest createdAt date to filter for. If it’s not provided, it defaults to
	// current day. Format: YYYY-MM-DD or an ISO 8601 string. Please note that your
	// Mercury transactions on your Dashboard might have their postedAt date displayed,
	// as opposed to createdAt
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// The ID of the transaction to end the page before (exclusive). When provided,
	// results will end just before this ID and work backwards. Use this for reverse
	// pagination or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Name of mercuryCategory you want to filter on. Merchant Type in the UI.
	MercuryCategory param.Opt[string] `query:"mercuryCategory,omitzero" json:"-"`
	// Latest postedAt date to filter for. Format: YYYY-MM-DD or an ISO 8601 string
	PostedEnd param.Opt[string] `query:"postedEnd,omitzero" json:"-"`
	// Earliest postedAt date to filter for. Format: YYYY-MM-DD or an ISO 8601 string
	PostedStart param.Opt[string] `query:"postedStart,omitzero" json:"-"`
	// Search term to look for in transaction descriptions.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Earliest createdAt date to filter for. If not provided, it defaults to the date
	// of your first transaction. Format: YYYY-MM-DD or an ISO 8601 string. Please note
	// that your Mercury transactions on your Dashboard might have their postedAt date
	// displayed, as opposed to createdAt
	Start param.Opt[string] `query:"start,omitzero" json:"-"`
	// The ID of the transaction to start the page after (exclusive). When provided,
	// results will begin with the transaction immediately following this ID. Use this
	// for standard forward pagination to get the next page of results. Cannot be
	// combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// The ID of the resource to start the page at (inclusive). When provided, results
	// will begin with and include the resource with this ID. Use this to retrieve a
	// specific page when you know the exact starting point. Cannot be combined with
	// start_after or end_before.
	StartAt   param.Opt[string] `query:"start_at,omitzero" json:"-"`
	AccountID []string          `query:"accountId,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order TransactionListParamsOrder `query:"order,omitzero" json:"-"`
	// Any of "pending", "sent", "cancelled", "failed", "reversed", "blocked".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type TransactionListParamsOrder string

const (
	TransactionListParamsOrderAsc  TransactionListParamsOrder = "asc"
	TransactionListParamsOrderDesc TransactionListParamsOrder = "desc"
)
