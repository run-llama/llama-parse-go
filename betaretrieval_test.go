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

func TestBetaRetrievalGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Retrieval.Get(context.TODO(), llamacloud.BetaRetrievalGetParams{
		IndexID:        "idx-abc123",
		Query:          "What are the key findings?",
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CustomFilters: map[string]*llamacloud.BetaRetrievalGetParamsCustomFilterUnion{
			"foo": {
				OfFilterTypeUnionStrIntBoolFloat: &llamacloud.BetaRetrievalGetParamsCustomFilterFilterTypeUnionStrIntBoolFloat{
					Operator: "eq",
					Value: llamacloud.BetaRetrievalGetParamsCustomFilterFilterTypeUnionStrIntBoolFloatValueUnion{
						OfString: llamacloud.String("string"),
					},
				},
			},
		},
		FullTextPipelineWeight: llamacloud.Float(0),
		NumCandidates:          llamacloud.Int(0),
		Rerank: llamacloud.BetaRetrievalGetParamsRerank{
			Enabled: llamacloud.Bool(true),
			TopN:    llamacloud.Int(5),
		},
		ScoreThreshold: llamacloud.Float(0),
		StaticFilters: llamacloud.BetaRetrievalGetParamsStaticFilters{
			ParsedDirectoryFileID: llamacloud.BetaRetrievalGetParamsStaticFiltersParsedDirectoryFileID{
				Operator: "eq",
				Value: llamacloud.BetaRetrievalGetParamsStaticFiltersParsedDirectoryFileIDValueUnion{
					OfString: llamacloud.String("string"),
				},
			},
		},
		TopK:                 llamacloud.Int(10),
		VectorPipelineWeight: llamacloud.Float(0),
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRetrievalFindWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Retrieval.Find(context.TODO(), llamacloud.BetaRetrievalFindParams{
		IndexID:          "idx-abc123",
		OrganizationID:   llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:        llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FileName:         llamacloud.String("file_name"),
		FileNameContains: llamacloud.String("file_name_contains"),
		PageSize:         llamacloud.Int(0),
		PageToken:        llamacloud.String("page_token"),
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRetrievalGrepWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Retrieval.Grep(context.TODO(), llamacloud.BetaRetrievalGrepParams{
		FileID:         "file_id",
		IndexID:        "idx-abc123",
		Pattern:        "revenue|profit",
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ContextChars:   llamacloud.Int(0),
		PageSize:       llamacloud.Int(0),
		PageToken:      llamacloud.String("page_token"),
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaRetrievalReadWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Retrieval.Read(context.TODO(), llamacloud.BetaRetrievalReadParams{
		FileID:         "file_id",
		IndexID:        "idx-abc123",
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		MaxLength:      llamacloud.Int(0),
		Offset:         llamacloud.Int(0),
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
