// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
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

type CursorPagePage struct {
	NextPage     string `json:"nextPage"`
	PreviousPage string `json:"previousPage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPage     respjson.Field
		PreviousPage respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CursorPagePage) RawJSON() string { return r.JSON.raw }
func (r *CursorPagePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CursorPage[T any] struct {
	Page CursorPagePage `json:"page"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorPage[T]) GetNextPage() (res *CursorPage[T], err error) {
	if len(r.data) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	if r.cfg.Request.URL.Query().Has("end_before") {
		next := r.Page.PreviousPage
		if next == "" {
			return nil, nil
		}
		err = cfg.Apply(option.WithQuery("end_before", next))
		if err != nil {
			return nil, err
		}
	} else {
		next := r.Page.NextPage
		if next == "" {
			return nil, nil
		}
		err = cfg.Apply(option.WithQuery("start_after", next))
		if err != nil {
			return nil, err
		}
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

func (r *CursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorPageAutoPager[T any] struct {
	page *CursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorPageAutoPager[T any](page *CursorPage[T], err error) *CursorPageAutoPager[T] {
	return &CursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.data) == 0 {
		return false
	}
	if r.idx >= len(r.page.data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.data) == 0 {
			return false
		}
	}
	r.cur = r.page.data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorPageAutoPager[T]) Index() int {
	return r.run
}

type OffsetPage[T any] struct {
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OffsetPage[T]) RawJSON() string { return r.JSON.raw }
func (r *OffsetPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OffsetPage[T]) GetNextPage() (res *OffsetPage[T], err error) {
	if len(r.data) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	q := cfg.Request.URL.Query()
	offset, err := strconv.ParseInt(q.Get("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	length := int64(len(r.Data))
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

func (r *OffsetPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OffsetPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetPageAutoPager[T any] struct {
	page *OffsetPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOffsetPageAutoPager[T any](page *OffsetPage[T], err error) *OffsetPageAutoPager[T] {
	return &OffsetPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.data) == 0 {
		return false
	}
	if r.idx >= len(r.page.data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.data) == 0 {
			return false
		}
	}
	r.cur = r.page.data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetPageAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetPageAutoPager[T]) Index() int {
	return r.run
}

type TreasuryCursorPage[T any] struct {
	Cursor int64 `json:"cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursor      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r TreasuryCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *TreasuryCursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TreasuryCursorPage[T]) GetNextPage() (res *TreasuryCursorPage[T], err error) {
	if len(r.data) == 0 {
		return nil, nil
	}
	next := r.Cursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
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

func (r *TreasuryCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TreasuryCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TreasuryCursorPageAutoPager[T any] struct {
	page *TreasuryCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewTreasuryCursorPageAutoPager[T any](page *TreasuryCursorPage[T], err error) *TreasuryCursorPageAutoPager[T] {
	return &TreasuryCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TreasuryCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.data) == 0 {
		return false
	}
	if r.idx >= len(r.page.data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.data) == 0 {
			return false
		}
	}
	r.cur = r.page.data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TreasuryCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *TreasuryCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *TreasuryCursorPageAutoPager[T]) Index() int {
	return r.run
}

type TokenPage[T any] struct {
	NextPageToken string `json:"next_page_token,nullable"`
	PrevPageToken string `json:"prev_page_token,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageToken respjson.Field
		PrevPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r TokenPage[T]) RawJSON() string { return r.JSON.raw }
func (r *TokenPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TokenPage[T]) GetNextPage() (res *TokenPage[T], err error) {
	if len(r.data) == 0 {
		return nil, nil
	}
	next := r.NextPageToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("page_token", next))
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

func (r *TokenPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TokenPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TokenPageAutoPager[T any] struct {
	page *TokenPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewTokenPageAutoPager[T any](page *TokenPage[T], err error) *TokenPageAutoPager[T] {
	return &TokenPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TokenPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.data) == 0 {
		return false
	}
	if r.idx >= len(r.page.data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.data) == 0 {
			return false
		}
	}
	r.cur = r.page.data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TokenPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *TokenPageAutoPager[T]) Err() error {
	return r.err
}

func (r *TokenPageAutoPager[T]) Index() int {
	return r.run
}

type TokenPageCamel[T any] struct {
	NextPageToken string `json:"nextPageToken,nullable"`
	PrevPageToken string `json:"prevPageToken,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageToken respjson.Field
		PrevPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r TokenPageCamel[T]) RawJSON() string { return r.JSON.raw }
func (r *TokenPageCamel[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *TokenPageCamel[T]) GetNextPage() (res *TokenPageCamel[T], err error) {
	if len(r.data) == 0 {
		return nil, nil
	}
	next := r.NextPageToken
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("page_token", next))
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

func (r *TokenPageCamel[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &TokenPageCamel[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type TokenPageCamelAutoPager[T any] struct {
	page *TokenPageCamel[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewTokenPageCamelAutoPager[T any](page *TokenPageCamel[T], err error) *TokenPageCamelAutoPager[T] {
	return &TokenPageCamelAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *TokenPageCamelAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.data) == 0 {
		return false
	}
	if r.idx >= len(r.page.data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.data) == 0 {
			return false
		}
	}
	r.cur = r.page.data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *TokenPageCamelAutoPager[T]) Current() T {
	return r.cur
}

func (r *TokenPageCamelAutoPager[T]) Err() error {
	return r.err
}

func (r *TokenPageCamelAutoPager[T]) Index() int {
	return r.run
}
