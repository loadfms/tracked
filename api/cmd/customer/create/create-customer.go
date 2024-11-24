package main

import (
	"context"
	"encoding/json"
	"tracked/internal/customers"
	"tracked/pkg/responses"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ReqBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody ReqBody
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return responses.BadRequest(err.Error()), nil
	}

	customer, err := customers.NewCustomer(reqBody.Name, reqBody.Email, reqBody.Password)
	if err != nil {
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	client := dynamodb.NewFromConfig(cfg)

	customerRepo := customers.NewCustomerRepository(client)
	err = customerRepo.CreateCustomer(customer)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	return responses.Success("created"), nil
}

func main() {
	lambda.Start(Handler)
}
