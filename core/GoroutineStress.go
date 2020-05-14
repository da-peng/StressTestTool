package core

import (
	"fmt"
	"math"
	"time"
)

// GoruntineStress Goruntine并发核心代码
// 1. testPlan 定义压力测试计划
// 2. generateTestData 定义压力测试测试数据生成方法
// 3. transaction 定义压测的事务
func GoruntineStress(testPlan []GoroutineStressTest, generateTestData func(nums int, iter int, count int) []map[string]interface{}, transaction func(userTestData map[string]interface{})) {
	// 测试计划中 配置的梯度

	for nums, grads := range testPlan {
		//思考时间
		thinkTime := grads.ThinkTime
		// 迭代节流值， 迭代间隔时间
		thinkTimeThrottle := time.Tick(time.Duration(thinkTime) * time.Second)
		// 测试时间 /用户思考时间
		a := float64(grads.DurationTime*60) / float64(grads.ThinkTime)

		// 迭代次数,  这里可以强制设置成 自己想要的值
		iterTimes := int(math.Ceil(a))

		// QPS 或 TPS
		QPS := grads.QPS

		// 压力测试各梯度 描述
		fmt.Printf("压测第[%d]梯度，开始时间[%s]\n", nums, time.Now().Format("2006-01-02-15-04-05"))
		fmt.Printf("1s内有[%d]个人操作\n", QPS)

		fmt.Printf("每个人操作间隔[%d]秒，连续操作[%d]次\n", thinkTime, iterTimes)

		// iteration 并发迭代次数
		for iter := 0; iter < iterTimes; iter++ {
			// 获取这个梯度所需要的测试数据 参数1. 第几梯度 2. 第几个迭代 3.总共需要返回多少数据
			iterTestData := generateTestData(nums, iter, QPS)

			// goroutine并发 userFlag 用户标示，用于获取每个用户的测试数据
			for userIndex := 0; userIndex < QPS; userIndex++ {

				go func(iterTestData []map[string]interface{}, userIndex int) {
					userTestData := iterTestData[userIndex]
					// 单个接口 或 事务
					// 如定义事务 需要 注意 思考时间
					transaction(userTestData)

				}(iterTestData, userIndex)

			}
			//迭代间的思考时间
			<-thinkTimeThrottle
		}

		// 每个梯度时长预留10s
		if len(testPlan) > nums {
			time.Sleep(10 * time.Second)
			fmt.Printf("去往压测第[%d]梯度\n", nums+1)
		}
	}
}
