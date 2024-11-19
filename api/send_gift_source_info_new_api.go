package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
)

// SendGiftSourceInfoNew Create a new superior gift pack.
func (client *PlanXClient) SendGiftSourceInfoNew(request *SendGiftSourceInfoNewRequest) (response *SendGiftSourceInfoNewResponse, err error) {
	response = CreateSendGiftSourceInfoNewResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// SendGiftSourceInfoNewWithChan Create a new superior gift pack.
func (client *PlanXClient) SendGiftSourceInfoNewWithChan(request *SendGiftSourceInfoNewRequest) (<-chan *SendGiftSourceInfoNewResponse, <-chan error) {
	responseChan := make(chan *SendGiftSourceInfoNewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendGiftSourceInfoNew(request)
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

// SendGiftSourceInfoNewRequest is the request struct for api SendGiftSourceInfoNew
type SendGiftSourceInfoNewRequest struct {
	*requests.BaseRequest
	PoolId         string `position:"Body" name:"poolId" binding:"required"`                  //The unique identifier of the capital pool
	GiftName       string `position:"Body" name:"giftName" binding:"required"`                //Gift name
	GiftType       string `position:"Body" name:"giftType" binding:"required"`                //api.GiftTypeAirdrop,api.GiftTypeRedemption
	Quantity       int    `position:"Body" name:"quantity" binding:"required,gte=1,lte=2000"` //Generate quantity,1<=quantity<=2000
	PriceAmount    string `position:"Body" name:"priceAmount" binding:"required,gt=0"`        //Number of tokens for a single Gift
	ExpiresSeconds int64  `position:"Body" name:"expiresSeconds" binding:"required,gte=1"`    //Expiration time.Expires after specified number of seconds
}

// SendGiftSourceInfoNewResponse is the request struct for api SendGiftSourceInfoNew
type SendGiftSourceInfoNewResponse struct {
	*responses.BaseResponse
	Data *ResponseGiftSourceInfo `json:"data"`
}

// CreateSendGiftSourceInfoNewRequest creates a request to invoke SendGiftSourceInfoNew API
func CreateSendGiftSourceInfoNewRequest(poolId, giftName, giftType string, quantity int, priceAmount string, expiresSeconds int64) (request *SendGiftSourceInfoNewRequest) {
	request = &SendGiftSourceInfoNewRequest{
		BaseRequest:    requests.NewPostRequest("/v1/api/gift/info/new"),
		PoolId:         poolId,
		GiftName:       giftName,
		GiftType:       giftType,
		Quantity:       quantity,
		PriceAmount:    priceAmount,
		ExpiresSeconds: expiresSeconds,
	}
	return
}

// CreateSendGiftSourceInfoNewResponse creates a response to parse from SendGiftSourceInfoNew response
func CreateSendGiftSourceInfoNewResponse() (response *SendGiftSourceInfoNewResponse) {
	response = &SendGiftSourceInfoNewResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
