package core

import (
	"fmt"
	"math"
	"time"
)

// LoopStress 正常轮训的 
func LoopStress(confs []LoopStressTest, requestMethod func(int, int)) {
	var throttle <-chan time.Time

	for _, conf := range confs {
		// 1000ms内每个请求间隔
		requestDuration := 1000 / conf.QPS
		// 节流值,1000ms 内请求间隔时间
		throttle = time.Tick(time.Duration(requestDuration) * time.Millisecond)

		// （前1s）与（后1s）的间隔时间思考时间
		thinkTimeThrottle := time.Tick(time.Duration(conf.ThinkTime) * time.Second)
		// 迭代次数
		iterNums := int(math.Ceil(float64(conf.DurationTime*60) / float64(conf.ThinkTime)))
		// 并发数
		concurrency := conf.Concurrency

		fmt.Printf("1s内有[%d]个人操作\n", concurrency)
		fmt.Printf("每个人间隔[%d]秒，连续操作[%d]次\n", conf.ThinkTime, iterNums)

		for i := 0; i < iterNums; i++ {
			for index := 0; index < concurrency; index++ {
				//同步接口调用方法
				requestMethod(index, i)
				<-throttle
			}
			<-thinkTimeThrottle
		}
	}

}
