// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/mercury-go/internal/apijson"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// Manage credit accounts
//
// CreditService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCreditService] method instead.
type CreditService struct {
	Options []option.RequestOption
}

// NewCreditService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCreditService(opts ...option.RequestOption) (r CreditService) {
	r = CreditService{}
	r.Options = opts
	return
}

// Retrieve a list of all credit accounts for the organization.
func (r *CreditService) List(ctx context.Context, opts ...option.RequestOption) (res *CreditListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "credit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CreditListResponse struct {
	Accounts []CreditListResponseAccount `json:"accounts" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accounts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditListResponse) RawJSON() string { return r.JSON.raw }
func (r *CreditListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditListResponseAccount struct {
	// ID for a Mercury account.
	ID               string  `json:"id" api:"required" format:"uuid"`
	AvailableBalance float64 `json:"availableBalance" api:"required"`
	CreatedAt        string  `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	CurrentBalance   float64 `json:"currentBalance" api:"required"`
	// Any of "active", "deleted", "pending", "archived".
	Status AccountStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		AvailableBalance respjson.Field
		CreatedAt        respjson.Field
		CurrentBalance   respjson.Field
		Status           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditListResponseAccount) RawJSON() string { return r.JSON.raw }
func (r *CreditListResponseAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
