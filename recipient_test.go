// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/MercuryTechnologies/mercury-go"
	"github.com/MercuryTechnologies/mercury-go/internal/testutil"
	"github.com/MercuryTechnologies/mercury-go/option"
)

func TestRecipientNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Recipients.New(context.TODO(), mercury.RecipientNewParams{
		Emails: []string{"string"},
		Name:   "name",
		Address: mercury.AddressDataParam{
			Address1:   "address1",
			City:       "city",
			PostalCode: "postalCode",
			Address2:   mercury.String("address2"),
			State:      mercury.UsStateAl,
		},
		CheckInfo: mercury.CheckInfoRawParam{
			Address: mercury.AddressWithoutNameParam{
				Address1:   "address1",
				City:       "city",
				Country:    "country",
				PostalCode: "postalCode",
				Region:     "region",
				Address2:   mercury.String("address2"),
			},
		},
		ContactEmail: mercury.String("contactEmail"),
		DomesticWireRoutingInfo: mercury.DomesticWireRoutingInfoRawParam{
			AccountNumber: "accountNumber",
			Address: mercury.AddressWithoutNameParam{
				Address1:   "address1",
				City:       "city",
				Country:    "country",
				PostalCode: "postalCode",
				Region:     "region",
				Address2:   mercury.String("address2"),
			},
			RoutingNumber:       "routingNumber",
			DefaultForBenefitOf: mercury.String("defaultForBenefitOf"),
		},
		ElectronicRoutingInfo: mercury.ElectronicRoutingInfoRawParam{
			AccountNumber: "accountNumber",
			Address: mercury.AddressWithoutNameParam{
				Address1:   "address1",
				City:       "city",
				Country:    "country",
				PostalCode: "postalCode",
				Region:     "region",
				Address2:   mercury.String("address2"),
			},
			ElectronicAccountType: mercury.ElectronicAccountTypeBusinessChecking,
			RoutingNumber:         "routingNumber",
		},
		Nickname: mercury.String("nickname"),
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecipientUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Recipients.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.RecipientUpdateParams{
			Address: mercury.AddressDataParam{
				Address1:   "address1",
				City:       "city",
				PostalCode: "postalCode",
				Address2:   mercury.String("address2"),
				State:      mercury.UsStateAl,
			},
			CheckInfo: mercury.CheckInfoRawParam{
				Address: mercury.AddressWithoutNameParam{
					Address1:   "address1",
					City:       "city",
					Country:    "country",
					PostalCode: "postalCode",
					Region:     "region",
					Address2:   mercury.String("address2"),
				},
			},
			ContactEmail: mercury.String("contactEmail"),
			DomesticWireRoutingInfo: mercury.DomesticWireRoutingInfoRawParam{
				AccountNumber: "accountNumber",
				Address: mercury.AddressWithoutNameParam{
					Address1:   "address1",
					City:       "city",
					Country:    "country",
					PostalCode: "postalCode",
					Region:     "region",
					Address2:   mercury.String("address2"),
				},
				RoutingNumber:       "routingNumber",
				DefaultForBenefitOf: mercury.String("defaultForBenefitOf"),
			},
			ElectronicRoutingInfo: mercury.ElectronicRoutingInfoRawParam{
				AccountNumber: "accountNumber",
				Address: mercury.AddressWithoutNameParam{
					Address1:   "address1",
					City:       "city",
					Country:    "country",
					PostalCode: "postalCode",
					Region:     "region",
					Address2:   mercury.String("address2"),
				},
				ElectronicAccountType: mercury.ElectronicAccountTypeBusinessChecking,
				RoutingNumber:         "routingNumber",
			},
			Emails:   []string{"string"},
			Name:     mercury.String("name"),
			Nickname: mercury.String("nickname"),
		},
	)
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecipientListWithOptionalParams(t *testing.T) {
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
	_, err := client.Recipients.List(context.TODO(), mercury.RecipientListParams{
		EndBefore:  mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Limit:      mercury.Int(1),
		Order:      mercury.RecipientListParamsOrderAsc,
		StartAfter: mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecipientAttach(t *testing.T) {
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
	err := client.Recipients.Attach(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.RecipientAttachParams{
			File: io.Reader(bytes.NewBuffer([]byte("Example data"))),
		},
	)
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecipientGet(t *testing.T) {
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
	_, err := client.Recipients.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRecipientListAttachmentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Recipients.ListAttachments(context.TODO(), mercury.RecipientListAttachmentsParams{
		EndBefore:  mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Limit:      mercury.Int(1),
		Order:      mercury.RecipientListAttachmentsParamsOrderAsc,
		StartAfter: mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
