package utils

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
			TraceTime.DNSTime = time.Since(dnsStart).Milliseconds() ////表示自从t时刻以后过了多长时间，是一个时间段
		},
		GetConn: func(hostPort string) {
			connectStart = time.Now()
		},

		GotConn: func(info httptrace.GotConnInfo) {
			TraceTime.ConnectTime = time.Since(connectStart).Milliseconds()
			//func Since(t Time) Duration //表示自从t时刻以后过了多长时间，是一个时间段，相当于time.Now().Sub(t)
		},

		WroteRequest: func(info httptrace.WroteRequestInfo) {
			reqStart = time.Now()
		},

		GotFirstResponseByte: func() {
			TraceTime.ResponseTime = time.Now().Sub(reqStart).Milliseconds()
		},
	}

	transport := &http.Transport{
		DisableKeepAlives: true,
	}

	httpClient := &http.Client{Transport: transport, Timeout: time.Duration(3000) * time.Second}

	return httpClient, httpTrace
}
