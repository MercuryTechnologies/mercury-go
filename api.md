# Ar

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#Attachment">Attachment</a>

Methods:

- <code title="get /ar/attachments/{attachmentId}">client.Ar.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArService.GetAttachment">GetAttachment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, attachmentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#Attachment">Attachment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Customers

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AddressInputParam">AddressInputParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CustomerResponse">CustomerResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArCustomerListResponse">ArCustomerListResponse</a>

Methods:

- <code title="post /ar/customers">client.Ar.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArCustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArCustomerNewParams">ArCustomerNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CustomerResponse">CustomerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ar/customers/{customerId}">client.Ar.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArCustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CustomerResponse">CustomerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ar/customers/{customerId}">client.Ar.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArCustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArCustomerUpdateParams">ArCustomerUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CustomerResponse">CustomerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ar/customers">client.Ar.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArCustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArCustomerListParams">ArCustomerListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArCustomerListResponse">ArCustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /ar/customers/{customerId}">client.Ar.Customers.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArCustomerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Invoices

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#LineItemDataParam">LineItemDataParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#InvoiceResponse">InvoiceResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#LineItemData">LineItemData</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#PaymentLinkStatus">PaymentLinkStatus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceListResponse">ArInvoiceListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceListAttachmentsResponse">ArInvoiceListAttachmentsResponse</a>

Methods:

- <code title="post /ar/invoices">client.Ar.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceNewParams">ArInvoiceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ar/invoices/{invoiceId}">client.Ar.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ar/invoices/{invoiceId}">client.Ar.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceUpdateParams">ArInvoiceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#InvoiceResponse">InvoiceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ar/invoices">client.Ar.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceListParams">ArInvoiceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceListResponse">ArInvoiceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ar/invoices/{invoiceId}/cancel">client.Ar.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /ar/invoices/{invoiceId}/pdf">client.Ar.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceService.DownloadPdf">DownloadPdf</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ar/invoices/{invoiceId}/attachments">client.Ar.Invoices.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceService.ListAttachments">ListAttachments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, invoiceID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ArInvoiceListAttachmentsResponse">ArInvoiceListAttachmentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Books

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#JournalEntry">JournalEntry</a>

Methods:

- <code title="post /books/agent-ledger-templates">client.Books.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookService.NewAgentLedgerTemplate">NewAgentLedgerTemplate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookNewAgentLedgerTemplateParams">BookNewAgentLedgerTemplateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#LedgerTemplate">LedgerTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /books/journal-entry/{booksId}/{tealJournalEntryId}">client.Books.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookService.GetJournalEntry">GetJournalEntry</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tealJournalEntryID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookGetJournalEntryParams">BookGetJournalEntryParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#JournalEntry">JournalEntry</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AgentLedgerTemplate

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#LedgerTemplate">LedgerTemplate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentLedgerTemplateDeleteResponse">BookAgentLedgerTemplateDeleteResponse</a>

Methods:

- <code title="put /books/agent-ledger-template/{ledgerId}">client.Books.AgentLedgerTemplate.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentLedgerTemplateService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ledgerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentLedgerTemplateUpdateParams">BookAgentLedgerTemplateUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#LedgerTemplate">LedgerTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /books/agent-ledger-template/{ledgerId}">client.Books.AgentLedgerTemplate.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentLedgerTemplateService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ledgerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentLedgerTemplateDeleteResponse">BookAgentLedgerTemplateDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## JournalEntries

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryListResponse">BookJournalEntryListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryBulkDeleteResponse">BookJournalEntryBulkDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryBulkUpdateResponse">BookJournalEntryBulkUpdateResponse</a>

Methods:

- <code title="get /books/journal-entries/{booksId}">client.Books.JournalEntries.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, booksID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryListParams">BookJournalEntryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryListResponse">BookJournalEntryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /books/journal-entries/{booksId}">client.Books.JournalEntries.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryService.BulkDelete">BulkDelete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, booksID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryBulkDeleteParams">BookJournalEntryBulkDeleteParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryBulkDeleteResponse">BookJournalEntryBulkDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /books/journal-entries/{booksId}">client.Books.JournalEntries.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryService.BulkUpdate">BulkUpdate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, booksID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryBulkUpdateParams">BookJournalEntryBulkUpdateParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryBulkUpdateResponse">BookJournalEntryBulkUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /books/journal-entries/{booksId}">client.Books.JournalEntries.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryService.NewMultiple">NewMultiple</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, booksID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookJournalEntryNewMultipleParams">BookJournalEntryNewMultipleParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#JournalEntry">JournalEntry</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AgentCoaTemplates

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ChartOfAccountsTemplateForBooks">ChartOfAccountsTemplateForBooks</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ChartOfAccountsTemplateForTeal">ChartOfAccountsTemplateForTeal</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentCoaTemplateListResponse">BookAgentCoaTemplateListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentCoaTemplateDeleteResponse">BookAgentCoaTemplateDeleteResponse</a>

