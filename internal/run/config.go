package run

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Urls             []string      // 要检测的 url
	ConcurrencyLimit int           // 最大并发数
	RequestTimeout   time.Duration // 请求超时时间
	InputPath        string        // 输入文件路径
	OutputPath       string        // 输出文件路径
}

var version = "0.0.0-dev"

func ParseFlags() (*Config, error) {
	concurrencyLimit := flag.Int("c", 5, "Concurrency limit")
	requestTimeout := flag.Int("t", 5, "Request timeout in seconds")
	inputPath := flag.String("f", "", "Input file path, one url per line")
	outputPath := flag.String("o", "", "Output CSV file path")
	ver := flag.Bool("v", false, "Show version")

	flag.Usage = func() {
		fmt.Println("Welcome to urlcheck!")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *ver {
		println("version:", version)
		os.Exit(0)
	}

	if *inputPath == "" && len(flag.Args()) == 0 {
		return nil, fmt.Errorf("please specify an input file path or at least one url")
	}

	return &Config{
		ConcurrencyLimit: *concurrencyLimit,
		RequestTimeout:   time.Duration(*requestTimeout) * time.Second,
		InputPath:        *inputPath,
		OutputPath:       *outputPath,
	}, nil
}
