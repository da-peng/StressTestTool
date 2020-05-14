package http

import (
	"StressTestTool/utils"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"strconv"
	"strings"
)

// APIInfo  1. 接口请求方法 2.接口url 3. 请求数据格式 JSON/Form 4. 请求头
type APIInfo struct {
	Method   string
	URL      string
	DataType string
	Headers  map[string]string
}

// BuildRequest  创建request 实例 1.params 请求参数
func (api *APIInfo) BuildRequest(params []byte) *http.Request {

	headers := make(http.Header)
	// var preRequest *http.Request
	if api.DataType == "Json" {
		headers.Set("Content-Type", "application/json;charset=UTF-8")
	} else if api.DataType == "Form" {
		headers.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	reuqest, err := http.NewRequest(api.Method, api.URL, bytes.NewBuffer(params))

	if err != nil {
		utils.Errors("创建请求实例报错->" + err.Error())
		return nil
	}
	//便利Header配置 设置 headers
	for k, v := range api.Headers {
		if strings.ToLower(k) == "cookie" {
			// 此处可以 自定义加入cookie的中间处理
		}
		headers.Set(k, v)
	}
	// 请求头
	reuqest.Header = headers

	// 返回request实例
	return reuqest
}

// ErrorCount 全局变量 统计错误数量
var ErrorCount int64

// DoRequest 返回 Reponse 和 TraceTime 请求开始时间/耗时/结束时间
func DoRequest(request *http.Request) (string, int64) {

	if request != nil {
		utils.Errors("请求实例为空——>请排查代码")
	}

	ctx := context.Background()
	// 定义httpTrace
	httpClient, httpTrace := httpTrace()

	request = request.WithContext(httptrace.WithClientTrace(ctx, httpTrace))

	response, err := httpClient.Do(request)

	if err != nil {
		utils.Errors("网络异常->" + err.Error())
		// 统计错误数量,一般是网络错误
		ErrorCount++
		return "", 0
	}

	statusCode := response.StatusCode
	if statusCode != 200 {
		// 统计错误数量，例如：400，503，非200
		utils.Errors("响应状态码异常->" + strconv.Itoa(statusCode))
		ErrorCount++
	}

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		utils.Errors("响应数据转换错误->" + err.Error())
	}
	responseContent := string(responseBytes)

	fmt.Printf("响应内容:[%s]\n,响应时间:[%d]ms", responseContent, TraceTime.ResponseTime)

	return responseContent, TraceTime.ResponseTime
}
