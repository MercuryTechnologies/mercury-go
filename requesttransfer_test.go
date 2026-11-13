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

func TestRequestTransferGet(t *testing.T) {
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
	_, err := client.RequestTransfer.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRequestTransferRequestTransferWithOptionalParams(t *testing.T) {
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
	_, err := client.RequestTransfer.RequestTransfer(context.TODO(), mercury.RequestTransferRequestTransferParams{
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

func TestRequestTransferGetRequestTransferWithOptionalParams(t *testing.T) {
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
	_, err := client.RequestTransfer.GetRequestTransfer(context.TODO(), mercury.RequestTransferGetRequestTransferParams{
		DestinationAccountID: mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		EndBefore:            mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Limit:                mercury.Int(1),
		Order:                mercury.RequestTransferGetRequestTransferParamsOrderAsc,
		SourceAccountID:      mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		StartAfter:           mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		StartAt:              mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:               mercury.RequestTransferGetRequestTransferParamsStatusPendingApproval,
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
