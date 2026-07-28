// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/run-llama/llama-parse-go"
	"github.com/run-llama/llama-parse-go/internal/testutil"
	"github.com/run-llama/llama-parse-go/option"
)

func TestClassifyNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Classify.New(context.TODO(), llamacloud.ClassifyNewParams{
		ClassifyCreateRequest: llamacloud.ClassifyCreateRequestParam{
			Configuration: llamacloud.ClassifyConfigurationParam{
				Rules: []llamacloud.ClassifyConfigurationRuleParam{{
					Description: "contains invoice number, line items, and total amount",
					Type:        "invoice",
				}},
				Mode: llamacloud.ClassifyConfigurationModeFast,
				ParsingConfiguration: llamacloud.ClassifyConfigurationParsingConfigurationParam{
					Lang:        llamacloud.String("en"),
					MaxPages:    llamacloud.Int(10),
					TargetPages: llamacloud.String("1,3,5-7"),
				},
			},
			ConfigurationID:         llamacloud.String("cfg-11111111-2222-3333-4444-555555555555"),
			FileID:                  llamacloud.String("dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
			FileInput:               llamacloud.String("dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
			ParseJobID:              llamacloud.String("pjb-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
			TransactionID:           llamacloud.String("tx-unique-idempotency-key"),
			WebhookConfigurationIDs: []string{"whc-...", "whc-..."},
			WebhookConfigurations: []llamacloud.ClassifyCreateRequestWebhookConfigurationParam{{
				WebhookEvents: []string{"parse.success", "parse.error"},
				WebhookHeaders: map[string]string{
					"Authorization": "Bearer sk-...",
				},
				WebhookOutputFormat:  llamacloud.String("json"),
				WebhookSigningSecret: llamacloud.String("whsec_..."),
				WebhookURL:           llamacloud.String("https://example.com/webhooks/llamacloud"),
			}},
		},
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClassifyListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Classify.List(context.TODO(), llamacloud.ClassifyListParams{
		ConfigurationID:     llamacloud.String("cfg-11111111-2222-3333-4444-555555555555"),
		CreatedAtOnOrAfter:  llamacloud.Time(time.Now()),
		CreatedAtOnOrBefore: llamacloud.Time(time.Now()),
		JobIDs:              []string{"string", "string"},
		OrganizationID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:            llamacloud.Int(1),
		PageToken:           llamacloud.String("page_token"),
		ProjectID:           llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:              llamacloud.ClassifyListParamsStatusCompleted,
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClassifyGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Classify.Get(
		context.TODO(),
		"job_id",
		llamacloud.ClassifyGetParams{
			OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		},
	)
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
