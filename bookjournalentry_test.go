// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mercury_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/mercury-go"
	"github.com/stainless-sdks/mercury-go/internal/testutil"
	"github.com/stainless-sdks/mercury-go/option"
)

func TestBookJournalEntryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Books.JournalEntries.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.BookJournalEntryListParams{
			Amount:         mercury.Float(0),
			CreationSource: mercury.BookJournalEntryListParamsCreationSourceManual,
			EndDate:        mercury.Time(time.Now()),
			Expand:         mercury.BookJournalEntryListParamsExpandLineEntries,
			Keywords:       mercury.String("keywords"),
			LedgerID:       mercury.String("7JRNsKwy2Lw66caxVU7WGC"),
			LedgerIDGroups: []string{"7JRNsKwy2Lw66caxVU7WGC"},
			Limit:          mercury.Int(0),
			MaxAmount:      mercury.Float(0),
			MinAmount:      mercury.Float(0),
			PageToken:      mercury.String("page_token"),
			Sort:           mercury.BookJournalEntryListParamsSortDescription,
			StartDate:      mercury.Time(time.Now()),
			TimeZone:       mercury.String("time_zone"),
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

func TestBookJournalEntryBulkDelete(t *testing.T) {
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
	_, err := client.Books.JournalEntries.BulkDelete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.BookJournalEntryBulkDeleteParams{
			JournalEntryIDs: []string{"7JRNsKwy2Lw66caxVU7WGC"},
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

func TestBookJournalEntryBulkUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Books.JournalEntries.BulkUpdate(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.BookJournalEntryBulkUpdateParams{
			Body: []mercury.BookJournalEntryBulkUpdateParamsBody{{
				ID:          "7JRNsKwy2Lw66caxVU7WGC",
				Datetime:    mercury.String("2016-07-22T00:00:00Z"),
				Description: mercury.String("description"),
				LineEntryChanges: mercury.BookJournalEntryBulkUpdateParamsBodyLineEntryChanges{
					Create: []mercury.BookJournalEntryBulkUpdateParamsBodyLineEntryChangesCreate{{
						Amount:      0,
						DebitCredit: "debit",
						Description: "description",
						LedgerID:    "7JRNsKwy2Lw66caxVU7WGC",
						TagIDs:      []string{"7JRNsKwy2Lw66caxVU7WGC"},
					}},
					Delete: []string{"7JRNsKwy2Lw66caxVU7WGC"},
					Update: []mercury.BookJournalEntryBulkUpdateParamsBodyLineEntryChangesUpdate{{
						ID:          "7JRNsKwy2Lw66caxVU7WGC",
						Amount:      mercury.Float(0),
						DebitCredit: "debit",
						Description: mercury.String("description"),
						LedgerID:    mercury.String("7JRNsKwy2Lw66caxVU7WGC"),
						TagIDs:      []string{"7JRNsKwy2Lw66caxVU7WGC"},
					}},
				},
			}},
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

func TestBookJournalEntryNewMultipleWithOptionalParams(t *testing.T) {
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
	_, err := client.Books.JournalEntries.NewMultiple(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		mercury.BookJournalEntryNewMultipleParams{
			Expand:     mercury.BookJournalEntryNewMultipleParamsExpandLineEntries,
			ModifiedBy: mercury.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Body: []mercury.BookJournalEntryNewMultipleParamsBody{{
				Datetime:    "2016-07-22T00:00:00Z",
				Description: "description",
				LineEntries: []mercury.BookJournalEntryNewMultipleParamsBodyLineEntry{{
					Amount:      0,
					DebitCredit: "debit",
					Description: "description",
					LedgerID:    "7JRNsKwy2Lw66caxVU7WGC",
					TagIDs:      []string{"7JRNsKwy2Lw66caxVU7WGC"},
				}},
			}},
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
