// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/llamacloud-prod-go"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/testutil"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
)

func TestUsage(t *testing.T) {
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
	parsing, err := client.Parsing.New(context.TODO(), llamacloudprod.ParsingNewParams{
		Tier:    llamacloudprod.ParsingNewParamsTierAgentic,
		Version: llamacloudprod.ParsingNewParamsVersionLatest,
		FileID:  llamacloudprod.String("abc1234"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", parsing.ID)
}
