// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/MercuryTechnologies/mercury-go/internal/apijson"
	shimjson "github.com/MercuryTechnologies/mercury-go/internal/encoding/json"
	"github.com/MercuryTechnologies/mercury-go/internal/requestconfig"
	"github.com/MercuryTechnologies/mercury-go/option"
	"github.com/MercuryTechnologies/mercury-go/packages/param"
	"github.com/MercuryTechnologies/mercury-go/packages/respjson"
)

// Endpoints for partner onboarding integration
//
// OnboardingService contains methods and other services that help with interacting
// with the mercury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOnboardingService] method instead.
type OnboardingService struct {
	options []option.RequestOption
}

// NewOnboardingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOnboardingService(opts ...option.RequestOption) (r OnboardingService) {
	r = OnboardingService{}
	r.options = opts
	return
}

// Submit onboarding data for applicants to pre-fill their Mercury application
func (r *OnboardingService) Submit(ctx context.Context, body OnboardingSubmitParams, opts ...option.RequestOption) (res *SubmitOnboardingDataResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "submit-onboarding-data"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The properties BeneficialOwners, Partner are required.
type SubmitOnboardingDataParams struct {
	BeneficialOwners []SubmitOnboardingDataParamsBeneficialOwner `json:"beneficialOwners,omitzero" api:"required"`
	Partner          string                                      `json:"partner" api:"required"`
	InviteEmail      param.Opt[string]                           `json:"inviteEmail,omitzero"`
	WebhookURL       param.Opt[string]                           `json:"webhookURL,omitzero"`
	About            SubmitOnboardingDataParamsAbout             `json:"about,omitzero"`
	// Any of "PendingEINApplication", "DefaultApplication".
	ApplicationType         SubmitOnboardingDataParamsApplicationType         `json:"applicationType,omitzero"`
	BusinessContactDetails  SubmitOnboardingDataParamsBusinessContactDetails  `json:"businessContactDetails,omitzero"`
	BusinessLegalAddress    SubmitOnboardingDataParamsBusinessLegalAddress    `json:"businessLegalAddress,omitzero"`
	BusinessPhysicalAddress SubmitOnboardingDataParamsBusinessPhysicalAddress `json:"businessPhysicalAddress,omitzero"`
	FormationDetails        SubmitOnboardingDataParamsFormationDetails        `json:"formationDetails,omitzero"`
	paramObj
}

func (r SubmitOnboardingDataParams) MarshalJSON() (data []byte, err error) {
	type shadow SubmitOnboardingDataParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubmitOnboardingDataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Beneficial Owner's Information Gathered From The Onboarding API
type SubmitOnboardingDataParamsBeneficialOwner struct {
	// Address line 1 of Beneficial Owner's address
	Address1 param.Opt[string] `json:"address1,omitzero"`
	// Address line 2 of Beneficial Owner's address
	Address2 param.Opt[string] `json:"address2,omitzero"`
	// City of Beneficial Owner's address
	City param.Opt[string] `json:"city,omitzero"`
	// Country of Beneficial Owner's address
	Country param.Opt[string] `json:"country,omitzero"`
	// Beneficial Owner's Date of Birth
	DateOfBirth param.Opt[time.Time] `json:"dateOfBirth,omitzero" format:"date"`
	// Beneficial Owner's Email Address
	Email param.Opt[string] `json:"email,omitzero"`
	// Beneficial Owner's First Name
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// Beneficial Owner's Identification File
	IdentificationBlob param.Opt[string] `json:"identificationBlob,omitzero"`
	// Beneficial Owner's Last Name
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// Beneficial Owner's Alternate Job Title
	OtherJobTitle param.Opt[string] `json:"otherJobTitle,omitzero"`
	// Beneficial Owner's Ownership Percentage
	PercentOwnership param.Opt[float64] `json:"percentOwnership,omitzero"`
	// Beneficial Owner's Phone Number
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// Postal Code of Beneficial Owner's address
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// Region or State of Beneficial Owner's address
	Region param.Opt[string] `json:"region,omitzero"`
	// State or Region of Beneficial Owner's address (Deprecated)
	State param.Opt[string] `json:"state,omitzero"`
	// Beneficial Owner's Citizenship Status
	//
	// Any of "USCitizen", "USResident", "NonResident".
	CitizenshipStatus string `json:"citizenshipStatus,omitzero"`
	// Beneficial Owner's Identification File Type
	//
	// Any of "Passport", "DriversLicense", "StateID", "AlienRegistrationCard",
	// "EmployeeAuthorizationDocument", "VerifiedByThirdParty".
	IdentificationType string `json:"identificationType,omitzero"`
	// Beneficial Owner's pep status
	//
	// Any of "IsPep", "IsNotPep".
	IsPep string `json:"isPep,omitzero"`
	// Beneficial Owner's Job Title
	//
	// Any of "ChiefExecutiveOfficer", "ChiefOperatingOfficer",
	// "ChiefTechnologyOfficer", "ChiefFinancialOfficer", "Founder", "President",
	// "GeneralPartner", "FinanceTeam", "Other".
	JobTitle string `json:"jobTitle,omitzero"`
	// Beneficial Owner's Social Profile Websites
	SocialProfileLinks []string `json:"socialProfileLinks,omitzero"`
	paramObj
}

func (r SubmitOnboardingDataParamsBeneficialOwner) MarshalJSON() (data []byte, err error) {
	type shadow SubmitOnboardingDataParamsBeneficialOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubmitOnboardingDataParamsBeneficialOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SubmitOnboardingDataParamsBeneficialOwner](
		"citizenshipStatus", "USCitizen", "USResident", "NonResident",
	)
	apijson.RegisterFieldValidator[SubmitOnboardingDataParamsBeneficialOwner](
		"identificationType", "Passport", "DriversLicense", "StateID", "AlienRegistrationCard", "EmployeeAuthorizationDocument", "VerifiedByThirdParty",
	)
	apijson.RegisterFieldValidator[SubmitOnboardingDataParamsBeneficialOwner](
		"isPep", "IsPep", "IsNotPep",
	)
	apijson.RegisterFieldValidator[SubmitOnboardingDataParamsBeneficialOwner](
		"jobTitle", "ChiefExecutiveOfficer", "ChiefOperatingOfficer", "ChiefTechnologyOfficer", "ChiefFinancialOfficer", "Founder", "President", "GeneralPartner", "FinanceTeam", "Other",
	)
}

type SubmitOnboardingDataParamsAbout struct {
	CountryOfOperation param.Opt[string] `json:"countryOfOperation,omitzero"`
	Description        param.Opt[string] `json:"description,omitzero"`
	Industry           param.Opt[string] `json:"industry,omitzero"`
	LegalBusinessName  param.Opt[string] `json:"legalBusinessName,omitzero"`
	Website            param.Opt[string] `json:"website,omitzero"`
	// The countries where the company operates.
	CountriesOfOperations []string `json:"countriesOfOperations,omitzero"`
	paramObj
}

func (r SubmitOnboardingDataParamsAbout) MarshalJSON() (data []byte, err error) {
	type shadow SubmitOnboardingDataParamsAbout
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubmitOnboardingDataParamsAbout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubmitOnboardingDataParamsApplicationType string

const (
	SubmitOnboardingDataParamsApplicationTypePendingEinApplication SubmitOnboardingDataParamsApplicationType = "PendingEINApplication"
	SubmitOnboardingDataParamsApplicationTypeDefaultApplication    SubmitOnboardingDataParamsApplicationType = "DefaultApplication"
)

type SubmitOnboardingDataParamsBusinessContactDetails struct {
	Address1    param.Opt[string] `json:"address1,omitzero"`
	Address2    param.Opt[string] `json:"address2,omitzero"`
	City        param.Opt[string] `json:"city,omitzero"`
	Country     param.Opt[string] `json:"country,omitzero"`
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	PostalCode  param.Opt[string] `json:"postalCode,omitzero"`
	State       param.Opt[string] `json:"state,omitzero"`
	paramObj
}

func (r SubmitOnboardingDataParamsBusinessContactDetails) MarshalJSON() (data []byte, err error) {
	type shadow SubmitOnboardingDataParamsBusinessContactDetails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubmitOnboardingDataParamsBusinessContactDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubmitOnboardingDataParamsBusinessLegalAddress struct {
	Address1   param.Opt[string] `json:"address1,omitzero"`
	Address2   param.Opt[string] `json:"address2,omitzero"`
	City       param.Opt[string] `json:"city,omitzero"`
	Country    param.Opt[string] `json:"country,omitzero"`
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	Region     param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r SubmitOnboardingDataParamsBusinessLegalAddress) MarshalJSON() (data []byte, err error) {
	type shadow SubmitOnboardingDataParamsBusinessLegalAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubmitOnboardingDataParamsBusinessLegalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubmitOnboardingDataParamsBusinessPhysicalAddress struct {
	Address1   param.Opt[string] `json:"address1,omitzero"`
	Address2   param.Opt[string] `json:"address2,omitzero"`
	City       param.Opt[string] `json:"city,omitzero"`
	Country    param.Opt[string] `json:"country,omitzero"`
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	Region     param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r SubmitOnboardingDataParamsBusinessPhysicalAddress) MarshalJSON() (data []byte, err error) {
	type shadow SubmitOnboardingDataParamsBusinessPhysicalAddress
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubmitOnboardingDataParamsBusinessPhysicalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FederalEin, FormationDocumentFileBlob are required.
type SubmitOnboardingDataParamsFormationDetails struct {
	// Field should be null (no value), 'Pending' (value will be provided at a later
	// date), or a valid value
	FederalEin param.Opt[string] `json:"federalEin,omitzero" api:"required"`
	// Field should be null (no value), 'Pending' (value will be provided at a later
	// date), or a valid value
	FormationDocumentFileBlob param.Opt[string] `json:"formationDocumentFileBlob,omitzero" api:"required"`
	CompanyOriginCountry      param.Opt[string] `json:"companyOriginCountry,omitzero"`
	EinDocumentFileBlob       param.Opt[string] `json:"einDocumentFileBlob,omitzero"`
	EInDocumentFileBlob       param.Opt[string] `json:"eINDocumentFileBlob,omitzero"`
	ForeignBusinessNumber     param.Opt[string] `json:"foreignBusinessNumber,omitzero"`
	// Any of "CCorp", "LLC", "LLP", "NonProfit", "Partnership",
	// "ProfessionalAssociation", "ProfessionalCorporation", "SCorp",
	// "GeneralPartnership", "LimitedPartnership", "JointVenture",
	// "LLCTaxedAsSoleProprietorship", "SoleProprietorship", "ExemptedCompany",
	// "Limited".
	CompanyStructure string `json:"companyStructure,omitzero"`
	// Any of "ArticlesOfIncorporation", "ArticlesOfOrganization",
	// "CertificateOfFormation", "PartnershipAgreement",
	// "SecretaryOfStateRegistrationPage".
	FormationDocumentType string `json:"formationDocumentType,omitzero"`
	paramObj
}

func (r SubmitOnboardingDataParamsFormationDetails) MarshalJSON() (data []byte, err error) {
	type shadow SubmitOnboardingDataParamsFormationDetails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubmitOnboardingDataParamsFormationDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SubmitOnboardingDataParamsFormationDetails](
		"companyStructure", "CCorp", "LLC", "LLP", "NonProfit", "Partnership", "ProfessionalAssociation", "ProfessionalCorporation", "SCorp", "GeneralPartnership", "LimitedPartnership", "JointVenture", "LLCTaxedAsSoleProprietorship", "SoleProprietorship", "ExemptedCompany", "Limited",
	)
	apijson.RegisterFieldValidator[SubmitOnboardingDataParamsFormationDetails](
		"formationDocumentType", "ArticlesOfIncorporation", "ArticlesOfOrganization", "CertificateOfFormation", "PartnershipAgreement", "SecretaryOfStateRegistrationPage",
	)
}

type SubmitOnboardingDataResponse struct {
	// ID for this saved onboarding data
	OnboardingDataID string `json:"onboardingDataId" api:"required" format:"uuid"`
	SignupLink       string `json:"signupLink" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OnboardingDataID respjson.Field
		SignupLink       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubmitOnboardingDataResponse) RawJSON() string { return r.JSON.raw }
func (r *SubmitOnboardingDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OnboardingSubmitParams struct {
	SubmitOnboardingDataParams SubmitOnboardingDataParams
	paramObj
}

func (r OnboardingSubmitParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SubmitOnboardingDataParams)
}
func (r *OnboardingSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
