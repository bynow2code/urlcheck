package run

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bynow2code/urlcheck/internal/checker"
)

func Run(cfg *Config) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errors := make(chan error, 1)
	done := make(chan struct{})

	go func() {
		defer close(done)

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

	var runErr error
	select {
	case <-ctx.Done():
		fmt.Println("ℹ️ Received stop signal, shutting down...")
	case runErr = <-errors:
		fmt.Fprintf(os.Stderr, "❌ Exiting due to error: %v\n", runErr)
		stop()
	case <-done:
		return nil
	}

	select {
	case <-done:
		if runErr == nil {
			fmt.Println("✅ Graceful shutdown completed.")
		}
	case <-time.After(5 * time.Second):
		fmt.Println("⚠️ Shutdown timed out, forcing exit.")
	}

	return nil
}
