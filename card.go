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

// Manage bank accounts
//
// CardService contains methods and other services that help with interacting with
// the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardService] method instead.
type CardService struct {
	Options []option.RequestOption
}

// NewCardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCardService(opts ...option.RequestOption) (r CardService) {
	r = CardService{}
	r.Options = opts
	return
}

// Retrieve all debit and credit cards associated with a specific account.
func (r *CardService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *CardListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/%s/cards", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CardListResponse struct {
	Cards []CardListResponseCard `json:"cards" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cards       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardListResponse) RawJSON() string { return r.JSON.raw }
func (r *CardListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardListResponseCard struct {
	CardID         string `json:"cardId" api:"required"`
	CreatedAt      string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	LastFourDigits string `json:"lastFourDigits" api:"required"`
	NameOnCard     string `json:"nameOnCard" api:"required"`
	// Any of "visa", "mastercard".
	Network string `json:"network" api:"required"`
	// Any of "active", "frozen", "cancelled", "inactive", "expired", "suspended".
	Status string `json:"status" api:"required"`
	// Any of "inactive", "active", "locked".
	PhysicalCardStatus string `json:"physicalCardStatus" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CardID             respjson.Field
		CreatedAt          respjson.Field
		LastFourDigits     respjson.Field
		NameOnCard         respjson.Field
		Network            respjson.Field
		Status             respjson.Field
		PhysicalCardStatus respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardListResponseCard) RawJSON() string { return r.JSON.raw }
func (r *CardListResponseCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
