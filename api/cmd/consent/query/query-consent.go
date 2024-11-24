package main

import (
	"context"
	"fmt"
	"tracked/internal/consent"
	"tracked/pkg/responses"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	siteUUID := fmt.Sprintf("SITE##%s", request.PathParameters["siteUUID"])

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	client := dynamodb.NewFromConfig(cfg)

	consentRepo := consent.NewConsentRepository(client)
	items, err := consentRepo.QueryConsentBySiteUUID(siteUUID)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	return responses.Success(items), nil
}

func main() {
	lambda.Start(Handler)
}
