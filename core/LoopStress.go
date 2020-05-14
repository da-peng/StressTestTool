package core

import (
	"fmt"
	"math"
	"time"
)

// LoopStress for轮训核心代码
func LoopStress(testPlan []LoopStressTest, transaction func(int, int)) {

	for nums, grads := range testPlan {
		//思考时间
		thinkTime := grads.ThinkTime
		// （前1s）与（后1s）的间隔时间思考时间
		thinkTimeThrottle := time.Tick(time.Duration(grads.ThinkTime) * time.Second)
		// 迭代次数
		iterationTimes := int(math.Ceil(float64(grads.DurationTime*60) / float64(grads.ThinkTime)))
		// 并发数
		loopNums := grads.LoopNums

		fmt.Printf("压测第{%d}梯度，开始时间{%s}\n", nums, time.Now().Format("2006-01-02-15-04-05"))
		fmt.Printf("1s内有[%d]个人操作\n", loopNums)
		fmt.Printf("每个人操作间隔[%d]秒，连续操作[%d]次\n", thinkTime, iterationTimes)

		for iteration := 0; iteration < iterationTimes; iteration++ {
			for userFlag := 0; userFlag < loopNums; userFlag++ {
				//同步接口调用方法
				transaction(userFlag, iteration)
			}
			<-thinkTimeThrottle
		}
	}

}
