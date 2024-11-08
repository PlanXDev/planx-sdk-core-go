package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
)

// GetGiftClaimWaiting Get all available sub-gift packs.
func (client *PlanXClient) GetGiftClaimWaiting(request *GetGiftClaimWaitingRequest) (response *GetGiftClaimWaitingResponse, err error) {
	response = CreateGetGiftClaimWaitingResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// GetGiftClaimWaitingWithChan Get all available sub-gift packs.
func (client *PlanXClient) GetGiftClaimWaitingWithChan(request *GetGiftClaimWaitingRequest) (<-chan *GetGiftClaimWaitingResponse, <-chan error) {
	responseChan := make(chan *GetGiftClaimWaitingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGiftClaimWaiting(request)
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

// GetGiftClaimWaitingRequest is the request struct for api GetGiftClaimWaiting
type GetGiftClaimWaitingRequest struct {
	*requests.BaseRequest
}

// GetGiftClaimWaitingResponse is the request struct for api GetGiftClaimWaiting
type GetGiftClaimWaitingResponse struct {
	*responses.BaseResponse
	Data []*ResponseGift `json:"data"`
}

// CreateGetGiftClaimWaitingRequest creates a request to invoke GetGiftClaimWaiting API
func CreateGetGiftClaimWaitingRequest() (request *GetGiftClaimWaitingRequest) {
	request = &GetGiftClaimWaitingRequest{
		BaseRequest: requests.NewGetRequest("/v1/api/gift/claim/waiting"),
	}
	return
}

// CreateGetGiftClaimWaitingResponse creates a response to parse from GetGiftClaimWaiting response
func CreateGetGiftClaimWaitingResponse() (response *GetGiftClaimWaitingResponse) {
	response = &GetGiftClaimWaitingResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
