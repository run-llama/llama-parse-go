// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/llamacloud-prod-go"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/testutil"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
)

func TestClassifierJobNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Classifier.Jobs.New(context.TODO(), llamacloudprod.ClassifierJobNewParams{
		FileIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		Rules: []llamacloudprod.ClassifierRuleParam{{
			Description: "contains invoice number, line items, and total amount",
			Type:        "invoice",
		}},
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Mode:           llamacloudprod.ClassifierJobNewParamsModeFast,
		ParsingConfiguration: llamacloudprod.ClassifyParsingConfigurationParam{
			Lang:        llamacloudprod.ParsingLanguagesAf,
			MaxPages:    llamacloudprod.Int(0),
			TargetPages: []int64{0},
		},
		WebhookConfigurations: []llamacloudprod.ClassifierJobNewParamsWebhookConfiguration{{
			WebhookEvents: []string{"string"},
			WebhookHeaders: map[string]any{
				"foo": "bar",
			},
			WebhookOutputFormat: "json",
			WebhookURL:          llamacloudprod.String("https:"),
		}},
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClassifierJobListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Classifier.Jobs.List(context.TODO(), llamacloudprod.ClassifierJobListParams{
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:       llamacloudprod.Int(0),
		PageToken:      llamacloudprod.String("page_token"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClassifierJobGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Classifier.Jobs.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.ClassifierJobGetParams{
			OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClassifierJobGetResultsWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Classifier.Jobs.GetResults(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.ClassifierJobGetResultsParams{
			OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
