package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
)

// GetAccountBalancesPool Obtain the balance information of all fund pool accounts.
func (client *PlanXClient) GetAccountBalancesPool(request *GetAccountBalancesPoolRequest) (response *GetAccountBalancesPoolResponse, err error) {
	response = CreateGetAccountBalancesPoolResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// GetAccountBalancesPoolWithChan Obtain the balance information of all fund pool accounts.
func (client *PlanXClient) GetAccountBalancesPoolWithChan(request *GetAccountBalancesPoolRequest) (<-chan *GetAccountBalancesPoolResponse, <-chan error) {
	responseChan := make(chan *GetAccountBalancesPoolResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccountBalancesPool(request)
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

// GetAccountBalancesPoolRequest is the request struct for api GetAccountBalancesPool
type GetAccountBalancesPoolRequest struct {
	*requests.BaseRequest
}

// GetAccountBalancesPoolResponse is the request struct for api GetAccountBalancesPool
type GetAccountBalancesPoolResponse struct {
	*responses.BaseResponse
	Data []*ResponseBalancesPool `json:"data"`
}

// CreateGetAccountBalancesPoolRequest creates a request to invoke GetAccountBalancesPool API
func CreateGetAccountBalancesPoolRequest() (request *GetAccountBalancesPoolRequest) {
	request = &GetAccountBalancesPoolRequest{
		BaseRequest: requests.NewGetRequest("/v1/api/account/balances/pool"),
	}
	return
}

// CreateGetAccountBalancesPoolResponse creates a response to parse from GetAccountBalancesPool response
func CreateGetAccountBalancesPoolResponse() (response *GetAccountBalancesPoolResponse) {
	response = &GetAccountBalancesPoolResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
