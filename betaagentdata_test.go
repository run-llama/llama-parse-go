// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/run-llama/llama-parse-go"
	"github.com/run-llama/llama-parse-go/internal/testutil"
	"github.com/run-llama/llama-parse-go/option"
)

func TestBetaAgentDataNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.AgentData.New(context.TODO(), llamacloudprod.BetaAgentDataNewParams{
		Data: map[string]any{
			"foo": "bar",
		},
		DeploymentName: "deployment_name",
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Collection:     llamacloudprod.String("collection"),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaAgentDataUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.AgentData.Update(
		context.TODO(),
		"item_id",
		llamacloudprod.BetaAgentDataUpdateParams{
			Data: map[string]any{
				"foo": "bar",
			},
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

func TestBetaAgentDataDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.AgentData.Delete(
		context.TODO(),
		"item_id",
		llamacloudprod.BetaAgentDataDeleteParams{
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

func TestBetaAgentDataAggregateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.AgentData.Aggregate(context.TODO(), llamacloudprod.BetaAgentDataAggregateParams{
		DeploymentName: "deployment_name",
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Collection:     llamacloudprod.String("collection"),
		Count:          llamacloudprod.Bool(true),
		Filter: map[string]llamacloudprod.BetaAgentDataAggregateParamsFilter{
			"foo": {
				Eq: llamacloudprod.BetaAgentDataAggregateParamsFilterEqUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Excludes: []*llamacloudprod.BetaAgentDataAggregateParamsFilterExcludeUnion{{
					OfFloat: llamacloudprod.Float(0),
				}},
				Gt: llamacloudprod.BetaAgentDataAggregateParamsFilterGtUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Gte: llamacloudprod.BetaAgentDataAggregateParamsFilterGteUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Includes: []*llamacloudprod.BetaAgentDataAggregateParamsFilterIncludeUnion{{
					OfFloat: llamacloudprod.Float(0),
				}},
				Lt: llamacloudprod.BetaAgentDataAggregateParamsFilterLtUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Lte: llamacloudprod.BetaAgentDataAggregateParamsFilterLteUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Ne: llamacloudprod.BetaAgentDataAggregateParamsFilterNeUnion{
					OfFloat: llamacloudprod.Float(0),
				},
			},
		},
		First:     llamacloudprod.Bool(true),
		GroupBy:   []string{"string"},
		Offset:    llamacloudprod.Int(0),
		OrderBy:   llamacloudprod.String("order_by"),
		PageSize:  llamacloudprod.Int(0),
		PageToken: llamacloudprod.String("page_token"),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaAgentDataDeleteByQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.AgentData.DeleteByQuery(context.TODO(), llamacloudprod.BetaAgentDataDeleteByQueryParams{
		DeploymentName: "deployment_name",
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Collection:     llamacloudprod.String("collection"),
		Filter: map[string]llamacloudprod.BetaAgentDataDeleteByQueryParamsFilter{
			"foo": {
				Eq: llamacloudprod.BetaAgentDataDeleteByQueryParamsFilterEqUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Excludes: []*llamacloudprod.BetaAgentDataDeleteByQueryParamsFilterExcludeUnion{{
					OfFloat: llamacloudprod.Float(0),
				}},
				Gt: llamacloudprod.BetaAgentDataDeleteByQueryParamsFilterGtUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Gte: llamacloudprod.BetaAgentDataDeleteByQueryParamsFilterGteUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Includes: []*llamacloudprod.BetaAgentDataDeleteByQueryParamsFilterIncludeUnion{{
					OfFloat: llamacloudprod.Float(0),
				}},
				Lt: llamacloudprod.BetaAgentDataDeleteByQueryParamsFilterLtUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Lte: llamacloudprod.BetaAgentDataDeleteByQueryParamsFilterLteUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Ne: llamacloudprod.BetaAgentDataDeleteByQueryParamsFilterNeUnion{
					OfFloat: llamacloudprod.Float(0),
				},
			},
		},
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaAgentDataGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.AgentData.Get(
		context.TODO(),
		"item_id",
		llamacloudprod.BetaAgentDataGetParams{
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

func TestBetaAgentDataSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.AgentData.Search(context.TODO(), llamacloudprod.BetaAgentDataSearchParams{
		DeploymentName: "deployment_name",
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Collection:     llamacloudprod.String("collection"),
		Filter: map[string]llamacloudprod.BetaAgentDataSearchParamsFilter{
			"foo": {
				Eq: llamacloudprod.BetaAgentDataSearchParamsFilterEqUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Excludes: []*llamacloudprod.BetaAgentDataSearchParamsFilterExcludeUnion{{
					OfFloat: llamacloudprod.Float(0),
				}},
				Gt: llamacloudprod.BetaAgentDataSearchParamsFilterGtUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Gte: llamacloudprod.BetaAgentDataSearchParamsFilterGteUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Includes: []*llamacloudprod.BetaAgentDataSearchParamsFilterIncludeUnion{{
					OfFloat: llamacloudprod.Float(0),
				}},
				Lt: llamacloudprod.BetaAgentDataSearchParamsFilterLtUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Lte: llamacloudprod.BetaAgentDataSearchParamsFilterLteUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				Ne: llamacloudprod.BetaAgentDataSearchParamsFilterNeUnion{
					OfFloat: llamacloudprod.Float(0),
				},
			},
		},
		IncludeTotal: llamacloudprod.Bool(true),
		Offset:       llamacloudprod.Int(0),
		OrderBy:      llamacloudprod.String("order_by"),
		PageSize:     llamacloudprod.Int(0),
		PageToken:    llamacloudprod.String("page_token"),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
