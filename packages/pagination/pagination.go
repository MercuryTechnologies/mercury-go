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

type CursorIDPage[T any] struct {
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
func (r CursorIDPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CursorIDPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CursorIDPage[T]) GetNextPage() (res *CursorIDPage[T], err error) {
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

func (r *CursorIDPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CursorIDPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CursorIDPageAutoPager[T any] struct {
	page *CursorIDPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCursorIDPageAutoPager[T any](page *CursorIDPage[T], err error) *CursorIDPageAutoPager[T] {
	return &CursorIDPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CursorIDPageAutoPager[T]) Next() bool {
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

func (r *CursorIDPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CursorIDPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CursorIDPageAutoPager[T]) Index() int {
	return r.run
}
