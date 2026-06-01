// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/MercuryTechnologies/mercury-go"
	"github.com/MercuryTechnologies/mercury-go/internal/testutil"
	"github.com/MercuryTechnologies/mercury-go/option"
)

func TestOnboardingApplyWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := mercury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Onboarding.Apply(context.TODO(), mercury.OnboardingApplyParams{
		SubmitOnboardingDataParams: mercury.SubmitOnboardingDataParams{
			BeneficialOwners: []mercury.BeneficialOwnerParam{{
				Address1:           mercury.String("address1"),
				Address2:           mercury.String("address2"),
				CitizenshipStatus:  mercury.BeneficialOwnerCitizenshipStatusUsCitizen,
				City:               mercury.String("city"),
				Country:            mercury.String("country"),
				DateOfBirth:        mercury.Time(time.Now()),
				Email:              mercury.String("email"),
				FirstName:          mercury.String("firstName"),
				IdentificationBlob: mercury.String("identificationBlob"),
				IdentificationType: mercury.BeneficialOwnerIdentificationTypePassport,
				IsPep:              mercury.BeneficialOwnerIsPepIsPep,
				JobTitle:           mercury.BeneficialOwnerJobTitleChiefExecutiveOfficer,
				LastName:           mercury.String("lastName"),
				OtherJobTitle:      mercury.String("otherJobTitle"),
				PercentOwnership:   mercury.Float(0),
				PhoneNumber:        mercury.String("phoneNumber"),
				PostalCode:         mercury.String("postalCode"),
				Region:             mercury.String("region"),
				SocialProfileLinks: []string{"string"},
				State:              mercury.String("state"),
			}},
			Partner: "partner",
			About: mercury.OnboardingDataAboutParam{
				CountriesOfOperations: []string{"string"},
				CountryOfOperation:    mercury.String("countryOfOperation"),
				Description:           mercury.String("description"),
				Industry:              mercury.String("industry"),
				LegalBusinessName:     mercury.String("legalBusinessName"),
				Website:               mercury.String("website"),
			},
			ApplicationType: mercury.ApplicationTypePendingEinApplication,
			BusinessContactDetails: mercury.BusinessContactDetailsParam{
				Address1:    mercury.String("address1"),
				Address2:    mercury.String("address2"),
				City:        mercury.String("city"),
				Country:     mercury.String("country"),
				PhoneNumber: mercury.String("phoneNumber"),
				PostalCode:  mercury.String("postalCode"),
				State:       mercury.String("state"),
			},
			BusinessLegalAddress: mercury.BusinessAddressParam{
				Address1:   mercury.String("address1"),
				Address2:   mercury.String("address2"),
				City:       mercury.String("city"),
				Country:    mercury.String("country"),
				PostalCode: mercury.String("postalCode"),
				Region:     mercury.String("region"),
			},
			BusinessPhysicalAddress: mercury.BusinessAddressParam{
				Address1:   mercury.String("address1"),
				Address2:   mercury.String("address2"),
				City:       mercury.String("city"),
				Country:    mercury.String("country"),
				PostalCode: mercury.String("postalCode"),
				Region:     mercury.String("region"),
			},
			FormationDetails: mercury.FormationDetailsParam{
				FederalEin:                mercury.String("12-3456789"),
				FormationDocumentFileBlob: mercury.String("12-3456789"),
				CompanyOriginCountry:      mercury.String("companyOriginCountry"),
				CompanyStructure:          mercury.FormationDetailsCompanyStructureCCorp,
				EinDocumentFileBlob:       mercury.String("einDocumentFileBlob"),
				EInDocumentFileBlob:       mercury.String("eINDocumentFileBlob"),
				ForeignBusinessNumber:     mercury.String("foreignBusinessNumber"),
				FormationDocumentType:     mercury.FormationDetailsFormationDocumentTypeArticlesOfIncorporation,
			},
			InviteEmail: mercury.String("inviteEmail"),
			WebhookURL:  mercury.String("webhookURL"),
		},
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
