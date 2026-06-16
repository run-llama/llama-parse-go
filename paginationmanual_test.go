// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod_test

import (
	"context"
	"os"
	"testing"

	"github.com/run-llama/llama-parse-go"
	"github.com/run-llama/llama-parse-go/internal/testutil"
	"github.com/run-llama/llama-parse-go/option"
)

func TestManualPagination(t *testing.T) {
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
	page, err := client.Extract.List(context.TODO(), llamacloudprod.ExtractListParams{
		PageSize: llamacloudprod.Int(20),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, extract := range page.Items {
		t.Logf("%+v\n", extract.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, extract := range page.Items {
			t.Logf("%+v\n", extract.ID)
		}
	}
}
