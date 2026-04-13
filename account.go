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
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/pagination"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// Manage bank accounts
//
// AccountService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of accounts. Supports cursor-based pagination with
// limit, order, start_after, and end_before query parameters.
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *pagination.CursorIDAccounts[Account], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "accounts"
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

// Retrieve a paginated list of accounts. Supports cursor-based pagination with
// limit, order, start_after, and end_before query parameters.
func (r *AccountService) ListAutoPaging(ctx context.Context, query AccountListParams, opts ...option.RequestOption) *pagination.CursorIDAccountsAutoPager[Account] {
	return pagination.NewCursorIDAccountsAutoPager(r.List(ctx, query, opts...))
}

// Get account by ID
func (r *AccountService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("account/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Account struct {
	// ID for a Mercury account.
	ID                string  `json:"id" api:"required" format:"uuid"`
	AccountNumber     string  `json:"accountNumber" api:"required"`
	AvailableBalance  float64 `json:"availableBalance" api:"required"`
	CreatedAt         string  `json:"createdAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	CurrentBalance    float64 `json:"currentBalance" api:"required"`
	DashboardLink     string  `json:"dashboardLink" api:"required"`
	Kind              string  `json:"kind" api:"required"`
	LegalBusinessName string  `json:"legalBusinessName" api:"required"`
	Name              string  `json:"name" api:"required"`
	RoutingNumber     string  `json:"routingNumber" api:"required"`
	// Any of "active", "deleted", "pending", "archived".
	Status AccountStatus `json:"status" api:"required"`
	// Any of "mercury", "external", "recipient".
	Type                   AccountType `json:"type" api:"required"`
	CanReceiveTransactions bool        `json:"canReceiveTransactions" api:"nullable"`
	Nickname               string      `json:"nickname" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccountNumber          respjson.Field
		AvailableBalance       respjson.Field
		CreatedAt              respjson.Field
		CurrentBalance         respjson.Field
		DashboardLink          respjson.Field
		Kind                   respjson.Field
		LegalBusinessName      respjson.Field
		Name                   respjson.Field
		RoutingNumber          respjson.Field
		Status                 respjson.Field
		Type                   respjson.Field
		CanReceiveTransactions respjson.Field
		Nickname               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountType string

const (
	AccountTypeMercury   AccountType = "mercury"
	AccountTypeExternal  AccountType = "external"
	AccountTypeRecipient AccountType = "recipient"
)

type AccountStatus string

const (
	AccountStatusActive   AccountStatus = "active"
	AccountStatusDeleted  AccountStatus = "deleted"
	AccountStatusPending  AccountStatus = "pending"
	AccountStatusArchived AccountStatus = "archived"
)

// Represents an expense category for transaction classification.
type CategoryData struct {
	// ID for the category
	ID string `json:"id" api:"required" format:"uuid"`
	// The name of the category
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
func (r CategoryData) RawJSON() string { return r.JSON.raw }
func (r *CategoryData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MercuryCategory string

const (
	MercuryCategoryOther                MercuryCategory = "Other"
	MercuryCategoryAdvertising          MercuryCategory = "Advertising"
	MercuryCategoryAirlines             MercuryCategory = "Airlines"
	MercuryCategoryAlcoholAndBars       MercuryCategory = "AlcoholAndBars"
	MercuryCategoryBooksAndNewspaper    MercuryCategory = "BooksAndNewspaper"
	MercuryCategoryCarRental            MercuryCategory = "CarRental"
	MercuryCategoryCharity              MercuryCategory = "Charity"
	MercuryCategoryClothing             MercuryCategory = "Clothing"
	MercuryCategoryConferences          MercuryCategory = "Conferences"
	MercuryCategoryEducation            MercuryCategory = "Education"
	MercuryCategoryElectronics          MercuryCategory = "Electronics"
	MercuryCategoryEntertainment        MercuryCategory = "Entertainment"
	MercuryCategoryFacilitiesExpenses   MercuryCategory = "FacilitiesExpenses"
	MercuryCategoryFees                 MercuryCategory = "Fees"
	MercuryCategoryFoodDelivery         MercuryCategory = "FoodDelivery"
	MercuryCategoryFuelAndGas           MercuryCategory = "FuelAndGas"
	MercuryCategoryGambling             MercuryCategory = "Gambling"
	MercuryCategoryGovernmentServices   MercuryCategory = "GovernmentServices"
	MercuryCategoryGrocery              MercuryCategory = "Grocery"
	MercuryCategoryGroundTransportation MercuryCategory = "GroundTransportation"
	MercuryCategoryInsurance            MercuryCategory = "Insurance"
	MercuryCategoryInternetAndTelephone MercuryCategory = "InternetAndTelephone"
	MercuryCategoryLegal                MercuryCategory = "Legal"
	MercuryCategoryLodging              MercuryCategory = "Lodging"
	MercuryCategoryMedical              MercuryCategory = "Medical"
	MercuryCategoryMemberships          MercuryCategory = "Memberships"
	MercuryCategoryOfficeSupplies       MercuryCategory = "OfficeSupplies"
	MercuryCategoryOtherTravel          MercuryCategory = "OtherTravel"
	MercuryCategoryParking              MercuryCategory = "Parking"
	MercuryCategoryPolitical            MercuryCategory = "Political"
	MercuryCategoryProfessionalServices MercuryCategory = "ProfessionalServices"
	MercuryCategoryRestaurants          MercuryCategory = "Restaurants"
	MercuryCategoryRetail               MercuryCategory = "Retail"
	MercuryCategoryRideshareAndTaxis    MercuryCategory = "RideshareAndTaxis"
	MercuryCategoryShipping             MercuryCategory = "Shipping"
	MercuryCategorySoftware             MercuryCategory = "Software"
	MercuryCategoryTaxes                MercuryCategory = "Taxes"
	MercuryCategoryUtilities            MercuryCategory = "Utilities"
	MercuryCategoryVehicleExpenses      MercuryCategory = "VehicleExpenses"
)

type AccountListParams struct {
	// The ID of the account to end the page before (exclusive). When provided, results
	// will end just before this ID and work backwards. Use this for reverse pagination
	// or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the account to start the page after (exclusive). When provided,
	// results will begin with the account immediately following this ID. Use this for
	// standard forward pagination to get the next page of results. Cannot be combined
	// with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order AccountListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type AccountListParamsOrder string

const (
	AccountListParamsOrderAsc  AccountListParamsOrder = "asc"
	AccountListParamsOrderDesc AccountListParamsOrder = "desc"
)
