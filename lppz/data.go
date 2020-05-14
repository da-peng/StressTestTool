package lppz

import (
	"StressTestTool/utils"
	"strconv"
)

// generateTestData 生产 每次迭代对应的测试数据
func generateTestData(nums int, iter int, count int) []RequestParmas {

	var requestParmas []RequestParmas
	/**
	* 这里面的参数可以根据 1.当前压测梯度 2.当前迭代数 3. 总数据
	 */
	start := 1000
	for index := 0; index < count; index++ {
		requestParmas = append(requestParmas,
			RequestParmas{
				MixNick:   strconv.Itoa(start),
				GameLevel: iter + 1,
				GameScore: (iter + 1) * 1000,
			})
		start++
	}

	params := utils.StructToMap(requestParmas)
	return params
}
