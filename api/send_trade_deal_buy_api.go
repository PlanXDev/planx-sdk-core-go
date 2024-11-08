package api

import (
	"planx-sdk-core-go/core/requests"
	"planx-sdk-core-go/core/responses"
)

// SendTradeDealBuy The trade of purchasing the specified ID will deduct the corresponding funds from the market account.
//     * Please confirm whether the funds are sufficient when calling.
func (client *PlanXClient) SendTradeDealBuy(request *SendTradeDealBuyRequest) (response *SendTradeDealBuyResponse, err error) {
	response = CreateSendTradeDealBuyResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// SendTradeDealBuyWithChan The trade of purchasing the specified ID will deduct the corresponding funds from the market account.
//     * Please confirm whether the funds are sufficient when calling.
func (client *PlanXClient) SendTradeDealBuyWithChan(request *SendTradeDealBuyRequest) (<-chan *SendTradeDealBuyResponse, <-chan error) {
	responseChan := make(chan *SendTradeDealBuyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendTradeDealBuy(request)
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

// SendTradeDealBuyRequest is the request struct for api SendTradeDealBuy
type SendTradeDealBuyRequest struct {
	*requests.BaseRequest
	TradeId string `position:"Body" name:"tradeId" binding:"required"` //Unique identifier of the trade
}

// SendTradeDealBuyResponse is the request struct for api SendTradeDealBuy
type SendTradeDealBuyResponse struct {
	*responses.BaseResponse
	Data *ResponseTradeToken `json:"data"`
}

// CreateSendTradeDealBuyRequest creates a request to invoke SendTradeDealBuy API
func CreateSendTradeDealBuyRequest(tradeId string) (request *SendTradeDealBuyRequest) {
	request = &SendTradeDealBuyRequest{
		BaseRequest: requests.NewPostRequest("/v1/api/trade/deal/buy"),
		TradeId:     tradeId,
	}
	return
}

// CreateSendTradeDealBuyResponse creates a response to parse from SendTradeDealBuy response
func CreateSendTradeDealBuyResponse() (response *SendTradeDealBuyResponse) {
	response = &SendTradeDealBuyResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
