package main

import (
	"context"
	"fmt"
	"strings"
	"tracked/internal/constants"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt"
)

func handler(ctx context.Context, request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	authorizationToken := request.AuthorizationToken
	if authorizationToken == "" {
		return generateAuthResponse("Unauthorized", "Missing authorization token", 401), nil
	}

	tokenParts := strings.Split(authorizationToken, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return generateAuthResponse("Unauthorized", fmt.Sprintf("Invalid token: %s", authorizationToken), 401), nil
	}

	tokenString := tokenParts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(constants.JWTTokenSalt), nil
	})

	if err != nil {
		return generateAuthResponse("Unauthorized", fmt.Sprintf("Invalid token: %s", authorizationToken), 401), nil
	}

	if !token.Valid {
		return generateAuthResponse("Unauthorized", fmt.Sprintf("Invalid token: %s", authorizationToken), 401), nil
	}

	return generateAuthResponse("Allow", "", 200), nil
}

func generateAuthResponse(effect string, message string, statusCode int) events.APIGatewayCustomAuthorizerResponse {
	return events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: "user",
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{"*"},
				},
			},
		},
		Context: map[string]interface{}{
			"message": message,
		},
	}
}

func main() {
	lambda.Start(handler)
}
