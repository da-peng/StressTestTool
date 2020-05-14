package core

// GoroutineStressTest Goruntine 协程不是并发的，而Goroutine支持并发的。因此Goroutine可以理解为一种Go语言的协程。同时它可以运行在一个或多个线程上。
type GoroutineStressTest struct {
	DurationTime int //迭代持续时间 单位 min
	QPS          int //QPS或TPS 注意定义事务 和事务操作间的思考时间 单位 s
	ThinkTime    int // 事务间的思考时间(迭代间隔时间)：单位s
}

// LoopStressTest for轮训
type LoopStressTest struct {
	DurationTime int //迭代持续时间  : 单位min
	ThinkTime    int //思考时间，迭代间隔时间: 单位s
	LoopNums     int //1次轮训次数 这里看成并发数
}
