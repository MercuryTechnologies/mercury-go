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

func TestSpendReimbursementNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Spend.Reimbursements.New(context.TODO(), mercury.SpendReimbursementNewParams{
		APIV1CreateReimbursementRequest: mercury.APIV1CreateReimbursementRequestParam{
			Amount:       0,
			Currency:     "currency",
			DatePaidFor:  time.Now(),
			MerchantName: "x",
			CategoryID:   mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Notes:        mercury.String("x"),
		},
		IdempotencyKey: mercury.String("Idempotency-Key"),
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSpendReimbursementGet(t *testing.T) {
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
	_, err := client.Spend.Reimbursements.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSpendReimbursementUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Spend.Reimbursements.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.SpendReimbursementUpdateParams{
			APIV1CreateReimbursementRequest: mercury.APIV1CreateReimbursementRequestParam{
				Amount:       0,
				Currency:     "currency",
				DatePaidFor:  time.Now(),
				MerchantName: "x",
				CategoryID:   mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Notes:        mercury.String("x"),
			},
			IdempotencyKey: mercury.String("Idempotency-Key"),
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

func TestSpendReimbursementListWithOptionalParams(t *testing.T) {
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
	_, err := client.Spend.Reimbursements.List(context.TODO(), mercury.SpendReimbursementListParams{
		EndBefore:  mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Limit:      mercury.Int(1),
		Order:      mercury.SpendReimbursementListParamsOrderAsc,
		StartAfter: mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		StartAt:    mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSpendReimbursementApproveWithOptionalParams(t *testing.T) {
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
	err := client.Spend.Reimbursements.Approve(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.SpendReimbursementApproveParams{
			IdempotencyKey: mercury.String("Idempotency-Key"),
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

func TestSpendReimbursementCancelWithOptionalParams(t *testing.T) {
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
	err := client.Spend.Reimbursements.Cancel(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.SpendReimbursementCancelParams{
			IdempotencyKey: mercury.String("Idempotency-Key"),
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

func TestSpendReimbursementDenyWithOptionalParams(t *testing.T) {
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
	err := client.Spend.Reimbursements.Deny(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.SpendReimbursementDenyParams{
			Reason:         "x",
			IdempotencyKey: mercury.String("Idempotency-Key"),
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

func TestSpendReimbursementMarkPaidWithOptionalParams(t *testing.T) {
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
	err := client.Spend.Reimbursements.MarkPaid(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.SpendReimbursementMarkPaidParams{
			Note:           mercury.String("x"),
			IdempotencyKey: mercury.String("Idempotency-Key"),
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
