package core

import (
	"fmt"
	"math"
	"time"
)

// GoruntineStress Goruntine并发核心代码
func GoruntineStress(testPlan []GoroutineStressTest, transaction func(int, int)) {
	// 测试计划中 配置的梯度
	for nums, grads := range testPlan {
		//思考时间
		thinkTime := grads.ThinkTime
		// 迭代节流值， 迭代间隔时间
		thinkTimeThrottle := time.Tick(time.Duration(thinkTime) * time.Second)
		// 测试时间 /用户思考时间
		a := float64(grads.DurationTime*60) / float64(grads.ThinkTime)
		// 迭代次数
		iterationTimes := int(math.Ceil(a))

		// QPS 或 TPS
		QPS := grads.QPS

		// 压力测试各梯度 描述
		fmt.Printf("压测第{%d}梯度，开始时间{%s}\n", nums, time.Now().Format("2006-01-02-15-04-05"))
		fmt.Printf("1s内有[%d]个人操作\n", QPS)
		fmt.Printf("每个人操作间隔[%d]秒，连续操作[%d]次\n", thinkTime, iterationTimes)

		// iteration 并发迭代次数
		for iteration := 0; iteration < iterationTimes; iteration++ {
			// goroutine并发 userFlag 用户标示，用于获取每个用户的测试数据
			for userFlag := 0; userFlag < QPS; userFlag++ {

				go func(userFlag int, iteration int) {
					// 单个接口 或 事务
					// 如定义事务 需要 注意 思考时间
					transaction(userFlag, iteration)

				}(userFlag, iteration)

			}
			//迭代间的思考时间
			<-thinkTimeThrottle
		}

		// 每个梯度时长预留10s
		if len(testPlan) > nums {
			time.Sleep(10 * time.Second)
			fmt.Printf("去往压测第{%d}梯度\n", nums+1)
		}
	}
}
