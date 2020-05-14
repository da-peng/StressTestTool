package http

import (
	"net/http"
	"net/http/httptrace"
	"time"
)

// HTTPTraceTime struct  include DNSTime ConnectTime ResponseTime
type HTTPTraceTime struct {
	DNSTime      int64
	ConnectTime  int64
	ResponseTime int64
}

//TraceTime 全局变量
var TraceTime HTTPTraceTime

func httpTrace() (*http.Client, *httptrace.ClientTrace) {
	var dnsStart, connectStart, reqStart time.Time

	httpTrace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStart = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			// 从DNS 开始到完成 耗时
			TraceTime.DNSTime = time.Since(dnsStart).Milliseconds()
		},
		GetConn: func(hostPort string) {
			connectStart = time.Now()
		},

		GotConn: func(info httptrace.GotConnInfo) {
			// 从connectStart 开始建立连接 到 完成 耗时
			TraceTime.ConnectTime = time.Since(connectStart).Milliseconds()
			//func Since(t Time) Duration //表示自从t时刻以后过了多长时间，是一个时间段，相当于time.Now().Sub(t)
		},

		WroteRequest: func(info httptrace.WroteRequestInfo) {
			reqStart = time.Now()
		},

		GotFirstResponseByte: func() {
			// 从请求开始 到 响应 的耗时
			TraceTime.ResponseTime = time.Now().Sub(reqStart).Milliseconds()
		},
	}

	transport := &http.Transport{
		DisableKeepAlives: true,
	}
	// 配置 httpClient 超时 1min
	httpClient := &http.Client{Transport: transport, Timeout: time.Duration(60 * time.Second)}

	return httpClient, httpTrace
}
