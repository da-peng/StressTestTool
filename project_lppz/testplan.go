package lppz

import (
	"StressTestTool/core"
)

//GoruntineStress Goruntine并发 设置不同梯度的数据
func GoroutineStressTest() []core.GoroutineStressTest {

	stepOne := core.GoroutineStressTest{
		DurationTime: 1,
		QPS:          100,
		ThinkTime:    6,
	}

	stepTwo := core.GoroutineStressTest{
		DurationTime: 1,
		QPS:          100,
		ThinkTime:    6,
	}
	
	stressTestPlan := [stepOne,stepTwo]
	return stressTestPlan
}

//LoopStressTest 利用同步的方式  实现压测试场景
func LoopStressTest() []core.LoopStressTest {
	//
	stepOne := core.LoopStressTest {
		DurationTime: 10,
		ThinkTime:    10,
		LoopNums:  10,
	}

	stepTwo := core.LoopStressTest {
		DurationTime: 10,
		ThinkTime:    10,
		LoopNums:  10,
	}

	stressTestPlan := [stepOne,stepTwo]
	return stressTestPlan
}
