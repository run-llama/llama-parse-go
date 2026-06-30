// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/run-llama/llama-parse-go/internal/apijson"
	"github.com/run-llama/llama-parse-go/internal/apiquery"
	shimjson "github.com/run-llama/llama-parse-go/internal/encoding/json"
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/pagination"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
	"github.com/run-llama/llama-parse-go/shared/constant"
)

// ConfigurationService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigurationService] method instead.
type ConfigurationService struct {
	options []option.RequestOption
}

// NewConfigurationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigurationService(opts ...option.RequestOption) (r ConfigurationService) {
	r = ConfigurationService{}
	r.options = opts
	return
}

// Upsert a product configuration; updates if one with the same name + product
// type + project exists, otherwise creates.
func (r *ConfigurationService) New(ctx context.Context, params ConfigurationNewParams, opts ...option.RequestOption) (res *ConfigurationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/beta/configurations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get a single product configuration by ID.
func (r *ConfigurationService) Get(ctx context.Context, configID string, query ConfigurationGetParams, opts ...option.RequestOption) (res *ConfigurationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/configurations/%s", url.PathEscape(configID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Update an existing product configuration.
func (r *ConfigurationService) Update(ctx context.Context, configID string, params ConfigurationUpdateParams, opts ...option.RequestOption) (res *ConfigurationResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/beta/configurations/%s", url.PathEscape(configID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// List product configurations for the current project.
func (r *ConfigurationService) List(ctx context.Context, query ConfigurationListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[ConfigurationResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v1/beta/configurations"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List product configurations for the current project.
func (r *ConfigurationService) ListAutoPaging(ctx context.Context, query ConfigurationListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[ConfigurationResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Delete a product configuration.
func (r *ConfigurationService) Delete(ctx context.Context, configID string, body ConfigurationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/beta/configurations/%s", url.PathEscape(configID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// Typed parameters for a _classify v2_ product configuration.
type ClassifyV2ParametersResp struct {
	// Product type.
	ProductType constant.ClassifyV2 `json:"product_type" default:"classify_v2"`
	// Classify rules to evaluate against the document (at least one required)
	Rules []ClassifyV2ParametersRuleResp `json:"rules" api:"required"`
	// Classify execution mode
	//
	// Any of "FAST".
	Mode ClassifyV2ParametersMode `json:"mode"`
	// Parsing configuration for classify jobs.
	ParsingConfiguration ClassifyV2ParametersParsingConfigurationResp `json:"parsing_configuration" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductType          respjson.Field
		Rules                respjson.Field
		Mode                 respjson.Field
		ParsingConfiguration respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyV2ParametersResp) RawJSON() string { return r.JSON.raw }
func (r *ClassifyV2ParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ClassifyV2ParametersResp to a ClassifyV2Parameters.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ClassifyV2Parameters.Overrides()
func (r ClassifyV2ParametersResp) ToParam() ClassifyV2Parameters {
	return param.Override[ClassifyV2Parameters](json.RawMessage(r.RawJSON()))
}

// A rule for classifying documents.
type ClassifyV2ParametersRuleResp struct {
	// Natural language criteria for matching this rule
	Description string `json:"description" api:"required"`
	// Document type to assign when rule matches
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyV2ParametersRuleResp) RawJSON() string { return r.JSON.raw }
func (r *ClassifyV2ParametersRuleResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Classify execution mode
type ClassifyV2ParametersMode string

const (
	ClassifyV2ParametersModeFast ClassifyV2ParametersMode = "FAST"
)

// Parsing configuration for classify jobs.
type ClassifyV2ParametersParsingConfigurationResp struct {
	// ISO 639-1 language code for the document
	Lang string `json:"lang"`
	// Maximum number of pages to process. Omit for no limit.
	MaxPages int64 `json:"max_pages" api:"nullable"`
	// Comma-separated page numbers or ranges to process (1-based). Omit to process all
	// pages.
	TargetPages string `json:"target_pages" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lang        respjson.Field
		MaxPages    respjson.Field
		TargetPages respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClassifyV2ParametersParsingConfigurationResp) RawJSON() string { return r.JSON.raw }
func (r *ClassifyV2ParametersParsingConfigurationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed parameters for a _classify v2_ product configuration.
//
// The properties ProductType, Rules are required.
type ClassifyV2Parameters struct {
	// Classify rules to evaluate against the document (at least one required)
	Rules []ClassifyV2ParametersRule `json:"rules,omitzero" api:"required"`
	// Parsing configuration for classify jobs.
	ParsingConfiguration ClassifyV2ParametersParsingConfiguration `json:"parsing_configuration,omitzero"`
	// Classify execution mode
	//
	// Any of "FAST".
	Mode ClassifyV2ParametersMode `json:"mode,omitzero"`
	// Product type.
	//
	// This field can be elided, and will marshal its zero value as "classify_v2".
	ProductType constant.ClassifyV2 `json:"product_type" default:"classify_v2"`
	paramObj
}

func (r ClassifyV2Parameters) MarshalJSON() (data []byte, err error) {
	type shadow ClassifyV2Parameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifyV2Parameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A rule for classifying documents.
//
// The properties Description, Type are required.
type ClassifyV2ParametersRule struct {
	// Natural language criteria for matching this rule
	Description string `json:"description" api:"required"`
	// Document type to assign when rule matches
	Type string `json:"type" api:"required"`
	paramObj
}

func (r ClassifyV2ParametersRule) MarshalJSON() (data []byte, err error) {
	type shadow ClassifyV2ParametersRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifyV2ParametersRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parsing configuration for classify jobs.
type ClassifyV2ParametersParsingConfiguration struct {
	// Maximum number of pages to process. Omit for no limit.
	MaxPages param.Opt[int64] `json:"max_pages,omitzero"`
	// Comma-separated page numbers or ranges to process (1-based). Omit to process all
	// pages.
	TargetPages param.Opt[string] `json:"target_pages,omitzero"`
	// ISO 639-1 language code for the document
	Lang param.Opt[string] `json:"lang,omitzero"`
	paramObj
}

func (r ClassifyV2ParametersParsingConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow ClassifyV2ParametersParsingConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClassifyV2ParametersParsingConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for creating a product configuration.
type ConfigurationCreate struct {
	// Human-readable name for this configuration.
	Name string `json:"name" api:"required"`
	// Product-specific configuration parameters.
	Parameters ConfigurationCreateParametersUnion `json:"parameters" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigurationCreate) RawJSON() string { return r.JSON.raw }
func (r *ConfigurationCreate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConfigurationCreate to a ConfigurationCreateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConfigurationCreateParam.Overrides()
func (r ConfigurationCreate) ToParam() ConfigurationCreateParam {
	return param.Override[ConfigurationCreateParam](json.RawMessage(r.RawJSON()))
}

// ConfigurationCreateParametersUnion contains all possible properties and values
// from [ClassifyV2ParametersResp], [ExtractV2ParametersResp],
// [ParseV2ParametersResp], [SplitV1ParametersResp],
// [ConfigurationCreateParametersSpreadsheetV1], [UntypedParametersResp].
//
// Use the [ConfigurationCreateParametersUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConfigurationCreateParametersUnion struct {
	// Any of "classify_v2", "extract_v2", "parse_v2", "split_v1", "spreadsheet_v1",
	// "unknown".
	ProductType string `json:"product_type"`
	// This field is from variant [ClassifyV2ParametersResp].
	Rules []ClassifyV2ParametersRuleResp `json:"rules"`
	// This field is from variant [ClassifyV2ParametersResp].
	Mode ClassifyV2ParametersMode `json:"mode"`
	// This field is from variant [ClassifyV2ParametersResp].
	ParsingConfiguration ClassifyV2ParametersParsingConfigurationResp `json:"parsing_configuration"`
	// This field is from variant [ExtractV2ParametersResp].
	DataSchema map[string]*ExtractV2ParametersDataSchemaUnionResp `json:"data_schema"`
	// This field is from variant [ExtractV2ParametersResp].
	CiteSources bool `json:"cite_sources"`
	// This field is from variant [ExtractV2ParametersResp].
	ConfidenceScores bool `json:"confidence_scores"`
	// This field is from variant [ExtractV2ParametersResp].
	ExtractionTarget ExtractV2ParametersExtractionTarget `json:"extraction_target"`
	// This field is from variant [ExtractV2ParametersResp].
	MaxPages int64 `json:"max_pages"`
	// This field is from variant [ExtractV2ParametersResp].
	ParseConfigID string `json:"parse_config_id"`
	// This field is from variant [ExtractV2ParametersResp].
	ParseTier string `json:"parse_tier"`
	// This field is from variant [ExtractV2ParametersResp].
	SystemPrompt string `json:"system_prompt"`
	// This field is from variant [ExtractV2ParametersResp].
	TargetPages string `json:"target_pages"`
	Tier        string `json:"tier"`
	Version     string `json:"version"`
	// This field is from variant [ParseV2ParametersResp].
	AgenticOptions ParseV2ParametersAgenticOptionsResp `json:"agentic_options"`
	// This field is from variant [ParseV2ParametersResp].
	ClientName string `json:"client_name"`
	// This field is from variant [ParseV2ParametersResp].
	CropBox ParseV2ParametersCropBoxResp `json:"crop_box"`
	// This field is from variant [ParseV2ParametersResp].
	DisableCache bool `json:"disable_cache"`
	// This field is from variant [ParseV2ParametersResp].
	FastOptions any `json:"fast_options"`
	// This field is from variant [ParseV2ParametersResp].
	InputOptions ParseV2ParametersInputOptionsResp `json:"input_options"`
	// This field is from variant [ParseV2ParametersResp].
	OutputOptions ParseV2ParametersOutputOptionsResp `json:"output_options"`
	// This field is from variant [ParseV2ParametersResp].
	PageRanges ParseV2ParametersPageRangesResp `json:"page_ranges"`
	// This field is from variant [ParseV2ParametersResp].
	ProcessingControl ParseV2ParametersProcessingControlResp `json:"processing_control"`
	// This field is from variant [ParseV2ParametersResp].
	ProcessingOptions ParseV2ParametersProcessingOptionsResp `json:"processing_options"`
	// This field is from variant [ParseV2ParametersResp].
	WebhookConfigurationIDs []string `json:"webhook_configuration_ids"`
	// This field is from variant [ParseV2ParametersResp].
	WebhookConfigurations []ParseV2ParametersWebhookConfigurationResp `json:"webhook_configurations"`
	// This field is from variant [SplitV1ParametersResp].
	Categories []SplitCategory `json:"categories"`
	// This field is from variant [SplitV1ParametersResp].
	SplittingStrategy SplitV1ParametersSplittingStrategyResp `json:"splitting_strategy"`
	// This field is from variant [ConfigurationCreateParametersSpreadsheetV1].
	ExtractionRange string `json:"extraction_range"`
	// This field is from variant [ConfigurationCreateParametersSpreadsheetV1].
	FlattenHierarchicalTables bool `json:"flatten_hierarchical_tables"`
	// This field is from variant [ConfigurationCreateParametersSpreadsheetV1].
	GenerateAdditionalMetadata bool `json:"generate_additional_metadata"`
	// This field is from variant [ConfigurationCreateParametersSpreadsheetV1].
	IncludeHiddenCells bool `json:"include_hidden_cells"`
	// This field is from variant [ConfigurationCreateParametersSpreadsheetV1].
	SheetNames []string `json:"sheet_names"`
	// This field is from variant [ConfigurationCreateParametersSpreadsheetV1].
	Specialization string `json:"specialization"`
	// This field is from variant [ConfigurationCreateParametersSpreadsheetV1].
	TableMergeSensitivity string `json:"table_merge_sensitivity"`
	// This field is from variant [ConfigurationCreateParametersSpreadsheetV1].
	UseExperimentalProcessing bool `json:"use_experimental_processing"`
	JSON                      struct {
		ProductType                respjson.Field
		Rules                      respjson.Field
		Mode                       respjson.Field
		ParsingConfiguration       respjson.Field
		DataSchema                 respjson.Field
		CiteSources                respjson.Field
		ConfidenceScores           respjson.Field
		ExtractionTarget           respjson.Field
		MaxPages                   respjson.Field
		ParseConfigID              respjson.Field
		ParseTier                  respjson.Field
		SystemPrompt               respjson.Field
		TargetPages                respjson.Field
		Tier                       respjson.Field
		Version                    respjson.Field
		AgenticOptions             respjson.Field
		ClientName                 respjson.Field
		CropBox                    respjson.Field
		DisableCache               respjson.Field
		FastOptions                respjson.Field
		InputOptions               respjson.Field
		OutputOptions              respjson.Field
		PageRanges                 respjson.Field
		ProcessingControl          respjson.Field
		ProcessingOptions          respjson.Field
		WebhookConfigurationIDs    respjson.Field
		WebhookConfigurations      respjson.Field
		Categories                 respjson.Field
		SplittingStrategy          respjson.Field
		ExtractionRange            respjson.Field
		FlattenHierarchicalTables  respjson.Field
		GenerateAdditionalMetadata respjson.Field
		IncludeHiddenCells         respjson.Field
		SheetNames                 respjson.Field
		Specialization             respjson.Field
		TableMergeSensitivity      respjson.Field
		UseExperimentalProcessing  respjson.Field
		raw                        string
	} `json:"-"`
}

// anyConfigurationCreateParameters is implemented by each variant of
// [ConfigurationCreateParametersUnion] to add type safety for the return type of
// [ConfigurationCreateParametersUnion.AsAny]
type anyConfigurationCreateParameters interface {
	implConfigurationCreateParametersUnion()
}

func (ClassifyV2ParametersResp) implConfigurationCreateParametersUnion()                   {}
func (ExtractV2ParametersResp) implConfigurationCreateParametersUnion()                    {}
func (ParseV2ParametersResp) implConfigurationCreateParametersUnion()                      {}
func (SplitV1ParametersResp) implConfigurationCreateParametersUnion()                      {}
func (ConfigurationCreateParametersSpreadsheetV1) implConfigurationCreateParametersUnion() {}
func (UntypedParametersResp) implConfigurationCreateParametersUnion()                      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConfigurationCreateParametersUnion.AsAny().(type) {
//	case llamacloudprod.ClassifyV2ParametersResp:
//	case llamacloudprod.ExtractV2ParametersResp:
//	case llamacloudprod.ParseV2ParametersResp:
//	case llamacloudprod.SplitV1ParametersResp:
//	case llamacloudprod.ConfigurationCreateParametersSpreadsheetV1:
//	case llamacloudprod.UntypedParametersResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConfigurationCreateParametersUnion) AsAny() anyConfigurationCreateParameters {
	switch u.ProductType {
	case "classify_v2":
		return u.AsClassifyV2()
	case "extract_v2":
		return u.AsExtractV2()
	case "parse_v2":
		return u.AsParseV2()
	case "split_v1":
		return u.AsSplitV1()
	case "spreadsheet_v1":
		return u.AsSpreadsheetV1()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConfigurationCreateParametersUnion) AsClassifyV2() (v ClassifyV2ParametersResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigurationCreateParametersUnion) AsExtractV2() (v ExtractV2ParametersResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigurationCreateParametersUnion) AsParseV2() (v ParseV2ParametersResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigurationCreateParametersUnion) AsSplitV1() (v SplitV1ParametersResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigurationCreateParametersUnion) AsSpreadsheetV1() (v ConfigurationCreateParametersSpreadsheetV1) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigurationCreateParametersUnion) AsUnknown() (v UntypedParametersResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConfigurationCreateParametersUnion) RawJSON() string { return u.JSON.raw }

func (r *ConfigurationCreateParametersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed parameters for a _spreadsheet v1_ product configuration.
type ConfigurationCreateParametersSpreadsheetV1 struct {
	// Product type.
	ProductType constant.SpreadsheetV1 `json:"product_type" default:"spreadsheet_v1"`
	// A1 notation of the range to extract a single region from. If None, the entire
	// sheet is used.
	ExtractionRange string `json:"extraction_range" api:"nullable"`
	// Return a flattened dataframe when a detected table is recognized as
	// hierarchical.
	FlattenHierarchicalTables bool `json:"flatten_hierarchical_tables"`
	// Whether to generate additional metadata (title, description) for each extracted
	// region.
	GenerateAdditionalMetadata bool `json:"generate_additional_metadata"`
	// Whether to include hidden cells when extracting regions from the spreadsheet.
	IncludeHiddenCells bool `json:"include_hidden_cells"`
	// The names of the sheets to extract regions from. If empty, all sheets will be
	// processed.
	SheetNames []string `json:"sheet_names" api:"nullable"`
	// Optional specialization mode for domain-specific extraction. Supported values:
	// 'financial-standard', 'financial-enhanced', 'financial-precise'. Default None
	// uses the general-purpose pipeline.
	Specialization string `json:"specialization" api:"nullable"`
	// Influences how likely similar-looking regions are merged into a single table.
	// Useful for spreadsheets that either have sparse tables (strong merging) or many
	// distinct tables close together (weak merging).
	//
	// Any of "strong", "weak".
	TableMergeSensitivity string `json:"table_merge_sensitivity"`
	// Enables experimental processing. Accuracy may be impacted.
	UseExperimentalProcessing bool `json:"use_experimental_processing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductType                respjson.Field
		ExtractionRange            respjson.Field
		FlattenHierarchicalTables  respjson.Field
		GenerateAdditionalMetadata respjson.Field
		IncludeHiddenCells         respjson.Field
		SheetNames                 respjson.Field
		Specialization             respjson.Field
		TableMergeSensitivity      respjson.Field
		UseExperimentalProcessing  respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigurationCreateParametersSpreadsheetV1) RawJSON() string { return r.JSON.raw }
func (r *ConfigurationCreateParametersSpreadsheetV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for creating a product configuration.
//
// The properties Name, Parameters are required.
type ConfigurationCreateParam struct {
	// Human-readable name for this configuration.
	Name string `json:"name" api:"required"`
	// Product-specific configuration parameters.
	Parameters ConfigurationCreateParametersUnionParam `json:"parameters,omitzero" api:"required"`
	paramObj
}

func (r ConfigurationCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigurationCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigurationCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConfigurationCreateParametersUnionParam struct {
	OfClassifyV2    *ClassifyV2Parameters                            `json:",omitzero,inline"`
	OfExtractV2     *ExtractV2Parameters                             `json:",omitzero,inline"`
	OfParseV2       *ParseV2Parameters                               `json:",omitzero,inline"`
	OfSplitV1       *SplitV1Parameters                               `json:",omitzero,inline"`
	OfSpreadsheetV1 *ConfigurationCreateParametersSpreadsheetV1Param `json:",omitzero,inline"`
	OfUnknown       *UntypedParameters                               `json:",omitzero,inline"`
	paramUnion
}

func (u ConfigurationCreateParametersUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfClassifyV2,
		u.OfExtractV2,
		u.OfParseV2,
		u.OfSplitV1,
		u.OfSpreadsheetV1,
		u.OfUnknown)
}
func (u *ConfigurationCreateParametersUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ConfigurationCreateParametersUnionParam](
		"product_type",
		apijson.Discriminator[ClassifyV2Parameters]("classify_v2"),
		apijson.Discriminator[ExtractV2Parameters]("extract_v2"),
		apijson.Discriminator[ParseV2Parameters]("parse_v2"),
		apijson.Discriminator[SplitV1Parameters]("split_v1"),
		apijson.Discriminator[ConfigurationCreateParametersSpreadsheetV1Param]("spreadsheet_v1"),
		apijson.Discriminator[UntypedParameters]("unknown"),
	)
}

// Typed parameters for a _spreadsheet v1_ product configuration.
//
// The property ProductType is required.
type ConfigurationCreateParametersSpreadsheetV1Param struct {
	// A1 notation of the range to extract a single region from. If None, the entire
	// sheet is used.
	ExtractionRange param.Opt[string] `json:"extraction_range,omitzero"`
	// Optional specialization mode for domain-specific extraction. Supported values:
	// 'financial-standard', 'financial-enhanced', 'financial-precise'. Default None
	// uses the general-purpose pipeline.
	Specialization param.Opt[string] `json:"specialization,omitzero"`
	// Return a flattened dataframe when a detected table is recognized as
	// hierarchical.
	FlattenHierarchicalTables param.Opt[bool] `json:"flatten_hierarchical_tables,omitzero"`
	// Whether to generate additional metadata (title, description) for each extracted
	// region.
	GenerateAdditionalMetadata param.Opt[bool] `json:"generate_additional_metadata,omitzero"`
	// Whether to include hidden cells when extracting regions from the spreadsheet.
	IncludeHiddenCells param.Opt[bool] `json:"include_hidden_cells,omitzero"`
	// Enables experimental processing. Accuracy may be impacted.
	UseExperimentalProcessing param.Opt[bool] `json:"use_experimental_processing,omitzero"`
	// The names of the sheets to extract regions from. If empty, all sheets will be
	// processed.
	SheetNames []string `json:"sheet_names,omitzero"`
	// Influences how likely similar-looking regions are merged into a single table.
	// Useful for spreadsheets that either have sparse tables (strong merging) or many
	// distinct tables close together (weak merging).
	//
	// Any of "strong", "weak".
	TableMergeSensitivity string `json:"table_merge_sensitivity,omitzero"`
	// Product type.
	//
	// This field can be elided, and will marshal its zero value as "spreadsheet_v1".
	ProductType constant.SpreadsheetV1 `json:"product_type" default:"spreadsheet_v1"`
	paramObj
}

func (r ConfigurationCreateParametersSpreadsheetV1Param) MarshalJSON() (data []byte, err error) {
	type shadow ConfigurationCreateParametersSpreadsheetV1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigurationCreateParametersSpreadsheetV1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConfigurationCreateParametersSpreadsheetV1Param](
		"table_merge_sensitivity", "strong", "weak",
	)
}

// Response schema for a single product configuration.
type ConfigurationResponse struct {
	// Unique configuration ID.
	ID string `json:"id" api:"required"`
	// Configuration name.
	Name string `json:"name" api:"required"`
	// Product-specific configuration parameters.
	Parameters ConfigurationResponseParametersUnion `json:"parameters" api:"required"`
	// Product type.
	//
	// Any of "classify_v2", "extract_v2", "parse_v2", "split_v1", "spreadsheet_v1",
	// "unknown".
	ProductType ConfigurationResponseProductType `json:"product_type" api:"required"`
	// Version identifier (datetime string).
	Version string `json:"version" api:"required"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Parameters  respjson.Field
		ProductType respjson.Field
		Version     respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigurationResponse) RawJSON() string { return r.JSON.raw }
func (r *ConfigurationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConfigurationResponseParametersUnion contains all possible properties and values
// from [ClassifyV2ParametersResp], [ExtractV2ParametersResp],
// [ParseV2ParametersResp], [SplitV1ParametersResp],
// [ConfigurationResponseParametersSpreadsheetV1], [UntypedParametersResp].
//
// Use the [ConfigurationResponseParametersUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConfigurationResponseParametersUnion struct {
	// Any of "classify_v2", "extract_v2", "parse_v2", "split_v1", "spreadsheet_v1",
	// "unknown".
	ProductType string `json:"product_type"`
	// This field is from variant [ClassifyV2ParametersResp].
	Rules []ClassifyV2ParametersRuleResp `json:"rules"`
	// This field is from variant [ClassifyV2ParametersResp].
	Mode ClassifyV2ParametersMode `json:"mode"`
	// This field is from variant [ClassifyV2ParametersResp].
	ParsingConfiguration ClassifyV2ParametersParsingConfigurationResp `json:"parsing_configuration"`
	// This field is from variant [ExtractV2ParametersResp].
	DataSchema map[string]*ExtractV2ParametersDataSchemaUnionResp `json:"data_schema"`
	// This field is from variant [ExtractV2ParametersResp].
	CiteSources bool `json:"cite_sources"`
	// This field is from variant [ExtractV2ParametersResp].
	ConfidenceScores bool `json:"confidence_scores"`
	// This field is from variant [ExtractV2ParametersResp].
	ExtractionTarget ExtractV2ParametersExtractionTarget `json:"extraction_target"`
	// This field is from variant [ExtractV2ParametersResp].
	MaxPages int64 `json:"max_pages"`
	// This field is from variant [ExtractV2ParametersResp].
	ParseConfigID string `json:"parse_config_id"`
	// This field is from variant [ExtractV2ParametersResp].
	ParseTier string `json:"parse_tier"`
	// This field is from variant [ExtractV2ParametersResp].
	SystemPrompt string `json:"system_prompt"`
	// This field is from variant [ExtractV2ParametersResp].
	TargetPages string `json:"target_pages"`
	Tier        string `json:"tier"`
	Version     string `json:"version"`
	// This field is from variant [ParseV2ParametersResp].
	AgenticOptions ParseV2ParametersAgenticOptionsResp `json:"agentic_options"`
	// This field is from variant [ParseV2ParametersResp].
	ClientName string `json:"client_name"`
	// This field is from variant [ParseV2ParametersResp].
	CropBox ParseV2ParametersCropBoxResp `json:"crop_box"`
	// This field is from variant [ParseV2ParametersResp].
	DisableCache bool `json:"disable_cache"`
	// This field is from variant [ParseV2ParametersResp].
	FastOptions any `json:"fast_options"`
	// This field is from variant [ParseV2ParametersResp].
	InputOptions ParseV2ParametersInputOptionsResp `json:"input_options"`
	// This field is from variant [ParseV2ParametersResp].
	OutputOptions ParseV2ParametersOutputOptionsResp `json:"output_options"`
	// This field is from variant [ParseV2ParametersResp].
	PageRanges ParseV2ParametersPageRangesResp `json:"page_ranges"`
	// This field is from variant [ParseV2ParametersResp].
	ProcessingControl ParseV2ParametersProcessingControlResp `json:"processing_control"`
	// This field is from variant [ParseV2ParametersResp].
	ProcessingOptions ParseV2ParametersProcessingOptionsResp `json:"processing_options"`
	// This field is from variant [ParseV2ParametersResp].
	WebhookConfigurationIDs []string `json:"webhook_configuration_ids"`
	// This field is from variant [ParseV2ParametersResp].
	WebhookConfigurations []ParseV2ParametersWebhookConfigurationResp `json:"webhook_configurations"`
	// This field is from variant [SplitV1ParametersResp].
	Categories []SplitCategory `json:"categories"`
	// This field is from variant [SplitV1ParametersResp].
	SplittingStrategy SplitV1ParametersSplittingStrategyResp `json:"splitting_strategy"`
	// This field is from variant [ConfigurationResponseParametersSpreadsheetV1].
	ExtractionRange string `json:"extraction_range"`
	// This field is from variant [ConfigurationResponseParametersSpreadsheetV1].
	FlattenHierarchicalTables bool `json:"flatten_hierarchical_tables"`
	// This field is from variant [ConfigurationResponseParametersSpreadsheetV1].
	GenerateAdditionalMetadata bool `json:"generate_additional_metadata"`
	// This field is from variant [ConfigurationResponseParametersSpreadsheetV1].
	IncludeHiddenCells bool `json:"include_hidden_cells"`
	// This field is from variant [ConfigurationResponseParametersSpreadsheetV1].
	SheetNames []string `json:"sheet_names"`
	// This field is from variant [ConfigurationResponseParametersSpreadsheetV1].
	Specialization string `json:"specialization"`
	// This field is from variant [ConfigurationResponseParametersSpreadsheetV1].
	TableMergeSensitivity string `json:"table_merge_sensitivity"`
	// This field is from variant [ConfigurationResponseParametersSpreadsheetV1].
	UseExperimentalProcessing bool `json:"use_experimental_processing"`
	JSON                      struct {
		ProductType                respjson.Field
		Rules                      respjson.Field
		Mode                       respjson.Field
		ParsingConfiguration       respjson.Field
		DataSchema                 respjson.Field
		CiteSources                respjson.Field
		ConfidenceScores           respjson.Field
		ExtractionTarget           respjson.Field
		MaxPages                   respjson.Field
		ParseConfigID              respjson.Field
		ParseTier                  respjson.Field
		SystemPrompt               respjson.Field
		TargetPages                respjson.Field
		Tier                       respjson.Field
		Version                    respjson.Field
		AgenticOptions             respjson.Field
		ClientName                 respjson.Field
		CropBox                    respjson.Field
		DisableCache               respjson.Field
		FastOptions                respjson.Field
		InputOptions               respjson.Field
		OutputOptions              respjson.Field
		PageRanges                 respjson.Field
		ProcessingControl          respjson.Field
		ProcessingOptions          respjson.Field
		WebhookConfigurationIDs    respjson.Field
		WebhookConfigurations      respjson.Field
		Categories                 respjson.Field
		SplittingStrategy          respjson.Field
		ExtractionRange            respjson.Field
		FlattenHierarchicalTables  respjson.Field
		GenerateAdditionalMetadata respjson.Field
		IncludeHiddenCells         respjson.Field
		SheetNames                 respjson.Field
		Specialization             respjson.Field
		TableMergeSensitivity      respjson.Field
		UseExperimentalProcessing  respjson.Field
		raw                        string
	} `json:"-"`
}

// anyConfigurationResponseParameters is implemented by each variant of
// [ConfigurationResponseParametersUnion] to add type safety for the return type of
// [ConfigurationResponseParametersUnion.AsAny]
type anyConfigurationResponseParameters interface {
	implConfigurationResponseParametersUnion()
}

func (ClassifyV2ParametersResp) implConfigurationResponseParametersUnion()                     {}
func (ExtractV2ParametersResp) implConfigurationResponseParametersUnion()                      {}
func (ParseV2ParametersResp) implConfigurationResponseParametersUnion()                        {}
func (SplitV1ParametersResp) implConfigurationResponseParametersUnion()                        {}
func (ConfigurationResponseParametersSpreadsheetV1) implConfigurationResponseParametersUnion() {}
func (UntypedParametersResp) implConfigurationResponseParametersUnion()                        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConfigurationResponseParametersUnion.AsAny().(type) {
//	case llamacloudprod.ClassifyV2ParametersResp:
//	case llamacloudprod.ExtractV2ParametersResp:
//	case llamacloudprod.ParseV2ParametersResp:
//	case llamacloudprod.SplitV1ParametersResp:
//	case llamacloudprod.ConfigurationResponseParametersSpreadsheetV1:
//	case llamacloudprod.UntypedParametersResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConfigurationResponseParametersUnion) AsAny() anyConfigurationResponseParameters {
	switch u.ProductType {
	case "classify_v2":
		return u.AsClassifyV2()
	case "extract_v2":
		return u.AsExtractV2()
	case "parse_v2":
		return u.AsParseV2()
	case "split_v1":
		return u.AsSplitV1()
	case "spreadsheet_v1":
		return u.AsSpreadsheetV1()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConfigurationResponseParametersUnion) AsClassifyV2() (v ClassifyV2ParametersResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigurationResponseParametersUnion) AsExtractV2() (v ExtractV2ParametersResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigurationResponseParametersUnion) AsParseV2() (v ParseV2ParametersResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigurationResponseParametersUnion) AsSplitV1() (v SplitV1ParametersResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigurationResponseParametersUnion) AsSpreadsheetV1() (v ConfigurationResponseParametersSpreadsheetV1) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConfigurationResponseParametersUnion) AsUnknown() (v UntypedParametersResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConfigurationResponseParametersUnion) RawJSON() string { return u.JSON.raw }

func (r *ConfigurationResponseParametersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed parameters for a _spreadsheet v1_ product configuration.
type ConfigurationResponseParametersSpreadsheetV1 struct {
	// Product type.
	ProductType constant.SpreadsheetV1 `json:"product_type" default:"spreadsheet_v1"`
	// A1 notation of the range to extract a single region from. If None, the entire
	// sheet is used.
	ExtractionRange string `json:"extraction_range" api:"nullable"`
	// Return a flattened dataframe when a detected table is recognized as
	// hierarchical.
	FlattenHierarchicalTables bool `json:"flatten_hierarchical_tables"`
	// Whether to generate additional metadata (title, description) for each extracted
	// region.
	GenerateAdditionalMetadata bool `json:"generate_additional_metadata"`
	// Whether to include hidden cells when extracting regions from the spreadsheet.
	IncludeHiddenCells bool `json:"include_hidden_cells"`
	// The names of the sheets to extract regions from. If empty, all sheets will be
	// processed.
	SheetNames []string `json:"sheet_names" api:"nullable"`
	// Optional specialization mode for domain-specific extraction. Supported values:
	// 'financial-standard', 'financial-enhanced', 'financial-precise'. Default None
	// uses the general-purpose pipeline.
	Specialization string `json:"specialization" api:"nullable"`
	// Influences how likely similar-looking regions are merged into a single table.
	// Useful for spreadsheets that either have sparse tables (strong merging) or many
	// distinct tables close together (weak merging).
	//
	// Any of "strong", "weak".
	TableMergeSensitivity string `json:"table_merge_sensitivity"`
	// Enables experimental processing. Accuracy may be impacted.
	UseExperimentalProcessing bool `json:"use_experimental_processing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductType                respjson.Field
		ExtractionRange            respjson.Field
		FlattenHierarchicalTables  respjson.Field
		GenerateAdditionalMetadata respjson.Field
		IncludeHiddenCells         respjson.Field
		SheetNames                 respjson.Field
		Specialization             respjson.Field
		TableMergeSensitivity      respjson.Field
		UseExperimentalProcessing  respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigurationResponseParametersSpreadsheetV1) RawJSON() string { return r.JSON.raw }
func (r *ConfigurationResponseParametersSpreadsheetV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Product type.
type ConfigurationResponseProductType string

const (
	ConfigurationResponseProductTypeClassifyV2    ConfigurationResponseProductType = "classify_v2"
	ConfigurationResponseProductTypeExtractV2     ConfigurationResponseProductType = "extract_v2"
	ConfigurationResponseProductTypeParseV2       ConfigurationResponseProductType = "parse_v2"
	ConfigurationResponseProductTypeSplitV1       ConfigurationResponseProductType = "split_v1"
	ConfigurationResponseProductTypeSpreadsheetV1 ConfigurationResponseProductType = "spreadsheet_v1"
	ConfigurationResponseProductTypeUnknown       ConfigurationResponseProductType = "unknown"
)

// Typed parameters for an _extract v2_ product configuration.
type ExtractV2ParametersResp struct {
	// JSON Schema defining the fields to extract. Validate with the /schema/validate
	// endpoint first.
	DataSchema map[string]*ExtractV2ParametersDataSchemaUnionResp `json:"data_schema" api:"required"`
	// Product type.
	ProductType constant.ExtractV2 `json:"product_type" default:"extract_v2"`
	// Include citations in results
	CiteSources bool `json:"cite_sources"`
	// Include confidence scores in results
	ConfidenceScores bool `json:"confidence_scores"`
	// Granularity of extraction: per_doc returns one object per document, per_page
	// returns one object per page, per_table_row returns one object per table row
	//
	// Any of "per_doc", "per_page", "per_table_row".
	ExtractionTarget ExtractV2ParametersExtractionTarget `json:"extraction_target"`
	// Maximum number of pages to process. Omit for no limit.
	MaxPages int64 `json:"max_pages" api:"nullable"`
	// Saved parse configuration ID to control how the document is parsed before
	// extraction
	ParseConfigID string `json:"parse_config_id" api:"nullable"`
	// Parse tier to use before extraction. Defaults to the extract tier if not
	// specified.
	ParseTier string `json:"parse_tier" api:"nullable"`
	// Custom system prompt to guide extraction behavior
	SystemPrompt string `json:"system_prompt" api:"nullable"`
	// Comma-separated page numbers or ranges to process (1-based). Omit to process all
	// pages.
	TargetPages string `json:"target_pages" api:"nullable"`
	// Extract tier: cost_effective (5 credits/page) or agentic (15 credits/page)
	//
	// Any of "agentic", "cost_effective".
	Tier ExtractV2ParametersTier `json:"tier"`
	// Use 'latest' for the latest release for the selected tier or a date string
	// (YYYY-MM-DD format) to pin to the nearest release at or before that date. Job
	// responses always report the concrete resolved version the job runs, fixed at job
	// creation; saved configurations keep the value as provided.
	Version string `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSchema       respjson.Field
		ProductType      respjson.Field
		CiteSources      respjson.Field
		ConfidenceScores respjson.Field
		ExtractionTarget respjson.Field
		MaxPages         respjson.Field
		ParseConfigID    respjson.Field
		ParseTier        respjson.Field
		SystemPrompt     respjson.Field
		TargetPages      respjson.Field
		Tier             respjson.Field
		Version          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractV2ParametersResp) RawJSON() string { return r.JSON.raw }
func (r *ExtractV2ParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExtractV2ParametersResp to a ExtractV2Parameters.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExtractV2Parameters.Overrides()
func (r ExtractV2ParametersResp) ToParam() ExtractV2Parameters {
	return param.Override[ExtractV2Parameters](json.RawMessage(r.RawJSON()))
}

// ExtractV2ParametersDataSchemaUnionResp contains all possible properties and
// values from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractV2etersDataSchemaMapItem OfAnyArray OfString OfFloat
// OfBool]
type ExtractV2ParametersDataSchemaUnionResp struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractV2etersDataSchemaMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfExtractV2etersDataSchemaMapItem respjson.Field
		OfAnyArray                        respjson.Field
		OfString                          respjson.Field
		OfFloat                           respjson.Field
		OfBool                            respjson.Field
		raw                               string
	} `json:"-"`
}

func (u ExtractV2ParametersDataSchemaUnionResp) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2ParametersDataSchemaUnionResp) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2ParametersDataSchemaUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2ParametersDataSchemaUnionResp) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractV2ParametersDataSchemaUnionResp) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractV2ParametersDataSchemaUnionResp) RawJSON() string { return u.JSON.raw }

func (r *ExtractV2ParametersDataSchemaUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Granularity of extraction: per_doc returns one object per document, per_page
// returns one object per page, per_table_row returns one object per table row
type ExtractV2ParametersExtractionTarget string

const (
	ExtractV2ParametersExtractionTargetPerDoc      ExtractV2ParametersExtractionTarget = "per_doc"
	ExtractV2ParametersExtractionTargetPerPage     ExtractV2ParametersExtractionTarget = "per_page"
	ExtractV2ParametersExtractionTargetPerTableRow ExtractV2ParametersExtractionTarget = "per_table_row"
)

// Extract tier: cost_effective (5 credits/page) or agentic (15 credits/page)
type ExtractV2ParametersTier string

const (
	ExtractV2ParametersTierAgentic       ExtractV2ParametersTier = "agentic"
	ExtractV2ParametersTierCostEffective ExtractV2ParametersTier = "cost_effective"
)

// Typed parameters for an _extract v2_ product configuration.
//
// The properties DataSchema, ProductType are required.
type ExtractV2Parameters struct {
	// JSON Schema defining the fields to extract. Validate with the /schema/validate
	// endpoint first.
	DataSchema map[string]*ExtractV2ParametersDataSchemaUnion `json:"data_schema,omitzero" api:"required"`
	// Maximum number of pages to process. Omit for no limit.
	MaxPages param.Opt[int64] `json:"max_pages,omitzero"`
	// Saved parse configuration ID to control how the document is parsed before
	// extraction
	ParseConfigID param.Opt[string] `json:"parse_config_id,omitzero"`
	// Parse tier to use before extraction. Defaults to the extract tier if not
	// specified.
	ParseTier param.Opt[string] `json:"parse_tier,omitzero"`
	// Custom system prompt to guide extraction behavior
	SystemPrompt param.Opt[string] `json:"system_prompt,omitzero"`
	// Comma-separated page numbers or ranges to process (1-based). Omit to process all
	// pages.
	TargetPages param.Opt[string] `json:"target_pages,omitzero"`
	// Include citations in results
	CiteSources param.Opt[bool] `json:"cite_sources,omitzero"`
	// Include confidence scores in results
	ConfidenceScores param.Opt[bool] `json:"confidence_scores,omitzero"`
	// Use 'latest' for the latest release for the selected tier or a date string
	// (YYYY-MM-DD format) to pin to the nearest release at or before that date. Job
	// responses always report the concrete resolved version the job runs, fixed at job
	// creation; saved configurations keep the value as provided.
	Version param.Opt[string] `json:"version,omitzero"`
	// Granularity of extraction: per_doc returns one object per document, per_page
	// returns one object per page, per_table_row returns one object per table row
	//
	// Any of "per_doc", "per_page", "per_table_row".
	ExtractionTarget ExtractV2ParametersExtractionTarget `json:"extraction_target,omitzero"`
	// Extract tier: cost_effective (5 credits/page) or agentic (15 credits/page)
	//
	// Any of "agentic", "cost_effective".
	Tier ExtractV2ParametersTier `json:"tier,omitzero"`
	// Product type.
	//
	// This field can be elided, and will marshal its zero value as "extract_v2".
	ProductType constant.ExtractV2 `json:"product_type" default:"extract_v2"`
	paramObj
}

func (r ExtractV2Parameters) MarshalJSON() (data []byte, err error) {
	type shadow ExtractV2Parameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractV2Parameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractV2ParametersDataSchemaUnion struct {
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractV2ParametersDataSchemaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfAnyArray,
		u.OfString,
		u.OfFloat,
		u.OfBool)
}
func (u *ExtractV2ParametersDataSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Configuration for LlamaParse v2 document parsing.
//
// Includes tier selection, processing options, output formatting, page targeting,
// and webhook delivery. Refer to the LlamaParse documentation for details on each
// field.
type ParseV2ParametersResp struct {
	// Product type.
	ProductType constant.ParseV2 `json:"product_type" default:"parse_v2"`
	// Parsing tier: 'fast' (rule-based, cheapest), 'cost_effective' (balanced),
	// 'agentic' (AI-powered with custom prompts), or 'agentic_plus' (premium AI with
	// highest accuracy)
	//
	// Any of "agentic", "agentic_plus", "cost_effective", "fast".
	Tier ParseV2ParametersTier `json:"tier" api:"required"`
	// Version for the selected tier. Use `latest`, or pin one of that tier's dated
	// versions.
	//
	// Current `latest` by tier:
	//
	// - `fast`: `2025-12-11`
	// - `cost_effective`: `2026-06-26`
	// - `agentic`: `2026-06-18`
	// - `agentic_plus`: `2026-06-18`
	//
	// Full list: `GET /api/v2/parse/versions`.
	Version ParseV2ParametersVersion `json:"version" api:"required"`
	// Options for AI-powered parsing tiers (cost_effective, agentic, agentic_plus).
	//
	// These options customize how the AI processes and interprets document content.
	// Only applicable when using non-fast tiers.
	AgenticOptions ParseV2ParametersAgenticOptionsResp `json:"agentic_options" api:"nullable"`
	// Identifier for the client/application making the request. Used for analytics and
	// debugging. Example: 'my-app-v2'
	ClientName string `json:"client_name" api:"nullable"`
	// Crop boundaries to process only a portion of each page. Values are ratios 0-1
	// from page edges
	CropBox ParseV2ParametersCropBoxResp `json:"crop_box"`
	// Bypass result caching and force re-parsing. Use when document content may have
	// changed or you need fresh results
	DisableCache bool `json:"disable_cache" api:"nullable"`
	// Options for fast tier parsing (rule-based, no AI).
	//
	// Fast tier uses deterministic algorithms for text extraction without AI
	// enhancement. It's the fastest and most cost-effective option, best suited for
	// simple documents with standard layouts. Currently has no configurable options
	// but reserved for future expansion.
	FastOptions any `json:"fast_options" api:"nullable"`
	// Format-specific options (HTML, PDF, spreadsheet, presentation). Applied based on
	// detected input file type
	InputOptions ParseV2ParametersInputOptionsResp `json:"input_options"`
	// Output formatting options for markdown, text, and extracted images
	OutputOptions ParseV2ParametersOutputOptionsResp `json:"output_options"`
	// Page selection: limit total pages or specify exact pages to process
	PageRanges ParseV2ParametersPageRangesResp `json:"page_ranges"`
	// Job execution controls including timeouts and failure thresholds
	ProcessingControl ParseV2ParametersProcessingControlResp `json:"processing_control"`
	// Document processing options including OCR, table extraction, and chart parsing
	ProcessingOptions ParseV2ParametersProcessingOptionsResp `json:"processing_options"`
	// IDs of saved webhook configurations to notify for this job.
	WebhookConfigurationIDs []string `json:"webhook_configuration_ids" api:"nullable"`
	// Webhook endpoints for job status notifications. Multiple webhooks can be
	// configured for different events or services
	WebhookConfigurations []ParseV2ParametersWebhookConfigurationResp `json:"webhook_configurations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductType             respjson.Field
		Tier                    respjson.Field
		Version                 respjson.Field
		AgenticOptions          respjson.Field
		ClientName              respjson.Field
		CropBox                 respjson.Field
		DisableCache            respjson.Field
		FastOptions             respjson.Field
		InputOptions            respjson.Field
		OutputOptions           respjson.Field
		PageRanges              respjson.Field
		ProcessingControl       respjson.Field
		ProcessingOptions       respjson.Field
		WebhookConfigurationIDs respjson.Field
		WebhookConfigurations   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ParseV2ParametersResp to a ParseV2Parameters.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ParseV2Parameters.Overrides()
func (r ParseV2ParametersResp) ToParam() ParseV2Parameters {
	return param.Override[ParseV2Parameters](json.RawMessage(r.RawJSON()))
}

// Parsing tier: 'fast' (rule-based, cheapest), 'cost_effective' (balanced),
// 'agentic' (AI-powered with custom prompts), or 'agentic_plus' (premium AI with
// highest accuracy)
type ParseV2ParametersTier string

const (
	ParseV2ParametersTierAgentic       ParseV2ParametersTier = "agentic"
	ParseV2ParametersTierAgenticPlus   ParseV2ParametersTier = "agentic_plus"
	ParseV2ParametersTierCostEffective ParseV2ParametersTier = "cost_effective"
	ParseV2ParametersTierFast          ParseV2ParametersTier = "fast"
)

// Version for the selected tier. Use `latest`, or pin one of that tier's dated
// versions.
//
// Current `latest` by tier:
//
// - `fast`: `2025-12-11`
// - `cost_effective`: `2026-06-26`
// - `agentic`: `2026-06-18`
// - `agentic_plus`: `2026-06-18`
//
// Full list: `GET /api/v2/parse/versions`.
type ParseV2ParametersVersion string

const (
	ParseV2ParametersVersionLatest     ParseV2ParametersVersion = "latest"
	ParseV2ParametersVersion2026_06_26 ParseV2ParametersVersion = "2026-06-26"
	ParseV2ParametersVersion2026_06_18 ParseV2ParametersVersion = "2026-06-18"
	ParseV2ParametersVersion2025_12_11 ParseV2ParametersVersion = "2025-12-11"
)

// Options for AI-powered parsing tiers (cost_effective, agentic, agentic_plus).
//
// These options customize how the AI processes and interprets document content.
// Only applicable when using non-fast tiers.
type ParseV2ParametersAgenticOptionsResp struct {
	// Custom instructions for the AI parser. Use to guide extraction behavior, specify
	// output formatting, or provide domain-specific context. Example: 'Extract
	// financial tables with currency symbols. Format dates as YYYY-MM-DD.'
	CustomPrompt string `json:"custom_prompt" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomPrompt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersAgenticOptionsResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersAgenticOptionsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Crop boundaries to process only a portion of each page. Values are ratios 0-1
// from page edges
type ParseV2ParametersCropBoxResp struct {
	// Bottom boundary as ratio (0-1). 0=top edge, 1=bottom edge. Content below this
	// line is excluded
	Bottom float64 `json:"bottom" api:"nullable"`
	// Left boundary as ratio (0-1). 0=left edge, 1=right edge. Content left of this
	// line is excluded
	Left float64 `json:"left" api:"nullable"`
	// Right boundary as ratio (0-1). 0=left edge, 1=right edge. Content right of this
	// line is excluded
	Right float64 `json:"right" api:"nullable"`
	// Top boundary as ratio (0-1). 0=top edge, 1=bottom edge. Content above this line
	// is excluded
	Top float64 `json:"top" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bottom      respjson.Field
		Left        respjson.Field
		Right       respjson.Field
		Top         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersCropBoxResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersCropBoxResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format-specific options (HTML, PDF, spreadsheet, presentation). Applied based on
// detected input file type
type ParseV2ParametersInputOptionsResp struct {
	// HTML/web page parsing options (applies to .html, .htm files)
	HTML ParseV2ParametersInputOptionsHTMLResp `json:"html"`
	// Image parsing options (applies to .jpg, .jpeg, .png, .webp files)
	Image ParseV2ParametersInputOptionsImageResp `json:"image"`
	// PDF-specific parsing options (applies to .pdf files)
	Pdf any `json:"pdf"`
	// Presentation parsing options (applies to .pptx, .ppt, .odp, .key files)
	Presentation ParseV2ParametersInputOptionsPresentationResp `json:"presentation"`
	// Spreadsheet parsing options (applies to .xlsx, .xls, .csv, .ods files)
	Spreadsheet ParseV2ParametersInputOptionsSpreadsheetResp `json:"spreadsheet"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTML         respjson.Field
		Image        respjson.Field
		Pdf          respjson.Field
		Presentation respjson.Field
		Spreadsheet  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersInputOptionsResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersInputOptionsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTML/web page parsing options (applies to .html, .htm files)
type ParseV2ParametersInputOptionsHTMLResp struct {
	// Force all HTML elements to be visible by overriding CSS display/visibility
	// properties. Useful for parsing pages with hidden content or collapsed sections
	MakeAllElementsVisible bool `json:"make_all_elements_visible" api:"nullable"`
	// Remove fixed-position elements (headers, footers, floating buttons) that appear
	// on every page render
	RemoveFixedElements bool `json:"remove_fixed_elements" api:"nullable"`
	// Remove navigation elements (nav bars, sidebars, menus) to focus on main content
	RemoveNavigationElements bool `json:"remove_navigation_elements" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MakeAllElementsVisible   respjson.Field
		RemoveFixedElements      respjson.Field
		RemoveNavigationElements respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersInputOptionsHTMLResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersInputOptionsHTMLResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image parsing options (applies to .jpg, .jpeg, .png, .webp files)
type ParseV2ParametersInputOptionsImageResp struct {
	// Detect documents photographed with a camera (e.g. phone scans of receipts or
	// forms), then crop, perspective-correct, and flatten uneven lighting and shadows
	// before parsing. Supports JPEG, PNG, WebP, and HEIC/HEIF inputs. Improves results
	// when the document is tilted or surrounded by background. Images that already
	// look like clean scans are left untouched
	CameraPhotoCorrection bool `json:"camera_photo_correction" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CameraPhotoCorrection respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersInputOptionsImageResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersInputOptionsImageResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Presentation parsing options (applies to .pptx, .ppt, .odp, .key files)
type ParseV2ParametersInputOptionsPresentationResp struct {
	// Extract content positioned outside the visible slide area. Some presentations
	// have hidden notes or content that extends beyond slide boundaries
	OutOfBoundsContent bool `json:"out_of_bounds_content" api:"nullable"`
	// Skip extraction of embedded chart data tables. When true, only the visual
	// representation of charts is captured, not the underlying data
	SkipEmbeddedData bool `json:"skip_embedded_data" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OutOfBoundsContent respjson.Field
		SkipEmbeddedData   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersInputOptionsPresentationResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersInputOptionsPresentationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spreadsheet parsing options (applies to .xlsx, .xls, .csv, .ods files)
type ParseV2ParametersInputOptionsSpreadsheetResp struct {
	// Detect and extract multiple tables within a single sheet. Useful when
	// spreadsheets contain several data regions separated by blank rows/columns
	DetectSubTablesInSheets bool `json:"detect_sub_tables_in_sheets" api:"nullable"`
	// Compute formula results instead of extracting formula text. Use when you need
	// calculated values rather than formula definitions
	ForceFormulaComputationInSheets bool `json:"force_formula_computation_in_sheets" api:"nullable"`
	// Parse hidden sheets in addition to visible ones. By default, hidden sheets are
	// skipped
	IncludeHiddenSheets bool `json:"include_hidden_sheets" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectSubTablesInSheets         respjson.Field
		ForceFormulaComputationInSheets respjson.Field
		IncludeHiddenSheets             respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersInputOptionsSpreadsheetResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersInputOptionsSpreadsheetResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Output formatting options for markdown, text, and extracted images
type ParseV2ParametersOutputOptionsResp struct {
	// Optional additional output artifacts to save alongside the primary parse output.
	// Each value opts in to generating and persisting one extra file; the empty list
	// (default) saves none. The three accepted values are: 'stripped_md' — per-page
	// markdown stripped of formatting (links, bold/italic, images, HTML), saved as
	// JSON for full-text-search indexing; fetch via
	// `expand=stripped_markdown_content_metadata`. 'concatenated_stripped_txt' — all
	// stripped pages concatenated into a single plain-text file with `\n\n---\n\n`
	// between pages, useful for feeding the document into search or embedding
	// pipelines as one blob; fetch via
	// `expand=concatenated_stripped_markdown_content_metadata`. 'word_bbox' — raw
	// word-level bounding boxes (one JSON object per word, with page number and
	// x/y/w/h coordinates) saved as JSONL, useful for highlighting or grounding
	// extracted answers back to the source document; fetch via
	// `expand=raw_words_content_metadata`.
	AdditionalOutputs []string `json:"additional_outputs"`
	// Extract the printed page number as it appears in the document (e.g., 'Page 5 of
	// 10', 'v', 'A-3'). Useful for referencing original page numbers
	ExtractPrintedPageNumber bool `json:"extract_printed_page_number" api:"nullable"`
	// Bounding-box granularity levels to compute for the parse. 'word' computes one
	// bounding box per detected word; 'line' computes one per text line; 'cell'
	// computes one per table cell. Multiple levels can be requested. Empty list
	// (default) disables granular bboxes — only item-level layout boxes are returned
	// on the result. When set, the computed boxes are not inlined on the result items;
	// they are written to a separate `grounded_items` sidecar (JSONL, one row per
	// page) and exposed as `result_content_metadata.grounded_items` (a presigned
	// download URL) on the parse result. Each row matches the `GroundedJsonItem`
	// shape.
	//
	// Any of "cell", "line", "word".
	GranularBboxes []string `json:"granular_bboxes"`
	// Image categories to extract and save. Options: 'screenshot' (full page renders
	// useful for visual QA), 'embedded' (images found within the document), 'layout'
	// (cropped regions from layout detection like figures and diagrams). Empty list
	// saves no images
	//
	// Any of "embedded", "layout", "screenshot".
	ImagesToSave []string `json:"images_to_save"`
	// Markdown formatting options including table styles and link annotations
	Markdown ParseV2ParametersOutputOptionsMarkdownResp `json:"markdown"`
	// Spatial text output options for preserving document layout structure
	SpatialText ParseV2ParametersOutputOptionsSpatialTextResp `json:"spatial_text"`
	// Options for exporting tables as XLSX spreadsheets
	TablesAsSpreadsheet ParseV2ParametersOutputOptionsTablesAsSpreadsheetResp `json:"tables_as_spreadsheet"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalOutputs        respjson.Field
		ExtractPrintedPageNumber respjson.Field
		GranularBboxes           respjson.Field
		ImagesToSave             respjson.Field
		Markdown                 respjson.Field
		SpatialText              respjson.Field
		TablesAsSpreadsheet      respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersOutputOptionsResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersOutputOptionsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown formatting options including table styles and link annotations
type ParseV2ParametersOutputOptionsMarkdownResp struct {
	// Add link annotations to markdown output in the format [text](url). When false,
	// only the link text is included
	AnnotateLinks bool `json:"annotate_links" api:"nullable"`
	// Embed images directly in markdown as base64 data URIs instead of extracting them
	// as separate files. Useful for self-contained markdown output
	InlineImages bool `json:"inline_images" api:"nullable"`
	// Table formatting options including markdown vs HTML format and merging behavior
	Tables ParseV2ParametersOutputOptionsMarkdownTablesResp `json:"tables"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnnotateLinks respjson.Field
		InlineImages  respjson.Field
		Tables        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersOutputOptionsMarkdownResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersOutputOptionsMarkdownResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Table formatting options including markdown vs HTML format and merging behavior
type ParseV2ParametersOutputOptionsMarkdownTablesResp struct {
	// Remove extra whitespace padding in markdown table cells for more compact output
	CompactMarkdownTables bool `json:"compact_markdown_tables" api:"nullable"`
	// Separator string for multiline cell content in markdown tables. Example:
	// '&lt;br&gt;' to preserve line breaks, ' ' to join with spaces
	MarkdownTableMultilineSeparator string `json:"markdown_table_multiline_separator" api:"nullable"`
	// Automatically merge tables that span multiple pages into a single table. The
	// merged table appears on the first page with merged_from_pages metadata
	MergeContinuedTables bool `json:"merge_continued_tables" api:"nullable"`
	// Output tables as markdown pipe tables instead of HTML &lt;table&gt; tags.
	// Markdown tables are simpler but cannot represent complex structures like merged
	// cells
	OutputTablesAsMarkdown bool `json:"output_tables_as_markdown" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompactMarkdownTables           respjson.Field
		MarkdownTableMultilineSeparator respjson.Field
		MergeContinuedTables            respjson.Field
		OutputTablesAsMarkdown          respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersOutputOptionsMarkdownTablesResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersOutputOptionsMarkdownTablesResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spatial text output options for preserving document layout structure
type ParseV2ParametersOutputOptionsSpatialTextResp struct {
	// Keep multi-column layouts intact instead of linearizing columns into sequential
	// text. Automatically enabled for non-fast tiers
	DoNotUnrollColumns bool `json:"do_not_unroll_columns" api:"nullable"`
	// Maintain consistent text column alignment across page boundaries. Automatically
	// enabled for document-level parsing modes
	PreserveLayoutAlignmentAcrossPages bool `json:"preserve_layout_alignment_across_pages" api:"nullable"`
	// Include text below the normal size threshold. Useful for footnotes, watermarks,
	// or fine print that might otherwise be filtered out
	PreserveVerySmallText bool `json:"preserve_very_small_text" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DoNotUnrollColumns                 respjson.Field
		PreserveLayoutAlignmentAcrossPages respjson.Field
		PreserveVerySmallText              respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersOutputOptionsSpatialTextResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersOutputOptionsSpatialTextResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for exporting tables as XLSX spreadsheets
type ParseV2ParametersOutputOptionsTablesAsSpreadsheetResp struct {
	// Whether this option is enabled
	Enable bool `json:"enable" api:"nullable"`
	// Automatically generate descriptive sheet names from table context (headers,
	// surrounding text) instead of using generic names like 'Table_1'
	GuessSheetName bool `json:"guess_sheet_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enable         respjson.Field
		GuessSheetName respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersOutputOptionsTablesAsSpreadsheetResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersOutputOptionsTablesAsSpreadsheetResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page selection: limit total pages or specify exact pages to process
type ParseV2ParametersPageRangesResp struct {
	// Maximum number of pages to process. Pages are processed in order starting from
	// page 1. If both max_pages and target_pages are set, target_pages takes
	// precedence
	MaxPages int64 `json:"max_pages" api:"nullable"`
	// Comma-separated list of specific pages to process using 1-based indexing.
	// Supports individual pages and ranges. Examples: '1,3,5' (pages 1, 3, 5), '1-5'
	// (pages 1 through 5 inclusive), '1,3,5-8,10' (pages 1, 3, 5-8, and 10). Pages are
	// sorted and deduplicated automatically. Duplicate pages cause an error
	TargetPages string `json:"target_pages" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxPages    respjson.Field
		TargetPages respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersPageRangesResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersPageRangesResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution controls including timeouts and failure thresholds
type ParseV2ParametersProcessingControlResp struct {
	// Quality thresholds that determine when a job should fail vs complete with
	// partial results
	JobFailureConditions ParseV2ParametersProcessingControlJobFailureConditionsResp `json:"job_failure_conditions"`
	// Timeout settings for job execution. Increase for large or complex documents
	Timeouts ParseV2ParametersProcessingControlTimeoutsResp `json:"timeouts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobFailureConditions respjson.Field
		Timeouts             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingControlResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersProcessingControlResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quality thresholds that determine when a job should fail vs complete with
// partial results
type ParseV2ParametersProcessingControlJobFailureConditionsResp struct {
	// Maximum ratio of pages allowed to fail before the job fails (0-1). Example: 0.1
	// means job fails if more than 10% of pages fail. Default is 0.05 (5%)
	AllowedPageFailureRatio float64 `json:"allowed_page_failure_ratio" api:"nullable"`
	// Fail the job if a problematic font is detected that may cause incorrect text
	// extraction. Buggy fonts can produce garbled or missing characters
	FailOnBuggyFont bool `json:"fail_on_buggy_font" api:"nullable"`
	// Fail the entire job if any embedded image cannot be extracted. By default, image
	// extraction errors are logged but don't fail the job
	FailOnImageExtractionError bool `json:"fail_on_image_extraction_error" api:"nullable"`
	// Fail the entire job if OCR fails on any image. By default, OCR errors result in
	// empty text for that image
	FailOnImageOcrError bool `json:"fail_on_image_ocr_error" api:"nullable"`
	// Fail the entire job if markdown cannot be reconstructed for any page. By
	// default, failed pages use fallback text extraction
	FailOnMarkdownReconstructionError bool `json:"fail_on_markdown_reconstruction_error" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedPageFailureRatio           respjson.Field
		FailOnBuggyFont                   respjson.Field
		FailOnImageExtractionError        respjson.Field
		FailOnImageOcrError               respjson.Field
		FailOnMarkdownReconstructionError respjson.Field
		ExtraFields                       map[string]respjson.Field
		raw                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingControlJobFailureConditionsResp) RawJSON() string {
	return r.JSON.raw
}
func (r *ParseV2ParametersProcessingControlJobFailureConditionsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Timeout settings for job execution. Increase for large or complex documents
type ParseV2ParametersProcessingControlTimeoutsResp struct {
	// Base timeout for the job in seconds (max 7200 = 2 hours). This is the minimum
	// time allowed regardless of document size
	BaseInSeconds int64 `json:"base_in_seconds" api:"nullable"`
	// Additional timeout per page in seconds (max 300 = 5 minutes). Total timeout =
	// base + (this value × page count)
	ExtraTimePerPageInSeconds int64 `json:"extra_time_per_page_in_seconds" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseInSeconds             respjson.Field
		ExtraTimePerPageInSeconds respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingControlTimeoutsResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersProcessingControlTimeoutsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Document processing options including OCR, table extraction, and chart parsing
type ParseV2ParametersProcessingOptionsResp struct {
	// Use aggressive heuristics to detect table boundaries, even without visible
	// borders. Useful for documents with borderless or complex tables
	AggressiveTableExtraction bool `json:"aggressive_table_extraction" api:"nullable"`
	// Conditional processing rules that apply different parsing options based on page
	// content, document structure, or filename patterns. Each entry defines trigger
	// conditions and the parsing configuration to apply when triggered
	AutoModeConfiguration []ParseV2ParametersProcessingOptionsAutoModeConfigurationResp `json:"auto_mode_configuration" api:"nullable"`
	// Cost optimizer configuration for reducing parsing costs on simpler pages.
	//
	// When enabled, the parser analyzes each page and routes simpler pages to faster,
	// cheaper processing while preserving quality for complex pages. Only works with
	// 'agentic' or 'agentic_plus' tiers.
	CostOptimizer ParseV2ParametersProcessingOptionsCostOptimizerResp `json:"cost_optimizer" api:"nullable"`
	// Disable automatic heuristics including outlined table extraction and adaptive
	// long table handling. Use when heuristics produce incorrect results
	DisableHeuristics bool `json:"disable_heuristics" api:"nullable"`
	// Options for ignoring specific text types (diagonal, hidden, text in images)
	Ignore ParseV2ParametersProcessingOptionsIgnoreResp `json:"ignore"`
	// OCR configuration including language detection settings
	OcrParameters ParseV2ParametersProcessingOptionsOcrParametersResp `json:"ocr_parameters"`
	// Enable AI-powered chart analysis. Modes: 'efficient' (fast, lower cost),
	// 'agentic' (balanced), 'agentic_plus' (highest accuracy). Automatically enables
	// extract_layout and precise_bounding_box when set
	//
	// Any of "agentic", "agentic_plus", "efficient".
	SpecializedChartParsing string `json:"specialized_chart_parsing" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggressiveTableExtraction respjson.Field
		AutoModeConfiguration     respjson.Field
		CostOptimizer             respjson.Field
		DisableHeuristics         respjson.Field
		Ignore                    respjson.Field
		OcrParameters             respjson.Field
		SpecializedChartParsing   respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingOptionsResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersProcessingOptionsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single auto mode rule with trigger conditions and parsing configuration.
//
// Auto mode allows conditional parsing where different configurations are applied
// based on page content, structure, or filename. When triggers match, the
// parsing_conf overrides default settings for that page.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationResp struct {
	// Parsing configuration to apply when trigger conditions are met
	ParsingConf ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfResp `json:"parsing_conf" api:"required"`
	// Single glob pattern to match against filename
	FilenameMatchGlob string `json:"filename_match_glob" api:"nullable"`
	// List of glob patterns to match against filename
	FilenameMatchGlobList []string `json:"filename_match_glob_list" api:"nullable"`
	// Regex pattern to match against filename
	FilenameRegexp string `json:"filename_regexp" api:"nullable"`
	// Regex mode flags (e.g., 'i' for case-insensitive)
	FilenameRegexpMode string `json:"filename_regexp_mode" api:"nullable"`
	// Trigger if page contains a full-page image (scanned page detection)
	FullPageImageInPage bool `json:"full_page_image_in_page" api:"nullable"`
	// Threshold for full page image detection (0.0-1.0, default 0.8)
	FullPageImageInPageThreshold ParseV2ParametersProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnionResp `json:"full_page_image_in_page_threshold" api:"nullable"`
	// Trigger if page contains non-screenshot images
	ImageInPage bool `json:"image_in_page" api:"nullable"`
	// Trigger if page contains this layout element type
	LayoutElementInPage string `json:"layout_element_in_page" api:"nullable"`
	// Confidence threshold for layout element detection
	LayoutElementInPageConfidenceThreshold ParseV2ParametersProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnionResp `json:"layout_element_in_page_confidence_threshold" api:"nullable"`
	// Trigger if page has more than N charts
	PageContainsAtLeastNCharts ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnionResp `json:"page_contains_at_least_n_charts" api:"nullable"`
	// Trigger if page has more than N images
	PageContainsAtLeastNImages ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnionResp `json:"page_contains_at_least_n_images" api:"nullable"`
	// Trigger if page has more than N layout elements
	PageContainsAtLeastNLayoutElements ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnionResp `json:"page_contains_at_least_n_layout_elements" api:"nullable"`
	// Trigger if page has more than N lines
	PageContainsAtLeastNLines ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnionResp `json:"page_contains_at_least_n_lines" api:"nullable"`
	// Trigger if page has more than N links
	PageContainsAtLeastNLinks ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnionResp `json:"page_contains_at_least_n_links" api:"nullable"`
	// Trigger if page has more than N numeric words
	PageContainsAtLeastNNumbers ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnionResp `json:"page_contains_at_least_n_numbers" api:"nullable"`
	// Trigger if page has more than N% numeric words
	PageContainsAtLeastNPercentNumbers ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnionResp `json:"page_contains_at_least_n_percent_numbers" api:"nullable"`
	// Trigger if page has more than N tables
	PageContainsAtLeastNTables ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnionResp `json:"page_contains_at_least_n_tables" api:"nullable"`
	// Trigger if page has more than N words
	PageContainsAtLeastNWords ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnionResp `json:"page_contains_at_least_n_words" api:"nullable"`
	// Trigger if page has fewer than N charts
	PageContainsAtMostNCharts ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnionResp `json:"page_contains_at_most_n_charts" api:"nullable"`
	// Trigger if page has fewer than N images
	PageContainsAtMostNImages ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnionResp `json:"page_contains_at_most_n_images" api:"nullable"`
	// Trigger if page has fewer than N layout elements
	PageContainsAtMostNLayoutElements ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnionResp `json:"page_contains_at_most_n_layout_elements" api:"nullable"`
	// Trigger if page has fewer than N lines
	PageContainsAtMostNLines ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnionResp `json:"page_contains_at_most_n_lines" api:"nullable"`
	// Trigger if page has fewer than N links
	PageContainsAtMostNLinks ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnionResp `json:"page_contains_at_most_n_links" api:"nullable"`
	// Trigger if page has fewer than N numeric words
	PageContainsAtMostNNumbers ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnionResp `json:"page_contains_at_most_n_numbers" api:"nullable"`
	// Trigger if page has fewer than N% numeric words
	PageContainsAtMostNPercentNumbers ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnionResp `json:"page_contains_at_most_n_percent_numbers" api:"nullable"`
	// Trigger if page has fewer than N tables
	PageContainsAtMostNTables ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnionResp `json:"page_contains_at_most_n_tables" api:"nullable"`
	// Trigger if page has fewer than N words
	PageContainsAtMostNWords ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnionResp `json:"page_contains_at_most_n_words" api:"nullable"`
	// Trigger if page has more than N characters
	PageLongerThanNChars ParseV2ParametersProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnionResp `json:"page_longer_than_n_chars" api:"nullable"`
	// Trigger on pages with markdown extraction errors
	PageMdError bool `json:"page_md_error" api:"nullable"`
	// Trigger if page has fewer than N characters
	PageShorterThanNChars ParseV2ParametersProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnionResp `json:"page_shorter_than_n_chars" api:"nullable"`
	// Regex pattern to match in page content
	RegexpInPage string `json:"regexp_in_page" api:"nullable"`
	// Regex mode flags for regexp_in_page
	RegexpInPageMode string `json:"regexp_in_page_mode" api:"nullable"`
	// Trigger if page contains a table
	TableInPage bool `json:"table_in_page" api:"nullable"`
	// Trigger if page text/markdown contains this string
	TextInPage string `json:"text_in_page" api:"nullable"`
	// How to combine multiple trigger conditions: 'and' (all conditions must match,
	// this is the default) or 'or' (any single condition can trigger)
	TriggerMode string `json:"trigger_mode" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ParsingConf                            respjson.Field
		FilenameMatchGlob                      respjson.Field
		FilenameMatchGlobList                  respjson.Field
		FilenameRegexp                         respjson.Field
		FilenameRegexpMode                     respjson.Field
		FullPageImageInPage                    respjson.Field
		FullPageImageInPageThreshold           respjson.Field
		ImageInPage                            respjson.Field
		LayoutElementInPage                    respjson.Field
		LayoutElementInPageConfidenceThreshold respjson.Field
		PageContainsAtLeastNCharts             respjson.Field
		PageContainsAtLeastNImages             respjson.Field
		PageContainsAtLeastNLayoutElements     respjson.Field
		PageContainsAtLeastNLines              respjson.Field
		PageContainsAtLeastNLinks              respjson.Field
		PageContainsAtLeastNNumbers            respjson.Field
		PageContainsAtLeastNPercentNumbers     respjson.Field
		PageContainsAtLeastNTables             respjson.Field
		PageContainsAtLeastNWords              respjson.Field
		PageContainsAtMostNCharts              respjson.Field
		PageContainsAtMostNImages              respjson.Field
		PageContainsAtMostNLayoutElements      respjson.Field
		PageContainsAtMostNLines               respjson.Field
		PageContainsAtMostNLinks               respjson.Field
		PageContainsAtMostNNumbers             respjson.Field
		PageContainsAtMostNPercentNumbers      respjson.Field
		PageContainsAtMostNTables              respjson.Field
		PageContainsAtMostNWords               respjson.Field
		PageLongerThanNChars                   respjson.Field
		PageMdError                            respjson.Field
		PageShorterThanNChars                  respjson.Field
		RegexpInPage                           respjson.Field
		RegexpInPageMode                       respjson.Field
		TableInPage                            respjson.Field
		TextInPage                             respjson.Field
		TriggerMode                            respjson.Field
		ExtraFields                            map[string]respjson.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingOptionsAutoModeConfigurationResp) RawJSON() string {
	return r.JSON.raw
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parsing configuration to apply when trigger conditions are met
type ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfResp struct {
	// Whether to use adaptive long table handling
	AdaptiveLongTable bool `json:"adaptive_long_table" api:"nullable"`
	// Whether to use aggressive table extraction
	AggressiveTableExtraction bool `json:"aggressive_table_extraction" api:"nullable"`
	// Crop box options for auto mode parsing configuration.
	CropBox ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfCropBoxResp `json:"crop_box" api:"nullable"`
	// Custom AI instructions for matched pages. Overrides the base custom_prompt
	CustomPrompt string `json:"custom_prompt" api:"nullable"`
	// Whether to extract layout information
	ExtractLayout bool `json:"extract_layout" api:"nullable"`
	// Whether to use high resolution OCR
	HighResOcr bool `json:"high_res_ocr" api:"nullable"`
	// Ignore options for auto mode parsing configuration.
	Ignore ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfIgnoreResp `json:"ignore" api:"nullable"`
	// Primary language of the document
	Language string `json:"language" api:"nullable"`
	// Whether to use outlined table extraction
	OutlinedTableExtraction bool `json:"outlined_table_extraction" api:"nullable"`
	// Presentation-specific options for auto mode parsing configuration.
	Presentation ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfPresentationResp `json:"presentation" api:"nullable"`
	// Spatial text options for auto mode parsing configuration.
	SpatialText ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfSpatialTextResp `json:"spatial_text" api:"nullable"`
	// Enable specialized chart parsing with the specified mode
	//
	// Any of "agentic", "agentic_plus", "efficient".
	SpecializedChartParsing string `json:"specialized_chart_parsing" api:"nullable"`
	// Override the parsing tier for matched pages. Must be paired with version
	//
	// Any of "agentic", "agentic_plus", "cost_effective", "fast".
	Tier string `json:"tier" api:"nullable"`
	// Version for the override tier. Required when `tier` is set. Use `latest`, or pin
	// one of that tier's dated versions.
	//
	// Current `latest` by tier:
	//
	// - `fast`: `2025-12-11`
	// - `cost_effective`: `2026-06-26`
	// - `agentic`: `2026-06-18`
	// - `agentic_plus`: `2026-06-18`
	//
	// Full list: `GET /api/v2/parse/versions`.
	Version string `json:"version" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdaptiveLongTable         respjson.Field
		AggressiveTableExtraction respjson.Field
		CropBox                   respjson.Field
		CustomPrompt              respjson.Field
		ExtractLayout             respjson.Field
		HighResOcr                respjson.Field
		Ignore                    respjson.Field
		Language                  respjson.Field
		OutlinedTableExtraction   respjson.Field
		Presentation              respjson.Field
		SpatialText               respjson.Field
		SpecializedChartParsing   respjson.Field
		Tier                      respjson.Field
		Version                   respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfResp) RawJSON() string {
	return r.JSON.raw
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Crop box options for auto mode parsing configuration.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfCropBoxResp struct {
	// Bottom boundary of crop box as ratio (0-1)
	Bottom float64 `json:"bottom" api:"nullable"`
	// Left boundary of crop box as ratio (0-1)
	Left float64 `json:"left" api:"nullable"`
	// Right boundary of crop box as ratio (0-1)
	Right float64 `json:"right" api:"nullable"`
	// Top boundary of crop box as ratio (0-1)
	Top float64 `json:"top" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bottom      respjson.Field
		Left        respjson.Field
		Right       respjson.Field
		Top         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfCropBoxResp) RawJSON() string {
	return r.JSON.raw
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfCropBoxResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ignore options for auto mode parsing configuration.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfIgnoreResp struct {
	// Whether to ignore diagonal text in the document
	IgnoreDiagonalText bool `json:"ignore_diagonal_text" api:"nullable"`
	// Whether to ignore hidden text in the document
	IgnoreHiddenText bool `json:"ignore_hidden_text" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IgnoreDiagonalText respjson.Field
		IgnoreHiddenText   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfIgnoreResp) RawJSON() string {
	return r.JSON.raw
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfIgnoreResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Presentation-specific options for auto mode parsing configuration.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfPresentationResp struct {
	// Extract out of bounds content in presentation slides
	OutOfBoundsContent bool `json:"out_of_bounds_content" api:"nullable"`
	// Skip extraction of embedded data for charts in presentation slides
	SkipEmbeddedData bool `json:"skip_embedded_data" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OutOfBoundsContent respjson.Field
		SkipEmbeddedData   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfPresentationResp) RawJSON() string {
	return r.JSON.raw
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfPresentationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spatial text options for auto mode parsing configuration.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfSpatialTextResp struct {
	// Keep column structure intact without unrolling
	DoNotUnrollColumns bool `json:"do_not_unroll_columns" api:"nullable"`
	// Preserve text alignment across page boundaries
	PreserveLayoutAlignmentAcrossPages bool `json:"preserve_layout_alignment_across_pages" api:"nullable"`
	// Include very small text in spatial output
	PreserveVerySmallText bool `json:"preserve_very_small_text" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DoNotUnrollColumns                 respjson.Field
		PreserveLayoutAlignmentAcrossPages respjson.Field
		PreserveVerySmallText              respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfSpatialTextResp) RawJSON() string {
	return r.JSON.raw
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfSpatialTextResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnionResp
// contains all possible properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnionResp struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnionResp) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnionResp
// contains all possible properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnionResp struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnionResp) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParseV2ParametersProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnionResp
// contains all possible properties and values from [int64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfInt OfString]
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnionResp struct {
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfInt    respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnionResp) AsInt() (v int64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnionResp) RawJSON() string {
	return u.JSON.raw
}

func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cost optimizer configuration for reducing parsing costs on simpler pages.
//
// When enabled, the parser analyzes each page and routes simpler pages to faster,
// cheaper processing while preserving quality for complex pages. Only works with
// 'agentic' or 'agentic_plus' tiers.
type ParseV2ParametersProcessingOptionsCostOptimizerResp struct {
	// Enable cost-optimized parsing. Routes simpler pages to faster processing while
	// complex pages use full AI analysis. May reduce speed on some documents.
	// IMPORTANT: Only available with 'agentic' or 'agentic_plus' tiers
	Enable bool `json:"enable" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enable      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingOptionsCostOptimizerResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersProcessingOptionsCostOptimizerResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for ignoring specific text types (diagonal, hidden, text in images)
type ParseV2ParametersProcessingOptionsIgnoreResp struct {
	// Skip text rotated at an angle (not horizontal/vertical). Useful for ignoring
	// watermarks or decorative angled text
	IgnoreDiagonalText bool `json:"ignore_diagonal_text" api:"nullable"`
	// Skip text marked as hidden in the document structure. Some PDFs contain
	// invisible text layers used for accessibility or search indexing
	IgnoreHiddenText bool `json:"ignore_hidden_text" api:"nullable"`
	// Skip OCR text extraction from embedded images. Use when images contain
	// irrelevant text (watermarks, logos) that shouldn't be in the output
	IgnoreTextInImage bool `json:"ignore_text_in_image" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IgnoreDiagonalText respjson.Field
		IgnoreHiddenText   respjson.Field
		IgnoreTextInImage  respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingOptionsIgnoreResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersProcessingOptionsIgnoreResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OCR configuration including language detection settings
type ParseV2ParametersProcessingOptionsOcrParametersResp struct {
	// Languages to use for OCR text recognition. Specify multiple languages if
	// document contains mixed-language content. Order matters - put primary language
	// first. Example: ['en', 'es'] for English with Spanish
	Languages []ParsingLanguages `json:"languages" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Languages   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersProcessingOptionsOcrParametersResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersProcessingOptionsOcrParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook configuration for receiving parsing job notifications.
//
// Webhooks are called when specified events occur during job processing. Configure
// multiple webhook configurations to send to different endpoints.
type ParseV2ParametersWebhookConfigurationResp struct {
	// Events that trigger this webhook. Options: 'parse.success' (job completed),
	// 'parse.error' (job failed), 'parse.partial_success' (some pages failed),
	// 'parse.pending', 'parse.running', 'parse.cancelled'. If not specified, webhook
	// fires for all events
	WebhookEvents []string `json:"webhook_events" api:"nullable"`
	// Custom HTTP headers to include in webhook requests. Use for authentication
	// tokens or custom routing. Example: {'Authorization': 'Bearer xyz'}
	WebhookHeaders map[string]any `json:"webhook_headers" api:"nullable"`
	// Format of the webhook payload body. 'string' (default) sends the payload as a
	// JSON-encoded string; 'json' sends it as a JSON object.
	//
	// Any of "json", "string".
	WebhookOutputFormat string `json:"webhook_output_format" api:"nullable"`
	// Shared signing secret used to sign webhook deliveries. When set, each request
	// includes an HMAC-SHA256 signature of the request body in the 'LC-Signature'
	// header (value 'sha256=<hex>'). Recompute the HMAC over the raw request body with
	// this secret to verify the delivery is authentic.
	WebhookSigningSecret string `json:"webhook_signing_secret" api:"nullable"`
	// HTTPS URL to receive webhook POST requests. Must be publicly accessible
	WebhookURL string `json:"webhook_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WebhookEvents        respjson.Field
		WebhookHeaders       respjson.Field
		WebhookOutputFormat  respjson.Field
		WebhookSigningSecret respjson.Field
		WebhookURL           respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParseV2ParametersWebhookConfigurationResp) RawJSON() string { return r.JSON.raw }
func (r *ParseV2ParametersWebhookConfigurationResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for LlamaParse v2 document parsing.
//
// Includes tier selection, processing options, output formatting, page targeting,
// and webhook delivery. Refer to the LlamaParse documentation for details on each
// field.
//
// The properties ProductType, Tier, Version are required.
type ParseV2Parameters struct {
	// Parsing tier: 'fast' (rule-based, cheapest), 'cost_effective' (balanced),
	// 'agentic' (AI-powered with custom prompts), or 'agentic_plus' (premium AI with
	// highest accuracy)
	//
	// Any of "agentic", "agentic_plus", "cost_effective", "fast".
	Tier ParseV2ParametersTier `json:"tier,omitzero" api:"required"`
	// Version for the selected tier. Use `latest`, or pin one of that tier's dated
	// versions.
	//
	// Current `latest` by tier:
	//
	// - `fast`: `2025-12-11`
	// - `cost_effective`: `2026-06-26`
	// - `agentic`: `2026-06-18`
	// - `agentic_plus`: `2026-06-18`
	//
	// Full list: `GET /api/v2/parse/versions`.
	Version ParseV2ParametersVersion `json:"version,omitzero" api:"required"`
	// Identifier for the client/application making the request. Used for analytics and
	// debugging. Example: 'my-app-v2'
	ClientName param.Opt[string] `json:"client_name,omitzero"`
	// Bypass result caching and force re-parsing. Use when document content may have
	// changed or you need fresh results
	DisableCache param.Opt[bool] `json:"disable_cache,omitzero"`
	// Options for AI-powered parsing tiers (cost_effective, agentic, agentic_plus).
	//
	// These options customize how the AI processes and interprets document content.
	// Only applicable when using non-fast tiers.
	AgenticOptions ParseV2ParametersAgenticOptions `json:"agentic_options,omitzero"`
	// Options for fast tier parsing (rule-based, no AI).
	//
	// Fast tier uses deterministic algorithms for text extraction without AI
	// enhancement. It's the fastest and most cost-effective option, best suited for
	// simple documents with standard layouts. Currently has no configurable options
	// but reserved for future expansion.
	FastOptions any `json:"fast_options,omitzero"`
	// IDs of saved webhook configurations to notify for this job.
	WebhookConfigurationIDs []string `json:"webhook_configuration_ids,omitzero"`
	// Crop boundaries to process only a portion of each page. Values are ratios 0-1
	// from page edges
	CropBox ParseV2ParametersCropBox `json:"crop_box,omitzero"`
	// Format-specific options (HTML, PDF, spreadsheet, presentation). Applied based on
	// detected input file type
	InputOptions ParseV2ParametersInputOptions `json:"input_options,omitzero"`
	// Output formatting options for markdown, text, and extracted images
	OutputOptions ParseV2ParametersOutputOptions `json:"output_options,omitzero"`
	// Page selection: limit total pages or specify exact pages to process
	PageRanges ParseV2ParametersPageRanges `json:"page_ranges,omitzero"`
	// Job execution controls including timeouts and failure thresholds
	ProcessingControl ParseV2ParametersProcessingControl `json:"processing_control,omitzero"`
	// Document processing options including OCR, table extraction, and chart parsing
	ProcessingOptions ParseV2ParametersProcessingOptions `json:"processing_options,omitzero"`
	// Webhook endpoints for job status notifications. Multiple webhooks can be
	// configured for different events or services
	WebhookConfigurations []ParseV2ParametersWebhookConfiguration `json:"webhook_configurations,omitzero"`
	// Product type.
	//
	// This field can be elided, and will marshal its zero value as "parse_v2".
	ProductType constant.ParseV2 `json:"product_type" default:"parse_v2"`
	paramObj
}

func (r ParseV2Parameters) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2Parameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2Parameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for AI-powered parsing tiers (cost_effective, agentic, agentic_plus).
//
// These options customize how the AI processes and interprets document content.
// Only applicable when using non-fast tiers.
type ParseV2ParametersAgenticOptions struct {
	// Custom instructions for the AI parser. Use to guide extraction behavior, specify
	// output formatting, or provide domain-specific context. Example: 'Extract
	// financial tables with currency symbols. Format dates as YYYY-MM-DD.'
	CustomPrompt param.Opt[string] `json:"custom_prompt,omitzero"`
	paramObj
}

func (r ParseV2ParametersAgenticOptions) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersAgenticOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersAgenticOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Crop boundaries to process only a portion of each page. Values are ratios 0-1
// from page edges
type ParseV2ParametersCropBox struct {
	// Bottom boundary as ratio (0-1). 0=top edge, 1=bottom edge. Content below this
	// line is excluded
	Bottom param.Opt[float64] `json:"bottom,omitzero"`
	// Left boundary as ratio (0-1). 0=left edge, 1=right edge. Content left of this
	// line is excluded
	Left param.Opt[float64] `json:"left,omitzero"`
	// Right boundary as ratio (0-1). 0=left edge, 1=right edge. Content right of this
	// line is excluded
	Right param.Opt[float64] `json:"right,omitzero"`
	// Top boundary as ratio (0-1). 0=top edge, 1=bottom edge. Content above this line
	// is excluded
	Top param.Opt[float64] `json:"top,omitzero"`
	paramObj
}

func (r ParseV2ParametersCropBox) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersCropBox
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersCropBox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format-specific options (HTML, PDF, spreadsheet, presentation). Applied based on
// detected input file type
type ParseV2ParametersInputOptions struct {
	// HTML/web page parsing options (applies to .html, .htm files)
	HTML ParseV2ParametersInputOptionsHTML `json:"html,omitzero"`
	// Image parsing options (applies to .jpg, .jpeg, .png, .webp files)
	Image ParseV2ParametersInputOptionsImage `json:"image,omitzero"`
	// PDF-specific parsing options (applies to .pdf files)
	Pdf any `json:"pdf,omitzero"`
	// Presentation parsing options (applies to .pptx, .ppt, .odp, .key files)
	Presentation ParseV2ParametersInputOptionsPresentation `json:"presentation,omitzero"`
	// Spreadsheet parsing options (applies to .xlsx, .xls, .csv, .ods files)
	Spreadsheet ParseV2ParametersInputOptionsSpreadsheet `json:"spreadsheet,omitzero"`
	paramObj
}

func (r ParseV2ParametersInputOptions) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersInputOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersInputOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTML/web page parsing options (applies to .html, .htm files)
type ParseV2ParametersInputOptionsHTML struct {
	// Force all HTML elements to be visible by overriding CSS display/visibility
	// properties. Useful for parsing pages with hidden content or collapsed sections
	MakeAllElementsVisible param.Opt[bool] `json:"make_all_elements_visible,omitzero"`
	// Remove fixed-position elements (headers, footers, floating buttons) that appear
	// on every page render
	RemoveFixedElements param.Opt[bool] `json:"remove_fixed_elements,omitzero"`
	// Remove navigation elements (nav bars, sidebars, menus) to focus on main content
	RemoveNavigationElements param.Opt[bool] `json:"remove_navigation_elements,omitzero"`
	paramObj
}

func (r ParseV2ParametersInputOptionsHTML) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersInputOptionsHTML
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersInputOptionsHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image parsing options (applies to .jpg, .jpeg, .png, .webp files)
type ParseV2ParametersInputOptionsImage struct {
	// Detect documents photographed with a camera (e.g. phone scans of receipts or
	// forms), then crop, perspective-correct, and flatten uneven lighting and shadows
	// before parsing. Supports JPEG, PNG, WebP, and HEIC/HEIF inputs. Improves results
	// when the document is tilted or surrounded by background. Images that already
	// look like clean scans are left untouched
	CameraPhotoCorrection param.Opt[bool] `json:"camera_photo_correction,omitzero"`
	paramObj
}

func (r ParseV2ParametersInputOptionsImage) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersInputOptionsImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersInputOptionsImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Presentation parsing options (applies to .pptx, .ppt, .odp, .key files)
type ParseV2ParametersInputOptionsPresentation struct {
	// Extract content positioned outside the visible slide area. Some presentations
	// have hidden notes or content that extends beyond slide boundaries
	OutOfBoundsContent param.Opt[bool] `json:"out_of_bounds_content,omitzero"`
	// Skip extraction of embedded chart data tables. When true, only the visual
	// representation of charts is captured, not the underlying data
	SkipEmbeddedData param.Opt[bool] `json:"skip_embedded_data,omitzero"`
	paramObj
}

func (r ParseV2ParametersInputOptionsPresentation) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersInputOptionsPresentation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersInputOptionsPresentation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spreadsheet parsing options (applies to .xlsx, .xls, .csv, .ods files)
type ParseV2ParametersInputOptionsSpreadsheet struct {
	// Detect and extract multiple tables within a single sheet. Useful when
	// spreadsheets contain several data regions separated by blank rows/columns
	DetectSubTablesInSheets param.Opt[bool] `json:"detect_sub_tables_in_sheets,omitzero"`
	// Compute formula results instead of extracting formula text. Use when you need
	// calculated values rather than formula definitions
	ForceFormulaComputationInSheets param.Opt[bool] `json:"force_formula_computation_in_sheets,omitzero"`
	// Parse hidden sheets in addition to visible ones. By default, hidden sheets are
	// skipped
	IncludeHiddenSheets param.Opt[bool] `json:"include_hidden_sheets,omitzero"`
	paramObj
}

func (r ParseV2ParametersInputOptionsSpreadsheet) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersInputOptionsSpreadsheet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersInputOptionsSpreadsheet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Output formatting options for markdown, text, and extracted images
type ParseV2ParametersOutputOptions struct {
	// Extract the printed page number as it appears in the document (e.g., 'Page 5 of
	// 10', 'v', 'A-3'). Useful for referencing original page numbers
	ExtractPrintedPageNumber param.Opt[bool] `json:"extract_printed_page_number,omitzero"`
	// Optional additional output artifacts to save alongside the primary parse output.
	// Each value opts in to generating and persisting one extra file; the empty list
	// (default) saves none. The three accepted values are: 'stripped_md' — per-page
	// markdown stripped of formatting (links, bold/italic, images, HTML), saved as
	// JSON for full-text-search indexing; fetch via
	// `expand=stripped_markdown_content_metadata`. 'concatenated_stripped_txt' — all
	// stripped pages concatenated into a single plain-text file with `\n\n---\n\n`
	// between pages, useful for feeding the document into search or embedding
	// pipelines as one blob; fetch via
	// `expand=concatenated_stripped_markdown_content_metadata`. 'word_bbox' — raw
	// word-level bounding boxes (one JSON object per word, with page number and
	// x/y/w/h coordinates) saved as JSONL, useful for highlighting or grounding
	// extracted answers back to the source document; fetch via
	// `expand=raw_words_content_metadata`.
	AdditionalOutputs []string `json:"additional_outputs,omitzero"`
	// Bounding-box granularity levels to compute for the parse. 'word' computes one
	// bounding box per detected word; 'line' computes one per text line; 'cell'
	// computes one per table cell. Multiple levels can be requested. Empty list
	// (default) disables granular bboxes — only item-level layout boxes are returned
	// on the result. When set, the computed boxes are not inlined on the result items;
	// they are written to a separate `grounded_items` sidecar (JSONL, one row per
	// page) and exposed as `result_content_metadata.grounded_items` (a presigned
	// download URL) on the parse result. Each row matches the `GroundedJsonItem`
	// shape.
	//
	// Any of "cell", "line", "word".
	GranularBboxes []string `json:"granular_bboxes,omitzero"`
	// Image categories to extract and save. Options: 'screenshot' (full page renders
	// useful for visual QA), 'embedded' (images found within the document), 'layout'
	// (cropped regions from layout detection like figures and diagrams). Empty list
	// saves no images
	//
	// Any of "embedded", "layout", "screenshot".
	ImagesToSave []string `json:"images_to_save,omitzero"`
	// Markdown formatting options including table styles and link annotations
	Markdown ParseV2ParametersOutputOptionsMarkdown `json:"markdown,omitzero"`
	// Spatial text output options for preserving document layout structure
	SpatialText ParseV2ParametersOutputOptionsSpatialText `json:"spatial_text,omitzero"`
	// Options for exporting tables as XLSX spreadsheets
	TablesAsSpreadsheet ParseV2ParametersOutputOptionsTablesAsSpreadsheet `json:"tables_as_spreadsheet,omitzero"`
	paramObj
}

func (r ParseV2ParametersOutputOptions) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersOutputOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersOutputOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown formatting options including table styles and link annotations
type ParseV2ParametersOutputOptionsMarkdown struct {
	// Add link annotations to markdown output in the format [text](url). When false,
	// only the link text is included
	AnnotateLinks param.Opt[bool] `json:"annotate_links,omitzero"`
	// Embed images directly in markdown as base64 data URIs instead of extracting them
	// as separate files. Useful for self-contained markdown output
	InlineImages param.Opt[bool] `json:"inline_images,omitzero"`
	// Table formatting options including markdown vs HTML format and merging behavior
	Tables ParseV2ParametersOutputOptionsMarkdownTables `json:"tables,omitzero"`
	paramObj
}

func (r ParseV2ParametersOutputOptionsMarkdown) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersOutputOptionsMarkdown
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersOutputOptionsMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Table formatting options including markdown vs HTML format and merging behavior
type ParseV2ParametersOutputOptionsMarkdownTables struct {
	// Remove extra whitespace padding in markdown table cells for more compact output
	CompactMarkdownTables param.Opt[bool] `json:"compact_markdown_tables,omitzero"`
	// Separator string for multiline cell content in markdown tables. Example:
	// '&lt;br&gt;' to preserve line breaks, ' ' to join with spaces
	MarkdownTableMultilineSeparator param.Opt[string] `json:"markdown_table_multiline_separator,omitzero"`
	// Automatically merge tables that span multiple pages into a single table. The
	// merged table appears on the first page with merged_from_pages metadata
	MergeContinuedTables param.Opt[bool] `json:"merge_continued_tables,omitzero"`
	// Output tables as markdown pipe tables instead of HTML &lt;table&gt; tags.
	// Markdown tables are simpler but cannot represent complex structures like merged
	// cells
	OutputTablesAsMarkdown param.Opt[bool] `json:"output_tables_as_markdown,omitzero"`
	paramObj
}

func (r ParseV2ParametersOutputOptionsMarkdownTables) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersOutputOptionsMarkdownTables
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersOutputOptionsMarkdownTables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spatial text output options for preserving document layout structure
type ParseV2ParametersOutputOptionsSpatialText struct {
	// Keep multi-column layouts intact instead of linearizing columns into sequential
	// text. Automatically enabled for non-fast tiers
	DoNotUnrollColumns param.Opt[bool] `json:"do_not_unroll_columns,omitzero"`
	// Maintain consistent text column alignment across page boundaries. Automatically
	// enabled for document-level parsing modes
	PreserveLayoutAlignmentAcrossPages param.Opt[bool] `json:"preserve_layout_alignment_across_pages,omitzero"`
	// Include text below the normal size threshold. Useful for footnotes, watermarks,
	// or fine print that might otherwise be filtered out
	PreserveVerySmallText param.Opt[bool] `json:"preserve_very_small_text,omitzero"`
	paramObj
}

func (r ParseV2ParametersOutputOptionsSpatialText) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersOutputOptionsSpatialText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersOutputOptionsSpatialText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for exporting tables as XLSX spreadsheets
type ParseV2ParametersOutputOptionsTablesAsSpreadsheet struct {
	// Whether this option is enabled
	Enable param.Opt[bool] `json:"enable,omitzero"`
	// Automatically generate descriptive sheet names from table context (headers,
	// surrounding text) instead of using generic names like 'Table_1'
	GuessSheetName param.Opt[bool] `json:"guess_sheet_name,omitzero"`
	paramObj
}

func (r ParseV2ParametersOutputOptionsTablesAsSpreadsheet) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersOutputOptionsTablesAsSpreadsheet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersOutputOptionsTablesAsSpreadsheet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page selection: limit total pages or specify exact pages to process
type ParseV2ParametersPageRanges struct {
	// Maximum number of pages to process. Pages are processed in order starting from
	// page 1. If both max_pages and target_pages are set, target_pages takes
	// precedence
	MaxPages param.Opt[int64] `json:"max_pages,omitzero"`
	// Comma-separated list of specific pages to process using 1-based indexing.
	// Supports individual pages and ranges. Examples: '1,3,5' (pages 1, 3, 5), '1-5'
	// (pages 1 through 5 inclusive), '1,3,5-8,10' (pages 1, 3, 5-8, and 10). Pages are
	// sorted and deduplicated automatically. Duplicate pages cause an error
	TargetPages param.Opt[string] `json:"target_pages,omitzero"`
	paramObj
}

func (r ParseV2ParametersPageRanges) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersPageRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersPageRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution controls including timeouts and failure thresholds
type ParseV2ParametersProcessingControl struct {
	// Quality thresholds that determine when a job should fail vs complete with
	// partial results
	JobFailureConditions ParseV2ParametersProcessingControlJobFailureConditions `json:"job_failure_conditions,omitzero"`
	// Timeout settings for job execution. Increase for large or complex documents
	Timeouts ParseV2ParametersProcessingControlTimeouts `json:"timeouts,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingControl) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingControl
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quality thresholds that determine when a job should fail vs complete with
// partial results
type ParseV2ParametersProcessingControlJobFailureConditions struct {
	// Maximum ratio of pages allowed to fail before the job fails (0-1). Example: 0.1
	// means job fails if more than 10% of pages fail. Default is 0.05 (5%)
	AllowedPageFailureRatio param.Opt[float64] `json:"allowed_page_failure_ratio,omitzero"`
	// Fail the job if a problematic font is detected that may cause incorrect text
	// extraction. Buggy fonts can produce garbled or missing characters
	FailOnBuggyFont param.Opt[bool] `json:"fail_on_buggy_font,omitzero"`
	// Fail the entire job if any embedded image cannot be extracted. By default, image
	// extraction errors are logged but don't fail the job
	FailOnImageExtractionError param.Opt[bool] `json:"fail_on_image_extraction_error,omitzero"`
	// Fail the entire job if OCR fails on any image. By default, OCR errors result in
	// empty text for that image
	FailOnImageOcrError param.Opt[bool] `json:"fail_on_image_ocr_error,omitzero"`
	// Fail the entire job if markdown cannot be reconstructed for any page. By
	// default, failed pages use fallback text extraction
	FailOnMarkdownReconstructionError param.Opt[bool] `json:"fail_on_markdown_reconstruction_error,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingControlJobFailureConditions) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingControlJobFailureConditions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingControlJobFailureConditions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Timeout settings for job execution. Increase for large or complex documents
type ParseV2ParametersProcessingControlTimeouts struct {
	// Base timeout for the job in seconds (max 7200 = 2 hours). This is the minimum
	// time allowed regardless of document size
	BaseInSeconds param.Opt[int64] `json:"base_in_seconds,omitzero"`
	// Additional timeout per page in seconds (max 300 = 5 minutes). Total timeout =
	// base + (this value × page count)
	ExtraTimePerPageInSeconds param.Opt[int64] `json:"extra_time_per_page_in_seconds,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingControlTimeouts) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingControlTimeouts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingControlTimeouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Document processing options including OCR, table extraction, and chart parsing
type ParseV2ParametersProcessingOptions struct {
	// Use aggressive heuristics to detect table boundaries, even without visible
	// borders. Useful for documents with borderless or complex tables
	AggressiveTableExtraction param.Opt[bool] `json:"aggressive_table_extraction,omitzero"`
	// Disable automatic heuristics including outlined table extraction and adaptive
	// long table handling. Use when heuristics produce incorrect results
	DisableHeuristics param.Opt[bool] `json:"disable_heuristics,omitzero"`
	// Conditional processing rules that apply different parsing options based on page
	// content, document structure, or filename patterns. Each entry defines trigger
	// conditions and the parsing configuration to apply when triggered
	AutoModeConfiguration []ParseV2ParametersProcessingOptionsAutoModeConfiguration `json:"auto_mode_configuration,omitzero"`
	// Cost optimizer configuration for reducing parsing costs on simpler pages.
	//
	// When enabled, the parser analyzes each page and routes simpler pages to faster,
	// cheaper processing while preserving quality for complex pages. Only works with
	// 'agentic' or 'agentic_plus' tiers.
	CostOptimizer ParseV2ParametersProcessingOptionsCostOptimizer `json:"cost_optimizer,omitzero"`
	// Enable AI-powered chart analysis. Modes: 'efficient' (fast, lower cost),
	// 'agentic' (balanced), 'agentic_plus' (highest accuracy). Automatically enables
	// extract_layout and precise_bounding_box when set
	//
	// Any of "agentic", "agentic_plus", "efficient".
	SpecializedChartParsing string `json:"specialized_chart_parsing,omitzero"`
	// Options for ignoring specific text types (diagonal, hidden, text in images)
	Ignore ParseV2ParametersProcessingOptionsIgnore `json:"ignore,omitzero"`
	// OCR configuration including language detection settings
	OcrParameters ParseV2ParametersProcessingOptionsOcrParameters `json:"ocr_parameters,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingOptions) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ParseV2ParametersProcessingOptions](
		"specialized_chart_parsing", "agentic", "agentic_plus", "efficient",
	)
}

// A single auto mode rule with trigger conditions and parsing configuration.
//
// Auto mode allows conditional parsing where different configurations are applied
// based on page content, structure, or filename. When triggers match, the
// parsing_conf overrides default settings for that page.
//
// The property ParsingConf is required.
type ParseV2ParametersProcessingOptionsAutoModeConfiguration struct {
	// Parsing configuration to apply when trigger conditions are met
	ParsingConf ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConf `json:"parsing_conf,omitzero" api:"required"`
	// Single glob pattern to match against filename
	FilenameMatchGlob param.Opt[string] `json:"filename_match_glob,omitzero"`
	// Regex pattern to match against filename
	FilenameRegexp param.Opt[string] `json:"filename_regexp,omitzero"`
	// Regex mode flags (e.g., 'i' for case-insensitive)
	FilenameRegexpMode param.Opt[string] `json:"filename_regexp_mode,omitzero"`
	// Trigger if page contains a full-page image (scanned page detection)
	FullPageImageInPage param.Opt[bool] `json:"full_page_image_in_page,omitzero"`
	// Trigger if page contains non-screenshot images
	ImageInPage param.Opt[bool] `json:"image_in_page,omitzero"`
	// Trigger if page contains this layout element type
	LayoutElementInPage param.Opt[string] `json:"layout_element_in_page,omitzero"`
	// Trigger on pages with markdown extraction errors
	PageMdError param.Opt[bool] `json:"page_md_error,omitzero"`
	// Regex pattern to match in page content
	RegexpInPage param.Opt[string] `json:"regexp_in_page,omitzero"`
	// Regex mode flags for regexp_in_page
	RegexpInPageMode param.Opt[string] `json:"regexp_in_page_mode,omitzero"`
	// Trigger if page contains a table
	TableInPage param.Opt[bool] `json:"table_in_page,omitzero"`
	// Trigger if page text/markdown contains this string
	TextInPage param.Opt[string] `json:"text_in_page,omitzero"`
	// How to combine multiple trigger conditions: 'and' (all conditions must match,
	// this is the default) or 'or' (any single condition can trigger)
	TriggerMode param.Opt[string] `json:"trigger_mode,omitzero"`
	// List of glob patterns to match against filename
	FilenameMatchGlobList []string `json:"filename_match_glob_list,omitzero"`
	// Threshold for full page image detection (0.0-1.0, default 0.8)
	FullPageImageInPageThreshold ParseV2ParametersProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnion `json:"full_page_image_in_page_threshold,omitzero"`
	// Confidence threshold for layout element detection
	LayoutElementInPageConfidenceThreshold ParseV2ParametersProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnion `json:"layout_element_in_page_confidence_threshold,omitzero"`
	// Trigger if page has more than N charts
	PageContainsAtLeastNCharts ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnion `json:"page_contains_at_least_n_charts,omitzero"`
	// Trigger if page has more than N images
	PageContainsAtLeastNImages ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnion `json:"page_contains_at_least_n_images,omitzero"`
	// Trigger if page has more than N layout elements
	PageContainsAtLeastNLayoutElements ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnion `json:"page_contains_at_least_n_layout_elements,omitzero"`
	// Trigger if page has more than N lines
	PageContainsAtLeastNLines ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnion `json:"page_contains_at_least_n_lines,omitzero"`
	// Trigger if page has more than N links
	PageContainsAtLeastNLinks ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnion `json:"page_contains_at_least_n_links,omitzero"`
	// Trigger if page has more than N numeric words
	PageContainsAtLeastNNumbers ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnion `json:"page_contains_at_least_n_numbers,omitzero"`
	// Trigger if page has more than N% numeric words
	PageContainsAtLeastNPercentNumbers ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnion `json:"page_contains_at_least_n_percent_numbers,omitzero"`
	// Trigger if page has more than N tables
	PageContainsAtLeastNTables ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnion `json:"page_contains_at_least_n_tables,omitzero"`
	// Trigger if page has more than N words
	PageContainsAtLeastNWords ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnion `json:"page_contains_at_least_n_words,omitzero"`
	// Trigger if page has fewer than N charts
	PageContainsAtMostNCharts ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnion `json:"page_contains_at_most_n_charts,omitzero"`
	// Trigger if page has fewer than N images
	PageContainsAtMostNImages ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnion `json:"page_contains_at_most_n_images,omitzero"`
	// Trigger if page has fewer than N layout elements
	PageContainsAtMostNLayoutElements ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnion `json:"page_contains_at_most_n_layout_elements,omitzero"`
	// Trigger if page has fewer than N lines
	PageContainsAtMostNLines ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnion `json:"page_contains_at_most_n_lines,omitzero"`
	// Trigger if page has fewer than N links
	PageContainsAtMostNLinks ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnion `json:"page_contains_at_most_n_links,omitzero"`
	// Trigger if page has fewer than N numeric words
	PageContainsAtMostNNumbers ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnion `json:"page_contains_at_most_n_numbers,omitzero"`
	// Trigger if page has fewer than N% numeric words
	PageContainsAtMostNPercentNumbers ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnion `json:"page_contains_at_most_n_percent_numbers,omitzero"`
	// Trigger if page has fewer than N tables
	PageContainsAtMostNTables ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnion `json:"page_contains_at_most_n_tables,omitzero"`
	// Trigger if page has fewer than N words
	PageContainsAtMostNWords ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnion `json:"page_contains_at_most_n_words,omitzero"`
	// Trigger if page has more than N characters
	PageLongerThanNChars ParseV2ParametersProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnion `json:"page_longer_than_n_chars,omitzero"`
	// Trigger if page has fewer than N characters
	PageShorterThanNChars ParseV2ParametersProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnion `json:"page_shorter_than_n_chars,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingOptionsAutoModeConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingOptionsAutoModeConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parsing configuration to apply when trigger conditions are met
type ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConf struct {
	// Whether to use adaptive long table handling
	AdaptiveLongTable param.Opt[bool] `json:"adaptive_long_table,omitzero"`
	// Whether to use aggressive table extraction
	AggressiveTableExtraction param.Opt[bool] `json:"aggressive_table_extraction,omitzero"`
	// Custom AI instructions for matched pages. Overrides the base custom_prompt
	CustomPrompt param.Opt[string] `json:"custom_prompt,omitzero"`
	// Whether to extract layout information
	ExtractLayout param.Opt[bool] `json:"extract_layout,omitzero"`
	// Whether to use high resolution OCR
	HighResOcr param.Opt[bool] `json:"high_res_ocr,omitzero"`
	// Primary language of the document
	Language param.Opt[string] `json:"language,omitzero"`
	// Whether to use outlined table extraction
	OutlinedTableExtraction param.Opt[bool] `json:"outlined_table_extraction,omitzero"`
	// Crop box options for auto mode parsing configuration.
	CropBox ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfCropBox `json:"crop_box,omitzero"`
	// Ignore options for auto mode parsing configuration.
	Ignore ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfIgnore `json:"ignore,omitzero"`
	// Presentation-specific options for auto mode parsing configuration.
	Presentation ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfPresentation `json:"presentation,omitzero"`
	// Spatial text options for auto mode parsing configuration.
	SpatialText ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfSpatialText `json:"spatial_text,omitzero"`
	// Enable specialized chart parsing with the specified mode
	//
	// Any of "agentic", "agentic_plus", "efficient".
	SpecializedChartParsing string `json:"specialized_chart_parsing,omitzero"`
	// Override the parsing tier for matched pages. Must be paired with version
	//
	// Any of "agentic", "agentic_plus", "cost_effective", "fast".
	Tier string `json:"tier,omitzero"`
	// Version for the override tier. Required when `tier` is set. Use `latest`, or pin
	// one of that tier's dated versions.
	//
	// Current `latest` by tier:
	//
	// - `fast`: `2025-12-11`
	// - `cost_effective`: `2026-06-26`
	// - `agentic`: `2026-06-18`
	// - `agentic_plus`: `2026-06-18`
	//
	// Full list: `GET /api/v2/parse/versions`.
	Version string `json:"version,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConf) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConf](
		"specialized_chart_parsing", "agentic", "agentic_plus", "efficient",
	)
	apijson.RegisterFieldValidator[ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConf](
		"tier", "agentic", "agentic_plus", "cost_effective", "fast",
	)
}

// Crop box options for auto mode parsing configuration.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfCropBox struct {
	// Bottom boundary of crop box as ratio (0-1)
	Bottom param.Opt[float64] `json:"bottom,omitzero"`
	// Left boundary of crop box as ratio (0-1)
	Left param.Opt[float64] `json:"left,omitzero"`
	// Right boundary of crop box as ratio (0-1)
	Right param.Opt[float64] `json:"right,omitzero"`
	// Top boundary of crop box as ratio (0-1)
	Top param.Opt[float64] `json:"top,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfCropBox) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfCropBox
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfCropBox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ignore options for auto mode parsing configuration.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfIgnore struct {
	// Whether to ignore diagonal text in the document
	IgnoreDiagonalText param.Opt[bool] `json:"ignore_diagonal_text,omitzero"`
	// Whether to ignore hidden text in the document
	IgnoreHiddenText param.Opt[bool] `json:"ignore_hidden_text,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfIgnore) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfIgnore
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfIgnore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Presentation-specific options for auto mode parsing configuration.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfPresentation struct {
	// Extract out of bounds content in presentation slides
	OutOfBoundsContent param.Opt[bool] `json:"out_of_bounds_content,omitzero"`
	// Skip extraction of embedded data for charts in presentation slides
	SkipEmbeddedData param.Opt[bool] `json:"skip_embedded_data,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfPresentation) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfPresentation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfPresentation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spatial text options for auto mode parsing configuration.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfSpatialText struct {
	// Keep column structure intact without unrolling
	DoNotUnrollColumns param.Opt[bool] `json:"do_not_unroll_columns,omitzero"`
	// Preserve text alignment across page boundaries
	PreserveLayoutAlignmentAcrossPages param.Opt[bool] `json:"preserve_layout_alignment_across_pages,omitzero"`
	// Include very small text in spatial output
	PreserveVerySmallText param.Opt[bool] `json:"preserve_very_small_text,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfSpatialText) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfSpatialText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingOptionsAutoModeConfigurationParsingConfSpatialText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParseV2ParametersProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParseV2ParametersProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParseV2ParametersProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Cost optimizer configuration for reducing parsing costs on simpler pages.
//
// When enabled, the parser analyzes each page and routes simpler pages to faster,
// cheaper processing while preserving quality for complex pages. Only works with
// 'agentic' or 'agentic_plus' tiers.
type ParseV2ParametersProcessingOptionsCostOptimizer struct {
	// Enable cost-optimized parsing. Routes simpler pages to faster processing while
	// complex pages use full AI analysis. May reduce speed on some documents.
	// IMPORTANT: Only available with 'agentic' or 'agentic_plus' tiers
	Enable param.Opt[bool] `json:"enable,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingOptionsCostOptimizer) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingOptionsCostOptimizer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingOptionsCostOptimizer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for ignoring specific text types (diagonal, hidden, text in images)
type ParseV2ParametersProcessingOptionsIgnore struct {
	// Skip text rotated at an angle (not horizontal/vertical). Useful for ignoring
	// watermarks or decorative angled text
	IgnoreDiagonalText param.Opt[bool] `json:"ignore_diagonal_text,omitzero"`
	// Skip text marked as hidden in the document structure. Some PDFs contain
	// invisible text layers used for accessibility or search indexing
	IgnoreHiddenText param.Opt[bool] `json:"ignore_hidden_text,omitzero"`
	// Skip OCR text extraction from embedded images. Use when images contain
	// irrelevant text (watermarks, logos) that shouldn't be in the output
	IgnoreTextInImage param.Opt[bool] `json:"ignore_text_in_image,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingOptionsIgnore) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingOptionsIgnore
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingOptionsIgnore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OCR configuration including language detection settings
type ParseV2ParametersProcessingOptionsOcrParameters struct {
	// Languages to use for OCR text recognition. Specify multiple languages if
	// document contains mixed-language content. Order matters - put primary language
	// first. Example: ['en', 'es'] for English with Spanish
	Languages []ParsingLanguages `json:"languages,omitzero"`
	paramObj
}

func (r ParseV2ParametersProcessingOptionsOcrParameters) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersProcessingOptionsOcrParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersProcessingOptionsOcrParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook configuration for receiving parsing job notifications.
//
// Webhooks are called when specified events occur during job processing. Configure
// multiple webhook configurations to send to different endpoints.
type ParseV2ParametersWebhookConfiguration struct {
	// Shared signing secret used to sign webhook deliveries. When set, each request
	// includes an HMAC-SHA256 signature of the request body in the 'LC-Signature'
	// header (value 'sha256=<hex>'). Recompute the HMAC over the raw request body with
	// this secret to verify the delivery is authentic.
	WebhookSigningSecret param.Opt[string] `json:"webhook_signing_secret,omitzero"`
	// HTTPS URL to receive webhook POST requests. Must be publicly accessible
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// Events that trigger this webhook. Options: 'parse.success' (job completed),
	// 'parse.error' (job failed), 'parse.partial_success' (some pages failed),
	// 'parse.pending', 'parse.running', 'parse.cancelled'. If not specified, webhook
	// fires for all events
	WebhookEvents []string `json:"webhook_events,omitzero"`
	// Custom HTTP headers to include in webhook requests. Use for authentication
	// tokens or custom routing. Example: {'Authorization': 'Bearer xyz'}
	WebhookHeaders map[string]any `json:"webhook_headers,omitzero"`
	// Format of the webhook payload body. 'string' (default) sends the payload as a
	// JSON-encoded string; 'json' sends it as a JSON object.
	//
	// Any of "json", "string".
	WebhookOutputFormat string `json:"webhook_output_format,omitzero"`
	paramObj
}

func (r ParseV2ParametersWebhookConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow ParseV2ParametersWebhookConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParseV2ParametersWebhookConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ParseV2ParametersWebhookConfiguration](
		"webhook_output_format", "json", "string",
	)
}

// Typed parameters for a _split v1_ product configuration.
type SplitV1ParametersResp struct {
	// Categories to split documents into.
	Categories []SplitCategory `json:"categories" api:"required"`
	// Product type.
	ProductType constant.SplitV1 `json:"product_type" default:"split_v1"`
	// Strategy for splitting documents.
	SplittingStrategy SplitV1ParametersSplittingStrategyResp `json:"splitting_strategy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories        respjson.Field
		ProductType       respjson.Field
		SplittingStrategy respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitV1ParametersResp) RawJSON() string { return r.JSON.raw }
func (r *SplitV1ParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SplitV1ParametersResp to a SplitV1Parameters.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SplitV1Parameters.Overrides()
func (r SplitV1ParametersResp) ToParam() SplitV1Parameters {
	return param.Override[SplitV1Parameters](json.RawMessage(r.RawJSON()))
}

// Strategy for splitting documents.
type SplitV1ParametersSplittingStrategyResp struct {
	// Controls handling of pages that don't match any category. 'include': pages can
	// be grouped as 'uncategorized' and included in results. 'forbid': all pages must
	// be assigned to a defined category. 'omit': pages can be classified as
	// 'uncategorized' but are excluded from results.
	//
	// Any of "forbid", "include", "omit".
	AllowUncategorized string `json:"allow_uncategorized"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowUncategorized respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SplitV1ParametersSplittingStrategyResp) RawJSON() string { return r.JSON.raw }
func (r *SplitV1ParametersSplittingStrategyResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Typed parameters for a _split v1_ product configuration.
//
// The properties Categories, ProductType are required.
type SplitV1Parameters struct {
	// Categories to split documents into.
	Categories []SplitCategoryParam `json:"categories,omitzero" api:"required"`
	// Strategy for splitting documents.
	SplittingStrategy SplitV1ParametersSplittingStrategy `json:"splitting_strategy,omitzero"`
	// Product type.
	//
	// This field can be elided, and will marshal its zero value as "split_v1".
	ProductType constant.SplitV1 `json:"product_type" default:"split_v1"`
	paramObj
}

func (r SplitV1Parameters) MarshalJSON() (data []byte, err error) {
	type shadow SplitV1Parameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SplitV1Parameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Strategy for splitting documents.
type SplitV1ParametersSplittingStrategy struct {
	// Controls handling of pages that don't match any category. 'include': pages can
	// be grouped as 'uncategorized' and included in results. 'forbid': all pages must
	// be assigned to a defined category. 'omit': pages can be classified as
	// 'uncategorized' but are excluded from results.
	//
	// Any of "forbid", "include", "omit".
	AllowUncategorized string `json:"allow_uncategorized,omitzero"`
	paramObj
}

func (r SplitV1ParametersSplittingStrategy) MarshalJSON() (data []byte, err error) {
	type shadow SplitV1ParametersSplittingStrategy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SplitV1ParametersSplittingStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SplitV1ParametersSplittingStrategy](
		"allow_uncategorized", "forbid", "include", "omit",
	)
}

// Catch-all for configurations without a dedicated typed schema.
//
// Accepts arbitrary JSON fields alongside `product_type`.
type UntypedParametersResp struct {
	// Product type.
	ProductType constant.Unknown `json:"product_type" default:"unknown"`
	ExtraFields map[string]any   `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProductType respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UntypedParametersResp) RawJSON() string { return r.JSON.raw }
func (r *UntypedParametersResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UntypedParametersResp to a UntypedParameters.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UntypedParameters.Overrides()
func (r UntypedParametersResp) ToParam() UntypedParameters {
	return param.Override[UntypedParameters](json.RawMessage(r.RawJSON()))
}

func NewUntypedParameters() UntypedParameters {
	return UntypedParameters{
		ProductType: "unknown",
	}
}

// Catch-all for configurations without a dedicated typed schema.
//
// Accepts arbitrary JSON fields alongside `product_type`.
//
// This struct has a constant value, construct it with [NewUntypedParameters].
type UntypedParameters struct {
	// Product type.
	ProductType constant.Unknown `json:"product_type" default:"unknown"`
	ExtraFields map[string]any   `json:"-"`
	paramObj
}

func (r UntypedParameters) MarshalJSON() (data []byte, err error) {
	type shadow UntypedParameters
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *UntypedParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigurationNewParams struct {
	// Request body for creating a product configuration.
	ConfigurationCreate ConfigurationCreateParam
	OrganizationID      param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID           param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r ConfigurationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ConfigurationCreate)
}
func (r *ConfigurationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ConfigurationNewParams]'s query parameters as `url.Values`.
func (r ConfigurationNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigurationGetParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ConfigurationGetParams]'s query parameters as `url.Values`.
func (r ConfigurationGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigurationUpdateParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Updated name (omit to leave unchanged).
	Name param.Opt[string] `json:"name,omitzero"`
	// Updated parameters (omit to leave unchanged).
	Parameters ConfigurationUpdateParamsParametersUnion `json:"parameters,omitzero"`
	paramObj
}

func (r ConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConfigurationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigurationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ConfigurationUpdateParams]'s query parameters as
// `url.Values`.
func (r ConfigurationUpdateParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConfigurationUpdateParamsParametersUnion struct {
	OfClassifyV2    *ClassifyV2Parameters                             `json:",omitzero,inline"`
	OfExtractV2     *ExtractV2Parameters                              `json:",omitzero,inline"`
	OfParseV2       *ParseV2Parameters                                `json:",omitzero,inline"`
	OfSplitV1       *SplitV1Parameters                                `json:",omitzero,inline"`
	OfSpreadsheetV1 *ConfigurationUpdateParamsParametersSpreadsheetV1 `json:",omitzero,inline"`
	OfUnknown       *UntypedParameters                                `json:",omitzero,inline"`
	paramUnion
}

func (u ConfigurationUpdateParamsParametersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfClassifyV2,
		u.OfExtractV2,
		u.OfParseV2,
		u.OfSplitV1,
		u.OfSpreadsheetV1,
		u.OfUnknown)
}
func (u *ConfigurationUpdateParamsParametersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ConfigurationUpdateParamsParametersUnion](
		"product_type",
		apijson.Discriminator[ClassifyV2Parameters]("classify_v2"),
		apijson.Discriminator[ExtractV2Parameters]("extract_v2"),
		apijson.Discriminator[ParseV2Parameters]("parse_v2"),
		apijson.Discriminator[SplitV1Parameters]("split_v1"),
		apijson.Discriminator[ConfigurationUpdateParamsParametersSpreadsheetV1]("spreadsheet_v1"),
		apijson.Discriminator[UntypedParameters]("unknown"),
	)
}

// Typed parameters for a _spreadsheet v1_ product configuration.
//
// The property ProductType is required.
type ConfigurationUpdateParamsParametersSpreadsheetV1 struct {
	// A1 notation of the range to extract a single region from. If None, the entire
	// sheet is used.
	ExtractionRange param.Opt[string] `json:"extraction_range,omitzero"`
	// Optional specialization mode for domain-specific extraction. Supported values:
	// 'financial-standard', 'financial-enhanced', 'financial-precise'. Default None
	// uses the general-purpose pipeline.
	Specialization param.Opt[string] `json:"specialization,omitzero"`
	// Return a flattened dataframe when a detected table is recognized as
	// hierarchical.
	FlattenHierarchicalTables param.Opt[bool] `json:"flatten_hierarchical_tables,omitzero"`
	// Whether to generate additional metadata (title, description) for each extracted
	// region.
	GenerateAdditionalMetadata param.Opt[bool] `json:"generate_additional_metadata,omitzero"`
	// Whether to include hidden cells when extracting regions from the spreadsheet.
	IncludeHiddenCells param.Opt[bool] `json:"include_hidden_cells,omitzero"`
	// Enables experimental processing. Accuracy may be impacted.
	UseExperimentalProcessing param.Opt[bool] `json:"use_experimental_processing,omitzero"`
	// The names of the sheets to extract regions from. If empty, all sheets will be
	// processed.
	SheetNames []string `json:"sheet_names,omitzero"`
	// Influences how likely similar-looking regions are merged into a single table.
	// Useful for spreadsheets that either have sparse tables (strong merging) or many
	// distinct tables close together (weak merging).
	//
	// Any of "strong", "weak".
	TableMergeSensitivity string `json:"table_merge_sensitivity,omitzero"`
	// Product type.
	//
	// This field can be elided, and will marshal its zero value as "spreadsheet_v1".
	ProductType constant.SpreadsheetV1 `json:"product_type" default:"spreadsheet_v1"`
	paramObj
}

func (r ConfigurationUpdateParamsParametersSpreadsheetV1) MarshalJSON() (data []byte, err error) {
	type shadow ConfigurationUpdateParamsParametersSpreadsheetV1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigurationUpdateParamsParametersSpreadsheetV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConfigurationUpdateParamsParametersSpreadsheetV1](
		"table_merge_sensitivity", "strong", "weak",
	)
}

type ConfigurationListParams struct {
	// Filter by configuration name.
	Name           param.Opt[string] `query:"name,omitzero" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	// Number of items per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Pagination token.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Return only the latest version per configuration name.
	LatestOnly param.Opt[bool] `query:"latest_only,omitzero" json:"-"`
	// Filter by one or more product types. Repeat the parameter for multiple values.
	//
	// Any of "classify_v2", "extract_v2", "parse_v2", "split_v1", "spreadsheet_v1",
	// "unknown".
	ProductType []string `query:"product_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConfigurationListParams]'s query parameters as
// `url.Values`.
func (r ConfigurationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigurationDeleteParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [ConfigurationDeleteParams]'s query parameters as
// `url.Values`.
func (r ConfigurationDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
