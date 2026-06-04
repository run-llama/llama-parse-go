// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/llamacloud-prod-go"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/testutil"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
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
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Parsing.New(context.TODO(), llamacloudprod.ParsingNewParams{
		Tier:           llamacloudprod.ParsingNewParamsTierFast,
		Version:        llamacloudprod.ParsingNewParamsVersionLatest,
		OrganizationID: llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ProjectID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		AgenticOptions: llamacloudprod.ParsingNewParamsAgenticOptions{
			CustomPrompt: llamacloudprod.String("custom_prompt"),
		},
		ClientName: llamacloudprod.String("client_name"),
		CropBox: llamacloudprod.ParsingNewParamsCropBox{
			Bottom: llamacloudprod.Float(0),
			Left:   llamacloudprod.Float(0),
			Right:  llamacloudprod.Float(0),
			Top:    llamacloudprod.Float(0),
		},
		DisableCache: llamacloudprod.Bool(true),
		FastOptions:  map[string]any{},
		FileID:       llamacloudprod.String("file_id"),
		HTTPProxy:    llamacloudprod.String("https:"),
		InputOptions: llamacloudprod.ParsingNewParamsInputOptions{
			HTML: llamacloudprod.ParsingNewParamsInputOptionsHTML{
				MakeAllElementsVisible:   llamacloudprod.Bool(true),
				RemoveFixedElements:      llamacloudprod.Bool(true),
				RemoveNavigationElements: llamacloudprod.Bool(true),
			},
			Pdf: map[string]any{},
			Presentation: llamacloudprod.ParsingNewParamsInputOptionsPresentation{
				OutOfBoundsContent: llamacloudprod.Bool(true),
				SkipEmbeddedData:   llamacloudprod.Bool(true),
			},
			Spreadsheet: llamacloudprod.ParsingNewParamsInputOptionsSpreadsheet{
				DetectSubTablesInSheets:         llamacloudprod.Bool(true),
				ForceFormulaComputationInSheets: llamacloudprod.Bool(true),
				IncludeHiddenSheets:             llamacloudprod.Bool(true),
			},
		},
		OutputOptions: llamacloudprod.ParsingNewParamsOutputOptions{
			AdditionalOutputs:        []string{"stripped_md", "concatenated_stripped_txt", "word_bbox"},
			ExtractPrintedPageNumber: llamacloudprod.Bool(true),
			ImagesToSave:             []string{"screenshot"},
			Markdown: llamacloudprod.ParsingNewParamsOutputOptionsMarkdown{
				AnnotateLinks: llamacloudprod.Bool(true),
				InlineImages:  llamacloudprod.Bool(true),
				Tables: llamacloudprod.ParsingNewParamsOutputOptionsMarkdownTables{
					CompactMarkdownTables:           llamacloudprod.Bool(true),
					MarkdownTableMultilineSeparator: llamacloudprod.String("markdown_table_multiline_separator"),
					MergeContinuedTables:            llamacloudprod.Bool(true),
					OutputTablesAsMarkdown:          llamacloudprod.Bool(true),
				},
			},
			SpatialText: llamacloudprod.ParsingNewParamsOutputOptionsSpatialText{
				DoNotUnrollColumns:                 llamacloudprod.Bool(true),
				PreserveLayoutAlignmentAcrossPages: llamacloudprod.Bool(true),
				PreserveVerySmallText:              llamacloudprod.Bool(true),
			},
			TablesAsSpreadsheet: llamacloudprod.ParsingNewParamsOutputOptionsTablesAsSpreadsheet{
				Enable:         llamacloudprod.Bool(true),
				GuessSheetName: llamacloudprod.Bool(true),
			},
		},
		PageRanges: llamacloudprod.ParsingNewParamsPageRanges{
			MaxPages:    llamacloudprod.Int(1),
			TargetPages: llamacloudprod.String("target_pages"),
		},
		ProcessingControl: llamacloudprod.ParsingNewParamsProcessingControl{
			JobFailureConditions: llamacloudprod.ParsingNewParamsProcessingControlJobFailureConditions{
				AllowedPageFailureRatio:           llamacloudprod.Float(1),
				FailOnBuggyFont:                   llamacloudprod.Bool(true),
				FailOnImageExtractionError:        llamacloudprod.Bool(true),
				FailOnImageOcrError:               llamacloudprod.Bool(true),
				FailOnMarkdownReconstructionError: llamacloudprod.Bool(true),
			},
			Timeouts: llamacloudprod.ParsingNewParamsProcessingControlTimeouts{
				BaseInSeconds:             llamacloudprod.Int(1),
				ExtraTimePerPageInSeconds: llamacloudprod.Int(1),
			},
		},
		ProcessingOptions: llamacloudprod.ParsingNewParamsProcessingOptions{
			AggressiveTableExtraction: llamacloudprod.Bool(true),
			AutoModeConfiguration: []llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfiguration{{
				ParsingConf: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConf{
					AdaptiveLongTable:         llamacloudprod.Bool(true),
					AggressiveTableExtraction: llamacloudprod.Bool(true),
					CropBox: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfCropBox{
						Bottom: llamacloudprod.Float(0),
						Left:   llamacloudprod.Float(0),
						Right:  llamacloudprod.Float(0),
						Top:    llamacloudprod.Float(0),
					},
					CustomPrompt:  llamacloudprod.String("custom_prompt"),
					ExtractLayout: llamacloudprod.Bool(true),
					HighResOcr:    llamacloudprod.Bool(true),
					Ignore: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfIgnore{
						IgnoreDiagonalText: llamacloudprod.Bool(true),
						IgnoreHiddenText:   llamacloudprod.Bool(true),
					},
					Language:                llamacloudprod.String("language"),
					OutlinedTableExtraction: llamacloudprod.Bool(true),
					Presentation: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfPresentation{
						OutOfBoundsContent: llamacloudprod.Bool(true),
						SkipEmbeddedData:   llamacloudprod.Bool(true),
					},
					SpatialText: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfSpatialText{
						DoNotUnrollColumns:                 llamacloudprod.Bool(true),
						PreserveLayoutAlignmentAcrossPages: llamacloudprod.Bool(true),
						PreserveVerySmallText:              llamacloudprod.Bool(true),
					},
					SpecializedChartParsing: "agentic_plus",
					Tier:                    "agentic",
					Version:                 "latest",
				},
				FilenameMatchGlob:     llamacloudprod.String("*.txt"),
				FilenameMatchGlobList: []string{"string"},
				FilenameRegexp:        llamacloudprod.String("filename_regexp"),
				FilenameRegexpMode:    llamacloudprod.String("filename_regexp_mode"),
				FullPageImageInPage:   llamacloudprod.Bool(true),
				FullPageImageInPageThreshold: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				ImageInPage:         llamacloudprod.Bool(true),
				LayoutElementInPage: llamacloudprod.String("layout_element_in_page"),
				LayoutElementInPageConfidenceThreshold: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnion{
					OfFloat: llamacloudprod.Float(0),
				},
				PageContainsAtLeastNCharts: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtLeastNImages: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtLeastNLayoutElements: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtLeastNLines: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtLeastNLinks: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtLeastNNumbers: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtLeastNPercentNumbers: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtLeastNTables: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtLeastNWords: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtMostNCharts: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtMostNImages: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtMostNLayoutElements: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtMostNLines: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtMostNLinks: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtMostNNumbers: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtMostNPercentNumbers: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtMostNTables: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageContainsAtMostNWords: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageLongerThanNChars: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnion{
					OfInt: llamacloudprod.Int(0),
				},
				PageMdError: llamacloudprod.Bool(true),
				PageShorterThanNChars: llamacloudprod.ParsingNewParamsProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnion{
					OfInt: llamacloudprod.Int(0),
				},
				RegexpInPage:     llamacloudprod.String("regexp_in_page"),
				RegexpInPageMode: llamacloudprod.String("regexp_in_page_mode"),
				TableInPage:      llamacloudprod.Bool(true),
				TextInPage:       llamacloudprod.String("text_in_page"),
				TriggerMode:      llamacloudprod.String("trigger_mode"),
			}},
			CostOptimizer: llamacloudprod.ParsingNewParamsProcessingOptionsCostOptimizer{
				Enable: llamacloudprod.Bool(true),
			},
			DisableHeuristics: llamacloudprod.Bool(true),
			Ignore: llamacloudprod.ParsingNewParamsProcessingOptionsIgnore{
				IgnoreDiagonalText: llamacloudprod.Bool(true),
				IgnoreHiddenText:   llamacloudprod.Bool(true),
				IgnoreTextInImage:  llamacloudprod.Bool(true),
			},
			OcrParameters: llamacloudprod.ParsingNewParamsProcessingOptionsOcrParameters{
				Languages: []llamacloudprod.ParsingLanguages{llamacloudprod.ParsingLanguagesAf},
			},
			SpecializedChartParsing: "agentic_plus",
		},
		SourceURL: llamacloudprod.String("https:"),
		WebhookConfigurations: []llamacloudprod.ParsingNewParamsWebhookConfiguration{{
			WebhookEvents: []string{"parse.success", "parse.error"},
			WebhookHeaders: map[string]any{
				"foo": "bar",
			},
			WebhookOutputFormat: "json",
			WebhookURL:          llamacloudprod.String("https:"),
		}},
	})
	if err != nil {
		var apierr *llamacloudprod.Error
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
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Parsing.List(context.TODO(), llamacloudprod.ParsingListParams{
		CreatedAtOnOrAfter:  llamacloudprod.Time(time.Now()),
		CreatedAtOnOrBefore: llamacloudprod.Time(time.Now()),
		JobIDs:              []string{"string", "string"},
		OrganizationID:      llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		PageSize:            llamacloudprod.Int(0),
		PageToken:           llamacloudprod.String("page_token"),
		ProjectID:           llamacloudprod.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Status:              llamacloudprod.ParsingListParamsStatusPending,
	})
	if err != nil {
		var apierr *llamacloudprod.Error
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
	client := llamacloudprod.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Parsing.Get(
		context.TODO(),
		"job_id",
		llamacloudprod.ParsingGetParams{
			Expand:         []string{"string"},
			ImageFilenames: llamacloudprod.String("image_filenames"),
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
