package api

import (
	"planx-sdk-core-go/core/requests"
	"planx-sdk-core-go/core/responses"
)

// GetTradePendingList Get all unfinished transactions.
func (client *PlanXClient) GetTradePendingList(request *GetTradePendingListRequest) (response *GetTradePendingListResponse, err error) {
	response = CreateGetTradePendingListResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// GetTradePendingListWithChan Get all unfinished transactions.
func (client *PlanXClient) GetTradePendingListWithChan(request *GetTradePendingListRequest) (<-chan *GetTradePendingListResponse, <-chan error) {
	responseChan := make(chan *GetTradePendingListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTradePendingList(request)
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

// GetTradePendingListRequest is the request struct for api GetTradePendingList
type GetTradePendingListRequest struct {
	*requests.BaseRequest
}

// GetTradePendingListResponse is the request struct for api GetTradePendingList
type GetTradePendingListResponse struct {
	*responses.BaseResponse
	Data []*ResponseTradeToken `json:"data"`
}

// CreateGetTradePendingListRequest creates a request to invoke GetTradePendingList API
func CreateGetTradePendingListRequest() (request *GetTradePendingListRequest) {
	request = &GetTradePendingListRequest{
		BaseRequest: requests.NewGetRequest("/v1/api/trade/pending/list"),
	}
	return
}

// CreateGetTradePendingListResponse creates a response to parse from GetTradePendingList response
func CreateGetTradePendingListResponse() (response *GetTradePendingListResponse) {
	response = &GetTradePendingListResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
