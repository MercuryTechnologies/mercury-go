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

// Manage Spend Management receipts
//
// SpendReceiptService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpendReceiptService] method instead.
type SpendReceiptService struct {
	options []option.RequestOption
}

// NewSpendReceiptService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpendReceiptService(opts ...option.RequestOption) (r SpendReceiptService) {
	r = SpendReceiptService{}
	r.options = opts
	return
}

// Retrieve directly uploaded receipts whose matching result is pending, ordered by
// creation time, ascending (oldest first) by default. Upload URL requests do not
// appear until matching starts. Requires see-all expense permission. Use GET
// /spend/receipts/{receiptId} to poll terminal outcomes. Supports cursor-based
// pagination with limit, order, start_at, start_after, and end_before query
// parameters.
func (r *SpendReceiptService) List(ctx context.Context, query SpendReceiptListParams, opts ...option.RequestOption) (res *pagination.CursorIDReceipts[Receipt], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "spend/receipts"
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

// Retrieve directly uploaded receipts whose matching result is pending, ordered by
// creation time, ascending (oldest first) by default. Upload URL requests do not
// appear until matching starts. Requires see-all expense permission. Use GET
// /spend/receipts/{receiptId} to poll terminal outcomes. Supports cursor-based
// pagination with limit, order, start_at, start_after, and end_before query
// parameters.
func (r *SpendReceiptService) ListAutoPaging(ctx context.Context, query SpendReceiptListParams, opts ...option.RequestOption) *pagination.CursorIDReceiptsAutoPager[Receipt] {
	return pagination.NewCursorIDReceiptsAutoPager(r.List(ctx, query, opts...))
}

// Request a presigned URL to upload a receipt file. The URL expires after five
// minutes; replay the request with the same Idempotency-Key to receive a fresh URL
// for the same receipt. Upload the file, then start matching via the match
// endpoint.
func (r *SpendReceiptService) NewUploadURL(ctx context.Context, params SpendReceiptNewUploadURLParams, opts ...option.RequestOption) (res *ReceiptUploadURL, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	path := "spend/receipts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a directly uploaded receipt and its auto-matching status by id.
// Uploaders can retrieve their own receipts; reading another uploader's receipt
// requires see-all expense permission.
func (r *SpendReceiptService) Get(ctx context.Context, receiptID string, opts ...option.RequestOption) (res *Receipt, err error) {
	opts = slices.Concat(r.options, opts)
	if receiptID == "" {
		err = errors.New("missing required receiptId parameter")
		return nil, err
	}
	path := fmt.Sprintf("spend/receipts/%s", receiptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Start matching an uploaded receipt against the organization's transactions
// (OCR + matcher). Only the uploader can start matching. Poll GET
// /spend/receipts/{receiptId} for the outcome. Returns 409 when matching has
// already started. Returns 409 when another request restarted matching, and 502
// when processing fails; retry both with a new Idempotency-Key.
func (r *SpendReceiptService) Match(ctx context.Context, receiptID string, body SpendReceiptMatchParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", body.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if receiptID == "" {
		err = errors.New("missing required receiptId parameter")
		return err
	}
	path := fmt.Sprintf("spend/receipts/%s/match", receiptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Public view of a receipt and its matching state.
type Receipt struct {
	// ID for the receipt
	ID        string `json:"id" api:"required" format:"uuid"`
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	FileName  string `json:"fileName" api:"required"`
	// Any of "pending", "matched", "unmatched", "notAttempted", "failed".
	Status ReceiptStatus `json:"status" api:"required"`
	// ID for this transaction
	MatchedTransactionID string `json:"matchedTransactionId" api:"nullable" format:"uuid"`
	// When Mercury began processing the direct upload; absent until matching starts.
	UploadedAt string `json:"uploadedAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		FileName             respjson.Field
		Status               respjson.Field
		MatchedTransactionID respjson.Field
		UploadedAt           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Receipt) RawJSON() string { return r.JSON.raw }
func (r *Receipt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReceiptStatus string

const (
	ReceiptStatusPending      ReceiptStatus = "pending"
	ReceiptStatusMatched      ReceiptStatus = "matched"
	ReceiptStatusUnmatched    ReceiptStatus = "unmatched"
	ReceiptStatusNotAttempted ReceiptStatus = "notAttempted"
	ReceiptStatusFailed       ReceiptStatus = "failed"
)

// The presigned upload URL plus the id of the receipt now being tracked. Upload
// the file to @upload@, then start matching via the match endpoint.
type ReceiptUploadURL struct {
	// ID for the receipt
	ID string `json:"id" api:"required" format:"uuid"`
	// The presigned URL to upload the receipt file to.
	Upload string `json:"upload" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Upload      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReceiptUploadURL) RawJSON() string { return r.JSON.raw }
func (r *ReceiptUploadURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendReceiptListParams struct {
	// The ID of the receipt to end the page before (exclusive). When provided, results
	// will end just before this ID and work backwards. Use this for reverse pagination
	// or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the receipt to start the page after (exclusive). When provided,
	// results will begin with the receipt immediately following this ID. Use this for
	// standard forward pagination to get the next page of results. Cannot be combined
	// with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// The ID of the receipt to start the page at (inclusive). Cannot be combined with
	// start_after or end_before.
	StartAt param.Opt[string] `query:"start_at,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order SpendReceiptListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SpendReceiptListParams]'s query parameters as `url.Values`.
func (r SpendReceiptListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type SpendReceiptListParamsOrder string

const (
	SpendReceiptListParamsOrderAsc  SpendReceiptListParamsOrder = "asc"
	SpendReceiptListParamsOrderDesc SpendReceiptListParamsOrder = "desc"
)

type SpendReceiptNewUploadURLParams struct {
	FileName       string            `json:"fileName" api:"required"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}

func (r SpendReceiptNewUploadURLParams) MarshalJSON() (data []byte, err error) {
	type shadow SpendReceiptNewUploadURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpendReceiptNewUploadURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpendReceiptMatchParams struct {
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	paramObj
}