Methods:

- <code title="post /books/agent-coa-templates">client.Books.AgentCoaTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentCoaTemplateService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentCoaTemplateNewParams">BookAgentCoaTemplateNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ChartOfAccountsTemplateForTeal">ChartOfAccountsTemplateForTeal</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /books/agent-coa-template/{coaTemplateId}">client.Books.AgentCoaTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentCoaTemplateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, coaTemplateID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentCoaTemplateGetParams">BookAgentCoaTemplateGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ChartOfAccountsTemplateForTeal">ChartOfAccountsTemplateForTeal</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /books/agent-coa-templates">client.Books.AgentCoaTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentCoaTemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentCoaTemplateListParams">BookAgentCoaTemplateListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentCoaTemplateListResponse">BookAgentCoaTemplateListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /books/agent-coa-template/{coaTemplateId}">client.Books.AgentCoaTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentCoaTemplateService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, coaTemplateID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#BookAgentCoaTemplateDeleteResponse">BookAgentCoaTemplateDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Categories

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CategoryListResponse">CategoryListResponse</a>

Methods:

- <code title="get /categories">client.Categories.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CategoryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CategoryListParams">CategoryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CategoryListResponse">CategoryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Credit

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CreditListResponse">CreditListResponse</a>

Methods:

- <code title="get /credit">client.Credit.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CreditService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CreditListResponse">CreditListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Events

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#APIEvent">APIEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#EventListResponse">EventListResponse</a>

Methods:

- <code title="get /events/{eventId}">client.Events.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#EventService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, eventID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#APIEvent">APIEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /events">client.Events.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#EventService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#EventListParams">EventListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#EventListResponse">EventListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Organization

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#OrganizationGetResponse">OrganizationGetResponse</a>

Methods:

- <code title="get /organization">client.Organization.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#OrganizationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#OrganizationGetResponse">OrganizationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RequestSendMoney

Methods:

- <code title="get /request-send-money/{requestId}">client.RequestSendMoney.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RequestSendMoneyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#SendMoneyApproval">SendMoneyApproval</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Safes

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#UsState">UsState</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#APISafeRequest">APISafeRequest</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#UsState">UsState</a>

Methods:

- <code title="get /safes/{safeRequestId}">client.Safes.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#SafeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, safeRequestID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#APISafeRequest">APISafeRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /safes">client.Safes.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#SafeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#APISafeRequest">APISafeRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /safes/{safeRequestId}/document">client.Safes.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#SafeService.DownloadDocument">DownloadDocument</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, safeRequestID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Statements

Methods:

- <code title="get /statements/{statementId}/pdf">client.Statements.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#StatementService.DownloadPdf">DownloadPdf</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, statementID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transfer

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransferNewResponse">TransferNewResponse</a>

Methods:

- <code title="post /transfer">client.Transfer.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransferService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransferNewParams">TransferNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransferNewResponse">TransferNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Treasury

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryListResponse">TreasuryListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryGetStatementsResponse">TreasuryGetStatementsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryGetTransactionsResponse">TreasuryGetTransactionsResponse</a>

Methods:

- <code title="get /treasury">client.Treasury.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryListParams">TreasuryListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryListResponse">TreasuryListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /treasury/{treasuryId}/statements">client.Treasury.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryService.GetStatements">GetStatements</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, treasuryID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryGetStatementsParams">TreasuryGetStatementsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryGetStatementsResponse">TreasuryGetStatementsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /treasury/{treasuryId}/transactions">client.Treasury.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryService.GetTransactions">GetTransactions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, treasuryID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryGetTransactionsParams">TreasuryGetTransactionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TreasuryGetTransactionsResponse">TreasuryGetTransactionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#UserDetails">UserDetails</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#UserListResponse">UserListResponse</a>

Methods:

- <code title="get /users/{userId}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#UserDetails">UserDetails</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#UserListParams">UserListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#UserListResponse">UserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#APIWebhook">APIWebhook</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookListResponse">WebhookListResponse</a>

Methods:

