package cookie

import (
	"context"
	"tracked/internal/constants"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type CookieRepository struct {
	client *dynamodb.Client
}

func NewCookieRepository(client *dynamodb.Client) *CookieRepository {
	return &CookieRepository{client: client}
}

func (r *CookieRepository) CreateCookie(cookie *Cookie) error {
	av, err := marshalMap(cookie)
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

func (r *CookieRepository) QueryCookieBySiteUUID(siteUUID string) (*[]Cookie, error) {
	tableName := constants.TableName

	input := &dynamodb.QueryInput{
		TableName:              &tableName,
		KeyConditionExpression: aws.String("pk = :pk and begins_with(sk, :skPrefix)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk":       &types.AttributeValueMemberS{Value: siteUUID},
			":skPrefix": &types.AttributeValueMemberS{Value: "COOKIE##"},
		},
	}

	result, err := r.client.Query(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var privacypolicies []Cookie
	err = attributevalue.UnmarshalListOfMaps(result.Items, &privacypolicies)
	if err != nil {
		return nil, err
	}

	return &privacypolicies, nil
}

func marshalMap(cookie *Cookie) (map[string]types.AttributeValue, error) {
	av, err := attributevalue.MarshalMap(cookie)
	if err != nil {
		return nil, err
	}

	return av, nil
}
