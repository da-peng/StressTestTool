package lppz

import "strings"

var conditons = []string{
	"店铺优惠券 139-5",
	"良品铺子新年福袋",
	"良品铺子大坚果礼盒",
	"潮流合伙人同款卫衣",
	"潮流合伙人同款短袖",
	"潮流合伙人老夫子",
	"店铺100元无门槛优惠券",
}

// 全局变量 统计奖品 数量
var prizes = make(map[string]int)
var requestCounts int = 0
var responseTimes []int64

func isExist(content string, condition string) bool {
	return strings.Contains(content, condition)
}

// Statistics 统计
func statistics(content string, responseTime int64) {
	// 1.统计 响应时间
	responseTimes = append(responseTimes, responseTime)
	// 2.统计 请求总数
	requestCounts++

	for _, conditon := range conditons {
		if isExist(content, conditon) {
			_, ok := prizes[conditon]
			if ok {
				prizes[conditon]++
			} else {
				prizes[conditon] = 1
			}
		}
	}
}
