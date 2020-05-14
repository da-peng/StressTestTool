package lppz

import (
	"StressTestTool/http"
	"strings"
)

// LppzXXLLotteryGoRequestBody 请求结构体
type LppzXXLLotteryGoRequestBody struct {
	MixNick   string `json:"mixNick"`
	GameLevel int    `json:"gameLevel"`
	GameScore int    `json:"gameScore"`
}

// LppzXXLLotteryGoAPIInfo  接口信息
func XXLLLotteryGoAPIInfo() http.APIInfo {


	header := map[string]string{
		"token": "d3eAHi0ut9E4fCsCzpSgGprwZycr4USb",
	}

	return http.APIInfo{
		Method:   "POST",
		URL:      "https://lppzxxl-test.meizhidev.com/apis/" + "lppz/lotterys/go",
		DataType: "Json", //Or Form
		Headers: header
	}
}


// 全局变量 统计奖品 数量
var statistics = make(map[string]int)

var keys =[
	"店铺优惠券 139-5",
	"良品铺子新年福袋",
	"良品铺子大坚果礼盒",
	"潮流合伙人同款卫衣",
	"潮流合伙人同款短袖",
	"潮流合伙人老夫子"
	"店铺100元无门槛优惠券",
]

func Statistics(responseContent string) {

}
