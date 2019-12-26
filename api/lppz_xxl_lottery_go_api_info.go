package api

import (
	"StressTestTool/base"
	"strings"
)

type LppzXXLLotteryGoRequestBody struct {
	MixNick   string `json:"mixNick"`
	GameLevel int    `json:"gameLevel"`
	GameScore int    `json:"gameScore"`
}

func LppzXXLLotteryGoApiInfo() base.APIInfo {
	return base.APIInfo{
		Method:   "POST",
		URL:      "http://lppzxxl-test-server.meizhidev.com/" + "lppz/lotterys/go",
		DataType: "Json", //Or Form
		Headers: map[string]string{
			"token": "d3eAHi0ut9E4fCsCzpSgGprwZycr4USb",
		},
	}
}

var (
	Lppz_xxl7 = 0
	Lppz_xxl6 = 0
	Lppz_xxl5 = 0
	Lppz_xxl4 = 0
	Lppz_xxl3 = 0
	Lppz_xxl2 = 0
	Lppz_xxl1 = 0
	Lppz_xxl0 = 0
)

func LppzXXLStatistics(responseContent string) {
	if strings.Contains(responseContent, "店铺优惠券 139-5") {
		Lppz_xxl1++
	} else if strings.Contains(responseContent, "良品铺子新年福袋") {
		Lppz_xxl2++
	} else if strings.Contains(responseContent, "良品铺子大坚果礼盒") {
		Lppz_xxl3++
	} else if strings.Contains(responseContent, "潮流合伙人同款卫衣") {
		Lppz_xxl4++
	} else if strings.Contains(responseContent, "潮流合伙人同款短袖") {
		Lppz_xxl5++
	} else if strings.Contains(responseContent, "潮流合伙人老夫子") {
		Lppz_xxl6++
	} else if strings.Contains(responseContent, "店铺100元无门槛优惠券") {
		Lppz_xxl7++
	}
}
