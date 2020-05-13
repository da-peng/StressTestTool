package test

import (
	"StressTestTool/api"
	"StressTestTool/base"
	"StressTestTool/utils"
	"testing"
)

// test 后缀的文件，内的函数不能被外部引用

//APISingleRequest  单个接口请求方法  拿到API 的url header method 数据类型 等信息
func LotteryGoAPISingleRequest(requestBody []byte) int64 {

	apiInfo := api.LotteryGoApiInfo()

	preRequest := apiInfo.BuildRequest(requestBody)
	responseContent, responseTime := base.RealRequest(preRequest)

	api.Statistics(responseContent)

	return responseTime
}

// TestLotteryGoAPI 单个接口测试
func TestLotteryGoAPI(t *testing.T) {
	lotteryGoRequestBody := api.LotteryGoRequestBody{
		MixNick: "aa",
	}
	requestBody := utils.StructToJsonBytes(lotteryGoRequestBody)

	LotteryGoAPISingleRequest(requestBody)

}
