package utility

const (
	DynamodbMasterDataTypeBlog = "blg"

	DynamodbAttributeNamePartitionKey = "pk"
	DynamodbAttributeNameSortKey      = "sk"
)

const (
	EnvKeyApiEnv = "API_ENV"

	EnvKeyDynamodbEndpoint            = "DYNAMODB_ENDPOINT"
	EnvKeyDynamodbRegion              = "DYNAMODB_REGION"
	EnvKeyDynamodbSharedConfigProfile = "DYNAMODB_SHARED_CONFIG_PROFILE"
	EnvKeyDynamodbTableNameMaster     = "DYNAMODB_TABLE_NAME_MASTER"
)
