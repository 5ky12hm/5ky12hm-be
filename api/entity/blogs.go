package entity

import (
	"api/utility"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Blog struct {
	PrimaryKey
	DataType  string `dynamodbav:"dtp"`
	Id        string `dynamodbav:"bid"`
	Title     string `dynamodbav:"tit"`
	Body      string `dynamodbav:"bdy"`
	Tag       string `dynamodbav:"tag"`
	CreatedAt string `dynamodbav:"cat"`
	UpdatedAt string `dynamodbav:"uat"`
}

func (e *Blog) GetKey() (map[string]types.AttributeValue, error) {
	pk, err := attributevalue.Marshal(utility.DynamodbMasterDataTypeBlog)
	if err != nil {
		return nil, err
	}
	sk, err := attributevalue.Marshal(e.Id)
	if err != nil {
		return nil, err
	}
	return map[string]types.AttributeValue{
		utility.DynamodbAttributeNamePartitionKey: pk,
		utility.DynamodbAttributeNameSortKey:      sk,
	}, nil
}
