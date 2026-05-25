// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/llamacloud-prod-go/internal/apijson"
	shimjson "github.com/stainless-sdks/llamacloud-prod-go/internal/encoding/json"
	"github.com/stainless-sdks/llamacloud-prod-go/internal/requestconfig"
	"github.com/stainless-sdks/llamacloud-prod-go/option"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/param"
	"github.com/stainless-sdks/llamacloud-prod-go/packages/respjson"
	"github.com/stainless-sdks/llamacloud-prod-go/shared"
)

// PipelineDataSourceService contains methods and other services that help with
// interacting with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPipelineDataSourceService] method instead.
type PipelineDataSourceService struct {
	options []option.RequestOption
}

// NewPipelineDataSourceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPipelineDataSourceService(opts ...option.RequestOption) (r PipelineDataSourceService) {
	r = PipelineDataSourceService{}
	r.options = opts
	return
}

// Update the configuration of a data source in a pipeline.
func (r *PipelineDataSourceService) Update(ctx context.Context, dataSourceID string, params PipelineDataSourceUpdateParams, opts ...option.RequestOption) (res *PipelineDataSource, err error) {
	opts = slices.Concat(r.options, opts)
	if params.PipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	if dataSourceID == "" {
		err = errors.New("missing required data_source_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/data-sources/%s", params.PipelineID, dataSourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Get data sources for a pipeline.
func (r *PipelineDataSourceService) GetDataSources(ctx context.Context, pipelineID string, opts ...option.RequestOption) (res *[]PipelineDataSource, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/data-sources", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the status of a data source for a pipeline.
func (r *PipelineDataSourceService) GetStatus(ctx context.Context, dataSourceID string, query PipelineDataSourceGetStatusParams, opts ...option.RequestOption) (res *ManagedIngestionStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.PipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	if dataSourceID == "" {
		err = errors.New("missing required data_source_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/data-sources/%s/status", query.PipelineID, dataSourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Run incremental ingestion: pull upstream changes from the data source into the
// data sink.
func (r *PipelineDataSourceService) Sync(ctx context.Context, dataSourceID string, params PipelineDataSourceSyncParams, opts ...option.RequestOption) (res *Pipeline, err error) {
	opts = slices.Concat(r.options, opts)
	if params.PipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	if dataSourceID == "" {
		err = errors.New("missing required data_source_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/data-sources/%s/sync", params.PipelineID, dataSourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Add data sources to a pipeline.
func (r *PipelineDataSourceService) UpdateDataSources(ctx context.Context, pipelineID string, body PipelineDataSourceUpdateDataSourcesParams, opts ...option.RequestOption) (res *[]PipelineDataSource, err error) {
	opts = slices.Concat(r.options, opts)
	if pipelineID == "" {
		err = errors.New("missing required pipeline_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/pipelines/%s/data-sources", pipelineID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Schema for a data source in a pipeline.
type PipelineDataSource struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Component that implements the data source
	Component PipelineDataSourceComponentUnion `json:"component" api:"required"`
	// The ID of the data source.
	DataSourceID string `json:"data_source_id" api:"required" format:"uuid"`
	// The last time the data source was automatically synced.
	LastSyncedAt time.Time `json:"last_synced_at" api:"required" format:"date-time"`
	// The name of the data source.
	Name string `json:"name" api:"required"`
	// The ID of the pipeline.
	PipelineID string `json:"pipeline_id" api:"required" format:"uuid"`
	ProjectID  string `json:"project_id" api:"required" format:"uuid"`
	// Any of "S3", "AZURE_STORAGE_BLOB", "GOOGLE_DRIVE", "MICROSOFT_ONEDRIVE",
	// "MICROSOFT_SHAREPOINT", "SLACK", "NOTION_PAGE", "CONFLUENCE", "JIRA", "JIRA_V2",
	// "BOX".
	SourceType PipelineDataSourceSourceType `json:"source_type" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Custom metadata that will be present on all data loaded from the data source
	CustomMetadata map[string]*PipelineDataSourceCustomMetadataUnion `json:"custom_metadata" api:"nullable"`
	// The status of the data source in the pipeline.
	//
	// Any of "NOT_STARTED", "IN_PROGRESS", "SUCCESS", "ERROR", "CANCELLED".
	Status PipelineDataSourceStatus `json:"status" api:"nullable"`
	// The last time the status was updated.
	StatusUpdatedAt time.Time `json:"status_updated_at" api:"nullable" format:"date-time"`
	// The interval at which the data source should be synced.
	SyncInterval float64 `json:"sync_interval" api:"nullable"`
	// The id of the user who set the sync schedule.
	SyncScheduleSetBy string `json:"sync_schedule_set_by" api:"nullable"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// Version metadata for the data source
	VersionMetadata DataSourceReaderVersionMetadata `json:"version_metadata" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Component         respjson.Field
		DataSourceID      respjson.Field
		LastSyncedAt      respjson.Field
		Name              respjson.Field
		PipelineID        respjson.Field
		ProjectID         respjson.Field
		SourceType        respjson.Field
		CreatedAt         respjson.Field
		CustomMetadata    respjson.Field
		Status            respjson.Field
		StatusUpdatedAt   respjson.Field
		SyncInterval      respjson.Field
		SyncScheduleSetBy respjson.Field
		UpdatedAt         respjson.Field
		VersionMetadata   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PipelineDataSource) RawJSON() string { return r.JSON.raw }
func (r *PipelineDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PipelineDataSourceComponentUnion contains all possible properties and values
// from [map[string]any], [shared.CloudS3DataSource],
// [shared.CloudAzStorageBlobDataSource], [shared.CloudGoogleDriveDataSource],
// [shared.CloudOneDriveDataSource], [shared.CloudSharepointDataSource],
// [shared.CloudSlackDataSource], [shared.CloudNotionPageDataSource],
// [shared.CloudConfluenceDataSource], [shared.CloudJiraDataSource],
// [shared.CloudJiraDataSourceV2], [shared.CloudBoxDataSource].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPipelineDataSourceComponentMapItem]
type PipelineDataSourceComponentUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPipelineDataSourceComponentMapItem any `json:",inline"`
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
		OfPipelineDataSourceComponentMapItem respjson.Field
		Bucket                               respjson.Field
		AwsAccessID                          respjson.Field
		AwsAccessSecret                      respjson.Field
		ClassName                            respjson.Field
		Prefix                               respjson.Field
		RegexPattern                         respjson.Field
		S3EndpointURL                        respjson.Field
		SupportsAccessControl                respjson.Field
		AccountURL                           respjson.Field
		ContainerName                        respjson.Field
		AccountKey                           respjson.Field
		AccountName                          respjson.Field
		Blob                                 respjson.Field
		ClientID                             respjson.Field
		ClientSecret                         respjson.Field
		TenantID                             respjson.Field
		FolderID                             respjson.Field
		ServiceAccountKey                    respjson.Field
		UserPrincipalName                    respjson.Field
		FolderPath                           respjson.Field
		RequiredExts                         respjson.Field
		DriveName                            respjson.Field
		ExcludePathPatterns                  respjson.Field
		GetPermissions                       respjson.Field
		IncludePathPatterns                  respjson.Field
		SiteID                               respjson.Field
		SiteName                             respjson.Field
		SlackToken                           respjson.Field
		ChannelIDs                           respjson.Field
		ChannelPatterns                      respjson.Field
		EarliestDate                         respjson.Field
		EarliestDateTimestamp                respjson.Field
		LatestDate                           respjson.Field
		LatestDateTimestamp                  respjson.Field
		IntegrationToken                     respjson.Field
		DatabaseIDs                          respjson.Field
		PageIDs                              respjson.Field
		AuthenticationMechanism              respjson.Field
		ServerURL                            respjson.Field
		APIToken                             respjson.Field
		Cql                                  respjson.Field
		FailureHandling                      respjson.Field
		IndexRestrictedPages                 respjson.Field
		KeepMarkdownFormat                   respjson.Field
		Label                                respjson.Field
		SpaceKey                             respjson.Field
		UserName                             respjson.Field
		Query                                respjson.Field
		CloudID                              respjson.Field
		Email                                respjson.Field
		APIVersion                           respjson.Field
		Expand                               respjson.Field
		Fields                               respjson.Field
		RequestsPerMinute                    respjson.Field
		DeveloperToken                       respjson.Field
		EnterpriseID                         respjson.Field
		UserID                               respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u PipelineDataSourceComponentUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceComponentUnion) AsCloudS3DataSource() (v shared.CloudS3DataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceComponentUnion) AsCloudAzStorageBlobDataSource() (v shared.CloudAzStorageBlobDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceComponentUnion) AsCloudGoogleDriveDataSource() (v shared.CloudGoogleDriveDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceComponentUnion) AsCloudOneDriveDataSource() (v shared.CloudOneDriveDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceComponentUnion) AsCloudSharepointDataSource() (v shared.CloudSharepointDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceComponentUnion) AsCloudSlackDataSource() (v shared.CloudSlackDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceComponentUnion) AsCloudNotionPageDataSource() (v shared.CloudNotionPageDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceComponentUnion) AsCloudConfluenceDataSource() (v shared.CloudConfluenceDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceComponentUnion) AsCloudJiraDataSource() (v shared.CloudJiraDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceComponentUnion) AsCloudJiraDataSourceV2() (v shared.CloudJiraDataSourceV2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceComponentUnion) AsCloudBoxDataSource() (v shared.CloudBoxDataSource) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PipelineDataSourceComponentUnion) RawJSON() string { return u.JSON.raw }

func (r *PipelineDataSourceComponentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineDataSourceSourceType string

const (
	PipelineDataSourceSourceTypeS3                  PipelineDataSourceSourceType = "S3"
	PipelineDataSourceSourceTypeAzureStorageBlob    PipelineDataSourceSourceType = "AZURE_STORAGE_BLOB"
	PipelineDataSourceSourceTypeGoogleDrive         PipelineDataSourceSourceType = "GOOGLE_DRIVE"
	PipelineDataSourceSourceTypeMicrosoftOnedrive   PipelineDataSourceSourceType = "MICROSOFT_ONEDRIVE"
	PipelineDataSourceSourceTypeMicrosoftSharepoint PipelineDataSourceSourceType = "MICROSOFT_SHAREPOINT"
	PipelineDataSourceSourceTypeSlack               PipelineDataSourceSourceType = "SLACK"
	PipelineDataSourceSourceTypeNotionPage          PipelineDataSourceSourceType = "NOTION_PAGE"
	PipelineDataSourceSourceTypeConfluence          PipelineDataSourceSourceType = "CONFLUENCE"
	PipelineDataSourceSourceTypeJira                PipelineDataSourceSourceType = "JIRA"
	PipelineDataSourceSourceTypeJiraV2              PipelineDataSourceSourceType = "JIRA_V2"
	PipelineDataSourceSourceTypeBox                 PipelineDataSourceSourceType = "BOX"
)

// PipelineDataSourceCustomMetadataUnion contains all possible properties and
// values from [map[string]any], [[]any], [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfPipelineDataSourceCustomMetadataMapItem OfAnyArray OfString
// OfFloat OfBool]
type PipelineDataSourceCustomMetadataUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfPipelineDataSourceCustomMetadataMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfPipelineDataSourceCustomMetadataMapItem respjson.Field
		OfAnyArray                                respjson.Field
		OfString                                  respjson.Field
		OfFloat                                   respjson.Field
		OfBool                                    respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u PipelineDataSourceCustomMetadataUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceCustomMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceCustomMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceCustomMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PipelineDataSourceCustomMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PipelineDataSourceCustomMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *PipelineDataSourceCustomMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the data source in the pipeline.
type PipelineDataSourceStatus string

const (
	PipelineDataSourceStatusNotStarted PipelineDataSourceStatus = "NOT_STARTED"
	PipelineDataSourceStatusInProgress PipelineDataSourceStatus = "IN_PROGRESS"
	PipelineDataSourceStatusSuccess    PipelineDataSourceStatus = "SUCCESS"
	PipelineDataSourceStatusError      PipelineDataSourceStatus = "ERROR"
	PipelineDataSourceStatusCancelled  PipelineDataSourceStatus = "CANCELLED"
)

type PipelineDataSourceUpdateParams struct {
	PipelineID string `path:"pipeline_id" api:"required" format:"uuid" json:"-"`
	// The interval at which the data source should be synced.
	SyncInterval param.Opt[float64] `json:"sync_interval,omitzero"`
	paramObj
}

func (r PipelineDataSourceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PipelineDataSourceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineDataSourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineDataSourceGetStatusParams struct {
	PipelineID string `path:"pipeline_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type PipelineDataSourceSyncParams struct {
	PipelineID      string   `path:"pipeline_id" api:"required" format:"uuid" json:"-"`
	PipelineFileIDs []string `json:"pipeline_file_ids,omitzero" format:"uuid"`
	paramObj
}

func (r PipelineDataSourceSyncParams) MarshalJSON() (data []byte, err error) {
	type shadow PipelineDataSourceSyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineDataSourceSyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PipelineDataSourceUpdateDataSourcesParams struct {
	Body []PipelineDataSourceUpdateDataSourcesParamsBody
	paramObj
}

func (r PipelineDataSourceUpdateDataSourcesParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *PipelineDataSourceUpdateDataSourcesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for creating an association between a data source and a pipeline.
//
// The property DataSourceID is required.
type PipelineDataSourceUpdateDataSourcesParamsBody struct {
	// The ID of the data source.
	DataSourceID string `json:"data_source_id" api:"required" format:"uuid"`
	// The interval at which the data source should be synced. Valid values are: 21600,
	// 43200, 86400
	SyncInterval param.Opt[float64] `json:"sync_interval,omitzero"`
	paramObj
}

func (r PipelineDataSourceUpdateDataSourcesParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow PipelineDataSourceUpdateDataSourcesParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PipelineDataSourceUpdateDataSourcesParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
