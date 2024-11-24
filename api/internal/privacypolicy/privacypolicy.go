package privacypolicy

import (
	"time"
	"tracked/internal/helpers"
)

type PrivacyPolicy struct {
	PK          string    `dynamodbav:"pk"`
	SK          string    `dynamodbav:"sk"`
	Content     string    `json:"content" dynamodbav:"content"`
	Responsible string    `json:"responsible" dynamodbav:"responsible"`
	UpdatedAt   time.Time `dynamodbav:"updated_at"`
}

func NewPrivacyPolicy(siteUUID string, content string, responsible string) (*PrivacyPolicy, error) {
	sk := helpers.GenerateSK("PRIVACYPOLICY")

	return &PrivacyPolicy{
		PK:          siteUUID,
		SK:          sk,
		Content:     content,
		Responsible: responsible,
		UpdatedAt:   time.Now(),
	}, nil
}
