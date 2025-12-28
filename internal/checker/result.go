package checker

import "time"

type result struct {
	url        string        // 要检测的 url
	statusCode int           // 响应状态码
	cost       time.Duration // 花费时间
	errMsg     string        // 失败原因
}
