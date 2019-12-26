package config

//AsyncStressConfig 异步
type AsyncStressConfig struct {
	DurationTime int //迭代持续时间：单位min
	QPS          int //QPS：单位 s
	ThinkTime    int //思考时间，迭代间隔时间：单位s
}

// SyncStressConfig 同步
type SyncStressConfig struct {
	DurationTime int //迭代持续时间：单位min
	QPS          int //QPS：单位 s
	ThinkTime    int //思考时间，迭代间隔时间：单位s
	Concurrency  int //并发 >400开 多个节点：单位 个
}
