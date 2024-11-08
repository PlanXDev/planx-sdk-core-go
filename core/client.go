package core

import (
	"context"
	"fmt"
	"github.com/PlanXDev/planx-sdk-core-go/core/credential"
	"github.com/PlanXDev/planx-sdk-core-go/core/errors"
	"github.com/PlanXDev/planx-sdk-core-go/core/requests"
	"github.com/PlanXDev/planx-sdk-core-go/core/responses"
	"github.com/PlanXDev/planx-sdk-core-go/core/sign"
	"github.com/PlanXDev/planx-sdk-core-go/core/util"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var Version = "1.0.0.Final"

var defaultConnectTimeout = 5 * time.Second
var defaultReadTimeout = 10 * time.Second

var hookDo = func(fn func(req *http.Request) (*http.Response, error)) func(req *http.Request) (*http.Response, error) {
	return fn
}

type Client struct {
	logger         *Logger
	config         *Config
	readTimeout    time.Duration
	connectTimeout time.Duration
	asyncTaskQueue chan func()
	isOpenAsync    bool
	credential     *credential.SecretKeyCredential
}

func (client *Client) Init() (err error) {
	panic("not support yet")
}

// EnableAsync enable the async task queue
func (client *Client) EnableAsync(routinePoolSize, maxTaskQueueSize int) {
	if client.isOpenAsync {
		fmt.Println("warning: Please not call EnableAsync repeatedly")
		return
	}
	client.isOpenAsync = true
	client.asyncTaskQueue = make(chan func(), maxTaskQueueSize)
	for i := 0; i < routinePoolSize; i++ {
		go func() {
			for {
				task, notClosed := <-client.asyncTaskQueue
				if !notClosed {
					return
				} else {
					task()
				}
			}
		}()
	}
}

func (client *Client) AddAsyncTask(task func()) (err error) {
	if client.asyncTaskQueue != nil {
		if client.isOpenAsync {
			client.asyncTaskQueue <- task
		}
	} else {
		err = errors.NewClientError(errors.AsyncFunctionNotEnabledCode, errors.AsyncFunctionNotEnabledMessage, nil)
	}
	return
}

func (client *Client) InitWithOptions(config *Config, credential *credential.SecretKeyCredential) (err error) {

	client.config = config

	client.credential = credential

	if config.EnableAsync {
		client.EnableAsync(config.GoRoutinePoolSize, config.MaxTaskQueueSize)
	}

	return
}

func (client *Client) getTimeout(hClient *http.Client, request requests.ComRequest) (time.Duration, time.Duration) {
	readTimeout := defaultReadTimeout
	connectTimeout := defaultConnectTimeout

	reqReadTimeout := request.GetReadTimeout()
	reqConnectTimeout := request.GetConnectTimeout()
	if reqReadTimeout != 0*time.Millisecond {
		readTimeout = reqReadTimeout
	} else if client.readTimeout != 0*time.Millisecond {
		readTimeout = client.readTimeout
	} else if hClient.Timeout != 0 {
		readTimeout = hClient.Timeout
	}

	if reqConnectTimeout != 0*time.Millisecond {
		connectTimeout = reqConnectTimeout
	} else if client.connectTimeout != 0*time.Millisecond {
		connectTimeout = client.connectTimeout
	}
	return readTimeout, connectTimeout
}

func (client *Client) newClientTimeout(request requests.ComRequest) *http.Client {
	hClient := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:           50,
			IdleConnTimeout:        60 * time.Second,
			TLSHandshakeTimeout:    5 * time.Second,
			ExpectContinueTimeout:  1 * time.Second,
			MaxResponseHeaderBytes: 5 * 1024,
		},
	}
	readTimeout, connectTimeout := client.getTimeout(hClient, request)
	hClient.Timeout = readTimeout
	if trans, ok := hClient.Transport.(*http.Transport); ok && trans != nil {
		trans.DialContext = timeout(connectTimeout)
		hClient.Transport = trans
	} else if hClient.Transport == nil {
		hClient.Transport = &http.Transport{
			DialContext: timeout(connectTimeout),
		}
	}
	return hClient
}

func (client *Client) buildRequestWithSign(request requests.ComRequest) (httpRequest *http.Request, err error) {
	// init request params
	err = requests.InitParams(request)
	if err != nil {
		return
	}
	request.GetHeaders()[requests.AppIDName] = client.credential.AppId
	request.GetHeaders()[requests.NonceName] = util.RandomString(16)
	httpRequest, err = buildHttpRequest(request, sign.NewSign(client.credential.SecretKey), client.credential.PathUrl)
	return
}

