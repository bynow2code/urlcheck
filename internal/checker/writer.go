package checker

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
)

type resultWriter interface {
	write(context.Context, <-chan result) error
}

type stdoutWriter struct{}

func newStdoutWriter() *stdoutWriter {
	return &stdoutWriter{}
}

func (sw *stdoutWriter) write(ctx context.Context, results <-chan result) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case res, ok := <-results:
			if !ok {
				return nil
			}

			if res.errMsg == "" {
				fmt.Printf(" ✅ url:[%s] Status Code:[%d] Duration:[%.2f ms]\n", res.url, res.statusCode, res.cost.Seconds()*1000)
			} else {
				fmt.Printf(" ❌ url:[%s] Status Code:[%d] Duration:[%.2f ms] Error:[%s]\n", res.url, res.statusCode, res.cost.Seconds()*1000, res.errMsg)
			}
		}
	}
}

type csvWriter struct {
	path string
}

func newCSVWriter(path string) *csvWriter {
	return &csvWriter{path: path}
}

func (cw *csvWriter) write(ctx context.Context, results <-chan result) error {
	file, err := os.Create(cw.path)
	if err != nil {
		return err
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	err = w.Write([]string{"url", "Status Code", "Duration(ms)", "Error"})
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case res, ok := <-results:
			if !ok {
				return nil
			}

			err = w.Write([]string{
				res.url,
				fmt.Sprintf("%d", res.statusCode),
				fmt.Sprintf("%.2f", res.cost.Seconds()*1000),
				res.errMsg,
			})
			if err != nil {
				return err
			}
		}
	}

}
