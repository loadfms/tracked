package responses

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func JSONResponse(status int, data interface{}) events.APIGatewayProxyResponse {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return events.APIGatewayProxyResponse{}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type, Authorization",
		},
		Body: string(jsonData),
	}
}

func BadRequest(data interface{}) events.APIGatewayProxyResponse {
	return JSONResponse(http.StatusBadRequest, map[string]interface{}{"error": data})
}

func InternalServerError(data interface{}) events.APIGatewayProxyResponse {
	return JSONResponse(http.StatusInternalServerError, map[string]interface{}{"error": data})
}

func Success(data interface{}) events.APIGatewayProxyResponse {
	return JSONResponse(http.StatusOK, map[string]interface{}{"data": data})
}

func Unauthorized(data interface{}) events.APIGatewayProxyResponse {
	return JSONResponse(http.StatusUnauthorized, map[string]interface{}{"error": data})
}