func (client *Client) DoActionWithSign(request requests.ComRequest, response responses.ComResponse) (err error) {
	fieldMap := make(map[string]string)
	initLogMsg(fieldMap)
	defer func() {
		client.printLog(fieldMap, err)
	}()
	httpRequest, err := client.buildRequestWithSign(request)
	if err != nil {
		return
	}
	var httpResponse *http.Response
	for retryTimes := 0; retryTimes <= client.config.MaxRetryTime; retryTimes++ {
		if retryTimes > 0 {
			client.printLog(fieldMap, err)
			initLogMsg(fieldMap)
		}
		putMsgToMap(fieldMap, httpRequest)
		startTime := time.Now()
		fieldMap["{start_time}"] = startTime.Format("2006-01-02 15:04:05")
		httpResponse, err = hookDo(client.newClientTimeout(request).Do)(httpRequest)
		fieldMap["{cost}"] = time.Since(startTime).String()
		if err == nil {
			fieldMap["{code}"] = strconv.Itoa(httpResponse.StatusCode)
			fieldMap["{res_headers}"] = TransToString(httpResponse.Header)
		}
		// receive error
		if err != nil {
			if !client.config.AutoRetry {
				return
			} else if retryTimes >= client.config.MaxRetryTime {
				// timeout but reached the max retry times, return
				times := strconv.Itoa(retryTimes + 1)
				timeoutErrorMsg := fmt.Sprintf(errors.TimeoutErrorMessage, times, times)
				if strings.Contains(err.Error(), "Client.Timeout") {
					timeoutErrorMsg += " Read timeout. Please set a valid ReadTimeout."
				} else {
					timeoutErrorMsg += " Connect timeout. Please set a valid ConnectTimeout."
				}
				err = errors.NewClientError(errors.TimeoutErrorCode, timeoutErrorMsg, err)
				return
			}
		}
		if isCertificateError(err) {
			return
		}

		//  if status code >= 500 or timeout, will trigger retry
		if client.config.AutoRetry && (err != nil || isServerError(httpResponse)) {
			// rewrite signatureNonce and signature
			httpRequest, err = client.buildRequestWithSign(request)
			// buildHttpRequest(request, finalSigner, regionId)
			if err != nil {
				return
			}
			continue
		}
		break
	}
	err = responses.Unmarshal(response, httpResponse)
	fieldMap["{res_body}"] = response.GetHttpContentString()
	return
}

func (client *Client) Shutdown() {
	if client.asyncTaskQueue != nil {
		close(client.asyncTaskQueue)
	}

	client.isOpenAsync = false
}

func timeout(connectTimeout time.Duration) func(cxt context.Context, net, addr string) (c net.Conn, err error) {
	return func(ctx context.Context, network, address string) (net.Conn, error) {
		return (&net.Dialer{
			Timeout:   connectTimeout,
			DualStack: true,
		}).DialContext(ctx, network, address)
	}
}

func buildHttpRequest(request requests.ComRequest, singer sign.PlanXSign, baseUrl string) (httpRequest *http.Request, err error) {
	if len(request.GetFormParams()) > 0 {
		// 转map准备验证签名
		json2Map := util.ReflectStruct2Map(request.GetFormParams())
		singer.ParamsOriginal = json2Map
	}
	singer.ParamsOriginal[requests.AppIDName] = request.GetHeaders()[requests.AppIDName]
	singer.ParamsOriginal[requests.NonceName] = request.GetHeaders()[requests.NonceName]
	singer.Sign()
	request.GetHeaders()[requests.SignName] = singer.SignStr
	requestMethod := request.GetMethod()
	requestUrl := baseUrl + request.GetBuildUrl()
	delete(request.GetFormParams(), requests.AppIDName)
	delete(request.GetFormParams(), requests.NonceName)
	body := request.GetBodyReader()
	httpRequest, err = http.NewRequest(requestMethod, requestUrl, body)
	if err != nil {
		return
	}
	for key, value := range request.GetHeaders() {
		httpRequest.Header[key] = []string{value}
	}
	return
}

func putMsgToMap(fieldMap map[string]string, request *http.Request) {
	fieldMap["{host}"] = request.Host
	fieldMap["{method}"] = request.Method
	fieldMap["{uri}"] = request.URL.RequestURI()
	fieldMap["{pid}"] = strconv.Itoa(os.Getpid())
	fieldMap["{version}"] = strings.Split(request.Proto, "/")[1]
	hostname, _ := os.Hostname()
	fieldMap["{hostname}"] = hostname
	fieldMap["{req_headers}"] = TransToString(request.Header)
	fieldMap["{target}"] = request.URL.Path + request.URL.RawQuery
}

func isServerError(httpResponse *http.Response) bool {
	return httpResponse.StatusCode >= http.StatusInternalServerError
}

func isCertificateError(err error) bool {
	if err != nil && strings.Contains(err.Error(), "x509: certificate signed by unknown authority") {
		return true
	}
	return false
}
