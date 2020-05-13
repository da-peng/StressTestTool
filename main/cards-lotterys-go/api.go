package main

import (
	_ "StressTestTool/"
	"strings"
)

type LotteryGoRequestBody struct {
	MixNick string `json:"mixNick"`
}

func LotteryGoApiInfo() ApiInfo {
	return base.APIInfo{
		Method:   "POST",
		URL:      "https://lppzcards-test-server.meizhidev.com/apis/" + "lppz/lotterys/go",
		DataType: "Json", //Or Form
		Headers: map[string]string{
			"token": "d3eAHi0ut9E4fCsCzpSgGprwZycr4USb",
		},
	}
}

var (
	Level7 = 0
	Level6 = 0
	Level5 = 0
	Level4 = 0
	Level3 = 0
	Level2 = 0
	Level1 = 0
	Level0 = 0
)

func Statistics(responseContent string) {
	if strings.Contains(responseContent, "潮") {
		Level1++
	} else if strings.Contains(responseContent, "流") {
		Level2++
	} else if strings.Contains(responseContent, "合") {
		Level3++
	} else if strings.Contains(responseContent, "伙") {
		Level4++
	} else if strings.Contains(responseContent, "人") {
		Level5++
	}
}
