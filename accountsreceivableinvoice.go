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
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/pagination"
	"github.com/stainless-sdks/mercury-go/packages/param"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// Manage invoices
//
// AccountsReceivableInvoiceService contains methods and other services that help
// with interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountsReceivableInvoiceService] method instead.
type AccountsReceivableInvoiceService struct {
	Options []option.RequestOption
}

// NewAccountsReceivableInvoiceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountsReceivableInvoiceService(opts ...option.RequestOption) (r AccountsReceivableInvoiceService) {
	r = AccountsReceivableInvoiceService{}
	r.Options = opts
	return
}

// Create a new invoice for the organization
func (r *AccountsReceivableInvoiceService) New(ctx context.Context, body AccountsReceivableInvoiceNewParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ar/invoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve details of an invoice by its ID
func (r *AccountsReceivableInvoiceService) Get(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoiceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("ar/invoices/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing invoice
func (r *AccountsReceivableInvoiceService) Update(ctx context.Context, invoiceID string, body AccountsReceivableInvoiceUpdateParams, opts ...option.RequestOption) (res *Invoice, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoiceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("ar/invoices/%s", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of invoices. Supports cursor-based pagination with
// limit, order, start_after, and end_before query parameters.
func (r *AccountsReceivableInvoiceService) List(ctx context.Context, query AccountsReceivableInvoiceListParams, opts ...option.RequestOption) (res *pagination.CursorIDArInvoices[AccountsReceivableInvoiceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ar/invoices"
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

// Retrieve a paginated list of invoices. Supports cursor-based pagination with
// limit, order, start_after, and end_before query parameters.
func (r *AccountsReceivableInvoiceService) ListAutoPaging(ctx context.Context, query AccountsReceivableInvoiceListParams, opts ...option.RequestOption) *pagination.CursorIDArInvoicesAutoPager[AccountsReceivableInvoiceListResponse] {
	return pagination.NewCursorIDArInvoicesAutoPager(r.List(ctx, query, opts...))
}

// Cancel an invoice. This action cannot be undone.
func (r *AccountsReceivableInvoiceService) Cancel(ctx context.Context, invoiceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if invoiceID == "" {
		err = errors.New("missing required invoiceId parameter")
		return err
	}
	path := fmt.Sprintf("ar/invoices/%s/cancel", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Downloads a PDF file for the specified invoice. The response includes a
// Content-Disposition header set to 'attachment' with the filename.
func (r *AccountsReceivableInvoiceService) DownloadPdf(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if invoiceID == "" {
		err = errors.New("missing required invoiceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("ar/invoices/%s/pdf", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a list of all attachments for a specific invoice
func (r *AccountsReceivableInvoiceService) ListAttachments(ctx context.Context, invoiceID string, opts ...option.RequestOption) (res *AccountsReceivableInvoiceListAttachmentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if invoiceID == "" {
		err = errors.New("missing required invoiceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("ar/invoices/%s/attachments", invoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The response type for an invoice in the api.
type Invoice struct {
	// ID for the invoice.
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether or not the invoice can be paid via ach debit.
	ACHDebitEnabled bool `json:"achDebitEnabled" api:"required"`
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount" api:"required"`
	// Emails to be CCed on invoice notifications/reminders.
	CcEmails []string `json:"ccEmails" api:"required"`
	// The timestamp when the invoice was created.
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Whether or not the invoice can be paid via credit card. Requires stripe to be
	// setup for the Mercury account.
	CreditCardEnabled bool `json:"creditCardEnabled" api:"required"`
	// The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint
	// to list your customers and find the corresponding id, or create a new customer
	// first.
	CustomerID string `json:"customerId" api:"required" format:"uuid"`
	// ID for a Mercury account.
	DestinationAccountID string `json:"destinationAccountId" api:"required" format:"uuid"`
	// The due date the invoice should be paid by.
	DueDate time.Time `json:"dueDate" api:"required" format:"date"`
	// The date of the invoice, set by the invoice creator and likely to be context
	// specific to the type of transaction. i.e. it could be a date a service was
	// performed, it does not need to be the date the invoice was created.
	InvoiceDate time.Time `json:"invoiceDate" api:"required" format:"date"`
	// The payer facing invoice number/identifier.
	InvoiceNumber string `json:"invoiceNumber" api:"required"`
	// The line items for the invoice.
	LineItems []LineItemData `json:"lineItems" api:"required"`
	// Public slug for an invoice. Used to construct the pay page URL as well as the
	// URL to retrieve the PDF of the invoice.
	Slug string `json:"slug" api:"required"`
	// The status of the invoice.
	//
	// Any of "Unpaid", "Paid", "Cancelled", "Processing".
	Status PaymentLinkStatus `json:"status" api:"required"`
	// The timestamp when the invoice was updated.
	UpdatedAt string `json:"updatedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Whether or not the invoice payment instructions will show the real account and
	// routing number for the destination account or use virtual account numbers
	// instead.
	UseRealAccountNumber bool `json:"useRealAccountNumber" api:"required"`
	// The time when the invoice was canceled.
	CanceledAt string `json:"canceledAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Internal note for the invoice, visible by users in the mercury organization but
	// not visible to payers.
	InternalNote string `json:"internalNote" api:"nullable"`
	// Memo for the payer of the invoice.
	PayerMemo string `json:"payerMemo" api:"nullable"`
	// Purchase order number for the invoice if applicable.
	PoNumber string `json:"poNumber" api:"nullable"`
	// The end date for the service period this invoice covers, if applicable.
	// YYYY-MM-DD
	ServicePeriodEndDate time.Time `json:"servicePeriodEndDate" api:"nullable" format:"date"`
	// The start date for the service period this invoice covers, if applicable.
	// YYYY-MM-DD
	ServicePeriodStartDate time.Time `json:"servicePeriodStartDate" api:"nullable" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		ACHDebitEnabled        respjson.Field
		Amount                 respjson.Field
		CcEmails               respjson.Field
		CreatedAt              respjson.Field
		CreditCardEnabled      respjson.Field
		CustomerID             respjson.Field
		DestinationAccountID   respjson.Field
		DueDate                respjson.Field
		InvoiceDate            respjson.Field
		InvoiceNumber          respjson.Field
		LineItems              respjson.Field
		Slug                   respjson.Field
		Status                 respjson.Field
		UpdatedAt              respjson.Field
		UseRealAccountNumber   respjson.Field
		CanceledAt             respjson.Field
		InternalNote           respjson.Field
		PayerMemo              respjson.Field
		PoNumber               respjson.Field
		ServicePeriodEndDate   respjson.Field
		ServicePeriodStartDate respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Invoice) RawJSON() string { return r.JSON.raw }
func (r *Invoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data for an invoice line item
type LineItemData struct {
	// the name of the line item
	Name string `json:"name" api:"required"`
	// the quantity of this item
	Quantity float64 `json:"quantity" api:"required"`
	// A dollar amount
	UnitPrice float64 `json:"unitPrice" api:"required"`
	// the sales tax applied to this item
	SalesTaxRate float64 `json:"salesTaxRate" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		Quantity     respjson.Field
		UnitPrice    respjson.Field
		SalesTaxRate respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LineItemData) RawJSON() string { return r.JSON.raw }
func (r *LineItemData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LineItemData to a LineItemDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LineItemDataParam.Overrides()
func (r LineItemData) ToParam() LineItemDataParam {
	return param.Override[LineItemDataParam](json.RawMessage(r.RawJSON()))
}

// Data for an invoice line item
//
// The properties Name, Quantity, UnitPrice are required.
type LineItemDataParam struct {
	// the name of the line item
	Name string `json:"name" api:"required"`
	// the quantity of this item
	Quantity float64 `json:"quantity" api:"required"`
	// A dollar amount
	UnitPrice float64 `json:"unitPrice" api:"required"`
	// the sales tax applied to this item
	SalesTaxRate param.Opt[float64] `json:"salesTaxRate,omitzero"`
	paramObj
}

func (r LineItemDataParam) MarshalJSON() (data []byte, err error) {
	type shadow LineItemDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LineItemDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentLinkStatus string

const (
	PaymentLinkStatusUnpaid     PaymentLinkStatus = "Unpaid"
	PaymentLinkStatusPaid       PaymentLinkStatus = "Paid"
	PaymentLinkStatusCancelled  PaymentLinkStatus = "Cancelled"
	PaymentLinkStatusProcessing PaymentLinkStatus = "Processing"
)

// Response data for Accounts Receivable invoices API Endpoint
type AccountsReceivableInvoiceListResponse struct {
	// ID for the invoice.
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether or not the invoice can be paid via ach debit.
	ACHDebitEnabled bool `json:"achDebitEnabled" api:"required"`
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount" api:"required"`
	// Emails to be CCed on invoice notifications/reminders.
	CcEmails []string `json:"ccEmails" api:"required"`
	// The timestamp when the invoice was created.
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Whether or not the invoice can be paid via credit card. Requires stripe to be
	// setup for the Mercury account.
	CreditCardEnabled bool `json:"creditCardEnabled" api:"required"`
	// The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint
	// to list your customers and find the corresponding id, or create a new customer
	// first.
	CustomerID string `json:"customerId" api:"required" format:"uuid"`
	// ID for a Mercury account.
	DestinationAccountID string `json:"destinationAccountId" api:"required" format:"uuid"`
	// The due date the invoice should be paid by.
	DueDate time.Time `json:"dueDate" api:"required" format:"date"`
	// The date of the invoice, set by the invoice creator and likely to be context
	// specific to the type of transaction. i.e. it could be a date a service was
	// performed, it does not need to be the date the invoice was created.
	InvoiceDate time.Time `json:"invoiceDate" api:"required" format:"date"`
	// The payer facing invoice number/identifier.
	InvoiceNumber string `json:"invoiceNumber" api:"required"`
	// A unique identifier used to build public URLs for this invoice. Use it to
	// construct the payment page URL (https://app.mercury.com/pay/{slug}) or fetch the
	// invoice PDF via /api/v1/ar/invoices/{slug}/pdf.
	Slug string `json:"slug" api:"required"`
	// The status of the invoice.
	//
	// Any of "Unpaid", "Paid", "Cancelled", "Processing".
	Status PaymentLinkStatus `json:"status" api:"required"`
	// The timestamp when the invoice was updated.
	UpdatedAt string `json:"updatedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Whether or not the invoice payment instructions will show the real account and
	// routing number for the destination account or use virtual account numbers
	// instead.
	UseRealAccountNumber bool `json:"useRealAccountNumber" api:"required"`
	// The time when the invoice was canceled.
	CanceledAt string `json:"canceledAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Internal note for the invoice, visible by users in the mercury organization but
	// not visible to payers.
	InternalNote string `json:"internalNote" api:"nullable"`
	// Memo for the payer of the invoice.
	PayerMemo string `json:"payerMemo" api:"nullable"`
	// Purchase order number for the invoice if applicable.
	PoNumber string `json:"poNumber" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		ACHDebitEnabled      respjson.Field
		Amount               respjson.Field
		CcEmails             respjson.Field
		CreatedAt            respjson.Field
		CreditCardEnabled    respjson.Field
		CustomerID           respjson.Field
		DestinationAccountID respjson.Field
		DueDate              respjson.Field
		InvoiceDate          respjson.Field
		InvoiceNumber        respjson.Field
		Slug                 respjson.Field
		Status               respjson.Field
		UpdatedAt            respjson.Field
		UseRealAccountNumber respjson.Field
		CanceledAt           respjson.Field
		InternalNote         respjson.Field
		PayerMemo            respjson.Field
		PoNumber             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountsReceivableInvoiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountsReceivableInvoiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The response type for fetching attachments related to an AR Invoice.
type AccountsReceivableInvoiceListAttachmentsResponse struct {
	// The list of attachments
	Attachments []Attachment `json:"attachments" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachments respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountsReceivableInvoiceListAttachmentsResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountsReceivableInvoiceListAttachmentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountsReceivableInvoiceNewParams struct {
	// Whether or not the invoice can be paid via ACH debit.
	ACHDebitEnabled bool `json:"achDebitEnabled" api:"required"`
	// Emails to be CCed on invoice notifications/reminders.
	CcEmails []string `json:"ccEmails,omitzero" api:"required"`
	// Whether or not the invoice can be paid via credit card. Requires Stripe to be
	// setup for the Mercury account.
	CreditCardEnabled bool `json:"creditCardEnabled" api:"required"`
	// The customer who will receive the invoice. Use the /api/v1/ar/customers endpoint
	// to list your customers and find the corresponding id, or create a new customer
	// first.
	CustomerID string `json:"customerId" api:"required" format:"uuid"`
	// ID for a Mercury account.
	DestinationAccountID string `json:"destinationAccountId" api:"required" format:"uuid"`
	// The due date the invoice should be paid by. YYYY-MM-DD
	DueDate time.Time `json:"dueDate" api:"required" format:"date"`
	// The date of the invoice, set by the invoice creator and likely to be context
	// specific to the type of transaction. For example, it could be a date a service
	// was performed. YYYY-MM-DD
	InvoiceDate time.Time `json:"invoiceDate" api:"required" format:"date"`
	// The line items for the invoice
	LineItems []LineItemDataParam `json:"lineItems,omitzero" api:"required"`
	// Whether or not the invoice payment instructions will show the real account and
	// routing number for the destination account or use virtual account numbers
	// instead. Virtual accounts are safer and are preferred in most cases.
	UseRealAccountNumber bool `json:"useRealAccountNumber" api:"required"`
	// Internal note for the invoice, visible by users in the organization but not
	// visible to payers.
	InternalNote param.Opt[string] `json:"internalNote,omitzero"`
	// The payer facing invoice number/identifier.
	InvoiceNumber param.Opt[string] `json:"invoiceNumber,omitzero"`
	// Memo for the payer of the invoice.
	PayerMemo param.Opt[string] `json:"payerMemo,omitzero"`
	// Purchase order number for the invoice, if applicable.
	PoNumber param.Opt[string] `json:"poNumber,omitzero"`
	// The end date for the service period this invoice covers, if applicable.
	// YYYY-MM-DD
	ServicePeriodEndDate param.Opt[time.Time] `json:"servicePeriodEndDate,omitzero" format:"date"`
	// The start date for the service period this invoice covers, if applicable.
	// YYYY-MM-DD
	ServicePeriodStartDate param.Opt[time.Time] `json:"servicePeriodStartDate,omitzero" format:"date"`
	// Rules for emailing the new invoice to payers. Can be "DontSend" to skip sending
	// or "SendNow" to send immediately. If omitted, defaults to sending immediately.
	//
	// Any of "DontSend", "SendNow".
	SendEmailOption AccountsReceivableInvoiceNewParamsSendEmailOption `json:"sendEmailOption,omitzero"`
	paramObj
}

func (r AccountsReceivableInvoiceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountsReceivableInvoiceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountsReceivableInvoiceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rules for emailing the new invoice to payers. Can be "DontSend" to skip sending
// or "SendNow" to send immediately. If omitted, defaults to sending immediately.
type AccountsReceivableInvoiceNewParamsSendEmailOption string

const (
	AccountsReceivableInvoiceNewParamsSendEmailOptionDontSend AccountsReceivableInvoiceNewParamsSendEmailOption = "DontSend"
	AccountsReceivableInvoiceNewParamsSendEmailOptionSendNow  AccountsReceivableInvoiceNewParamsSendEmailOption = "SendNow"
)

type AccountsReceivableInvoiceUpdateParams struct {
	// Whether or not the invoice can be paid via ACH debit.
	ACHDebitEnabled bool `json:"achDebitEnabled" api:"required"`
	// List of emails to be CCed on notifications/reminders.
	CcEmails []string `json:"ccEmails,omitzero" api:"required"`
	// Whether or not the invoice can be paid via credit card. Requires Stripe to be
	// setup for the Mercury account.
	CreditCardEnabled bool `json:"creditCardEnabled" api:"required"`
	// The date the invoice should be paid by. YYYY-MM-DD
	DueDate time.Time `json:"dueDate" api:"required" format:"date"`
	// The date of the invoice, set by the invoice creator. Does not have to be the day
	// the invoice was created. It can be business specific i.e. service/sale date.
	// YYYY-MM-DD
	InvoiceDate time.Time `json:"invoiceDate" api:"required" format:"date"`
	// The invoice number.
	InvoiceNumber string `json:"invoiceNumber" api:"required"`
	// The line items for the invoice
	LineItems []LineItemDataParam `json:"lineItems,omitzero" api:"required"`
	// Whether or not the invoice payment instructions will show the real account and
	// routing number for the destination account or use virtual account numbers
	// instead.
	UseRealAccountNumber bool `json:"useRealAccountNumber" api:"required"`
	// Internal note for the invoice, visible by users in the organization but not
	// visible to payers.
	InternalNote param.Opt[string] `json:"internalNote,omitzero"`
	// Memo for the payer of the invoice.
	PayerMemo param.Opt[string] `json:"payerMemo,omitzero"`
	// The purchase order number for the invoice if applicable.
	PoNumber param.Opt[string] `json:"poNumber,omitzero"`
	// The end date for the service period this invoice covers, if applicable.
	// YYYY-MM-DD
	ServicePeriodEndDate param.Opt[time.Time] `json:"servicePeriodEndDate,omitzero" format:"date"`
	// The start date for the service period this invoice covers, if applicable.
	// YYYY-MM-DD
	ServicePeriodStartDate param.Opt[time.Time] `json:"servicePeriodStartDate,omitzero" format:"date"`
	paramObj
}

func (r AccountsReceivableInvoiceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountsReceivableInvoiceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountsReceivableInvoiceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountsReceivableInvoiceListParams struct {
	// The ID of the invoice to end the page before (exclusive). When provided, results
	// will end just before this ID and work backwards. Use this for reverse pagination
	// or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the invoice to start the page after (exclusive). When provided,
	// results will begin with the invoice immediately following this ID. Use this for
	// standard forward pagination to get the next page of results. Cannot be combined
	// with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order AccountsReceivableInvoiceListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountsReceivableInvoiceListParams]'s query parameters as
// `url.Values`.
func (r AccountsReceivableInvoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type AccountsReceivableInvoiceListParamsOrder string

const (
	AccountsReceivableInvoiceListParamsOrderAsc  AccountsReceivableInvoiceListParamsOrder = "asc"
	AccountsReceivableInvoiceListParamsOrderDesc AccountsReceivableInvoiceListParamsOrder = "desc"
)
