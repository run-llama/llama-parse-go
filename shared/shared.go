// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/run-llama/llama-parse-go/internal/apijson"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Cloud AstraDB Vector Store.
//
// This class is used to store the configuration for an AstraDB vector store, so
// that it can be created and used in LlamaCloud.
//
// Args: token (str): The Astra DB Application Token to use. api_endpoint (str):
// The Astra DB JSON API endpoint for your database. collection_name (str):
// Collection name to use. If not existing, it will be created. embedding_dimension
// (int): Length of the embedding vectors in use. keyspace (optional[str]): The
// keyspace to use. If not provided, 'default_keyspace'
type CloudAstraDBVectorStore struct {
	// The Astra DB Application Token to use
	Token string `json:"token" api:"required" format:"password"`
	// The Astra DB JSON API endpoint for your database
	APIEndpoint string `json:"api_endpoint" api:"required"`
	// Collection name to use. If not existing, it will be created
	CollectionName string `json:"collection_name" api:"required"`
	// Length of the embedding vectors in use
	EmbeddingDimension int64  `json:"embedding_dimension" api:"required"`
	ClassName          string `json:"class_name"`
	// The keyspace to use. If not provided, 'default_keyspace'
	Keyspace string `json:"keyspace" api:"nullable"`
	// Any of true.
	SupportsNestedMetadataFilters bool `json:"supports_nested_metadata_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token                         respjson.Field
		APIEndpoint                   respjson.Field
		CollectionName                respjson.Field
		EmbeddingDimension            respjson.Field
		ClassName                     respjson.Field
		Keyspace                      respjson.Field
		SupportsNestedMetadataFilters respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudAstraDBVectorStore) RawJSON() string { return r.JSON.raw }
func (r *CloudAstraDBVectorStore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudAstraDBVectorStore to a CloudAstraDBVectorStoreParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudAstraDBVectorStoreParam.Overrides()
func (r CloudAstraDBVectorStore) ToParam() CloudAstraDBVectorStoreParam {
	return param.Override[CloudAstraDBVectorStoreParam](json.RawMessage(r.RawJSON()))
}

// Cloud AstraDB Vector Store.
//
// This class is used to store the configuration for an AstraDB vector store, so
// that it can be created and used in LlamaCloud.
//
// Args: token (str): The Astra DB Application Token to use. api_endpoint (str):
// The Astra DB JSON API endpoint for your database. collection_name (str):
// Collection name to use. If not existing, it will be created. embedding_dimension
// (int): Length of the embedding vectors in use. keyspace (optional[str]): The
// keyspace to use. If not provided, 'default_keyspace'
//
// The properties Token, APIEndpoint, CollectionName, EmbeddingDimension are
// required.
type CloudAstraDBVectorStoreParam struct {
	// The Astra DB Application Token to use
	Token string `json:"token" api:"required" format:"password"`
	// The Astra DB JSON API endpoint for your database
	APIEndpoint string `json:"api_endpoint" api:"required"`
	// Collection name to use. If not existing, it will be created
	CollectionName string `json:"collection_name" api:"required"`
	// Length of the embedding vectors in use
	EmbeddingDimension int64 `json:"embedding_dimension" api:"required"`
	// The keyspace to use. If not provided, 'default_keyspace'
	Keyspace  param.Opt[string] `json:"keyspace,omitzero"`
	ClassName param.Opt[string] `json:"class_name,omitzero"`
	// Any of true.
	SupportsNestedMetadataFilters bool `json:"supports_nested_metadata_filters,omitzero"`
	paramObj
}

func (r CloudAstraDBVectorStoreParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudAstraDBVectorStoreParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudAstraDBVectorStoreParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CloudAstraDBVectorStoreParam](
		"supports_nested_metadata_filters", true,
	)
}

type CloudAzStorageBlobDataSource struct {
	// The Azure Storage Blob account URL to use for authentication.
	AccountURL string `json:"account_url" api:"required"`
	// The name of the Azure Storage Blob container to read from.
	ContainerName string `json:"container_name" api:"required"`
	// The Azure Storage Blob account key to use for authentication.
	AccountKey string `json:"account_key" api:"nullable" format:"password"`
	// The Azure Storage Blob account name to use for authentication.
	AccountName string `json:"account_name" api:"nullable"`
	// The blob name to read from.
	Blob      string `json:"blob" api:"nullable"`
	ClassName string `json:"class_name"`
	// The Azure AD client ID to use for authentication.
	ClientID string `json:"client_id" api:"nullable"`
	// The Azure AD client secret to use for authentication.
	ClientSecret string `json:"client_secret" api:"nullable" format:"password"`
	// The prefix of the Azure Storage Blob objects to read from.
	Prefix                string `json:"prefix" api:"nullable"`
	SupportsAccessControl bool   `json:"supports_access_control"`
	// The Azure AD tenant ID to use for authentication.
	TenantID string `json:"tenant_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountURL            respjson.Field
		ContainerName         respjson.Field
		AccountKey            respjson.Field
		AccountName           respjson.Field
		Blob                  respjson.Field
		ClassName             respjson.Field
		ClientID              respjson.Field
		ClientSecret          respjson.Field
		Prefix                respjson.Field
		SupportsAccessControl respjson.Field
		TenantID              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudAzStorageBlobDataSource) RawJSON() string { return r.JSON.raw }
