package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
	"github.com/shopspring/decimal"
)

// SendAccountMarketInvest Add funds to a designated token market account.
func (client *PlanXClient) SendAccountMarketInvest(request *SendAccountMarketInvestRequest) (response *SendAccountMarketInvestResponse, err error) {
	response = CreateSendAccountMarketInvestResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// SendAccountMarketInvestWithChan Add funds to a designated token market account.
func (client *PlanXClient) SendAccountMarketInvestWithChan(request *SendAccountMarketInvestRequest) (<-chan *SendAccountMarketInvestResponse, <-chan error) {
	responseChan := make(chan *SendAccountMarketInvestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendAccountMarketInvest(request)
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

// SendAccountMarketInvestRequest is the request struct for api SendAccountMarketInvest
type SendAccountMarketInvestRequest struct {
	*requests.BaseRequest
	ExternalOrderId string          `position:"Body" name:"externalOrderId" binding:"required"`   //External order ID,This parameter is used to mark this operation
	SourceSymbol    string          `position:"Body" name:"sourceSymbol" binding:"required"`      //Token name on the left
	TargetSymbol    string          `position:"Body" name:"targetSymbol" binding:"required"`      //Token name on the right
	TargetAmount    decimal.Decimal `position:"Body" name:"targetAmount" binding:"required,gt=0"` //The amount of tokens on the right
}

// SendAccountMarketInvestResponse is the request struct for api SendAccountMarketInvest
type SendAccountMarketInvestResponse struct {
	*responses.BaseResponse
	Data []*ResponseGiftSourceDetail `json:"data"`
}

// CreateSendAccountMarketInvestRequest creates a request to invoke SendAccountMarketInvest API
func CreateSendAccountMarketInvestRequest(externalOrderId, sourceSymbol, TargetSymbol string, targetAmount decimal.Decimal) (request *SendAccountMarketInvestRequest) {
	request = &SendAccountMarketInvestRequest{
		BaseRequest:     requests.NewPostRequest("/v1/api/account/market/invest"),
		ExternalOrderId: externalOrderId,
		SourceSymbol:    sourceSymbol,
		TargetSymbol:    TargetSymbol,
		TargetAmount:    targetAmount,
	}
	return
}

// CreateSendAccountMarketInvestResponse creates a response to parse from SendAccountMarketInvest response
func CreateSendAccountMarketInvestResponse() (response *SendAccountMarketInvestResponse) {
	response = &SendAccountMarketInvestResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
