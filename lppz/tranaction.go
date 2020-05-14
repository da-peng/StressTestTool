package lppz

import (
	"StressTestTool/http"
	"StressTestTool/utils"
)

//DoRequest 请求
func DoRequest(params []byte) (string, int64) {

	request := XXLLLotteryGo.BuildRequest(params)
	responseContent, responseTime := http.DoRequest(request)

	return responseContent, responseTime
}

//Tranaction 事务定义
func Tranaction(userTestData map[string]interface{}) {
	// 发送请求
	//1. 测试数据
	params := utils.MapToJSON(userTestData)
	//2. 发送请求，触发事务，这里可自定义
	responseContent, responseTime := DoRequest(params)

	//3. 统计
	statistics(responseContent, responseTime)

}
