# Customers

Params Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#AddressInputParam">AddressInputParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CustomerAddress">CustomerAddress</a>

Methods:

- <code title="post /ar/customers">client.Customers.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CustomerNewParams">CustomerNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ar/customers/{customerId}">client.Customers.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CustomerUpdateParams">CustomerUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ar/customers">client.Customers.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CustomerListParams">CustomerListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDArCustomers">CursorIDArCustomers</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Customer">Customer</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /ar/customers/{customerId}">client.Customers.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CustomerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /ar/customers/{customerId}">client.Customers.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Customer">Customer</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Invoices

Params Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#LineItemDataParam">LineItemDataParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Invoice">Invoice</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#LineItemData">LineItemData</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentLinkStatus">PaymentLinkStatus</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceListResponse">InvoiceListResponse</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceListAttachmentsResponse">InvoiceListAttachmentsResponse</a>

Methods:

- <code title="post /ar/invoices">client.Invoices.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceNewParams">InvoiceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Invoice">Invoice</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ar/invoices/{invoiceId}">client.Invoices.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceUpdateParams">InvoiceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Invoice">Invoice</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ar/invoices">client.Invoices.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceListParams">InvoiceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDArInvoices">CursorIDArInvoices</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceListResponse">InvoiceListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ar/invoices/{invoiceId}/cancel">client.Invoices.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /ar/invoices/{invoiceId}/pdf">client.Invoices.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ar/invoices/{invoiceId}">client.Invoices.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Invoice">Invoice</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ar/invoices/{invoiceId}/attachments">client.Invoices.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceService.ListAttachments">ListAttachments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceListAttachmentsResponse">InvoiceListAttachmentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Attachments

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Attachment">Attachment</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceAttachmentListResponse">InvoiceAttachmentListResponse</a>

Methods:

- <code title="get /ar/invoices/{invoiceId}/attachments">client.Invoices.Attachments.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceAttachmentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceAttachmentListResponse">InvoiceAttachmentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ar/attachments/{attachmentId}">client.Invoices.Attachments.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InvoiceAttachmentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, attachmentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Attachment">Attachment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cards

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CardListResponse">CardListResponse</a>

Methods:

- <code title="get /account/{accountId}/cards">client.Cards.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CardListResponse">CardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Categories

Methods:

- <code title="get /categories">client.Categories.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CategoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CategoryListParams">CategoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDCategories">CursorIDCategories</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CategoryData">CategoryData</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Credit

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CreditListResponse">CreditListResponse</a>

Methods:

- <code title="get /credit">client.Credit.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CreditService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CreditListResponse">CreditListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Events

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Event">Event</a>

Methods:

- <code title="get /events">client.Events.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#EventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#EventListParams">EventListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDEvents">CursorIDEvents</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Event">Event</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /events/{eventId}">client.Events.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#EventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Event">Event</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Org

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#OrgGetResponse">OrgGetResponse</a>

Methods:

- <code title="get /organization">client.Org.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#OrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#OrgGetResponse">OrgGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

Params Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SendMoneyPaymentMethod">SendMoneyPaymentMethod</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SendMoneyApproval">SendMoneyApproval</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SendMoneyPaymentMethod">SendMoneyPaymentMethod</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentTransferResponse">PaymentTransferResponse</a>

Methods:

- <code title="post /account/{accountId}/transactions">client.Payments.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentNewParams">PaymentNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /request-send-money">client.Payments.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentListParams">PaymentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDRequestSendMoney">CursorIDRequestSendMoney</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SendMoneyApproval">SendMoneyApproval</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /request-send-money/{requestId}">client.Payments.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SendMoneyApproval">SendMoneyApproval</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /account/{accountId}/request-send-money">client.Payments.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentService.Request">Request</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentRequestParams">PaymentRequestParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SendMoneyApproval">SendMoneyApproval</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /transfer">client.Payments.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentService.Transfer">Transfer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentTransferParams">PaymentTransferParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#PaymentTransferResponse">PaymentTransferResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Safes

Params Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#UsState">UsState</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SafeRequest">SafeRequest</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#UsState">UsState</a>

Methods:

- <code title="get /safes">client.Safes.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SafeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SafeRequest">SafeRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /safes/{safeRequestId}/document">client.Safes.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SafeService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, safeRequestID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /safes/{safeRequestId}">client.Safes.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SafeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, safeRequestID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SafeRequest">SafeRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Statements

Methods:

- <code title="get /statements/{statementId}/pdf">client.Statements.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#StatementService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, statementID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#StatementAccountListResponse">StatementAccountListResponse</a>

Methods:

- <code title="get /account/{accountId}/statements">client.Statements.Accounts.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#StatementAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#StatementAccountListParams">StatementAccountListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDAccountStatements">CursorIDAccountStatements</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#StatementAccountListResponse">StatementAccountListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Treasury

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#StatementTreasuryListResponse">StatementTreasuryListResponse</a>

Methods:

