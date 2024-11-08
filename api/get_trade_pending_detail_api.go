package api

import (
	"planx-sdk-core-go/core/requests"
	"planx-sdk-core-go/core/responses"
)

// GetTradePendingDetail Get details according to trade ID.
func (client *PlanXClient) GetTradePendingDetail(request *GetTradePendingDetailRequest) (response *GetTradePendingDetailResponse, err error) {
	response = CreateGetTradePendingDetailResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// GetTradePendingDetailWithChan Get details according to trade ID.
func (client *PlanXClient) GetTradePendingDetailWithChan(request *GetTradePendingDetailRequest) (<-chan *GetTradePendingDetailResponse, <-chan error) {
	responseChan := make(chan *GetTradePendingDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTradePendingDetail(request)
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

// GetTradePendingDetailRequest is the request struct for api GetTradePendingDetail
type GetTradePendingDetailRequest struct {
	*requests.BaseRequest
	TradeId string `position:"Body" name:"tradeId" binding:"required"` //Unique identifier of the trade
}

// GetTradePendingDetailResponse is the request struct for api GetTradePendingDetail
type GetTradePendingDetailResponse struct {
	*responses.BaseResponse
	Data []*ResponseTradeToken `json:"data"`
}

// CreateGetTradePendingDetailRequest creates a request to invoke GetTradePendingDetail API
func CreateGetTradePendingDetailRequest(tradeId string) (request *GetTradePendingDetailRequest) {
	request = &GetTradePendingDetailRequest{
		BaseRequest: requests.NewPostRequest("/v1/api/trade/pending/detail"),
		TradeId:     tradeId,
	}
	return
}

// CreateGetTradePendingDetailResponse creates a response to parse from GetTradePendingDetail response
func CreateGetTradePendingDetailResponse() (response *GetTradePendingDetailResponse) {
	response = &GetTradePendingDetailResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
