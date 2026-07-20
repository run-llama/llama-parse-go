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

func TestRetrieverNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Retrievers.New(context.TODO(), llamacloud.RetrieverNewParams{
		RetrieverCreate: llamacloud.RetrieverCreateParam{
			Name: "x",
			Pipelines: []llamacloud.RetrieverPipelineParam{{
				Description: llamacloud.String("description"),
				Name:        llamacloud.String("x"),
				PipelineID:  "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				PresetRetrievalParameters: llamacloud.PresetRetrievalParams{
					Alpha:                       llamacloud.Float(0),
					ClassName:                   llamacloud.String("class_name"),
					DenseSimilarityCutoff:       llamacloud.Float(0),
					DenseSimilarityTopK:         llamacloud.Int(1),
					EnableReranking:             llamacloud.Bool(true),
					FilesTopK:                   llamacloud.Int(1),
					RerankTopN:                  llamacloud.Int(1),
					RetrievalMode:               llamacloud.RetrievalModeAutoRouted,
					RetrieveImageNodes:          llamacloud.Bool(true),
					RetrievePageFigureNodes:     llamacloud.Bool(true),
					RetrievePageScreenshotNodes: llamacloud.Bool(true),
					SearchFilters: llamacloud.MetadataFiltersParam{
						Filters: []llamacloud.MetadataFiltersFilterUnionParam{{
							OfMetadataFilter: &llamacloud.MetadataFiltersFilterMetadataFilterParam{
								Key: "key",
								Value: llamacloud.MetadataFiltersFilterMetadataFilterValueUnionParam{
									OfFloat: llamacloud.Float(0),
								},
								Operator: "!=",
							},
						}},
						Condition: llamacloud.MetadataFiltersConditionAnd,
					},
					SearchFiltersInferenceSchema: map[string]*llamacloud.PresetRetrievalParamsSearchFiltersInferenceSchemaUnion{
						"foo": {
							OfAnyMap: map[string]any{
								"foo": "bar",
							},
						},
					},
					SparseSimilarityTopK: llamacloud.Int(1),
				},
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

func TestRetrieverUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Retrievers.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloud.RetrieverUpdateParams{
			Pipelines: []llamacloud.RetrieverPipelineParam{{
				Description: llamacloud.String("description"),
				Name:        llamacloud.String("x"),
				PipelineID:  "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				PresetRetrievalParameters: llamacloud.PresetRetrievalParams{
					Alpha:                       llamacloud.Float(0),
					ClassName:                   llamacloud.String("class_name"),
					DenseSimilarityCutoff:       llamacloud.Float(0),
					DenseSimilarityTopK:         llamacloud.Int(1),
					EnableReranking:             llamacloud.Bool(true),
					FilesTopK:                   llamacloud.Int(1),
					RerankTopN:                  llamacloud.Int(1),
					RetrievalMode:               llamacloud.RetrievalModeAutoRouted,
					RetrieveImageNodes:          llamacloud.Bool(true),
					RetrievePageFigureNodes:     llamacloud.Bool(true),
					RetrievePageScreenshotNodes: llamacloud.Bool(true),
					SearchFilters: llamacloud.MetadataFiltersParam{
						Filters: []llamacloud.MetadataFiltersFilterUnionParam{{
							OfMetadataFilter: &llamacloud.MetadataFiltersFilterMetadataFilterParam{
								Key: "key",
								Value: llamacloud.MetadataFiltersFilterMetadataFilterValueUnionParam{
									OfFloat: llamacloud.Float(0),
								},
								Operator: "!=",
							},
						}},
						Condition: llamacloud.MetadataFiltersConditionAnd,
					},
					SearchFiltersInferenceSchema: map[string]*llamacloud.PresetRetrievalParamsSearchFiltersInferenceSchemaUnion{
						"foo": {
							OfAnyMap: map[string]any{
								"foo": "bar",
							},
						},
					},
					SparseSimilarityTopK: llamacloud.Int(1),
				},
			}},
			OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Name:           llamacloud.String("name"),
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

func TestRetrieverListWithOptionalParams(t *testing.T) {
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
	_, err := client.Retrievers.List(context.TODO(), llamacloud.RetrieverListParams{
		Name:           llamacloud.String("name"),
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

func TestRetrieverDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Retrievers.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloud.RetrieverDeleteParams{
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

func TestRetrieverGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Retrievers.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloud.RetrieverGetParams{
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

func TestRetrieverSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Retrievers.Search(context.TODO(), llamacloud.RetrieverSearchParams{
		Query:          "x",
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Mode:           llamacloud.CompositeRetrievalModeFull,
		Pipelines: []llamacloud.RetrieverPipelineParam{{
			Description: llamacloud.String("description"),
			Name:        llamacloud.String("x"),
			PipelineID:  "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			PresetRetrievalParameters: llamacloud.PresetRetrievalParams{
				Alpha:                       llamacloud.Float(0),
				ClassName:                   llamacloud.String("class_name"),
				DenseSimilarityCutoff:       llamacloud.Float(0),
				DenseSimilarityTopK:         llamacloud.Int(1),
				EnableReranking:             llamacloud.Bool(true),
				FilesTopK:                   llamacloud.Int(1),
				RerankTopN:                  llamacloud.Int(1),
				RetrievalMode:               llamacloud.RetrievalModeAutoRouted,
				RetrieveImageNodes:          llamacloud.Bool(true),
				RetrievePageFigureNodes:     llamacloud.Bool(true),
				RetrievePageScreenshotNodes: llamacloud.Bool(true),
				SearchFilters: llamacloud.MetadataFiltersParam{
					Filters: []llamacloud.MetadataFiltersFilterUnionParam{{
						OfMetadataFilter: &llamacloud.MetadataFiltersFilterMetadataFilterParam{
							Key: "key",
							Value: llamacloud.MetadataFiltersFilterMetadataFilterValueUnionParam{
								OfFloat: llamacloud.Float(0),
							},
							Operator: "!=",
						},
					}},
					Condition: llamacloud.MetadataFiltersConditionAnd,
				},
				SearchFiltersInferenceSchema: map[string]*llamacloud.PresetRetrievalParamsSearchFiltersInferenceSchemaUnion{
					"foo": {
						OfAnyMap: map[string]any{
							"foo": "bar",
						},
					},
				},
				SparseSimilarityTopK: llamacloud.Int(1),
			},
		}},
		RerankConfig: llamacloud.ReRankConfigParam{
			TopN: llamacloud.Int(1),
			Type: llamacloud.ReRankConfigTypeBedrock,
		},
		RerankTopN: llamacloud.Int(0),
	})
	if err != nil {
		var apierr *llamacloud.Error
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
	client := llamacloud.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Retrievers.Upsert(context.TODO(), llamacloud.RetrieverUpsertParams{
		RetrieverCreate: llamacloud.RetrieverCreateParam{
			Name: "x",
			Pipelines: []llamacloud.RetrieverPipelineParam{{
				Description: llamacloud.String("description"),
				Name:        llamacloud.String("x"),
				PipelineID:  "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				PresetRetrievalParameters: llamacloud.PresetRetrievalParams{
					Alpha:                       llamacloud.Float(0),
					ClassName:                   llamacloud.String("class_name"),
					DenseSimilarityCutoff:       llamacloud.Float(0),
					DenseSimilarityTopK:         llamacloud.Int(1),
					EnableReranking:             llamacloud.Bool(true),
					FilesTopK:                   llamacloud.Int(1),
					RerankTopN:                  llamacloud.Int(1),
					RetrievalMode:               llamacloud.RetrievalModeAutoRouted,
					RetrieveImageNodes:          llamacloud.Bool(true),
					RetrievePageFigureNodes:     llamacloud.Bool(true),
					RetrievePageScreenshotNodes: llamacloud.Bool(true),
					SearchFilters: llamacloud.MetadataFiltersParam{
						Filters: []llamacloud.MetadataFiltersFilterUnionParam{{
							OfMetadataFilter: &llamacloud.MetadataFiltersFilterMetadataFilterParam{
								Key: "key",
								Value: llamacloud.MetadataFiltersFilterMetadataFilterValueUnionParam{
									OfFloat: llamacloud.Float(0),
								},
								Operator: "!=",
							},
						}},
						Condition: llamacloud.MetadataFiltersConditionAnd,
					},
					SearchFiltersInferenceSchema: map[string]*llamacloud.PresetRetrievalParamsSearchFiltersInferenceSchemaUnion{
						"foo": {
							OfAnyMap: map[string]any{
								"foo": "bar",
							},
						},
					},
					SparseSimilarityTopK: llamacloud.Int(1),
				},
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
