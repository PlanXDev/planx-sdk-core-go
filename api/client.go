package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core"
	"github.com/PlanXDev/planx-sdk-core-go/core/credential"
)

// PlanXClient is the sdk client struct, each func corresponds to an OpenAPI
type PlanXClient struct {
	core.Client
}

// NewClientWithOptions creates a sdk client with Config/SecretKeyCredential
// this is the common api to create a sdk client
func NewClientWithOptions(config *core.Config, credential *credential.SecretKeyCredential) (client *PlanXClient, err error) {
	client = &PlanXClient{}
	err = client.InitWithOptions(config, credential)
	return
}