- <code title="get /treasury/{treasuryId}/statements">client.Statements.Treasury.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#StatementTreasuryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, treasuryID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#StatementTreasuryListParams">StatementTreasuryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDAccountStatements">CursorIDAccountStatements</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#StatementTreasuryListResponse">StatementTreasuryListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Treasury

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TreasuryListResponse">TreasuryListResponse</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TreasuryGetTransactionsResponse">TreasuryGetTransactionsResponse</a>

Methods:

- <code title="get /treasury">client.Treasury.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TreasuryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TreasuryListParams">TreasuryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDAccounts">CursorIDAccounts</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TreasuryListResponse">TreasuryListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /treasury/{treasuryId}/transactions">client.Treasury.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TreasuryService.GetTransactions">GetTransactions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, treasuryID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TreasuryGetTransactionsParams">TreasuryGetTransactionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TreasuryGetTransactionsResponse">TreasuryGetTransactionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#User">User</a>

Methods:

- <code title="get /users">client.Users.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#UserListParams">UserListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDUsers">CursorIDUsers</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#User">User</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users/{userId}">client.Users.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Webhook">Webhook</a>

Methods:

- <code title="post /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#WebhookNewParams">WebhookNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /webhooks/{webhookEndpointId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookEndpointID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#WebhookUpdateParams">WebhookUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#WebhookListParams">WebhookListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDWebhooks">CursorIDWebhooks</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Webhook">Webhook</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /webhooks/{webhookEndpointId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookEndpointID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /webhooks/{webhookEndpointId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookEndpointID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Webhook">Webhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /webhooks/{webhookEndpointId}/verify">client.Webhooks.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#WebhookService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookEndpointID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#WebhookVerifyParams">WebhookVerifyParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#AccountStatus">AccountStatus</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CategoryData">CategoryData</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#MercuryCategory">MercuryCategory</a>

Methods:

- <code title="get /accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#AccountListParams">AccountListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDAccounts">CursorIDAccounts</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Account">Account</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account/{accountId}">client.Accounts.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Recipients

Params Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#AddressDataParam">AddressDataParam</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#AddressWithoutNameParam">AddressWithoutNameParam</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CheckInfoRawParam">CheckInfoRawParam</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#DomesticWireRoutingInfoRawParam">DomesticWireRoutingInfoRawParam</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#ElectronicAccountType">ElectronicAccountType</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#ElectronicRoutingInfoRawParam">ElectronicRoutingInfoRawParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#AddressData">AddressData</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#AddressWithoutName">AddressWithoutName</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CheckInfo">CheckInfo</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#DomesticWireRoutingInfo">DomesticWireRoutingInfo</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#ElectronicAccountType">ElectronicAccountType</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#ElectronicRoutingInfo">ElectronicRoutingInfo</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#InternationalWireRoutingInfo">InternationalWireRoutingInfo</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RealTimePaymentRoutingInfo">RealTimePaymentRoutingInfo</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Recipient">Recipient</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientAttachment">RecipientAttachment</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#SwiftBankAccountType">SwiftBankAccountType</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TaxFormType">TaxFormType</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientListAttachmentsResponse">RecipientListAttachmentsResponse</a>

Methods:

- <code title="post /recipients">client.Recipients.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientNewParams">RecipientNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Recipient">Recipient</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /recipient/{recipientId}">client.Recipients.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, recipientID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientUpdateParams">RecipientUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Recipient">Recipient</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /recipients">client.Recipients.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientListParams">RecipientListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDRecipients">CursorIDRecipients</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Recipient">Recipient</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /recipient/{recipientId}/attachments">client.Recipients.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientService.Attach">Attach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, recipientID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientAttachParams">RecipientAttachParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /recipient/{recipientId}">client.Recipients.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, recipientID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Recipient">Recipient</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /recipients/attachments">client.Recipients.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientService.ListAttachments">ListAttachments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientListAttachmentsParams">RecipientListAttachmentsParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDRecipientAttachments">CursorIDRecipientAttachments</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RecipientListAttachmentsResponse">RecipientListAttachmentsResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#CurrencyExchangeInfo">CurrencyExchangeInfo</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#GlAllocation">GlAllocation</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#MerchantData">MerchantData</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#RelatedTransactionData">RelatedTransactionData</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Transaction">Transaction</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TransactionAttachment">TransactionAttachment</a>
- <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TransactionMethodData">TransactionMethodData</a>

Methods:

- <code title="patch /transaction/{transactionId}">client.Transactions.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TransactionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TransactionUpdateParams">TransactionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TransactionListParams">TransactionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go/packages/pagination#CursorIDTransactions">CursorIDTransactions</a>[<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Transaction">Transaction</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /transaction/{transactionId}/attachments">client.Transactions.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TransactionService.Attach">Attach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TransactionAttachParams">TransactionAttachParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /transaction/{transactionId}">client.Transactions.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/MercuryTechnologies/mercury-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
