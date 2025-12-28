package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bynow2code/urlcheck/internal/run"
)

type CheckResult struct {
	Url    string        // 要检测的 url
	Code   int           // 响应状态码
	Cost   time.Duration // 花费时间
	ErrMsg string        // 失败原因，空则成功
}

func main() {
	cfg, err := run.ParseFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Invalid command-line arguments: %v\n", err)
		os.Exit(1)
	}

	if err := run.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Exiting due to error: %v\n", err)
		os.Exit(1)
	}
}
