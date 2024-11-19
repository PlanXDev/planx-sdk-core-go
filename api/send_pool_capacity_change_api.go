package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
)

// SendPoolCapacityChange Increase or decrease the assets of the fund pool.
func (client *PlanXClient) SendPoolCapacityChange(request *SendPoolCapacityChangeRequest) (response *SendPoolCapacityChangeResponse, err error) {
	response = CreateSendPoolCapacityChangeResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// SendPoolCapacityChangeWithChan Increase or decrease the assets of the fund pool.
func (client *PlanXClient) SendPoolCapacityChangeWithChan(request *SendPoolCapacityChangeRequest) (<-chan *SendPoolCapacityChangeResponse, <-chan error) {
	responseChan := make(chan *SendPoolCapacityChangeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendPoolCapacityChange(request)
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

// SendPoolCapacityChangeRequest is the request struct for api SendPoolCapacityChange
type SendPoolCapacityChangeRequest struct {
	*requests.BaseRequest
	PoolId         string `position:"Body" name:"poolId" binding:"required"`            //The unique identifier of the capital pool
	SourceSymbol   string `position:"Body" name:"sourceSymbol" binding:"required"`      //Token name on the left
	SourceAmount   string `position:"Body" name:"sourceAmount" binding:"required,gt=0"` //The amount of tokens on the left
	CapacityAction string `position:"Body" name:"capacityAction" binding:"required"`    //api.PoolCapacityActionIncrease，api.PoolCapacityActionDecrease
}

// SendPoolCapacityChangeResponse is the request struct for api SendPoolCapacityChange
type SendPoolCapacityChangeResponse struct {
	*responses.BaseResponse
	Data *ResponsePool `json:"data"`
}

// CreateSendPoolCapacityChangeRequest creates a request to invoke SendPoolCapacityChange API
func CreateSendPoolCapacityChangeRequest(poolId, sourceSymbol, capacityAction string, sourceAmount string) (request *SendPoolCapacityChangeRequest) {
	request = &SendPoolCapacityChangeRequest{
		BaseRequest:    requests.NewPostRequest("/v1/api/pool/capacity/change"),
		PoolId:         poolId,
		SourceSymbol:   sourceSymbol,
		SourceAmount:   sourceAmount,
		CapacityAction: capacityAction,
	}
	return
}

// CreateSendPoolCapacityChangeResponse creates a response to parse from SendPoolCapacityChange response
func CreateSendPoolCapacityChangeResponse() (response *SendPoolCapacityChangeResponse) {
	response = &SendPoolCapacityChangeResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
