// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"reflect"

	"github.com/stainless-sdks/mercury-go/internal/apijson"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/param"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type CursorIDAccounts[T any] struct {
	Accounts []T `json:"accounts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorIDAccounts[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDAccounts[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDAccounts[T]) GetNextPage() (res *CursorIDAccounts[T], err error) {
	if len(r.Accounts) == 0 {
		return nil, nil
	}
	items := r.Accounts
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("start_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorIDAccounts[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDAccounts[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDAccountsAutoPager[T any] struct {
	page *CursorIDAccounts[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDAccountsAutoPager[T any](page *CursorIDAccounts[T], err error) *CursorIDAccountsAutoPager[T] {
	return &CursorIDAccountsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDAccountsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Accounts) == 0 {
		return false
	}
	if r.idx >= len(r.page.Accounts) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Accounts) == 0 {
			return false
		}
	}
	r.cur = r.page.Accounts[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDAccountsAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDAccountsAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDAccountsAutoPager[T]) Index() int {
	return r.run
}

type CursorIDTransactions[T any] struct {
	Transactions []T `json:"transactions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transactions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorIDTransactions[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDTransactions[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDTransactions[T]) GetNextPage() (res *CursorIDTransactions[T], err error) {
	if len(r.Transactions) == 0 {
		return nil, nil
	}
	items := r.Transactions
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("start_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorIDTransactions[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDTransactions[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDTransactionsAutoPager[T any] struct {
	page *CursorIDTransactions[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDTransactionsAutoPager[T any](page *CursorIDTransactions[T], err error) *CursorIDTransactionsAutoPager[T] {
	return &CursorIDTransactionsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDTransactionsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Transactions) == 0 {
		return false
	}
	if r.idx >= len(r.page.Transactions) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Transactions) == 0 {
			return false
		}
	}
	r.cur = r.page.Transactions[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDTransactionsAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDTransactionsAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDTransactionsAutoPager[T]) Index() int {
	return r.run
}

