package main

//AsyncStressScene 利用协程异步的方式  实现压测试场景
func AsyncStressScene() []AsyncStressConfig {

	conf := []AsyncStressConfig{
		AsyncStressConfig{
			DurationTime: 1,
			QPS:          100,
			ThinkTime:    6,
		},
	}
	return conf
}

//SyncStressScene 利用同步的方式  实现压测试场景
func SyncStressScene() []SyncStressConfig {

	conf := []SyncStressConfig{
		SyncStressConfig{
			DurationTime: 10,
			QPS:          10,
			ThinkTime:    10,
			Concurrency:  10,
		},
		SyncStressConfig{
			DurationTime: 10,
			QPS:          10,
			ThinkTime:    10,
			Concurrency:  10,
		},
	}
	return conf
}
