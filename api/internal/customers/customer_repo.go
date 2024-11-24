package customers

import (
	"context"
	"errors"
	"tracked/internal/constants"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type CustomerRepository struct {
	client *dynamodb.Client
}

func NewCustomerRepository(client *dynamodb.Client) *CustomerRepository {
	return &CustomerRepository{client: client}
}

func (r *CustomerRepository) CreateCustomer(customer *Customer) error {
	av, err := marshalMap(customer)
	if err != nil {
		return err
	}

	alreadyExists, _ := r.GetCustomerByEmail(customer.Email)
	if alreadyExists != nil {
		return errors.New("Customer already exists")
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

func (r *CustomerRepository) GetCustomerByEmail(email string) (*Customer, error) {
	tableName := constants.TableName
	pk := GeneratePKByEmail(email)

	key := map[string]types.AttributeValue{
		"pk": &types.AttributeValueMemberS{Value: pk},
		"sk": &types.AttributeValueMemberS{Value: pk},
	}

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: &tableName,
	}

	result, err := r.client.GetItem(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, errors.New("Customer not found")
	}

	var user Customer
	err = unmarshalMap(result.Item, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func marshalMap(customer *Customer) (map[string]types.AttributeValue, error) {
	av, err := attributevalue.MarshalMap(customer)
	if err != nil {
		return nil, err
	}

	return av, nil
}

func unmarshalMap(av map[string]types.AttributeValue, v interface{}) error {
	err := attributevalue.UnmarshalMap(av, v)
	if err != nil {
		return err
	}

	return nil
}
