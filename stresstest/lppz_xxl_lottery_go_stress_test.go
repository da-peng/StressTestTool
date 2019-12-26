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

//LppzXXLLotteryGoAPISingleRequest  单个接口请求方法  拿到API 的url header method 数据类型 等信息
func LppzXXLLotteryGoAPISingleRequest(requestBody []byte) int64 {

	apiInfo := api.LppzXXLLotteryGoApiInfo()

	preRequest := apiInfo.BuildRequest(requestBody)
	responseContent, responseTime := base.RealRequest(preRequest)

	api.LppzXXLStatistics(responseContent)

	return responseTime
}

//createTestDatas 批量准备接口请求数据
func lppzXXLcreateMultiRequestBody(start int, count int, iterNums int) []api.LppzXXLLotteryGoRequestBody {

	var requestBodys []api.LppzXXLLotteryGoRequestBody
	flag := start
	for index := 0; index < count; index++ {
		requestBodys = append(requestBodys,
			api.LppzXXLLotteryGoRequestBody{
				MixNick:   strconv.Itoa(flag),
				GameLevel: iterNums + 1,
				GameScore: (iterNums + 1) * 7000,
			})
		flag++
	}
	return requestBodys
}

// 全局变量
// var lppzXXLRequestBodys []api.LppzXXLLotteryGoRequestBody = lppzXXLcreateMultiRequestBody(10, 100) // 创建了100个测试请求内容

var lppzXXLrequestCountStatistics int = 0
var lppzXXLresponseTimeStatistics []int64

func LppzXXLLotteryGoAPIRequestMethod(index int, iterNum int) {

	lppzlotteryGoRequestBody := lppzXXLcreateMultiRequestBody(10, 100, iterNum)[index]

	requestBody := utils.StructToJsonBytes(lppzlotteryGoRequestBody)
	// 发送请求
	responseTime := LppzXXLLotteryGoAPISingleRequest(requestBody)

	lppzXXLresponseTimeStatistics = append(lppzXXLresponseTimeStatistics, responseTime)

	lppzXXLrequestCountStatistics++
}

func TestLppzXXLScene(t *testing.T) {

	convey.Convey("游戏关卡获奖接口调用", t, func() {
		conf := config.AsyncStressScene()

		stress.AsyncStress(conf, LppzXXLLotteryGoAPIRequestMethod)
		utils.StressResponseTimeScatter(lppzXXLresponseTimeStatistics)
		fmt.Println(lppzXXLrequestCountStatistics)

		fmt.Printf("店铺优惠券 139-5[%d]良品铺子新年福袋[%d]良品铺子大坚果礼盒[%d]潮流合伙人同款卫衣[%d]潮流合伙人同款短袖[%d]潮流合伙人老夫子[%d]店铺100元无门槛优惠券[%d]", api.Lppz_xxl1, api.Lppz_xxl2, api.Lppz_xxl3, api.Lppz_xxl4, api.Lppz_xxl5, api.Lppz_xxl6, api.Lppz_xxl7)
	})
}
