package run

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bynow2code/urlcheck/internal/checker"
)

func Run(cfg *Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	errors := make(chan error, 1)
	doneCh := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(doneCh)

		if err := checker.RunUrlChecker(
			ctx,
			checker.WithConcurrencyLimit(cfg.ConcurrencyLimit),
			checker.WithRequestTimeout(cfg.RequestTimeout),
			checker.WithInputPath(cfg.InputPath),
			checker.WithOutputPath(cfg.OutputPath),
		); err != nil {
			select {
			case errors <- err:
			default:
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigChan:
		fmt.Println("ℹ️ Received stop signal, shutting down...")
	case err := <-errors:
		fmt.Fprintf(os.Stderr, "❌ Exiting due to error: %v\n", err)
	case <-doneCh:
	}

	cancel()

	waitDone := make(chan struct{})
	go func() {
		wg.Wait()
		close(waitDone)
	}()

	select {
	case <-waitDone:
	case <-time.After(5 * time.Second):
		fmt.Println("⚠️ Shutdown timed out, forcing exit.")
	}

	return nil
}
