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

	"github.com/MercuryTechnologies/mercury-go/internal/apiform"
	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	"github.com/MercuryTechnologies/mercury-go/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/pagination"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// Manage payment recipients
//
// RecipientAttachmentService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecipientAttachmentService] method instead.
type RecipientAttachmentService struct {
	Options []option.RequestOption
}

// NewRecipientAttachmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRecipientAttachmentService(opts ...option.RequestOption) (r RecipientAttachmentService) {
	r = RecipientAttachmentService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of all recipient tax form attachments across all
// recipients in the organization. Use cursor parameters (start_after, end_before)
// for pagination.
func (r *RecipientAttachmentService) List(ctx context.Context, query RecipientAttachmentListParams, opts ...option.RequestOption) (res *pagination.CursorIDRecipientAttachments[RecipientAttachmentListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "recipients/attachments"
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

// Retrieve a paginated list of all recipient tax form attachments across all
// recipients in the organization. Use cursor parameters (start_after, end_before)
// for pagination.
func (r *RecipientAttachmentService) ListAutoPaging(ctx context.Context, query RecipientAttachmentListParams, opts ...option.RequestOption) *pagination.CursorIDRecipientAttachmentsAutoPager[RecipientAttachmentListResponse] {
	return pagination.NewCursorIDRecipientAttachmentsAutoPager(r.List(ctx, query, opts...))
}

// Upload a tax form attachment for a recipient. The file is uploaded via
// multipart/form-data. Supported file types include PDF, images (PNG, JPG, GIF),
// and common document formats. The attachment will be associated as a tax document
// for the recipient.
func (r *RecipientAttachmentService) Attach(ctx context.Context, recipientID string, body RecipientAttachmentAttachParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if recipientID == "" {
		err = errors.New("missing required recipientId parameter")
		return err
	}
	path := fmt.Sprintf("recipient/%s/attachments", recipientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type RecipientAttachment struct {
	// Name of the uploaded file
	FileName string `json:"fileName" api:"required"`
	// Timestamp when the attachment was uploaded
	UploadedAt string `json:"uploadedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Presigned URL to download the attachment (valid for 12 hours)
	URL string `json:"url" api:"required"`
	// The tax form type (W-9 for US persons, W-8BEN for foreign individuals, W-8BEN-E
	// for foreign entities)
	//
	// Any of "w9", "w8BEN", "w8BENE", "unknown".
	FormType TaxFormType `json:"formType" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileName    respjson.Field
		UploadedAt  respjson.Field
		URL         respjson.Field
		FormType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientAttachment) RawJSON() string { return r.JSON.raw }
func (r *RecipientAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientAttachmentListResponse struct {
	// ID for the recipient tax form attachment
	ID string `json:"id" api:"required" format:"uuid"`
	// Name of the uploaded file
	FileName string `json:"fileName" api:"required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId" api:"required" format:"uuid"`
	// Timestamp when the attachment was uploaded
	UploadedAt string `json:"uploadedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Presigned URL to download the attachment (valid for 12 hours)
	URL string `json:"url" api:"required"`
	// The tax form type (W-9, W-8BEN, W-8BEN-E, or Unknown)
	//
	// Any of "w9", "w8BEN", "w8BENE", "unknown".
	FormType TaxFormType `json:"formType" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FileName    respjson.Field
		RecipientID respjson.Field
		UploadedAt  respjson.Field
		URL         respjson.Field
		FormType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientAttachmentListResponse) RawJSON() string { return r.JSON.raw }
func (r *RecipientAttachmentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientAttachmentListParams struct {
	// The ID of the recipient attachment to end the page before (exclusive). When
	// provided, results will end just before this ID and work backwards. Use this for
	// reverse pagination or to retrieve previous pages. Cannot be combined with
	// start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the recipient attachment to start the page after (exclusive). When
	// provided, results will begin with the recipient attachment immediately following
	// this ID. Use this for standard forward pagination to get the next page of
	// results. Cannot be combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order RecipientAttachmentListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecipientAttachmentListParams]'s query parameters as
// `url.Values`.
func (r RecipientAttachmentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type RecipientAttachmentListParamsOrder string

const (
	RecipientAttachmentListParamsOrderAsc  RecipientAttachmentListParamsOrder = "asc"
	RecipientAttachmentListParamsOrderDesc RecipientAttachmentListParamsOrder = "desc"
)

type RecipientAttachmentAttachParams struct {
	// The file to upload (tax form document)
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r RecipientAttachmentAttachParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