- <code title="post /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookNewParams">WebhookNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#APIWebhook">APIWebhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks/{webhookEndpointId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookEndpointID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#APIWebhook">APIWebhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /webhooks/{webhookEndpointId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookEndpointID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookUpdateParams">WebhookUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#APIWebhook">APIWebhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookListParams">WebhookListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookListResponse">WebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /webhooks/{webhookEndpointId}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookEndpointID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /webhooks/{webhookEndpointId}/verify">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookEndpointID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#WebhookVerifyParams">WebhookVerifyParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Account

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#SendMoneyPaymentMethod">SendMoneyPaymentMethod</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountStatus">AccountStatus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CategoryData">CategoryData</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#MercuryCategory">MercuryCategory</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#SendMoneyApproval">SendMoneyApproval</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#SendMoneyPaymentMethod">SendMoneyPaymentMethod</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#Transaction">Transaction</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountListResponse">AccountListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountListCardsResponse">AccountListCardsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountListStatementsResponse">AccountListStatementsResponse</a>

Methods:

- <code title="get /account/{accountId}">client.Account.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts">client.Account.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountListParams">AccountListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountListResponse">AccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account/{accountId}/cards">client.Account.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountService.ListCards">ListCards</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountListCardsResponse">AccountListCardsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account/{accountId}/statements">client.Account.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountService.ListStatements">ListStatements</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountListStatementsParams">AccountListStatementsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountListStatementsResponse">AccountListStatementsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /account/{accountId}/request-send-money">client.Account.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountService.RequestSendMoney">RequestSendMoney</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountRequestSendMoneyParams">AccountRequestSendMoneyParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#SendMoneyApproval">SendMoneyApproval</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /account/{accountId}/transaction/{transactionId}">client.Account.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountService.GetTransaction">GetTransaction</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountGetTransactionParams">AccountGetTransactionParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountTransactionListResponse">AccountTransactionListResponse</a>

Methods:

- <code title="get /account/{accountId}/transactions">client.Account.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountTransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountTransactionListParams">AccountTransactionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountTransactionListResponse">AccountTransactionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /account/{accountId}/transactions">client.Account.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountTransactionService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AccountTransactionSendParams">AccountTransactionSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Recipients

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AddressDataParam">AddressDataParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AddressWithoutNameParam">AddressWithoutNameParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#CheckInfoRawParam">CheckInfoRawParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#DomesticWireRoutingInfoRawParam">DomesticWireRoutingInfoRawParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ElectronicAccountType">ElectronicAccountType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ElectronicRoutingInfoRawParam">ElectronicRoutingInfoRawParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AddressData">AddressData</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#AddressWithoutName">AddressWithoutName</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#DomesticWireRoutingInfo">DomesticWireRoutingInfo</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ElectronicAccountType">ElectronicAccountType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#ElectronicRoutingInfo">ElectronicRoutingInfo</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#InternationalWireRoutingInfo">InternationalWireRoutingInfo</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientInfo">RecipientInfo</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#SwiftBankAccountType">SwiftBankAccountType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TaxFormType">TaxFormType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientListResponse">RecipientListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientListAttachmentsResponse">RecipientListAttachmentsResponse</a>

Methods:

- <code title="post /recipients">client.Recipients.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientNewParams">RecipientNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientInfo">RecipientInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /recipient/{recipientId}">client.Recipients.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, recipientID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientInfo">RecipientInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /recipient/{recipientId}">client.Recipients.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, recipientID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientUpdateParams">RecipientUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientInfo">RecipientInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /recipients">client.Recipients.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientListParams">RecipientListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientListResponse">RecipientListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /recipients/attachments">client.Recipients.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientService.ListAttachments">ListAttachments</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientListAttachmentsParams">RecipientListAttachmentsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientListAttachmentsResponse">RecipientListAttachmentsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /recipient/{recipientId}/attachments">client.Recipients.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientService.UploadAttachment">UploadAttachment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, recipientID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#RecipientUploadAttachmentParams">RecipientUploadAttachmentParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransactionListResponse">TransactionListResponse</a>

Methods:

- <code title="get /transaction/{transactionId}">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /transaction/{transactionId}">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransactionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransactionUpdateParams">TransactionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransactionListParams">TransactionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransactionListResponse">TransactionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /transaction/{transactionId}/attachments">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransactionService.UploadAttachment">UploadAttachment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go">mercury</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/mercury-go#TransactionUploadAttachmentParams">TransactionUploadAttachmentParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
