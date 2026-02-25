// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/mercury-go"
	"github.com/stainless-sdks/mercury-go/internal/testutil"
	"github.com/stainless-sdks/mercury-go/option"
)

func TestAccountTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.Transactions.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.AccountTransactionListParams{
			CategoryID:      mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			End:             mercury.String("end"),
			Limit:           mercury.Int(1),
			MercuryCategory: mercury.String("mercuryCategory"),
			Offset:          mercury.Int(0),
			Order:           mercury.AccountTransactionListParamsOrderAsc,
			RequestID:       mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Search:          mercury.String("search"),
			Start:           mercury.String("start"),
			Status:          mercury.AccountTransactionListParamsStatusPending,
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

func TestAccountTransactionSendWithOptionalParams(t *testing.T) {
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
	_, err := client.Account.Transactions.Send(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.AccountTransactionSendParams{
			Amount:         0.01,
			IdempotencyKey: "idempotencyKey",
			PaymentMethod:  mercury.AccountTransactionSendParamsPaymentMethodACH,
			RecipientID:    "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			ExternalMemo:   mercury.String("externalMemo"),
			Note:           mercury.String("note"),
			Purpose: mercury.AccountTransactionSendParamsPurpose{
				Simple: mercury.AccountTransactionSendParamsPurposeSimple{
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
