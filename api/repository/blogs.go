package repository

import (
	"api/client"
	"api/entity"
	"api/utility"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type BlogsRepository interface {
	FindById(string) (*entity.Blog, error)
	FindAll() ([]*entity.Blog, error)
}

type blogsRepository struct {
	dynamodbClient client.DynamodbClient
	tableName      *string
}

func NewBlogsRepository(
	dynamodbClient client.DynamodbClient,
) BlogsRepository {
	return &blogsRepository{
		dynamodbClient: dynamodbClient,
		tableName:      aws.String(os.Getenv(utility.EnvKeyDynamodbTableNameMaster)),
	}
}

func (r *blogsRepository) FindById(id string) (*entity.Blog, error) {
	blog := &entity.Blog{
		Id: id,
	}
	key, err := blog.GetKey()
	if err != nil {
		return nil, err
	}
	res, err := r.dynamodbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: r.tableName,
		Key:       key,
	})
	if err != nil {
		return nil, err
	}
	if res.Item == nil {
		return nil, nil
	}
	if err := attributevalue.UnmarshalMap(res.Item, blog); err != nil {
		return nil, err
	}
	return blog, nil
}

func (r *blogsRepository) FindAll() ([]*entity.Blog, error) {
	keyCond := expression.Key(utility.DynamodbAttributeNamePartitionKey).Equal(
		expression.Value(utility.DynamodbMasterDataTypeBlog),
	)
	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		return nil, err
	}
	paginator := dynamodb.NewQueryPaginator(r.dynamodbClient, &dynamodb.QueryInput{
		TableName:                 r.tableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	items := []map[string]types.AttributeValue{}
	for paginator.HasMorePages() {
		res, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		items = append(items, res.Items...)
	}
	blogs := []*entity.Blog{}
	if err := attributevalue.UnmarshalListOfMaps(items, &blogs); err != nil {
		return nil, err
	}
	return blogs, nil
}
