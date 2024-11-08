package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
)

// GetAccountBalancesBase Get balance information for all token accounts.
func (client *PlanXClient) GetAccountBalancesBase(request *GetAccountBalancesBaseRequest) (response *GetAccountBalancesBaseResponse, err error) {
	response = CreateGetAccountBalancesBaseResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// GetAccountBalancesBaseWithChan Get balance information for all token accounts.
func (client *PlanXClient) GetAccountBalancesBaseWithChan(request *GetAccountBalancesBaseRequest) (<-chan *GetAccountBalancesBaseResponse, <-chan error) {
	responseChan := make(chan *GetAccountBalancesBaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccountBalancesBase(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetAccountBalancesBaseRequest is the request struct for api GetAccountBalancesBase
type GetAccountBalancesBaseRequest struct {
	*requests.BaseRequest
}

// GetAccountBalancesBaseResponse is the request struct for api GetAccountBalancesBase
type GetAccountBalancesBaseResponse struct {
	*responses.BaseResponse
	Data []*ResponseBalancesBase `json:"data"`
}

// CreateGetAccountBalancesBaseRequest creates a request to invoke GetAccountBalancesBase API
func CreateGetAccountBalancesBaseRequest() (request *GetAccountBalancesBaseRequest) {
	request = &GetAccountBalancesBaseRequest{
		BaseRequest: requests.NewGetRequest("/v1/api/account/balances/base"),
	}
	return
}

// CreateGetAccountBalancesBaseResponse creates a response to parse from GetAccountBalancesBase response
func CreateGetAccountBalancesBaseResponse() (response *GetAccountBalancesBaseResponse) {
	response = &GetAccountBalancesBaseResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
