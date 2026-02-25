// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"reflect"
	"strconv"

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

type OffsetAccountTransactions[T any] struct {
	Total        int64 `json:"total"`
	Transactions []T   `json:"transactions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Total        respjson.Field
		Transactions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OffsetAccountTransactions[T]) RawJSON() string { return r.JSON.raw }
func (r *OffsetAccountTransactions[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OffsetAccountTransactions[T]) GetNextPage() (res *OffsetAccountTransactions[T], err error) {
	if len(r.Transactions) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	q := cfg.Request.URL.Query()
	offset, err := strconv.ParseInt(q.Get("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	length := int64(len(r.Transactions))
	next := offset + length

	if next < r.Total && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
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

func (r *OffsetAccountTransactions[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OffsetAccountTransactions[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetAccountTransactionsAutoPager[T any] struct {
	page *OffsetAccountTransactions[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOffsetAccountTransactionsAutoPager[T any](page *OffsetAccountTransactions[T], err error) *OffsetAccountTransactionsAutoPager[T] {
	return &OffsetAccountTransactionsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetAccountTransactionsAutoPager[T]) Next() bool {
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

func (r *OffsetAccountTransactionsAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetAccountTransactionsAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetAccountTransactionsAutoPager[T]) Index() int {
	return r.run
}

type CursorIDArCustomers[T any] struct {
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
func (r CursorIDArCustomers[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDArCustomers[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDArCustomers[T]) GetNextPage() (res *CursorIDArCustomers[T], err error) {
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

func (r *CursorIDArCustomers[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDArCustomers[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDArCustomersAutoPager[T any] struct {
	page *CursorIDArCustomers[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDArCustomersAutoPager[T any](page *CursorIDArCustomers[T], err error) *CursorIDArCustomersAutoPager[T] {
	return &CursorIDArCustomersAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDArCustomersAutoPager[T]) Next() bool {
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

func (r *CursorIDArCustomersAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDArCustomersAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDArCustomersAutoPager[T]) Index() int {
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
