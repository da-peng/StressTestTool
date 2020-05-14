package lppz

import (
	"StressTestTool/core"
	"StressTestTool/utils"
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 消消乐压力测试
func TestXXLStressTest(t *testing.T) {

	convey.Convey("游戏关卡获奖接口调用", t, func() {
		//1.testPlan
		testPlan := GoroutineStressTest()
		//2.运行压测
		core.GoruntineStress(testPlan, generateTestData, Tranaction)

		//3.RT数据绘图
		utils.StatisticsOfRT(responseTimes)

		//4.打印抽奖情况
		fmt.Printf("总计抽奖:{%d}次", requestCounts)

		for k, v := range prizes {
			fmt.Printf("奖品:[%s]->[%d]个", k, v)
		}

	})
}
