package sites

import (
	"context"
	"tracked/internal/constants"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type SiteRepository struct {
	client *dynamodb.Client
}

func NewSiteRepository(client *dynamodb.Client) *SiteRepository {
	return &SiteRepository{client: client}
}

func (r *SiteRepository) CreateSite(site *Site) error {
	av, err := marshalMap(site)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(constants.TableName),
	}

	_, err = r.client.PutItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

func (r *SiteRepository) QuerySitesByWorkspace(workspaceUUID string) (*[]Site, error) {
	tableName := constants.TableName

	input := &dynamodb.QueryInput{
		TableName:              &tableName,
		KeyConditionExpression: aws.String("pk = :pk and begins_with(sk, :skPrefix)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk":       &types.AttributeValueMemberS{Value: workspaceUUID},
			":skPrefix": &types.AttributeValueMemberS{Value: "SITE##"},
		},
	}

	result, err := r.client.Query(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var workspaces []Site
	err = attributevalue.UnmarshalListOfMaps(result.Items, &workspaces)
	if err != nil {
		return nil, err
	}

	return &workspaces, nil
}

func marshalMap(site *Site) (map[string]types.AttributeValue, error) {
	av, err := attributevalue.MarshalMap(site)
	if err != nil {
		return nil, err
	}

	return av, nil
}
