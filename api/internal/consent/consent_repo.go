package consent

import (
	"context"
	"tracked/internal/constants"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ConsentRepository struct {
	client *dynamodb.Client
}

func NewConsentRepository(client *dynamodb.Client) *ConsentRepository {
	return &ConsentRepository{client: client}
}

func (r *ConsentRepository) CreateConsent(consent *Consent) error {
	av, err := marshalMap(consent)
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

func (r *ConsentRepository) QueryConsentBySiteUUID(siteUUID string) (*[]Consent, error) {
	tableName := constants.TableName

	input := &dynamodb.QueryInput{
		TableName:              &tableName,
		KeyConditionExpression: aws.String("pk = :pk and begins_with(sk, :skPrefix)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk":       &types.AttributeValueMemberS{Value: siteUUID},
			":skPrefix": &types.AttributeValueMemberS{Value: "CONSENT##"},
		},
	}

	result, err := r.client.Query(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var privacypolicies []Consent
	err = attributevalue.UnmarshalListOfMaps(result.Items, &privacypolicies)
	if err != nil {
		return nil, err
	}

	return &privacypolicies, nil
}

func marshalMap(consent *Consent) (map[string]types.AttributeValue, error) {
	av, err := attributevalue.MarshalMap(consent)
	if err != nil {
		return nil, err
	}

	return av, nil
}
