package checker

import (
	"flag"
	"os"
)

type urlReader interface {
	read() ([]string, error)
}

type stdinReader struct {
}

func newStdinReader() *stdinReader {
	return &stdinReader{}
}

func (sr *stdinReader) read() ([]string, error) {
	return flag.Args(), nil
}

type fileReader struct {
	path string
}

func newFileReader(path string) *fileReader {
	return &fileReader{path: path}
}

func (fr *fileReader) read() ([]string, error) {
	bytes, err := os.ReadFile(fr.path)
	if err != nil {
		return nil, err
	}

	var urls []string
	start := 0
	content := string(bytes)

	for i := 0; i < len(content); i++ {
		if content[i] == '\n' {
			line := content[start:i]
			// 处理Windows换行符(\r\n)
			if len(line) > 0 && line[len(line)-1] == '\r' {
				line = line[:len(line)-1]
			}
			if len(line) > 0 {
				urls = append(urls, line)
			}
			start = i + 1
		}
	}

	// 处理最后一行（文件末尾没有换行符的情况）
	if start < len(content) {
		line := content[start:]
		if len(line) > 0 && line[len(line)-1] == '\r' {
			line = line[:len(line)-1]
		}
		if len(line) > 0 {
			urls = append(urls, line)
		}
	}

	return urls, nil
}
