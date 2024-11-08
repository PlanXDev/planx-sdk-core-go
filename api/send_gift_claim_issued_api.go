package api

import (
	"planx-sdk-core-go/core/requests"
	"planx-sdk-core-go/core/responses"
)

// SendGiftClaimIssued Send a sub-gift pack through the superior gift pack code.
func (client *PlanXClient) SendGiftClaimIssued(request *SendGiftClaimIssuedRequest) (response *SendGiftClaimIssuedResponse, err error) {
	response = CreateSendGiftClaimIssuedResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// SendGiftClaimIssuedWithChan Send a sub-gift pack through the superior gift pack code.
func (client *PlanXClient) SendGiftClaimIssuedWithChan(request *SendGiftClaimIssuedRequest) (<-chan *SendGiftClaimIssuedResponse, <-chan error) {
	responseChan := make(chan *SendGiftClaimIssuedResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendGiftClaimIssued(request)
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

// SendGiftClaimIssuedRequest is the request struct for api SendGiftClaimIssued
type SendGiftClaimIssuedRequest struct {
	*requests.BaseRequest
	GiftSourceQrCode string `position:"Body" name:"giftSourceQrCode"  binding:"required"` //The unique CODE of the gift source
}

// SendGiftClaimIssuedResponse is the request struct for api SendGiftClaimIssued
type SendGiftClaimIssuedResponse struct {
	*responses.BaseResponse
	Data *ResponseIssuedGift `json:"data"`
}

// CreateSendGiftClaimIssuedRequest creates a request to invoke SendGiftClaimIssued API
func CreateSendGiftClaimIssuedRequest(giftSourceQrCode string) (request *SendGiftClaimIssuedRequest) {
	request = &SendGiftClaimIssuedRequest{
		BaseRequest:      requests.NewPostRequest("/v1/api/gift/claim/issued"),
		GiftSourceQrCode: giftSourceQrCode,
	}
	return
}

// CreateSendGiftClaimIssuedResponse creates a response to parse from SendGiftClaimIssued response
func CreateSendGiftClaimIssuedResponse() (response *SendGiftClaimIssuedResponse) {
	response = &SendGiftClaimIssuedResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
