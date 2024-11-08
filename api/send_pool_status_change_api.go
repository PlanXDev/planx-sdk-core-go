package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
)

// SendPoolStatusChange Modify the active status of the capital pool.
func (client *PlanXClient) SendPoolStatusChange(request *SendPoolStatusChangeRequest) (response *SendPoolStatusChangeResponse, err error) {
	response = CreateSendPoolStatusChangeResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// SendPoolStatusChangeWithChan Modify the active status of the capital pool.
func (client *PlanXClient) SendPoolStatusChangeWithChan(request *SendPoolStatusChangeRequest) (<-chan *SendPoolStatusChangeResponse, <-chan error) {
	responseChan := make(chan *SendPoolStatusChangeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendPoolStatusChange(request)
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

// SendPoolStatusChangeRequest is the request struct for api SendPoolStatusChange
type SendPoolStatusChangeRequest struct {
	*requests.BaseRequest
	PoolId       string `position:"Body" name:"poolId" binding:"required"`       //The unique identifier of the capital pool
	SourceSymbol string `position:"Body" name:"sourceSymbol" binding:"required"` //Token name on the left
	TargetSymbol string `position:"Body" name:"targetSymbol" binding:"required"` //Token name on the right
	PoolStatus   string `position:"Body" name:"poolStatus" binding:"required"`   //api.PoolStatusActive，api.PoolStatusInActive
}

// SendPoolStatusChangeResponse is the request struct for api SendPoolStatusChange
type SendPoolStatusChangeResponse struct {
	*responses.BaseResponse
	Data *ResponsePool `json:"data"`
}

// CreateSendPoolStatusChangeRequest creates a request to invoke SendPoolStatusChange API
func CreateSendPoolStatusChangeRequest(poolId, sourceSymbol, targetSymbol, poolStatus string) (request *SendPoolStatusChangeRequest) {
	request = &SendPoolStatusChangeRequest{
		BaseRequest:  requests.NewPostRequest("/v1/api/pool/status/change"),
		PoolId:       poolId,
		SourceSymbol: sourceSymbol,
		TargetSymbol: targetSymbol,
		PoolStatus:   poolStatus,
	}
	return
}

// CreateSendPoolStatusChangeResponse creates a response to parse from SendPoolStatusChange response
func CreateSendPoolStatusChangeResponse() (response *SendPoolStatusChangeResponse) {
	response = &SendPoolStatusChangeResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
