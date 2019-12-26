package main

import (
	"StressTestTool/api"
	"StressTestTool/base"
	"StressTestTool/config"
	"StressTestTool/stress"
	"StressTestTool/utils"
	"fmt"
	"strconv"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

//----------

//APISingleRequest  单个接口请求方法  拿到API 的url header method 数据类型 等信息
func LotteryGoAPISingleRequest(requestBody []byte) int64 {

	apiInfo := api.LotteryGoApiInfo()

	preRequest := apiInfo.BuildRequest(requestBody)
	responseContent, responseTime := base.RealRequest(preRequest)

	api.Statistics(responseContent)

	return responseTime
}

//createTestDatas 批量准备接口请求数据
func createMultiRequestBody(start int, count int) []api.LotteryGoRequestBody {

	var requestBodys []api.LotteryGoRequestBody
	flag := start
	for index := 0; index < count; index++ {
		requestBodys = append(requestBodys,
			api.LotteryGoRequestBody{
				MixNick: strconv.Itoa(flag),
			})
		flag++
	}
	return requestBodys
}

// 全局变量
var requestBodys []api.LotteryGoRequestBody = createMultiRequestBody(10, 100) // 创建了100个测试请求内容

var requestCountStatistics int = 0
var responseTimeStatistics []int64

func LotteryGoAPIRequestMethod(index int, iterNum int) {

	lotteryGoRequestBody := requestBodys[index]

	requestBody := utils.StructToJsonBytes(lotteryGoRequestBody)
	// 发送请求
	responseTime := LotteryGoAPISingleRequest(requestBody)

	responseTimeStatistics = append(responseTimeStatistics, responseTime)

	requestCountStatistics++
}

func TestScene(t *testing.T) {

	convey.Convey("抽奖接口调用", t, func() {
		conf := config.AsyncStressScene()

		stress.AsyncStress(conf, LotteryGoAPIRequestMethod)

		fmt.Println(requestCountStatistics)
		utils.StressResponseTimeScatter(responseTimeStatistics)
		fmt.Printf("潮[%d]流[%d]合[%d]伙[%d]人[%d]", api.Level1, api.Level2, api.Level3, api.Level4, api.Level5)
	})
}
