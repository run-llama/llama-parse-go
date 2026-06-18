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
	"github.com/run-llama/llama-parse-go/internal/requestconfig"
	"github.com/run-llama/llama-parse-go/option"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
	"github.com/run-llama/llama-parse-go/shared"
)

// DataSourceService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataSourceService] method instead.
type DataSourceService struct {
	options []option.RequestOption
}

// NewDataSourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataSourceService(opts ...option.RequestOption) (r DataSourceService) {
	r = DataSourceService{}
	r.options = opts
	return
}

// Create a new data source.
func (r *DataSourceService) New(ctx context.Context, params DataSourceNewParams, opts ...option.RequestOption) (res *DataSource, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/data-sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a data source by ID.
func (r *DataSourceService) Update(ctx context.Context, dataSourceID string, body DataSourceUpdateParams, opts ...option.RequestOption) (res *DataSource, err error) {
	opts = slices.Concat(r.options, opts)
	if dataSourceID == "" {
		err = errors.New("missing required data_source_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/data-sources/%s", dataSourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List data sources for a given project. If project_id is not provided, uses the
// default project.
func (r *DataSourceService) List(ctx context.Context, query DataSourceListParams, opts ...option.RequestOption) (res *[]DataSource, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/data-sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a data source by ID.
func (r *DataSourceService) Delete(ctx context.Context, dataSourceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if dataSourceID == "" {
		err = errors.New("missing required data_source_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/data-sources/%s", dataSourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a data source by ID.
func (r *DataSourceService) Get(ctx context.Context, dataSourceID string, opts ...option.RequestOption) (res *DataSource, err error) {
	opts = slices.Concat(r.options, opts)
	if dataSourceID == "" {
		err = errors.New("missing required data_source_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/data-sources/%s", dataSourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Schema for a data source.
type DataSource struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Component that implements the data source
	Component DataSourceComponentUnion `json:"component" api:"required"`
	// The name of the data source.
	Name      string `json:"name" api:"required"`
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// Any of "S3", "AZURE_STORAGE_BLOB", "GOOGLE_DRIVE", "MICROSOFT_ONEDRIVE",
	// "MICROSOFT_SHAREPOINT", "SLACK", "NOTION_PAGE", "CONFLUENCE", "JIRA", "JIRA_V2",
	// "BOX".
	SourceType DataSourceSourceType `json:"source_type" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Custom metadata that will be present on all data loaded from the data source
	CustomMetadata map[string]*DataSourceCustomMetadataUnion `json:"custom_metadata" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Version metadata for the data source
	VersionMetadata DataSourceReaderVersionMetadata `json:"version_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Component       respjson.Field
		Name            respjson.Field
		ProjectID       respjson.Field
		SourceType      respjson.Field
		CreatedAt       respjson.Field
		CustomMetadata  respjson.Field
		UpdatedAt       respjson.Field
		VersionMetadata respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataSource) RawJSON() string { return r.JSON.raw }
func (r *DataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataSourceComponentUnion contains all possible properties and values from
// [map[string]any], [shared.CloudS3DataSource],
// [shared.CloudAzStorageBlobDataSource], [shared.CloudGoogleDriveDataSource],
// [shared.CloudOneDriveDataSource], [shared.CloudSharepointDataSource],
// [shared.CloudSlackDataSource], [shared.CloudNotionPageDataSource],
// [shared.CloudConfluenceDataSource], [shared.CloudJiraDataSource],
// [shared.CloudJiraDataSourceV2], [shared.CloudBoxDataSource].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfDataSourceComponentMapItem]
type DataSourceComponentUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfDataSourceComponentMapItem any `json:",inline"`
	// This field is from variant [shared.CloudS3DataSource].
	Bucket string `json:"bucket"`
	// This field is from variant [shared.CloudS3DataSource].
	AwsAccessID string `json:"aws_access_id"`
	// This field is from variant [shared.CloudS3DataSource].
	AwsAccessSecret string `json:"aws_access_secret"`
	ClassName       string `json:"class_name"`
	Prefix          string `json:"prefix"`
	// This field is from variant [shared.CloudS3DataSource].
	RegexPattern string `json:"regex_pattern"`
	// This field is from variant [shared.CloudS3DataSource].
	S3EndpointURL         string `json:"s3_endpoint_url"`
	SupportsAccessControl bool   `json:"supports_access_control"`
	// This field is from variant [shared.CloudAzStorageBlobDataSource].
	AccountURL string `json:"account_url"`
	// This field is from variant [shared.CloudAzStorageBlobDataSource].
	ContainerName string `json:"container_name"`
	// This field is from variant [shared.CloudAzStorageBlobDataSource].
	AccountKey string `json:"account_key"`
	// This field is from variant [shared.CloudAzStorageBlobDataSource].
	AccountName string `json:"account_name"`
	// This field is from variant [shared.CloudAzStorageBlobDataSource].
	Blob         string `json:"blob"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	TenantID     string `json:"tenant_id"`
	FolderID     string `json:"folder_id"`
	// This field is from variant [shared.CloudGoogleDriveDataSource].
	ServiceAccountKey map[string]string `json:"service_account_key"`
	// This field is from variant [shared.CloudOneDriveDataSource].
	UserPrincipalName string   `json:"user_principal_name"`
	FolderPath        string   `json:"folder_path"`
	RequiredExts      []string `json:"required_exts"`
	// This field is from variant [shared.CloudSharepointDataSource].
	DriveName string `json:"drive_name"`
	// This field is from variant [shared.CloudSharepointDataSource].
	ExcludePathPatterns []string `json:"exclude_path_patterns"`
	GetPermissions      bool     `json:"get_permissions"`
	// This field is from variant [shared.CloudSharepointDataSource].
	IncludePathPatterns []string `json:"include_path_patterns"`
	// This field is from variant [shared.CloudSharepointDataSource].
	SiteID string `json:"site_id"`
	// This field is from variant [shared.CloudSharepointDataSource].
	SiteName string `json:"site_name"`
	// This field is from variant [shared.CloudSlackDataSource].
	SlackToken string `json:"slack_token"`
	// This field is from variant [shared.CloudSlackDataSource].
	ChannelIDs string `json:"channel_ids"`
	// This field is from variant [shared.CloudSlackDataSource].
	ChannelPatterns string `json:"channel_patterns"`
	// This field is from variant [shared.CloudSlackDataSource].
	EarliestDate string `json:"earliest_date"`
	// This field is from variant [shared.CloudSlackDataSource].
	EarliestDateTimestamp float64 `json:"earliest_date_timestamp"`
	// This field is from variant [shared.CloudSlackDataSource].
	LatestDate string `json:"latest_date"`
	// This field is from variant [shared.CloudSlackDataSource].
	LatestDateTimestamp float64 `json:"latest_date_timestamp"`
	// This field is from variant [shared.CloudNotionPageDataSource].
	IntegrationToken string `json:"integration_token"`
	// This field is from variant [shared.CloudNotionPageDataSource].
	DatabaseIDs             string `json:"database_ids"`
	PageIDs                 string `json:"page_ids"`
	AuthenticationMechanism string `json:"authentication_mechanism"`
	ServerURL               string `json:"server_url"`
	APIToken                string `json:"api_token"`
	// This field is from variant [shared.CloudConfluenceDataSource].
	Cql string `json:"cql"`
	// This field is from variant [shared.CloudConfluenceDataSource].
	FailureHandling shared.FailureHandlingConfig `json:"failure_handling"`
	// This field is from variant [shared.CloudConfluenceDataSource].
	IndexRestrictedPages bool `json:"index_restricted_pages"`
	// This field is from variant [shared.CloudConfluenceDataSource].
	KeepMarkdownFormat bool `json:"keep_markdown_format"`
	// This field is from variant [shared.CloudConfluenceDataSource].
	Label string `json:"label"`
	// This field is from variant [shared.CloudConfluenceDataSource].
	SpaceKey string `json:"space_key"`
	// This field is from variant [shared.CloudConfluenceDataSource].
	SyncPermissions bool `json:"sync_permissions"`
	// This field is from variant [shared.CloudConfluenceDataSource].
	UserName string `json:"user_name"`
	Query    string `json:"query"`
	CloudID  string `json:"cloud_id"`
	Email    string `json:"email"`
	// This field is from variant [shared.CloudJiraDataSourceV2].
	APIVersion shared.CloudJiraDataSourceV2APIVersion `json:"api_version"`
	// This field is from variant [shared.CloudJiraDataSourceV2].
	Expand string `json:"expand"`
	// This field is from variant [shared.CloudJiraDataSourceV2].
	Fields []string `json:"fields"`
	// This field is from variant [shared.CloudJiraDataSourceV2].
	RequestsPerMinute int64 `json:"requests_per_minute"`
	// This field is from variant [shared.CloudBoxDataSource].
	DeveloperToken string `json:"developer_token"`
	// This field is from variant [shared.CloudBoxDataSource].
	EnterpriseID string `json:"enterprise_id"`
	// This field is from variant [shared.CloudBoxDataSource].
	UserID string `json:"user_id"`
	JSON   struct {
		OfDataSourceComponentMapItem respjson.Field
		Bucket                       respjson.Field
		AwsAccessID                  respjson.Field
		AwsAccessSecret              respjson.Field
		ClassName                    respjson.Field
		Prefix                       respjson.Field
		RegexPattern                 respjson.Field
		S3EndpointURL                respjson.Field
		SupportsAccessControl        respjson.Field
		AccountURL                   respjson.Field
		ContainerName                respjson.Field
		AccountKey                   respjson.Field
		AccountName                  respjson.Field
		Blob                         respjson.Field
		ClientID                     respjson.Field
		ClientSecret                 respjson.Field
		TenantID                     respjson.Field
		FolderID                     respjson.Field
		ServiceAccountKey            respjson.Field
		UserPrincipalName            respjson.Field
		FolderPath                   respjson.Field
		RequiredExts                 respjson.Field
		DriveName                    respjson.Field
		ExcludePathPatterns          respjson.Field
		GetPermissions               respjson.Field
		IncludePathPatterns          respjson.Field
		SiteID                       respjson.Field
		SiteName                     respjson.Field
		SlackToken                   respjson.Field
		ChannelIDs                   respjson.Field
		ChannelPatterns              respjson.Field
		EarliestDate                 respjson.Field
		EarliestDateTimestamp        respjson.Field
		LatestDate                   respjson.Field
		LatestDateTimestamp          respjson.Field
		IntegrationToken             respjson.Field
		DatabaseIDs                  respjson.Field
		PageIDs                      respjson.Field
		AuthenticationMechanism      respjson.Field
		ServerURL                    respjson.Field
		APIToken                     respjson.Field
		Cql                          respjson.Field
		FailureHandling              respjson.Field
		IndexRestrictedPages         respjson.Field
		KeepMarkdownFormat           respjson.Field
		Label                        respjson.Field
		SpaceKey                     respjson.Field
		SyncPermissions              respjson.Field
		UserName                     respjson.Field
		Query                        respjson.Field
		CloudID                      respjson.Field
		Email                        respjson.Field
		APIVersion                   respjson.Field
		Expand                       respjson.Field
		Fields                       respjson.Field
		RequestsPerMinute            respjson.Field
		DeveloperToken               respjson.Field
		EnterpriseID                 respjson.Field
		UserID                       respjson.Field
		raw                          string
	} `json:"-"`
}

func (u DataSourceComponentUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceComponentUnion) AsCloudS3DataSource() (v shared.CloudS3DataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceComponentUnion) AsCloudAzStorageBlobDataSource() (v shared.CloudAzStorageBlobDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceComponentUnion) AsCloudGoogleDriveDataSource() (v shared.CloudGoogleDriveDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceComponentUnion) AsCloudOneDriveDataSource() (v shared.CloudOneDriveDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceComponentUnion) AsCloudSharepointDataSource() (v shared.CloudSharepointDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceComponentUnion) AsCloudSlackDataSource() (v shared.CloudSlackDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceComponentUnion) AsCloudNotionPageDataSource() (v shared.CloudNotionPageDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceComponentUnion) AsCloudConfluenceDataSource() (v shared.CloudConfluenceDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceComponentUnion) AsCloudJiraDataSource() (v shared.CloudJiraDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceComponentUnion) AsCloudJiraDataSourceV2() (v shared.CloudJiraDataSourceV2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceComponentUnion) AsCloudBoxDataSource() (v shared.CloudBoxDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataSourceComponentUnion) RawJSON() string { return u.JSON.raw }

func (r *DataSourceComponentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataSourceSourceType string

const (
	DataSourceSourceTypeS3                  DataSourceSourceType = "S3"
	DataSourceSourceTypeAzureStorageBlob    DataSourceSourceType = "AZURE_STORAGE_BLOB"
	DataSourceSourceTypeGoogleDrive         DataSourceSourceType = "GOOGLE_DRIVE"
	DataSourceSourceTypeMicrosoftOnedrive   DataSourceSourceType = "MICROSOFT_ONEDRIVE"
	DataSourceSourceTypeMicrosoftSharepoint DataSourceSourceType = "MICROSOFT_SHAREPOINT"
	DataSourceSourceTypeSlack               DataSourceSourceType = "SLACK"
	DataSourceSourceTypeNotionPage          DataSourceSourceType = "NOTION_PAGE"
	DataSourceSourceTypeConfluence          DataSourceSourceType = "CONFLUENCE"
	DataSourceSourceTypeJira                DataSourceSourceType = "JIRA"
	DataSourceSourceTypeJiraV2              DataSourceSourceType = "JIRA_V2"
	DataSourceSourceTypeBox                 DataSourceSourceType = "BOX"
)

// DataSourceCustomMetadataUnion contains all possible properties and values from
// [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfDataSourceCustomMetadataMapItem OfAnyArray OfString OfFloat
// OfBool]
type DataSourceCustomMetadataUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfDataSourceCustomMetadataMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfDataSourceCustomMetadataMapItem respjson.Field
		OfAnyArray                        respjson.Field
		OfString                          respjson.Field
		OfFloat                           respjson.Field
		OfBool                            respjson.Field
		raw                               string
	} `json:"-"`
}

func (u DataSourceCustomMetadataUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceCustomMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceCustomMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceCustomMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSourceCustomMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataSourceCustomMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *DataSourceCustomMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataSourceReaderVersionMetadata struct {
	// The version of the reader to use for this data source.
	//
	// Any of "1.0", "2.0", "2.1".
	ReaderVersion DataSourceReaderVersionMetadataReaderVersion `json:"reader_version" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReaderVersion respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataSourceReaderVersionMetadata) RawJSON() string { return r.JSON.raw }
func (r *DataSourceReaderVersionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The version of the reader to use for this data source.
type DataSourceReaderVersionMetadataReaderVersion string

const (
	DataSourceReaderVersionMetadataReaderVersion1_0 DataSourceReaderVersionMetadataReaderVersion = "1.0"
	DataSourceReaderVersionMetadataReaderVersion2_0 DataSourceReaderVersionMetadataReaderVersion = "2.0"
	DataSourceReaderVersionMetadataReaderVersion2_1 DataSourceReaderVersionMetadataReaderVersion = "2.1"
)

type DataSourceNewParams struct {
	// Component that implements the data source
	Component DataSourceNewParamsComponentUnion `json:"component,omitzero" api:"required"`
	// The name of the data source.
	Name string `json:"name" api:"required"`
	// Any of "S3", "AZURE_STORAGE_BLOB", "GOOGLE_DRIVE", "MICROSOFT_ONEDRIVE",
	// "MICROSOFT_SHAREPOINT", "SLACK", "NOTION_PAGE", "CONFLUENCE", "JIRA", "JIRA_V2",
	// "BOX".
	SourceType     DataSourceNewParamsSourceType `json:"source_type,omitzero" api:"required"`
	OrganizationID param.Opt[string]             `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string]             `query:"project_id,omitzero" format:"uuid" json:"-"`
	// Custom metadata that will be present on all data loaded from the data source
	CustomMetadata map[string]*DataSourceNewParamsCustomMetadataUnion `json:"custom_metadata,omitzero"`
	paramObj
}

func (r DataSourceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DataSourceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataSourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [DataSourceNewParams]'s query parameters as `url.Values`.
func (r DataSourceNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DataSourceNewParamsComponentUnion struct {
	OfAnyMap                       map[string]any                            `json:",omitzero,inline"`
	OfCloudS3DataSource            *shared.CloudS3DataSourceParam            `json:",omitzero,inline"`
	OfCloudAzStorageBlobDataSource *shared.CloudAzStorageBlobDataSourceParam `json:",omitzero,inline"`
	OfCloudGoogleDriveDataSource   *shared.CloudGoogleDriveDataSourceParam   `json:",omitzero,inline"`
	OfCloudOneDriveDataSource      *shared.CloudOneDriveDataSourceParam      `json:",omitzero,inline"`
	OfCloudSharepointDataSource    *shared.CloudSharepointDataSourceParam    `json:",omitzero,inline"`
	OfCloudSlackDataSource         *shared.CloudSlackDataSourceParam         `json:",omitzero,inline"`
	OfCloudNotionPageDataSource    *shared.CloudNotionPageDataSourceParam    `json:",omitzero,inline"`
	OfCloudConfluenceDataSource    *shared.CloudConfluenceDataSourceParam    `json:",omitzero,inline"`
	OfCloudJiraDataSource          *shared.CloudJiraDataSourceParam          `json:",omitzero,inline"`
	OfCloudJiraDataSourceV2        *shared.CloudJiraDataSourceV2Param        `json:",omitzero,inline"`
	OfCloudBoxDataSource           *shared.CloudBoxDataSourceParam           `json:",omitzero,inline"`
	paramUnion
}

func (u DataSourceNewParamsComponentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfCloudS3DataSource,
		u.OfCloudAzStorageBlobDataSource,
		u.OfCloudGoogleDriveDataSource,
		u.OfCloudOneDriveDataSource,
		u.OfCloudSharepointDataSource,
		u.OfCloudSlackDataSource,
		u.OfCloudNotionPageDataSource,
		u.OfCloudConfluenceDataSource,
		u.OfCloudJiraDataSource,
		u.OfCloudJiraDataSourceV2,
		u.OfCloudBoxDataSource)
}
func (u *DataSourceNewParamsComponentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type DataSourceNewParamsSourceType string

const (
	DataSourceNewParamsSourceTypeS3                  DataSourceNewParamsSourceType = "S3"
	DataSourceNewParamsSourceTypeAzureStorageBlob    DataSourceNewParamsSourceType = "AZURE_STORAGE_BLOB"
	DataSourceNewParamsSourceTypeGoogleDrive         DataSourceNewParamsSourceType = "GOOGLE_DRIVE"
	DataSourceNewParamsSourceTypeMicrosoftOnedrive   DataSourceNewParamsSourceType = "MICROSOFT_ONEDRIVE"
	DataSourceNewParamsSourceTypeMicrosoftSharepoint DataSourceNewParamsSourceType = "MICROSOFT_SHAREPOINT"
	DataSourceNewParamsSourceTypeSlack               DataSourceNewParamsSourceType = "SLACK"
	DataSourceNewParamsSourceTypeNotionPage          DataSourceNewParamsSourceType = "NOTION_PAGE"
	DataSourceNewParamsSourceTypeConfluence          DataSourceNewParamsSourceType = "CONFLUENCE"
	DataSourceNewParamsSourceTypeJira                DataSourceNewParamsSourceType = "JIRA"
	DataSourceNewParamsSourceTypeJiraV2              DataSourceNewParamsSourceType = "JIRA_V2"
	DataSourceNewParamsSourceTypeBox                 DataSourceNewParamsSourceType = "BOX"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DataSourceNewParamsCustomMetadataUnion struct {
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u DataSourceNewParamsCustomMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfAnyArray,
		u.OfString,
		u.OfFloat,
		u.OfBool)
}
func (u *DataSourceNewParamsCustomMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type DataSourceUpdateParams struct {
	// Any of "S3", "AZURE_STORAGE_BLOB", "GOOGLE_DRIVE", "MICROSOFT_ONEDRIVE",
	// "MICROSOFT_SHAREPOINT", "SLACK", "NOTION_PAGE", "CONFLUENCE", "JIRA", "JIRA_V2",
	// "BOX".
	SourceType DataSourceUpdateParamsSourceType `json:"source_type,omitzero" api:"required"`
	// The name of the data source.
	Name param.Opt[string] `json:"name,omitzero"`
	// Component that implements the data source
	Component DataSourceUpdateParamsComponentUnion `json:"component,omitzero"`
	// Custom metadata that will be present on all data loaded from the data source
	CustomMetadata map[string]*DataSourceUpdateParamsCustomMetadataUnion `json:"custom_metadata,omitzero"`
	paramObj
}

func (r DataSourceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DataSourceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataSourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataSourceUpdateParamsSourceType string

const (
	DataSourceUpdateParamsSourceTypeS3                  DataSourceUpdateParamsSourceType = "S3"
	DataSourceUpdateParamsSourceTypeAzureStorageBlob    DataSourceUpdateParamsSourceType = "AZURE_STORAGE_BLOB"
	DataSourceUpdateParamsSourceTypeGoogleDrive         DataSourceUpdateParamsSourceType = "GOOGLE_DRIVE"
	DataSourceUpdateParamsSourceTypeMicrosoftOnedrive   DataSourceUpdateParamsSourceType = "MICROSOFT_ONEDRIVE"
	DataSourceUpdateParamsSourceTypeMicrosoftSharepoint DataSourceUpdateParamsSourceType = "MICROSOFT_SHAREPOINT"
	DataSourceUpdateParamsSourceTypeSlack               DataSourceUpdateParamsSourceType = "SLACK"
	DataSourceUpdateParamsSourceTypeNotionPage          DataSourceUpdateParamsSourceType = "NOTION_PAGE"
	DataSourceUpdateParamsSourceTypeConfluence          DataSourceUpdateParamsSourceType = "CONFLUENCE"
	DataSourceUpdateParamsSourceTypeJira                DataSourceUpdateParamsSourceType = "JIRA"
	DataSourceUpdateParamsSourceTypeJiraV2              DataSourceUpdateParamsSourceType = "JIRA_V2"
	DataSourceUpdateParamsSourceTypeBox                 DataSourceUpdateParamsSourceType = "BOX"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DataSourceUpdateParamsComponentUnion struct {
	OfAnyMap                       map[string]any                            `json:",omitzero,inline"`
	OfCloudS3DataSource            *shared.CloudS3DataSourceParam            `json:",omitzero,inline"`
	OfCloudAzStorageBlobDataSource *shared.CloudAzStorageBlobDataSourceParam `json:",omitzero,inline"`
	OfCloudGoogleDriveDataSource   *shared.CloudGoogleDriveDataSourceParam   `json:",omitzero,inline"`
	OfCloudOneDriveDataSource      *shared.CloudOneDriveDataSourceParam      `json:",omitzero,inline"`
	OfCloudSharepointDataSource    *shared.CloudSharepointDataSourceParam    `json:",omitzero,inline"`
	OfCloudSlackDataSource         *shared.CloudSlackDataSourceParam         `json:",omitzero,inline"`
	OfCloudNotionPageDataSource    *shared.CloudNotionPageDataSourceParam    `json:",omitzero,inline"`
	OfCloudConfluenceDataSource    *shared.CloudConfluenceDataSourceParam    `json:",omitzero,inline"`
	OfCloudJiraDataSource          *shared.CloudJiraDataSourceParam          `json:",omitzero,inline"`
	OfCloudJiraDataSourceV2        *shared.CloudJiraDataSourceV2Param        `json:",omitzero,inline"`
	OfCloudBoxDataSource           *shared.CloudBoxDataSourceParam           `json:",omitzero,inline"`
	paramUnion
}

func (u DataSourceUpdateParamsComponentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfCloudS3DataSource,
		u.OfCloudAzStorageBlobDataSource,
		u.OfCloudGoogleDriveDataSource,
		u.OfCloudOneDriveDataSource,
		u.OfCloudSharepointDataSource,
		u.OfCloudSlackDataSource,
		u.OfCloudNotionPageDataSource,
		u.OfCloudConfluenceDataSource,
		u.OfCloudJiraDataSource,
		u.OfCloudJiraDataSourceV2,
		u.OfCloudBoxDataSource)
}
func (u *DataSourceUpdateParamsComponentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DataSourceUpdateParamsCustomMetadataUnion struct {
	OfAnyMap   map[string]any     `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u DataSourceUpdateParamsCustomMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfAnyArray,
		u.OfString,
		u.OfFloat,
		u.OfBool)
}
func (u *DataSourceUpdateParamsCustomMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type DataSourceListParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [DataSourceListParams]'s query parameters as `url.Values`.
func (r DataSourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
