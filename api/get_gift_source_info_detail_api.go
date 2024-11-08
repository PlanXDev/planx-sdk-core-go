package api

import (
	"planx-sdk-core-go/core/requests"
	"planx-sdk-core-go/core/responses"
)

// GetGiftSourceInfoDetail Check the details of the superior gift pack corresponding to the Code.
func (client *PlanXClient) GetGiftSourceInfoDetail(request *GetGiftSourceInfoDetailRequest) (response *GetGiftSourceInfoDetailResponse, err error) {
	response = CreateGetGiftSourceInfoDetailResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// GetGiftSourceInfoDetailWithChan Check the details of the superior gift pack corresponding to the Code.
func (client *PlanXClient) GetGiftSourceInfoDetailWithChan(request *GetGiftSourceInfoDetailRequest) (<-chan *GetGiftSourceInfoDetailResponse, <-chan error) {
	responseChan := make(chan *GetGiftSourceInfoDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGiftSourceInfoDetail(request)
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

// GetGiftSourceInfoDetailRequest is the request struct for api GetGiftSourceInfoDetail
type GetGiftSourceInfoDetailRequest struct {
	*requests.BaseRequest
	GiftSourceQrCode string `position:"Body" name:"giftSourceQrCode" binding:"required"` //The unique CODE of the gift source
}

// GetGiftSourceInfoDetailResponse is the request struct for api GetGiftSourceInfoDetail
type GetGiftSourceInfoDetailResponse struct {
	*responses.BaseResponse
	Data *ResponseGiftSourceDetail `json:"data"`
}

// CreateGetGiftSourceInfoDetailRequest creates a request to invoke GetGiftSourceInfoDetail API
func CreateGetGiftSourceInfoDetailRequest(giftSourceQrCode string) (request *GetGiftSourceInfoDetailRequest) {
	request = &GetGiftSourceInfoDetailRequest{
		BaseRequest:      requests.NewPostRequest("/v1/api/gift/info/detail"),
		GiftSourceQrCode: giftSourceQrCode,
	}
	return
}

// CreateGetGiftSourceInfoDetailResponse creates a response to parse from GetGiftSourceInfoDetail response
func CreateGetGiftSourceInfoDetailResponse() (response *GetGiftSourceInfoDetailResponse) {
	response = &GetGiftSourceInfoDetailResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
