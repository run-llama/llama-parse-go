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

func TestPipelineNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.New(context.TODO(), llamacloudprod.PipelineNewParams{
		PipelineCreate: llamacloudprod.PipelineCreateParam{
			Name: "x",
			DataSink: llamacloudprod.DataSinkCreateParam{
				Component: llamacloudprod.DataSinkCreateComponentUnionParam{
					OfAnyMap: map[string]any{
						"foo": "bar",
					},
				},
				Name:     "name",
				SinkType: llamacloudprod.DataSinkCreateSinkTypePinecone,
			},
			DataSinkID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			EmbeddingConfig: llamacloudprod.PipelineCreateEmbeddingConfigUnionParam{
				OfAzureEmbedding: &llamacloudprod.AzureOpenAIEmbeddingConfigParam{
					Component: llamacloudprod.AzureOpenAIEmbeddingParam{
						AdditionalKwargs: map[string]any{
							"foo": "bar",
						},
						APIBase:         llamacloudprod.String("api_base"),
						APIKey:          llamacloudprod.String("api_key"),
						APIVersion:      llamacloudprod.String("api_version"),
						AzureDeployment: llamacloudprod.String("azure_deployment"),
						AzureEndpoint:   llamacloudprod.String("azure_endpoint"),
						ClassName:       llamacloudprod.String("class_name"),
						DefaultHeaders: map[string]string{
							"foo": "string",
						},
						Dimensions:     llamacloudprod.Int(0),
						EmbedBatchSize: llamacloudprod.Int(1),
						MaxRetries:     llamacloudprod.Int(0),
						ModelName:      llamacloudprod.String("model_name"),
						NumWorkers:     llamacloudprod.Int(0),
						ReuseClient:    llamacloudprod.Bool(true),
						Timeout:        llamacloudprod.Float(0),
					},
					Type: llamacloudprod.AzureOpenAIEmbeddingConfigTypeAzureEmbedding,
				},
			},
			EmbeddingModelConfigID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LlamaParseParameters: llamacloudprod.LlamaParseParameters{
				AdaptiveLongTable:                        llamacloudprod.Bool(true),
				AggressiveTableExtraction:                llamacloudprod.Bool(true),
				AnnotateLinks:                            llamacloudprod.Bool(true),
				AutoMode:                                 llamacloudprod.Bool(true),
				AutoModeConfigurationJson:                llamacloudprod.String("auto_mode_configuration_json"),
				AutoModeTriggerOnImageInPage:             llamacloudprod.Bool(true),
				AutoModeTriggerOnRegexpInPage:            llamacloudprod.String("auto_mode_trigger_on_regexp_in_page"),
				AutoModeTriggerOnTableInPage:             llamacloudprod.Bool(true),
				AutoModeTriggerOnTextInPage:              llamacloudprod.String("auto_mode_trigger_on_text_in_page"),
				AzureOpenAIAPIVersion:                    llamacloudprod.String("azure_openai_api_version"),
				AzureOpenAIDeploymentName:                llamacloudprod.String("azure_openai_deployment_name"),
				AzureOpenAIEndpoint:                      llamacloudprod.String("azure_openai_endpoint"),
				AzureOpenAIKey:                           llamacloudprod.String("azure_openai_key"),
				BboxBottom:                               llamacloudprod.Float(0),
				BboxLeft:                                 llamacloudprod.Float(0),
				BboxRight:                                llamacloudprod.Float(0),
				BboxTop:                                  llamacloudprod.Float(0),
				BoundingBox:                              llamacloudprod.String("bounding_box"),
				CompactMarkdownTable:                     llamacloudprod.Bool(true),
				ComplementalFormattingInstruction:        llamacloudprod.String("complemental_formatting_instruction"),
				ContentGuidelineInstruction:              llamacloudprod.String("content_guideline_instruction"),
				ContinuousMode:                           llamacloudprod.Bool(true),
				DisableImageExtraction:                   llamacloudprod.Bool(true),
				DisableOcr:                               llamacloudprod.Bool(true),
				DisableReconstruction:                    llamacloudprod.Bool(true),
				DoNotCache:                               llamacloudprod.Bool(true),
				DoNotUnrollColumns:                       llamacloudprod.Bool(true),
				EnableCostOptimizer:                      llamacloudprod.Bool(true),
				ExtractCharts:                            llamacloudprod.Bool(true),
				ExtractLayout:                            llamacloudprod.Bool(true),
				ExtractPrintedPageNumber:                 llamacloudprod.Bool(true),
				FastMode:                                 llamacloudprod.Bool(true),
				FormattingInstruction:                    llamacloudprod.String("formatting_instruction"),
				Gpt4oAPIKey:                              llamacloudprod.String("gpt4o_api_key"),
				Gpt4oMode:                                llamacloudprod.Bool(true),
				GuessXlsxSheetName:                       llamacloudprod.Bool(true),
				HideFooters:                              llamacloudprod.Bool(true),
				HideHeaders:                              llamacloudprod.Bool(true),
				HighResOcr:                               llamacloudprod.Bool(true),
				HTMLMakeAllElementsVisible:               llamacloudprod.Bool(true),
				HTMLRemoveFixedElements:                  llamacloudprod.Bool(true),
				HTMLRemoveNavigationElements:             llamacloudprod.Bool(true),
				HTTPProxy:                                llamacloudprod.String("http_proxy"),
				IgnoreDocumentElementsForLayoutDetection: llamacloudprod.Bool(true),
				ImagesToSave:                             []string{"screenshot"},
				InlineImagesInMarkdown:                   llamacloudprod.Bool(true),
				InputS3Path:                              llamacloudprod.String("input_s3_path"),
				InputS3Region:                            llamacloudprod.String("input_s3_region"),
				InputURL:                                 llamacloudprod.String("input_url"),
				InternalIsScreenshotJob:                  llamacloudprod.Bool(true),
				InvalidateCache:                          llamacloudprod.Bool(true),
				IsFormattingInstruction:                  llamacloudprod.Bool(true),
				JobTimeoutExtraTimePerPageInSeconds:      llamacloudprod.Float(0),
				JobTimeoutInSeconds:                      llamacloudprod.Float(0),
				KeepPageSeparatorWhenMergingTables:       llamacloudprod.Bool(true),
				Languages:                                []llamacloudprod.ParsingLanguages{llamacloudprod.ParsingLanguagesAf},
				LayoutAware:                              llamacloudprod.Bool(true),
				LineLevelBoundingBox:                     llamacloudprod.Bool(true),
				MarkdownTableMultilineHeaderSeparator:    llamacloudprod.String("markdown_table_multiline_header_separator"),
				MaxPages:                                 llamacloudprod.Int(0),
				MaxPagesEnforced:                         llamacloudprod.Int(0),
				MergeTablesAcrossPagesInMarkdown:         llamacloudprod.Bool(true),
				Model:                                    llamacloudprod.String("model"),
				OutlinedTableExtraction:                  llamacloudprod.Bool(true),
				OutputPdfOfDocument:                      llamacloudprod.Bool(true),
				OutputS3PathPrefix:                       llamacloudprod.String("output_s3_path_prefix"),
				OutputS3Region:                           llamacloudprod.String("output_s3_region"),
				OutputTablesAsHTML:                       llamacloudprod.Bool(true),
				PageErrorTolerance:                       llamacloudprod.Float(0),
				PageFooterPrefix:                         llamacloudprod.String("page_footer_prefix"),
				PageFooterSuffix:                         llamacloudprod.String("page_footer_suffix"),
				PageHeaderPrefix:                         llamacloudprod.String("page_header_prefix"),
				PageHeaderSuffix:                         llamacloudprod.String("page_header_suffix"),
				PagePrefix:                               llamacloudprod.String("page_prefix"),
				PageSeparator:                            llamacloudprod.String("page_separator"),
				PageSuffix:                               llamacloudprod.String("page_suffix"),
				ParseMode:                                llamacloudprod.ParsingModeParsePageWithoutLlm,
				ParsingInstruction:                       llamacloudprod.String("parsing_instruction"),
				PreciseBoundingBox:                       llamacloudprod.Bool(true),
				PremiumMode:                              llamacloudprod.Bool(true),
				PresentationOutOfBoundsContent:           llamacloudprod.Bool(true),
				PresentationSkipEmbeddedData:             llamacloudprod.Bool(true),
				PreserveLayoutAlignmentAcrossPages:       llamacloudprod.Bool(true),
				PreserveVerySmallText:                    llamacloudprod.Bool(true),
				Preset:                                   llamacloudprod.String("preset"),
				Priority:                                 llamacloudprod.LlamaParseParametersPriorityLow,
				ProjectID:                                llamacloudprod.String("project_id"),
				RemoveHiddenText:                         llamacloudprod.Bool(true),
				ReplaceFailedPageMode:                    llamacloudprod.FailPageModeRawText,
				ReplaceFailedPageWithErrorMessagePrefix:  llamacloudprod.String("replace_failed_page_with_error_message_prefix"),
				ReplaceFailedPageWithErrorMessageSuffix:  llamacloudprod.String("replace_failed_page_with_error_message_suffix"),
				SaveImages:                               llamacloudprod.Bool(true),
				SkipDiagonalText:                         llamacloudprod.Bool(true),
				SpecializedChartParsingAgentic:           llamacloudprod.Bool(true),
				SpecializedChartParsingEfficient:         llamacloudprod.Bool(true),
				SpecializedChartParsingPlus:              llamacloudprod.Bool(true),
				SpecializedImageParsing:                  llamacloudprod.Bool(true),
				SpreadsheetExtractSubTables:              llamacloudprod.Bool(true),
				SpreadsheetForceFormulaComputation:       llamacloudprod.Bool(true),
				SpreadsheetIncludeHiddenSheets:           llamacloudprod.Bool(true),
				StrictModeBuggyFont:                      llamacloudprod.Bool(true),
				StrictModeImageExtraction:                llamacloudprod.Bool(true),
				StrictModeImageOcr:                       llamacloudprod.Bool(true),
				StrictModeReconstruction:                 llamacloudprod.Bool(true),
				StructuredOutput:                         llamacloudprod.Bool(true),
				StructuredOutputJsonSchema:               llamacloudprod.String("structured_output_json_schema"),
				StructuredOutputJsonSchemaName:           llamacloudprod.String("structured_output_json_schema_name"),
				SystemPrompt:                             llamacloudprod.String("system_prompt"),
				SystemPromptAppend:                       llamacloudprod.String("system_prompt_append"),
				TakeScreenshot:                           llamacloudprod.Bool(true),
				TargetPages:                              llamacloudprod.String("target_pages"),
				Tier:                                     llamacloudprod.String("tier"),
				UseVendorMultimodalModel:                 llamacloudprod.Bool(true),
				UserPrompt:                               llamacloudprod.String("user_prompt"),
				VendorMultimodalAPIKey:                   llamacloudprod.String("vendor_multimodal_api_key"),
				VendorMultimodalModelName:                llamacloudprod.String("vendor_multimodal_model_name"),
				Version:                                  llamacloudprod.String("version"),
				WebhookConfigurations: []llamacloudprod.LlamaParseParametersWebhookConfiguration{{
					WebhookEvents: []string{"parse.success", "parse.error"},
					WebhookHeaders: map[string]string{
						"Authorization": "Bearer sk-...",
					},
					WebhookOutputFormat: llamacloudprod.String("json"),
					WebhookURL:          llamacloudprod.String("https://example.com/webhooks/llamacloud"),
				}},
				WebhookURL: llamacloudprod.String("webhook_url"),
			},
			ManagedPipelineID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			MetadataConfig: llamacloudprod.PipelineMetadataConfigParam{
				ExcludedEmbedMetadataKeys: []string{"string"},
				ExcludedLlmMetadataKeys:   []string{"string"},
			},
			PipelineType: llamacloudprod.PipelineTypePlayground,
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
			SparseModelConfig: llamacloudprod.SparseModelConfigParam{
				ClassName: llamacloudprod.String("class_name"),
				ModelType: llamacloudprod.SparseModelConfigModelTypeSplade,
			},
			Status: llamacloudprod.String("status"),
			TransformConfig: llamacloudprod.PipelineCreateTransformConfigUnionParam{
				OfAutoTransformConfig: &llamacloudprod.AutoTransformConfigParam{
					ChunkOverlap: llamacloudprod.Int(0),
					ChunkSize:    llamacloudprod.Int(1),
					Mode:         llamacloudprod.AutoTransformConfigModeAuto,
				},
			},
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

func TestPipelineGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.PipelineGetParams{
			Query:                       "x",
			OrganizationID:              llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:                   llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
			SearchFiltersInferenceSchema: map[string]*llamacloudprod.PipelineGetParamsSearchFiltersInferenceSchemaUnion{
				"foo": {
					OfAnyMap: map[string]any{
						"foo": "bar",
					},
				},
			},
			SparseSimilarityTopK: llamacloudprod.Int(1),
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

func TestPipelineUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.PipelineUpdateParams{
			DataSink: llamacloudprod.DataSinkCreateParam{
				Component: llamacloudprod.DataSinkCreateComponentUnionParam{
					OfAnyMap: map[string]any{
						"foo": "bar",
					},
				},
				Name:     "name",
				SinkType: llamacloudprod.DataSinkCreateSinkTypePinecone,
			},
			DataSinkID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			EmbeddingConfig: llamacloudprod.PipelineUpdateParamsEmbeddingConfigUnion{
				OfAzureEmbedding: &llamacloudprod.AzureOpenAIEmbeddingConfigParam{
					Component: llamacloudprod.AzureOpenAIEmbeddingParam{
						AdditionalKwargs: map[string]any{
							"foo": "bar",
						},
						APIBase:         llamacloudprod.String("api_base"),
						APIKey:          llamacloudprod.String("api_key"),
						APIVersion:      llamacloudprod.String("api_version"),
						AzureDeployment: llamacloudprod.String("azure_deployment"),
						AzureEndpoint:   llamacloudprod.String("azure_endpoint"),
						ClassName:       llamacloudprod.String("class_name"),
						DefaultHeaders: map[string]string{
							"foo": "string",
						},
						Dimensions:     llamacloudprod.Int(0),
						EmbedBatchSize: llamacloudprod.Int(1),
						MaxRetries:     llamacloudprod.Int(0),
						ModelName:      llamacloudprod.String("model_name"),
						NumWorkers:     llamacloudprod.Int(0),
						ReuseClient:    llamacloudprod.Bool(true),
						Timeout:        llamacloudprod.Float(0),
					},
					Type: llamacloudprod.AzureOpenAIEmbeddingConfigTypeAzureEmbedding,
				},
			},
			EmbeddingModelConfigID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LlamaParseParameters: llamacloudprod.LlamaParseParameters{
				AdaptiveLongTable:                        llamacloudprod.Bool(true),
				AggressiveTableExtraction:                llamacloudprod.Bool(true),
				AnnotateLinks:                            llamacloudprod.Bool(true),
				AutoMode:                                 llamacloudprod.Bool(true),
				AutoModeConfigurationJson:                llamacloudprod.String("auto_mode_configuration_json"),
				AutoModeTriggerOnImageInPage:             llamacloudprod.Bool(true),
				AutoModeTriggerOnRegexpInPage:            llamacloudprod.String("auto_mode_trigger_on_regexp_in_page"),
				AutoModeTriggerOnTableInPage:             llamacloudprod.Bool(true),
				AutoModeTriggerOnTextInPage:              llamacloudprod.String("auto_mode_trigger_on_text_in_page"),
				AzureOpenAIAPIVersion:                    llamacloudprod.String("azure_openai_api_version"),
				AzureOpenAIDeploymentName:                llamacloudprod.String("azure_openai_deployment_name"),
				AzureOpenAIEndpoint:                      llamacloudprod.String("azure_openai_endpoint"),
				AzureOpenAIKey:                           llamacloudprod.String("azure_openai_key"),
				BboxBottom:                               llamacloudprod.Float(0),
				BboxLeft:                                 llamacloudprod.Float(0),
				BboxRight:                                llamacloudprod.Float(0),
				BboxTop:                                  llamacloudprod.Float(0),
				BoundingBox:                              llamacloudprod.String("bounding_box"),
				CompactMarkdownTable:                     llamacloudprod.Bool(true),
				ComplementalFormattingInstruction:        llamacloudprod.String("complemental_formatting_instruction"),
				ContentGuidelineInstruction:              llamacloudprod.String("content_guideline_instruction"),
				ContinuousMode:                           llamacloudprod.Bool(true),
				DisableImageExtraction:                   llamacloudprod.Bool(true),
				DisableOcr:                               llamacloudprod.Bool(true),
				DisableReconstruction:                    llamacloudprod.Bool(true),
				DoNotCache:                               llamacloudprod.Bool(true),
				DoNotUnrollColumns:                       llamacloudprod.Bool(true),
				EnableCostOptimizer:                      llamacloudprod.Bool(true),
				ExtractCharts:                            llamacloudprod.Bool(true),
				ExtractLayout:                            llamacloudprod.Bool(true),
				ExtractPrintedPageNumber:                 llamacloudprod.Bool(true),
				FastMode:                                 llamacloudprod.Bool(true),
				FormattingInstruction:                    llamacloudprod.String("formatting_instruction"),
				Gpt4oAPIKey:                              llamacloudprod.String("gpt4o_api_key"),
				Gpt4oMode:                                llamacloudprod.Bool(true),
				GuessXlsxSheetName:                       llamacloudprod.Bool(true),
				HideFooters:                              llamacloudprod.Bool(true),
				HideHeaders:                              llamacloudprod.Bool(true),
				HighResOcr:                               llamacloudprod.Bool(true),
				HTMLMakeAllElementsVisible:               llamacloudprod.Bool(true),
				HTMLRemoveFixedElements:                  llamacloudprod.Bool(true),
				HTMLRemoveNavigationElements:             llamacloudprod.Bool(true),
				HTTPProxy:                                llamacloudprod.String("http_proxy"),
				IgnoreDocumentElementsForLayoutDetection: llamacloudprod.Bool(true),
				ImagesToSave:                             []string{"screenshot"},
				InlineImagesInMarkdown:                   llamacloudprod.Bool(true),
				InputS3Path:                              llamacloudprod.String("input_s3_path"),
				InputS3Region:                            llamacloudprod.String("input_s3_region"),
				InputURL:                                 llamacloudprod.String("input_url"),
				InternalIsScreenshotJob:                  llamacloudprod.Bool(true),
				InvalidateCache:                          llamacloudprod.Bool(true),
				IsFormattingInstruction:                  llamacloudprod.Bool(true),
				JobTimeoutExtraTimePerPageInSeconds:      llamacloudprod.Float(0),
				JobTimeoutInSeconds:                      llamacloudprod.Float(0),
				KeepPageSeparatorWhenMergingTables:       llamacloudprod.Bool(true),
				Languages:                                []llamacloudprod.ParsingLanguages{llamacloudprod.ParsingLanguagesAf},
				LayoutAware:                              llamacloudprod.Bool(true),
				LineLevelBoundingBox:                     llamacloudprod.Bool(true),
				MarkdownTableMultilineHeaderSeparator:    llamacloudprod.String("markdown_table_multiline_header_separator"),
				MaxPages:                                 llamacloudprod.Int(0),
				MaxPagesEnforced:                         llamacloudprod.Int(0),
				MergeTablesAcrossPagesInMarkdown:         llamacloudprod.Bool(true),
				Model:                                    llamacloudprod.String("model"),
				OutlinedTableExtraction:                  llamacloudprod.Bool(true),
				OutputPdfOfDocument:                      llamacloudprod.Bool(true),
				OutputS3PathPrefix:                       llamacloudprod.String("output_s3_path_prefix"),
				OutputS3Region:                           llamacloudprod.String("output_s3_region"),
				OutputTablesAsHTML:                       llamacloudprod.Bool(true),
				PageErrorTolerance:                       llamacloudprod.Float(0),
				PageFooterPrefix:                         llamacloudprod.String("page_footer_prefix"),
				PageFooterSuffix:                         llamacloudprod.String("page_footer_suffix"),
				PageHeaderPrefix:                         llamacloudprod.String("page_header_prefix"),
				PageHeaderSuffix:                         llamacloudprod.String("page_header_suffix"),
				PagePrefix:                               llamacloudprod.String("page_prefix"),
				PageSeparator:                            llamacloudprod.String("page_separator"),
				PageSuffix:                               llamacloudprod.String("page_suffix"),
				ParseMode:                                llamacloudprod.ParsingModeParsePageWithoutLlm,
				ParsingInstruction:                       llamacloudprod.String("parsing_instruction"),
				PreciseBoundingBox:                       llamacloudprod.Bool(true),
				PremiumMode:                              llamacloudprod.Bool(true),
				PresentationOutOfBoundsContent:           llamacloudprod.Bool(true),
				PresentationSkipEmbeddedData:             llamacloudprod.Bool(true),
				PreserveLayoutAlignmentAcrossPages:       llamacloudprod.Bool(true),
				PreserveVerySmallText:                    llamacloudprod.Bool(true),
				Preset:                                   llamacloudprod.String("preset"),
				Priority:                                 llamacloudprod.LlamaParseParametersPriorityLow,
				ProjectID:                                llamacloudprod.String("project_id"),
				RemoveHiddenText:                         llamacloudprod.Bool(true),
				ReplaceFailedPageMode:                    llamacloudprod.FailPageModeRawText,
				ReplaceFailedPageWithErrorMessagePrefix:  llamacloudprod.String("replace_failed_page_with_error_message_prefix"),
				ReplaceFailedPageWithErrorMessageSuffix:  llamacloudprod.String("replace_failed_page_with_error_message_suffix"),
				SaveImages:                               llamacloudprod.Bool(true),
				SkipDiagonalText:                         llamacloudprod.Bool(true),
				SpecializedChartParsingAgentic:           llamacloudprod.Bool(true),
				SpecializedChartParsingEfficient:         llamacloudprod.Bool(true),
				SpecializedChartParsingPlus:              llamacloudprod.Bool(true),
				SpecializedImageParsing:                  llamacloudprod.Bool(true),
				SpreadsheetExtractSubTables:              llamacloudprod.Bool(true),
				SpreadsheetForceFormulaComputation:       llamacloudprod.Bool(true),
				SpreadsheetIncludeHiddenSheets:           llamacloudprod.Bool(true),
				StrictModeBuggyFont:                      llamacloudprod.Bool(true),
				StrictModeImageExtraction:                llamacloudprod.Bool(true),
				StrictModeImageOcr:                       llamacloudprod.Bool(true),
				StrictModeReconstruction:                 llamacloudprod.Bool(true),
				StructuredOutput:                         llamacloudprod.Bool(true),
				StructuredOutputJsonSchema:               llamacloudprod.String("structured_output_json_schema"),
				StructuredOutputJsonSchemaName:           llamacloudprod.String("structured_output_json_schema_name"),
				SystemPrompt:                             llamacloudprod.String("system_prompt"),
				SystemPromptAppend:                       llamacloudprod.String("system_prompt_append"),
				TakeScreenshot:                           llamacloudprod.Bool(true),
				TargetPages:                              llamacloudprod.String("target_pages"),
				Tier:                                     llamacloudprod.String("tier"),
				UseVendorMultimodalModel:                 llamacloudprod.Bool(true),
				UserPrompt:                               llamacloudprod.String("user_prompt"),
				VendorMultimodalAPIKey:                   llamacloudprod.String("vendor_multimodal_api_key"),
				VendorMultimodalModelName:                llamacloudprod.String("vendor_multimodal_model_name"),
				Version:                                  llamacloudprod.String("version"),
				WebhookConfigurations: []llamacloudprod.LlamaParseParametersWebhookConfiguration{{
					WebhookEvents: []string{"parse.success", "parse.error"},
					WebhookHeaders: map[string]string{
						"Authorization": "Bearer sk-...",
					},
					WebhookOutputFormat: llamacloudprod.String("json"),
					WebhookURL:          llamacloudprod.String("https://example.com/webhooks/llamacloud"),
				}},
				WebhookURL: llamacloudprod.String("webhook_url"),
			},
			ManagedPipelineID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			MetadataConfig: llamacloudprod.PipelineMetadataConfigParam{
				ExcludedEmbedMetadataKeys: []string{"string"},
				ExcludedLlmMetadataKeys:   []string{"string"},
			},
			Name: llamacloudprod.String("name"),
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
			SparseModelConfig: llamacloudprod.SparseModelConfigParam{
				ClassName: llamacloudprod.String("class_name"),
				ModelType: llamacloudprod.SparseModelConfigModelTypeSplade,
			},
			Status: llamacloudprod.String("status"),
			TransformConfig: llamacloudprod.PipelineUpdateParamsTransformConfigUnion{
				OfAutoTransformConfig: &llamacloudprod.AutoTransformConfigParam{
					ChunkOverlap: llamacloudprod.Int(0),
					ChunkSize:    llamacloudprod.Int(1),
					Mode:         llamacloudprod.AutoTransformConfigModeAuto,
				},
			},
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

func TestPipelineListWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.List(context.TODO(), llamacloudprod.PipelineListParams{
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PipelineName:   llamacloudprod.String("pipeline_name"),
		PipelineType:   llamacloudprod.PipelineTypePlayground,
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectName:    llamacloudprod.String("project_name"),
	})
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineDelete(t *testing.T) {
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
	err := client.Pipelines.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineGet(t *testing.T) {
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
	_, err := client.Pipelines.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *llamacloudprod.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineGetStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.GetStatus(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloudprod.PipelineGetStatusParams{
			FullDetails: llamacloudprod.Bool(true),
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

func TestPipelineUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Upsert(context.TODO(), llamacloudprod.PipelineUpsertParams{
		PipelineCreate: llamacloudprod.PipelineCreateParam{
			Name: "x",
			DataSink: llamacloudprod.DataSinkCreateParam{
				Component: llamacloudprod.DataSinkCreateComponentUnionParam{
					OfAnyMap: map[string]any{
						"foo": "bar",
					},
				},
				Name:     "name",
				SinkType: llamacloudprod.DataSinkCreateSinkTypePinecone,
			},
			DataSinkID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			EmbeddingConfig: llamacloudprod.PipelineCreateEmbeddingConfigUnionParam{
				OfAzureEmbedding: &llamacloudprod.AzureOpenAIEmbeddingConfigParam{
					Component: llamacloudprod.AzureOpenAIEmbeddingParam{
						AdditionalKwargs: map[string]any{
							"foo": "bar",
						},
						APIBase:         llamacloudprod.String("api_base"),
						APIKey:          llamacloudprod.String("api_key"),
						APIVersion:      llamacloudprod.String("api_version"),
						AzureDeployment: llamacloudprod.String("azure_deployment"),
						AzureEndpoint:   llamacloudprod.String("azure_endpoint"),
						ClassName:       llamacloudprod.String("class_name"),
						DefaultHeaders: map[string]string{
							"foo": "string",
						},
						Dimensions:     llamacloudprod.Int(0),
						EmbedBatchSize: llamacloudprod.Int(1),
						MaxRetries:     llamacloudprod.Int(0),
						ModelName:      llamacloudprod.String("model_name"),
						NumWorkers:     llamacloudprod.Int(0),
						ReuseClient:    llamacloudprod.Bool(true),
						Timeout:        llamacloudprod.Float(0),
					},
					Type: llamacloudprod.AzureOpenAIEmbeddingConfigTypeAzureEmbedding,
				},
			},
			EmbeddingModelConfigID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LlamaParseParameters: llamacloudprod.LlamaParseParameters{
				AdaptiveLongTable:                        llamacloudprod.Bool(true),
				AggressiveTableExtraction:                llamacloudprod.Bool(true),
				AnnotateLinks:                            llamacloudprod.Bool(true),
				AutoMode:                                 llamacloudprod.Bool(true),
				AutoModeConfigurationJson:                llamacloudprod.String("auto_mode_configuration_json"),
				AutoModeTriggerOnImageInPage:             llamacloudprod.Bool(true),
				AutoModeTriggerOnRegexpInPage:            llamacloudprod.String("auto_mode_trigger_on_regexp_in_page"),
				AutoModeTriggerOnTableInPage:             llamacloudprod.Bool(true),
				AutoModeTriggerOnTextInPage:              llamacloudprod.String("auto_mode_trigger_on_text_in_page"),
				AzureOpenAIAPIVersion:                    llamacloudprod.String("azure_openai_api_version"),
				AzureOpenAIDeploymentName:                llamacloudprod.String("azure_openai_deployment_name"),
				AzureOpenAIEndpoint:                      llamacloudprod.String("azure_openai_endpoint"),
				AzureOpenAIKey:                           llamacloudprod.String("azure_openai_key"),
				BboxBottom:                               llamacloudprod.Float(0),
				BboxLeft:                                 llamacloudprod.Float(0),
				BboxRight:                                llamacloudprod.Float(0),
				BboxTop:                                  llamacloudprod.Float(0),
				BoundingBox:                              llamacloudprod.String("bounding_box"),
				CompactMarkdownTable:                     llamacloudprod.Bool(true),
				ComplementalFormattingInstruction:        llamacloudprod.String("complemental_formatting_instruction"),
				ContentGuidelineInstruction:              llamacloudprod.String("content_guideline_instruction"),
				ContinuousMode:                           llamacloudprod.Bool(true),
				DisableImageExtraction:                   llamacloudprod.Bool(true),
				DisableOcr:                               llamacloudprod.Bool(true),
				DisableReconstruction:                    llamacloudprod.Bool(true),
				DoNotCache:                               llamacloudprod.Bool(true),
				DoNotUnrollColumns:                       llamacloudprod.Bool(true),
				EnableCostOptimizer:                      llamacloudprod.Bool(true),
				ExtractCharts:                            llamacloudprod.Bool(true),
				ExtractLayout:                            llamacloudprod.Bool(true),
				ExtractPrintedPageNumber:                 llamacloudprod.Bool(true),
				FastMode:                                 llamacloudprod.Bool(true),
				FormattingInstruction:                    llamacloudprod.String("formatting_instruction"),
				Gpt4oAPIKey:                              llamacloudprod.String("gpt4o_api_key"),
				Gpt4oMode:                                llamacloudprod.Bool(true),
				GuessXlsxSheetName:                       llamacloudprod.Bool(true),
				HideFooters:                              llamacloudprod.Bool(true),
				HideHeaders:                              llamacloudprod.Bool(true),
				HighResOcr:                               llamacloudprod.Bool(true),
				HTMLMakeAllElementsVisible:               llamacloudprod.Bool(true),
				HTMLRemoveFixedElements:                  llamacloudprod.Bool(true),
				HTMLRemoveNavigationElements:             llamacloudprod.Bool(true),
				HTTPProxy:                                llamacloudprod.String("http_proxy"),
				IgnoreDocumentElementsForLayoutDetection: llamacloudprod.Bool(true),
				ImagesToSave:                             []string{"screenshot"},
				InlineImagesInMarkdown:                   llamacloudprod.Bool(true),
				InputS3Path:                              llamacloudprod.String("input_s3_path"),
				InputS3Region:                            llamacloudprod.String("input_s3_region"),
				InputURL:                                 llamacloudprod.String("input_url"),
				InternalIsScreenshotJob:                  llamacloudprod.Bool(true),
				InvalidateCache:                          llamacloudprod.Bool(true),
				IsFormattingInstruction:                  llamacloudprod.Bool(true),
				JobTimeoutExtraTimePerPageInSeconds:      llamacloudprod.Float(0),
				JobTimeoutInSeconds:                      llamacloudprod.Float(0),
				KeepPageSeparatorWhenMergingTables:       llamacloudprod.Bool(true),
				Languages:                                []llamacloudprod.ParsingLanguages{llamacloudprod.ParsingLanguagesAf},
				LayoutAware:                              llamacloudprod.Bool(true),
				LineLevelBoundingBox:                     llamacloudprod.Bool(true),
				MarkdownTableMultilineHeaderSeparator:    llamacloudprod.String("markdown_table_multiline_header_separator"),
				MaxPages:                                 llamacloudprod.Int(0),
				MaxPagesEnforced:                         llamacloudprod.Int(0),
				MergeTablesAcrossPagesInMarkdown:         llamacloudprod.Bool(true),
				Model:                                    llamacloudprod.String("model"),
				OutlinedTableExtraction:                  llamacloudprod.Bool(true),
				OutputPdfOfDocument:                      llamacloudprod.Bool(true),
				OutputS3PathPrefix:                       llamacloudprod.String("output_s3_path_prefix"),
				OutputS3Region:                           llamacloudprod.String("output_s3_region"),
				OutputTablesAsHTML:                       llamacloudprod.Bool(true),
				PageErrorTolerance:                       llamacloudprod.Float(0),
				PageFooterPrefix:                         llamacloudprod.String("page_footer_prefix"),
				PageFooterSuffix:                         llamacloudprod.String("page_footer_suffix"),
				PageHeaderPrefix:                         llamacloudprod.String("page_header_prefix"),
				PageHeaderSuffix:                         llamacloudprod.String("page_header_suffix"),
				PagePrefix:                               llamacloudprod.String("page_prefix"),
				PageSeparator:                            llamacloudprod.String("page_separator"),
				PageSuffix:                               llamacloudprod.String("page_suffix"),
				ParseMode:                                llamacloudprod.ParsingModeParsePageWithoutLlm,
				ParsingInstruction:                       llamacloudprod.String("parsing_instruction"),
				PreciseBoundingBox:                       llamacloudprod.Bool(true),
				PremiumMode:                              llamacloudprod.Bool(true),
				PresentationOutOfBoundsContent:           llamacloudprod.Bool(true),
				PresentationSkipEmbeddedData:             llamacloudprod.Bool(true),
				PreserveLayoutAlignmentAcrossPages:       llamacloudprod.Bool(true),
				PreserveVerySmallText:                    llamacloudprod.Bool(true),
				Preset:                                   llamacloudprod.String("preset"),
				Priority:                                 llamacloudprod.LlamaParseParametersPriorityLow,
				ProjectID:                                llamacloudprod.String("project_id"),
				RemoveHiddenText:                         llamacloudprod.Bool(true),
				ReplaceFailedPageMode:                    llamacloudprod.FailPageModeRawText,
				ReplaceFailedPageWithErrorMessagePrefix:  llamacloudprod.String("replace_failed_page_with_error_message_prefix"),
				ReplaceFailedPageWithErrorMessageSuffix:  llamacloudprod.String("replace_failed_page_with_error_message_suffix"),
				SaveImages:                               llamacloudprod.Bool(true),
				SkipDiagonalText:                         llamacloudprod.Bool(true),
				SpecializedChartParsingAgentic:           llamacloudprod.Bool(true),
				SpecializedChartParsingEfficient:         llamacloudprod.Bool(true),
				SpecializedChartParsingPlus:              llamacloudprod.Bool(true),
				SpecializedImageParsing:                  llamacloudprod.Bool(true),
				SpreadsheetExtractSubTables:              llamacloudprod.Bool(true),
				SpreadsheetForceFormulaComputation:       llamacloudprod.Bool(true),
				SpreadsheetIncludeHiddenSheets:           llamacloudprod.Bool(true),
				StrictModeBuggyFont:                      llamacloudprod.Bool(true),
				StrictModeImageExtraction:                llamacloudprod.Bool(true),
				StrictModeImageOcr:                       llamacloudprod.Bool(true),
				StrictModeReconstruction:                 llamacloudprod.Bool(true),
				StructuredOutput:                         llamacloudprod.Bool(true),
				StructuredOutputJsonSchema:               llamacloudprod.String("structured_output_json_schema"),
				StructuredOutputJsonSchemaName:           llamacloudprod.String("structured_output_json_schema_name"),
				SystemPrompt:                             llamacloudprod.String("system_prompt"),
				SystemPromptAppend:                       llamacloudprod.String("system_prompt_append"),
				TakeScreenshot:                           llamacloudprod.Bool(true),
				TargetPages:                              llamacloudprod.String("target_pages"),
				Tier:                                     llamacloudprod.String("tier"),
				UseVendorMultimodalModel:                 llamacloudprod.Bool(true),
				UserPrompt:                               llamacloudprod.String("user_prompt"),
				VendorMultimodalAPIKey:                   llamacloudprod.String("vendor_multimodal_api_key"),
				VendorMultimodalModelName:                llamacloudprod.String("vendor_multimodal_model_name"),
				Version:                                  llamacloudprod.String("version"),
				WebhookConfigurations: []llamacloudprod.LlamaParseParametersWebhookConfiguration{{
					WebhookEvents: []string{"parse.success", "parse.error"},
					WebhookHeaders: map[string]string{
						"Authorization": "Bearer sk-...",
					},
					WebhookOutputFormat: llamacloudprod.String("json"),
					WebhookURL:          llamacloudprod.String("https://example.com/webhooks/llamacloud"),
				}},
				WebhookURL: llamacloudprod.String("webhook_url"),
			},
			ManagedPipelineID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			MetadataConfig: llamacloudprod.PipelineMetadataConfigParam{
				ExcludedEmbedMetadataKeys: []string{"string"},
				ExcludedLlmMetadataKeys:   []string{"string"},
			},
			PipelineType: llamacloudprod.PipelineTypePlayground,
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
			SparseModelConfig: llamacloudprod.SparseModelConfigParam{
				ClassName: llamacloudprod.String("class_name"),
				ModelType: llamacloudprod.SparseModelConfigModelTypeSplade,
			},
			Status: llamacloudprod.String("status"),
			TransformConfig: llamacloudprod.PipelineCreateTransformConfigUnionParam{
				OfAutoTransformConfig: &llamacloudprod.AutoTransformConfigParam{
					ChunkOverlap: llamacloudprod.Int(0),
					ChunkSize:    llamacloudprod.Int(1),
					Mode:         llamacloudprod.AutoTransformConfigModeAuto,
				},
			},
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
