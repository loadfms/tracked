package workspaces

import (
	"context"
	"tracked/internal/constants"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type WorkspaceRepository struct {
	client *dynamodb.Client
}

func NewWorkspaceRepository(client *dynamodb.Client) *WorkspaceRepository {
	return &WorkspaceRepository{client: client}
}

func (r *WorkspaceRepository) CreateWorkspace(workspace *Workspace) error {
	av, err := marshalMap(workspace)
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

func (r *WorkspaceRepository) QueryWorkspaceByCustomer(customerUUID string) (*[]Workspace, error) {
	tableName := constants.TableName

	input := &dynamodb.QueryInput{
		TableName:              &tableName,
		KeyConditionExpression: aws.String("pk = :pk and begins_with(sk, :skPrefix)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk":       &types.AttributeValueMemberS{Value: customerUUID},
			":skPrefix": &types.AttributeValueMemberS{Value: "WORKSPACE##"},
		},
	}

	result, err := r.client.Query(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var workspaces []Workspace
	err = attributevalue.UnmarshalListOfMaps(result.Items, &workspaces)
	if err != nil {
		return nil, err
	}

	return &workspaces, nil
}

func marshalMap(workspace *Workspace) (map[string]types.AttributeValue, error) {
	av, err := attributevalue.MarshalMap(workspace)
	if err != nil {
		return nil, err
	}

	return av, nil
}
