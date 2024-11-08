package api

import (
	"github.com/shopspring/decimal"
	"planx-sdk-core-go/core/requests"
	"planx-sdk-core-go/core/responses"
)

// SendPoolInfoNew Transfer an equal amount of tokens from the base account to create a new funding pool.
func (client *PlanXClient) SendPoolInfoNew(request *SendPoolInfoNewRequest) (response *SendPoolInfoNewResponse, err error) {
	response = CreateSendPoolInfoNewResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// SendPoolInfoNewWithChan Transfer an equal amount of tokens from the base account to create a new funding pool.
func (client *PlanXClient) SendPoolInfoNewWithChan(request *SendPoolInfoNewRequest) (<-chan *SendPoolInfoNewResponse, <-chan error) {
	responseChan := make(chan *SendPoolInfoNewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendPoolInfoNew(request)
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

// SendPoolInfoNewRequest is the request struct for api SendPoolInfoNew
type SendPoolInfoNewRequest struct {
	*requests.BaseRequest
	SourceSymbol string          `position:"Body" name:"sourceSymbol" binding:"required"`      //Token name on the left
	TargetSymbol string          `position:"Body" name:"targetSymbol" binding:"required"`      //Token name on the right
	SourceAmount decimal.Decimal `position:"Body" name:"sourceAmount" binding:"required,gt=0"` //The amount of tokens on the left
	TargetAmount decimal.Decimal `position:"Body" name:"targetAmount" binding:"required,gt=0"` //The amount of tokens on the right
}

// SendPoolInfoNewResponse is the request struct for api SendPoolInfoNew
type SendPoolInfoNewResponse struct {
	*responses.BaseResponse
	Data *ResponsePool `json:"data"`
}

// CreateSendPoolInfoNewRequest creates a request to invoke SendPoolInfoNew API
func CreateSendPoolInfoNewRequest(sourceSymbol, targetSymbol string, sourceAmount, targetAmount decimal.Decimal) (request *SendPoolInfoNewRequest) {
	request = &SendPoolInfoNewRequest{
		BaseRequest:  requests.NewPostRequest("/v1/api/pool/info/new"),
		SourceSymbol: sourceSymbol,
		TargetSymbol: targetSymbol,
		SourceAmount: sourceAmount,
		TargetAmount: targetAmount,
	}
	return
}

// CreateSendPoolInfoNewResponse creates a response to parse from SendPoolInfoNew response
func CreateSendPoolInfoNewResponse() (response *SendPoolInfoNewResponse) {
	response = &SendPoolInfoNewResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
