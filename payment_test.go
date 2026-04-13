// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/MercuryTechnologies/mercury-go"
	"github.com/MercuryTechnologies/mercury-go/internal/testutil"
	"github.com/MercuryTechnologies/mercury-go/option"
)

func TestPaymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.PaymentNewParams{
			Amount:         0.01,
			IdempotencyKey: "idempotencyKey",
			PaymentMethod:  mercury.PaymentNewParamsPaymentMethodACH,
			RecipientID:    "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ExternalMemo:   mercury.String("externalMemo"),
			Note:           mercury.String("note"),
			Purpose: mercury.PaymentNewParamsPurpose{
				Simple: mercury.PaymentNewParamsPurposeSimple{
					Category:       "Employee",
					AdditionalInfo: mercury.String("additionalInfo"),
				},
			},
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

func TestPaymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.List(context.TODO(), mercury.PaymentListParams{
		AccountID:  mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		EndBefore:  mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Limit:      mercury.Int(1),
		StartAfter: mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:     mercury.PaymentListParamsStatusPendingApproval,
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentGet(t *testing.T) {
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
	_, err := client.Payments.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentRequestWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.Request(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.PaymentRequestParams{
			Amount:         0.01,
			IdempotencyKey: "idempotencyKey",
			PaymentMethod:  mercury.SendMoneyPaymentMethodACH,
			RecipientID:    "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ExternalMemo:   mercury.String("externalMemo"),
			Note:           mercury.String("note"),
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

func TestPaymentTransferWithOptionalParams(t *testing.T) {
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
	_, err := client.Payments.Transfer(context.TODO(), mercury.PaymentTransferParams{
		Amount:               0.01,
		DestinationAccountID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		IdempotencyKey:       "idempotencyKey",
		SourceAccountID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		Note:                 mercury.String("note"),
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
