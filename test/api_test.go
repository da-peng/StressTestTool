package test

import (
	"StressTestTool/http"
	"StressTestTool/lppz"
	"StressTestTool/utils"
	"testing"
)

// test 后缀的文件，内的函数不能被外部引用

//APISingleRequest  单个接口请求方法  拿到API 的url header method 数据类型 等信息
func DoRequest(params []byte) (string, int64) {

	request := lppz.XXLLLotteryGo.BuildRequest(params)
	responseContent, responseTime := http.DoRequest(request)

	return responseContent, responseTime
}

// TestLotteryGoAPI 单个接口测试
func TestLotteryGoAPI(t *testing.T) {
	lotteryGoRequestBody := lppz.RequestParmas{
		MixNick:   "aa",
		GameLevel: 1,
		GameScore: 2,
	}
	requestBody := utils.StructToJSON(lotteryGoRequestBody)

	DoRequest(requestBody)

}
