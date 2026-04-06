// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/mercury-go/internal/apiform"
	"github.com/stainless-sdks/mercury-go/internal/apijson"
	"github.com/stainless-sdks/mercury-go/internal/apiquery"
	"github.com/stainless-sdks/mercury-go/internal/requestconfig"
	"github.com/stainless-sdks/mercury-go/option"
	"github.com/stainless-sdks/mercury-go/packages/pagination"
	"github.com/stainless-sdks/mercury-go/packages/param"
	"github.com/stainless-sdks/mercury-go/packages/respjson"
)

// Manage payment recipients
//
// RecipientService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecipientService] method instead.
type RecipientService struct {
	Options []option.RequestOption
}

// NewRecipientService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecipientService(opts ...option.RequestOption) (r RecipientService) {
	r = RecipientService{}
	r.Options = opts
	return
}

// Create a new recipient for making payments
func (r *RecipientService) New(ctx context.Context, body RecipientNewParams, opts ...option.RequestOption) (res *Recipient, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "recipients"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve details of a specific recipient by ID
func (r *RecipientService) Get(ctx context.Context, recipientID string, opts ...option.RequestOption) (res *Recipient, err error) {
	opts = slices.Concat(r.Options, opts)
	if recipientID == "" {
		err = errors.New("missing required recipientId parameter")
		return nil, err
	}
	path := fmt.Sprintf("recipient/%s", recipientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update an existing recipient's information
func (r *RecipientService) Update(ctx context.Context, recipientID string, body RecipientUpdateParams, opts ...option.RequestOption) (res *Recipient, err error) {
	opts = slices.Concat(r.Options, opts)
	if recipientID == "" {
		err = errors.New("missing required recipientId parameter")
		return nil, err
	}
	path := fmt.Sprintf("recipient/%s", recipientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a paginated list of all recipients. Use cursor parameters (start_after,
// end_before) for pagination.
func (r *RecipientService) List(ctx context.Context, query RecipientListParams, opts ...option.RequestOption) (res *pagination.CursorIDRecipients[Recipient], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "recipients"
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

// Retrieve a paginated list of all recipients. Use cursor parameters (start_after,
// end_before) for pagination.
func (r *RecipientService) ListAutoPaging(ctx context.Context, query RecipientListParams, opts ...option.RequestOption) *pagination.CursorIDRecipientsAutoPager[Recipient] {
	return pagination.NewCursorIDRecipientsAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a paginated list of all recipient tax form attachments across all
// recipients in the organization. Use cursor parameters (start_after, end_before)
// for pagination.
func (r *RecipientService) ListAttachments(ctx context.Context, query RecipientListAttachmentsParams, opts ...option.RequestOption) (res *pagination.CursorIDRecipientAttachments[RecipientListAttachmentsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "recipients/attachments"
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

// Retrieve a paginated list of all recipient tax form attachments across all
// recipients in the organization. Use cursor parameters (start_after, end_before)
// for pagination.
func (r *RecipientService) ListAttachmentsAutoPaging(ctx context.Context, query RecipientListAttachmentsParams, opts ...option.RequestOption) *pagination.CursorIDRecipientAttachmentsAutoPager[RecipientListAttachmentsResponse] {
	return pagination.NewCursorIDRecipientAttachmentsAutoPager(r.ListAttachments(ctx, query, opts...))
}

// Upload a tax form attachment for a recipient. The file is uploaded via
// multipart/form-data. Supported file types include PDF, images (PNG, JPG, GIF),
// and common document formats. The attachment will be associated as a tax document
// for the recipient.
func (r *RecipientService) UploadAttachment(ctx context.Context, recipientID string, body RecipientUploadAttachmentParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if recipientID == "" {
		err = errors.New("missing required recipientId parameter")
		return err
	}
	path := fmt.Sprintf("recipient/%s/attachments", recipientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type Address struct {
	Address1   string `json:"address1" api:"required"`
	City       string `json:"city" api:"required"`
	Country    string `json:"country" api:"required"`
	Name       string `json:"name" api:"required"`
	PostalCode string `json:"postalCode" api:"required"`
	Region     string `json:"region" api:"required"`
	Address2   string `json:"address2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address1    respjson.Field
		City        respjson.Field
		Country     respjson.Field
		Name        respjson.Field
		PostalCode  respjson.Field
		Region      respjson.Field
		Address2    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Address) RawJSON() string { return r.JSON.raw }
func (r *Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressData struct {
	Address1   string `json:"address1" api:"required"`
	City       string `json:"city" api:"required"`
	PostalCode string `json:"postalCode" api:"required"`
	Address2   string `json:"address2" api:"nullable"`
	// Any of "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA", "HI",
	// "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS",
	// "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR",
	// "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY".
	State UsState `json:"state" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address1    respjson.Field
		City        respjson.Field
		PostalCode  respjson.Field
		Address2    respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressData) RawJSON() string { return r.JSON.raw }
func (r *AddressData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AddressData to a AddressDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AddressDataParam.Overrides()
func (r AddressData) ToParam() AddressDataParam {
	return param.Override[AddressDataParam](json.RawMessage(r.RawJSON()))
}

// The properties Address1, City, PostalCode are required.
type AddressDataParam struct {
	Address1   string            `json:"address1" api:"required"`
	City       string            `json:"city" api:"required"`
	PostalCode string            `json:"postalCode" api:"required"`
	Address2   param.Opt[string] `json:"address2,omitzero"`
	// Any of "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA", "HI",
	// "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS",
	// "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR",
	// "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY".
	State UsState `json:"state,omitzero"`
	paramObj
}

func (r AddressDataParam) MarshalJSON() (data []byte, err error) {
	type shadow AddressDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AddressWithoutName struct {
	Address1   string `json:"address1" api:"required"`
	City       string `json:"city" api:"required"`
	Country    string `json:"country" api:"required"`
	PostalCode string `json:"postalCode" api:"required"`
	Region     string `json:"region" api:"required"`
	Address2   string `json:"address2" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address1    respjson.Field
		City        respjson.Field
		Country     respjson.Field
		PostalCode  respjson.Field
		Region      respjson.Field
		Address2    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddressWithoutName) RawJSON() string { return r.JSON.raw }
func (r *AddressWithoutName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AddressWithoutName to a AddressWithoutNameParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AddressWithoutNameParam.Overrides()
func (r AddressWithoutName) ToParam() AddressWithoutNameParam {
	return param.Override[AddressWithoutNameParam](json.RawMessage(r.RawJSON()))
}

// The properties Address1, City, Country, PostalCode, Region are required.
type AddressWithoutNameParam struct {
	Address1   string            `json:"address1" api:"required"`
	City       string            `json:"city" api:"required"`
	Country    string            `json:"country" api:"required"`
	PostalCode string            `json:"postalCode" api:"required"`
	Region     string            `json:"region" api:"required"`
	Address2   param.Opt[string] `json:"address2,omitzero"`
	paramObj
}

func (r AddressWithoutNameParam) MarshalJSON() (data []byte, err error) {
	type shadow AddressWithoutNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AddressWithoutNameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CheckInfo struct {
	Address AddressWithoutName `json:"address" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CheckInfo) RawJSON() string { return r.JSON.raw }
func (r *CheckInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Address is required.
type CheckInfoRawParam struct {
	// Mailing address for sending a physical check.
	Address AddressWithoutNameParam `json:"address,omitzero" api:"required"`
	paramObj
}

func (r CheckInfoRawParam) MarshalJSON() (data []byte, err error) {
	type shadow CheckInfoRawParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CheckInfoRawParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomesticWireRoutingInfo struct {
	AccountNumber string             `json:"accountNumber" api:"required"`
	RoutingNumber string             `json:"routingNumber" api:"required"`
	Address       AddressWithoutName `json:"address" api:"nullable"`
	BankName      string             `json:"bankName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountNumber respjson.Field
		RoutingNumber respjson.Field
		Address       respjson.Field
		BankName      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomesticWireRoutingInfo) RawJSON() string { return r.JSON.raw }
func (r *DomesticWireRoutingInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccountNumber, Address, RoutingNumber are required.
type DomesticWireRoutingInfoRawParam struct {
	// The account number of the bank account to use for domestic wire payments.
	AccountNumber string `json:"accountNumber" api:"required"`
	// The address of the bank account to use for domestic wire payments. This has to
	// be the recipient's legal address.
	Address AddressWithoutNameParam `json:"address,omitzero" api:"required"`
	// The routing number of the bank account to use for domestic wire payments.
	RoutingNumber string `json:"routingNumber" api:"required"`
	// The name of the beneficiary of the domestic wire. This is the name of the entity
	// that will receive the domestic wire.
	DefaultForBenefitOf param.Opt[string] `json:"defaultForBenefitOf,omitzero"`
	paramObj
}

func (r DomesticWireRoutingInfoRawParam) MarshalJSON() (data []byte, err error) {
	type shadow DomesticWireRoutingInfoRawParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DomesticWireRoutingInfoRawParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ElectronicAccountType string

const (
	ElectronicAccountTypeBusinessChecking ElectronicAccountType = "businessChecking"
	ElectronicAccountTypeBusinessSavings  ElectronicAccountType = "businessSavings"
	ElectronicAccountTypePersonalChecking ElectronicAccountType = "personalChecking"
	ElectronicAccountTypePersonalSavings  ElectronicAccountType = "personalSavings"
)

type ElectronicRoutingInfo struct {
	AccountNumber string `json:"accountNumber" api:"required"`
	// Any of "businessChecking", "businessSavings", "personalChecking",
	// "personalSavings".
	ElectronicAccountType ElectronicAccountType `json:"electronicAccountType" api:"required"`
	RoutingNumber         string                `json:"routingNumber" api:"required"`
	Address               AddressWithoutName    `json:"address" api:"nullable"`
	BankName              string                `json:"bankName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountNumber         respjson.Field
		ElectronicAccountType respjson.Field
		RoutingNumber         respjson.Field
		Address               respjson.Field
		BankName              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ElectronicRoutingInfo) RawJSON() string { return r.JSON.raw }
func (r *ElectronicRoutingInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccountNumber, Address, ElectronicAccountType, RoutingNumber are
// required.
type ElectronicRoutingInfoRawParam struct {
	// The account number of the bank account to use for ACH payments.
	AccountNumber string `json:"accountNumber" api:"required"`
	// The address of the bank account to use for ACH payments. This has to be the
	// recipient's legal address.
	Address AddressWithoutNameParam `json:"address,omitzero" api:"required"`
	// The type of bank account to use for ACH payments.
	//
	// Any of "businessChecking", "businessSavings", "personalChecking",
	// "personalSavings".
	ElectronicAccountType ElectronicAccountType `json:"electronicAccountType,omitzero" api:"required"`
	// The routing number of the bank account to use for ACH payments.
	RoutingNumber string `json:"routingNumber" api:"required"`
	paramObj
}

func (r ElectronicRoutingInfoRawParam) MarshalJSON() (data []byte, err error) {
	type shadow ElectronicRoutingInfoRawParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ElectronicRoutingInfoRawParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfo struct {
	CountrySpecific   InternationalWireRoutingInfoCountrySpecific   `json:"countrySpecific" api:"required"`
	Iban              string                                        `json:"iban" api:"required"`
	SwiftCode         string                                        `json:"swiftCode" api:"required"`
	Address           AddressWithoutName                            `json:"address" api:"nullable"`
	BankDetails       InternationalWireRoutingInfoBankDetails       `json:"bankDetails" api:"nullable"`
	CorrespondentInfo InternationalWireRoutingInfoCorrespondentInfo `json:"correspondentInfo" api:"nullable"`
	EmailAddress      string                                        `json:"emailAddress" api:"nullable"`
	PhoneNumber       string                                        `json:"phoneNumber" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountrySpecific   respjson.Field
		Iban              respjson.Field
		SwiftCode         respjson.Field
		Address           respjson.Field
		BankDetails       respjson.Field
		CorrespondentInfo respjson.Field
		EmailAddress      respjson.Field
		PhoneNumber       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfo) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecific struct {
	Australia         InternationalWireRoutingInfoCountrySpecificAustralia         `json:"australia" api:"nullable"`
	Brazil            InternationalWireRoutingInfoCountrySpecificBrazil            `json:"brazil" api:"nullable"`
	Canada            InternationalWireRoutingInfoCountrySpecificCanada            `json:"canada" api:"nullable"`
	Chile             InternationalWireRoutingInfoCountrySpecificChile             `json:"chile" api:"nullable"`
	Colombia          InternationalWireRoutingInfoCountrySpecificColombia          `json:"colombia" api:"nullable"`
	DominicanRepublic InternationalWireRoutingInfoCountrySpecificDominicanRepublic `json:"dominicanRepublic" api:"nullable"`
	Honduras          InternationalWireRoutingInfoCountrySpecificHonduras          `json:"honduras" api:"nullable"`
	India             InternationalWireRoutingInfoCountrySpecificIndia             `json:"india" api:"nullable"`
	Kazakhstan        InternationalWireRoutingInfoCountrySpecificKazakhstan        `json:"kazakhstan" api:"nullable"`
	Pakistan          InternationalWireRoutingInfoCountrySpecificPakistan          `json:"pakistan" api:"nullable"`
	Paraguay          InternationalWireRoutingInfoCountrySpecificParaguay          `json:"paraguay" api:"nullable"`
	Philippines       InternationalWireRoutingInfoCountrySpecificPhilippines       `json:"philippines" api:"nullable"`
	Russia            InternationalWireRoutingInfoCountrySpecificRussia            `json:"russia" api:"nullable"`
	SouthAfrica       InternationalWireRoutingInfoCountrySpecificSouthAfrica       `json:"southAfrica" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Australia         respjson.Field
		Brazil            respjson.Field
		Canada            respjson.Field
		Chile             respjson.Field
		Colombia          respjson.Field
		DominicanRepublic respjson.Field
		Honduras          respjson.Field
		India             respjson.Field
		Kazakhstan        respjson.Field
		Pakistan          respjson.Field
		Paraguay          respjson.Field
		Philippines       respjson.Field
		Russia            respjson.Field
		SouthAfrica       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecific) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecific) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificAustralia struct {
	BsbCode string `json:"bsbCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BsbCode     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificAustralia) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificAustralia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificBrazil struct {
	LegalID string `json:"legalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LegalID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificBrazil) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificBrazil) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificCanada struct {
	BankCode      string `json:"bankCode" api:"required"`
	TransitNumber string `json:"transitNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BankCode      respjson.Field
		TransitNumber respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificCanada) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificCanada) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificChile struct {
	LegalID string `json:"legalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LegalID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificChile) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificChile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificColombia struct {
	LegalID string `json:"legalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LegalID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificColombia) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificColombia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificDominicanRepublic struct {
	// Any of "checking", "savings".
	AccountType SwiftBankAccountType `json:"accountType" api:"required"`
	LegalID     string               `json:"legalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountType respjson.Field
		LegalID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificDominicanRepublic) RawJSON() string {
	return r.JSON.raw
}
func (r *InternationalWireRoutingInfoCountrySpecificDominicanRepublic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificHonduras struct {
	// Any of "checking", "savings".
	AccountType SwiftBankAccountType `json:"accountType" api:"required"`
	LegalID     string               `json:"legalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountType respjson.Field
		LegalID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificHonduras) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificHonduras) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificIndia struct {
	IfscCode string `json:"ifscCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IfscCode    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificIndia) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificIndia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificKazakhstan struct {
	LegalID string `json:"legalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LegalID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificKazakhstan) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificKazakhstan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificPakistan struct {
	LegalID string `json:"legalId" api:"required"`
	// Any of "CNIC", "SNIC", "Passport", "NTN".
	LegalIDType string `json:"legalIdType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LegalID     respjson.Field
		LegalIDType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificPakistan) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificPakistan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificParaguay struct {
	LegalID string `json:"legalId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LegalID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificParaguay) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificParaguay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificPhilippines struct {
	RoutingNumber string `json:"routingNumber" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RoutingNumber respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificPhilippines) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificPhilippines) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificRussia struct {
	Inn string `json:"inn" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inn         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificRussia) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificRussia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCountrySpecificSouthAfrica struct {
	BranchCode string `json:"branchCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BranchCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCountrySpecificSouthAfrica) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCountrySpecificSouthAfrica) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoBankDetails struct {
	BankCityState string `json:"bankCityState" api:"required"`
	BankCountry   string `json:"bankCountry" api:"required"`
	BankName      string `json:"bankName" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BankCityState respjson.Field
		BankCountry   respjson.Field
		BankName      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoBankDetails) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoBankDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InternationalWireRoutingInfoCorrespondentInfo struct {
	BankName      string `json:"bankName" api:"nullable"`
	RoutingNumber string `json:"routingNumber" api:"nullable"`
	SwiftCode     string `json:"swiftCode" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BankName      respjson.Field
		RoutingNumber respjson.Field
		SwiftCode     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InternationalWireRoutingInfoCorrespondentInfo) RawJSON() string { return r.JSON.raw }
func (r *InternationalWireRoutingInfoCorrespondentInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RealTimePaymentRoutingInfo struct {
	AccountNumber string             `json:"accountNumber" api:"required"`
	RoutingNumber string             `json:"routingNumber" api:"required"`
	Address       AddressWithoutName `json:"address" api:"nullable"`
	BankName      string             `json:"bankName" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountNumber respjson.Field
		RoutingNumber respjson.Field
		Address       respjson.Field
		BankName      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RealTimePaymentRoutingInfo) RawJSON() string { return r.JSON.raw }
func (r *RealTimePaymentRoutingInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Recipient struct {
	// ID for a Mercury account.
	ID          string                `json:"id" api:"required" format:"uuid"`
	Attachments []RecipientAttachment `json:"attachments" api:"required"`
	// Any of "ach", "check", "domesticWire", "internationalWire", "realTimePayment".
	DefaultPaymentMethod RecipientDefaultPaymentMethod `json:"defaultPaymentMethod" api:"required"`
	Emails               []string                      `json:"emails" api:"required"`
	Name                 string                        `json:"name" api:"required"`
	// Any of "active", "deleted".
	Status                       RecipientStatus              `json:"status" api:"required"`
	Address                      Address                      `json:"address" api:"nullable"`
	CheckInfo                    CheckInfo                    `json:"checkInfo" api:"nullable"`
	ContactEmail                 string                       `json:"contactEmail" api:"nullable"`
	DateLastPaid                 string                       `json:"dateLastPaid" api:"nullable" format:"yyyy-mm-ddThh:MM:ssZ"`
	DefaultAddress               AddressWithoutName           `json:"defaultAddress" api:"nullable"`
	DomesticWireRoutingInfo      DomesticWireRoutingInfo      `json:"domesticWireRoutingInfo" api:"nullable"`
	ElectronicRoutingInfo        ElectronicRoutingInfo        `json:"electronicRoutingInfo" api:"nullable"`
	InternationalWireRoutingInfo InternationalWireRoutingInfo `json:"internationalWireRoutingInfo" api:"nullable"`
	IsBusiness                   bool                         `json:"isBusiness" api:"nullable"`
	Nickname                     string                       `json:"nickname" api:"nullable"`
	RealTimePaymentRoutingInfo   RealTimePaymentRoutingInfo   `json:"realTimePaymentRoutingInfo" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		Attachments                  respjson.Field
		DefaultPaymentMethod         respjson.Field
		Emails                       respjson.Field
		Name                         respjson.Field
		Status                       respjson.Field
		Address                      respjson.Field
		CheckInfo                    respjson.Field
		ContactEmail                 respjson.Field
		DateLastPaid                 respjson.Field
		DefaultAddress               respjson.Field
		DomesticWireRoutingInfo      respjson.Field
		ElectronicRoutingInfo        respjson.Field
		InternationalWireRoutingInfo respjson.Field
		IsBusiness                   respjson.Field
		Nickname                     respjson.Field
		RealTimePaymentRoutingInfo   respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Recipient) RawJSON() string { return r.JSON.raw }
func (r *Recipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientDefaultPaymentMethod string

const (
	RecipientDefaultPaymentMethodACH               RecipientDefaultPaymentMethod = "ach"
	RecipientDefaultPaymentMethodCheck             RecipientDefaultPaymentMethod = "check"
	RecipientDefaultPaymentMethodDomesticWire      RecipientDefaultPaymentMethod = "domesticWire"
	RecipientDefaultPaymentMethodInternationalWire RecipientDefaultPaymentMethod = "internationalWire"
	RecipientDefaultPaymentMethodRealTimePayment   RecipientDefaultPaymentMethod = "realTimePayment"
)

type RecipientStatus string

const (
	RecipientStatusActive  RecipientStatus = "active"
	RecipientStatusDeleted RecipientStatus = "deleted"
)

type RecipientAttachment struct {
	// Name of the uploaded file
	FileName string `json:"fileName" api:"required"`
	// Timestamp when the attachment was uploaded
	UploadedAt string `json:"uploadedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Presigned URL to download the attachment (valid for 12 hours)
	URL string `json:"url" api:"required"`
	// The tax form type (W-9 for US persons, W-8BEN for foreign individuals, W-8BEN-E
	// for foreign entities)
	//
	// Any of "w9", "w8BEN", "w8BENE", "unknown".
	FormType TaxFormType `json:"formType" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileName    respjson.Field
		UploadedAt  respjson.Field
		URL         respjson.Field
		FormType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientAttachment) RawJSON() string { return r.JSON.raw }
func (r *RecipientAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwiftBankAccountType string

const (
	SwiftBankAccountTypeChecking SwiftBankAccountType = "checking"
	SwiftBankAccountTypeSavings  SwiftBankAccountType = "savings"
)

type TaxFormType string

const (
	TaxFormTypeW9      TaxFormType = "w9"
	TaxFormTypeW8Ben   TaxFormType = "w8BEN"
	TaxFormTypeW8Bene  TaxFormType = "w8BENE"
	TaxFormTypeUnknown TaxFormType = "unknown"
)

type RecipientListAttachmentsResponse struct {
	// ID for the recipient tax form attachment
	ID string `json:"id" api:"required" format:"uuid"`
	// Name of the uploaded file
	FileName string `json:"fileName" api:"required"`
	// ID for a Mercury account.
	RecipientID string `json:"recipientId" api:"required" format:"uuid"`
	// Timestamp when the attachment was uploaded
	UploadedAt string `json:"uploadedAt" api:"required" format:"yyyy-mm-ddThh:MM:ssZ"`
	// Presigned URL to download the attachment (valid for 12 hours)
	URL string `json:"url" api:"required"`
	// The tax form type (W-9, W-8BEN, W-8BEN-E, or Unknown)
	//
	// Any of "w9", "w8BEN", "w8BENE", "unknown".
	FormType TaxFormType `json:"formType" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		FileName    respjson.Field
		RecipientID respjson.Field
		UploadedAt  respjson.Field
		URL         respjson.Field
		FormType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecipientListAttachmentsResponse) RawJSON() string { return r.JSON.raw }
func (r *RecipientListAttachmentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientNewParams struct {
	Emails []string `json:"emails,omitzero" api:"required"`
	Name   string   `json:"name" api:"required"`
	// Contact email address of the recipient
	ContactEmail param.Opt[string] `json:"contactEmail,omitzero"`
	Nickname     param.Opt[string] `json:"nickname,omitzero"`
	// Deprecated. Use checkInfo instead.
	Address AddressDataParam `json:"address,omitzero"`
	// Information needed to send a physical check.
	CheckInfo CheckInfoRawParam `json:"checkInfo,omitzero"`
	// Information needed to send a domestic wire.
	DomesticWireRoutingInfo DomesticWireRoutingInfoRawParam `json:"domesticWireRoutingInfo,omitzero"`
	// Information needed to send an ACH.
	ElectronicRoutingInfo ElectronicRoutingInfoRawParam `json:"electronicRoutingInfo,omitzero"`
	paramObj
}

func (r RecipientNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RecipientNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientUpdateParams struct {
	// Contact email address of the recipient
	ContactEmail param.Opt[string] `json:"contactEmail,omitzero"`
	Name         param.Opt[string] `json:"name,omitzero"`
	Nickname     param.Opt[string] `json:"nickname,omitzero"`
	// Deprecated. Use checkInfo instead.
	Address AddressDataParam `json:"address,omitzero"`
	// Information needed to send a check.
	CheckInfo CheckInfoRawParam `json:"checkInfo,omitzero"`
	// Information needed to send a domestic wire.
	DomesticWireRoutingInfo DomesticWireRoutingInfoRawParam `json:"domesticWireRoutingInfo,omitzero"`
	// Information needed to send an ACH.
	ElectronicRoutingInfo ElectronicRoutingInfoRawParam `json:"electronicRoutingInfo,omitzero"`
	Emails                []string                      `json:"emails,omitzero"`
	paramObj
}

func (r RecipientUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RecipientUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecipientUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecipientListParams struct {
	// The ID of the recipient to end the page before (exclusive). When provided,
	// results will end just before this ID and work backwards. Use this for reverse
	// pagination or to retrieve previous pages. Cannot be combined with start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the recipient to start the page after (exclusive). When provided,
	// results will begin with the recipient immediately following this ID. Use this
	// for standard forward pagination to get the next page of results. Cannot be
	// combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order RecipientListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecipientListParams]'s query parameters as `url.Values`.
func (r RecipientListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type RecipientListParamsOrder string

const (
	RecipientListParamsOrderAsc  RecipientListParamsOrder = "asc"
	RecipientListParamsOrderDesc RecipientListParamsOrder = "desc"
)

type RecipientListAttachmentsParams struct {
	// The ID of the recipient attachment to end the page before (exclusive). When
	// provided, results will end just before this ID and work backwards. Use this for
	// reverse pagination or to retrieve previous pages. Cannot be combined with
	// start_after.
	EndBefore param.Opt[string] `query:"end_before,omitzero" format:"uuid" json:"-"`
	// Maximum number of results to return. Allowed range: 1 to 1000. Defaults to 1000
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the recipient attachment to start the page after (exclusive). When
	// provided, results will begin with the recipient attachment immediately following
	// this ID. Use this for standard forward pagination to get the next page of
	// results. Cannot be combined with end_before.
	StartAfter param.Opt[string] `query:"start_after,omitzero" format:"uuid" json:"-"`
	// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
	//
	// Any of "asc", "desc".
	Order RecipientListAttachmentsParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RecipientListAttachmentsParams]'s query parameters as
// `url.Values`.
func (r RecipientListAttachmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order. Can be 'asc' or 'desc'. Defaults to 'asc'
type RecipientListAttachmentsParamsOrder string

const (
	RecipientListAttachmentsParamsOrderAsc  RecipientListAttachmentsParamsOrder = "asc"
	RecipientListAttachmentsParamsOrderDesc RecipientListAttachmentsParamsOrder = "desc"
)

type RecipientUploadAttachmentParams struct {
	// The file to upload (tax form document)
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r RecipientUploadAttachmentParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
