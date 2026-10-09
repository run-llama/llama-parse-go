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

func TestAlphaVerifyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Alpha.Verify.New(context.TODO(), llamacloud.AlphaVerifyNewParams{
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Configuration: llamacloud.AlphaVerifyNewParamsConfiguration{
			TargetPages: llamacloud.String("1,3,5-7"),
			Tier:        "agentic",
		},
		FileID:                  llamacloud.String("dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		FileInput:               llamacloud.String("dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		TransactionID:           llamacloud.String("tx-unique-idempotency-key"),
		WebhookConfigurationIDs: []string{"whc-...", "whc-..."},
		WebhookConfigurations: []llamacloud.AlphaVerifyNewParamsWebhookConfiguration{{
			WebhookEvents: []string{"parse.success", "parse.error"},
			WebhookHeaders: map[string]string{
				"Authorization": "Bearer sk-...",
			},
			WebhookOutputFormat:  llamacloud.String("json"),
			WebhookSigningSecret: llamacloud.String("whsec_..."),
			WebhookURL:           llamacloud.String("https://example.com/webhooks/llamacloud"),
		}},
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAlphaVerifyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Alpha.Verify.List(context.TODO(), llamacloud.AlphaVerifyListParams{
		CreatedAtOnOrAfter:  llamacloud.Time(time.Now()),
		CreatedAtOnOrBefore: llamacloud.Time(time.Now()),
		Expand:              []string{"string"},
		JobIDs:              []string{"string", "string"},
		OrganizationID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:            llamacloud.Int(1),
		PageToken:           llamacloud.String("page_token"),
		ProjectID:           llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:              llamacloud.AlphaVerifyListParamsStatusCancelled,
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAlphaVerifyCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Alpha.Verify.Cancel(
		context.TODO(),
		"job_id",
		llamacloud.AlphaVerifyCancelParams{
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

func TestAlphaVerifyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Alpha.Verify.Get(
		context.TODO(),
		"job_id",
		llamacloud.AlphaVerifyGetParams{
			Expand:         []string{"string"},
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

func TestAlphaVerifyGetDetailsWithOptionalParams(t *testing.T) {
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
	_, err := client.Alpha.Verify.GetDetails(
		context.TODO(),
		"job_id",
		llamacloud.AlphaVerifyGetDetailsParams{
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
