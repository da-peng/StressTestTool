package lppz

import "StressTestTool/http"

// XXLLotteryGoRequestParmas 请求结构体
type RequestParmas struct {
	MixNick   string `json:"mixNick"`
	GameLevel int    `json:"gameLevel"`
	GameScore int    `json:"gameScore"`
}

// XXLLLotteryGoAPIInfo  接口信息
var XXLLLotteryGo = http.APIInfo{
	Method:   "POST",
	URL:      "https://lppzxxl-test.meizhidev.com/apis/" + "lppz/lotterys/go",
	DataType: "Json", //Or Form
	Headers: map[string]string{
		"token": "d3eAHi0ut9E4fCsCzpSgGprwZycr4USb",
	},
}
