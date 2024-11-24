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
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody ReqBody
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return responses.BadRequest(err.Error()), nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	client := dynamodb.NewFromConfig(cfg)

	customerRepo := customers.NewCustomerRepository(client)

	customer, err := customerRepo.GetCustomerByEmail(reqBody.Email)
	if err != nil {
		return responses.BadRequest("invalid email or password" + err.Error()), nil
	}

	logginSucced := customers.CheckPassword(customer, reqBody.Password)
	if !logginSucced {
		return responses.Unauthorized("invalid email or password"), nil
	}

	customerToken, err := customers.GenerateJWTToken(customer)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	return responses.Success(customerToken), nil
}

func main() {
	lambda.Start(Handler)
}
