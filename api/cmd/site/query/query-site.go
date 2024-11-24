package main

import (
	"context"
	"fmt"
	"strings"
	"tracked/internal/customers"
	"tracked/internal/sites"
	"tracked/internal/workspaces"
	"tracked/pkg/responses"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	workspaceUUID := fmt.Sprintf("WORKSPACE##%s", request.PathParameters["workspaceUUID"])

	authorizationHeader := request.Headers["authorization"]
	if authorizationHeader == "" {
		return responses.Unauthorized("Authorization header missing"), nil
	}

	tokenParts := strings.Split(authorizationHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return responses.Unauthorized("Invalid token format"), nil
	}

	token := tokenParts[1]
	_, err := customers.GetCustomerUUIDFromToken(token)
	if err != nil {
		return responses.Unauthorized(err.Error()), nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	client := dynamodb.NewFromConfig(cfg)

	siteRepo := sites.NewSiteRepository(client)
	items, err := siteRepo.QuerySitesByWorkspace(workspaceUUID)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	if items == nil || len(*items) == 0 {
		return responses.Success([]workspaces.Workspace{}), nil
	}

	return responses.Success(items), nil
}

func main() {
	lambda.Start(Handler)
}
