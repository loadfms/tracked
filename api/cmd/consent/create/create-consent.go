package main

import (
	"context"
	"encoding/json"
	"tracked/internal/consent"
	"tracked/pkg/responses"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ReqBody struct {
	SiteUUID          string   `json:"site_uuid"`
	AnonymousID       string   `json:"anonymous_id"`
	AcceptedCookiesID []string `json:"accepted_cookies_id"`
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	ip := request.RequestContext.HTTP.SourceIP

	var reqBody ReqBody
	err := json.Unmarshal([]byte(request.Body), &reqBody)
	if err != nil {
		return responses.BadRequest(err.Error()), nil
	}

	reqConsent, err := consent.NewConsent(reqBody.SiteUUID, reqBody.AnonymousID, reqBody.AcceptedCookiesID, ip)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	client := dynamodb.NewFromConfig(cfg)

	consentRepo := consent.NewConsentRepository(client)
	err = consentRepo.CreateConsent(reqConsent)
	if err != nil {
		return responses.InternalServerError(err.Error()), nil
	}

	return responses.Success("created"), nil
}

func main() {
	lambda.Start(Handler)
}
