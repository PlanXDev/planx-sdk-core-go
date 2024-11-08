package planx_sdk_core_go

import (
	"github.com/PlanXDev/planx-sdk-core-go/api"
	"github.com/PlanXDev/planx-sdk-core-go/core"
	"github.com/PlanXDev/planx-sdk-core-go/core/credential"
	"testing"
)

func TestClient(t *testing.T) {
	baseUrl := "https://cp-api.planx.io"
	appId := "Go to the PlanX CP website to get"
	secretKey := "Go to the PlanX CP website to get"
	client, _ := api.NewClientWithOptions(core.NewConfig(), credential.NewAccessKeyCredential(baseUrl, appId, secretKey))
	base, err := client.GetAccountBalancesBase(api.CreateGetAccountBalancesBaseRequest())
	if err != nil {
		println(err.Error())
	}
	println(base.String())

	// Process your bussiness ...
}
