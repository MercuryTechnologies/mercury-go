// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	"github.com/MercuryTechnologies/mercury-go/internal/apiquery"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
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
	options []option.RequestOption
}

// NewCreditService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCreditService(opts ...option.RequestOption) (r CreditService) {
	r = CreditService{}
	r.options = opts
	return
}

// Retrieve a list of all credit accounts for the organization.
func (r *CreditService) List(ctx context.Context, opts ...option.RequestOption) (res *CreditListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "credit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve historical IO credit statements with cursor-based pagination. Each
// statement includes its billing period (startDate and endDate), closing balance
// (endingBalance), and opening balance (details.accountDetails.startingBalance),
// calculated using the same strict posting-date boundaries as the PDF. Balances
// use signed account amounts: amounts owed are negative. Filter by statement start
// date with start and end. Use getAccountStatements for depository accounts.
func (r *CreditService) GetStatements(ctx context.Context, accountID string, query CreditGetStatementsParams, opts ...option.RequestOption) (res *CreditGetStatementsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return nil, err
	}
	path := fmt.Sprintf("credit/%s/statements", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

type CreditGetStatementsResponse struct {
	Page       CreditGetStatementsResponsePage        `json:"page" api:"required"`
	Statements []CreditGetStatementsResponseStatement `json:"statements" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		Statements  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditGetStatementsResponse) RawJSON() string { return r.JSON.raw }
func (r *CreditGetStatementsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditGetStatementsResponsePage struct {
	// ID for the account statement
	NextPage string `json:"nextPage" format:"uuid"`
	// ID for the account statement
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
func (r CreditGetStatementsResponsePage) RawJSON() string { return r.JSON.raw }
func (r *CreditGetStatementsResponsePage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditGetStatementsResponseStatement struct {
	// ID for the account statement
	ID                  string  `json:"id" api:"required" format:"uuid"`
	CompanyLegalAddress Address `json:"companyLegalAddress" api:"required"`
	CompanyLegalName    string  `json:"companyLegalName" api:"required"`
	// Account statement payload with type discriminator
	Details     CreditGetStatementsResponseStatementDetails `json:"details" api:"required"`
	DownloadURL string                                      `json:"downloadUrl" api:"required"`
	EndDate     string                                      `json:"endDate" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// A dollar amount
	EndingBalance float64                       `json:"endingBalance" api:"required"`
	StartDate     string                        `json:"startDate" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	Transactions  []AccountStatementTransaction `json:"transactions" api:"required"`
	Ein           string                        `json:"ein" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CompanyLegalAddress respjson.Field
		CompanyLegalName    respjson.Field
		Details             respjson.Field
		DownloadURL         respjson.Field
		EndDate             respjson.Field
		EndingBalance       respjson.Field
		StartDate           respjson.Field
		Transactions        respjson.Field
		Ein                 respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditGetStatementsResponseStatement) RawJSON() string { return r.JSON.raw }
func (r *CreditGetStatementsResponseStatement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Account statement payload with type discriminator
type CreditGetStatementsResponseStatementDetails struct {
	// Account-specific details based on accountType
	AccountDetails CreditGetStatementsResponseStatementDetailsAccountDetailsUnion `json:"accountDetails" api:"required"`
	// Any of "depository", "credit".
	AccountType string `json:"accountType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountDetails respjson.Field
		AccountType    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditGetStatementsResponseStatementDetails) RawJSON() string { return r.JSON.raw }
func (r *CreditGetStatementsResponseStatementDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CreditGetStatementsResponseStatementDetailsAccountDetailsUnion contains all
// possible properties and values from
// [CreditGetStatementsResponseStatementDetailsAccountDetailsDepositoryAccountStatementPayload],
// [CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CreditGetStatementsResponseStatementDetailsAccountDetailsUnion struct {
	// This field is from variant
	// [CreditGetStatementsResponseStatementDetailsAccountDetailsDepositoryAccountStatementPayload].
	AccountNumber string `json:"accountNumber"`
	// This field is from variant
	// [CreditGetStatementsResponseStatementDetailsAccountDetailsDepositoryAccountStatementPayload].
	RoutingNumber string `json:"routingNumber"`
	// This field is from variant
	// [CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload].
	CreditLimit float64 `json:"creditLimit"`
	// This field is from variant
	// [CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload].
	CreditPeriodID string `json:"creditPeriodId"`
	// This field is from variant
	// [CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload].
	DueDate string `json:"dueDate"`
	// This field is from variant
	// [CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload].
	StartingBalance float64 `json:"startingBalance"`
	// This field is from variant
	// [CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload].
	TotalAutopayPayments float64 `json:"totalAutopayPayments"`
	// This field is from variant
	// [CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload].
	TotalCashbackEarned float64 `json:"totalCashbackEarned"`
	// This field is from variant
	// [CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload].
	TotalOneTimePayments float64 `json:"totalOneTimePayments"`
	// This field is from variant
	// [CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload].
	TotalTransactions float64 `json:"totalTransactions"`
	JSON              struct {
		AccountNumber        respjson.Field
		RoutingNumber        respjson.Field
		CreditLimit          respjson.Field
		CreditPeriodID       respjson.Field
		DueDate              respjson.Field
		StartingBalance      respjson.Field
		TotalAutopayPayments respjson.Field
		TotalCashbackEarned  respjson.Field
		TotalOneTimePayments respjson.Field
		TotalTransactions    respjson.Field
		raw                  string
	} `json:"-"`
}

func (u CreditGetStatementsResponseStatementDetailsAccountDetailsUnion) AsCreditGetStatementsResponseStatementDetailsAccountDetailsDepositoryAccountStatementPayload() (v CreditGetStatementsResponseStatementDetailsAccountDetailsDepositoryAccountStatementPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CreditGetStatementsResponseStatementDetailsAccountDetailsUnion) AsCreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload() (v CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CreditGetStatementsResponseStatementDetailsAccountDetailsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CreditGetStatementsResponseStatementDetailsAccountDetailsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditGetStatementsResponseStatementDetailsAccountDetailsDepositoryAccountStatementPayload struct {
	AccountNumber string `json:"accountNumber" api:"required"`
	RoutingNumber string `json:"routingNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountNumber respjson.Field
		RoutingNumber respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditGetStatementsResponseStatementDetailsAccountDetailsDepositoryAccountStatementPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *CreditGetStatementsResponseStatementDetailsAccountDetailsDepositoryAccountStatementPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload struct {
	// A dollar amount
	CreditLimit float64 `json:"creditLimit" api:"required"`
	// ID for the credit statement period
	CreditPeriodID string `json:"creditPeriodId" api:"required" format:"uuid"`
	DueDate        string `json:"dueDate" api:"required"`
	// A dollar amount
	StartingBalance float64 `json:"startingBalance" api:"required"`
	// A dollar amount
	TotalAutopayPayments float64 `json:"totalAutopayPayments" api:"required"`
	// A dollar amount
	TotalCashbackEarned float64 `json:"totalCashbackEarned" api:"required"`
	// A dollar amount
	TotalOneTimePayments float64 `json:"totalOneTimePayments" api:"required"`
	// A dollar amount
	TotalTransactions float64 `json:"totalTransactions" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreditLimit          respjson.Field
		CreditPeriodID       respjson.Field
		DueDate              respjson.Field
		StartingBalance      respjson.Field
		TotalAutopayPayments respjson.Field
		TotalCashbackEarned  respjson.Field
		TotalOneTimePayments respjson.Field
		TotalTransactions    respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload) RawJSON() string {
	return r.JSON.raw
}
func (r *CreditGetStatementsResponseStatementDetailsAccountDetailsCreditAccountStatementPayload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreditGetStatementsParams struct {
	// Filter statements where the period start date is on or before this date. If the
	// date is in the future, defaults to the current date. Format: YYYY-MM-DD
	End param.Opt[string] `query:"end,omitzero" json:"-"`
	// The ID of the statement to end the page before (exclusive). When provided,
	// results will end just before this ID and work backwards. Use this for reverse
	// pagination or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 100. Defaults to 25
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter statements where the period start date is on or after this date. Format:
	// YYYY-MM-DD
	Start param.Opt[string] `query:"start,omitzero" json:"-"`
	// The ID of the statement to start the page after (exclusive). When provided,
	// results will begin with the statement immediately following this ID. Use this
	// for standard forward pagination to get the next page of results. Cannot be
	// combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'
	//
	// Any of "asc", "desc".
	Order CreditGetStatementsParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CreditGetStatementsParams]'s query parameters as
// `url.Values`.
func (r CreditGetStatementsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'desc'
type CreditGetStatementsParamsOrder string

const (
	CreditGetStatementsParamsOrderAsc  CreditGetStatementsParamsOrder = "asc"
	CreditGetStatementsParamsOrderDesc CreditGetStatementsParamsOrder = "desc"
)
