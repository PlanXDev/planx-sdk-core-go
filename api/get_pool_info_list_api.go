package api

import (
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
)

// GetPoolInfoList Get all capital pool objects.
func (client *PlanXClient) GetPoolInfoList(request *GetPoolInfoListRequest) (response *GetPoolInfoListResponse, err error) {
	response = CreateGetPoolInfoListResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// GetPoolInfoListWithChan Get all capital pool objects.
func (client *PlanXClient) GetPoolInfoListWithChan(request *GetPoolInfoListRequest) (<-chan *GetPoolInfoListResponse, <-chan error) {
	responseChan := make(chan *GetPoolInfoListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPoolInfoList(request)
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

// GetPoolInfoListRequest is the request struct for api GetPoolInfoList
type GetPoolInfoListRequest struct {
	*requests.BaseRequest
}

// GetPoolInfoListResponse is the request struct for api GetPoolInfoList
type GetPoolInfoListResponse struct {
	*responses.BaseResponse
	Data []*ResponsePool `json:"data"`
}

// CreateGetPoolInfoListRequest creates a request to invoke GetPoolInfoList API
func CreateGetPoolInfoListRequest() (request *GetPoolInfoListRequest) {
	request = &GetPoolInfoListRequest{
		BaseRequest: requests.NewGetRequest("/v1/api/pool/info/list"),
	}
	return
}

// CreateGetPoolInfoListResponse creates a response to parse from GetPoolInfoList response
func CreateGetPoolInfoListResponse() (response *GetPoolInfoListResponse) {
	response = &GetPoolInfoListResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
