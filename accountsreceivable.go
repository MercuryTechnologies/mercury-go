// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// Manage invoices
//
// AccountsReceivableService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountsReceivableService] method instead.
type AccountsReceivableService struct {
	Options []option.RequestOption
	// Manage customers
	Customers AccountsReceivableCustomerService
	// Manage invoices
	Invoices AccountsReceivableInvoiceService
}

// NewAccountsReceivableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountsReceivableService(opts ...option.RequestOption) (r AccountsReceivableService) {
	r = AccountsReceivableService{}
	r.Options = opts
	r.Customers = NewAccountsReceivableCustomerService(opts...)
	r.Invoices = NewAccountsReceivableInvoiceService(opts...)
	return
}

// Retrieve attachment details including download URL
func (r *AccountsReceivableService) GetAttachment(ctx context.Context, attachmentID string, opts ...option.RequestOption) (res *Attachment, err error) {
	opts = slices.Concat(r.Options, opts)
	if attachmentID == "" {
		err = errors.New("missing required attachmentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("ar/attachments/%s", attachmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The object representing a file attachment for an invoice. The file is not a part
// of this object itself but information for where to download it will be in this
// object.
type Attachment struct {
	// ID for the attachment.
	ID string `json:"id" api:"required" format:"uuid"`
	// The filename for the file.
	FileName string `json:"fileName" api:"required"`
	// The signed download URL for the file itself.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FileName    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Attachment) RawJSON() string { return r.JSON.raw }
func (r *Attachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
