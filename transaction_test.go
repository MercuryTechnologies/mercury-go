// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/mercury-go"
	"github.com/stainless-sdks/mercury-go/internal/testutil"
	"github.com/stainless-sdks/mercury-go/option"
)

func TestTransactionGet(t *testing.T) {
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
	_, err := client.Transactions.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionUpdate(t *testing.T) {
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
	_, err := client.Transactions.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.TransactionUpdateParams{
			CategoryID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Note:       mercury.String("note"),
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

func TestTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.List(context.TODO(), mercury.TransactionListParams{
		AccountID:       []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		CategoryID:      mercury.String("categoryId"),
		End:             mercury.String("end"),
		EndBefore:       mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Limit:           mercury.Int(1),
		MercuryCategory: mercury.String("mercuryCategory"),
		Order:           mercury.TransactionListParamsOrderAsc,
		PostedEnd:       mercury.String("postedEnd"),
		PostedStart:     mercury.String("postedStart"),
		Search:          mercury.String("search"),
		Start:           mercury.String("start"),
		StartAfter:      mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		StartAt:         mercury.String("start_at"),
		Status:          []string{"pending"},
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionUploadAttachmentWithOptionalParams(t *testing.T) {
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
	err := client.Transactions.UploadAttachment(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.TransactionUploadAttachmentParams{
			File:           io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			AttachmentType: mercury.TransactionUploadAttachmentParamsAttachmentTypeReceipt,
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
