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

	"github.com/stainless-sdks/llamacloud-prod-go/internal/apijson"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/apiquery"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/requestconfig"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/pagination"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/param"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/respjson"
)

// ParsingService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewParsingService] method instead.
type ParsingService struct {
	options []option.RequestOption
}

// NewParsingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewParsingService(opts ...option.RequestOption) (r ParsingService) {
	r = ParsingService{}
	r.options = opts
	return
}

// Parse a file by file ID or URL.
//
// Provide either `file_id` (a previously uploaded file) or `source_url` (a
// publicly accessible URL). Configure parsing with options like `tier`,
// `target_pages`, and `lang`.
//
// ## Tiers
//
// - `fast` — rule-based, cheapest, no AI
// - `cost_effective` — balanced speed and quality
// - `agentic` — full AI-powered parsing
// - `agentic_plus` — premium AI with specialized features
//
// The job runs asynchronously. Poll `GET /parse/{job_id}` with `expand=text` or
// `expand=markdown` to retrieve results.
func (r *ParsingService) New(ctx context.Context, params ParsingNewParams, opts ...option.RequestOption) (res *ParsingNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v2/parse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// List parse jobs for the current project.
//
// Filter by `status` or creation date range. Results are paginated — use
// `page_token` from the response to fetch subsequent pages.
func (r *ParsingService) List(ctx context.Context, query ParsingListParams, opts ...option.RequestOption) (res *pagination.PaginatedCursor[ParsingListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/v2/parse"
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

// List parse jobs for the current project.
//
// Filter by `status` or creation date range. Results are paginated — use
// `page_token` from the response to fetch subsequent pages.
func (r *ParsingService) ListAutoPaging(ctx context.Context, query ParsingListParams, opts ...option.RequestOption) *pagination.PaginatedCursorAutoPager[ParsingListResponse] {
	return pagination.NewPaginatedCursorAutoPager(r.List(ctx, query, opts...))
}

// Retrieve a parse job with optional expanded content.
//
// By default returns job metadata only. Use `expand` to include parsed content:
//
// - `text` — plain text output
// - `markdown` — markdown output
// - `items` — structured page-by-page output
// - `job_metadata` — usage and processing details
//
// Content metadata fields (e.g. `text_content_metadata`) return presigned URLs for
// downloading large results.
func (r *ParsingService) Get(ctx context.Context, jobID string, query ParsingGetParams, opts ...option.RequestOption) (res *ParsingGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/parse/%s", url.PathEscape(jobID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Bounding box with coordinates and optional metadata.
type BBox struct {
	// Height of the bounding box
	H float64 `json:"h" api:"required"`
	// Width of the bounding box
	W float64 `json:"w" api:"required"`
	// X coordinate of the bounding box
	X float64 `json:"x" api:"required"`
	// Y coordinate of the bounding box
	Y float64 `json:"y" api:"required"`
	// Confidence score
	Confidence float64 `json:"confidence" api:"nullable"`
	// End index in the text
	EndIndex int64 `json:"end_index" api:"nullable"`
	// Label for the bounding box
	Label string `json:"label" api:"nullable"`
	// Optional visual text rotation angle in degrees. Omitted when unrotated.
	R float64 `json:"r" api:"nullable"`
	// Start index in the text
	StartIndex int64 `json:"start_index" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H           respjson.Field
		W           respjson.Field
		X           respjson.Field
		Y           respjson.Field
		Confidence  respjson.Field
		EndIndex    respjson.Field
		Label       respjson.Field
		R           respjson.Field
		StartIndex  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BBox) RawJSON() string { return r.JSON.raw }
func (r *BBox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CodeItem struct {
	// Markdown representation preserving formatting
	Md string `json:"md" api:"required"`
	// Code content
	Value string `json:"value" api:"required"`
	// List of bounding boxes
	Bbox []BBox `json:"bbox" api:"nullable"`
	// Programming language identifier
	Language string `json:"language" api:"nullable"`
	// Code block item type
	//
	// Any of "code".
	Type CodeItemType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Md          respjson.Field
		Value       respjson.Field
		Bbox        respjson.Field
		Language    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CodeItem) RawJSON() string { return r.JSON.raw }
func (r *CodeItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Code block item type
type CodeItemType string

const (
	CodeItemTypeCode CodeItemType = "code"
)

// Enum for representing the different available page error handling modes.
type FailPageMode string

const (
	FailPageModeRawText      FailPageMode = "raw_text"
	FailPageModeBlankPage    FailPageMode = "blank_page"
	FailPageModeErrorMessage FailPageMode = "error_message"
)

type FooterItem struct {
	// List of items within the footer
	Items []FooterItemItemUnion `json:"items" api:"required"`
	// Markdown representation preserving formatting
	Md string `json:"md" api:"required"`
	// List of bounding boxes
	Bbox []BBox `json:"bbox" api:"nullable"`
	// Page footer container
	//
	// Any of "footer".
	Type FooterItemType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Md          respjson.Field
		Bbox        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FooterItem) RawJSON() string { return r.JSON.raw }
func (r *FooterItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FooterItemItemUnion contains all possible properties and values from [TextItem],
// [HeadingItem], [ListItem], [CodeItem], [TableItem], [ImageItem], [LinkItem].
//
// Use the [FooterItemItemUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FooterItemItemUnion struct {
	Md    string `json:"md"`
	Value string `json:"value"`
	Bbox  []BBox `json:"bbox"`
	// Any of "text", "heading", "list", "code", "table", "image", "link".
	Type string `json:"type"`
	// This field is from variant [HeadingItem].
	Level int64 `json:"level"`
	// This field is from variant [ListItem].
	Items []ListItemItemUnion `json:"items"`
	// This field is from variant [ListItem].
	Ordered bool `json:"ordered"`
	// This field is from variant [CodeItem].
	Language string `json:"language"`
	// This field is from variant [TableItem].
	Csv string `json:"csv"`
	// This field is from variant [TableItem].
	HTML string `json:"html"`
	// This field is from variant [TableItem].
	Rows [][]*TableItemRowUnion `json:"rows"`
	// This field is from variant [TableItem].
	MergedFromPages []int64 `json:"merged_from_pages"`
	// This field is from variant [TableItem].
	MergedIntoPage int64 `json:"merged_into_page"`
	// This field is from variant [TableItem].
	ParseConcerns []TableItemParseConcern `json:"parse_concerns"`
	// This field is from variant [ImageItem].
	Caption string `json:"caption"`
	URL     string `json:"url"`
	// This field is from variant [LinkItem].
	Text string `json:"text"`
	JSON struct {
		Md              respjson.Field
		Value           respjson.Field
		Bbox            respjson.Field
		Type            respjson.Field
		Level           respjson.Field
		Items           respjson.Field
		Ordered         respjson.Field
		Language        respjson.Field
		Csv             respjson.Field
		HTML            respjson.Field
		Rows            respjson.Field
		MergedFromPages respjson.Field
		MergedIntoPage  respjson.Field
		ParseConcerns   respjson.Field
		Caption         respjson.Field
		URL             respjson.Field
		Text            respjson.Field
		raw             string
	} `json:"-"`
}

// anyFooterItemItem is implemented by each variant of [FooterItemItemUnion] to add
// type safety for the return type of [FooterItemItemUnion.AsAny]
type anyFooterItemItem interface {
	implFooterItemItemUnion()
}

func (TextItem) implFooterItemItemUnion()    {}
func (HeadingItem) implFooterItemItemUnion() {}
func (ListItem) implFooterItemItemUnion()    {}
func (CodeItem) implFooterItemItemUnion()    {}
func (TableItem) implFooterItemItemUnion()   {}
func (ImageItem) implFooterItemItemUnion()   {}
func (LinkItem) implFooterItemItemUnion()    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := FooterItemItemUnion.AsAny().(type) {
//	case llamacloudprod.TextItem:
//	case llamacloudprod.HeadingItem:
//	case llamacloudprod.ListItem:
//	case llamacloudprod.CodeItem:
//	case llamacloudprod.TableItem:
//	case llamacloudprod.ImageItem:
//	case llamacloudprod.LinkItem:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FooterItemItemUnion) AsAny() anyFooterItemItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "heading":
		return u.AsHeading()
	case "list":
		return u.AsList()
	case "code":
		return u.AsCode()
	case "table":
		return u.AsTable()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	}
	return nil
}

func (u FooterItemItemUnion) AsText() (v TextItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FooterItemItemUnion) AsHeading() (v HeadingItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FooterItemItemUnion) AsList() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FooterItemItemUnion) AsCode() (v CodeItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FooterItemItemUnion) AsTable() (v TableItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FooterItemItemUnion) AsImage() (v ImageItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FooterItemItemUnion) AsLink() (v LinkItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FooterItemItemUnion) RawJSON() string { return u.JSON.raw }

func (r *FooterItemItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page footer container
type FooterItemType string

const (
	FooterItemTypeFooter FooterItemType = "footer"
)

type HeaderItem struct {
	// List of items within the header
	Items []HeaderItemItemUnion `json:"items" api:"required"`
	// Markdown representation preserving formatting
	Md string `json:"md" api:"required"`
	// List of bounding boxes
	Bbox []BBox `json:"bbox" api:"nullable"`
	// Page header container
	//
	// Any of "header".
	Type HeaderItemType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Md          respjson.Field
		Bbox        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeaderItem) RawJSON() string { return r.JSON.raw }
func (r *HeaderItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HeaderItemItemUnion contains all possible properties and values from [TextItem],
// [HeadingItem], [ListItem], [CodeItem], [TableItem], [ImageItem], [LinkItem].
//
// Use the [HeaderItemItemUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type HeaderItemItemUnion struct {
	Md    string `json:"md"`
	Value string `json:"value"`
	Bbox  []BBox `json:"bbox"`
	// Any of "text", "heading", "list", "code", "table", "image", "link".
	Type string `json:"type"`
	// This field is from variant [HeadingItem].
	Level int64 `json:"level"`
	// This field is from variant [ListItem].
	Items []ListItemItemUnion `json:"items"`
	// This field is from variant [ListItem].
	Ordered bool `json:"ordered"`
	// This field is from variant [CodeItem].
	Language string `json:"language"`
	// This field is from variant [TableItem].
	Csv string `json:"csv"`
	// This field is from variant [TableItem].
	HTML string `json:"html"`
	// This field is from variant [TableItem].
	Rows [][]*TableItemRowUnion `json:"rows"`
	// This field is from variant [TableItem].
	MergedFromPages []int64 `json:"merged_from_pages"`
	// This field is from variant [TableItem].
	MergedIntoPage int64 `json:"merged_into_page"`
	// This field is from variant [TableItem].
	ParseConcerns []TableItemParseConcern `json:"parse_concerns"`
	// This field is from variant [ImageItem].
	Caption string `json:"caption"`
	URL     string `json:"url"`
	// This field is from variant [LinkItem].
	Text string `json:"text"`
	JSON struct {
		Md              respjson.Field
		Value           respjson.Field
		Bbox            respjson.Field
		Type            respjson.Field
		Level           respjson.Field
		Items           respjson.Field
		Ordered         respjson.Field
		Language        respjson.Field
		Csv             respjson.Field
		HTML            respjson.Field
		Rows            respjson.Field
		MergedFromPages respjson.Field
		MergedIntoPage  respjson.Field
		ParseConcerns   respjson.Field
		Caption         respjson.Field
		URL             respjson.Field
		Text            respjson.Field
		raw             string
	} `json:"-"`
}

// anyHeaderItemItem is implemented by each variant of [HeaderItemItemUnion] to add
// type safety for the return type of [HeaderItemItemUnion.AsAny]
type anyHeaderItemItem interface {
	implHeaderItemItemUnion()
}

func (TextItem) implHeaderItemItemUnion()    {}
func (HeadingItem) implHeaderItemItemUnion() {}
func (ListItem) implHeaderItemItemUnion()    {}
func (CodeItem) implHeaderItemItemUnion()    {}
func (TableItem) implHeaderItemItemUnion()   {}
func (ImageItem) implHeaderItemItemUnion()   {}
func (LinkItem) implHeaderItemItemUnion()    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := HeaderItemItemUnion.AsAny().(type) {
//	case llamacloudprod.TextItem:
//	case llamacloudprod.HeadingItem:
//	case llamacloudprod.ListItem:
//	case llamacloudprod.CodeItem:
//	case llamacloudprod.TableItem:
//	case llamacloudprod.ImageItem:
//	case llamacloudprod.LinkItem:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u HeaderItemItemUnion) AsAny() anyHeaderItemItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "heading":
		return u.AsHeading()
	case "list":
		return u.AsList()
	case "code":
		return u.AsCode()
	case "table":
		return u.AsTable()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	}
	return nil
}

func (u HeaderItemItemUnion) AsText() (v TextItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeaderItemItemUnion) AsHeading() (v HeadingItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeaderItemItemUnion) AsList() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeaderItemItemUnion) AsCode() (v CodeItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeaderItemItemUnion) AsTable() (v TableItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeaderItemItemUnion) AsImage() (v ImageItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeaderItemItemUnion) AsLink() (v LinkItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u HeaderItemItemUnion) RawJSON() string { return u.JSON.raw }

func (r *HeaderItemItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page header container
type HeaderItemType string

const (
	HeaderItemTypeHeader HeaderItemType = "header"
)

type HeadingItem struct {
	// Heading level (1-6)
	Level int64 `json:"level" api:"required"`
	// Markdown representation preserving formatting
	Md string `json:"md" api:"required"`
	// Heading text content
	Value string `json:"value" api:"required"`
	// List of bounding boxes
	Bbox []BBox `json:"bbox" api:"nullable"`
	// Heading item type
	//
	// Any of "heading".
	Type HeadingItemType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		Md          respjson.Field
		Value       respjson.Field
		Bbox        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HeadingItem) RawJSON() string { return r.JSON.raw }
func (r *HeadingItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Heading item type
type HeadingItemType string

const (
	HeadingItemTypeHeading HeadingItemType = "heading"
)

type ImageItem struct {
	// Image caption
	Caption string `json:"caption" api:"required"`
	// Markdown representation preserving formatting
	Md string `json:"md" api:"required"`
	// URL to the image
	URL string `json:"url" api:"required"`
	// List of bounding boxes
	Bbox []BBox `json:"bbox" api:"nullable"`
	// Image item type
	//
	// Any of "image".
	Type ImageItemType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Md          respjson.Field
		URL         respjson.Field
		Bbox        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageItem) RawJSON() string { return r.JSON.raw }
func (r *ImageItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image item type
type ImageItemType string

const (
	ImageItemTypeImage ImageItemType = "image"
)

type LinkItem struct {
	// Markdown representation preserving formatting
	Md string `json:"md" api:"required"`
	// Display text of the link
	Text string `json:"text" api:"required"`
	// URL of the link
	URL string `json:"url" api:"required"`
	// List of bounding boxes
	Bbox []BBox `json:"bbox" api:"nullable"`
	// Link item type
	//
	// Any of "link".
	Type LinkItemType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Md          respjson.Field
		Text        respjson.Field
		URL         respjson.Field
		Bbox        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkItem) RawJSON() string { return r.JSON.raw }
func (r *LinkItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Link item type
type LinkItemType string

const (
	LinkItemTypeLink LinkItemType = "link"
)

type ListItem struct {
	// List of nested text or list items
	Items []ListItemItemUnion `json:"items" api:"required"`
	// Markdown representation preserving formatting
	Md string `json:"md" api:"required"`
	// Whether the list is ordered or unordered
	Ordered bool `json:"ordered" api:"required"`
	// List of bounding boxes
	Bbox []BBox `json:"bbox" api:"nullable"`
	// List item type
	//
	// Any of "list".
	Type ListItemType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Md          respjson.Field
		Ordered     respjson.Field
		Bbox        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListItem) RawJSON() string { return r.JSON.raw }
func (r *ListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListItemItemUnion contains all possible properties and values from [TextItem],
// [ListItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ListItemItemUnion struct {
	Md string `json:"md"`
	// This field is from variant [TextItem].
	Value string `json:"value"`
	Bbox  []BBox `json:"bbox"`
	Type  string `json:"type"`
	// This field is from variant [ListItem].
	Items []ListItemItemUnion `json:"items"`
	// This field is from variant [ListItem].
	Ordered bool `json:"ordered"`
	JSON    struct {
		Md      respjson.Field
		Value   respjson.Field
		Bbox    respjson.Field
		Type    respjson.Field
		Items   respjson.Field
		Ordered respjson.Field
		raw     string
	} `json:"-"`
}

func (u ListItemItemUnion) AsTextItem() (v TextItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemItemUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListItemItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ListItemItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// List item type
type ListItemType string

const (
	ListItemTypeList ListItemType = "list"
)

// Enum for representing the languages supported by the parser.
type ParsingLanguages string

const (
	ParsingLanguagesAf         ParsingLanguages = "af"
	ParsingLanguagesAz         ParsingLanguages = "az"
	ParsingLanguagesBs         ParsingLanguages = "bs"
	ParsingLanguagesCs         ParsingLanguages = "cs"
	ParsingLanguagesCy         ParsingLanguages = "cy"
	ParsingLanguagesDa         ParsingLanguages = "da"
	ParsingLanguagesDe         ParsingLanguages = "de"
	ParsingLanguagesEn         ParsingLanguages = "en"
	ParsingLanguagesEs         ParsingLanguages = "es"
	ParsingLanguagesEt         ParsingLanguages = "et"
	ParsingLanguagesFr         ParsingLanguages = "fr"
	ParsingLanguagesGa         ParsingLanguages = "ga"
	ParsingLanguagesHr         ParsingLanguages = "hr"
	ParsingLanguagesHu         ParsingLanguages = "hu"
	ParsingLanguagesID         ParsingLanguages = "id"
	ParsingLanguagesIs         ParsingLanguages = "is"
	ParsingLanguagesIt         ParsingLanguages = "it"
	ParsingLanguagesKu         ParsingLanguages = "ku"
	ParsingLanguagesLa         ParsingLanguages = "la"
	ParsingLanguagesLt         ParsingLanguages = "lt"
	ParsingLanguagesLv         ParsingLanguages = "lv"
	ParsingLanguagesMi         ParsingLanguages = "mi"
	ParsingLanguagesMs         ParsingLanguages = "ms"
	ParsingLanguagesMt         ParsingLanguages = "mt"
	ParsingLanguagesNl         ParsingLanguages = "nl"
	ParsingLanguagesNo         ParsingLanguages = "no"
	ParsingLanguagesOc         ParsingLanguages = "oc"
	ParsingLanguagesPi         ParsingLanguages = "pi"
	ParsingLanguagesPl         ParsingLanguages = "pl"
	ParsingLanguagesPt         ParsingLanguages = "pt"
	ParsingLanguagesRo         ParsingLanguages = "ro"
	ParsingLanguagesRsLatin    ParsingLanguages = "rs_latin"
	ParsingLanguagesSk         ParsingLanguages = "sk"
	ParsingLanguagesSl         ParsingLanguages = "sl"
	ParsingLanguagesSq         ParsingLanguages = "sq"
	ParsingLanguagesSv         ParsingLanguages = "sv"
	ParsingLanguagesSw         ParsingLanguages = "sw"
	ParsingLanguagesTl         ParsingLanguages = "tl"
	ParsingLanguagesTr         ParsingLanguages = "tr"
	ParsingLanguagesUz         ParsingLanguages = "uz"
	ParsingLanguagesVi         ParsingLanguages = "vi"
	ParsingLanguagesAr         ParsingLanguages = "ar"
	ParsingLanguagesFa         ParsingLanguages = "fa"
	ParsingLanguagesUg         ParsingLanguages = "ug"
	ParsingLanguagesUr         ParsingLanguages = "ur"
	ParsingLanguagesBn         ParsingLanguages = "bn"
	ParsingLanguagesAs         ParsingLanguages = "as"
	ParsingLanguagesMni        ParsingLanguages = "mni"
	ParsingLanguagesRu         ParsingLanguages = "ru"
	ParsingLanguagesRsCyrillic ParsingLanguages = "rs_cyrillic"
	ParsingLanguagesBe         ParsingLanguages = "be"
	ParsingLanguagesBg         ParsingLanguages = "bg"
	ParsingLanguagesUk         ParsingLanguages = "uk"
	ParsingLanguagesMn         ParsingLanguages = "mn"
	ParsingLanguagesAbq        ParsingLanguages = "abq"
	ParsingLanguagesAdy        ParsingLanguages = "ady"
	ParsingLanguagesKbd        ParsingLanguages = "kbd"
	ParsingLanguagesAva        ParsingLanguages = "ava"
	ParsingLanguagesDar        ParsingLanguages = "dar"
	ParsingLanguagesInh        ParsingLanguages = "inh"
	ParsingLanguagesChe        ParsingLanguages = "che"
	ParsingLanguagesLbe        ParsingLanguages = "lbe"
	ParsingLanguagesLez        ParsingLanguages = "lez"
	ParsingLanguagesTab        ParsingLanguages = "tab"
	ParsingLanguagesTjk        ParsingLanguages = "tjk"
	ParsingLanguagesHi         ParsingLanguages = "hi"
	ParsingLanguagesMr         ParsingLanguages = "mr"
	ParsingLanguagesNe         ParsingLanguages = "ne"
	ParsingLanguagesBh         ParsingLanguages = "bh"
	ParsingLanguagesMai        ParsingLanguages = "mai"
	ParsingLanguagesAng        ParsingLanguages = "ang"
	ParsingLanguagesBho        ParsingLanguages = "bho"
	ParsingLanguagesMah        ParsingLanguages = "mah"
	ParsingLanguagesSck        ParsingLanguages = "sck"
	ParsingLanguagesNew        ParsingLanguages = "new"
	ParsingLanguagesGom        ParsingLanguages = "gom"
	ParsingLanguagesSa         ParsingLanguages = "sa"
	ParsingLanguagesBgc        ParsingLanguages = "bgc"
	ParsingLanguagesTh         ParsingLanguages = "th"
	ParsingLanguagesChSim      ParsingLanguages = "ch_sim"
	ParsingLanguagesChTra      ParsingLanguages = "ch_tra"
	ParsingLanguagesJa         ParsingLanguages = "ja"
	ParsingLanguagesKo         ParsingLanguages = "ko"
	ParsingLanguagesTa         ParsingLanguages = "ta"
	ParsingLanguagesTe         ParsingLanguages = "te"
	ParsingLanguagesKn         ParsingLanguages = "kn"
)

// Enum for representing the mode of parsing to be used.
type ParsingMode string

const (
	ParsingModeParsePageWithoutLlm      ParsingMode = "parse_page_without_llm"
	ParsingModeParsePageWithLlm         ParsingMode = "parse_page_with_llm"
	ParsingModeParsePageWithLvm         ParsingMode = "parse_page_with_lvm"
	ParsingModeParsePageWithAgent       ParsingMode = "parse_page_with_agent"
	ParsingModeParsePageWithLayoutAgent ParsingMode = "parse_page_with_layout_agent"
	ParsingModeParseDocumentWithLlm     ParsingMode = "parse_document_with_llm"
	ParsingModeParseDocumentWithLvm     ParsingMode = "parse_document_with_lvm"
	ParsingModeParseDocumentWithAgent   ParsingMode = "parse_document_with_agent"
)

// Enum for representing the status of a job
type StatusEnum string

const (
	StatusEnumPending        StatusEnum = "PENDING"
	StatusEnumSuccess        StatusEnum = "SUCCESS"
	StatusEnumError          StatusEnum = "ERROR"
	StatusEnumPartialSuccess StatusEnum = "PARTIAL_SUCCESS"
	StatusEnumCancelled      StatusEnum = "CANCELLED"
)

type TableItem struct {
	// CSV representation of the table
	Csv string `json:"csv" api:"required"`
	// HTML representation of the table
	HTML string `json:"html" api:"required"`
	// Markdown representation preserving formatting
	Md string `json:"md" api:"required"`
	// Table data as array of arrays (string, number, or null)
	Rows [][]*TableItemRowUnion `json:"rows" api:"required"`
	// List of bounding boxes
	Bbox []BBox `json:"bbox" api:"nullable"`
	// List of page numbers with tables that were merged into this table (e.g., [1, 2,
	// 3, 4])
	MergedFromPages []int64 `json:"merged_from_pages" api:"nullable"`
	// Populated when merged into another table. Page number where the full merged
	// table begins (used on empty tables).
	MergedIntoPage int64 `json:"merged_into_page" api:"nullable"`
	// Quality concerns detected during table extraction, indicating the table may have
	// issues
	ParseConcerns []TableItemParseConcern `json:"parse_concerns" api:"nullable"`
	// Table item type
	//
	// Any of "table".
	Type TableItemType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Csv             respjson.Field
		HTML            respjson.Field
		Md              respjson.Field
		Rows            respjson.Field
		Bbox            respjson.Field
		MergedFromPages respjson.Field
		MergedIntoPage  respjson.Field
		ParseConcerns   respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TableItem) RawJSON() string { return r.JSON.raw }
func (r *TableItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TableItemRowUnion contains all possible properties and values from [string],
// [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type TableItemRowUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u TableItemRowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableItemRowUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TableItemRowUnion) RawJSON() string { return u.JSON.raw }

func (r *TableItemRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TableItemParseConcern struct {
	// Human-readable details about the concern
	Details string `json:"details" api:"required"`
	// Type of parse concern (e.g. header_value_type_mismatch,
	// inconsistent_row_cell_count)
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Details     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TableItemParseConcern) RawJSON() string { return r.JSON.raw }
func (r *TableItemParseConcern) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Table item type
type TableItemType string

const (
	TableItemTypeTable TableItemType = "table"
)

type TextItem struct {
	// Markdown representation preserving formatting
	Md string `json:"md" api:"required"`
	// Text content
	Value string `json:"value" api:"required"`
	// List of bounding boxes
	Bbox []BBox `json:"bbox" api:"nullable"`
	// Text item type
	//
	// Any of "text".
	Type TextItemType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Md          respjson.Field
		Value       respjson.Field
		Bbox        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextItem) RawJSON() string { return r.JSON.raw }
func (r *TextItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text item type
type TextItemType string

const (
	TextItemTypeText TextItemType = "text"
)

// A parse job.
type ParsingNewResponse struct {
	// Unique parse job identifier
	ID string `json:"id" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
	//
	// Any of "PENDING", "RUNNING", "COMPLETED", "FAILED", "CANCELLED".
	Status ParsingNewResponseStatus `json:"status" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error details when status is FAILED
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Optional display name for this parse job
	Name string `json:"name" api:"nullable"`
	// Parsing tier used for this job
	Tier string `json:"tier" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ProjectID    respjson.Field
		Status       respjson.Field
		CreatedAt    respjson.Field
		ErrorMessage respjson.Field
		Name         respjson.Field
		Tier         respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ParsingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
type ParsingNewResponseStatus string

const (
	ParsingNewResponseStatusPending   ParsingNewResponseStatus = "PENDING"
	ParsingNewResponseStatusRunning   ParsingNewResponseStatus = "RUNNING"
	ParsingNewResponseStatusCompleted ParsingNewResponseStatus = "COMPLETED"
	ParsingNewResponseStatusFailed    ParsingNewResponseStatus = "FAILED"
	ParsingNewResponseStatusCancelled ParsingNewResponseStatus = "CANCELLED"
)

// A parse job.
type ParsingListResponse struct {
	// Unique parse job identifier
	ID string `json:"id" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
	//
	// Any of "PENDING", "RUNNING", "COMPLETED", "FAILED", "CANCELLED".
	Status ParsingListResponseStatus `json:"status" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error details when status is FAILED
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Optional display name for this parse job
	Name string `json:"name" api:"nullable"`
	// Parsing tier used for this job
	Tier string `json:"tier" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ProjectID    respjson.Field
		Status       respjson.Field
		CreatedAt    respjson.Field
		ErrorMessage respjson.Field
		Name         respjson.Field
		Tier         respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingListResponse) RawJSON() string { return r.JSON.raw }
func (r *ParsingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
type ParsingListResponseStatus string

const (
	ParsingListResponseStatusPending   ParsingListResponseStatus = "PENDING"
	ParsingListResponseStatusRunning   ParsingListResponseStatus = "RUNNING"
	ParsingListResponseStatusCompleted ParsingListResponseStatus = "COMPLETED"
	ParsingListResponseStatusFailed    ParsingListResponseStatus = "FAILED"
	ParsingListResponseStatusCancelled ParsingListResponseStatus = "CANCELLED"
)

// Parse result response with job status and optional content or metadata.
//
// The job field is always included. Other fields are included based on expand
// parameters.
type ParsingGetResponse struct {
	// Parse job status and metadata
	Job ParsingGetResponseJob `json:"job" api:"required"`
	// Metadata for all extracted images.
	ImagesContentMetadata ParsingGetResponseImagesContentMetadata `json:"images_content_metadata" api:"nullable"`
	// Structured JSON result (if requested)
	Items ParsingGetResponseItems `json:"items" api:"nullable"`
	// Job execution metadata (if requested)
	JobMetadata map[string]any `json:"job_metadata" api:"nullable"`
	// Markdown result (if requested)
	Markdown ParsingGetResponseMarkdown `json:"markdown" api:"nullable"`
	// Full raw markdown content (if requested)
	MarkdownFull string `json:"markdown_full" api:"nullable"`
	// Result containing metadata (page level and general) for the parsed document.
	Metadata      ParsingGetResponseMetadata `json:"metadata" api:"nullable"`
	RawParameters map[string]any             `json:"raw_parameters" api:"nullable"`
	// Metadata including size, existence, and presigned URLs for result files
	ResultContentMetadata map[string]ParsingGetResponseResultContentMetadata `json:"result_content_metadata" api:"nullable"`
	// Plain text result (if requested)
	Text ParsingGetResponseText `json:"text" api:"nullable"`
	// Full raw text content (if requested)
	TextFull string `json:"text_full" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Job                   respjson.Field
		ImagesContentMetadata respjson.Field
		Items                 respjson.Field
		JobMetadata           respjson.Field
		Markdown              respjson.Field
		MarkdownFull          respjson.Field
		Metadata              respjson.Field
		RawParameters         respjson.Field
		ResultContentMetadata respjson.Field
		Text                  respjson.Field
		TextFull              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parse job status and metadata
type ParsingGetResponseJob struct {
	// Unique parse job identifier
	ID string `json:"id" api:"required"`
	// Project this job belongs to
	ProjectID string `json:"project_id" api:"required"`
	// Current job status: PENDING, RUNNING, COMPLETED, FAILED, or CANCELLED
	//
	// Any of "PENDING", "RUNNING", "COMPLETED", "FAILED", "CANCELLED".
	Status string `json:"status" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error details when status is FAILED
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Optional display name for this parse job
	Name string `json:"name" api:"nullable"`
	// Parsing tier used for this job
	Tier string `json:"tier" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ProjectID    respjson.Field
		Status       respjson.Field
		CreatedAt    respjson.Field
		ErrorMessage respjson.Field
		Name         respjson.Field
		Tier         respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseJob) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for all extracted images.
type ParsingGetResponseImagesContentMetadata struct {
	// List of image metadata with presigned URLs
	Images []ParsingGetResponseImagesContentMetadataImage `json:"images" api:"required"`
	// Total number of extracted images
	TotalCount int64 `json:"total_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Images      respjson.Field
		TotalCount  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseImagesContentMetadata) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseImagesContentMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for a single extracted image.
type ParsingGetResponseImagesContentMetadataImage struct {
	// Image filename (e.g., 'image_0.png')
	Filename string `json:"filename" api:"required"`
	// Index of the image in the extraction order
	Index int64 `json:"index" api:"required"`
	// Bounding box for an image on its page.
	Bbox ParsingGetResponseImagesContentMetadataImageBbox `json:"bbox" api:"nullable"`
	// Image category: 'screenshot' (full page), 'embedded' (images in document), or
	// 'layout' (cropped from layout detection)
	//
	// Any of "screenshot", "embedded", "layout".
	Category string `json:"category" api:"nullable"`
	// MIME type of the image
	ContentType string `json:"content_type" api:"nullable"`
	// Presigned URL to download the image
	PresignedURL string `json:"presigned_url" api:"nullable"`
	// Deprecated: always returns None. Will be removed in a future release.
	//
	// Deprecated: deprecated
	SizeBytes int64 `json:"size_bytes" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filename     respjson.Field
		Index        respjson.Field
		Bbox         respjson.Field
		Category     respjson.Field
		ContentType  respjson.Field
		PresignedURL respjson.Field
		SizeBytes    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseImagesContentMetadataImage) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseImagesContentMetadataImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bounding box for an image on its page.
type ParsingGetResponseImagesContentMetadataImageBbox struct {
	// Height of the bounding box
	H int64 `json:"h" api:"required"`
	// Width of the bounding box
	W int64 `json:"w" api:"required"`
	// X coordinate of the bounding box
	X int64 `json:"x" api:"required"`
	// Y coordinate of the bounding box
	Y int64 `json:"y" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		H           respjson.Field
		W           respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseImagesContentMetadataImageBbox) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseImagesContentMetadataImageBbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured JSON result (if requested)
type ParsingGetResponseItems struct {
	// List of structured pages or failed page entries
	Pages []ParsingGetResponseItemsPageUnion `json:"pages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pages       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseItems) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseItems) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParsingGetResponseItemsPageUnion contains all possible properties and values
// from [ParsingGetResponseItemsPageStructuredResultPage],
// [ParsingGetResponseItemsPageFailedStructuredPage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ParsingGetResponseItemsPageUnion struct {
	// This field is from variant [ParsingGetResponseItemsPageStructuredResultPage].
	Items []ParsingGetResponseItemsPageStructuredResultPageItemUnion `json:"items"`
	// This field is from variant [ParsingGetResponseItemsPageStructuredResultPage].
	PageHeight float64 `json:"page_height"`
	PageNumber int64   `json:"page_number"`
	// This field is from variant [ParsingGetResponseItemsPageStructuredResultPage].
	PageWidth float64 `json:"page_width"`
	Success   bool    `json:"success"`
	// This field is from variant [ParsingGetResponseItemsPageFailedStructuredPage].
	Error string `json:"error"`
	JSON  struct {
		Items      respjson.Field
		PageHeight respjson.Field
		PageNumber respjson.Field
		PageWidth  respjson.Field
		Success    respjson.Field
		Error      respjson.Field
		raw        string
	} `json:"-"`
}

func (u ParsingGetResponseItemsPageUnion) AsStructuredResultPage() (v ParsingGetResponseItemsPageStructuredResultPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsingGetResponseItemsPageUnion) AsFailedStructuredPage() (v ParsingGetResponseItemsPageFailedStructuredPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParsingGetResponseItemsPageUnion) RawJSON() string { return u.JSON.raw }

func (r *ParsingGetResponseItemsPageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParsingGetResponseItemsPageStructuredResultPage struct {
	// List of structured items on the page
	Items []ParsingGetResponseItemsPageStructuredResultPageItemUnion `json:"items" api:"required"`
	// Height of the page in points
	PageHeight float64 `json:"page_height" api:"required"`
	// Page number of the document
	PageNumber int64 `json:"page_number" api:"required"`
	// Width of the page in points
	PageWidth float64 `json:"page_width" api:"required"`
	// Success indicator
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		PageHeight  respjson.Field
		PageNumber  respjson.Field
		PageWidth   respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseItemsPageStructuredResultPage) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseItemsPageStructuredResultPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParsingGetResponseItemsPageStructuredResultPageItemUnion contains all possible
// properties and values from [TextItem], [HeadingItem], [ListItem], [CodeItem],
// [TableItem], [ImageItem], [LinkItem], [HeaderItem], [FooterItem].
//
// Use the [ParsingGetResponseItemsPageStructuredResultPageItemUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ParsingGetResponseItemsPageStructuredResultPageItemUnion struct {
	Md    string `json:"md"`
	Value string `json:"value"`
	Bbox  []BBox `json:"bbox"`
	// Any of "text", "heading", "list", "code", "table", "image", "link", "header",
	// "footer".
	Type string `json:"type"`
	// This field is from variant [HeadingItem].
	Level int64 `json:"level"`
	// This field is a union of [[]ListItemItemUnion], [[]HeaderItemItemUnion],
	// [[]FooterItemItemUnion]
	Items ParsingGetResponseItemsPageStructuredResultPageItemUnionItems `json:"items"`
	// This field is from variant [ListItem].
	Ordered bool `json:"ordered"`
	// This field is from variant [CodeItem].
	Language string `json:"language"`
	// This field is from variant [TableItem].
	Csv string `json:"csv"`
	// This field is from variant [TableItem].
	HTML string `json:"html"`
	// This field is from variant [TableItem].
	Rows [][]*TableItemRowUnion `json:"rows"`
	// This field is from variant [TableItem].
	MergedFromPages []int64 `json:"merged_from_pages"`
	// This field is from variant [TableItem].
	MergedIntoPage int64 `json:"merged_into_page"`
	// This field is from variant [TableItem].
	ParseConcerns []TableItemParseConcern `json:"parse_concerns"`
	// This field is from variant [ImageItem].
	Caption string `json:"caption"`
	URL     string `json:"url"`
	// This field is from variant [LinkItem].
	Text string `json:"text"`
	JSON struct {
		Md              respjson.Field
		Value           respjson.Field
		Bbox            respjson.Field
		Type            respjson.Field
		Level           respjson.Field
		Items           respjson.Field
		Ordered         respjson.Field
		Language        respjson.Field
		Csv             respjson.Field
		HTML            respjson.Field
		Rows            respjson.Field
		MergedFromPages respjson.Field
		MergedIntoPage  respjson.Field
		ParseConcerns   respjson.Field
		Caption         respjson.Field
		URL             respjson.Field
		Text            respjson.Field
		raw             string
	} `json:"-"`
}

// anyParsingGetResponseItemsPageStructuredResultPageItem is implemented by each
// variant of [ParsingGetResponseItemsPageStructuredResultPageItemUnion] to add
// type safety for the return type of
// [ParsingGetResponseItemsPageStructuredResultPageItemUnion.AsAny]
type anyParsingGetResponseItemsPageStructuredResultPageItem interface {
	implParsingGetResponseItemsPageStructuredResultPageItemUnion()
}

func (TextItem) implParsingGetResponseItemsPageStructuredResultPageItemUnion()    {}
func (HeadingItem) implParsingGetResponseItemsPageStructuredResultPageItemUnion() {}
func (ListItem) implParsingGetResponseItemsPageStructuredResultPageItemUnion()    {}
func (CodeItem) implParsingGetResponseItemsPageStructuredResultPageItemUnion()    {}
func (TableItem) implParsingGetResponseItemsPageStructuredResultPageItemUnion()   {}
func (ImageItem) implParsingGetResponseItemsPageStructuredResultPageItemUnion()   {}
func (LinkItem) implParsingGetResponseItemsPageStructuredResultPageItemUnion()    {}
func (HeaderItem) implParsingGetResponseItemsPageStructuredResultPageItemUnion()  {}
func (FooterItem) implParsingGetResponseItemsPageStructuredResultPageItemUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ParsingGetResponseItemsPageStructuredResultPageItemUnion.AsAny().(type) {
//	case llamacloudprod.TextItem:
//	case llamacloudprod.HeadingItem:
//	case llamacloudprod.ListItem:
//	case llamacloudprod.CodeItem:
//	case llamacloudprod.TableItem:
//	case llamacloudprod.ImageItem:
//	case llamacloudprod.LinkItem:
//	case llamacloudprod.HeaderItem:
//	case llamacloudprod.FooterItem:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ParsingGetResponseItemsPageStructuredResultPageItemUnion) AsAny() anyParsingGetResponseItemsPageStructuredResultPageItem {
	switch u.Type {
	case "text":
		return u.AsText()
	case "heading":
		return u.AsHeading()
	case "list":
		return u.AsList()
	case "code":
		return u.AsCode()
	case "table":
		return u.AsTable()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "header":
		return u.AsHeader()
	case "footer":
		return u.AsFooter()
	}
	return nil
}

func (u ParsingGetResponseItemsPageStructuredResultPageItemUnion) AsText() (v TextItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsingGetResponseItemsPageStructuredResultPageItemUnion) AsHeading() (v HeadingItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsingGetResponseItemsPageStructuredResultPageItemUnion) AsList() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsingGetResponseItemsPageStructuredResultPageItemUnion) AsCode() (v CodeItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsingGetResponseItemsPageStructuredResultPageItemUnion) AsTable() (v TableItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsingGetResponseItemsPageStructuredResultPageItemUnion) AsImage() (v ImageItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsingGetResponseItemsPageStructuredResultPageItemUnion) AsLink() (v LinkItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsingGetResponseItemsPageStructuredResultPageItemUnion) AsHeader() (v HeaderItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsingGetResponseItemsPageStructuredResultPageItemUnion) AsFooter() (v FooterItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParsingGetResponseItemsPageStructuredResultPageItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ParsingGetResponseItemsPageStructuredResultPageItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParsingGetResponseItemsPageStructuredResultPageItemUnionItems is an implicit
// subunion of [ParsingGetResponseItemsPageStructuredResultPageItemUnion].
// ParsingGetResponseItemsPageStructuredResultPageItemUnionItems provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ParsingGetResponseItemsPageStructuredResultPageItemUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfItems]
type ParsingGetResponseItemsPageStructuredResultPageItemUnionItems struct {
	// This field will be present if the value is a [[]ListItemItemUnion] instead of an
	// object.
	OfItems []ListItemItemUnion `json:",inline"`
	JSON    struct {
		OfItems respjson.Field
		raw     string
	} `json:"-"`
}

func (r *ParsingGetResponseItemsPageStructuredResultPageItemUnionItems) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParsingGetResponseItemsPageFailedStructuredPage struct {
	// Error message describing the failure
	Error string `json:"error" api:"required"`
	// Page number of the document
	PageNumber int64 `json:"page_number" api:"required"`
	// Failure indicator
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		PageNumber  respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseItemsPageFailedStructuredPage) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseItemsPageFailedStructuredPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown result (if requested)
type ParsingGetResponseMarkdown struct {
	// List of markdown pages or failed page entries
	Pages []ParsingGetResponseMarkdownPageUnion `json:"pages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pages       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseMarkdown) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParsingGetResponseMarkdownPageUnion contains all possible properties and values
// from [ParsingGetResponseMarkdownPageMarkdownResultPage],
// [ParsingGetResponseMarkdownPageFailedMarkdownPage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ParsingGetResponseMarkdownPageUnion struct {
	// This field is from variant [ParsingGetResponseMarkdownPageMarkdownResultPage].
	Markdown   string `json:"markdown"`
	PageNumber int64  `json:"page_number"`
	Success    bool   `json:"success"`
	// This field is from variant [ParsingGetResponseMarkdownPageMarkdownResultPage].
	Footer string `json:"footer"`
	// This field is from variant [ParsingGetResponseMarkdownPageMarkdownResultPage].
	Header string `json:"header"`
	// This field is from variant [ParsingGetResponseMarkdownPageFailedMarkdownPage].
	Error string `json:"error"`
	JSON  struct {
		Markdown   respjson.Field
		PageNumber respjson.Field
		Success    respjson.Field
		Footer     respjson.Field
		Header     respjson.Field
		Error      respjson.Field
		raw        string
	} `json:"-"`
}

func (u ParsingGetResponseMarkdownPageUnion) AsMarkdownResultPage() (v ParsingGetResponseMarkdownPageMarkdownResultPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParsingGetResponseMarkdownPageUnion) AsFailedMarkdownPage() (v ParsingGetResponseMarkdownPageFailedMarkdownPage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParsingGetResponseMarkdownPageUnion) RawJSON() string { return u.JSON.raw }

func (r *ParsingGetResponseMarkdownPageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParsingGetResponseMarkdownPageMarkdownResultPage struct {
	// Markdown content of the page
	Markdown string `json:"markdown" api:"required"`
	// Page number of the document
	PageNumber int64 `json:"page_number" api:"required"`
	// Success indicator
	Success bool `json:"success" api:"required"`
	// Footer of the page in markdown
	Footer string `json:"footer" api:"nullable"`
	// Header of the page in markdown
	Header string `json:"header" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Markdown    respjson.Field
		PageNumber  respjson.Field
		Success     respjson.Field
		Footer      respjson.Field
		Header      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseMarkdownPageMarkdownResultPage) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseMarkdownPageMarkdownResultPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParsingGetResponseMarkdownPageFailedMarkdownPage struct {
	// Error message describing the failure
	Error string `json:"error" api:"required"`
	// Page number of the document
	PageNumber int64 `json:"page_number" api:"required"`
	// Failure indicator
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		PageNumber  respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseMarkdownPageFailedMarkdownPage) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseMarkdownPageFailedMarkdownPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result containing metadata (page level and general) for the parsed document.
type ParsingGetResponseMetadata struct {
	// List of page metadata entries
	Pages []ParsingGetResponseMetadataPage `json:"pages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pages       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page-level metadata including confidence scores and presentation-specific data.
type ParsingGetResponseMetadataPage struct {
	// Page number of the document
	PageNumber int64 `json:"page_number" api:"required"`
	// Confidence score for the page parsing (0-1)
	Confidence float64 `json:"confidence" api:"nullable"`
	// Whether cost-optimized parsing was used for the page
	CostOptimized bool `json:"cost_optimized" api:"nullable"`
	// Original orientation angle of the page in degrees
	OriginalOrientationAngle int64 `json:"original_orientation_angle" api:"nullable"`
	// Printed page number as it appears in the document
	PrintedPageNumber string `json:"printed_page_number" api:"nullable"`
	// Section name from presentation slides
	SlideSectionName string `json:"slide_section_name" api:"nullable"`
	// Speaker notes from presentation slides
	SpeakerNotes string `json:"speaker_notes" api:"nullable"`
	// Whether auto mode was triggered for the page
	TriggeredAutoMode bool `json:"triggered_auto_mode" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber               respjson.Field
		Confidence               respjson.Field
		CostOptimized            respjson.Field
		OriginalOrientationAngle respjson.Field
		PrintedPageNumber        respjson.Field
		SlideSectionName         respjson.Field
		SpeakerNotes             respjson.Field
		TriggeredAutoMode        respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseMetadataPage) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseMetadataPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about a specific result type stored in S3.
type ParsingGetResponseResultContentMetadata struct {
	// Size of the result file in bytes
	SizeBytes int64 `json:"size_bytes" api:"required"`
	// Whether the result file exists in S3
	Exists bool `json:"exists"`
	// Presigned URL to download the result file
	PresignedURL string `json:"presigned_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SizeBytes    respjson.Field
		Exists       respjson.Field
		PresignedURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseResultContentMetadata) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseResultContentMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Plain text result (if requested)
type ParsingGetResponseText struct {
	// List of text pages
	Pages []ParsingGetResponseTextPage `json:"pages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pages       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseText) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParsingGetResponseTextPage struct {
	// Page number of the document
	PageNumber int64 `json:"page_number" api:"required"`
	// Plain text content of the page
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber  respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParsingGetResponseTextPage) RawJSON() string { return r.JSON.raw }
func (r *ParsingGetResponseTextPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParsingNewParams struct {
	// Parsing tier: 'fast' (rule-based, cheapest), 'cost_effective' (balanced),
	// 'agentic' (AI-powered with custom prompts), or 'agentic_plus' (premium AI with
	// highest accuracy)
	//
	// Any of "fast", "cost_effective", "agentic", "agentic_plus".
	Tier ParsingNewParamsTier `json:"tier,omitzero" api:"required"`
	// Version for the selected tier. Use `latest`, or pin one of that tier's dated
	// versions.
	//
	// Current `latest` by tier:
	//
	// - `fast`: `2025-12-11`
	// - `cost_effective`: `2026-05-28`
	// - `agentic`: `2026-05-26`
	// - `agentic_plus`: `2026-05-26`
	//
	// Full list: `GET /api/v2/parse/versions`.
	Version        ParsingNewParamsVersion `json:"version,omitzero" api:"required"`
	OrganizationID param.Opt[string]       `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string]       `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Identifier for the client/application making the request. Used for analytics and
	// debugging. Example: 'my-app-v2'
	ClientName param.Opt[string] `json:"client_name,omitzero"`
	// Bypass result caching and force re-parsing. Use when document content may have
	// changed or you need fresh results
	DisableCache param.Opt[bool] `json:"disable_cache,omitzero"`
	// ID of an existing file in the project to parse. Mutually exclusive with
	// source_url
	FileID param.Opt[string] `json:"file_id,omitzero"`
	// HTTP/HTTPS proxy for fetching source_url. Ignored if using file_id
	HTTPProxy param.Opt[string] `json:"http_proxy,omitzero"`
	// Public URL of the document to parse. Mutually exclusive with file_id
	SourceURL param.Opt[string] `json:"source_url,omitzero"`
	// Options for AI-powered parsing tiers (cost_effective, agentic, agentic_plus).
	//
	// These options customize how the AI processes and interprets document content.
	// Only applicable when using non-fast tiers.
	AgenticOptions ParsingNewParamsAgenticOptions `json:"agentic_options,omitzero"`
	// Options for fast tier parsing (rule-based, no AI).
	//
	// Fast tier uses deterministic algorithms for text extraction without AI
	// enhancement. It's the fastest and most cost-effective option, best suited for
	// simple documents with standard layouts. Currently has no configurable options
	// but reserved for future expansion.
	FastOptions any `json:"fast_options,omitzero"`
	// Crop boundaries to process only a portion of each page. Values are ratios 0-1
	// from page edges
	CropBox ParsingNewParamsCropBox `json:"crop_box,omitzero"`
	// Format-specific options (HTML, PDF, spreadsheet, presentation). Applied based on
	// detected input file type
	InputOptions ParsingNewParamsInputOptions `json:"input_options,omitzero"`
	// Output formatting options for markdown, text, and extracted images
	OutputOptions ParsingNewParamsOutputOptions `json:"output_options,omitzero"`
	// Page selection: limit total pages or specify exact pages to process
	PageRanges ParsingNewParamsPageRanges `json:"page_ranges,omitzero"`
	// Job execution controls including timeouts and failure thresholds
	ProcessingControl ParsingNewParamsProcessingControl `json:"processing_control,omitzero"`
	// Document processing options including OCR, table extraction, and chart parsing
	ProcessingOptions ParsingNewParamsProcessingOptions `json:"processing_options,omitzero"`
	// Webhook endpoints for job status notifications. Multiple webhooks can be
	// configured for different events or services
	WebhookConfigurations []ParsingNewParamsWebhookConfiguration `json:"webhook_configurations,omitzero"`
	paramObj
}

func (r ParsingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [ParsingNewParams]'s query parameters as `url.Values`.
func (r ParsingNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Parsing tier: 'fast' (rule-based, cheapest), 'cost_effective' (balanced),
// 'agentic' (AI-powered with custom prompts), or 'agentic_plus' (premium AI with
// highest accuracy)
type ParsingNewParamsTier string

const (
	ParsingNewParamsTierFast          ParsingNewParamsTier = "fast"
	ParsingNewParamsTierCostEffective ParsingNewParamsTier = "cost_effective"
	ParsingNewParamsTierAgentic       ParsingNewParamsTier = "agentic"
	ParsingNewParamsTierAgenticPlus   ParsingNewParamsTier = "agentic_plus"
)

// Version for the selected tier. Use `latest`, or pin one of that tier's dated
// versions.
//
// Current `latest` by tier:
//
// - `fast`: `2025-12-11`
// - `cost_effective`: `2026-05-28`
// - `agentic`: `2026-05-26`
// - `agentic_plus`: `2026-05-26`
//
// Full list: `GET /api/v2/parse/versions`.
type ParsingNewParamsVersion string

const (
	ParsingNewParamsVersionLatest     ParsingNewParamsVersion = "latest"
	ParsingNewParamsVersion2026_05_28 ParsingNewParamsVersion = "2026-05-28"
	ParsingNewParamsVersion2026_05_26 ParsingNewParamsVersion = "2026-05-26"
	ParsingNewParamsVersion2025_12_11 ParsingNewParamsVersion = "2025-12-11"
)

// Options for AI-powered parsing tiers (cost_effective, agentic, agentic_plus).
//
// These options customize how the AI processes and interprets document content.
// Only applicable when using non-fast tiers.
type ParsingNewParamsAgenticOptions struct {
	// Custom instructions for the AI parser. Use to guide extraction behavior, specify
	// output formatting, or provide domain-specific context. Example: 'Extract
	// financial tables with currency symbols. Format dates as YYYY-MM-DD.'
	CustomPrompt param.Opt[string] `json:"custom_prompt,omitzero"`
	paramObj
}

func (r ParsingNewParamsAgenticOptions) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsAgenticOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsAgenticOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Crop boundaries to process only a portion of each page. Values are ratios 0-1
// from page edges
type ParsingNewParamsCropBox struct {
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

func (r ParsingNewParamsCropBox) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsCropBox
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsCropBox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Format-specific options (HTML, PDF, spreadsheet, presentation). Applied based on
// detected input file type
type ParsingNewParamsInputOptions struct {
	// HTML/web page parsing options (applies to .html, .htm files)
	HTML ParsingNewParamsInputOptionsHTML `json:"html,omitzero"`
	// PDF-specific parsing options (applies to .pdf files)
	Pdf any `json:"pdf,omitzero"`
	// Presentation parsing options (applies to .pptx, .ppt, .odp, .key files)
	Presentation ParsingNewParamsInputOptionsPresentation `json:"presentation,omitzero"`
	// Spreadsheet parsing options (applies to .xlsx, .xls, .csv, .ods files)
	Spreadsheet ParsingNewParamsInputOptionsSpreadsheet `json:"spreadsheet,omitzero"`
	paramObj
}

func (r ParsingNewParamsInputOptions) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsInputOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsInputOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTML/web page parsing options (applies to .html, .htm files)
type ParsingNewParamsInputOptionsHTML struct {
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

func (r ParsingNewParamsInputOptionsHTML) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsInputOptionsHTML
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsInputOptionsHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Presentation parsing options (applies to .pptx, .ppt, .odp, .key files)
type ParsingNewParamsInputOptionsPresentation struct {
	// Extract content positioned outside the visible slide area. Some presentations
	// have hidden notes or content that extends beyond slide boundaries
	OutOfBoundsContent param.Opt[bool] `json:"out_of_bounds_content,omitzero"`
	// Skip extraction of embedded chart data tables. When true, only the visual
	// representation of charts is captured, not the underlying data
	SkipEmbeddedData param.Opt[bool] `json:"skip_embedded_data,omitzero"`
	paramObj
}

func (r ParsingNewParamsInputOptionsPresentation) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsInputOptionsPresentation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsInputOptionsPresentation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spreadsheet parsing options (applies to .xlsx, .xls, .csv, .ods files)
type ParsingNewParamsInputOptionsSpreadsheet struct {
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

func (r ParsingNewParamsInputOptionsSpreadsheet) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsInputOptionsSpreadsheet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsInputOptionsSpreadsheet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Output formatting options for markdown, text, and extracted images
type ParsingNewParamsOutputOptions struct {
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
	// Image categories to extract and save. Options: 'screenshot' (full page renders
	// useful for visual QA), 'embedded' (images found within the document), 'layout'
	// (cropped regions from layout detection like figures and diagrams). Empty list
	// saves no images
	//
	// Any of "screenshot", "embedded", "layout".
	ImagesToSave []string `json:"images_to_save,omitzero"`
	// Markdown formatting options including table styles and link annotations
	Markdown ParsingNewParamsOutputOptionsMarkdown `json:"markdown,omitzero"`
	// Spatial text output options for preserving document layout structure
	SpatialText ParsingNewParamsOutputOptionsSpatialText `json:"spatial_text,omitzero"`
	// Options for exporting tables as XLSX spreadsheets
	TablesAsSpreadsheet ParsingNewParamsOutputOptionsTablesAsSpreadsheet `json:"tables_as_spreadsheet,omitzero"`
	paramObj
}

func (r ParsingNewParamsOutputOptions) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsOutputOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsOutputOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Markdown formatting options including table styles and link annotations
type ParsingNewParamsOutputOptionsMarkdown struct {
	// Add link annotations to markdown output in the format [text](url). When false,
	// only the link text is included
	AnnotateLinks param.Opt[bool] `json:"annotate_links,omitzero"`
	// Embed images directly in markdown as base64 data URIs instead of extracting them
	// as separate files. Useful for self-contained markdown output
	InlineImages param.Opt[bool] `json:"inline_images,omitzero"`
	// Table formatting options including markdown vs HTML format and merging behavior
	Tables ParsingNewParamsOutputOptionsMarkdownTables `json:"tables,omitzero"`
	paramObj
}

func (r ParsingNewParamsOutputOptionsMarkdown) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsOutputOptionsMarkdown
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsOutputOptionsMarkdown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Table formatting options including markdown vs HTML format and merging behavior
type ParsingNewParamsOutputOptionsMarkdownTables struct {
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

func (r ParsingNewParamsOutputOptionsMarkdownTables) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsOutputOptionsMarkdownTables
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsOutputOptionsMarkdownTables) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spatial text output options for preserving document layout structure
type ParsingNewParamsOutputOptionsSpatialText struct {
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

func (r ParsingNewParamsOutputOptionsSpatialText) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsOutputOptionsSpatialText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsOutputOptionsSpatialText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for exporting tables as XLSX spreadsheets
type ParsingNewParamsOutputOptionsTablesAsSpreadsheet struct {
	// Whether this option is enabled
	Enable param.Opt[bool] `json:"enable,omitzero"`
	// Automatically generate descriptive sheet names from table context (headers,
	// surrounding text) instead of using generic names like 'Table_1'
	GuessSheetName param.Opt[bool] `json:"guess_sheet_name,omitzero"`
	paramObj
}

func (r ParsingNewParamsOutputOptionsTablesAsSpreadsheet) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsOutputOptionsTablesAsSpreadsheet
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsOutputOptionsTablesAsSpreadsheet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page selection: limit total pages or specify exact pages to process
type ParsingNewParamsPageRanges struct {
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

func (r ParsingNewParamsPageRanges) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsPageRanges
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsPageRanges) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution controls including timeouts and failure thresholds
type ParsingNewParamsProcessingControl struct {
	// Quality thresholds that determine when a job should fail vs complete with
	// partial results
	JobFailureConditions ParsingNewParamsProcessingControlJobFailureConditions `json:"job_failure_conditions,omitzero"`
	// Timeout settings for job execution. Increase for large or complex documents
	Timeouts ParsingNewParamsProcessingControlTimeouts `json:"timeouts,omitzero"`
	paramObj
}

func (r ParsingNewParamsProcessingControl) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingControl
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingControl) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quality thresholds that determine when a job should fail vs complete with
// partial results
type ParsingNewParamsProcessingControlJobFailureConditions struct {
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

func (r ParsingNewParamsProcessingControlJobFailureConditions) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingControlJobFailureConditions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingControlJobFailureConditions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Timeout settings for job execution. Increase for large or complex documents
type ParsingNewParamsProcessingControlTimeouts struct {
	// Base timeout for the job in seconds (max 7200 = 2 hours). This is the minimum
	// time allowed regardless of document size
	BaseInSeconds param.Opt[int64] `json:"base_in_seconds,omitzero"`
	// Additional timeout per page in seconds (max 300 = 5 minutes). Total timeout =
	// base + (this value × page count)
	ExtraTimePerPageInSeconds param.Opt[int64] `json:"extra_time_per_page_in_seconds,omitzero"`
	paramObj
}

func (r ParsingNewParamsProcessingControlTimeouts) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingControlTimeouts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingControlTimeouts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Document processing options including OCR, table extraction, and chart parsing
type ParsingNewParamsProcessingOptions struct {
	// Use aggressive heuristics to detect table boundaries, even without visible
	// borders. Useful for documents with borderless or complex tables
	AggressiveTableExtraction param.Opt[bool] `json:"aggressive_table_extraction,omitzero"`
	// Disable automatic heuristics including outlined table extraction and adaptive
	// long table handling. Use when heuristics produce incorrect results
	DisableHeuristics param.Opt[bool] `json:"disable_heuristics,omitzero"`
	// Conditional processing rules that apply different parsing options based on page
	// content, document structure, or filename patterns. Each entry defines trigger
	// conditions and the parsing configuration to apply when triggered
	AutoModeConfiguration []ParsingNewParamsProcessingOptionsAutoModeConfiguration `json:"auto_mode_configuration,omitzero"`
	// Cost optimizer configuration for reducing parsing costs on simpler pages.
	//
	// When enabled, the parser analyzes each page and routes simpler pages to faster,
	// cheaper processing while preserving quality for complex pages. Only works with
	// 'agentic' or 'agentic_plus' tiers.
	CostOptimizer ParsingNewParamsProcessingOptionsCostOptimizer `json:"cost_optimizer,omitzero"`
	// Enable AI-powered chart analysis. Modes: 'efficient' (fast, lower cost),
	// 'agentic' (balanced), 'agentic_plus' (highest accuracy). Automatically enables
	// extract_layout and precise_bounding_box when set
	//
	// Any of "agentic_plus", "agentic", "efficient".
	SpecializedChartParsing string `json:"specialized_chart_parsing,omitzero"`
	// Options for ignoring specific text types (diagonal, hidden, text in images)
	Ignore ParsingNewParamsProcessingOptionsIgnore `json:"ignore,omitzero"`
	// OCR configuration including language detection settings
	OcrParameters ParsingNewParamsProcessingOptionsOcrParameters `json:"ocr_parameters,omitzero"`
	paramObj
}

func (r ParsingNewParamsProcessingOptions) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ParsingNewParamsProcessingOptions](
		"specialized_chart_parsing", "agentic_plus", "agentic", "efficient",
	)
}

// A single auto mode rule with trigger conditions and parsing configuration.
//
// Auto mode allows conditional parsing where different configurations are applied
// based on page content, structure, or filename. When triggers match, the
// parsing_conf overrides default settings for that page.
//
// The property ParsingConf is required.
type ParsingNewParamsProcessingOptionsAutoModeConfiguration struct {
	// Parsing configuration to apply when trigger conditions are met
	ParsingConf ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConf `json:"parsing_conf,omitzero" api:"required"`
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
	FullPageImageInPageThreshold ParsingNewParamsProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnion `json:"full_page_image_in_page_threshold,omitzero"`
	// Confidence threshold for layout element detection
	LayoutElementInPageConfidenceThreshold ParsingNewParamsProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnion `json:"layout_element_in_page_confidence_threshold,omitzero"`
	// Trigger if page has more than N charts
	PageContainsAtLeastNCharts ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnion `json:"page_contains_at_least_n_charts,omitzero"`
	// Trigger if page has more than N images
	PageContainsAtLeastNImages ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnion `json:"page_contains_at_least_n_images,omitzero"`
	// Trigger if page has more than N layout elements
	PageContainsAtLeastNLayoutElements ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnion `json:"page_contains_at_least_n_layout_elements,omitzero"`
	// Trigger if page has more than N lines
	PageContainsAtLeastNLines ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnion `json:"page_contains_at_least_n_lines,omitzero"`
	// Trigger if page has more than N links
	PageContainsAtLeastNLinks ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnion `json:"page_contains_at_least_n_links,omitzero"`
	// Trigger if page has more than N numeric words
	PageContainsAtLeastNNumbers ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnion `json:"page_contains_at_least_n_numbers,omitzero"`
	// Trigger if page has more than N% numeric words
	PageContainsAtLeastNPercentNumbers ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnion `json:"page_contains_at_least_n_percent_numbers,omitzero"`
	// Trigger if page has more than N tables
	PageContainsAtLeastNTables ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnion `json:"page_contains_at_least_n_tables,omitzero"`
	// Trigger if page has more than N words
	PageContainsAtLeastNWords ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnion `json:"page_contains_at_least_n_words,omitzero"`
	// Trigger if page has fewer than N charts
	PageContainsAtMostNCharts ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnion `json:"page_contains_at_most_n_charts,omitzero"`
	// Trigger if page has fewer than N images
	PageContainsAtMostNImages ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnion `json:"page_contains_at_most_n_images,omitzero"`
	// Trigger if page has fewer than N layout elements
	PageContainsAtMostNLayoutElements ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnion `json:"page_contains_at_most_n_layout_elements,omitzero"`
	// Trigger if page has fewer than N lines
	PageContainsAtMostNLines ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnion `json:"page_contains_at_most_n_lines,omitzero"`
	// Trigger if page has fewer than N links
	PageContainsAtMostNLinks ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnion `json:"page_contains_at_most_n_links,omitzero"`
	// Trigger if page has fewer than N numeric words
	PageContainsAtMostNNumbers ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnion `json:"page_contains_at_most_n_numbers,omitzero"`
	// Trigger if page has fewer than N% numeric words
	PageContainsAtMostNPercentNumbers ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnion `json:"page_contains_at_most_n_percent_numbers,omitzero"`
	// Trigger if page has fewer than N tables
	PageContainsAtMostNTables ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnion `json:"page_contains_at_most_n_tables,omitzero"`
	// Trigger if page has fewer than N words
	PageContainsAtMostNWords ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnion `json:"page_contains_at_most_n_words,omitzero"`
	// Trigger if page has more than N characters
	PageLongerThanNChars ParsingNewParamsProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnion `json:"page_longer_than_n_chars,omitzero"`
	// Trigger if page has fewer than N characters
	PageShorterThanNChars ParsingNewParamsProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnion `json:"page_shorter_than_n_chars,omitzero"`
	paramObj
}

func (r ParsingNewParamsProcessingOptionsAutoModeConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingOptionsAutoModeConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingOptionsAutoModeConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parsing configuration to apply when trigger conditions are met
type ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConf struct {
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
	CropBox ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfCropBox `json:"crop_box,omitzero"`
	// Ignore options for auto mode parsing configuration.
	Ignore ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfIgnore `json:"ignore,omitzero"`
	// Presentation-specific options for auto mode parsing configuration.
	Presentation ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfPresentation `json:"presentation,omitzero"`
	// Spatial text options for auto mode parsing configuration.
	SpatialText ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfSpatialText `json:"spatial_text,omitzero"`
	// Enable specialized chart parsing with the specified mode
	//
	// Any of "agentic_plus", "agentic", "efficient".
	SpecializedChartParsing string `json:"specialized_chart_parsing,omitzero"`
	// Override the parsing tier for matched pages. Must be paired with version
	//
	// Any of "fast", "cost_effective", "agentic", "agentic_plus".
	Tier string `json:"tier,omitzero"`
	// Version for the override tier. Required when `tier` is set. Use `latest`, or pin
	// one of that tier's dated versions.
	//
	// Current `latest` by tier:
	//
	// - `fast`: `2025-12-11`
	// - `cost_effective`: `2026-05-28`
	// - `agentic`: `2026-05-26`
	// - `agentic_plus`: `2026-05-26`
	//
	// Full list: `GET /api/v2/parse/versions`.
	Version string `json:"version,omitzero"`
	paramObj
}

func (r ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConf) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConf](
		"specialized_chart_parsing", "agentic_plus", "agentic", "efficient",
	)
	apijson.RegisterFieldValidator[ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConf](
		"tier", "fast", "cost_effective", "agentic", "agentic_plus",
	)
}

// Crop box options for auto mode parsing configuration.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfCropBox struct {
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

func (r ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfCropBox) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfCropBox
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfCropBox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ignore options for auto mode parsing configuration.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfIgnore struct {
	// Whether to ignore diagonal text in the document
	IgnoreDiagonalText param.Opt[bool] `json:"ignore_diagonal_text,omitzero"`
	// Whether to ignore hidden text in the document
	IgnoreHiddenText param.Opt[bool] `json:"ignore_hidden_text,omitzero"`
	paramObj
}

func (r ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfIgnore) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfIgnore
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfIgnore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Presentation-specific options for auto mode parsing configuration.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfPresentation struct {
	// Extract out of bounds content in presentation slides
	OutOfBoundsContent param.Opt[bool] `json:"out_of_bounds_content,omitzero"`
	// Skip extraction of embedded data for charts in presentation slides
	SkipEmbeddedData param.Opt[bool] `json:"skip_embedded_data,omitzero"`
	paramObj
}

func (r ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfPresentation) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfPresentation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfPresentation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spatial text options for auto mode parsing configuration.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfSpatialText struct {
	// Keep column structure intact without unrolling
	DoNotUnrollColumns param.Opt[bool] `json:"do_not_unroll_columns,omitzero"`
	// Preserve text alignment across page boundaries
	PreserveLayoutAlignmentAcrossPages param.Opt[bool] `json:"preserve_layout_alignment_across_pages,omitzero"`
	// Include very small text in spatial output
	PreserveVerySmallText param.Opt[bool] `json:"preserve_very_small_text,omitzero"`
	paramObj
}

func (r ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfSpatialText) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfSpatialText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingOptionsAutoModeConfigurationParsingConfSpatialText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationFullPageImageInPageThresholdUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationLayoutElementInPageConfidenceThresholdUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNChartsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNImagesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLayoutElementsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNLinksUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNNumbersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNPercentNumbersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNTablesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtLeastNWordsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNChartsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNImagesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLayoutElementsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNLinksUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNNumbersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNPercentNumbersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNTablesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageContainsAtMostNWordsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageLongerThanNCharsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ParsingNewParamsProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnion struct {
	OfInt    param.Opt[int64]  `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ParsingNewParamsProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfInt, u.OfString)
}
func (u *ParsingNewParamsProcessingOptionsAutoModeConfigurationPageShorterThanNCharsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Cost optimizer configuration for reducing parsing costs on simpler pages.
//
// When enabled, the parser analyzes each page and routes simpler pages to faster,
// cheaper processing while preserving quality for complex pages. Only works with
// 'agentic' or 'agentic_plus' tiers.
type ParsingNewParamsProcessingOptionsCostOptimizer struct {
	// Enable cost-optimized parsing. Routes simpler pages to faster processing while
	// complex pages use full AI analysis. May reduce speed on some documents.
	// IMPORTANT: Only available with 'agentic' or 'agentic_plus' tiers
	Enable param.Opt[bool] `json:"enable,omitzero"`
	paramObj
}

func (r ParsingNewParamsProcessingOptionsCostOptimizer) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingOptionsCostOptimizer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingOptionsCostOptimizer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for ignoring specific text types (diagonal, hidden, text in images)
type ParsingNewParamsProcessingOptionsIgnore struct {
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

func (r ParsingNewParamsProcessingOptionsIgnore) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingOptionsIgnore
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingOptionsIgnore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OCR configuration including language detection settings
type ParsingNewParamsProcessingOptionsOcrParameters struct {
	// Languages to use for OCR text recognition. Specify multiple languages if
	// document contains mixed-language content. Order matters - put primary language
	// first. Example: ['en', 'es'] for English with Spanish
	Languages []ParsingLanguages `json:"languages,omitzero"`
	paramObj
}

func (r ParsingNewParamsProcessingOptionsOcrParameters) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsProcessingOptionsOcrParameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsProcessingOptionsOcrParameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook configuration for receiving parsing job notifications.
//
// Webhooks are called when specified events occur during job processing. Configure
// multiple webhook configurations to send to different endpoints.
type ParsingNewParamsWebhookConfiguration struct {
	// HTTPS URL to receive webhook POST requests. Must be publicly accessible
	WebhookURL param.Opt[string] `json:"webhook_url,omitzero"`
	// Events that trigger this webhook. Options: 'parse.success' (job completed),
	// 'parse.failure' (job failed), 'parse.partial' (some pages failed). If not
	// specified, webhook fires for all events
	WebhookEvents []string `json:"webhook_events,omitzero"`
	// Custom HTTP headers to include in webhook requests. Use for authentication
	// tokens or custom routing. Example: {'Authorization': 'Bearer xyz'}
	WebhookHeaders map[string]any `json:"webhook_headers,omitzero"`
	paramObj
}

func (r ParsingNewParamsWebhookConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow ParsingNewParamsWebhookConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ParsingNewParamsWebhookConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParsingListParams struct {
	// Include items created at or after this timestamp (inclusive)
	CreatedAtOnOrAfter param.Opt[time.Time] `query:"created_at_on_or_after,omitzero" format:"date-time" json:"-"`
	// Include items created at or before this timestamp (inclusive)
	CreatedAtOnOrBefore param.Opt[time.Time] `query:"created_at_on_or_before,omitzero" format:"date-time" json:"-"`
	OrganizationID      param.Opt[string]    `query:"organization_id,omitzero" format:"uuid" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Token for pagination
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	ProjectID param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Filter by specific job IDs
	JobIDs []string `query:"job_ids,omitzero" json:"-"`
	// Filter by job status (PENDING, RUNNING, COMPLETED, FAILED, CANCELLED)
	//
	// Any of "PENDING", "RUNNING", "COMPLETED", "FAILED", "CANCELLED".
	Status ParsingListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ParsingListParams]'s query parameters as `url.Values`.
func (r ParsingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by job status (PENDING, RUNNING, COMPLETED, FAILED, CANCELLED)
type ParsingListParamsStatus string

const (
	ParsingListParamsStatusPending   ParsingListParamsStatus = "PENDING"
	ParsingListParamsStatusRunning   ParsingListParamsStatus = "RUNNING"
	ParsingListParamsStatusCompleted ParsingListParamsStatus = "COMPLETED"
	ParsingListParamsStatusFailed    ParsingListParamsStatus = "FAILED"
	ParsingListParamsStatusCancelled ParsingListParamsStatus = "CANCELLED"
)

type ParsingGetParams struct {
	// Filter to specific image filenames (optional). Example: image_0.png,image_1.jpg
	ImageFilenames param.Opt[string] `query:"image_filenames,omitzero" json:"-"`
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Fields to include: text, markdown, items, metadata, job_metadata,
	// text_content_metadata, markdown_content_metadata, items_content_metadata,
	// metadata_content_metadata, raw_words_content_metadata, xlsx_content_metadata,
	// output_pdf_content_metadata, images_content_metadata. Metadata fields include
	// presigned URLs.
	Expand []string `query:"expand,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ParsingGetParams]'s query parameters as `url.Values`.
func (r ParsingGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
