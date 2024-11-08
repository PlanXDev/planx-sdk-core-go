package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
)

// SendPoolFundDestroy Destroys the fund pool and returns the available amount to the base account.
func (client *PlanXClient) SendPoolFundDestroy(request *SendPoolFundDestroyRequest) (response *SendPoolFundDestroyResponse, err error) {
	response = CreateSendPoolFundDestroyResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// SendPoolFundDestroyWithChan Destroys the fund pool and returns the available amount to the base account.
func (client *PlanXClient) SendPoolFundDestroyWithChan(request *SendPoolFundDestroyRequest) (<-chan *SendPoolFundDestroyResponse, <-chan error) {
	responseChan := make(chan *SendPoolFundDestroyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendPoolFundDestroy(request)
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

// SendPoolFundDestroyRequest is the request struct for api SendPoolFundDestroy
type SendPoolFundDestroyRequest struct {
	*requests.BaseRequest
	PoolId       string `position:"Body" name:"poolId" binding:"required"`       //The unique identifier of the capital pool
	SourceSymbol string `position:"Body" name:"sourceSymbol" binding:"required"` //Token name on the left
	TargetSymbol string `position:"Body" name:"targetSymbol" binding:"required"` //Token name on the right
}

// SendPoolFundDestroyResponse is the request struct for api SendPoolFundDestroy
type SendPoolFundDestroyResponse struct {
	*responses.BaseResponse
	Data *ResponsePool `json:"data"`
}

// CreateSendPoolFundDestroyRequest creates a request to invoke SendPoolFundDestroy API
func CreateSendPoolFundDestroyRequest(poolId, sourceSymbol, targetSymbol string) (request *SendPoolFundDestroyRequest) {
	request = &SendPoolFundDestroyRequest{
		BaseRequest:  requests.NewPostRequest("/v1/api/pool/fund/destroy"),
		PoolId:       poolId,
		SourceSymbol: sourceSymbol,
		TargetSymbol: targetSymbol,
	}
	return
}

// CreateSendPoolFundDestroyResponse creates a response to parse from SendPoolFundDestroy response
func CreateSendPoolFundDestroyResponse() (response *SendPoolFundDestroyResponse) {
	response = &SendPoolFundDestroyResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
