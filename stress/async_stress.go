package stress

import (
	"StressTestTool/config"
	"fmt"
	"math"
	"time"
)

// AsyncStress 异步压测
func AsyncStress(confs []config.AsyncStressConfig, requestMethod func(int, int)) {

	for _, conf := range confs {
		// 迭代节流值， 迭代间隔时间
		groutineThinkTimeThrottle := time.Tick(time.Duration(conf.ThinkTime) * time.Second)
		// 迭代次数
		a := float64(conf.DurationTime*60) / float64(conf.ThinkTime)
		groutineTimes := int(math.Ceil(a))

		thinkTime := conf.ThinkTime
		qps := conf.QPS

		fmt.Printf("1s内有[%d]个人操作\n", qps)
		fmt.Printf("每个人间隔[%d]秒，连续操作[%d]次\n", thinkTime, groutineTimes)
		// 开始迭代
		for i := 0; i < groutineTimes; i++ {
			fmt.Printf("测试开始时间[%s]\n", time.Now().Format("2006-01-02-15-04-05"))
			//开始协程并发
			for index := 0; index < qps; index++ {
				// 异步接口调用方法
				go requestMethod(index, i)
			}
			<-groutineThinkTimeThrottle
		}
		if len(confs) > 1 {
			time.Sleep(10 * time.Second)
		}
	}
}
