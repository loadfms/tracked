package main

import (
	"context"
	"encoding/json"
	"strings"
	"tracked/internal/customers"
	"tracked/internal/workspaces"
	"tracked/pkg/responses"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ReqBody struct {
	Name string `json:"name"`
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody ReqBody
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return responses.BadRequest(err.Error()), nil
	}

	authorizationHeader := request.Headers["authorization"]
	if authorizationHeader == "" {
		return responses.Unauthorized("Authorization header missing"), nil
	}

	tokenParts := strings.Split(authorizationHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return responses.Unauthorized("Invalid token format"), nil
	}

	token := tokenParts[1]
	customerUUID, err := customers.GetCustomerUUIDFromToken(token)
	if err != nil {
		return responses.Unauthorized(err.Error()), nil
	}

	workspace, err := workspaces.NewWorkspace(reqBody.Name, customerUUID)
	if err != nil {
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	client := dynamodb.NewFromConfig(cfg)

	workspaceRepo := workspaces.NewWorkspaceRepository(client)
	err = workspaceRepo.CreateWorkspace(workspace)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	return responses.Success("created"), nil
}

func main() {
	lambda.Start(Handler)
}
