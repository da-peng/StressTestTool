

# GoroutineStressTest Goroutine并发核心 说明

```
func GoroutineStress(testPlan []GoroutineStressTest, generateTestData func(nums int, iter int, count int) []map[string]interface{}, transaction func(userTestData map[string]interface{})) 
```
## 理由
资料[Go并发编程之并发和Goroutine](https://www.jianshu.com/p/80f69dad849f)
1. 比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务


## 配置步骤

### testPlan  []GoroutineStressTest 定义压测计划
目前支持梯度
例如QPS或TPS 100/200/300 梯度压测 ，每个梯度间隔默认10s

> GoruntineStress 并发方式配置
```
//GoroutineStressTest Goruntine并发 设置不同梯度的数据
func GoroutineStressTest() []core.GoroutineStressTest {
    //梯度1
	stepOne := core.GoroutineStressTest{
		DurationTime: 1, //单位min
		QPS:          100, //100
		ThinkTime:    6, // 思考时间 单位s
	}
    //梯度2
	stepTwo := core.GoroutineStressTest{
		DurationTime: 1,
		QPS:          200,
		ThinkTime:    6,
	}

	stressTestPlan := []core.GoroutineStressTest{stepOne, stepTwo}

	return stressTestPlan
}
```
###  generateTestData func(nums int, iter int, count int) 定义测试数据生成方法,可以根据条件 每个梯度每次迭代每个用户 都生成一组测试数据
看具体需求，自定义配置

参数:
1. nums 当前压测梯度 
2. iter 当前迭代数 
3. count = QPS = TPS 单次迭代总数据

```
// generateTestData 生产 每次迭代对应的测试数据
func generateTestData(nums int, iter int, count int) []XXLLotteryGoRequestParmas {

	var requestBodys []XXLLotteryGoRequestParmas
	/**
	* 这里面的参数可以根据 1.当前压测梯度 2.当前迭代数 3. 总数据
	*/
	start := 1000
	for index := 0; index < count; index++ {
		requestBodys = append(requestBodys,
			XXLLotteryGoRequestParmas{
				MixNick:   strconv.Itoa(start),
				GameLevel: iter + 1,
				GameScore: (iter + 1) * 1000,
			})
		start++
	}
	return requestBodys
}
```

###  transaction 定义压测的事务

```
func Tranaction(userTestData map[string]interface{}) {
	// 发送请求
	//1. 测试数据
	params := utils.MapToJSON(userTestData)
	//2. 发送请求，触发事务，这里可自定义
	responseContent, responseTime := DoRequest(params)

	//3. 统计
	statistics(responseContent, responseTime)

}
```





