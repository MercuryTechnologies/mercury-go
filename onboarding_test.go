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

func TestOnboardingSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Onboarding.Submit(context.TODO(), mercury.OnboardingSubmitParams{
		SubmitOnboardingDataParams: mercury.SubmitOnboardingDataParams{
			BeneficialOwners: []mercury.SubmitOnboardingDataParamsBeneficialOwner{{
				Address1:           mercury.String("address1"),
				Address2:           mercury.String("address2"),
				CitizenshipStatus:  "USCitizen",
				City:               mercury.String("city"),
				Country:            mercury.String("country"),
				DateOfBirth:        mercury.Time(time.Now()),
				Email:              mercury.String("email"),
				FirstName:          mercury.String("firstName"),
				IdentificationBlob: mercury.String("identificationBlob"),
				IdentificationType: "Passport",
				IsPep:              "IsPep",
				JobTitle:           "ChiefExecutiveOfficer",
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
			About: mercury.SubmitOnboardingDataParamsAbout{
				CountriesOfOperations: []string{"string"},
				CountryOfOperation:    mercury.String("countryOfOperation"),
				Description:           mercury.String("description"),
				Industry:              mercury.String("industry"),
				LegalBusinessName:     mercury.String("legalBusinessName"),
				Website:               mercury.String("website"),
			},
			ApplicationType: mercury.SubmitOnboardingDataParamsApplicationTypePendingEinApplication,
			BusinessContactDetails: mercury.SubmitOnboardingDataParamsBusinessContactDetails{
				Address1:    mercury.String("address1"),
				Address2:    mercury.String("address2"),
				City:        mercury.String("city"),
				Country:     mercury.String("country"),
				PhoneNumber: mercury.String("phoneNumber"),
				PostalCode:  mercury.String("postalCode"),
				State:       mercury.String("state"),
			},
			BusinessLegalAddress: mercury.SubmitOnboardingDataParamsBusinessLegalAddress{
				Address1:   mercury.String("address1"),
				Address2:   mercury.String("address2"),
				City:       mercury.String("city"),
				Country:    mercury.String("country"),
				PostalCode: mercury.String("postalCode"),
				Region:     mercury.String("region"),
			},
			BusinessPhysicalAddress: mercury.SubmitOnboardingDataParamsBusinessPhysicalAddress{
				Address1:   mercury.String("address1"),
				Address2:   mercury.String("address2"),
				City:       mercury.String("city"),
				Country:    mercury.String("country"),
				PostalCode: mercury.String("postalCode"),
				Region:     mercury.String("region"),
			},
			FormationDetails: mercury.SubmitOnboardingDataParamsFormationDetails{
				FederalEin:                mercury.String("12-3456789"),
				FormationDocumentFileBlob: mercury.String("12-3456789"),
				CompanyOriginCountry:      mercury.String("companyOriginCountry"),
				CompanyStructure:          "CCorp",
				EinDocumentFileBlob:       mercury.String("einDocumentFileBlob"),
				EInDocumentFileBlob:       mercury.String("eINDocumentFileBlob"),
				ForeignBusinessNumber:     mercury.String("foreignBusinessNumber"),
				FormationDocumentType:     "ArticlesOfIncorporation",
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
