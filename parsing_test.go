// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloud_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/run-llama/llama-parse-go"
	"github.com/run-llama/llama-parse-go/internal/testutil"
	"github.com/run-llama/llama-parse-go/option"
)

func TestParsingNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Parsing.New(context.TODO(), llamacloud.ParsingNewParams{
		Tier:           llamacloud.ParsingNewParamsTierFast,
		Version:        llamacloud.ParsingNewParamsVersionLatest,
		OrganizationID: llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		AgenticOptions: llamacloud.ParsingNewParamsAgenticOptions{
			CustomPrompt: llamacloud.String("custom_prompt"),
		},
		ClientName:      llamacloud.String("client_name"),
		ConfigurationID: llamacloud.String("configuration_id"),
		CropBox: llamacloud.ParsingNewParamsCropBox{
			Bottom: llamacloud.Float(0),
			Left:   llamacloud.Float(0),
			Right:  llamacloud.Float(0),
			Top:    llamacloud.Float(0),
		},
		DisableCache: llamacloud.Bool(true),
		FastOptions:  map[string]any{},
		FileID:       llamacloud.String("file_id"),
		HTTPProxy:    llamacloud.String("https:"),
		InputOptions: llamacloud.ParsingNewParamsInputOptions{
			HTML: llamacloud.ParsingNewParamsInputOptionsHTML{
				MakeAllElementsVisible:   llamacloud.Bool(true),
				RemoveFixedElements:      llamacloud.Bool(true),
				RemoveNavigationElements: llamacloud.Bool(true),
			},
			Image: llamacloud.ParsingNewParamsInputOptionsImage{
				CameraPhotoCorrection: llamacloud.Bool(true),
			},
			Pdf: map[string]any{},
			Presentation: llamacloud.ParsingNewParamsInputOptionsPresentation{
				OutOfBoundsContent: llamacloud.Bool(true),
				SkipEmbeddedData:   llamacloud.Bool(true),
			},
			Spreadsheet: llamacloud.ParsingNewParamsInputOptionsSpreadsheet{
				DetectSubTablesInSheets:         llamacloud.Bool(true),
				ForceFormulaComputationInSheets: llamacloud.Bool(true),
				IncludeHiddenSheets:             llamacloud.Bool(true),
			},
		},
		OutputOptions: llamacloud.ParsingNewParamsOutputOptions{
			AdditionalOutputs:        []string{"stripped_md", "concatenated_stripped_txt", "word_bbox"},
			ExtractPrintedPageNumber: llamacloud.Bool(true),
			GranularBboxes:           []string{"word", "line", "cell"},
			ImagesToSave:             []string{"embedded"},
			Markdown: llamacloud.ParsingNewParamsOutputOptionsMarkdown{
				AnnotateLinks:     llamacloud.Bool(true),
				AnnotateRevisions: llamacloud.Bool(true),
				InlineImages:      llamacloud.Bool(true),
				Tables: llamacloud.ParsingNewParamsOutputOptionsMarkdownTables{
					CompactMarkdownTables:           llamacloud.Bool(true),
					MarkdownTableMultilineSeparator: llamacloud.String("markdown_table_multiline_separator"),
					MergeContinuedTables:            llamacloud.Bool(true),
					OutputTablesAsMarkdown:          llamacloud.Bool(true),
				},
			},
			SaveOutputPdf: llamacloud.Bool(true),
			SpatialText: llamacloud.ParsingNewParamsOutputOptionsSpatialText{
				DoNotUnrollColumns:                 llamacloud.Bool(true),
				PreserveLayoutAlignmentAcrossPages: llamacloud.Bool(true),
				PreserveVerySmallText:              llamacloud.Bool(true),
			},
			TablesAsSpreadsheet: llamacloud.ParsingNewParamsOutputOptionsTablesAsSpreadsheet{
				Enable:         llamacloud.Bool(true),
				GuessSheetName: llamacloud.Bool(true),
			},
		},
		PageRanges: llamacloud.ParsingNewParamsPageRanges{
			MaxPages:    llamacloud.Int(1),
			TargetPages: llamacloud.String("target_pages"),
		},
		ProcessingControl: llamacloud.ParsingNewParamsProcessingControl{
			JobFailureConditions: llamacloud.ParsingNewParamsProcessingControlJobFailureConditions{
				AllowedPageFailureRatio:           llamacloud.Float(1),
				FailOnBuggyFont:                   llamacloud.Bool(true),
				FailOnImageExtractionError:        llamacloud.Bool(true),
				FailOnImageOcrError:               llamacloud.Bool(true),
				FailOnMarkdownReconstructionError: llamacloud.Bool(true),
			},
			Timeouts: llamacloud.ParsingNewParamsProcessingControlTimeouts{
				BaseInSeconds:             llamacloud.Int(1),
				ExtraTimePerPageInSeconds: llamacloud.Int(1),
			},
		},
		ProcessingOptions: llamacloud.ParsingNewParamsProcessingOptions{
			AggressiveTableExtraction: llamacloud.Bool(true),
			AutoModeConfiguration: []llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfiguration{{
				ParsingConf: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConf{
					AdaptiveLongTable:         llamacloud.Bool(true),
					AggressiveTableExtraction: llamacloud.Bool(true),
					CropBox: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfCropBox{
						Bottom: llamacloud.Float(0),
						Left:   llamacloud.Float(0),
						Right:  llamacloud.Float(0),
						Top:    llamacloud.Float(0),
					},
					CustomPrompt:  llamacloud.String("custom_prompt"),
					ExtractLayout: llamacloud.Bool(true),
					HighResOcr:    llamacloud.Bool(true),
					Ignore: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfIgnore{
						IgnoreDiagonalText: llamacloud.Bool(true),
						IgnoreHiddenText:   llamacloud.Bool(true),
					},
					Language:                llamacloud.String("language"),
					OutlinedTableExtraction: llamacloud.Bool(true),
					Presentation: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfPresentation{
						OutOfBoundsContent: llamacloud.Bool(true),
						SkipEmbeddedData:   llamacloud.Bool(true),
					},
					SpatialText: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfSpatialText{
						DoNotUnrollColumns:                 llamacloud.Bool(true),
						PreserveLayoutAlignmentAcrossPages: llamacloud.Bool(true),
						PreserveVerySmallText:              llamacloud.Bool(true),
					},
					SpecializedChartParsing: "agentic",
					Tier:                    "agentic",
					Version:                 "latest",
				},
				FilenameMatchGlob:     llamacloud.String("*.txt"),
				FilenameMatchGlobList: []string{"string"},
				FilenameRegexp:        llamacloud.String("filename_regexp"),
				FilenameRegexpMode:    llamacloud.String("filename_regexp_mode"),
				FullPageImageInPage:   llamacloud.Bool(true),
				FullPageImageInPageThreshold: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnion{
					OfFloat: llamacloud.Float(0),
				},
				ImageInPage:         llamacloud.Bool(true),
				LayoutElementInPage: llamacloud.String("layout_element_in_page"),
				LayoutElementInPageConfidenceThreshold: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnion{
					OfFloat: llamacloud.Float(0),
				},
				PageContainsAtLeastNCharts: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtLeastNImages: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtLeastNLayoutElements: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtLeastNLines: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtLeastNLinks: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtLeastNNumbers: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtLeastNPercentNumbers: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtLeastNTables: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtLeastNWords: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtMostNCharts: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtMostNImages: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtMostNLayoutElements: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtMostNLines: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtMostNLinks: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtMostNNumbers: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtMostNPercentNumbers: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtMostNTables: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnion{
					OfInt: llamacloud.Int(0),
				},
				PageContainsAtMostNWords: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnion{
					OfInt: llamacloud.Int(0),
				},
				PageLongerThanNChars: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnion{
					OfInt: llamacloud.Int(0),
				},
				PageMdError: llamacloud.Bool(true),
				PageShorterThanNChars: llamacloud.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnion{
					OfInt: llamacloud.Int(0),
				},
				RegexpInPage:     llamacloud.String("regexp_in_page"),
				RegexpInPageMode: llamacloud.String("regexp_in_page_mode"),
				TableInPage:      llamacloud.Bool(true),
				TextInPage:       llamacloud.String("text_in_page"),
				TriggerMode:      llamacloud.String("trigger_mode"),
			}},
			ConfidenceScoreEffort: "high",
			CostOptimizer: llamacloud.ParsingNewParamsProcessingOptionsCostOptimizer{
				Enable: llamacloud.Bool(true),
			},
			DisableHeuristics: llamacloud.Bool(true),
			Forms:             "enrich",
			Ignore: llamacloud.ParsingNewParamsProcessingOptionsIgnore{
				IgnoreDiagonalText: llamacloud.Bool(true),
				IgnoreHiddenText:   llamacloud.Bool(true),
				IgnoreTextInImage:  llamacloud.Bool(true),
			},
			OcrParameters: llamacloud.ParsingNewParamsProcessingOptionsOcrParameters{
				Languages: []llamacloud.ParsingLanguages{llamacloud.ParsingLanguagesAbq},
			},
			SpecializedChartParsing: "agentic",
		},
		SourceURL: llamacloud.String("https:"),
		UserMetadata: map[string]string{
			"owner": "jerry",
			"team":  "research",
		},
		WebhookConfigurationIDs: []string{"whc-...", "whc-..."},
		WebhookConfigurations: []llamacloud.ParsingNewParamsWebhookConfiguration{{
			WebhookEvents: []string{"parse.success", "parse.error"},
			WebhookHeaders: map[string]any{
				"foo": "bar",
			},
			WebhookOutputFormat:  "json",
			WebhookSigningSecret: llamacloud.String("webhook_signing_secret"),
			WebhookURL:           llamacloud.String("https:"),
		}},
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestParsingListWithOptionalParams(t *testing.T) {
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
	_, err := client.Parsing.List(context.TODO(), llamacloud.ParsingListParams{
		CreatedAtOnOrAfter:  llamacloud.Time(time.Now()),
		CreatedAtOnOrBefore: llamacloud.Time(time.Now()),
		JobIDs:              []string{"string", "string"},
		OrganizationID:      llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:            llamacloud.Int(0),
		PageToken:           llamacloud.String("page_token"),
		ProjectID:           llamacloud.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:              llamacloud.ParsingListParamsStatusCancelled,
	})
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestParsingCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Parsing.Cancel(
		context.TODO(),
		"job_id",
		llamacloud.ParsingCancelParams{
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

func TestParsingGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Parsing.Get(
		context.TODO(),
		"job_id",
		llamacloud.ParsingGetParams{
			Expand:         []string{"string"},
			ImageFilenames: llamacloud.String("image_filenames"),
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

func TestParsingListVersions(t *testing.T) {
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
	_, err := client.Parsing.ListVersions(context.TODO())
	if err != nil {
		var apierr *llamacloud.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
