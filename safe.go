// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/mercury-go/internal/apijson"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// SafeService contains methods and other services that help with interacting with
// the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSafeService] method instead.
type SafeService struct {
	Options []option.RequestOption
}

// NewSafeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSafeService(opts ...option.RequestOption) (r SafeService) {
	r = SafeService{}
	r.Options = opts
	return
}

// Retrieve a specific SAFE request by its ID.
func (r *SafeService) Get(ctx context.Context, safeRequestID string, opts ...option.RequestOption) (res *SafeRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	if safeRequestID == "" {
		err = errors.New("missing required safeRequestId parameter")
		return
	}
	path := fmt.Sprintf("safes/%s", safeRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve all SAFE (Simple Agreement for Future Equity) requests for your
// organization.
func (r *SafeService) List(ctx context.Context, opts ...option.RequestOption) (res *[]SafeRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/json;charset=utf-8")}, opts...)
	path := "safes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Download the PDF document for a specific SAFE request. Returns binary PDF data
// with a Content-Disposition header.
func (r *SafeService) DownloadDocument(ctx context.Context, safeRequestID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if safeRequestID == "" {
		err = errors.New("missing required safeRequestId parameter")
		return
	}
	path := fmt.Sprintf("safes/%s/document", safeRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A summary of a SAFE request.
type SafeRequest struct {
	// ID for the SAFE request
	ID                              string `json:"id" api:"required" format:"uuid"`
	DocumentURL                     string `json:"documentUrl" api:"required"`
	ExpiresAt                       string `json:"expiresAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	IncludesMostFavoredNationClause bool   `json:"includesMostFavoredNationClause" api:"required"`
	IncludesProRataRightsLetter     bool   `json:"includesProRataRightsLetter" api:"required"`
	// A positive dollar amount with at least 1 cent.
	InvestmentAmount float64   `json:"investmentAmount" api:"required"`
	InvestmentDate   time.Time `json:"investmentDate" api:"required" format:"date"`
	// Details about the investor buying the equity.
	Investor SafeRequestInvestor `json:"investor" api:"required"`
	// Details about the organization selling the equity
	Organization SafeRequestOrganization `json:"organization" api:"required"`
	// Any of "PreMoney", "PostMoney", "NoValuation".
	ValuationType SafeRequestValuationType `json:"valuationType" api:"required"`
	CanceledAt    string                   `json:"canceledAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	DiscountRate  float64                  `json:"discountRate" api:"nullable"`
	// Any of "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA", "HI",
	// "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS",
	// "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR",
	// "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY".
	GoverningState     UsState `json:"governingState" api:"nullable"`
	PaidAt             string  `json:"paidAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	SignedByInvestorAt string  `json:"signedByInvestorAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	SignedByOwnerAt    string  `json:"signedByOwnerAt" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	// A positive dollar amount with at least 1 cent.
	ValuationCap float64 `json:"valuationCap" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                              respjson.Field
		DocumentURL                     respjson.Field
		ExpiresAt                       respjson.Field
		IncludesMostFavoredNationClause respjson.Field
		IncludesProRataRightsLetter     respjson.Field
		InvestmentAmount                respjson.Field
		InvestmentDate                  respjson.Field
		Investor                        respjson.Field
		Organization                    respjson.Field
		ValuationType                   respjson.Field
		CanceledAt                      respjson.Field
		DiscountRate                    respjson.Field
		GoverningState                  respjson.Field
		PaidAt                          respjson.Field
		SignedByInvestorAt              respjson.Field
		SignedByOwnerAt                 respjson.Field
		ValuationCap                    respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SafeRequest) RawJSON() string { return r.JSON.raw }
func (r *SafeRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about the investor buying the equity.
type SafeRequestInvestor struct {
	// Any of "SafeRequestInvestorTypeIndividual",
	// "SafeRequestInvestorTypeVentureFund", "SafeRequestInvestorTypeOther".
	InvestorType      string `json:"investorType" api:"required"`
	LegalEntityName   string `json:"legalEntityName" api:"required"`
	SignatoryEmail    string `json:"signatoryEmail" api:"required"`
	SignatoryName     string `json:"signatoryName" api:"required"`
	AdditionalBylines string `json:"additionalBylines" api:"nullable"`
	Address           string `json:"address" api:"nullable"`
	SignatoryTitle    string `json:"signatoryTitle" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InvestorType      respjson.Field
		LegalEntityName   respjson.Field
		SignatoryEmail    respjson.Field
		SignatoryName     respjson.Field
		AdditionalBylines respjson.Field
		Address           respjson.Field
		SignatoryTitle    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SafeRequestInvestor) RawJSON() string { return r.JSON.raw }
func (r *SafeRequestInvestor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about the organization selling the equity
type SafeRequestOrganization struct {
	LegalEntityName string `json:"legalEntityName" api:"required"`
	SignatoryEmail  string `json:"signatoryEmail" api:"required"`
	SignatoryName   string `json:"signatoryName" api:"required"`
	SignatoryTitle  string `json:"signatoryTitle" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LegalEntityName respjson.Field
		SignatoryEmail  respjson.Field
		SignatoryName   respjson.Field
		SignatoryTitle  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SafeRequestOrganization) RawJSON() string { return r.JSON.raw }
func (r *SafeRequestOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SafeRequestValuationType string

const (
	SafeRequestValuationTypePreMoney    SafeRequestValuationType = "PreMoney"
	SafeRequestValuationTypePostMoney   SafeRequestValuationType = "PostMoney"
	SafeRequestValuationTypeNoValuation SafeRequestValuationType = "NoValuation"
)

type UsState string

const (
	UsStateAl UsState = "AL"
	UsStateAk UsState = "AK"
	UsStateAz UsState = "AZ"
	UsStateAr UsState = "AR"
	UsStateCa UsState = "CA"
	UsStateCo UsState = "CO"
	UsStateCt UsState = "CT"
	UsStateDe UsState = "DE"
	UsStateDc UsState = "DC"
	UsStateFl UsState = "FL"
	UsStateGa UsState = "GA"
	UsStateHi UsState = "HI"
	UsStateID UsState = "ID"
	UsStateIl UsState = "IL"
	UsStateIn UsState = "IN"
	UsStateIa UsState = "IA"
	UsStateKs UsState = "KS"
	UsStateKy UsState = "KY"
	UsStateLa UsState = "LA"
	UsStateMe UsState = "ME"
	UsStateMd UsState = "MD"
	UsStateMa UsState = "MA"
	UsStateMi UsState = "MI"
	UsStateMn UsState = "MN"
	UsStateMs UsState = "MS"
	UsStateMo UsState = "MO"
	UsStateMt UsState = "MT"
	UsStateNe UsState = "NE"
	UsStateNv UsState = "NV"
	UsStateNh UsState = "NH"
	UsStateNj UsState = "NJ"
	UsStateNm UsState = "NM"
	UsStateNy UsState = "NY"
	UsStateNc UsState = "NC"
	UsStateNd UsState = "ND"
	UsStateOh UsState = "OH"
	UsStateOk UsState = "OK"
	UsStateOr UsState = "OR"
	UsStatePa UsState = "PA"
	UsStateRi UsState = "RI"
	UsStateSc UsState = "SC"
	UsStateSd UsState = "SD"
	UsStateTn UsState = "TN"
	UsStateTx UsState = "TX"
	UsStateUt UsState = "UT"
	UsStateVt UsState = "VT"
	UsStateVa UsState = "VA"
	UsStateWa UsState = "WA"
	UsStateWv UsState = "WV"
	UsStateWi UsState = "WI"
	UsStateWy UsState = "WY"
)
