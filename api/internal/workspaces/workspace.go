package workspaces

import (
	"tracked/internal/helpers"
)

type Workspace struct {
	PK   string `dynamodbav:"pk"`
	SK   string `dynamodbav:"sk"`
	Name string `json:"name" dynamodbav:"name"`
}

func NewWorkspace(name string, customerUUID string) (*Workspace, error) {
	sk := helpers.GenerateSK("WORKSPACE")

	return &Workspace{
		Name: name,
		PK:   customerUUID,
		SK:   sk,
	}, nil
}
