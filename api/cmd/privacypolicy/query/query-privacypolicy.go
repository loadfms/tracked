package main

import (
	"context"
	"encoding/json"
	"fmt"
	"tracked/internal/privacypolicy"
	"tracked/pkg/responses"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ReqBody struct {
	SiteUUID    string `json:"site_uuid"`
	Content     string `json:"content"`
	Responsible string `json:"responsible"`
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	siteUUID := fmt.Sprintf("SITE##%s", request.PathParameters["siteUUID"])

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

	privacyPolicyRepo := privacypolicy.NewPrivacyPolicyRepository(client)
	items, err := privacyPolicyRepo.QueryPrivacyPolicyBySiteUUID(siteUUID)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	return responses.Success(items), nil
}

func main() {
	lambda.Start(Handler)
}
