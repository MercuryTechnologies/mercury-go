// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"github.com/MercuryTechnologies/mercury-go/option"
)

// SpendService contains methods and other services that help with interacting with
// the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpendService] method instead.
type SpendService struct {
	options []option.RequestOption
	// Manage Spend Management budgets
	Budgets SpendBudgetService
	// Manage Spend Management receipts
	Receipts SpendReceiptService
	// Manage Spend Management reimbursements. Automated access is permitted, including
	// from apps and AI agents. You're responsible for your API tokens and for the apps
	// and agents that use them. See
	// [Getting Started](https://docs.mercury.com/docs/getting-started) and sections 2
	// (Our Proprietary Rights) and 3.2 (Third-Party AI and Connected Services) of the
	// [Terms of Use](https://mercury.com/legal/terms).
	Reimbursements SpendReimbursementService
}

// NewSpendService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSpendService(opts ...option.RequestOption) (r SpendService) {
	r = SpendService{}
	r.options = opts
	r.Budgets = NewSpendBudgetService(opts...)
	r.Receipts = NewSpendReceiptService(opts...)
	r.Reimbursements = NewSpendReimbursementService(opts...)
	return
}
