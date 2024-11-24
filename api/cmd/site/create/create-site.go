package main

import (
	"context"
	"encoding/json"
	"tracked/internal/sites"
	"tracked/pkg/responses"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ReqBody struct {
	Site          string `json:"site"`
	Domain        string `json:"domain"`
	WorkspaceUUID string `json:"workspace_uuid"`
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody ReqBody
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return responses.BadRequest(err.Error()), nil
	}

	site, err := sites.NewSite(reqBody.Site, reqBody.WorkspaceUUID, reqBody.Domain)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	client := dynamodb.NewFromConfig(cfg)

	siteRepo := sites.NewSiteRepository(client)
	err = siteRepo.CreateSite(site)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	return responses.Success("created"), nil
}

func main() {
	lambda.Start(Handler)
}
