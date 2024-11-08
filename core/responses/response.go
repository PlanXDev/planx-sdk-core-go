package responses

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"planx-sdk-core-go/core/errors"
	"strings"
)

type ComResponse interface {
	IsSuccess() bool
	GetHttpStatus() int
	GetHttpHeaders() map[string][]string
	GetHttpContentString() string
	GetHttpContentBytes() []byte
	GetOriginHttpResponse() *http.Response
	parseFromHttpResponse(httpResponse *http.Response) error
	GetSuccess() bool
	GetCode() string
	GetMsg() string
	GetDesc() string
}

type BaseResponse struct {
	httpStatus         int
	httpHeaders        map[string][]string
	httpContentString  string
	httpContentBytes   []byte
	originHttpResponse *http.Response
	Code               string
	Success            bool
	Msg                string
	Desc               string
}

func (baseResponse *BaseResponse) GetSuccess() bool {
	return baseResponse.Success
}

func (baseResponse *BaseResponse) GetCode() string {
	return baseResponse.Code
}

func (baseResponse *BaseResponse) GetMsg() string {
	return baseResponse.Msg
}

func (baseResponse *BaseResponse) GetDesc() string {
	return baseResponse.Desc
}

func (baseResponse *BaseResponse) GetHttpStatus() int {
	return baseResponse.httpStatus
}

func (baseResponse *BaseResponse) GetHttpHeaders() map[string][]string {
	return baseResponse.httpHeaders
}

func (baseResponse *BaseResponse) GetHttpContentString() string {
	return baseResponse.httpContentString
}

func (baseResponse *BaseResponse) GetHttpContentBytes() []byte {
	return baseResponse.httpContentBytes
}

func (baseResponse *BaseResponse) GetOriginHttpResponse() *http.Response {
	return baseResponse.originHttpResponse
}

func (baseResponse *BaseResponse) IsSuccess() bool {
	if baseResponse.GetHttpStatus() >= 200 && baseResponse.GetHttpStatus() < 300 {
		return true
	}

	return false
}

func (baseResponse *BaseResponse) parseFromHttpResponse(httpResponse *http.Response) (err error) {
	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}
	baseResponse.httpStatus = httpResponse.StatusCode
	baseResponse.httpHeaders = httpResponse.Header
	baseResponse.httpContentBytes = body
	baseResponse.httpContentString = string(body)
	baseResponse.originHttpResponse = httpResponse
	return
}

func (baseResponse *BaseResponse) String() string {
	resultBuilder := bytes.Buffer{}
	// statusCode
	// resultBuilder.WriteString("\n")
	resultBuilder.WriteString(fmt.Sprintf("%s %s\n", baseResponse.originHttpResponse.Proto, baseResponse.originHttpResponse.Status))
	// httpHeaders
	//resultBuilder.WriteString("Headers:\n")
	for key, value := range baseResponse.httpHeaders {
		resultBuilder.WriteString(key + ": " + strings.Join(value, ";") + "\n")
	}
	// content
	resultBuilder.WriteString("Content:" + baseResponse.httpContentString + "\n")
	return resultBuilder.String()
}

// Unmarshal object from http response body to target Response
func Unmarshal(response ComResponse, httpResponse *http.Response) (err error) {
	err = response.parseFromHttpResponse(httpResponse)
	if err != nil {
		return
	}
	if !response.IsSuccess() {
		err = errors.NewClientError(string(rune(response.GetHttpStatus())), response.GetHttpContentString(), nil)
		return
	}

	if len(response.GetHttpContentBytes()) == 0 {
		return
	}
	err = json.Unmarshal(response.GetHttpContentBytes(), response)
	if err != nil {
		err = errors.NewClientError(errors.JsonUnmarshalErrorCode, errors.JsonUnmarshalErrorMessage, err)
	}
	if !response.GetSuccess() {
		if response.GetDesc() != "" {
			err = errors.NewClientError(errors.RequestErrorCode, response.GetDesc(), err)
		} else {
			err = errors.NewClientError(errors.RequestErrorCode, response.GetMsg(), err)
		}
	}
	return
}

func NewResponse() (response *BaseResponse) {
	return &BaseResponse{}
}
