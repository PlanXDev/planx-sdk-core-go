package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
)

// SendGiftInfoForceExpire Forced expiration gift pack.
func (client *PlanXClient) SendGiftInfoForceExpire(request *SendGiftInfoForceExpireRequest) (response *SendGiftInfoForceExpireResponse, err error) {
	response = CreateSendGiftInfoForceExpireResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// SendGiftInfoForceExpireWithChan Forced expiration gift pack.
func (client *PlanXClient) SendGiftInfoForceExpireWithChan(request *SendGiftInfoForceExpireRequest) (<-chan *SendGiftInfoForceExpireResponse, <-chan error) {
	responseChan := make(chan *SendGiftInfoForceExpireResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendGiftInfoForceExpire(request)
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

// SendGiftInfoForceExpireRequest is the request struct for api SendGiftInfoForceExpire
type SendGiftInfoForceExpireRequest struct {
	*requests.BaseRequest
	GiftSourceQrCode string `position:"Body" name:"giftSourceQrCode"  binding:"required"` //The unique CODE of the gift source
}

// SendGiftInfoForceExpireResponse is the request struct for api SendGiftInfoForceExpire
type SendGiftInfoForceExpireResponse struct {
	*responses.BaseResponse
	Data *ResponseGiftSourceExpire `json:"data"`
}

// CreateSendGiftInfoForceExpireRequest creates a request to invoke SendGiftInfoForceExpire API
func CreateSendGiftInfoForceExpireRequest(giftSourceQrCode string) (request *SendGiftInfoForceExpireRequest) {
	request = &SendGiftInfoForceExpireRequest{
		BaseRequest:      requests.NewPostRequest("/v1/api/gift/info/forceExpire"),
		GiftSourceQrCode: giftSourceQrCode,
	}
	return
}

// CreateSendGiftInfoForceExpireResponse creates a response to parse from SendGiftInfoForceExpire response
func CreateSendGiftInfoForceExpireResponse() (response *SendGiftInfoForceExpireResponse) {
	response = &SendGiftInfoForceExpireResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
