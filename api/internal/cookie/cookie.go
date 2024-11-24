package cookie

import (
	"tracked/internal/helpers"
)

type Cookie struct {
	PK       string `dynamodbav:"pk"`
	SK       string `dynamodbav:"sk"`
	Name     string `json:"name" dynamodbav:"name"`
	Purpose  string `json:"purpose" dynamodbav:"purpose"`
	Duration string `json:"duration" dynamodbav:"duration"`
	Provider string `json:"provider" dynamodbav:"provider"`
	Required bool   `json:"required" dynamodbav:"required"`
}

func NewCookie(siteUUID string, name string, purpose string, duration string, provider string, required bool) (*Cookie, error) {
	sk := helpers.GenerateSK("COOKIE")

	return &Cookie{
		PK:       siteUUID,
		SK:       sk,
		Name:     name,
		Purpose:  purpose,
		Duration: duration,
		Provider: provider,
		Required: required,
	}, nil
}
