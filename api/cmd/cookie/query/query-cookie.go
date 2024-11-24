package main

import (
	"context"
	"fmt"
	"strings"
	"tracked/internal/cookie"
	"tracked/internal/customers"
	"tracked/pkg/responses"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	siteUUID := fmt.Sprintf("SITE##%s", request.PathParameters["siteUUID"])

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

	cookieRepo := cookie.NewCookieRepository(client)
	items, err := cookieRepo.QueryCookieBySiteUUID(siteUUID)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	return responses.Success(items), nil
}

func main() {
	lambda.Start(Handler)
}
