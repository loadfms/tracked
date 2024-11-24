package main

import (
	"context"
	"encoding/json"
	"tracked/internal/cookie"
	"tracked/pkg/responses"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ReqBody struct {
	SiteUUID string `json:"site_uuid"`
	Name     string `json:"name"`
	Purpose  string `json:"purpose"`
	Duration string `json:"duration"`
	Provider string `json:"provider"`
	Required bool   `json:"required"`
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var reqBody ReqBody
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return responses.BadRequest(err.Error()), nil
	}

	reqCookie, err := cookie.NewCookie(reqBody.SiteUUID, reqBody.Name, reqBody.Purpose, reqBody.Duration, reqBody.Provider, reqBody.Required)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	client := dynamodb.NewFromConfig(cfg)

	cookieRepo := cookie.NewCookieRepository(client)
	err = cookieRepo.CreateCookie(reqCookie)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	return responses.Success("created"), nil
}

func main() {
	lambda.Start(Handler)
}
