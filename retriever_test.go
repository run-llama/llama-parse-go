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

func TestRetrieverNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Retrievers.New(context.TODO(), llamacloudprod.RetrieverNewParams{
		RetrieverCreate: llamacloudprod.RetrieverCreateParam{
			Name: "x",
			Pipelines: []llamacloudprod.RetrieverPipelineParam{{
				Description: llamacloudprod.String("description"),
				Name:        llamacloudprod.String("x"),
				PipelineID:  "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				PresetRetrievalParameters: llamacloudprod.PresetRetrievalParams{
					Alpha:                       llamacloudprod.Float(0),
					ClassName:                   llamacloudprod.String("class_name"),
					DenseSimilarityCutoff:       llamacloudprod.Float(0),
					DenseSimilarityTopK:         llamacloudprod.Int(1),
					EnableReranking:             llamacloudprod.Bool(true),
					FilesTopK:                   llamacloudprod.Int(1),
					RerankTopN:                  llamacloudprod.Int(1),
					RetrievalMode:               llamacloudprod.RetrievalModeChunks,
					RetrieveImageNodes:          llamacloudprod.Bool(true),
					RetrievePageFigureNodes:     llamacloudprod.Bool(true),
					RetrievePageScreenshotNodes: llamacloudprod.Bool(true),
					SearchFilters: llamacloudprod.MetadataFiltersParam{
						Filters: []llamacloudprod.MetadataFiltersFilterUnionParam{{
							OfMetadataFilter: &llamacloudprod.MetadataFiltersFilterMetadataFilterParam{
								Key: "key",
								Value: llamacloudprod.MetadataFiltersFilterMetadataFilterValueUnionParam{
									OfFloat: llamacloudprod.Float(0),
								},
								Operator: "==",
							},
						}},
						Condition: llamacloudprod.MetadataFiltersConditionAnd,
					},
					SearchFiltersInferenceSchema: map[string]*llamacloudprod.PresetRetrievalParamsSearchFiltersInferenceSchemaUnion{
						"foo": {
							OfAnyMap: map[string]any{
								"foo": "bar",
							},
						},
					},
					SparseSimilarityTopK: llamacloudprod.Int(1),
				},
			}},
		},
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRetrieverUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Retrievers.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.RetrieverUpdateParams{
			Pipelines: []llamacloudprod.RetrieverPipelineParam{{
				Description: llamacloudprod.String("description"),
				Name:        llamacloudprod.String("x"),
				PipelineID:  "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				PresetRetrievalParameters: llamacloudprod.PresetRetrievalParams{
					Alpha:                       llamacloudprod.Float(0),
					ClassName:                   llamacloudprod.String("class_name"),
					DenseSimilarityCutoff:       llamacloudprod.Float(0),
					DenseSimilarityTopK:         llamacloudprod.Int(1),
					EnableReranking:             llamacloudprod.Bool(true),
					FilesTopK:                   llamacloudprod.Int(1),
					RerankTopN:                  llamacloudprod.Int(1),
					RetrievalMode:               llamacloudprod.RetrievalModeChunks,
					RetrieveImageNodes:          llamacloudprod.Bool(true),
					RetrievePageFigureNodes:     llamacloudprod.Bool(true),
					RetrievePageScreenshotNodes: llamacloudprod.Bool(true),
					SearchFilters: llamacloudprod.MetadataFiltersParam{
						Filters: []llamacloudprod.MetadataFiltersFilterUnionParam{{
							OfMetadataFilter: &llamacloudprod.MetadataFiltersFilterMetadataFilterParam{
								Key: "key",
								Value: llamacloudprod.MetadataFiltersFilterMetadataFilterValueUnionParam{
									OfFloat: llamacloudprod.Float(0),
								},
								Operator: "==",
							},
						}},
						Condition: llamacloudprod.MetadataFiltersConditionAnd,
					},
					SearchFiltersInferenceSchema: map[string]*llamacloudprod.PresetRetrievalParamsSearchFiltersInferenceSchemaUnion{
						"foo": {
							OfAnyMap: map[string]any{
								"foo": "bar",
							},
						},
					},
					SparseSimilarityTopK: llamacloudprod.Int(1),
				},
			}},
			OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:           llamacloudprod.String("name"),
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

func TestRetrieverListWithOptionalParams(t *testing.T) {
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
	_, err := client.Retrievers.List(context.TODO(), llamacloudprod.RetrieverListParams{
		Name:           llamacloudprod.String("name"),
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestRetrieverDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Retrievers.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.RetrieverDeleteParams{
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

func TestRetrieverGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Retrievers.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.RetrieverGetParams{
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

func TestRetrieverSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Retrievers.Search(context.TODO(), llamacloudprod.RetrieverSearchParams{
		Query:          "x",
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Mode:           llamacloudprod.CompositeRetrievalModeRouting,
		Pipelines: []llamacloudprod.RetrieverPipelineParam{{
			Description: llamacloudprod.String("description"),
			Name:        llamacloudprod.String("x"),
			PipelineID:  "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			PresetRetrievalParameters: llamacloudprod.PresetRetrievalParams{
				Alpha:                       llamacloudprod.Float(0),
				ClassName:                   llamacloudprod.String("class_name"),
				DenseSimilarityCutoff:       llamacloudprod.Float(0),
				DenseSimilarityTopK:         llamacloudprod.Int(1),
				EnableReranking:             llamacloudprod.Bool(true),
				FilesTopK:                   llamacloudprod.Int(1),
				RerankTopN:                  llamacloudprod.Int(1),
				RetrievalMode:               llamacloudprod.RetrievalModeChunks,
				RetrieveImageNodes:          llamacloudprod.Bool(true),
				RetrievePageFigureNodes:     llamacloudprod.Bool(true),
				RetrievePageScreenshotNodes: llamacloudprod.Bool(true),
				SearchFilters: llamacloudprod.MetadataFiltersParam{
					Filters: []llamacloudprod.MetadataFiltersFilterUnionParam{{
						OfMetadataFilter: &llamacloudprod.MetadataFiltersFilterMetadataFilterParam{
							Key: "key",
							Value: llamacloudprod.MetadataFiltersFilterMetadataFilterValueUnionParam{
								OfFloat: llamacloudprod.Float(0),
							},
							Operator: "==",
						},
					}},
					Condition: llamacloudprod.MetadataFiltersConditionAnd,
				},
				SearchFiltersInferenceSchema: map[string]*llamacloudprod.PresetRetrievalParamsSearchFiltersInferenceSchemaUnion{
					"foo": {
						OfAnyMap: map[string]any{
							"foo": "bar",
						},
					},
				},
				SparseSimilarityTopK: llamacloudprod.Int(1),
			},
		}},
		RerankConfig: llamacloudprod.ReRankConfigParam{
			TopN: llamacloudprod.Int(1),
			Type: llamacloudprod.ReRankConfigTypeSystemDefault,
		},
		RerankTopN: llamacloudprod.Int(0),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRetrieverUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Retrievers.Upsert(context.TODO(), llamacloudprod.RetrieverUpsertParams{
		RetrieverCreate: llamacloudprod.RetrieverCreateParam{
			Name: "x",
			Pipelines: []llamacloudprod.RetrieverPipelineParam{{
				Description: llamacloudprod.String("description"),
				Name:        llamacloudprod.String("x"),
				PipelineID:  "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				PresetRetrievalParameters: llamacloudprod.PresetRetrievalParams{
					Alpha:                       llamacloudprod.Float(0),
					ClassName:                   llamacloudprod.String("class_name"),
					DenseSimilarityCutoff:       llamacloudprod.Float(0),
					DenseSimilarityTopK:         llamacloudprod.Int(1),
					EnableReranking:             llamacloudprod.Bool(true),
					FilesTopK:                   llamacloudprod.Int(1),
					RerankTopN:                  llamacloudprod.Int(1),
					RetrievalMode:               llamacloudprod.RetrievalModeChunks,
					RetrieveImageNodes:          llamacloudprod.Bool(true),
					RetrievePageFigureNodes:     llamacloudprod.Bool(true),
					RetrievePageScreenshotNodes: llamacloudprod.Bool(true),
					SearchFilters: llamacloudprod.MetadataFiltersParam{
						Filters: []llamacloudprod.MetadataFiltersFilterUnionParam{{
							OfMetadataFilter: &llamacloudprod.MetadataFiltersFilterMetadataFilterParam{
								Key: "key",
								Value: llamacloudprod.MetadataFiltersFilterMetadataFilterValueUnionParam{
									OfFloat: llamacloudprod.Float(0),
								},
								Operator: "==",
							},
						}},
						Condition: llamacloudprod.MetadataFiltersConditionAnd,
					},
					SearchFiltersInferenceSchema: map[string]*llamacloudprod.PresetRetrievalParamsSearchFiltersInferenceSchemaUnion{
						"foo": {
							OfAnyMap: map[string]any{
								"foo": "bar",
							},
						},
					},
					SparseSimilarityTopK: llamacloudprod.Int(1),
				},
			}},
		},
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
