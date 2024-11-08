package api

import (
	"planx-sdk-core-go/core/requests"
	"planx-sdk-core-go/core/responses"
)

// GetAccountBalancesMarket Get balance information for all market accounts.
func (client *PlanXClient) GetAccountBalancesMarket(request *GetAccountBalancesMarketRequest) (response *GetAccountBalancesMarketResponse, err error) {
	response = CreateGetAccountBalancesMarketResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// GetAccountBalancesMarketWithChan Get balance information for all market accounts.
func (client *PlanXClient) GetAccountBalancesMarketWithChan(request *GetAccountBalancesMarketRequest) (<-chan *GetAccountBalancesMarketResponse, <-chan error) {
	responseChan := make(chan *GetAccountBalancesMarketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAccountBalancesMarket(request)
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

// GetAccountBalancesMarketRequest is the request struct for api GetAccountBalancesMarket
type GetAccountBalancesMarketRequest struct {
	*requests.BaseRequest
}

// GetAccountBalancesMarketResponse is the request struct for api GetAccountBalancesMarket
type GetAccountBalancesMarketResponse struct {
	*responses.BaseResponse
	Data []*ResponseBalancesMarket `json:"data"`
}

// CreateGetAccountBalancesMarketRequest creates a request to invoke GetAccountBalancesMarket API
func CreateGetAccountBalancesMarketRequest() (request *GetAccountBalancesMarketRequest) {
	request = &GetAccountBalancesMarketRequest{
		BaseRequest: requests.NewGetRequest("/v1/api/account/balances/market"),
	}
	return
}

// CreateGetAccountBalancesMarketResponse creates a response to parse from GetAccountBalancesMarket response
func CreateGetAccountBalancesMarketResponse() (response *GetAccountBalancesMarketResponse) {
	response = &GetAccountBalancesMarketResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
