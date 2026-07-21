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

func TestBetaBatchNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Batch.New(context.TODO(), llamacloud.BetaBatchNewParams{
		JobConfig: llamacloud.BetaBatchNewParamsJobConfigUnion{
			OfBatchParseJobRecordCreate: &llamacloud.BetaBatchNewParamsJobConfigBatchParseJobRecordCreate{
				CorrelationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				JobName:       "parse_raw_file_job",
				Parameters: llamacloud.BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParameters{
					AdaptiveLongTable:                 llamacloud.Bool(true),
					AggressiveTableExtraction:         llamacloud.Bool(true),
					AnnotateLinks:                     llamacloud.Bool(true),
					AutoMode:                          llamacloud.Bool(true),
					AutoModeConfigurationJson:         llamacloud.String("auto_mode_configuration_json"),
					AutoModeTriggerOnImageInPage:      llamacloud.Bool(true),
					AutoModeTriggerOnRegexpInPage:     llamacloud.String("auto_mode_trigger_on_regexp_in_page"),
					AutoModeTriggerOnTableInPage:      llamacloud.Bool(true),
					AutoModeTriggerOnTextInPage:       llamacloud.String("auto_mode_trigger_on_text_in_page"),
					AzureOpenAIAPIVersion:             llamacloud.String("azure_openai_api_version"),
					AzureOpenAIDeploymentName:         llamacloud.String("azure_openai_deployment_name"),
					AzureOpenAIEndpoint:               llamacloud.String("azure_openai_endpoint"),
					AzureOpenAIKey:                    llamacloud.String("azure_openai_key"),
					BboxBottom:                        llamacloud.Float(0),
					BboxLeft:                          llamacloud.Float(0),
					BboxRight:                         llamacloud.Float(0),
					BboxTop:                           llamacloud.Float(0),
					BoundingBox:                       llamacloud.String("bounding_box"),
					CompactMarkdownTable:              llamacloud.Bool(true),
					ComplementalFormattingInstruction: llamacloud.String("complemental_formatting_instruction"),
					ConfidenceScoreEffort:             llamacloud.String("confidence_score_effort"),
					ContentGuidelineInstruction:       llamacloud.String("content_guideline_instruction"),
					ContinuousMode:                    llamacloud.Bool(true),
					CustomMetadata: map[string]any{
						"foo": "bar",
					},
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
					Lang:                                     llamacloud.String("lang"),
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
					OutputBucket:                             llamacloud.String("outputBucket"),
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
					PipelineID:                               llamacloud.String("pipeline_id"),
					PreciseBoundingBox:                       llamacloud.Bool(true),
					PremiumMode:                              llamacloud.Bool(true),
					PresentationOutOfBoundsContent:           llamacloud.Bool(true),
					PresentationSkipEmbeddedData:             llamacloud.Bool(true),
					PreserveLayoutAlignmentAcrossPages:       llamacloud.Bool(true),
					PreserveVerySmallText:                    llamacloud.Bool(true),
					Preset:                                   llamacloud.String("preset"),
					Priority:                                 "critical",
					ProjectID:                                llamacloud.String("project_id"),
					RemoveHiddenText:                         llamacloud.Bool(true),
					ReplaceFailedPageMode:                    llamacloud.FailPageModeBlankPage,
					ReplaceFailedPageWithErrorMessagePrefix:  llamacloud.String("replace_failed_page_with_error_message_prefix"),
					ReplaceFailedPageWithErrorMessageSuffix:  llamacloud.String("replace_failed_page_with_error_message_suffix"),
					ResourceInfo: map[string]any{
						"foo": "bar",
					},
					SaveImages:                         llamacloud.Bool(true),
					SkipDiagonalText:                   llamacloud.Bool(true),
					SpecializedChartParsingAgentic:     llamacloud.Bool(true),
					SpecializedChartParsingEfficient:   llamacloud.Bool(true),
					SpecializedChartParsingPlus:        llamacloud.Bool(true),
					SpecializedImageParsing:            llamacloud.Bool(true),
					SpreadsheetExtractSubTables:        llamacloud.Bool(true),
					SpreadsheetForceFormulaComputation: llamacloud.Bool(true),
					SpreadsheetIncludeHiddenSheets:     llamacloud.Bool(true),
					StrictModeBuggyFont:                llamacloud.Bool(true),
					StrictModeImageExtraction:          llamacloud.Bool(true),
					StrictModeImageOcr:                 llamacloud.Bool(true),
					StrictModeReconstruction:           llamacloud.Bool(true),
					StructuredOutput:                   llamacloud.Bool(true),
					StructuredOutputJsonSchema:         llamacloud.String("structured_output_json_schema"),
					StructuredOutputJsonSchemaName:     llamacloud.String("structured_output_json_schema_name"),
					SystemPrompt:                       llamacloud.String("system_prompt"),
					SystemPromptAppend:                 llamacloud.String("system_prompt_append"),
					TakeScreenshot:                     llamacloud.Bool(true),
					TargetPages:                        llamacloud.String("target_pages"),
					Tier:                               llamacloud.String("tier"),
					Type:                               "parse",
					UseVendorMultimodalModel:           llamacloud.Bool(true),
					UserPrompt:                         llamacloud.String("user_prompt"),
					VendorMultimodalAPIKey:             llamacloud.String("vendor_multimodal_api_key"),
					VendorMultimodalModelName:          llamacloud.String("vendor_multimodal_model_name"),
					Version:                            llamacloud.String("version"),
					WebhookConfigurations: []llamacloud.BetaBatchNewParamsJobConfigBatchParseJobRecordCreateParametersWebhookConfiguration{{
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
				ParentJobExecutionID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Partitions: map[string]string{
					"foo": "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				},
				ProjectID:  llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				SessionID:  llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				UserID:     llamacloud.String("user_id"),
				WebhookURL: llamacloud.String("webhook_url"),
			},
		},
		OrganizationID:         llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:              llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ContinueAsNewThreshold: llamacloud.Int(0),
		DirectoryID:            llamacloud.String("dir-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
		ItemIDs:                []string{"dfl-aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", "dfl-11111111-2222-3333-4444-555555555555"},
		PageSize:               llamacloud.Int(1),
		TemporalNamespace:      llamacloud.String("temporal-namespace"),
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaBatchListWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Batch.List(context.TODO(), llamacloud.BetaBatchListParams{
		DirectoryID:    llamacloud.String("directory_id"),
		JobType:        llamacloud.BetaBatchListParamsJobTypeClassify,
		Limit:          llamacloud.Int(1),
		Offset:         llamacloud.Int(0),
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:         llamacloud.BetaBatchListParamsStatusCancelled,
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBetaBatchCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Batch.Cancel(
		context.TODO(),
		"job_id",
		llamacloud.BetaBatchCancelParams{
			OrganizationID:    llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ProjectID:         llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Reason:            llamacloud.String("reason"),
			TemporalNamespace: llamacloud.String("temporal-namespace"),
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

func TestBetaBatchGetStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Beta.Batch.GetStatus(
		context.TODO(),
		"job_id",
		llamacloud.BetaBatchGetStatusParams{
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
