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

func TestBetaRetrievalGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Retrieval.Get(context.TODO(), llamacloudprod.BetaRetrievalGetParams{
		IndexID:        "idx-abc123",
		Query:          "What are the key findings?",
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CustomFilters: map[string]*llamacloudprod.BetaRetrievalGetParamsCustomFilterUnion{
			"foo": {
				OfFilterTypeUnionStrIntBoolFloat: &llamacloudprod.BetaRetrievalGetParamsCustomFilterFilterTypeUnionStrIntBoolFloat{
					Operator: "eq",
					Value: llamacloudprod.BetaRetrievalGetParamsCustomFilterFilterTypeUnionStrIntBoolFloatValueUnion{
						OfString: llamacloudprod.String("string"),
					},
				},
			},
		},
		FullTextPipelineWeight: llamacloudprod.Float(0),
		NumCandidates:          llamacloudprod.Int(0),
		Rerank: llamacloudprod.BetaRetrievalGetParamsRerank{
			Enabled: llamacloudprod.Bool(true),
			TopN:    llamacloudprod.Int(5),
		},
		ScoreThreshold: llamacloudprod.Float(0),
		StaticFilters: llamacloudprod.BetaRetrievalGetParamsStaticFilters{
			ParsedDirectoryFileID: llamacloudprod.BetaRetrievalGetParamsStaticFiltersParsedDirectoryFileID{
				Operator: "eq",
				Value: llamacloudprod.BetaRetrievalGetParamsStaticFiltersParsedDirectoryFileIDValueUnion{
					OfString: llamacloudprod.String("string"),
				},
			},
		},
		TopK:                 llamacloudprod.Int(10),
		VectorPipelineWeight: llamacloudprod.Float(0),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
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
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Retrieval.Find(context.TODO(), llamacloudprod.BetaRetrievalFindParams{
		IndexID:          "idx-abc123",
		OrganizationID:   llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:        llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		FileName:         llamacloudprod.String("file_name"),
		FileNameContains: llamacloudprod.String("file_name_contains"),
		PageSize:         llamacloudprod.Int(0),
		PageToken:        llamacloudprod.String("page_token"),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
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
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Retrieval.Grep(context.TODO(), llamacloudprod.BetaRetrievalGrepParams{
		FileID:         "file_id",
		IndexID:        "idx-abc123",
		Pattern:        "revenue|profit",
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ContextChars:   llamacloudprod.Int(0),
		PageSize:       llamacloudprod.Int(0),
		PageToken:      llamacloudprod.String("page_token"),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
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
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Beta.Retrieval.Read(context.TODO(), llamacloudprod.BetaRetrievalReadParams{
		FileID:         "file_id",
		IndexID:        "idx-abc123",
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		MaxLength:      llamacloudprod.Int(0),
		Offset:         llamacloudprod.Int(0),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
