// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/mercury-go/internal/apiform"
	"github.com/stainless-sdks/mercury-go/internal/apijson"
	"github.com/stainless-sdks/mercury-go/internal/apiquery"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/param"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// TransactionService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	Options []option.RequestOption
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r TransactionService) {
	r = TransactionService{}
	r.Options = opts
	return
}

// Retrieve a single transaction by its ID. Returns full transaction details
// including attachments, check images, and related metadata.
func (r *TransactionService) Get(ctx context.Context, transactionID string, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transaction/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the note and/or category of an existing transaction. Use null values to
// clear existing data.
func (r *TransactionService) Update(ctx context.Context, transactionID string, body TransactionUpdateParams, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transaction/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of all transactions across all accounts. Supports
// advanced filtering by date ranges, status, categories, and cursor-based
// pagination.
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *TransactionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload a file attachment to a transaction. The file is uploaded via
// multipart/form-data. Supported file types include PDF, images (PNG, JPG, GIF),
// and common document formats.
func (r *TransactionService) UploadAttachment(ctx context.Context, transactionID string, body TransactionUploadAttachmentParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transaction/%s/attachments", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type TransactionListResponse struct {
	Page         TransactionListResponsePage `json:"page,required"`
	Transactions []Transaction               `json:"transactions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page         respjson.Field
		Transactions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListResponsePage struct {
	NextPage     string `json:"nextPage" format:"uuid"`
	PreviousPage string `json:"previousPage" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPage     respjson.Field
		PreviousPage respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionListResponsePage) RawJSON() string { return r.JSON.raw }
func (r *TransactionListResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionUpdateParams struct {
	// Note update action. Omit field to keep current note, send null or empty string
	// to clear note, send text to set note.
	Note param.Opt[string] `json:"note,omitzero,required"`
	// ID for the category
	CategoryID string `json:"categoryId,required" format:"uuid"`
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
	// Latest createdAt date to filter for. If it's not provided, it defaults to
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

type TransactionUploadAttachmentParams struct {
	// The file to upload
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Type of attachment: 'receipt', 'bill', or 'other'. Defaults to 'other'.
	//
	// Any of "receipt", "bill", "other".
	AttachmentType TransactionUploadAttachmentParamsAttachmentType `json:"attachmentType,omitzero"`
	paramObj
}

func (r TransactionUploadAttachmentParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Type of attachment: 'receipt', 'bill', or 'other'. Defaults to 'other'.
type TransactionUploadAttachmentParamsAttachmentType string

const (
	TransactionUploadAttachmentParamsAttachmentTypeReceipt TransactionUploadAttachmentParamsAttachmentType = "receipt"
	TransactionUploadAttachmentParamsAttachmentTypeBill    TransactionUploadAttachmentParamsAttachmentType = "bill"
	TransactionUploadAttachmentParamsAttachmentTypeOther   TransactionUploadAttachmentParamsAttachmentType = "other"
)
