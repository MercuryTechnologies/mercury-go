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

func TestBookNewAgentLedgerTemplateWithOptionalParams(t *testing.T) {
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
	_, err := client.Books.NewAgentLedgerTemplate(context.TODO(), mercury.BookNewAgentLedgerTemplateParams{
		CoaTemplateID:        "7JRNsKwy2Lw66caxVU7WGC",
		DebitCredit:          mercury.BookNewAgentLedgerTemplateParamsDebitCreditDebit,
		Editable:             true,
		FinancialAccountType: mercury.BookNewAgentLedgerTemplateParamsFinancialAccountTypeBankAccount,
		LedgerSubType:        mercury.BookNewAgentLedgerTemplateParamsLedgerSubTypeCurrentAssets,
		LedgerType:           mercury.BookNewAgentLedgerTemplateParamsLedgerTypeAsset,
		Name:                 "name",
		ReportCashFlow:       true,
		SortCode:             "sort_code",
		ParentID:             mercury.String("7JRNsKwy2Lw66caxVU7WGC"),
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBookGetJournalEntryWithOptionalParams(t *testing.T) {
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
	_, err := client.Books.GetJournalEntry(
		context.TODO(),
		"7JRNsKwy2Lw66caxVU7WGC",
		mercury.BookGetJournalEntryParams{
			BooksID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Expand:  mercury.BookGetJournalEntryParamsExpandLineEntries,
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
