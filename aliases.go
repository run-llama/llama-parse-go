// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamacloudprod

import (
	"github.com/run-llama/llama-parse-go/internal/apierror"
	"github.com/run-llama/llama-parse-go/packages/param"
	"github.com/run-llama/llama-parse-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

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
// This is an alias to an internal type.
type CloudAstraDBVectorStore = shared.CloudAstraDBVectorStore

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
// This is an alias to an internal type.
type CloudAstraDBVectorStoreParam = shared.CloudAstraDBVectorStoreParam

// This is an alias to an internal type.
type CloudAzStorageBlobDataSource = shared.CloudAzStorageBlobDataSource

// This is an alias to an internal type.
type CloudAzStorageBlobDataSourceParam = shared.CloudAzStorageBlobDataSourceParam

// Cloud Azure AI Search Vector Store.
//
// This is an alias to an internal type.
type CloudAzureAISearchVectorStore = shared.CloudAzureAISearchVectorStore

// Cloud Azure AI Search Vector Store.
//
// This is an alias to an internal type.
type CloudAzureAISearchVectorStoreParam = shared.CloudAzureAISearchVectorStoreParam

// This is an alias to an internal type.
type CloudBoxDataSource = shared.CloudBoxDataSource

// The type of authentication to use (Developer Token or CCG)
//
// This is an alias to an internal type.
type CloudBoxDataSourceAuthenticationMechanism = shared.CloudBoxDataSourceAuthenticationMechanism

// Equals "ccg"
const CloudBoxDataSourceAuthenticationMechanismCcg = shared.CloudBoxDataSourceAuthenticationMechanismCcg

// Equals "developer_token"
const CloudBoxDataSourceAuthenticationMechanismDeveloperToken = shared.CloudBoxDataSourceAuthenticationMechanismDeveloperToken

// This is an alias to an internal type.
type CloudBoxDataSourceParam = shared.CloudBoxDataSourceParam

// This is an alias to an internal type.
type CloudConfluenceDataSource = shared.CloudConfluenceDataSource

// This is an alias to an internal type.
type CloudConfluenceDataSourceParam = shared.CloudConfluenceDataSourceParam

// This is an alias to an internal type.
type CloudGoogleDriveDataSource = shared.CloudGoogleDriveDataSource

// This is an alias to an internal type.
type CloudGoogleDriveDataSourceParam = shared.CloudGoogleDriveDataSourceParam

// Cloud Jira Data Source integrating JiraReader.
//
// This is an alias to an internal type.
type CloudJiraDataSource = shared.CloudJiraDataSource

// Cloud Jira Data Source integrating JiraReader.
//
// This is an alias to an internal type.
type CloudJiraDataSourceParam = shared.CloudJiraDataSourceParam

// Cloud Jira Data Source integrating JiraReaderV2.
//
// This is an alias to an internal type.
type CloudJiraDataSourceV2 = shared.CloudJiraDataSourceV2

// Jira REST API version to use (2 or 3). 3 supports Atlassian Document Format
// (ADF).
//
// This is an alias to an internal type.
type CloudJiraDataSourceV2APIVersion = shared.CloudJiraDataSourceV2APIVersion

// Equals "2"
const CloudJiraDataSourceV2APIVersion2 = shared.CloudJiraDataSourceV2APIVersion2

// Equals "3"
const CloudJiraDataSourceV2APIVersion3 = shared.CloudJiraDataSourceV2APIVersion3

// Cloud Jira Data Source integrating JiraReaderV2.
//
// This is an alias to an internal type.
type CloudJiraDataSourceV2Param = shared.CloudJiraDataSourceV2Param

// Cloud Milvus Vector Store.
//
// This is an alias to an internal type.
type CloudMilvusVectorStore = shared.CloudMilvusVectorStore

// Cloud Milvus Vector Store.
//
// This is an alias to an internal type.
type CloudMilvusVectorStoreParam = shared.CloudMilvusVectorStoreParam

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
// This is an alias to an internal type.
type CloudMongoDBAtlasVectorSearch = shared.CloudMongoDBAtlasVectorSearch

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
// This is an alias to an internal type.
type CloudMongoDBAtlasVectorSearchParam = shared.CloudMongoDBAtlasVectorSearchParam

// This is an alias to an internal type.
type CloudNotionPageDataSource = shared.CloudNotionPageDataSource

// This is an alias to an internal type.
type CloudNotionPageDataSourceParam = shared.CloudNotionPageDataSourceParam

// This is an alias to an internal type.
type CloudOneDriveDataSource = shared.CloudOneDriveDataSource

// This is an alias to an internal type.
type CloudOneDriveDataSourceParam = shared.CloudOneDriveDataSourceParam

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
// This is an alias to an internal type.
type CloudPineconeVectorStore = shared.CloudPineconeVectorStore

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
// This is an alias to an internal type.
type CloudPineconeVectorStoreParam = shared.CloudPineconeVectorStoreParam

// This is an alias to an internal type.
type CloudPostgresVectorStore = shared.CloudPostgresVectorStore

// This is an alias to an internal type.
type CloudPostgresVectorStoreParam = shared.CloudPostgresVectorStoreParam

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
// This is an alias to an internal type.
type CloudQdrantVectorStore = shared.CloudQdrantVectorStore

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
// This is an alias to an internal type.
type CloudQdrantVectorStoreParam = shared.CloudQdrantVectorStoreParam

// This is an alias to an internal type.
type CloudS3DataSource = shared.CloudS3DataSource

// This is an alias to an internal type.
type CloudS3DataSourceParam = shared.CloudS3DataSourceParam

// This is an alias to an internal type.
type CloudSharepointDataSource = shared.CloudSharepointDataSource

// This is an alias to an internal type.
type CloudSharepointDataSourceParam = shared.CloudSharepointDataSourceParam

// This is an alias to an internal type.
type CloudSlackDataSource = shared.CloudSlackDataSource

// This is an alias to an internal type.
type CloudSlackDataSourceParam = shared.CloudSlackDataSourceParam

// Configuration for handling different types of failures during data source
// processing.
//
// This is an alias to an internal type.
type FailureHandlingConfig = shared.FailureHandlingConfig

// Configuration for handling different types of failures during data source
// processing.
//
// This is an alias to an internal type.
type FailureHandlingConfigParam = shared.FailureHandlingConfigParam

// HNSW settings for PGVector.
//
// This is an alias to an internal type.
type PgVectorHnswSettings = shared.PgVectorHnswSettings

// The distance method to use.
//
// This is an alias to an internal type.
type PgVectorHnswSettingsDistanceMethod = shared.PgVectorHnswSettingsDistanceMethod

// Equals "cosine"
const PgVectorHnswSettingsDistanceMethodCosine = shared.PgVectorHnswSettingsDistanceMethodCosine

// Equals "hamming"
const PgVectorHnswSettingsDistanceMethodHamming = shared.PgVectorHnswSettingsDistanceMethodHamming

// Equals "ip"
const PgVectorHnswSettingsDistanceMethodIP = shared.PgVectorHnswSettingsDistanceMethodIP

// Equals "jaccard"
const PgVectorHnswSettingsDistanceMethodJaccard = shared.PgVectorHnswSettingsDistanceMethodJaccard

// Equals "l1"
const PgVectorHnswSettingsDistanceMethodL1 = shared.PgVectorHnswSettingsDistanceMethodL1

// Equals "l2"
const PgVectorHnswSettingsDistanceMethodL2 = shared.PgVectorHnswSettingsDistanceMethodL2

// The type of vector to use.
//
// This is an alias to an internal type.
type PgVectorHnswSettingsVectorType = shared.PgVectorHnswSettingsVectorType

// Equals "bit"
const PgVectorHnswSettingsVectorTypeBit = shared.PgVectorHnswSettingsVectorTypeBit

// Equals "half_vec"
const PgVectorHnswSettingsVectorTypeHalfVec = shared.PgVectorHnswSettingsVectorTypeHalfVec

// Equals "sparse_vec"
const PgVectorHnswSettingsVectorTypeSparseVec = shared.PgVectorHnswSettingsVectorTypeSparseVec

// Equals "vector"
const PgVectorHnswSettingsVectorTypeVector = shared.PgVectorHnswSettingsVectorTypeVector

// HNSW settings for PGVector.
//
// This is an alias to an internal type.
type PgVectorHnswSettingsParam = shared.PgVectorHnswSettingsParam
