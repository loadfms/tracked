package privacypolicy

import (
	"context"
	"tracked/internal/constants"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type PrivacyPolicyRepository struct {
	client *dynamodb.Client
}

func NewPrivacyPolicyRepository(client *dynamodb.Client) *PrivacyPolicyRepository {
	return &PrivacyPolicyRepository{client: client}
}

func (r *PrivacyPolicyRepository) CreatePrivacyPolicy(privacypolicy *PrivacyPolicy) error {
	av, err := marshalMap(privacypolicy)
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

func (r *PrivacyPolicyRepository) QueryPrivacyPolicyBySiteUUID(siteUUID string) (*[]PrivacyPolicy, error) {
	tableName := constants.TableName

	input := &dynamodb.QueryInput{
		TableName:              &tableName,
		KeyConditionExpression: aws.String("pk = :pk and begins_with(sk, :skPrefix)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk":       &types.AttributeValueMemberS{Value: siteUUID},
			":skPrefix": &types.AttributeValueMemberS{Value: "PRIVACYPOLICY##"},
		},
	}

	result, err := r.client.Query(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var privacypolicies []PrivacyPolicy
	err = attributevalue.UnmarshalListOfMaps(result.Items, &privacypolicies)
	if err != nil {
		return nil, err
	}

	return &privacypolicies, nil
}

func marshalMap(privacypolicy *PrivacyPolicy) (map[string]types.AttributeValue, error) {
	av, err := attributevalue.MarshalMap(privacypolicy)
	if err != nil {
		return nil, err
	}

	return av, nil
}
