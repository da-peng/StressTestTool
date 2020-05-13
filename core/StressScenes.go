package core

// GoroutineStress 协程不是并发的，而Goroutine支持并发的。因此Goroutine可以理解为一种Go语言的协程。同时它可以运行在一个或多个线程上。
type GoroutineStressTest struct {
	DurationTime int //迭代持续时间：单位min
	QPS          int //QPS：单位 s
	ThinkTime    int //思考时间，迭代间隔时间：单位s
}


// LoopStress 正常轮训的 
type LoopStressTest struct {
	DurationTime int //迭代持续时间：单位min
	QPS          int //QPS：单位 s
	ThinkTime    int //思考时间，迭代间隔时间：单位s
	Concurrency  int //并发 >400开 多个节点：单位 个
}
