// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/mercury-go"
	"github.com/stainless-sdks/mercury-go/internal/testutil"
	"github.com/stainless-sdks/mercury-go/option"
)

func TestAccountsReceivableInvoiceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountsReceivable.Invoices.New(context.TODO(), mercury.AccountsReceivableInvoiceNewParams{
		ACHDebitEnabled:      true,
		CcEmails:             []string{"string"},
		CreditCardEnabled:    true,
		CustomerID:           "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		DestinationAccountID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		DueDate:              time.Now(),
		InvoiceDate:          time.Now(),
		LineItems: []mercury.LineItemDataParam{{
			Name:         "name",
			Quantity:     0,
			UnitPrice:    0,
			SalesTaxRate: mercury.Float(0),
		}},
		UseRealAccountNumber:   true,
		InternalNote:           mercury.String("internalNote"),
		InvoiceNumber:          mercury.String("invoiceNumber"),
		PayerMemo:              mercury.String("payerMemo"),
		PoNumber:               mercury.String("poNumber"),
		SendEmailOption:        mercury.AccountsReceivableInvoiceNewParamsSendEmailOptionDontSend,
		ServicePeriodEndDate:   mercury.Time(time.Now()),
		ServicePeriodStartDate: mercury.Time(time.Now()),
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountsReceivableInvoiceGet(t *testing.T) {
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
	_, err := client.AccountsReceivable.Invoices.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountsReceivableInvoiceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountsReceivable.Invoices.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.AccountsReceivableInvoiceUpdateParams{
			ACHDebitEnabled:   true,
			CcEmails:          []string{"string"},
			CreditCardEnabled: true,
			DueDate:           time.Now(),
			InvoiceDate:       time.Now(),
			InvoiceNumber:     "invoiceNumber",
			LineItems: []mercury.LineItemDataParam{{
				Name:         "name",
				Quantity:     0,
				UnitPrice:    0,
				SalesTaxRate: mercury.Float(0),
			}},
			UseRealAccountNumber:   true,
			InternalNote:           mercury.String("internalNote"),
			PayerMemo:              mercury.String("payerMemo"),
			PoNumber:               mercury.String("poNumber"),
			ServicePeriodEndDate:   mercury.Time(time.Now()),
			ServicePeriodStartDate: mercury.Time(time.Now()),
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

func TestAccountsReceivableInvoiceListWithOptionalParams(t *testing.T) {
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
	_, err := client.AccountsReceivable.Invoices.List(context.TODO(), mercury.AccountsReceivableInvoiceListParams{
		EndBefore:  mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Limit:      mercury.Int(1),
		Order:      mercury.AccountsReceivableInvoiceListParamsOrderAsc,
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

func TestAccountsReceivableInvoiceCancel(t *testing.T) {
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
	err := client.AccountsReceivable.Invoices.Cancel(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountsReceivableInvoiceDownloadPdf(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := mercury.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.AccountsReceivable.Invoices.DownloadPdf(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAccountsReceivableInvoiceListAttachments(t *testing.T) {
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
	_, err := client.AccountsReceivable.Invoices.ListAttachments(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
