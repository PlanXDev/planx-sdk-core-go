package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
)

// GetGiftSourceInfoBatch Check the details of the superior gift pack corresponding to the Code.
func (client *PlanXClient) GetGiftSourceInfoBatch(request *GetGiftSourceInfoBatchRequest) (response *GetGiftSourceInfoBatchResponse, err error) {
	response = CreateGetGiftSourceInfoBatchResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// GetGiftSourceInfoBatchWithChan Check the details of the superior gift pack corresponding to the Code.
func (client *PlanXClient) GetGiftSourceInfoBatchWithChan(request *GetGiftSourceInfoBatchRequest) (<-chan *GetGiftSourceInfoBatchResponse, <-chan error) {
	responseChan := make(chan *GetGiftSourceInfoBatchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGiftSourceInfoBatch(request)
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

// GetGiftSourceInfoBatchRequest is the request struct for api GetGiftSourceInfoBatch
type GetGiftSourceInfoBatchRequest struct {
	*requests.BaseRequest
	GiftSourceQrCodes []string `position:"Body" name:"giftSourceQrCodes"` //The unique CODE of the gift source
}

// GetGiftSourceInfoBatchResponse is the request struct for api GetGiftSourceInfoBatch
type GetGiftSourceInfoBatchResponse struct {
	*responses.BaseResponse
	Data []*ResponseGiftSourceDetail `json:"data"`
}

// CreateGetGiftSourceInfoBatchRequest creates a request to invoke GetGiftSourceInfoBatch API
func CreateGetGiftSourceInfoBatchRequest(qrCodes []string) (request *GetGiftSourceInfoBatchRequest) {
	request = &GetGiftSourceInfoBatchRequest{
		BaseRequest:       requests.NewPostRequest("/v1/api/gift/info/batch"),
		GiftSourceQrCodes: qrCodes,
	}
	return
}

// CreateGetGiftSourceInfoBatchResponse creates a response to parse from GetGiftSourceInfoBatch response
func CreateGetGiftSourceInfoBatchResponse() (response *GetGiftSourceInfoBatchResponse) {
	response = &GetGiftSourceInfoBatchResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
