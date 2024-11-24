package consent

import (
	"fmt"
	"time"
)

type Consent struct {
	PK                string    `dynamodbav:"pk"`
	SK                string    `dynamodbav:"sk"`
	AcceptedCookiesID []string  `json:"acceptedCookiesID" dynamodbav:"acceptedCookiesID"`
	UpdatedAt         time.Time `json:"updatedAt" dynamodbav:"updatedAt"`
	Ip                string    `json:"ip" dynamodbav:"ip"`
}

func NewConsent(siteUUID string, anonymousID string, acceptedCookiesID []string, ip string) (*Consent, error) {
	sk := fmt.Sprintf("CONSENT##%s", anonymousID)

	return &Consent{
		PK:                siteUUID,
		SK:                sk,
		AcceptedCookiesID: acceptedCookiesID,
		UpdatedAt:         time.Now(),
		Ip:                ip,
	}, nil
}
