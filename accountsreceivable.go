// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"github.com/MercuryTechnologies/mercury-go/option"
)

// AccountsReceivableService contains methods and other services that help with
// interacting with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountsReceivableService] method instead.
type AccountsReceivableService struct {
	Options []option.RequestOption
	// Manage invoices
	Attachments AccountsReceivableAttachmentService
	// Manage customers
	Customers AccountsReceivableCustomerService
}

// NewAccountsReceivableService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountsReceivableService(opts ...option.RequestOption) (r AccountsReceivableService) {
	r = AccountsReceivableService{}
	r.Options = opts
	r.Attachments = NewAccountsReceivableAttachmentService(opts...)
	r.Customers = NewAccountsReceivableCustomerService(opts...)
	return
}
