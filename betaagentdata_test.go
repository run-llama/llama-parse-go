// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud_test

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
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.AgentData.New(context.TODO(), llamacloud.BetaAgentDataNewParams{
		Data: map[string]any{
			"foo": "bar",
		},
		DeploymentName: "deployment_name",
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Collection:     llamacloud.String("collection"),
	})
	if err != nil {
		var apierr *llamacloud.Error
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
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.AgentData.Update(
		context.TODO(),
		"item_id",
		llamacloud.BetaAgentDataUpdateParams{
			Data: map[string]any{
				"foo": "bar",
			},
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

func TestBetaAgentDataDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.AgentData.Delete(
		context.TODO(),
		"item_id",
		llamacloud.BetaAgentDataDeleteParams{
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

func TestBetaAgentDataAggregateWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.AgentData.Aggregate(context.TODO(), llamacloud.BetaAgentDataAggregateParams{
		DeploymentName: "deployment_name",
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Collection:     llamacloud.String("collection"),
		Count:          llamacloud.Bool(true),
		Filter: map[string]llamacloud.BetaAgentDataAggregateParamsFilter{
			"foo": {
				Eq: llamacloud.BetaAgentDataAggregateParamsFilterEqUnion{
					OfFloat: llamacloud.Float(0),
				},
				Excludes: []*llamacloud.BetaAgentDataAggregateParamsFilterExcludeUnion{{
					OfFloat: llamacloud.Float(0),
				}},
				Gt: llamacloud.BetaAgentDataAggregateParamsFilterGtUnion{
					OfFloat: llamacloud.Float(0),
				},
				Gte: llamacloud.BetaAgentDataAggregateParamsFilterGteUnion{
					OfFloat: llamacloud.Float(0),
				},
				Includes: []*llamacloud.BetaAgentDataAggregateParamsFilterIncludeUnion{{
					OfFloat: llamacloud.Float(0),
				}},
				Lt: llamacloud.BetaAgentDataAggregateParamsFilterLtUnion{
					OfFloat: llamacloud.Float(0),
				},
				Lte: llamacloud.BetaAgentDataAggregateParamsFilterLteUnion{
					OfFloat: llamacloud.Float(0),
				},
				Ne: llamacloud.BetaAgentDataAggregateParamsFilterNeUnion{
					OfFloat: llamacloud.Float(0),
				},
			},
		},
		First:     llamacloud.Bool(true),
		GroupBy:   []string{"string"},
		Offset:    llamacloud.Int(0),
		OrderBy:   llamacloud.String("order_by"),
		PageSize:  llamacloud.Int(0),
		PageToken: llamacloud.String("page_token"),
	})
	if err != nil {
		var apierr *llamacloud.Error
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
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.AgentData.DeleteByQuery(context.TODO(), llamacloud.BetaAgentDataDeleteByQueryParams{
		DeploymentName: "deployment_name",
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Collection:     llamacloud.String("collection"),
		Filter: map[string]llamacloud.BetaAgentDataDeleteByQueryParamsFilter{
			"foo": {
				Eq: llamacloud.BetaAgentDataDeleteByQueryParamsFilterEqUnion{
					OfFloat: llamacloud.Float(0),
				},
				Excludes: []*llamacloud.BetaAgentDataDeleteByQueryParamsFilterExcludeUnion{{
					OfFloat: llamacloud.Float(0),
				}},
				Gt: llamacloud.BetaAgentDataDeleteByQueryParamsFilterGtUnion{
					OfFloat: llamacloud.Float(0),
				},
				Gte: llamacloud.BetaAgentDataDeleteByQueryParamsFilterGteUnion{
					OfFloat: llamacloud.Float(0),
				},
				Includes: []*llamacloud.BetaAgentDataDeleteByQueryParamsFilterIncludeUnion{{
					OfFloat: llamacloud.Float(0),
				}},
				Lt: llamacloud.BetaAgentDataDeleteByQueryParamsFilterLtUnion{
					OfFloat: llamacloud.Float(0),
				},
				Lte: llamacloud.BetaAgentDataDeleteByQueryParamsFilterLteUnion{
					OfFloat: llamacloud.Float(0),
				},
				Ne: llamacloud.BetaAgentDataDeleteByQueryParamsFilterNeUnion{
					OfFloat: llamacloud.Float(0),
				},
			},
		},
	})
	if err != nil {
		var apierr *llamacloud.Error
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
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.AgentData.Get(
		context.TODO(),
		"item_id",
		llamacloud.BetaAgentDataGetParams{
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

func TestBetaAgentDataSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.AgentData.Search(context.TODO(), llamacloud.BetaAgentDataSearchParams{
		DeploymentName: "deployment_name",
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Collection:     llamacloud.String("collection"),
		Filter: map[string]llamacloud.BetaAgentDataSearchParamsFilter{
			"foo": {
				Eq: llamacloud.BetaAgentDataSearchParamsFilterEqUnion{
					OfFloat: llamacloud.Float(0),
				},
				Excludes: []*llamacloud.BetaAgentDataSearchParamsFilterExcludeUnion{{
					OfFloat: llamacloud.Float(0),
				}},
				Gt: llamacloud.BetaAgentDataSearchParamsFilterGtUnion{
					OfFloat: llamacloud.Float(0),
				},
				Gte: llamacloud.BetaAgentDataSearchParamsFilterGteUnion{
					OfFloat: llamacloud.Float(0),
				},
				Includes: []*llamacloud.BetaAgentDataSearchParamsFilterIncludeUnion{{
					OfFloat: llamacloud.Float(0),
				}},
				Lt: llamacloud.BetaAgentDataSearchParamsFilterLtUnion{
					OfFloat: llamacloud.Float(0),
				},
				Lte: llamacloud.BetaAgentDataSearchParamsFilterLteUnion{
					OfFloat: llamacloud.Float(0),
				},
				Ne: llamacloud.BetaAgentDataSearchParamsFilterNeUnion{
					OfFloat: llamacloud.Float(0),
				},
			},
		},
		IncludeTotal: llamacloud.Bool(true),
		Offset:       llamacloud.Int(0),
		OrderBy:      llamacloud.String("order_by"),
		PageSize:     llamacloud.Int(0),
		PageToken:    llamacloud.String("page_token"),
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
