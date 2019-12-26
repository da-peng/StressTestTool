package base

import (
	"StressTestTool/utils"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"strings"
)

// BuildRequest give requestBody parmas
func (its *APIInfo) BuildRequest(requestBody []byte) *http.Request {

	headers := make(http.Header)
	var preRequest *http.Request
	if its.DataType == "Json" {
		headers.Set("Content-Type", "application/json;charset=UTF-8")
	} else if its.DataType == "Form" {
		headers.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	fmt.Printf("请求Body:[%s]\n", requestBody)
	preRequest, err := http.NewRequest(its.Method, its.URL, bytes.NewBuffer(requestBody))

	if err != nil {
		utils.Errors(err.Error())
		return nil
	}
	for k, v := range its.Headers {
		if strings.ToLower(k) == "cookie" {

		}
		headers.Set(k, v)
	}

	preRequest.Header = headers
	return preRequest
}

var errorCount int64

// RealRequest return Reponse and TraceTime
func RealRequest(preRequest *http.Request) (string, int64) {
	if preRequest == nil {
		utils.Errors("error")
	}
	ctx := context.Background()
	httpClient, httpTrace := httpTrace()

	preRequest = preRequest.WithContext(httptrace.WithClientTrace(ctx, httpTrace))

	Response, err := httpClient.Do(preRequest)

	if err != nil {
		utils.Errors(err.Error())
		errorCount++
		return "", 0
	} else {
		statusCode := Response.StatusCode
		if statusCode != 200 {
			utils.Wrong("statusCode is not 200")
			errorCount++
		}
	}

	responseBytes, err := ioutil.ReadAll(Response.Body)

	if err != nil {
		utils.Errors(err.Error())
	}
	responseContent := string(responseBytes)

	fmt.Printf("响应信息:[%s]\n,响应时间:[%d]ms", responseContent, TraceTime.ResponseTime)

	return responseContent, TraceTime.ResponseTime
}
