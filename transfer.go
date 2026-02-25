// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"net/http"
	"slices"

	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// TransferService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransferService] method instead.
type TransferService struct {
	Options []option.RequestOption
}

// NewTransferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransferService(opts ...option.RequestOption) (r TransferService) {
	r = TransferService{}
	r.Options = opts
	return
}

// Transfer funds between two accounts within the same organization. Creates paired
// debit and credit transactions.
func (r *TransferService) New(ctx context.Context, body TransferNewParams, opts ...option.RequestOption) (res *TransferNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "transfer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response for POST /api/v1/transfer endpoint. Returns both the credit and debit
// transactions for the transfer.
type TransferNewResponse struct {
	CreditTransaction Transaction `json:"creditTransaction" api:"required"`
	DebitTransaction  Transaction `json:"debitTransaction" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditTransaction respjson.Field
		DebitTransaction  respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TransferNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferNewParams struct {
	// A positive dollar amount with at least 1 cent.
	Amount float64 `json:"amount" api:"required"`
	// ID for a Mercury account.
	DestinationAccountID string `json:"destinationAccountId" api:"required" format:"uuid"`
	IdempotencyKey       string `json:"idempotencyKey" api:"required"`
	// ID for a Mercury account.
	SourceAccountID string            `json:"sourceAccountId" api:"required" format:"uuid"`
	Note            param.Opt[string] `json:"note,omitzero"`
	paramObj
}

func (r TransferNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TransferNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
