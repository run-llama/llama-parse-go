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

func TestJobDataPointListWithOptionalParams(t *testing.T) {
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
	_, err := client.JobDataPoints.List(context.TODO(), llamacloud.JobDataPointListParams{
		JobType:             llamacloud.JobDataPointListParamsJobTypeParse,
		CreatedAtOnOrAfter:  llamacloud.Time(time.Now()),
		CreatedAtOnOrBefore: llamacloud.Time(time.Now()),
		Hours:               llamacloud.Int(24),
		OrganizationID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:            llamacloud.Int(100),
		PageToken:           llamacloud.String("page_token"),
		ProjectID:           llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:              []string{"completed", "failed"},
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
