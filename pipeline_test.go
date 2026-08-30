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

func TestPipelineNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.New(context.TODO(), llamacloud.PipelineNewParams{
		PipelineCreate: llamacloud.PipelineCreateParam{
			Name: "x",
			DataSink: llamacloud.DataSinkCreateParam{
				Component: llamacloud.DataSinkCreateComponentUnionParam{
					OfAnyMap: map[string]any{
						"foo": "bar",
					},
				},
				Name:     "name",
				SinkType: llamacloud.DataSinkCreateSinkTypeAstraDB,
			},
			DataSinkID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			EmbeddingConfig: llamacloud.PipelineCreateEmbeddingConfigUnionParam{
				OfAzureEmbedding: &llamacloud.AzureOpenAIEmbeddingConfigParam{
					Component: llamacloud.AzureOpenAIEmbeddingParam{
						AdditionalKwargs: map[string]any{
							"foo": "bar",
						},
						APIBase:         llamacloud.String("api_base"),
						APIKey:          llamacloud.String("api_key"),
						APIVersion:      llamacloud.String("api_version"),
						AzureDeployment: llamacloud.String("azure_deployment"),
						AzureEndpoint:   llamacloud.String("azure_endpoint"),
						ClassName:       llamacloud.String("class_name"),
						DefaultHeaders: map[string]string{
							"foo": "string",
						},
						Dimensions:     llamacloud.Int(0),
						EmbedBatchSize: llamacloud.Int(1),
						MaxRetries:     llamacloud.Int(0),
						ModelName:      llamacloud.String("model_name"),
						NumWorkers:     llamacloud.Int(0),
						ReuseClient:    llamacloud.Bool(true),
						Timeout:        llamacloud.Float(0),
					},
					Type: llamacloud.AzureOpenAIEmbeddingConfigTypeAzureEmbedding,
				},
			},
			EmbeddingModelConfigID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LlamaParseParameters: llamacloud.LlamaParseParameters{
				AdaptiveLongTable:                        llamacloud.Bool(true),
				AggressiveTableExtraction:                llamacloud.Bool(true),
				AnnotateLineNumbers:                      llamacloud.Bool(true),
				AnnotateLinks:                            llamacloud.Bool(true),
				AnnotateRevisions:                        llamacloud.Bool(true),
				AutoMode:                                 llamacloud.Bool(true),
				AutoModeConfigurationJson:                llamacloud.String("auto_mode_configuration_json"),
				AutoModeTriggerOnImageInPage:             llamacloud.Bool(true),
				AutoModeTriggerOnRegexpInPage:            llamacloud.String("auto_mode_trigger_on_regexp_in_page"),
				AutoModeTriggerOnTableInPage:             llamacloud.Bool(true),
				AutoModeTriggerOnTextInPage:              llamacloud.String("auto_mode_trigger_on_text_in_page"),
				AzureOpenAIAPIVersion:                    llamacloud.String("azure_openai_api_version"),
				AzureOpenAIDeploymentName:                llamacloud.String("azure_openai_deployment_name"),
				AzureOpenAIEndpoint:                      llamacloud.String("azure_openai_endpoint"),
				AzureOpenAIKey:                           llamacloud.String("azure_openai_key"),
				BboxBottom:                               llamacloud.Float(0),
				BboxLeft:                                 llamacloud.Float(0),
				BboxRight:                                llamacloud.Float(0),
				BboxTop:                                  llamacloud.Float(0),
				BoundingBox:                              llamacloud.String("bounding_box"),
				CompactMarkdownTable:                     llamacloud.Bool(true),
				ComplementalFormattingInstruction:        llamacloud.String("complemental_formatting_instruction"),
				ConfidenceScoreEffort:                    llamacloud.String("confidence_score_effort"),
				ContentGuidelineInstruction:              llamacloud.String("content_guideline_instruction"),
				ContinuousMode:                           llamacloud.Bool(true),
				DisableImageExtraction:                   llamacloud.Bool(true),
				DisableOcr:                               llamacloud.Bool(true),
				DisableReconstruction:                    llamacloud.Bool(true),
				DoNotCache:                               llamacloud.Bool(true),
				DoNotUnrollColumns:                       llamacloud.Bool(true),
				EnableCostOptimizer:                      llamacloud.Bool(true),
				ExtractCharts:                            llamacloud.Bool(true),
				ExtractLayout:                            llamacloud.Bool(true),
				ExtractPrintedPageNumber:                 llamacloud.Bool(true),
				FastMode:                                 llamacloud.Bool(true),
				FormattingInstruction:                    llamacloud.String("formatting_instruction"),
				Gpt4oAPIKey:                              llamacloud.String("gpt4o_api_key"),
				Gpt4oMode:                                llamacloud.Bool(true),
				GuessXlsxSheetName:                       llamacloud.Bool(true),
				HideFooters:                              llamacloud.Bool(true),
				HideHeaders:                              llamacloud.Bool(true),
				HighResOcr:                               llamacloud.Bool(true),
				HTMLMakeAllElementsVisible:               llamacloud.Bool(true),
				HTMLRemoveFixedElements:                  llamacloud.Bool(true),
				HTMLRemoveNavigationElements:             llamacloud.Bool(true),
				HTTPProxy:                                llamacloud.String("http_proxy"),
				IgnoreDocumentElementsForLayoutDetection: llamacloud.Bool(true),
				ImagesToSave:                             []string{"embedded"},
				InlineImagesInMarkdown:                   llamacloud.Bool(true),
				InputS3Path:                              llamacloud.String("input_s3_path"),
				InputS3Region:                            llamacloud.String("input_s3_region"),
				InputURL:                                 llamacloud.String("input_url"),
				InternalIsScreenshotJob:                  llamacloud.Bool(true),
				InvalidateCache:                          llamacloud.Bool(true),
				IsFormattingInstruction:                  llamacloud.Bool(true),
				JobTimeoutExtraTimePerPageInSeconds:      llamacloud.Float(0),
				JobTimeoutInSeconds:                      llamacloud.Float(0),
				KeepPageSeparatorWhenMergingTables:       llamacloud.Bool(true),
				Languages:                                []llamacloud.ParsingLanguages{llamacloud.ParsingLanguagesAbq},
				LayoutAware:                              llamacloud.Bool(true),
				LineLevelBoundingBox:                     llamacloud.Bool(true),
				MarkdownTableMultilineHeaderSeparator:    llamacloud.String("markdown_table_multiline_header_separator"),
				MaxPages:                                 llamacloud.Int(0),
				MaxPagesEnforced:                         llamacloud.Int(0),
				MergeTablesAcrossPagesInMarkdown:         llamacloud.Bool(true),
				Model:                                    llamacloud.String("model"),
				OutlinedTableExtraction:                  llamacloud.Bool(true),
				OutputPdfOfDocument:                      llamacloud.Bool(true),
				OutputS3PathPrefix:                       llamacloud.String("output_s3_path_prefix"),
				OutputS3Region:                           llamacloud.String("output_s3_region"),
				OutputTablesAsHTML:                       llamacloud.Bool(true),
				PageErrorTolerance:                       llamacloud.Float(0),
				PageFooterPrefix:                         llamacloud.String("page_footer_prefix"),
				PageFooterSuffix:                         llamacloud.String("page_footer_suffix"),
				PageHeaderPrefix:                         llamacloud.String("page_header_prefix"),
				PageHeaderSuffix:                         llamacloud.String("page_header_suffix"),
				PagePrefix:                               llamacloud.String("page_prefix"),
				PageSeparator:                            llamacloud.String("page_separator"),
				PageSuffix:                               llamacloud.String("page_suffix"),
				ParseMode:                                llamacloud.ParsingModeParseDocumentWithAgent,
				ParsingInstruction:                       llamacloud.String("parsing_instruction"),
				PreciseBoundingBox:                       llamacloud.Bool(true),
				PremiumMode:                              llamacloud.Bool(true),
				PresentationOutOfBoundsContent:           llamacloud.Bool(true),
				PresentationSkipEmbeddedData:             llamacloud.Bool(true),
				PreserveLayoutAlignmentAcrossPages:       llamacloud.Bool(true),
				PreserveVerySmallText:                    llamacloud.Bool(true),
				Preset:                                   llamacloud.String("preset"),
				Priority:                                 llamacloud.LlamaParseParametersPriorityCritical,
				ProjectID:                                llamacloud.String("project_id"),
				RemoveHiddenText:                         llamacloud.Bool(true),
				ReplaceFailedPageMode:                    llamacloud.FailPageModeBlankPage,
				ReplaceFailedPageWithErrorMessagePrefix:  llamacloud.String("replace_failed_page_with_error_message_prefix"),
				ReplaceFailedPageWithErrorMessageSuffix:  llamacloud.String("replace_failed_page_with_error_message_suffix"),
				SaveImages:                               llamacloud.Bool(true),
				SkipDiagonalText:                         llamacloud.Bool(true),
				SpecializedChartParsingAgentic:           llamacloud.Bool(true),
				SpecializedChartParsingEfficient:         llamacloud.Bool(true),
				SpecializedChartParsingPlus:              llamacloud.Bool(true),
				SpecializedImageParsing:                  llamacloud.Bool(true),
				SpreadsheetExtractSubTables:              llamacloud.Bool(true),
				SpreadsheetForceFormulaComputation:       llamacloud.Bool(true),
				SpreadsheetIncludeHiddenSheets:           llamacloud.Bool(true),
				StrictModeBuggyFont:                      llamacloud.Bool(true),
				StrictModeImageExtraction:                llamacloud.Bool(true),
				StrictModeImageOcr:                       llamacloud.Bool(true),
				StrictModeReconstruction:                 llamacloud.Bool(true),
				StructuredOutput:                         llamacloud.Bool(true),
				StructuredOutputJsonSchema:               llamacloud.String("structured_output_json_schema"),
				StructuredOutputJsonSchemaName:           llamacloud.String("structured_output_json_schema_name"),
				SystemPrompt:                             llamacloud.String("system_prompt"),
				SystemPromptAppend:                       llamacloud.String("system_prompt_append"),
				TakeScreenshot:                           llamacloud.Bool(true),
				TargetPages:                              llamacloud.String("target_pages"),
				Tier:                                     llamacloud.String("tier"),
				UseVendorMultimodalModel:                 llamacloud.Bool(true),
				UserPrompt:                               llamacloud.String("user_prompt"),
				VendorMultimodalAPIKey:                   llamacloud.String("vendor_multimodal_api_key"),
				VendorMultimodalModelName:                llamacloud.String("vendor_multimodal_model_name"),
				Version:                                  llamacloud.String("version"),
				WebhookConfigurations: []llamacloud.LlamaParseParametersWebhookConfiguration{{
					WebhookEvents: []string{"parse.success", "parse.error"},
					WebhookHeaders: map[string]string{
						"Authorization": "Bearer sk-...",
					},
					WebhookOutputFormat:  llamacloud.String("json"),
					WebhookSigningSecret: llamacloud.String("whsec_..."),
					WebhookURL:           llamacloud.String("https://example.com/webhooks/llamacloud"),
				}},
				WebhookURL: llamacloud.String("webhook_url"),
			},
			ManagedPipelineID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			MetadataConfig: llamacloud.PipelineMetadataConfigParam{
				ExcludedEmbedMetadataKeys: []string{"string"},
				ExcludedLlmMetadataKeys:   []string{"string"},
			},
			PipelineType: llamacloud.PipelineTypeManaged,
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
			SparseModelConfig: llamacloud.SparseModelConfigParam{
				ClassName: llamacloud.String("class_name"),
				ModelType: llamacloud.SparseModelConfigModelTypeAuto,
			},
			Status: llamacloud.String("status"),
			TransformConfig: llamacloud.PipelineCreateTransformConfigUnionParam{
				OfAutoTransformConfig: &llamacloud.AutoTransformConfigParam{
					ChunkOverlap: llamacloud.Int(0),
					ChunkSize:    llamacloud.Int(1),
					Mode:         llamacloud.AutoTransformConfigModeAuto,
				},
			},
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

func TestPipelineUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloud.PipelineUpdateParams{
			ProjectID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			DataSink: llamacloud.DataSinkCreateParam{
				Component: llamacloud.DataSinkCreateComponentUnionParam{
					OfAnyMap: map[string]any{
						"foo": "bar",
					},
				},
				Name:     "name",
				SinkType: llamacloud.DataSinkCreateSinkTypeAstraDB,
			},
			DataSinkID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			EmbeddingConfig: llamacloud.PipelineUpdateParamsEmbeddingConfigUnion{
				OfAzureEmbedding: &llamacloud.AzureOpenAIEmbeddingConfigParam{
					Component: llamacloud.AzureOpenAIEmbeddingParam{
						AdditionalKwargs: map[string]any{
							"foo": "bar",
						},
						APIBase:         llamacloud.String("api_base"),
						APIKey:          llamacloud.String("api_key"),
						APIVersion:      llamacloud.String("api_version"),
						AzureDeployment: llamacloud.String("azure_deployment"),
						AzureEndpoint:   llamacloud.String("azure_endpoint"),
						ClassName:       llamacloud.String("class_name"),
						DefaultHeaders: map[string]string{
							"foo": "string",
						},
						Dimensions:     llamacloud.Int(0),
						EmbedBatchSize: llamacloud.Int(1),
						MaxRetries:     llamacloud.Int(0),
						ModelName:      llamacloud.String("model_name"),
						NumWorkers:     llamacloud.Int(0),
						ReuseClient:    llamacloud.Bool(true),
						Timeout:        llamacloud.Float(0),
					},
					Type: llamacloud.AzureOpenAIEmbeddingConfigTypeAzureEmbedding,
				},
			},
			EmbeddingModelConfigID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LlamaParseParameters: llamacloud.LlamaParseParameters{
				AdaptiveLongTable:                        llamacloud.Bool(true),
				AggressiveTableExtraction:                llamacloud.Bool(true),
				AnnotateLineNumbers:                      llamacloud.Bool(true),
				AnnotateLinks:                            llamacloud.Bool(true),
				AnnotateRevisions:                        llamacloud.Bool(true),
				AutoMode:                                 llamacloud.Bool(true),
				AutoModeConfigurationJson:                llamacloud.String("auto_mode_configuration_json"),
				AutoModeTriggerOnImageInPage:             llamacloud.Bool(true),
				AutoModeTriggerOnRegexpInPage:            llamacloud.String("auto_mode_trigger_on_regexp_in_page"),
				AutoModeTriggerOnTableInPage:             llamacloud.Bool(true),
				AutoModeTriggerOnTextInPage:              llamacloud.String("auto_mode_trigger_on_text_in_page"),
				AzureOpenAIAPIVersion:                    llamacloud.String("azure_openai_api_version"),
				AzureOpenAIDeploymentName:                llamacloud.String("azure_openai_deployment_name"),
				AzureOpenAIEndpoint:                      llamacloud.String("azure_openai_endpoint"),
				AzureOpenAIKey:                           llamacloud.String("azure_openai_key"),
				BboxBottom:                               llamacloud.Float(0),
				BboxLeft:                                 llamacloud.Float(0),
				BboxRight:                                llamacloud.Float(0),
				BboxTop:                                  llamacloud.Float(0),
				BoundingBox:                              llamacloud.String("bounding_box"),
				CompactMarkdownTable:                     llamacloud.Bool(true),
				ComplementalFormattingInstruction:        llamacloud.String("complemental_formatting_instruction"),
				ConfidenceScoreEffort:                    llamacloud.String("confidence_score_effort"),
				ContentGuidelineInstruction:              llamacloud.String("content_guideline_instruction"),
				ContinuousMode:                           llamacloud.Bool(true),
				DisableImageExtraction:                   llamacloud.Bool(true),
				DisableOcr:                               llamacloud.Bool(true),
				DisableReconstruction:                    llamacloud.Bool(true),
				DoNotCache:                               llamacloud.Bool(true),
				DoNotUnrollColumns:                       llamacloud.Bool(true),
				EnableCostOptimizer:                      llamacloud.Bool(true),
				ExtractCharts:                            llamacloud.Bool(true),
				ExtractLayout:                            llamacloud.Bool(true),
				ExtractPrintedPageNumber:                 llamacloud.Bool(true),
				FastMode:                                 llamacloud.Bool(true),
				FormattingInstruction:                    llamacloud.String("formatting_instruction"),
				Gpt4oAPIKey:                              llamacloud.String("gpt4o_api_key"),
				Gpt4oMode:                                llamacloud.Bool(true),
				GuessXlsxSheetName:                       llamacloud.Bool(true),
				HideFooters:                              llamacloud.Bool(true),
				HideHeaders:                              llamacloud.Bool(true),
				HighResOcr:                               llamacloud.Bool(true),
				HTMLMakeAllElementsVisible:               llamacloud.Bool(true),
				HTMLRemoveFixedElements:                  llamacloud.Bool(true),
				HTMLRemoveNavigationElements:             llamacloud.Bool(true),
				HTTPProxy:                                llamacloud.String("http_proxy"),
				IgnoreDocumentElementsForLayoutDetection: llamacloud.Bool(true),
				ImagesToSave:                             []string{"embedded"},
				InlineImagesInMarkdown:                   llamacloud.Bool(true),
				InputS3Path:                              llamacloud.String("input_s3_path"),
				InputS3Region:                            llamacloud.String("input_s3_region"),
				InputURL:                                 llamacloud.String("input_url"),
				InternalIsScreenshotJob:                  llamacloud.Bool(true),
				InvalidateCache:                          llamacloud.Bool(true),
				IsFormattingInstruction:                  llamacloud.Bool(true),
				JobTimeoutExtraTimePerPageInSeconds:      llamacloud.Float(0),
				JobTimeoutInSeconds:                      llamacloud.Float(0),
				KeepPageSeparatorWhenMergingTables:       llamacloud.Bool(true),
				Languages:                                []llamacloud.ParsingLanguages{llamacloud.ParsingLanguagesAbq},
				LayoutAware:                              llamacloud.Bool(true),
				LineLevelBoundingBox:                     llamacloud.Bool(true),
				MarkdownTableMultilineHeaderSeparator:    llamacloud.String("markdown_table_multiline_header_separator"),
				MaxPages:                                 llamacloud.Int(0),
				MaxPagesEnforced:                         llamacloud.Int(0),
				MergeTablesAcrossPagesInMarkdown:         llamacloud.Bool(true),
				Model:                                    llamacloud.String("model"),
				OutlinedTableExtraction:                  llamacloud.Bool(true),
				OutputPdfOfDocument:                      llamacloud.Bool(true),
				OutputS3PathPrefix:                       llamacloud.String("output_s3_path_prefix"),
				OutputS3Region:                           llamacloud.String("output_s3_region"),
				OutputTablesAsHTML:                       llamacloud.Bool(true),
				PageErrorTolerance:                       llamacloud.Float(0),
				PageFooterPrefix:                         llamacloud.String("page_footer_prefix"),
				PageFooterSuffix:                         llamacloud.String("page_footer_suffix"),
				PageHeaderPrefix:                         llamacloud.String("page_header_prefix"),
				PageHeaderSuffix:                         llamacloud.String("page_header_suffix"),
				PagePrefix:                               llamacloud.String("page_prefix"),
				PageSeparator:                            llamacloud.String("page_separator"),
				PageSuffix:                               llamacloud.String("page_suffix"),
				ParseMode:                                llamacloud.ParsingModeParseDocumentWithAgent,
				ParsingInstruction:                       llamacloud.String("parsing_instruction"),
				PreciseBoundingBox:                       llamacloud.Bool(true),
				PremiumMode:                              llamacloud.Bool(true),
				PresentationOutOfBoundsContent:           llamacloud.Bool(true),
				PresentationSkipEmbeddedData:             llamacloud.Bool(true),
				PreserveLayoutAlignmentAcrossPages:       llamacloud.Bool(true),
				PreserveVerySmallText:                    llamacloud.Bool(true),
				Preset:                                   llamacloud.String("preset"),
				Priority:                                 llamacloud.LlamaParseParametersPriorityCritical,
				ProjectID:                                llamacloud.String("project_id"),
				RemoveHiddenText:                         llamacloud.Bool(true),
				ReplaceFailedPageMode:                    llamacloud.FailPageModeBlankPage,
				ReplaceFailedPageWithErrorMessagePrefix:  llamacloud.String("replace_failed_page_with_error_message_prefix"),
				ReplaceFailedPageWithErrorMessageSuffix:  llamacloud.String("replace_failed_page_with_error_message_suffix"),
				SaveImages:                               llamacloud.Bool(true),
				SkipDiagonalText:                         llamacloud.Bool(true),
				SpecializedChartParsingAgentic:           llamacloud.Bool(true),
				SpecializedChartParsingEfficient:         llamacloud.Bool(true),
				SpecializedChartParsingPlus:              llamacloud.Bool(true),
				SpecializedImageParsing:                  llamacloud.Bool(true),
				SpreadsheetExtractSubTables:              llamacloud.Bool(true),
				SpreadsheetForceFormulaComputation:       llamacloud.Bool(true),
				SpreadsheetIncludeHiddenSheets:           llamacloud.Bool(true),
				StrictModeBuggyFont:                      llamacloud.Bool(true),
				StrictModeImageExtraction:                llamacloud.Bool(true),
				StrictModeImageOcr:                       llamacloud.Bool(true),
				StrictModeReconstruction:                 llamacloud.Bool(true),
				StructuredOutput:                         llamacloud.Bool(true),
				StructuredOutputJsonSchema:               llamacloud.String("structured_output_json_schema"),
				StructuredOutputJsonSchemaName:           llamacloud.String("structured_output_json_schema_name"),
				SystemPrompt:                             llamacloud.String("system_prompt"),
				SystemPromptAppend:                       llamacloud.String("system_prompt_append"),
				TakeScreenshot:                           llamacloud.Bool(true),
				TargetPages:                              llamacloud.String("target_pages"),
				Tier:                                     llamacloud.String("tier"),
				UseVendorMultimodalModel:                 llamacloud.Bool(true),
				UserPrompt:                               llamacloud.String("user_prompt"),
				VendorMultimodalAPIKey:                   llamacloud.String("vendor_multimodal_api_key"),
				VendorMultimodalModelName:                llamacloud.String("vendor_multimodal_model_name"),
				Version:                                  llamacloud.String("version"),
				WebhookConfigurations: []llamacloud.LlamaParseParametersWebhookConfiguration{{
					WebhookEvents: []string{"parse.success", "parse.error"},
					WebhookHeaders: map[string]string{
						"Authorization": "Bearer sk-...",
					},
					WebhookOutputFormat:  llamacloud.String("json"),
					WebhookSigningSecret: llamacloud.String("whsec_..."),
					WebhookURL:           llamacloud.String("https://example.com/webhooks/llamacloud"),
				}},
				WebhookURL: llamacloud.String("webhook_url"),
			},
			ManagedPipelineID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			MetadataConfig: llamacloud.PipelineMetadataConfigParam{
				ExcludedEmbedMetadataKeys: []string{"string"},
				ExcludedLlmMetadataKeys:   []string{"string"},
			},
			Name: llamacloud.String("name"),
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
			SparseModelConfig: llamacloud.SparseModelConfigParam{
				ClassName: llamacloud.String("class_name"),
				ModelType: llamacloud.SparseModelConfigModelTypeAuto,
			},
			Status: llamacloud.String("status"),
			TransformConfig: llamacloud.PipelineUpdateParamsTransformConfigUnion{
				OfAutoTransformConfig: &llamacloud.AutoTransformConfigParam{
					ChunkOverlap: llamacloud.Int(0),
					ChunkSize:    llamacloud.Int(1),
					Mode:         llamacloud.AutoTransformConfigModeAuto,
				},
			},
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

func TestPipelineListWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.List(context.TODO(), llamacloud.PipelineListParams{
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PipelineName:   llamacloud.String("pipeline_name"),
		PipelineType:   llamacloud.PipelineTypeManaged,
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectName:    llamacloud.String("project_name"),
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPipelineDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Pipelines.Delete(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloud.PipelineDeleteParams{
			ProjectID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestPipelineGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloud.PipelineGetParams{
			ProjectID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestPipelineGetStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.GetStatus(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloud.PipelineGetStatusParams{
			FullDetails: llamacloud.Bool(true),
			ProjectID:   llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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

func TestPipelineRunSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.RunSearch(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		llamacloud.PipelineRunSearchParams{
			Query:                       "x",
			OrganizationID:              llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:                   llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
			SearchFiltersInferenceSchema: map[string]*llamacloud.PipelineRunSearchParamsSearchFiltersInferenceSchemaUnion{
				"foo": {
					OfAnyMap: map[string]any{
						"foo": "bar",
					},
				},
			},
			SparseSimilarityTopK: llamacloud.Int(1),
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

func TestPipelineUpsertWithOptionalParams(t *testing.T) {
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
	_, err := client.Pipelines.Upsert(context.TODO(), llamacloud.PipelineUpsertParams{
		PipelineCreate: llamacloud.PipelineCreateParam{
			Name: "x",
			DataSink: llamacloud.DataSinkCreateParam{
				Component: llamacloud.DataSinkCreateComponentUnionParam{
					OfAnyMap: map[string]any{
						"foo": "bar",
					},
				},
				Name:     "name",
				SinkType: llamacloud.DataSinkCreateSinkTypeAstraDB,
			},
			DataSinkID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			EmbeddingConfig: llamacloud.PipelineCreateEmbeddingConfigUnionParam{
				OfAzureEmbedding: &llamacloud.AzureOpenAIEmbeddingConfigParam{
					Component: llamacloud.AzureOpenAIEmbeddingParam{
						AdditionalKwargs: map[string]any{
							"foo": "bar",
						},
						APIBase:         llamacloud.String("api_base"),
						APIKey:          llamacloud.String("api_key"),
						APIVersion:      llamacloud.String("api_version"),
						AzureDeployment: llamacloud.String("azure_deployment"),
						AzureEndpoint:   llamacloud.String("azure_endpoint"),
						ClassName:       llamacloud.String("class_name"),
						DefaultHeaders: map[string]string{
							"foo": "string",
						},
						Dimensions:     llamacloud.Int(0),
						EmbedBatchSize: llamacloud.Int(1),
						MaxRetries:     llamacloud.Int(0),
						ModelName:      llamacloud.String("model_name"),
						NumWorkers:     llamacloud.Int(0),
						ReuseClient:    llamacloud.Bool(true),
						Timeout:        llamacloud.Float(0),
					},
					Type: llamacloud.AzureOpenAIEmbeddingConfigTypeAzureEmbedding,
				},
			},
			EmbeddingModelConfigID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			LlamaParseParameters: llamacloud.LlamaParseParameters{
				AdaptiveLongTable:                        llamacloud.Bool(true),
				AggressiveTableExtraction:                llamacloud.Bool(true),
				AnnotateLineNumbers:                      llamacloud.Bool(true),
				AnnotateLinks:                            llamacloud.Bool(true),
				AnnotateRevisions:                        llamacloud.Bool(true),
				AutoMode:                                 llamacloud.Bool(true),
				AutoModeConfigurationJson:                llamacloud.String("auto_mode_configuration_json"),
				AutoModeTriggerOnImageInPage:             llamacloud.Bool(true),
				AutoModeTriggerOnRegexpInPage:            llamacloud.String("auto_mode_trigger_on_regexp_in_page"),
				AutoModeTriggerOnTableInPage:             llamacloud.Bool(true),
				AutoModeTriggerOnTextInPage:              llamacloud.String("auto_mode_trigger_on_text_in_page"),
				AzureOpenAIAPIVersion:                    llamacloud.String("azure_openai_api_version"),
				AzureOpenAIDeploymentName:                llamacloud.String("azure_openai_deployment_name"),
				AzureOpenAIEndpoint:                      llamacloud.String("azure_openai_endpoint"),
				AzureOpenAIKey:                           llamacloud.String("azure_openai_key"),
				BboxBottom:                               llamacloud.Float(0),
				BboxLeft:                                 llamacloud.Float(0),
				BboxRight:                                llamacloud.Float(0),
				BboxTop:                                  llamacloud.Float(0),
				BoundingBox:                              llamacloud.String("bounding_box"),
				CompactMarkdownTable:                     llamacloud.Bool(true),
				ComplementalFormattingInstruction:        llamacloud.String("complemental_formatting_instruction"),
				ConfidenceScoreEffort:                    llamacloud.String("confidence_score_effort"),
				ContentGuidelineInstruction:              llamacloud.String("content_guideline_instruction"),
				ContinuousMode:                           llamacloud.Bool(true),
				DisableImageExtraction:                   llamacloud.Bool(true),
				DisableOcr:                               llamacloud.Bool(true),
				DisableReconstruction:                    llamacloud.Bool(true),
				DoNotCache:                               llamacloud.Bool(true),
				DoNotUnrollColumns:                       llamacloud.Bool(true),
				EnableCostOptimizer:                      llamacloud.Bool(true),
				ExtractCharts:                            llamacloud.Bool(true),
				ExtractLayout:                            llamacloud.Bool(true),
				ExtractPrintedPageNumber:                 llamacloud.Bool(true),
				FastMode:                                 llamacloud.Bool(true),
				FormattingInstruction:                    llamacloud.String("formatting_instruction"),
				Gpt4oAPIKey:                              llamacloud.String("gpt4o_api_key"),
				Gpt4oMode:                                llamacloud.Bool(true),
				GuessXlsxSheetName:                       llamacloud.Bool(true),
				HideFooters:                              llamacloud.Bool(true),
				HideHeaders:                              llamacloud.Bool(true),
				HighResOcr:                               llamacloud.Bool(true),
				HTMLMakeAllElementsVisible:               llamacloud.Bool(true),
				HTMLRemoveFixedElements:                  llamacloud.Bool(true),
				HTMLRemoveNavigationElements:             llamacloud.Bool(true),
				HTTPProxy:                                llamacloud.String("http_proxy"),
				IgnoreDocumentElementsForLayoutDetection: llamacloud.Bool(true),
				ImagesToSave:                             []string{"embedded"},
				InlineImagesInMarkdown:                   llamacloud.Bool(true),
				InputS3Path:                              llamacloud.String("input_s3_path"),
				InputS3Region:                            llamacloud.String("input_s3_region"),
				InputURL:                                 llamacloud.String("input_url"),
				InternalIsScreenshotJob:                  llamacloud.Bool(true),
				InvalidateCache:                          llamacloud.Bool(true),
				IsFormattingInstruction:                  llamacloud.Bool(true),
				JobTimeoutExtraTimePerPageInSeconds:      llamacloud.Float(0),
				JobTimeoutInSeconds:                      llamacloud.Float(0),
				KeepPageSeparatorWhenMergingTables:       llamacloud.Bool(true),
				Languages:                                []llamacloud.ParsingLanguages{llamacloud.ParsingLanguagesAbq},
				LayoutAware:                              llamacloud.Bool(true),
				LineLevelBoundingBox:                     llamacloud.Bool(true),
				MarkdownTableMultilineHeaderSeparator:    llamacloud.String("markdown_table_multiline_header_separator"),
				MaxPages:                                 llamacloud.Int(0),
				MaxPagesEnforced:                         llamacloud.Int(0),
				MergeTablesAcrossPagesInMarkdown:         llamacloud.Bool(true),
				Model:                                    llamacloud.String("model"),
				OutlinedTableExtraction:                  llamacloud.Bool(true),
				OutputPdfOfDocument:                      llamacloud.Bool(true),
				OutputS3PathPrefix:                       llamacloud.String("output_s3_path_prefix"),
				OutputS3Region:                           llamacloud.String("output_s3_region"),
				OutputTablesAsHTML:                       llamacloud.Bool(true),
				PageErrorTolerance:                       llamacloud.Float(0),
				PageFooterPrefix:                         llamacloud.String("page_footer_prefix"),
				PageFooterSuffix:                         llamacloud.String("page_footer_suffix"),
				PageHeaderPrefix:                         llamacloud.String("page_header_prefix"),
				PageHeaderSuffix:                         llamacloud.String("page_header_suffix"),
				PagePrefix:                               llamacloud.String("page_prefix"),
				PageSeparator:                            llamacloud.String("page_separator"),
				PageSuffix:                               llamacloud.String("page_suffix"),
				ParseMode:                                llamacloud.ParsingModeParseDocumentWithAgent,
				ParsingInstruction:                       llamacloud.String("parsing_instruction"),
				PreciseBoundingBox:                       llamacloud.Bool(true),
				PremiumMode:                              llamacloud.Bool(true),
				PresentationOutOfBoundsContent:           llamacloud.Bool(true),
				PresentationSkipEmbeddedData:             llamacloud.Bool(true),
				PreserveLayoutAlignmentAcrossPages:       llamacloud.Bool(true),
				PreserveVerySmallText:                    llamacloud.Bool(true),
				Preset:                                   llamacloud.String("preset"),
				Priority:                                 llamacloud.LlamaParseParametersPriorityCritical,
				ProjectID:                                llamacloud.String("project_id"),
				RemoveHiddenText:                         llamacloud.Bool(true),
				ReplaceFailedPageMode:                    llamacloud.FailPageModeBlankPage,
				ReplaceFailedPageWithErrorMessagePrefix:  llamacloud.String("replace_failed_page_with_error_message_prefix"),
				ReplaceFailedPageWithErrorMessageSuffix:  llamacloud.String("replace_failed_page_with_error_message_suffix"),
				SaveImages:                               llamacloud.Bool(true),
				SkipDiagonalText:                         llamacloud.Bool(true),
				SpecializedChartParsingAgentic:           llamacloud.Bool(true),
				SpecializedChartParsingEfficient:         llamacloud.Bool(true),
				SpecializedChartParsingPlus:              llamacloud.Bool(true),
				SpecializedImageParsing:                  llamacloud.Bool(true),
				SpreadsheetExtractSubTables:              llamacloud.Bool(true),
				SpreadsheetForceFormulaComputation:       llamacloud.Bool(true),
				SpreadsheetIncludeHiddenSheets:           llamacloud.Bool(true),
				StrictModeBuggyFont:                      llamacloud.Bool(true),
				StrictModeImageExtraction:                llamacloud.Bool(true),
				StrictModeImageOcr:                       llamacloud.Bool(true),
				StrictModeReconstruction:                 llamacloud.Bool(true),
				StructuredOutput:                         llamacloud.Bool(true),
				StructuredOutputJsonSchema:               llamacloud.String("structured_output_json_schema"),
				StructuredOutputJsonSchemaName:           llamacloud.String("structured_output_json_schema_name"),
				SystemPrompt:                             llamacloud.String("system_prompt"),
				SystemPromptAppend:                       llamacloud.String("system_prompt_append"),
				TakeScreenshot:                           llamacloud.Bool(true),
				TargetPages:                              llamacloud.String("target_pages"),
				Tier:                                     llamacloud.String("tier"),
				UseVendorMultimodalModel:                 llamacloud.Bool(true),
				UserPrompt:                               llamacloud.String("user_prompt"),
				VendorMultimodalAPIKey:                   llamacloud.String("vendor_multimodal_api_key"),
				VendorMultimodalModelName:                llamacloud.String("vendor_multimodal_model_name"),
				Version:                                  llamacloud.String("version"),
				WebhookConfigurations: []llamacloud.LlamaParseParametersWebhookConfiguration{{
					WebhookEvents: []string{"parse.success", "parse.error"},
					WebhookHeaders: map[string]string{
						"Authorization": "Bearer sk-...",
					},
					WebhookOutputFormat:  llamacloud.String("json"),
					WebhookSigningSecret: llamacloud.String("whsec_..."),
					WebhookURL:           llamacloud.String("https://example.com/webhooks/llamacloud"),
				}},
				WebhookURL: llamacloud.String("webhook_url"),
			},
			ManagedPipelineID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			MetadataConfig: llamacloud.PipelineMetadataConfigParam{
				ExcludedEmbedMetadataKeys: []string{"string"},
				ExcludedLlmMetadataKeys:   []string{"string"},
			},
			PipelineType: llamacloud.PipelineTypeManaged,
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
			SparseModelConfig: llamacloud.SparseModelConfigParam{
				ClassName: llamacloud.String("class_name"),
				ModelType: llamacloud.SparseModelConfigModelTypeAuto,
			},
			Status: llamacloud.String("status"),
			TransformConfig: llamacloud.PipelineCreateTransformConfigUnionParam{
				OfAutoTransformConfig: &llamacloud.AutoTransformConfigParam{
					ChunkOverlap: llamacloud.Int(0),
					ChunkSize:    llamacloud.Int(1),
					Mode:         llamacloud.AutoTransformConfigModeAuto,
				},
			},
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
