package api

import (
	"planx-sdk-core-go/core/requests"
	"planx-sdk-core-go/core/responses"
)

// GetGiftClaimInfoBatch Batch query the sub-gift packs details corresponding to the superior gift pack code.
func (client *PlanXClient) GetGiftClaimInfoBatch(request *GetGiftClaimInfoBatchRequest) (response *GetGiftClaimInfoBatchResponse, err error) {
	response = CreateGetGiftClaimInfoBatchResponse()
	err = client.DoActionWithSign(request, response)
	return
}

// GetGiftClaimInfoBatchWithChan Batch query the sub-gift packs details corresponding to the superior gift pack code.
func (client *PlanXClient) GetGiftClaimInfoBatchWithChan(request *GetGiftClaimInfoBatchRequest) (<-chan *GetGiftClaimInfoBatchResponse, <-chan error) {
	responseChan := make(chan *GetGiftClaimInfoBatchResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGiftClaimInfoBatch(request)
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

// GetGiftClaimInfoBatchRequest is the request struct for api GetGiftClaimInfoBatch
type GetGiftClaimInfoBatchRequest struct {
	*requests.BaseRequest
	GiftQrCode       []string `position:"Body" name:"giftQrCode"`       //The unique CODE of the gift source
	GiftSourceQrCode string   `position:"Body" name:"giftSourceQrCode"` //The unique CODE of the gift
}

// GetGiftClaimInfoBatchResponse is the request struct for api GetGiftClaimInfoBatch
type GetGiftClaimInfoBatchResponse struct {
	*responses.BaseResponse
	Data []*ResponseGift `json:"data"`
}

// CreateGetGiftClaimInfoBatchRequest creates a request to invoke GetGiftClaimInfoBatch API
func CreateGetGiftClaimInfoBatchRequest(giftQrCode []string, giftSourceQrCode string) (request *GetGiftClaimInfoBatchRequest) {
	request = &GetGiftClaimInfoBatchRequest{
		BaseRequest:      requests.NewPostRequest("/v1/api/gift/claim/info/batch"),
		GiftQrCode:       giftQrCode,
		GiftSourceQrCode: giftSourceQrCode,
	}
	return
}

// CreateGetGiftClaimInfoBatchResponse creates a response to parse from GetGiftClaimInfoBatch response
func CreateGetGiftClaimInfoBatchResponse() (response *GetGiftClaimInfoBatchResponse) {
	response = &GetGiftClaimInfoBatchResponse{
		BaseResponse: responses.NewResponse(),
	}
	return
}
