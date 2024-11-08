package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
)

// GetPoolInfoDetail Get the details of a fund pool.
func (client *PlanXClient) GetPoolInfoDetail(request *GetPoolInfoDetailRequest) (response *GetPoolInfoDetailResponse, err error) {
	response = CreateGetPoolInfoDetailResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// GetPoolInfoDetailWithChan Get the details of a fund pool.
func (client *PlanXClient) GetPoolInfoDetailWithChan(request *GetPoolInfoDetailRequest) (<-chan *GetPoolInfoDetailResponse, <-chan error) {
	responseChan := make(chan *GetPoolInfoDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPoolInfoDetail(request)
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

// GetPoolInfoDetailRequest is the request struct for api GetPoolInfoDetail
type GetPoolInfoDetailRequest struct {
	*requests.BaseRequest
	PoolId string `position:"Body" name:"poolId" binding:"required"` //The unique identifier of the capital pool
}

// GetPoolInfoDetailResponse is the request struct for api GetPoolInfoDetail
type GetPoolInfoDetailResponse struct {
	*responses.BaseResponse
	Data *ResponsePool `json:"data"`
}

// CreateGetPoolInfoDetailRequest creates a request to invoke GetPoolInfoDetail API
func CreateGetPoolInfoDetailRequest(poolId string) (request *GetPoolInfoDetailRequest) {
	request = &GetPoolInfoDetailRequest{
		BaseRequest: requests.NewPostRequest("/v1/api/pool/info/detail"),
		PoolId:      poolId,
	}
	return
}

// CreateGetPoolInfoDetailResponse creates a response to parse from GetPoolInfoDetail response
func CreateGetPoolInfoDetailResponse() (response *GetPoolInfoDetailResponse) {
	response = &GetPoolInfoDetailResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
