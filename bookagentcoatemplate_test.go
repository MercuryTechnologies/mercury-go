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

func TestBookAgentCoaTemplateNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Books.AgentCoaTemplates.New(context.TODO(), mercury.BookAgentCoaTemplateNewParams{
		AccountingBasis: mercury.BookAgentCoaTemplateNewParamsAccountingBasisCash,
		Name:            "name",
		Expand:          mercury.BookAgentCoaTemplateNewParamsExpandLedgers,
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBookAgentCoaTemplateGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Books.AgentCoaTemplates.Get(
		context.TODO(),
		"7JRNsKwy2Lw66caxVU7WGC",
		mercury.BookAgentCoaTemplateGetParams{
			Expand: mercury.BookAgentCoaTemplateGetParamsExpandLedgers,
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

func TestBookAgentCoaTemplateListWithOptionalParams(t *testing.T) {
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
	_, err := client.Books.AgentCoaTemplates.List(context.TODO(), mercury.BookAgentCoaTemplateListParams{
		Expand:    mercury.BookAgentCoaTemplateListParamsExpandLedgers,
		Limit:     mercury.Int(-9007199254740991),
		PageToken: mercury.String("page_token"),
	})
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBookAgentCoaTemplateDelete(t *testing.T) {
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
	_, err := client.Books.AgentCoaTemplates.Delete(context.TODO(), "7JRNsKwy2Lw66caxVU7WGC")
	if err != nil {
		var apierr *mercury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
