package sites

import (
	"tracked/internal/helpers"
)

type Site struct {
	PK     string `dynamodbav:"pk"`
	SK     string `dynamodbav:"sk"`
	Site   string `json:"site" dynamodbav:"site"`
	Domain string `json:"domain" dynamodbav:"domain"`
}

func NewSite(site string, workspaceUUID string, domain string) (*Site, error) {
	sk := helpers.GenerateSK("SITE")

	return &Site{
		Site:   site,
		PK:     workspaceUUID,
		SK:     sk,
		Domain: domain,
	}, nil
}