type CursorIDEvents[T any] struct {
	Events []T `json:"events"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorIDEvents[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDEvents[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDEvents[T]) GetNextPage() (res *CursorIDEvents[T], err error) {
	if len(r.Events) == 0 {
		return nil, nil
	}
	items := r.Events
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("start_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorIDEvents[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDEvents[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDEventsAutoPager[T any] struct {
	page *CursorIDEvents[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDEventsAutoPager[T any](page *CursorIDEvents[T], err error) *CursorIDEventsAutoPager[T] {
	return &CursorIDEventsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDEventsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Events) == 0 {
		return false
	}
	if r.idx >= len(r.page.Events) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Events) == 0 {
			return false
		}
	}
	r.cur = r.page.Events[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDEventsAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDEventsAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDEventsAutoPager[T]) Index() int {
	return r.run
}

type CursorIDCustomers[T any] struct {
	Customers []T `json:"customers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Customers   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorIDCustomers[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDCustomers[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDCustomers[T]) GetNextPage() (res *CursorIDCustomers[T], err error) {
	if len(r.Customers) == 0 {
		return nil, nil
	}
	items := r.Customers
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("start_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorIDCustomers[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDCustomers[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDCustomersAutoPager[T any] struct {
	page *CursorIDCustomers[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDCustomersAutoPager[T any](page *CursorIDCustomers[T], err error) *CursorIDCustomersAutoPager[T] {
	return &CursorIDCustomersAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDCustomersAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Customers) == 0 {
		return false
	}
	if r.idx >= len(r.page.Customers) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Customers) == 0 {
			return false
		}
	}
	r.cur = r.page.Customers[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDCustomersAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDCustomersAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDCustomersAutoPager[T]) Index() int {
	return r.run
}

type CursorIDInvoices[T any] struct {
	Invoices []T `json:"invoices"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Invoices    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorIDInvoices[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDInvoices[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDInvoices[T]) GetNextPage() (res *CursorIDInvoices[T], err error) {
	if len(r.Invoices) == 0 {
		return nil, nil
	}
	items := r.Invoices
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("start_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorIDInvoices[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDInvoices[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDInvoicesAutoPager[T any] struct {
	page *CursorIDInvoices[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDInvoicesAutoPager[T any](page *CursorIDInvoices[T], err error) *CursorIDInvoicesAutoPager[T] {
	return &CursorIDInvoicesAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDInvoicesAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Invoices) == 0 {
		return false
	}
	if r.idx >= len(r.page.Invoices) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Invoices) == 0 {
			return false
		}
	}
	r.cur = r.page.Invoices[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDInvoicesAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDInvoicesAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDInvoicesAutoPager[T]) Index() int {
	return r.run
}

type CursorIDWebhooks[T any] struct {
	Webhooks []T `json:"webhooks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Webhooks    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorIDWebhooks[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDWebhooks[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDWebhooks[T]) GetNextPage() (res *CursorIDWebhooks[T], err error) {
	if len(r.Webhooks) == 0 {
		return nil, nil
	}
	items := r.Webhooks
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("start_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorIDWebhooks[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDWebhooks[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDWebhooksAutoPager[T any] struct {
	page *CursorIDWebhooks[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDWebhooksAutoPager[T any](page *CursorIDWebhooks[T], err error) *CursorIDWebhooksAutoPager[T] {
	return &CursorIDWebhooksAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDWebhooksAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Webhooks) == 0 {
		return false
	}
	if r.idx >= len(r.page.Webhooks) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Webhooks) == 0 {
			return false
		}
	}
	r.cur = r.page.Webhooks[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDWebhooksAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDWebhooksAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDWebhooksAutoPager[T]) Index() int {
	return r.run
}

type CursorIDCategories[T any] struct {
	Categories []T `json:"categories"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorIDCategories[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDCategories[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDCategories[T]) GetNextPage() (res *CursorIDCategories[T], err error) {
	if len(r.Categories) == 0 {
		return nil, nil
	}
	items := r.Categories
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("start_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorIDCategories[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDCategories[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDCategoriesAutoPager[T any] struct {
	page *CursorIDCategories[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDCategoriesAutoPager[T any](page *CursorIDCategories[T], err error) *CursorIDCategoriesAutoPager[T] {
	return &CursorIDCategoriesAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDCategoriesAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Categories) == 0 {
		return false
	}
	if r.idx >= len(r.page.Categories) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Categories) == 0 {
			return false
		}
	}
	r.cur = r.page.Categories[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDCategoriesAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDCategoriesAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDCategoriesAutoPager[T]) Index() int {
	return r.run
}

type CursorIDStatements[T any] struct {
	Statements []T `json:"statements"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Statements  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorIDStatements[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDStatements[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDStatements[T]) GetNextPage() (res *CursorIDStatements[T], err error) {
	if len(r.Statements) == 0 {
		return nil, nil
	}
	items := r.Statements
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("start_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorIDStatements[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDStatements[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDStatementsAutoPager[T any] struct {
	page *CursorIDStatements[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDStatementsAutoPager[T any](page *CursorIDStatements[T], err error) *CursorIDStatementsAutoPager[T] {
	return &CursorIDStatementsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDStatementsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Statements) == 0 {
		return false
	}
	if r.idx >= len(r.page.Statements) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Statements) == 0 {
			return false
		}
	}
	r.cur = r.page.Statements[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDStatementsAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDStatementsAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDStatementsAutoPager[T]) Index() int {
	return r.run
}

type CursorIDRecipients[T any] struct {
	Recipients []T `json:"recipients"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Recipients  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorIDRecipients[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDRecipients[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDRecipients[T]) GetNextPage() (res *CursorIDRecipients[T], err error) {
	if len(r.Recipients) == 0 {
		return nil, nil
	}
	items := r.Recipients
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("start_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorIDRecipients[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDRecipients[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDRecipientsAutoPager[T any] struct {
	page *CursorIDRecipients[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDRecipientsAutoPager[T any](page *CursorIDRecipients[T], err error) *CursorIDRecipientsAutoPager[T] {
	return &CursorIDRecipientsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDRecipientsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Recipients) == 0 {
		return false
	}
	if r.idx >= len(r.page.Recipients) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Recipients) == 0 {
			return false
		}
	}
	r.cur = r.page.Recipients[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDRecipientsAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDRecipientsAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDRecipientsAutoPager[T]) Index() int {
	return r.run
}

type CursorIDUsers[T any] struct {
	Users []T `json:"users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorIDUsers[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDUsers[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDUsers[T]) GetNextPage() (res *CursorIDUsers[T], err error) {
	if len(r.Users) == 0 {
		return nil, nil
	}
	items := r.Users
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("start_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorIDUsers[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDUsers[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDUsersAutoPager[T any] struct {
	page *CursorIDUsers[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDUsersAutoPager[T any](page *CursorIDUsers[T], err error) *CursorIDUsersAutoPager[T] {
	return &CursorIDUsersAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDUsersAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Users) == 0 {
		return false
	}
	if r.idx >= len(r.page.Users) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Users) == 0 {
			return false
		}
	}
	r.cur = r.page.Users[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDUsersAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDUsersAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDUsersAutoPager[T]) Index() int {
	return r.run
}

type CursorIDAttachments[T any] struct {
	Attachments []T `json:"attachments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attachments respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorIDAttachments[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDAttachments[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDAttachments[T]) GetNextPage() (res *CursorIDAttachments[T], err error) {
	if len(r.Attachments) == 0 {
		return nil, nil
	}
	items := r.Attachments
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("start_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CursorIDAttachments[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDAttachments[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDAttachmentsAutoPager[T any] struct {
	page *CursorIDAttachments[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDAttachmentsAutoPager[T any](page *CursorIDAttachments[T], err error) *CursorIDAttachmentsAutoPager[T] {
	return &CursorIDAttachmentsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDAttachmentsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Attachments) == 0 {
		return false
	}
	if r.idx >= len(r.page.Attachments) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Attachments) == 0 {
			return false
		}
	}
	r.cur = r.page.Attachments[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorIDAttachmentsAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDAttachmentsAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDAttachmentsAutoPager[T]) Index() int {
	return r.run
}