func (r *CloudAzStorageBlobDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudAzStorageBlobDataSource to a
// CloudAzStorageBlobDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudAzStorageBlobDataSourceParam.Overrides()
func (r CloudAzStorageBlobDataSource) ToParam() CloudAzStorageBlobDataSourceParam {
	return param.Override[CloudAzStorageBlobDataSourceParam](json.RawMessage(r.RawJSON()))
}

// The properties AccountURL, ContainerName are required.
type CloudAzStorageBlobDataSourceParam struct {
	// The Azure Storage Blob account URL to use for authentication.
	AccountURL string `json:"account_url" api:"required"`
	// The name of the Azure Storage Blob container to read from.
	ContainerName string `json:"container_name" api:"required"`
	// The Azure Storage Blob account key to use for authentication.
	AccountKey param.Opt[string] `json:"account_key,omitzero" format:"password"`
	// The Azure Storage Blob account name to use for authentication.
	AccountName param.Opt[string] `json:"account_name,omitzero"`
	// The blob name to read from.
	Blob param.Opt[string] `json:"blob,omitzero"`
	// The Azure AD client ID to use for authentication.
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// The Azure AD client secret to use for authentication.
	ClientSecret param.Opt[string] `json:"client_secret,omitzero" format:"password"`
	// The prefix of the Azure Storage Blob objects to read from.
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// The Azure AD tenant ID to use for authentication.
	TenantID              param.Opt[string] `json:"tenant_id,omitzero"`
	ClassName             param.Opt[string] `json:"class_name,omitzero"`
	SupportsAccessControl param.Opt[bool]   `json:"supports_access_control,omitzero"`
	paramObj
}

func (r CloudAzStorageBlobDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudAzStorageBlobDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudAzStorageBlobDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cloud Azure AI Search Vector Store.
type CloudAzureAISearchVectorStore struct {
	SearchServiceAPIKey         string         `json:"search_service_api_key" api:"required" format:"password"`
	SearchServiceEndpoint       string         `json:"search_service_endpoint" api:"required"`
	ClassName                   string         `json:"class_name"`
	ClientID                    string         `json:"client_id" api:"nullable"`
	ClientSecret                string         `json:"client_secret" api:"nullable" format:"password"`
	EmbeddingDimension          int64          `json:"embedding_dimension" api:"nullable"`
	FilterableMetadataFieldKeys map[string]any `json:"filterable_metadata_field_keys" api:"nullable"`
	IndexName                   string         `json:"index_name" api:"nullable"`
	SearchServiceAPIVersion     string         `json:"search_service_api_version" api:"nullable"`
	// Any of true.
	SupportsNestedMetadataFilters bool   `json:"supports_nested_metadata_filters"`
	TenantID                      string `json:"tenant_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SearchServiceAPIKey           respjson.Field
		SearchServiceEndpoint         respjson.Field
		ClassName                     respjson.Field
		ClientID                      respjson.Field
		ClientSecret                  respjson.Field
		EmbeddingDimension            respjson.Field
		FilterableMetadataFieldKeys   respjson.Field
		IndexName                     respjson.Field
		SearchServiceAPIVersion       respjson.Field
		SupportsNestedMetadataFilters respjson.Field
		TenantID                      respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudAzureAISearchVectorStore) RawJSON() string { return r.JSON.raw }
func (r *CloudAzureAISearchVectorStore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudAzureAISearchVectorStore to a
// CloudAzureAISearchVectorStoreParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudAzureAISearchVectorStoreParam.Overrides()
func (r CloudAzureAISearchVectorStore) ToParam() CloudAzureAISearchVectorStoreParam {
	return param.Override[CloudAzureAISearchVectorStoreParam](json.RawMessage(r.RawJSON()))
}

// Cloud Azure AI Search Vector Store.
//
// The properties SearchServiceAPIKey, SearchServiceEndpoint are required.
type CloudAzureAISearchVectorStoreParam struct {
	SearchServiceAPIKey         string            `json:"search_service_api_key" api:"required" format:"password"`
	SearchServiceEndpoint       string            `json:"search_service_endpoint" api:"required"`
	ClientID                    param.Opt[string] `json:"client_id,omitzero"`
	ClientSecret                param.Opt[string] `json:"client_secret,omitzero" format:"password"`
	EmbeddingDimension          param.Opt[int64]  `json:"embedding_dimension,omitzero"`
	IndexName                   param.Opt[string] `json:"index_name,omitzero"`
	SearchServiceAPIVersion     param.Opt[string] `json:"search_service_api_version,omitzero"`
	TenantID                    param.Opt[string] `json:"tenant_id,omitzero"`
	ClassName                   param.Opt[string] `json:"class_name,omitzero"`
	FilterableMetadataFieldKeys map[string]any    `json:"filterable_metadata_field_keys,omitzero"`
	// Any of true.
	SupportsNestedMetadataFilters bool `json:"supports_nested_metadata_filters,omitzero"`
	paramObj
}

func (r CloudAzureAISearchVectorStoreParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudAzureAISearchVectorStoreParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudAzureAISearchVectorStoreParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CloudAzureAISearchVectorStoreParam](
		"supports_nested_metadata_filters", true,
	)
}

type CloudBoxDataSource struct {
	// The type of authentication to use (Developer Token or CCG)
	//
	// Any of "ccg", "developer_token".
	AuthenticationMechanism CloudBoxDataSourceAuthenticationMechanism `json:"authentication_mechanism" api:"required"`
	ClassName               string                                    `json:"class_name"`
	// Box API key used for identifying the application the user is authenticating with
	ClientID string `json:"client_id" api:"nullable"`
	// Box API secret used for making auth requests.
	ClientSecret string `json:"client_secret" api:"nullable" format:"password"`
	// Developer token for authentication if authentication_mechanism is
	// 'developer_token'.
	DeveloperToken string `json:"developer_token" api:"nullable" format:"password"`
	// Box Enterprise ID, if provided authenticates as service.
	EnterpriseID string `json:"enterprise_id" api:"nullable"`
	// The ID of the Box folder to read from.
	FolderID              string `json:"folder_id" api:"nullable"`
	SupportsAccessControl bool   `json:"supports_access_control"`
	// Box User ID, if provided authenticates as user.
	UserID string `json:"user_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthenticationMechanism respjson.Field
		ClassName               respjson.Field
		ClientID                respjson.Field
		ClientSecret            respjson.Field
		DeveloperToken          respjson.Field
		EnterpriseID            respjson.Field
		FolderID                respjson.Field
		SupportsAccessControl   respjson.Field
		UserID                  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudBoxDataSource) RawJSON() string { return r.JSON.raw }
func (r *CloudBoxDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudBoxDataSource to a CloudBoxDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudBoxDataSourceParam.Overrides()
func (r CloudBoxDataSource) ToParam() CloudBoxDataSourceParam {
	return param.Override[CloudBoxDataSourceParam](json.RawMessage(r.RawJSON()))
}

// The type of authentication to use (Developer Token or CCG)
type CloudBoxDataSourceAuthenticationMechanism string

const (
	CloudBoxDataSourceAuthenticationMechanismCcg            CloudBoxDataSourceAuthenticationMechanism = "ccg"
	CloudBoxDataSourceAuthenticationMechanismDeveloperToken CloudBoxDataSourceAuthenticationMechanism = "developer_token"
)

// The property AuthenticationMechanism is required.
type CloudBoxDataSourceParam struct {
	// The type of authentication to use (Developer Token or CCG)
	//
	// Any of "ccg", "developer_token".
	AuthenticationMechanism CloudBoxDataSourceAuthenticationMechanism `json:"authentication_mechanism,omitzero" api:"required"`
	// Box API key used for identifying the application the user is authenticating with
	ClientID param.Opt[string] `json:"client_id,omitzero"`
	// Box API secret used for making auth requests.
	ClientSecret param.Opt[string] `json:"client_secret,omitzero" format:"password"`
	// Developer token for authentication if authentication_mechanism is
	// 'developer_token'.
	DeveloperToken param.Opt[string] `json:"developer_token,omitzero" format:"password"`
	// Box Enterprise ID, if provided authenticates as service.
	EnterpriseID param.Opt[string] `json:"enterprise_id,omitzero"`
	// The ID of the Box folder to read from.
	FolderID param.Opt[string] `json:"folder_id,omitzero"`
	// Box User ID, if provided authenticates as user.
	UserID                param.Opt[string] `json:"user_id,omitzero"`
	ClassName             param.Opt[string] `json:"class_name,omitzero"`
	SupportsAccessControl param.Opt[bool]   `json:"supports_access_control,omitzero"`
	paramObj
}

func (r CloudBoxDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudBoxDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudBoxDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CloudConfluenceDataSource struct {
	// Type of Authentication for connecting to Confluence APIs.
	AuthenticationMechanism string `json:"authentication_mechanism" api:"required"`
	// The server URL of the Confluence instance.
	ServerURL string `json:"server_url" api:"required"`
	// The API token to use for authentication.
	APIToken  string `json:"api_token" api:"nullable" format:"password"`
	ClassName string `json:"class_name"`
	// The CQL query to use for fetching pages.
	Cql string `json:"cql" api:"nullable"`
	// Configuration for handling failures during processing. Key-value object
	// controlling failure handling behaviors.
	//
	// Example: { "skip_list_failures": true }
	//
	// Currently supports:
	//
	// - skip_list_failures: Skip failed batches/lists and continue processing
	FailureHandling FailureHandlingConfig `json:"failure_handling"`
	// Whether to index restricted pages.
	IndexRestrictedPages bool `json:"index_restricted_pages"`
	// Whether to keep the markdown format.
	KeepMarkdownFormat bool `json:"keep_markdown_format"`
	// The label to use for fetching pages.
	Label string `json:"label" api:"nullable"`
	// The page IDs of the Confluence to read from.
	PageIDs string `json:"page_ids" api:"nullable"`
	// The space key to read from.
	SpaceKey              string `json:"space_key" api:"nullable"`
	SupportsAccessControl bool   `json:"supports_access_control"`
	// Whether to fetch space-level permissions (allowed users/groups) and attach them
	// to document metadata for access control. Disable for Confluence Server/Data
	// Center versions whose permission APIs are unavailable (e.g. the JSON-RPC API
	// removed in Data Center 9.2.6+), which otherwise surface as 401 errors during
	// sync.
	SyncPermissions bool `json:"sync_permissions"`
	// The username to use for authentication.
	UserName string `json:"user_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthenticationMechanism respjson.Field
		ServerURL               respjson.Field
		APIToken                respjson.Field
		ClassName               respjson.Field
		Cql                     respjson.Field
		FailureHandling         respjson.Field
		IndexRestrictedPages    respjson.Field
		KeepMarkdownFormat      respjson.Field
		Label                   respjson.Field
		PageIDs                 respjson.Field
		SpaceKey                respjson.Field
		SupportsAccessControl   respjson.Field
		SyncPermissions         respjson.Field
		UserName                respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudConfluenceDataSource) RawJSON() string { return r.JSON.raw }
func (r *CloudConfluenceDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudConfluenceDataSource to a
// CloudConfluenceDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudConfluenceDataSourceParam.Overrides()
func (r CloudConfluenceDataSource) ToParam() CloudConfluenceDataSourceParam {
	return param.Override[CloudConfluenceDataSourceParam](json.RawMessage(r.RawJSON()))
}

// The properties AuthenticationMechanism, ServerURL are required.
type CloudConfluenceDataSourceParam struct {
	// Type of Authentication for connecting to Confluence APIs.
	AuthenticationMechanism string `json:"authentication_mechanism" api:"required"`
	// The server URL of the Confluence instance.
	ServerURL string `json:"server_url" api:"required"`
	// The API token to use for authentication.
	APIToken param.Opt[string] `json:"api_token,omitzero" format:"password"`
	// The CQL query to use for fetching pages.
	Cql param.Opt[string] `json:"cql,omitzero"`
	// The label to use for fetching pages.
	Label param.Opt[string] `json:"label,omitzero"`
	// The page IDs of the Confluence to read from.
	PageIDs param.Opt[string] `json:"page_ids,omitzero"`
	// The space key to read from.
	SpaceKey param.Opt[string] `json:"space_key,omitzero"`
	// The username to use for authentication.
	UserName  param.Opt[string] `json:"user_name,omitzero"`
	ClassName param.Opt[string] `json:"class_name,omitzero"`
	// Whether to index restricted pages.
	IndexRestrictedPages param.Opt[bool] `json:"index_restricted_pages,omitzero"`
	// Whether to keep the markdown format.
	KeepMarkdownFormat    param.Opt[bool] `json:"keep_markdown_format,omitzero"`
	SupportsAccessControl param.Opt[bool] `json:"supports_access_control,omitzero"`
	// Whether to fetch space-level permissions (allowed users/groups) and attach them
	// to document metadata for access control. Disable for Confluence Server/Data
	// Center versions whose permission APIs are unavailable (e.g. the JSON-RPC API
	// removed in Data Center 9.2.6+), which otherwise surface as 401 errors during
	// sync.
	SyncPermissions param.Opt[bool] `json:"sync_permissions,omitzero"`
	// Configuration for handling failures during processing. Key-value object
	// controlling failure handling behaviors.
	//
	// Example: { "skip_list_failures": true }
	//
	// Currently supports:
	//
	// - skip_list_failures: Skip failed batches/lists and continue processing
	FailureHandling FailureHandlingConfigParam `json:"failure_handling,omitzero"`
	paramObj
}

func (r CloudConfluenceDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudConfluenceDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudConfluenceDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CloudGoogleDriveDataSource struct {
	// The ID of the Google Drive folder to read from.
	FolderID  string `json:"folder_id" api:"required"`
	ClassName string `json:"class_name"`
	// A dictionary containing secret values
	ServiceAccountKey     map[string]string `json:"service_account_key" api:"nullable"`
	SupportsAccessControl bool              `json:"supports_access_control"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FolderID              respjson.Field
		ClassName             respjson.Field
		ServiceAccountKey     respjson.Field
		SupportsAccessControl respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudGoogleDriveDataSource) RawJSON() string { return r.JSON.raw }
func (r *CloudGoogleDriveDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudGoogleDriveDataSource to a
// CloudGoogleDriveDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudGoogleDriveDataSourceParam.Overrides()
func (r CloudGoogleDriveDataSource) ToParam() CloudGoogleDriveDataSourceParam {
	return param.Override[CloudGoogleDriveDataSourceParam](json.RawMessage(r.RawJSON()))
}

// The property FolderID is required.
type CloudGoogleDriveDataSourceParam struct {
	// The ID of the Google Drive folder to read from.
	FolderID              string            `json:"folder_id" api:"required"`
	ClassName             param.Opt[string] `json:"class_name,omitzero"`
	SupportsAccessControl param.Opt[bool]   `json:"supports_access_control,omitzero"`
	// A dictionary containing secret values
	ServiceAccountKey map[string]string `json:"service_account_key,omitzero"`
	paramObj
}

func (r CloudGoogleDriveDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudGoogleDriveDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudGoogleDriveDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cloud Jira Data Source integrating JiraReader.
type CloudJiraDataSource struct {
	// Type of Authentication for connecting to Jira APIs.
	AuthenticationMechanism string `json:"authentication_mechanism" api:"required"`
	// JQL (Jira Query Language) query to search.
	Query string `json:"query" api:"required"`
	// The API/ Access Token used for Basic, PAT and OAuth2 authentication.
	APIToken  string `json:"api_token" api:"nullable" format:"password"`
	ClassName string `json:"class_name"`
	// The cloud ID, used in case of OAuth2.
	CloudID string `json:"cloud_id" api:"nullable"`
	// The email address to use for authentication.
	Email string `json:"email" api:"nullable"`
	// The server url for Jira Cloud.
	ServerURL             string `json:"server_url" api:"nullable"`
	SupportsAccessControl bool   `json:"supports_access_control"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthenticationMechanism respjson.Field
		Query                   respjson.Field
		APIToken                respjson.Field
		ClassName               respjson.Field
		CloudID                 respjson.Field
		Email                   respjson.Field
		ServerURL               respjson.Field
		SupportsAccessControl   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudJiraDataSource) RawJSON() string { return r.JSON.raw }
func (r *CloudJiraDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudJiraDataSource to a CloudJiraDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudJiraDataSourceParam.Overrides()
func (r CloudJiraDataSource) ToParam() CloudJiraDataSourceParam {
	return param.Override[CloudJiraDataSourceParam](json.RawMessage(r.RawJSON()))
}

// Cloud Jira Data Source integrating JiraReader.
//
// The properties AuthenticationMechanism, Query are required.
type CloudJiraDataSourceParam struct {
	// Type of Authentication for connecting to Jira APIs.
	AuthenticationMechanism string `json:"authentication_mechanism" api:"required"`
	// JQL (Jira Query Language) query to search.
	Query string `json:"query" api:"required"`
	// The API/ Access Token used for Basic, PAT and OAuth2 authentication.
	APIToken param.Opt[string] `json:"api_token,omitzero" format:"password"`
	// The cloud ID, used in case of OAuth2.
	CloudID param.Opt[string] `json:"cloud_id,omitzero"`
	// The email address to use for authentication.
	Email param.Opt[string] `json:"email,omitzero"`
	// The server url for Jira Cloud.
	ServerURL             param.Opt[string] `json:"server_url,omitzero"`
	ClassName             param.Opt[string] `json:"class_name,omitzero"`
	SupportsAccessControl param.Opt[bool]   `json:"supports_access_control,omitzero"`
	paramObj
}

func (r CloudJiraDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudJiraDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudJiraDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cloud Jira Data Source integrating JiraReaderV2.
type CloudJiraDataSourceV2 struct {
	// Type of Authentication for connecting to Jira APIs.
	AuthenticationMechanism string `json:"authentication_mechanism" api:"required"`
	// JQL (Jira Query Language) query to search.
	Query string `json:"query" api:"required"`
	// The server url for Jira Cloud.
	ServerURL string `json:"server_url" api:"required"`
	// The API Access Token used for Basic, PAT and OAuth2 authentication.
	APIToken string `json:"api_token" api:"nullable" format:"password"`
	// Jira REST API version to use (2 or 3). 3 supports Atlassian Document Format
	// (ADF).
	//
	// Any of "2", "3".
	APIVersion CloudJiraDataSourceV2APIVersion `json:"api_version"`
	ClassName  string                          `json:"class_name"`
	// The cloud ID, used in case of OAuth2.
	CloudID string `json:"cloud_id" api:"nullable"`
	// The email address to use for authentication.
	Email string `json:"email" api:"nullable"`
	// Fields to expand in the response.
	Expand string `json:"expand" api:"nullable"`
	// List of fields to retrieve from Jira. If None, retrieves all fields.
	Fields []string `json:"fields" api:"nullable"`
	// Whether to fetch project role permissions and issue-level security
	GetPermissions bool `json:"get_permissions"`
	// Rate limit for Jira API requests per minute.
	RequestsPerMinute     int64 `json:"requests_per_minute" api:"nullable"`
	SupportsAccessControl bool  `json:"supports_access_control"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthenticationMechanism respjson.Field
		Query                   respjson.Field
		ServerURL               respjson.Field
		APIToken                respjson.Field
		APIVersion              respjson.Field
		ClassName               respjson.Field
		CloudID                 respjson.Field
		Email                   respjson.Field
		Expand                  respjson.Field
		Fields                  respjson.Field
		GetPermissions          respjson.Field
		RequestsPerMinute       respjson.Field
		SupportsAccessControl   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudJiraDataSourceV2) RawJSON() string { return r.JSON.raw }
func (r *CloudJiraDataSourceV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudJiraDataSourceV2 to a CloudJiraDataSourceV2Param.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudJiraDataSourceV2Param.Overrides()
func (r CloudJiraDataSourceV2) ToParam() CloudJiraDataSourceV2Param {
	return param.Override[CloudJiraDataSourceV2Param](json.RawMessage(r.RawJSON()))
}

// Jira REST API version to use (2 or 3). 3 supports Atlassian Document Format
// (ADF).
type CloudJiraDataSourceV2APIVersion string

const (
	CloudJiraDataSourceV2APIVersion2 CloudJiraDataSourceV2APIVersion = "2"
	CloudJiraDataSourceV2APIVersion3 CloudJiraDataSourceV2APIVersion = "3"
)

// Cloud Jira Data Source integrating JiraReaderV2.
//
// The properties AuthenticationMechanism, Query, ServerURL are required.
type CloudJiraDataSourceV2Param struct {
	// Type of Authentication for connecting to Jira APIs.
	AuthenticationMechanism string `json:"authentication_mechanism" api:"required"`
	// JQL (Jira Query Language) query to search.
	Query string `json:"query" api:"required"`
	// The server url for Jira Cloud.
	ServerURL string `json:"server_url" api:"required"`
	// The API Access Token used for Basic, PAT and OAuth2 authentication.
	APIToken param.Opt[string] `json:"api_token,omitzero" format:"password"`
	// The cloud ID, used in case of OAuth2.
	CloudID param.Opt[string] `json:"cloud_id,omitzero"`
	// The email address to use for authentication.
	Email param.Opt[string] `json:"email,omitzero"`
	// Fields to expand in the response.
	Expand param.Opt[string] `json:"expand,omitzero"`
	// Rate limit for Jira API requests per minute.
	RequestsPerMinute param.Opt[int64]  `json:"requests_per_minute,omitzero"`
	ClassName         param.Opt[string] `json:"class_name,omitzero"`
	// Whether to fetch project role permissions and issue-level security
	GetPermissions        param.Opt[bool] `json:"get_permissions,omitzero"`
	SupportsAccessControl param.Opt[bool] `json:"supports_access_control,omitzero"`
	// List of fields to retrieve from Jira. If None, retrieves all fields.
	Fields []string `json:"fields,omitzero"`
	// Jira REST API version to use (2 or 3). 3 supports Atlassian Document Format
	// (ADF).
	//
	// Any of "2", "3".
	APIVersion CloudJiraDataSourceV2APIVersion `json:"api_version,omitzero"`
	paramObj
}

func (r CloudJiraDataSourceV2Param) MarshalJSON() (data []byte, err error) {
	type shadow CloudJiraDataSourceV2Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudJiraDataSourceV2Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cloud Milvus Vector Store.
type CloudMilvusVectorStore struct {
	Uri                           string `json:"uri" api:"required"`
	Token                         string `json:"token" api:"nullable" format:"password"`
	ClassName                     string `json:"class_name"`
	CollectionName                string `json:"collection_name" api:"nullable"`
	EmbeddingDimension            int64  `json:"embedding_dimension" api:"nullable"`
	SupportsNestedMetadataFilters bool   `json:"supports_nested_metadata_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri                           respjson.Field
		Token                         respjson.Field
		ClassName                     respjson.Field
		CollectionName                respjson.Field
		EmbeddingDimension            respjson.Field
		SupportsNestedMetadataFilters respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudMilvusVectorStore) RawJSON() string { return r.JSON.raw }
func (r *CloudMilvusVectorStore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudMilvusVectorStore to a CloudMilvusVectorStoreParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudMilvusVectorStoreParam.Overrides()
func (r CloudMilvusVectorStore) ToParam() CloudMilvusVectorStoreParam {
	return param.Override[CloudMilvusVectorStoreParam](json.RawMessage(r.RawJSON()))
}

// Cloud Milvus Vector Store.
//
// The property Uri is required.
type CloudMilvusVectorStoreParam struct {
	Uri                           string            `json:"uri" api:"required"`
	Token                         param.Opt[string] `json:"token,omitzero" format:"password"`
	CollectionName                param.Opt[string] `json:"collection_name,omitzero"`
	EmbeddingDimension            param.Opt[int64]  `json:"embedding_dimension,omitzero"`
	ClassName                     param.Opt[string] `json:"class_name,omitzero"`
	SupportsNestedMetadataFilters param.Opt[bool]   `json:"supports_nested_metadata_filters,omitzero"`
	paramObj
}

func (r CloudMilvusVectorStoreParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudMilvusVectorStoreParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudMilvusVectorStoreParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cloud MongoDB Atlas Vector Store.
//
// This class is used to store the configuration for a MongoDB Atlas vector store,
// so that it can be created and used in LlamaCloud.
//
// Args: mongodb_uri (str): URI for connecting to MongoDB Atlas db_name (str): name
// of the MongoDB database collection_name (str): name of the MongoDB collection
// vector_index_name (str): name of the MongoDB Atlas vector index
// fulltext_index_name (str): name of the MongoDB Atlas full-text index
type CloudMongoDBAtlasVectorSearch struct {
	CollectionName                string `json:"collection_name" api:"required"`
	DBName                        string `json:"db_name" api:"required"`
	MongoDBUri                    string `json:"mongodb_uri" api:"required" format:"password"`
	ClassName                     string `json:"class_name"`
	EmbeddingDimension            int64  `json:"embedding_dimension" api:"nullable"`
	FulltextIndexName             string `json:"fulltext_index_name" api:"nullable"`
	SupportsNestedMetadataFilters bool   `json:"supports_nested_metadata_filters"`
	VectorIndexName               string `json:"vector_index_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionName                respjson.Field
		DBName                        respjson.Field
		MongoDBUri                    respjson.Field
		ClassName                     respjson.Field
		EmbeddingDimension            respjson.Field
		FulltextIndexName             respjson.Field
		SupportsNestedMetadataFilters respjson.Field
		VectorIndexName               respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudMongoDBAtlasVectorSearch) RawJSON() string { return r.JSON.raw }
func (r *CloudMongoDBAtlasVectorSearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudMongoDBAtlasVectorSearch to a
// CloudMongoDBAtlasVectorSearchParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudMongoDBAtlasVectorSearchParam.Overrides()
func (r CloudMongoDBAtlasVectorSearch) ToParam() CloudMongoDBAtlasVectorSearchParam {
	return param.Override[CloudMongoDBAtlasVectorSearchParam](json.RawMessage(r.RawJSON()))
}

// Cloud MongoDB Atlas Vector Store.
//
// This class is used to store the configuration for a MongoDB Atlas vector store,
// so that it can be created and used in LlamaCloud.
//
// Args: mongodb_uri (str): URI for connecting to MongoDB Atlas db_name (str): name
// of the MongoDB database collection_name (str): name of the MongoDB collection
// vector_index_name (str): name of the MongoDB Atlas vector index
// fulltext_index_name (str): name of the MongoDB Atlas full-text index
//
// The properties CollectionName, DBName, MongoDBUri are required.
type CloudMongoDBAtlasVectorSearchParam struct {
	CollectionName                string            `json:"collection_name" api:"required"`
	DBName                        string            `json:"db_name" api:"required"`
	MongoDBUri                    string            `json:"mongodb_uri" api:"required" format:"password"`
	EmbeddingDimension            param.Opt[int64]  `json:"embedding_dimension,omitzero"`
	FulltextIndexName             param.Opt[string] `json:"fulltext_index_name,omitzero"`
	VectorIndexName               param.Opt[string] `json:"vector_index_name,omitzero"`
	ClassName                     param.Opt[string] `json:"class_name,omitzero"`
	SupportsNestedMetadataFilters param.Opt[bool]   `json:"supports_nested_metadata_filters,omitzero"`
	paramObj
}

func (r CloudMongoDBAtlasVectorSearchParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudMongoDBAtlasVectorSearchParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudMongoDBAtlasVectorSearchParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CloudNotionPageDataSource struct {
	// The integration token to use for authentication.
	IntegrationToken string `json:"integration_token" api:"required" format:"password"`
	ClassName        string `json:"class_name"`
	// The Notion Database Id to read content from.
	DatabaseIDs string `json:"database_ids" api:"nullable"`
	// The Page ID's of the Notion to read from.
	PageIDs               string `json:"page_ids" api:"nullable"`
	SupportsAccessControl bool   `json:"supports_access_control"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IntegrationToken      respjson.Field
		ClassName             respjson.Field
		DatabaseIDs           respjson.Field
		PageIDs               respjson.Field
		SupportsAccessControl respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudNotionPageDataSource) RawJSON() string { return r.JSON.raw }
func (r *CloudNotionPageDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudNotionPageDataSource to a
// CloudNotionPageDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudNotionPageDataSourceParam.Overrides()
func (r CloudNotionPageDataSource) ToParam() CloudNotionPageDataSourceParam {
	return param.Override[CloudNotionPageDataSourceParam](json.RawMessage(r.RawJSON()))
}

// The property IntegrationToken is required.
type CloudNotionPageDataSourceParam struct {
	// The integration token to use for authentication.
	IntegrationToken string `json:"integration_token" api:"required" format:"password"`
	// The Notion Database Id to read content from.
	DatabaseIDs param.Opt[string] `json:"database_ids,omitzero"`
	// The Page ID's of the Notion to read from.
	PageIDs               param.Opt[string] `json:"page_ids,omitzero"`
	ClassName             param.Opt[string] `json:"class_name,omitzero"`
	SupportsAccessControl param.Opt[bool]   `json:"supports_access_control,omitzero"`
	paramObj
}

func (r CloudNotionPageDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudNotionPageDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudNotionPageDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CloudOneDriveDataSource struct {
	// The client ID to use for authentication.
	ClientID string `json:"client_id" api:"required"`
	// The client secret to use for authentication.
	ClientSecret string `json:"client_secret" api:"required" format:"password"`
	// The tenant ID to use for authentication.
	TenantID string `json:"tenant_id" api:"required"`
	// The user principal name to use for authentication.
	UserPrincipalName string `json:"user_principal_name" api:"required"`
	ClassName         string `json:"class_name"`
	// The ID of the OneDrive folder to read from.
	FolderID string `json:"folder_id" api:"nullable"`
	// The path of the OneDrive folder to read from.
	FolderPath string `json:"folder_path" api:"nullable"`
	// The list of required file extensions.
	RequiredExts []string `json:"required_exts" api:"nullable"`
	// Any of true.
	SupportsAccessControl bool `json:"supports_access_control"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID              respjson.Field
		ClientSecret          respjson.Field
		TenantID              respjson.Field
		UserPrincipalName     respjson.Field
		ClassName             respjson.Field
		FolderID              respjson.Field
		FolderPath            respjson.Field
		RequiredExts          respjson.Field
		SupportsAccessControl respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudOneDriveDataSource) RawJSON() string { return r.JSON.raw }
func (r *CloudOneDriveDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudOneDriveDataSource to a CloudOneDriveDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudOneDriveDataSourceParam.Overrides()
func (r CloudOneDriveDataSource) ToParam() CloudOneDriveDataSourceParam {
	return param.Override[CloudOneDriveDataSourceParam](json.RawMessage(r.RawJSON()))
}

// The properties ClientID, ClientSecret, TenantID, UserPrincipalName are required.
type CloudOneDriveDataSourceParam struct {
	// The client ID to use for authentication.
	ClientID string `json:"client_id" api:"required"`
	// The client secret to use for authentication.
	ClientSecret string `json:"client_secret" api:"required" format:"password"`
	// The tenant ID to use for authentication.
	TenantID string `json:"tenant_id" api:"required"`
	// The user principal name to use for authentication.
	UserPrincipalName string `json:"user_principal_name" api:"required"`
	// The ID of the OneDrive folder to read from.
	FolderID param.Opt[string] `json:"folder_id,omitzero"`
	// The path of the OneDrive folder to read from.
	FolderPath param.Opt[string] `json:"folder_path,omitzero"`
	ClassName  param.Opt[string] `json:"class_name,omitzero"`
	// The list of required file extensions.
	RequiredExts []string `json:"required_exts,omitzero"`
	// Any of true.
	SupportsAccessControl bool `json:"supports_access_control,omitzero"`
	paramObj
}

func (r CloudOneDriveDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudOneDriveDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudOneDriveDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CloudOneDriveDataSourceParam](
		"supports_access_control", true,
	)
}

// Cloud Pinecone Vector Store.
//
// This class is used to store the configuration for a Pinecone vector store, so
// that it can be created and used in LlamaCloud.
//
// Args: api_key (str): API key for authenticating with Pinecone index_name (str):
// name of the Pinecone index namespace (optional[str]): namespace to use in the
// Pinecone index insert_kwargs (optional[dict]): additional kwargs to pass during
// insertion
type CloudPineconeVectorStore struct {
	// The API key for authenticating with Pinecone
	APIKey       string         `json:"api_key" api:"required" format:"password"`
	IndexName    string         `json:"index_name" api:"required"`
	ClassName    string         `json:"class_name"`
	InsertKwargs map[string]any `json:"insert_kwargs" api:"nullable"`
	Namespace    string         `json:"namespace" api:"nullable"`
	// Any of true.
	SupportsNestedMetadataFilters bool `json:"supports_nested_metadata_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey                        respjson.Field
		IndexName                     respjson.Field
		ClassName                     respjson.Field
		InsertKwargs                  respjson.Field
		Namespace                     respjson.Field
		SupportsNestedMetadataFilters respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudPineconeVectorStore) RawJSON() string { return r.JSON.raw }
func (r *CloudPineconeVectorStore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudPineconeVectorStore to a
// CloudPineconeVectorStoreParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudPineconeVectorStoreParam.Overrides()
func (r CloudPineconeVectorStore) ToParam() CloudPineconeVectorStoreParam {
	return param.Override[CloudPineconeVectorStoreParam](json.RawMessage(r.RawJSON()))
}

// Cloud Pinecone Vector Store.
//
// This class is used to store the configuration for a Pinecone vector store, so
// that it can be created and used in LlamaCloud.
//
// Args: api_key (str): API key for authenticating with Pinecone index_name (str):
// name of the Pinecone index namespace (optional[str]): namespace to use in the
// Pinecone index insert_kwargs (optional[dict]): additional kwargs to pass during
// insertion
//
// The properties APIKey, IndexName are required.
type CloudPineconeVectorStoreParam struct {
	// The API key for authenticating with Pinecone
	APIKey       string            `json:"api_key" api:"required" format:"password"`
	IndexName    string            `json:"index_name" api:"required"`
	Namespace    param.Opt[string] `json:"namespace,omitzero"`
	ClassName    param.Opt[string] `json:"class_name,omitzero"`
	InsertKwargs map[string]any    `json:"insert_kwargs,omitzero"`
	// Any of true.
	SupportsNestedMetadataFilters bool `json:"supports_nested_metadata_filters,omitzero"`
	paramObj
}

func (r CloudPineconeVectorStoreParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudPineconeVectorStoreParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudPineconeVectorStoreParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CloudPineconeVectorStoreParam](
		"supports_nested_metadata_filters", true,
	)
}

type CloudPostgresVectorStore struct {
	Database   string `json:"database" api:"required"`
	EmbedDim   int64  `json:"embed_dim" api:"required"`
	Host       string `json:"host" api:"required"`
	Password   string `json:"password" api:"required" format:"password"`
	Port       int64  `json:"port" api:"required"`
	SchemaName string `json:"schema_name" api:"required"`
	TableName  string `json:"table_name" api:"required"`
	User       string `json:"user" api:"required"`
	ClassName  string `json:"class_name"`
	// HNSW settings for PGVector.
	HnswSettings                  PgVectorHnswSettings `json:"hnsw_settings" api:"nullable"`
	HybridSearch                  bool                 `json:"hybrid_search" api:"nullable"`
	PerformSetup                  bool                 `json:"perform_setup"`
	SupportsNestedMetadataFilters bool                 `json:"supports_nested_metadata_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Database                      respjson.Field
		EmbedDim                      respjson.Field
		Host                          respjson.Field
		Password                      respjson.Field
		Port                          respjson.Field
		SchemaName                    respjson.Field
		TableName                     respjson.Field
		User                          respjson.Field
		ClassName                     respjson.Field
		HnswSettings                  respjson.Field
		HybridSearch                  respjson.Field
		PerformSetup                  respjson.Field
		SupportsNestedMetadataFilters respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudPostgresVectorStore) RawJSON() string { return r.JSON.raw }
func (r *CloudPostgresVectorStore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudPostgresVectorStore to a
// CloudPostgresVectorStoreParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudPostgresVectorStoreParam.Overrides()
func (r CloudPostgresVectorStore) ToParam() CloudPostgresVectorStoreParam {
	return param.Override[CloudPostgresVectorStoreParam](json.RawMessage(r.RawJSON()))
}

// The properties Database, EmbedDim, Host, Password, Port, SchemaName, TableName,
// User are required.
type CloudPostgresVectorStoreParam struct {
	Database                      string            `json:"database" api:"required"`
	EmbedDim                      int64             `json:"embed_dim" api:"required"`
	Host                          string            `json:"host" api:"required"`
	Password                      string            `json:"password" api:"required" format:"password"`
	Port                          int64             `json:"port" api:"required"`
	SchemaName                    string            `json:"schema_name" api:"required"`
	TableName                     string            `json:"table_name" api:"required"`
	User                          string            `json:"user" api:"required"`
	HybridSearch                  param.Opt[bool]   `json:"hybrid_search,omitzero"`
	ClassName                     param.Opt[string] `json:"class_name,omitzero"`
	PerformSetup                  param.Opt[bool]   `json:"perform_setup,omitzero"`
	SupportsNestedMetadataFilters param.Opt[bool]   `json:"supports_nested_metadata_filters,omitzero"`
	// HNSW settings for PGVector.
	HnswSettings PgVectorHnswSettingsParam `json:"hnsw_settings,omitzero"`
	paramObj
}

func (r CloudPostgresVectorStoreParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudPostgresVectorStoreParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudPostgresVectorStoreParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cloud Qdrant Vector Store.
//
// This class is used to store the configuration for a Qdrant vector store, so that
// it can be created and used in LlamaCloud.
//
// Args: collection_name (str): name of the Qdrant collection url (str): url of the
// Qdrant instance api_key (str): API key for authenticating with Qdrant
// max_retries (int): maximum number of retries in case of a failure. Defaults to 3
// client_kwargs (dict): additional kwargs to pass to the Qdrant client
type CloudQdrantVectorStore struct {
	APIKey         string         `json:"api_key" api:"required" format:"password"`
	CollectionName string         `json:"collection_name" api:"required"`
	URL            string         `json:"url" api:"required"`
	ClassName      string         `json:"class_name"`
	ClientKwargs   map[string]any `json:"client_kwargs"`
	MaxRetries     int64          `json:"max_retries"`
	// Any of true.
	SupportsNestedMetadataFilters bool `json:"supports_nested_metadata_filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKey                        respjson.Field
		CollectionName                respjson.Field
		URL                           respjson.Field
		ClassName                     respjson.Field
		ClientKwargs                  respjson.Field
		MaxRetries                    respjson.Field
		SupportsNestedMetadataFilters respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudQdrantVectorStore) RawJSON() string { return r.JSON.raw }
func (r *CloudQdrantVectorStore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudQdrantVectorStore to a CloudQdrantVectorStoreParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudQdrantVectorStoreParam.Overrides()
func (r CloudQdrantVectorStore) ToParam() CloudQdrantVectorStoreParam {
	return param.Override[CloudQdrantVectorStoreParam](json.RawMessage(r.RawJSON()))
}

// Cloud Qdrant Vector Store.
//
// This class is used to store the configuration for a Qdrant vector store, so that
// it can be created and used in LlamaCloud.
//
// Args: collection_name (str): name of the Qdrant collection url (str): url of the
// Qdrant instance api_key (str): API key for authenticating with Qdrant
// max_retries (int): maximum number of retries in case of a failure. Defaults to 3
// client_kwargs (dict): additional kwargs to pass to the Qdrant client
//
// The properties APIKey, CollectionName, URL are required.
type CloudQdrantVectorStoreParam struct {
	APIKey         string            `json:"api_key" api:"required" format:"password"`
	CollectionName string            `json:"collection_name" api:"required"`
	URL            string            `json:"url" api:"required"`
	ClassName      param.Opt[string] `json:"class_name,omitzero"`
	MaxRetries     param.Opt[int64]  `json:"max_retries,omitzero"`
	ClientKwargs   map[string]any    `json:"client_kwargs,omitzero"`
	// Any of true.
	SupportsNestedMetadataFilters bool `json:"supports_nested_metadata_filters,omitzero"`
	paramObj
}

func (r CloudQdrantVectorStoreParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudQdrantVectorStoreParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudQdrantVectorStoreParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CloudQdrantVectorStoreParam](
		"supports_nested_metadata_filters", true,
	)
}

type CloudS3DataSource struct {
	// The name of the S3 bucket to read from.
	Bucket string `json:"bucket" api:"required"`
	// The AWS access ID to use for authentication.
	AwsAccessID string `json:"aws_access_id" api:"nullable"`
	// The AWS access secret to use for authentication.
	AwsAccessSecret string `json:"aws_access_secret" api:"nullable" format:"password"`
	ClassName       string `json:"class_name"`
	// The prefix of the S3 objects to read from.
	Prefix string `json:"prefix" api:"nullable"`
	// The regex pattern to filter S3 objects. Must be a valid regex pattern.
	RegexPattern string `json:"regex_pattern" api:"nullable"`
	// The S3 endpoint URL to use for authentication.
	S3EndpointURL         string `json:"s3_endpoint_url" api:"nullable"`
	SupportsAccessControl bool   `json:"supports_access_control"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket                respjson.Field
		AwsAccessID           respjson.Field
		AwsAccessSecret       respjson.Field
		ClassName             respjson.Field
		Prefix                respjson.Field
		RegexPattern          respjson.Field
		S3EndpointURL         respjson.Field
		SupportsAccessControl respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudS3DataSource) RawJSON() string { return r.JSON.raw }
func (r *CloudS3DataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudS3DataSource to a CloudS3DataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudS3DataSourceParam.Overrides()
func (r CloudS3DataSource) ToParam() CloudS3DataSourceParam {
	return param.Override[CloudS3DataSourceParam](json.RawMessage(r.RawJSON()))
}

// The property Bucket is required.
type CloudS3DataSourceParam struct {
	// The name of the S3 bucket to read from.
	Bucket string `json:"bucket" api:"required"`
	// The AWS access ID to use for authentication.
	AwsAccessID param.Opt[string] `json:"aws_access_id,omitzero"`
	// The AWS access secret to use for authentication.
	AwsAccessSecret param.Opt[string] `json:"aws_access_secret,omitzero" format:"password"`
	// The prefix of the S3 objects to read from.
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// The regex pattern to filter S3 objects. Must be a valid regex pattern.
	RegexPattern param.Opt[string] `json:"regex_pattern,omitzero"`
	// The S3 endpoint URL to use for authentication.
	S3EndpointURL         param.Opt[string] `json:"s3_endpoint_url,omitzero"`
	ClassName             param.Opt[string] `json:"class_name,omitzero"`
	SupportsAccessControl param.Opt[bool]   `json:"supports_access_control,omitzero"`
	paramObj
}

func (r CloudS3DataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudS3DataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudS3DataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CloudSharepointDataSource struct {
	// The client ID to use for authentication.
	ClientID string `json:"client_id" api:"required"`
	// The client secret to use for authentication.
	ClientSecret string `json:"client_secret" api:"required" format:"password"`
	// The tenant ID to use for authentication.
	TenantID  string `json:"tenant_id" api:"required"`
	ClassName string `json:"class_name"`
	// The name of the Sharepoint drive to read from.
	DriveName string `json:"drive_name" api:"nullable"`
	// List of regex patterns for file paths to exclude. Files whose paths (including
	// filename) match any pattern will be excluded. Example: ['/temp/', '/backup/',
	// '\.git/', '\.tmp$', '^~']
	ExcludePathPatterns []string `json:"exclude_path_patterns" api:"nullable"`
	// The ID of the Sharepoint folder to read from.
	FolderID string `json:"folder_id" api:"nullable"`
	// The path of the Sharepoint folder to read from.
	FolderPath string `json:"folder_path" api:"nullable"`
	// Whether to get permissions for the sharepoint site.
	GetPermissions bool `json:"get_permissions"`
	// List of regex patterns for file paths to include. Full paths (including
	// filename) must match at least one pattern to be included. Example: ['/reports/',
	// '/docs/.*\.pdf$', '^Report.*\.pdf$']
	IncludePathPatterns []string `json:"include_path_patterns" api:"nullable"`
	// The list of required file extensions.
	RequiredExts []string `json:"required_exts" api:"nullable"`
	// The ID of the SharePoint site to download from.
	SiteID string `json:"site_id" api:"nullable"`
	// The name of the SharePoint site to download from.
	SiteName string `json:"site_name" api:"nullable"`
	// Any of true.
	SupportsAccessControl bool `json:"supports_access_control"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID              respjson.Field
		ClientSecret          respjson.Field
		TenantID              respjson.Field
		ClassName             respjson.Field
		DriveName             respjson.Field
		ExcludePathPatterns   respjson.Field
		FolderID              respjson.Field
		FolderPath            respjson.Field
		GetPermissions        respjson.Field
		IncludePathPatterns   respjson.Field
		RequiredExts          respjson.Field
		SiteID                respjson.Field
		SiteName              respjson.Field
		SupportsAccessControl respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudSharepointDataSource) RawJSON() string { return r.JSON.raw }
func (r *CloudSharepointDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudSharepointDataSource to a
// CloudSharepointDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudSharepointDataSourceParam.Overrides()
func (r CloudSharepointDataSource) ToParam() CloudSharepointDataSourceParam {
	return param.Override[CloudSharepointDataSourceParam](json.RawMessage(r.RawJSON()))
}

// The properties ClientID, ClientSecret, TenantID are required.
type CloudSharepointDataSourceParam struct {
	// The client ID to use for authentication.
	ClientID string `json:"client_id" api:"required"`
	// The client secret to use for authentication.
	ClientSecret string `json:"client_secret" api:"required" format:"password"`
	// The tenant ID to use for authentication.
	TenantID string `json:"tenant_id" api:"required"`
	// The name of the Sharepoint drive to read from.
	DriveName param.Opt[string] `json:"drive_name,omitzero"`
	// The ID of the Sharepoint folder to read from.
	FolderID param.Opt[string] `json:"folder_id,omitzero"`
	// The path of the Sharepoint folder to read from.
	FolderPath param.Opt[string] `json:"folder_path,omitzero"`
	// The ID of the SharePoint site to download from.
	SiteID param.Opt[string] `json:"site_id,omitzero"`
	// The name of the SharePoint site to download from.
	SiteName  param.Opt[string] `json:"site_name,omitzero"`
	ClassName param.Opt[string] `json:"class_name,omitzero"`
	// Whether to get permissions for the sharepoint site.
	GetPermissions param.Opt[bool] `json:"get_permissions,omitzero"`
	// List of regex patterns for file paths to exclude. Files whose paths (including
	// filename) match any pattern will be excluded. Example: ['/temp/', '/backup/',
	// '\.git/', '\.tmp$', '^~']
	ExcludePathPatterns []string `json:"exclude_path_patterns,omitzero"`
	// List of regex patterns for file paths to include. Full paths (including
	// filename) must match at least one pattern to be included. Example: ['/reports/',
	// '/docs/.*\.pdf$', '^Report.*\.pdf$']
	IncludePathPatterns []string `json:"include_path_patterns,omitzero"`
	// The list of required file extensions.
	RequiredExts []string `json:"required_exts,omitzero"`
	// Any of true.
	SupportsAccessControl bool `json:"supports_access_control,omitzero"`
	paramObj
}

func (r CloudSharepointDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudSharepointDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudSharepointDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CloudSharepointDataSourceParam](
		"supports_access_control", true,
	)
}

type CloudSlackDataSource struct {
	// Slack Bot Token.
	SlackToken string `json:"slack_token" api:"required" format:"password"`
	// Slack Channel.
	ChannelIDs string `json:"channel_ids" api:"nullable"`
	// Slack Channel name pattern.
	ChannelPatterns string `json:"channel_patterns" api:"nullable"`
	ClassName       string `json:"class_name"`
	// Earliest date.
	EarliestDate string `json:"earliest_date" api:"nullable"`
	// Earliest date timestamp.
	EarliestDateTimestamp float64 `json:"earliest_date_timestamp" api:"nullable"`
	// Latest date.
	LatestDate string `json:"latest_date" api:"nullable"`
	// Latest date timestamp.
	LatestDateTimestamp   float64 `json:"latest_date_timestamp" api:"nullable"`
	SupportsAccessControl bool    `json:"supports_access_control"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SlackToken            respjson.Field
		ChannelIDs            respjson.Field
		ChannelPatterns       respjson.Field
		ClassName             respjson.Field
		EarliestDate          respjson.Field
		EarliestDateTimestamp respjson.Field
		LatestDate            respjson.Field
		LatestDateTimestamp   respjson.Field
		SupportsAccessControl respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloudSlackDataSource) RawJSON() string { return r.JSON.raw }
func (r *CloudSlackDataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CloudSlackDataSource to a CloudSlackDataSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CloudSlackDataSourceParam.Overrides()
func (r CloudSlackDataSource) ToParam() CloudSlackDataSourceParam {
	return param.Override[CloudSlackDataSourceParam](json.RawMessage(r.RawJSON()))
}

// The property SlackToken is required.
type CloudSlackDataSourceParam struct {
	// Slack Bot Token.
	SlackToken string `json:"slack_token" api:"required" format:"password"`
	// Slack Channel.
	ChannelIDs param.Opt[string] `json:"channel_ids,omitzero"`
	// Slack Channel name pattern.
	ChannelPatterns param.Opt[string] `json:"channel_patterns,omitzero"`
	// Earliest date.
	EarliestDate param.Opt[string] `json:"earliest_date,omitzero"`
	// Earliest date timestamp.
	EarliestDateTimestamp param.Opt[float64] `json:"earliest_date_timestamp,omitzero"`
	// Latest date.
	LatestDate param.Opt[string] `json:"latest_date,omitzero"`
	// Latest date timestamp.
	LatestDateTimestamp   param.Opt[float64] `json:"latest_date_timestamp,omitzero"`
	ClassName             param.Opt[string]  `json:"class_name,omitzero"`
	SupportsAccessControl param.Opt[bool]    `json:"supports_access_control,omitzero"`
	paramObj
}

func (r CloudSlackDataSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow CloudSlackDataSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CloudSlackDataSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for handling different types of failures during data source
// processing.
type FailureHandlingConfig struct {
	// Whether to skip failed batches/lists and continue processing
	SkipListFailures bool `json:"skip_list_failures"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SkipListFailures respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FailureHandlingConfig) RawJSON() string { return r.JSON.raw }
func (r *FailureHandlingConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FailureHandlingConfig to a FailureHandlingConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FailureHandlingConfigParam.Overrides()
func (r FailureHandlingConfig) ToParam() FailureHandlingConfigParam {
	return param.Override[FailureHandlingConfigParam](json.RawMessage(r.RawJSON()))
}

// Configuration for handling different types of failures during data source
// processing.
type FailureHandlingConfigParam struct {
	// Whether to skip failed batches/lists and continue processing
	SkipListFailures param.Opt[bool] `json:"skip_list_failures,omitzero"`
	paramObj
}

func (r FailureHandlingConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow FailureHandlingConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FailureHandlingConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HNSW settings for PGVector.
type PgVectorHnswSettings struct {
	// The distance method to use.
	//
	// Any of "cosine", "hamming", "ip", "jaccard", "l1", "l2".
	DistanceMethod PgVectorHnswSettingsDistanceMethod `json:"distance_method"`
	// The number of edges to use during the construction phase.
	EfConstruction int64 `json:"ef_construction"`
	// The number of edges to use during the search phase.
	EfSearch int64 `json:"ef_search"`
	// The number of bi-directional links created for each new element.
	M int64 `json:"m"`
	// The type of vector to use.
	//
	// Any of "bit", "half_vec", "sparse_vec", "vector".
	VectorType PgVectorHnswSettingsVectorType `json:"vector_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DistanceMethod respjson.Field
		EfConstruction respjson.Field
		EfSearch       respjson.Field
		M              respjson.Field
		VectorType     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PgVectorHnswSettings) RawJSON() string { return r.JSON.raw }
func (r *PgVectorHnswSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PgVectorHnswSettings to a PgVectorHnswSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PgVectorHnswSettingsParam.Overrides()
func (r PgVectorHnswSettings) ToParam() PgVectorHnswSettingsParam {
	return param.Override[PgVectorHnswSettingsParam](json.RawMessage(r.RawJSON()))
}

// The distance method to use.
type PgVectorHnswSettingsDistanceMethod string

const (
	PgVectorHnswSettingsDistanceMethodCosine  PgVectorHnswSettingsDistanceMethod = "cosine"
	PgVectorHnswSettingsDistanceMethodHamming PgVectorHnswSettingsDistanceMethod = "hamming"
	PgVectorHnswSettingsDistanceMethodIP      PgVectorHnswSettingsDistanceMethod = "ip"
	PgVectorHnswSettingsDistanceMethodJaccard PgVectorHnswSettingsDistanceMethod = "jaccard"
	PgVectorHnswSettingsDistanceMethodL1      PgVectorHnswSettingsDistanceMethod = "l1"
	PgVectorHnswSettingsDistanceMethodL2      PgVectorHnswSettingsDistanceMethod = "l2"
)

// The type of vector to use.
type PgVectorHnswSettingsVectorType string

const (
	PgVectorHnswSettingsVectorTypeBit       PgVectorHnswSettingsVectorType = "bit"
	PgVectorHnswSettingsVectorTypeHalfVec   PgVectorHnswSettingsVectorType = "half_vec"
	PgVectorHnswSettingsVectorTypeSparseVec PgVectorHnswSettingsVectorType = "sparse_vec"
	PgVectorHnswSettingsVectorTypeVector    PgVectorHnswSettingsVectorType = "vector"
)

// HNSW settings for PGVector.
type PgVectorHnswSettingsParam struct {
	// The number of edges to use during the construction phase.
	EfConstruction param.Opt[int64] `json:"ef_construction,omitzero"`
	// The number of edges to use during the search phase.
	EfSearch param.Opt[int64] `json:"ef_search,omitzero"`
	// The number of bi-directional links created for each new element.
	M param.Opt[int64] `json:"m,omitzero"`
	// The distance method to use.
	//
	// Any of "cosine", "hamming", "ip", "jaccard", "l1", "l2".
	DistanceMethod PgVectorHnswSettingsDistanceMethod `json:"distance_method,omitzero"`
	// The type of vector to use.
	//
	// Any of "bit", "half_vec", "sparse_vec", "vector".
	VectorType PgVectorHnswSettingsVectorType `json:"vector_type,omitzero"`
	paramObj
}

func (r PgVectorHnswSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow PgVectorHnswSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PgVectorHnswSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
