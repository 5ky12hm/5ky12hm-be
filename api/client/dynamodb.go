package client

import (
	"api/utility"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamodbClient interface {
	GetItem(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

func NewDynamodbClient() DynamodbClient {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithBaseEndpoint(os.Getenv(utility.EnvKeyDynamodbEndpoint)),
		config.WithSharedConfigProfile(os.Getenv(utility.EnvKeyDynamodbSharedConfigProfile)),
	)
	if err != nil {
		panic(err)
	}
	return dynamodb.NewFromConfig(cfg)
}
