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
	"slices"

	"github.com/MercuryTechnologies/mercury-go/internal/apiform"
	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// Manage transactions
//
// TransactionAttachmentService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionAttachmentService] method instead.
type TransactionAttachmentService struct {
	Options []option.RequestOption
}

// NewTransactionAttachmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionAttachmentService(opts ...option.RequestOption) (r TransactionAttachmentService) {
	r = TransactionAttachmentService{}
	r.Options = opts
	return
}

// Upload a file attachment to a transaction. The file is uploaded via
// multipart/form-data. Supported file types include PDF, images (PNG, JPG, GIF),
// and common document formats.
func (r *TransactionAttachmentService) Attach(ctx context.Context, transactionID string, body TransactionAttachmentAttachParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return err
	}
	path := fmt.Sprintf("transaction/%s/attachments", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type TransactionAttachment struct {
	// Any of "checkImage", "receipt", "other".
	AttachmentType TransactionAttachmentAttachmentType `json:"attachmentType" api:"required"`
	FileName       string                              `json:"fileName" api:"required"`
	URL            string                              `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachmentType respjson.Field
		FileName       respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionAttachment) RawJSON() string { return r.JSON.raw }
func (r *TransactionAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionAttachmentAttachmentType string

const (
	TransactionAttachmentAttachmentTypeCheckImage TransactionAttachmentAttachmentType = "checkImage"
	TransactionAttachmentAttachmentTypeReceipt    TransactionAttachmentAttachmentType = "receipt"
	TransactionAttachmentAttachmentTypeOther      TransactionAttachmentAttachmentType = "other"
)

type TransactionAttachmentAttachParams struct {
	// The file to upload
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	// Type of attachment: 'receipt', 'bill', or 'other'. Defaults to 'other'.
	//
	// Any of "receipt", "bill", "other".
	AttachmentType TransactionAttachmentAttachParamsAttachmentType `json:"attachmentType,omitzero"`
	paramObj
}

func (r TransactionAttachmentAttachParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type TransactionAttachmentAttachParamsAttachmentType string

const (
	TransactionAttachmentAttachParamsAttachmentTypeReceipt TransactionAttachmentAttachParamsAttachmentType = "receipt"
	TransactionAttachmentAttachParamsAttachmentTypeBill    TransactionAttachmentAttachParamsAttachmentType = "bill"
	TransactionAttachmentAttachParamsAttachmentTypeOther   TransactionAttachmentAttachParamsAttachmentType = "other"
)
