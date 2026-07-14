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
	shimjson "github.com/MercuryTechnologies/mercury-go/internal/encoding/json"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/pagination"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// Manage cards
//
// CardService contains methods and other services that help with interacting with
// the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardService] method instead.
type CardService struct {
	options []option.RequestOption
}

// NewCardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCardService(opts ...option.RequestOption) (r CardService) {
	r = CardService{}
	r.options = opts
	return
}

// Issue a new virtual card.
func (r *CardService) New(ctx context.Context, body CardNewParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a card's nickname or spending limits.
func (r *CardService) Update(ctx context.Context, cardID string, body CardUpdateParams, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of cards.
func (r *CardService) List(ctx context.Context, query CardListParams, opts ...option.RequestOption) (res *pagination.CursorIDCardsPagination[Card], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "cards"
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

// Retrieve a paginated list of cards.
func (r *CardService) ListAutoPaging(ctx context.Context, query CardListParams, opts ...option.RequestOption) *pagination.CursorIDCardsPaginationAutoPager[Card] {
	return pagination.NewCursorIDCardsPaginationAutoPager(r.List(ctx, query, opts...))
}

// Permanently cancel a card. This action cannot be undone.
func (r *CardService) Cancel(ctx context.Context, cardID string, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s/cancel", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Temporarily freeze a card. The card must be active.
func (r *CardService) Freeze(ctx context.Context, cardID string, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s/freeze", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Retrieve details of a specific card by its ID.
func (r *CardService) Get(ctx context.Context, cardID string, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Unfreeze a previously frozen card, restoring it to active status.
func (r *CardService) Unfreeze(ctx context.Context, cardID string, opts ...option.RequestOption) (res *Card, err error) {
	opts = slices.Concat(r.options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s/unfreeze", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type Card struct {
	// Unique identifier for the card.
	ID string `json:"id" api:"required" format:"uuid"`
	// The Mercury account this card is associated with.
	AccountID string `json:"accountId" api:"required"`
	// Mercury spend-category locks applied to this card, in no particular order. Empty
	// when the card has no category restrictions.
	CategoryLocks []MercuryCategory `json:"categoryLocks" api:"required"`
	// Timestamp when the card was issued.
	CreatedAt string `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Month and year the card expires.
	Expiration CardExpiration `json:"expiration" api:"required"`
	// Whether the card is a debit or credit card.
	//
	// Any of "debit", "credit".
	Kind CardKind `json:"kind" api:"required"`
	// Last four digits of the card's primary account number (PAN).
	LastFour string `json:"lastFour" api:"required"`
	// Cardholder name printed on the card.
	NameOnCard string `json:"nameOnCard" api:"required"`
	// Current lifecycle state of the card.
	//
	// Any of "active", "frozen", "cancelled", "inactive", "expired", "suspended".
	Status CardStatus `json:"status" api:"required"`
	// Whether the card is virtual (digital-only) or physical (printed, supports ATM).
	//
	// Any of "virtual", "physical".
	Type CardType `json:"type" api:"required"`
	// Timestamp of the last modification to the card or its settings.
	UpdatedAt string `json:"updatedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Mercury User who owns the card.
	UserID string `json:"userId" api:"required"`
	// Information about a merchant that can be used for spend controls like merchant
	// locking.
	MerchantLock CardMerchantLock `json:"merchantLock" api:"nullable"`
	// Optional user-assigned label for the card.
	Nickname string `json:"nickname" api:"nullable"`
	// Activation state of a physical card. Null for virtual cards.
	//
	// Any of "inactive", "active", "locked".
	PhysicalCardStatus CardPhysicalCardStatus `json:"physicalCardStatus" api:"nullable"`
	// Spending controls applied to a card
	SpendLimit CardSpendLimit `json:"spendLimit" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccountID          respjson.Field
		CategoryLocks      respjson.Field
		CreatedAt          respjson.Field
		Expiration         respjson.Field
		Kind               respjson.Field
		LastFour           respjson.Field
		NameOnCard         respjson.Field
		Status             respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		UserID             respjson.Field
		MerchantLock       respjson.Field
		Nickname           respjson.Field
		PhysicalCardStatus respjson.Field
		SpendLimit         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Card) RawJSON() string { return r.JSON.raw }
func (r *Card) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about a merchant that can be used for spend controls like merchant
// locking.
type CardMerchantLock struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardMerchantLock) RawJSON() string { return r.JSON.raw }
func (r *CardMerchantLock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Activation state of a physical card. Null for virtual cards.
type CardPhysicalCardStatus string

const (
	CardPhysicalCardStatusInactive CardPhysicalCardStatus = "inactive"
	CardPhysicalCardStatusActive   CardPhysicalCardStatus = "active"
	CardPhysicalCardStatusLocked   CardPhysicalCardStatus = "locked"
)

// Spending controls applied to a card
type CardSpendLimit struct {
	// Maximum total spend allowed per interval, in cents.
	AmountCents int64 `json:"amountCents" api:"required"`
	// Rolling window the limit applies to.
	//
	// Any of "daily", "weekly", "monthly", "yearly".
	Interval string `json:"interval" api:"required"`
	// Maximum ATM withdrawal allowed per interval, in cents. Null for virtual cards.
	AtmAmountCents int64 `json:"atmAmountCents" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountCents    respjson.Field
		Interval       respjson.Field
		AtmAmountCents respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardSpendLimit) RawJSON() string { return r.JSON.raw }
func (r *CardSpendLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Month and year the card expires.
type CardExpiration struct {
	// Calendar month.
	Month int64 `json:"month" api:"required"`
	// Four-digit calendar year.
	Year int64 `json:"year" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Month       respjson.Field
		Year        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardExpiration) RawJSON() string { return r.JSON.raw }
func (r *CardExpiration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardKind string

const (
	CardKindDebit  CardKind = "debit"
	CardKindCredit CardKind = "credit"
)

type CardListResponse struct {
	// List of cards in the current page.
	Cards []Card `json:"cards" api:"required"`
	// Pagination cursors for navigating to next/previous pages.
	Page CardListResponsePage `json:"page" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cards       respjson.Field
		Page        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardListResponse) RawJSON() string { return r.JSON.raw }
func (r *CardListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination cursors for navigating to next/previous pages.
type CardListResponsePage struct {
	// Unique identifier for a card
	NextPage string `json:"nextPage" format:"uuid"`
	// Unique identifier for a card
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
func (r CardListResponsePage) RawJSON() string { return r.JSON.raw }
func (r *CardListResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardStatus string

const (
	CardStatusActive    CardStatus = "active"
	CardStatusFrozen    CardStatus = "frozen"
	CardStatusCancelled CardStatus = "cancelled"
	CardStatusInactive  CardStatus = "inactive"
	CardStatusExpired   CardStatus = "expired"
	CardStatusSuspended CardStatus = "suspended"
)

type CardType string

const (
	CardTypeVirtual  CardType = "virtual"
	CardTypePhysical CardType = "physical"
)

// The properties Kind, Type, UserID are required.
type CreateCardRequestParam struct {
	// Whether to issue a debit or credit card.
	//
	// Any of "debit", "credit".
	Kind CardKind `json:"kind,omitzero" api:"required"`
	// The type of card to issue.
	//
	// Any of "virtual".
	Type CreateCardRequestType `json:"type,omitzero" api:"required"`
	// ID for the user
	UserID string `json:"userId" api:"required" format:"uuid"`
	// ID for a Mercury account.
	AccountID param.Opt[string] `json:"accountId,omitzero" format:"uuid"`
	// Optional user-assigned label for the card.
	Nickname param.Opt[string] `json:"nickname,omitzero"`
	// The spend-management state of a card. Used both as the input at issuance time
	// (the configuration the caller wants applied) and as the read-side response from
	// 'CardManagement.Issuing.CardSpendManagement.getCardSpendManagement'. All of the
	// card spend-management tables are collapsed into this single type.
	CardSpendManagementState CreateCardRequestCardSpendManagementStateParam `json:"cardSpendManagementState,omitzero"`
	// Spending controls to apply at issuance.
	SpendLimit CreateCardRequestSpendLimitParam `json:"spendLimit,omitzero"`
	paramObj
}

func (r CreateCardRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateCardRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateCardRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of card to issue.
type CreateCardRequestType string

const (
	CreateCardRequestTypeVirtual CreateCardRequestType = "virtual"
)

// The spend-management state of a card. Used both as the input at issuance time
// (the configuration the caller wants applied) and as the read-side response from
// 'CardManagement.Issuing.CardSpendManagement.getCardSpendManagement'. All of the
// card spend-management tables are collapsed into this single type.
//
// The property IsSingleUse is required.
type CreateCardRequestCardSpendManagementStateParam struct {
	// Whether this card auto-cancels itself on the first approved non-zero
	// authorization. Must be 'False' for physical cards.
	IsSingleUse bool `json:"isSingleUse" api:"required"`
	// Configuration for a card managed by the agentic spend-management agent.
	AgenticCardState CreateCardRequestCardSpendManagementStateAgenticCardStateParam `json:"agenticCardState,omitzero"`
	paramObj
}

func (r CreateCardRequestCardSpendManagementStateParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateCardRequestCardSpendManagementStateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateCardRequestCardSpendManagementStateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a card managed by the agentic spend-management agent.
type CreateCardRequestCardSpendManagementStateAgenticCardStateParam struct {
	// A dollar amount
	ManualApprovalThreshold param.Opt[float64] `json:"manualApprovalThreshold,omitzero"`
	paramObj
}

func (r CreateCardRequestCardSpendManagementStateAgenticCardStateParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateCardRequestCardSpendManagementStateAgenticCardStateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateCardRequestCardSpendManagementStateAgenticCardStateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spending controls to apply at issuance.
//
// The properties AmountCents, Interval are required.
type CreateCardRequestSpendLimitParam struct {
	// Maximum total spend allowed per interval, in cents.
	AmountCents int64 `json:"amountCents" api:"required"`
	// Rolling window the limit applies to.
	//
	// Any of "daily", "weekly", "monthly", "yearly".
	Interval string `json:"interval,omitzero" api:"required"`
	paramObj
}

func (r CreateCardRequestSpendLimitParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateCardRequestSpendLimitParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateCardRequestSpendLimitParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CreateCardRequestSpendLimitParam](
		"interval", "daily", "weekly", "monthly", "yearly",
	)
}

// At least one updateable field must be provided; requests where every field is
// the same are rejected.
//
// The property Nickname is required.
type UpdateCardRequestParam struct {
	// Nickname update action. Omit the field to keep the current nickname, send null
	// or an empty/whitespace-only string to clear it, or send a string to set it.
	Nickname param.Opt[string] `json:"nickname,omitzero" api:"required"`
	// Spending controls applied to a card
	SpendLimit UpdateCardRequestSpendLimitParam `json:"spendLimit,omitzero"`
	paramObj
}

func (r UpdateCardRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateCardRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateCardRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spending controls applied to a card
//
// The properties AmountCents, Interval are required.
type UpdateCardRequestSpendLimitParam struct {
	// Maximum total spend allowed per interval, in cents.
	AmountCents int64 `json:"amountCents" api:"required"`
	// Rolling window the limit applies to.
	//
	// Any of "daily", "weekly", "monthly", "yearly".
	Interval string `json:"interval,omitzero" api:"required"`
	// Maximum ATM withdrawal allowed per interval, in cents. Null for virtual cards.
	AtmAmountCents param.Opt[int64] `json:"atmAmountCents,omitzero"`
	paramObj
}

func (r UpdateCardRequestSpendLimitParam) MarshalJSON() (data []byte, err error) {
	type shadow UpdateCardRequestSpendLimitParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateCardRequestSpendLimitParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UpdateCardRequestSpendLimitParam](
		"interval", "daily", "weekly", "monthly", "yearly",
	)
}

type CardNewParams struct {
	CreateCardRequest CreateCardRequestParam
	paramObj
}

func (r CardNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateCardRequest)
}
func (r *CardNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardUpdateParams struct {
	// At least one updateable field must be provided; requests where every field is
	// the same are rejected.
	UpdateCardRequest UpdateCardRequestParam
	paramObj
}

func (r CardUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.UpdateCardRequest)
}
func (r *CardUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardListParams struct {
	// The ID of the card to end the page before (exclusive). When provided, results
	// will end just before this ID and work backwards. Use this for reverse pagination
	// or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 500
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the card to start the page after (exclusive). When provided, results
	// will begin with the card immediately following this ID. Use this for standard
	// forward pagination to get the next page of results. Cannot be combined with
	// end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Filter cards by the cardholder's user ID.
	UserID param.Opt[string] `query:"userId,omitzero" json:"-"`
	// Filter cards by one or more account IDs.
	AccountID []string `query:"accountId,omitzero" json:"-"`
	// Filter cards by kind (debit or credit).
	//
	// Any of "debit", "credit".
	Kind []string `query:"kind,omitzero" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order CardListParamsOrder `query:"order,omitzero" json:"-"`
	// Filter cards by one or more statuses.
	//
	// Any of "active", "frozen", "cancelled", "inactive", "expired", "suspended".
	Status []string `query:"status,omitzero" json:"-"`
	// Filter cards by type (virtual or physical).
	//
	// Any of "virtual", "physical".
	Type []string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CardListParams]'s query parameters as `url.Values`.
func (r CardListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type CardListParamsOrder string

const (
	CardListParamsOrderAsc  CardListParamsOrder = "asc"
	CardListParamsOrderDesc CardListParamsOrder = "desc"
)
