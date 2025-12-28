package checker

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type urlChecker struct {
	reader           urlReader
	concurrencyLimit int           // 最大并发数
	requestTimeout   time.Duration // 请求超时时间
	inputPath        string        // 输入文件路径
	outputPath       string        // 输出文件路径
	results          chan result
	writer           resultWriter
	wg               sync.WaitGroup
	ctx              context.Context
	cancel           context.CancelFunc
}

type Option func(*urlChecker) error

// WithConcurrencyLimit 设置最大并发数
func WithConcurrencyLimit(concurrencyLimit int) Option {
	return func(c *urlChecker) error {
		c.concurrencyLimit = concurrencyLimit
		return nil
	}
}

// WithRequestTimeout 设置请求超时时间
func WithRequestTimeout(requestTimeout time.Duration) Option {
	return func(c *urlChecker) error {
		c.requestTimeout = requestTimeout
		return nil
	}
}

// WithInputPath 设置输入文件路径
func WithInputPath(inputPath string) Option {
	return func(c *urlChecker) error {
		if inputPath != "" {
			c.reader = newFileReader(inputPath)
			c.inputPath = inputPath
		}
		return nil
	}
}

// WithOutputPath 设置输出文件路径
func WithOutputPath(outputPath string) Option {
	return func(c *urlChecker) error {
		if outputPath != "" {
			c.writer = newCSVWriter(outputPath)
			c.outputPath = outputPath
		}
		return nil
	}
}

// RunUrlChecker 运行检查器
func RunUrlChecker(ctx context.Context, opts ...Option) error {
	checker, err := newUrlChecker(ctx, opts...)
	if err != nil {
		return err
	}
	defer checker.close()

	if err := checker.run(); err != nil {
		return err
	}

	return checker.handleResults()
}

// 创建检查器
func newUrlChecker(parentCtx context.Context, opts ...Option) (*urlChecker, error) {
	ctx, cancel := context.WithCancel(parentCtx)

	checker := &urlChecker{
		reader:  newStdinReader(),
		results: make(chan result, 10),
		writer:  newStdoutWriter(),
		ctx:     ctx,
		cancel:  cancel,
	}

	for _, opt := range opts {
		if err := opt(checker); err != nil {
			return nil, err
		}
	}

	return checker, nil
}

// 并发检测 url
func (uc *urlChecker) run() error {
	urls, err := uc.reader.read()
	if err != nil {
		return err
	}

	uc.wg.Add(1)
	go uc.runCheks(urls)

	return nil
}

// 并行检测 url 执行逻辑
func (uc *urlChecker) runCheks(urls []string) {
	defer uc.wg.Done()

	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
		close(uc.results)
	}()

	semaphore := make(chan struct{}, uc.concurrencyLimit)

	for _, url := range urls {
		select {
		case <-uc.ctx.Done():
			return
		case semaphore <- struct{}{}:
			wg.Add(1)

			go func() {
				defer func() {
					wg.Done()
					<-semaphore
				}()

				res := checkSingleURL(uc.ctx, url, uc.requestTimeout)
				uc.results <- res
			}()
		}
	}
}

// 处理结果
func (uc *urlChecker) handleResults() error {
	return uc.writer.write(uc.ctx, uc.results)
}

func (uc *urlChecker) close() {
	uc.cancel()
	uc.wg.Wait()
}

// 检查单个 url
func checkSingleURL(ctx context.Context, url string, requestTimeout time.Duration) result {
	res := result{
		url: url,
	}

	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		res.errMsg = fmt.Sprintf("failed to create request : %v", err)
		return res
	}

	client := http.Client{}
	now := time.Now()

	response, err := client.Do(request)
	if err != nil {
		res.errMsg = fmt.Sprintf("request failed : %v", err)
		return res
	}
	defer response.Body.Close()

	res.statusCode = response.StatusCode
	res.cost = time.Since(now)

	return res
}
