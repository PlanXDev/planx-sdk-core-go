package planx_sdk_core_go

import (
	"encoding/json"
	"planx-sdk-core-go/api"
	"planx-sdk-core-go/core"
	"planx-sdk-core-go/core/credential"
	"testing"
)

func TestClient(t *testing.T) {
	config := core.NewConfig()
	config.AutoRetry = true
	accessKeyCredential := credential.NewAccessKeyCredential(
		"http://192.168.1.200:8886",
		"cg7Zqpb4",
		"842c2d804f5b5e9bba7a35436d9585b26bca4c69",
	)
	client, _ := api.NewClientWithOptions(config, accessKeyCredential)
	client.OpenLogger()
	request := api.CreateGetGiftSourceInfoBatchRequest([]string{
		"1E44BA0A-6D95-450B-9D0D-9F934FC5EFCA",
		"215FF713-D8C5-44DD-9597-9CCD722BE39E",
		"498DDBA1-E4AD-45CD-A9A1-27A11DDC201B",
	})
	response, err := client.GetGiftSourceInfoBatch(request)
	if err != nil {
		println(err.Error())
	} else {
		if len(response.Data) > 0 {
			for _, datum := range response.Data {
				marshal, _ := json.Marshal(datum)
				println(string(marshal))
			}
		}
	}
}
