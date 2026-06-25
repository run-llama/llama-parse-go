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
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/packages/respjson"
	"github.com/run-llama/llama-parse-go/shared"
)

// DataSinkService contains methods and other services that help with interacting
// with the llama-cloud API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDataSinkService] method instead.
type DataSinkService struct {
	options []option.RequestOption
}

// NewDataSinkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDataSinkService(opts ...option.RequestOption) (r DataSinkService) {
	r = DataSinkService{}
	r.options = opts
	return
}

// Create a new data sink.
func (r *DataSinkService) New(ctx context.Context, params DataSinkNewParams, opts ...option.RequestOption) (res *DataSink, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/data-sinks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Update a data sink by ID.
func (r *DataSinkService) Update(ctx context.Context, dataSinkID string, body DataSinkUpdateParams, opts ...option.RequestOption) (res *DataSink, err error) {
	opts = slices.Concat(r.options, opts)
	if dataSinkID == "" {
		err = errors.New("missing required data_sink_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/data-sinks/%s", dataSinkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// List data sinks for a given project.
func (r *DataSinkService) List(ctx context.Context, query DataSinkListParams, opts ...option.RequestOption) (res *[]DataSink, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/data-sinks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a data sink by ID.
func (r *DataSinkService) Delete(ctx context.Context, dataSinkID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if dataSinkID == "" {
		err = errors.New("missing required data_sink_id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/data-sinks/%s", dataSinkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get a data sink by ID.
func (r *DataSinkService) Get(ctx context.Context, dataSinkID string, opts ...option.RequestOption) (res *DataSink, err error) {
	opts = slices.Concat(r.options, opts)
	if dataSinkID == "" {
		err = errors.New("missing required data_sink_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/data-sinks/%s", dataSinkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Schema for a data sink.
type DataSink struct {
	// Unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// Component that implements the data sink
	Component DataSinkComponentUnion `json:"component" api:"required"`
	// The name of the data sink.
	Name      string `json:"name" api:"required"`
	ProjectID string `json:"project_id" api:"required" format:"uuid"`
	// Any of "ASTRA_DB", "AZUREAI_SEARCH", "MILVUS", "MONGODB_ATLAS", "PINECONE",
	// "POSTGRES", "QDRANT".
	SinkType DataSinkSinkType `json:"sink_type" api:"required"`
	// Creation datetime
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Update datetime
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Component   respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		SinkType    respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DataSink) RawJSON() string { return r.JSON.raw }
func (r *DataSink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DataSinkComponentUnion contains all possible properties and values from
// [map[string]any], [shared.CloudPineconeVectorStore],
// [shared.CloudPostgresVectorStore], [shared.CloudQdrantVectorStore],
// [shared.CloudAzureAISearchVectorStore], [shared.CloudMongoDBAtlasVectorSearch],
// [shared.CloudMilvusVectorStore], [shared.CloudAstraDBVectorStore].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfDataSinkComponentMapItem]
type DataSinkComponentUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfDataSinkComponentMapItem any    `json:",inline"`
	APIKey                     string `json:"api_key"`
	IndexName                  string `json:"index_name"`
	ClassName                  string `json:"class_name"`
	// This field is from variant [shared.CloudPineconeVectorStore].
	InsertKwargs map[string]any `json:"insert_kwargs"`
	// This field is from variant [shared.CloudPineconeVectorStore].
	Namespace                     string `json:"namespace"`
	SupportsNestedMetadataFilters bool   `json:"supports_nested_metadata_filters"`
	// This field is from variant [shared.CloudPostgresVectorStore].
	Database string `json:"database"`
	// This field is from variant [shared.CloudPostgresVectorStore].
	EmbedDim int64 `json:"embed_dim"`
	// This field is from variant [shared.CloudPostgresVectorStore].
	Host string `json:"host"`
	// This field is from variant [shared.CloudPostgresVectorStore].
	Password string `json:"password"`
	// This field is from variant [shared.CloudPostgresVectorStore].
	Port int64 `json:"port"`
	// This field is from variant [shared.CloudPostgresVectorStore].
	SchemaName string `json:"schema_name"`
	// This field is from variant [shared.CloudPostgresVectorStore].
	TableName string `json:"table_name"`
	// This field is from variant [shared.CloudPostgresVectorStore].
	User string `json:"user"`
	// This field is from variant [shared.CloudPostgresVectorStore].
	HnswSettings shared.PgVectorHnswSettings `json:"hnsw_settings"`
	// This field is from variant [shared.CloudPostgresVectorStore].
	HybridSearch bool `json:"hybrid_search"`
	// This field is from variant [shared.CloudPostgresVectorStore].
	PerformSetup   bool   `json:"perform_setup"`
	CollectionName string `json:"collection_name"`
	// This field is from variant [shared.CloudQdrantVectorStore].
	URL string `json:"url"`
	// This field is from variant [shared.CloudQdrantVectorStore].
	ClientKwargs map[string]any `json:"client_kwargs"`
	// This field is from variant [shared.CloudQdrantVectorStore].
	MaxRetries int64 `json:"max_retries"`
	// This field is from variant [shared.CloudAzureAISearchVectorStore].
	SearchServiceAPIKey string `json:"search_service_api_key"`
	// This field is from variant [shared.CloudAzureAISearchVectorStore].
	SearchServiceEndpoint string `json:"search_service_endpoint"`
	// This field is from variant [shared.CloudAzureAISearchVectorStore].
	ClientID string `json:"client_id"`
	// This field is from variant [shared.CloudAzureAISearchVectorStore].
	ClientSecret       string `json:"client_secret"`
	EmbeddingDimension int64  `json:"embedding_dimension"`
	// This field is from variant [shared.CloudAzureAISearchVectorStore].
	FilterableMetadataFieldKeys map[string]any `json:"filterable_metadata_field_keys"`
	// This field is from variant [shared.CloudAzureAISearchVectorStore].
	SearchServiceAPIVersion string `json:"search_service_api_version"`
	// This field is from variant [shared.CloudAzureAISearchVectorStore].
	TenantID string `json:"tenant_id"`
	// This field is from variant [shared.CloudMongoDBAtlasVectorSearch].
	DBName string `json:"db_name"`
	// This field is from variant [shared.CloudMongoDBAtlasVectorSearch].
	MongoDBUri string `json:"mongodb_uri"`
	// This field is from variant [shared.CloudMongoDBAtlasVectorSearch].
	FulltextIndexName string `json:"fulltext_index_name"`
	// This field is from variant [shared.CloudMongoDBAtlasVectorSearch].
	VectorIndexName string `json:"vector_index_name"`
	// This field is from variant [shared.CloudMilvusVectorStore].
	Uri   string `json:"uri"`
	Token string `json:"token"`
	// This field is from variant [shared.CloudAstraDBVectorStore].
	APIEndpoint string `json:"api_endpoint"`
	// This field is from variant [shared.CloudAstraDBVectorStore].
	Keyspace string `json:"keyspace"`
	JSON     struct {
		OfDataSinkComponentMapItem    respjson.Field
		APIKey                        respjson.Field
		IndexName                     respjson.Field
		ClassName                     respjson.Field
		InsertKwargs                  respjson.Field
		Namespace                     respjson.Field
		SupportsNestedMetadataFilters respjson.Field
		Database                      respjson.Field
		EmbedDim                      respjson.Field
		Host                          respjson.Field
		Password                      respjson.Field
		Port                          respjson.Field
		SchemaName                    respjson.Field
		TableName                     respjson.Field
		User                          respjson.Field
		HnswSettings                  respjson.Field
		HybridSearch                  respjson.Field
		PerformSetup                  respjson.Field
		CollectionName                respjson.Field
		URL                           respjson.Field
		ClientKwargs                  respjson.Field
		MaxRetries                    respjson.Field
		SearchServiceAPIKey           respjson.Field
		SearchServiceEndpoint         respjson.Field
		ClientID                      respjson.Field
		ClientSecret                  respjson.Field
		EmbeddingDimension            respjson.Field
		FilterableMetadataFieldKeys   respjson.Field
		SearchServiceAPIVersion       respjson.Field
		TenantID                      respjson.Field
		DBName                        respjson.Field
		MongoDBUri                    respjson.Field
		FulltextIndexName             respjson.Field
		VectorIndexName               respjson.Field
		Uri                           respjson.Field
		Token                         respjson.Field
		APIEndpoint                   respjson.Field
		Keyspace                      respjson.Field
		raw                           string
	} `json:"-"`
}

func (u DataSinkComponentUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSinkComponentUnion) AsCloudPineconeVectorStore() (v shared.CloudPineconeVectorStore) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSinkComponentUnion) AsCloudPostgresVectorStore() (v shared.CloudPostgresVectorStore) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSinkComponentUnion) AsCloudQdrantVectorStore() (v shared.CloudQdrantVectorStore) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSinkComponentUnion) AsCloudAzureAISearchVectorStore() (v shared.CloudAzureAISearchVectorStore) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSinkComponentUnion) AsCloudMongoDBAtlasVectorSearch() (v shared.CloudMongoDBAtlasVectorSearch) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSinkComponentUnion) AsCloudMilvusVectorStore() (v shared.CloudMilvusVectorStore) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DataSinkComponentUnion) AsCloudAstraDBVectorStore() (v shared.CloudAstraDBVectorStore) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DataSinkComponentUnion) RawJSON() string { return u.JSON.raw }

func (r *DataSinkComponentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataSinkSinkType string

const (
	DataSinkSinkTypeAstraDB       DataSinkSinkType = "ASTRA_DB"
	DataSinkSinkTypeAzureaiSearch DataSinkSinkType = "AZUREAI_SEARCH"
	DataSinkSinkTypeMilvus        DataSinkSinkType = "MILVUS"
	DataSinkSinkTypeMongoDBAtlas  DataSinkSinkType = "MONGODB_ATLAS"
	DataSinkSinkTypePinecone      DataSinkSinkType = "PINECONE"
	DataSinkSinkTypePostgres      DataSinkSinkType = "POSTGRES"
	DataSinkSinkTypeQdrant        DataSinkSinkType = "QDRANT"
)

type DataSinkNewParams struct {
	// Schema for creating a data sink.
	DataSinkCreate DataSinkCreateParam
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

func (r DataSinkNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DataSinkCreate)
}
func (r *DataSinkNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [DataSinkNewParams]'s query parameters as `url.Values`.
func (r DataSinkNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DataSinkUpdateParams struct {
	// Any of "ASTRA_DB", "AZUREAI_SEARCH", "MILVUS", "MONGODB_ATLAS", "PINECONE",
	// "POSTGRES", "QDRANT".
	SinkType DataSinkUpdateParamsSinkType `json:"sink_type,omitzero" api:"required"`
	// The name of the data sink.
	Name param.Opt[string] `json:"name,omitzero"`
	// Component that implements the data sink
	Component DataSinkUpdateParamsComponentUnion `json:"component,omitzero"`
	paramObj
}

func (r DataSinkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DataSinkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DataSinkUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DataSinkUpdateParamsSinkType string

const (
	DataSinkUpdateParamsSinkTypeAstraDB       DataSinkUpdateParamsSinkType = "ASTRA_DB"
	DataSinkUpdateParamsSinkTypeAzureaiSearch DataSinkUpdateParamsSinkType = "AZUREAI_SEARCH"
	DataSinkUpdateParamsSinkTypeMilvus        DataSinkUpdateParamsSinkType = "MILVUS"
	DataSinkUpdateParamsSinkTypeMongoDBAtlas  DataSinkUpdateParamsSinkType = "MONGODB_ATLAS"
	DataSinkUpdateParamsSinkTypePinecone      DataSinkUpdateParamsSinkType = "PINECONE"
	DataSinkUpdateParamsSinkTypePostgres      DataSinkUpdateParamsSinkType = "POSTGRES"
	DataSinkUpdateParamsSinkTypeQdrant        DataSinkUpdateParamsSinkType = "QDRANT"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DataSinkUpdateParamsComponentUnion struct {
	OfAnyMap                        map[string]any                             `json:",omitzero,inline"`
	OfCloudPineconeVectorStore      *shared.CloudPineconeVectorStoreParam      `json:",omitzero,inline"`
	OfCloudPostgresVectorStore      *shared.CloudPostgresVectorStoreParam      `json:",omitzero,inline"`
	OfCloudQdrantVectorStore        *shared.CloudQdrantVectorStoreParam        `json:",omitzero,inline"`
	OfCloudAzureAISearchVectorStore *shared.CloudAzureAISearchVectorStoreParam `json:",omitzero,inline"`
	OfCloudMongoDBAtlasVectorSearch *shared.CloudMongoDBAtlasVectorSearchParam `json:",omitzero,inline"`
	OfCloudMilvusVectorStore        *shared.CloudMilvusVectorStoreParam        `json:",omitzero,inline"`
	OfCloudAstraDBVectorStore       *shared.CloudAstraDBVectorStoreParam       `json:",omitzero,inline"`
	paramUnion
}

func (u DataSinkUpdateParamsComponentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap,
		u.OfCloudPineconeVectorStore,
		u.OfCloudPostgresVectorStore,
		u.OfCloudQdrantVectorStore,
		u.OfCloudAzureAISearchVectorStore,
		u.OfCloudMongoDBAtlasVectorSearch,
		u.OfCloudMilvusVectorStore,
		u.OfCloudAstraDBVectorStore)
}
func (u *DataSinkUpdateParamsComponentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type DataSinkListParams struct {
	OrganizationID param.Opt[string] `query:"organization_id,omitzero" format:"uuid" json:"-"`
	ProjectID      param.Opt[string] `query:"project_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [DataSinkListParams]'s query parameters as `url.Values`.
func (r DataSinkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
