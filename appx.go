package appx

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

var startTime = time.Now()

// Context of the application listening SIGINT and SIGTERM signals.
func Context() context.Context {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	return ctx
}

// Uptime of the application.
func Uptime() time.Duration {
	return time.Since(startTime)
}

// DoGCWhen one of the given signal happens (GC means runtime.GC()).
// Function is async, context is used to close underlying goroutine.
// In most cases appx.DoGCWhen(os.SIGUSR1) should be enough.
func DoGCWhen(ctx context.Context, signals ...os.Signal) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ch:
				runtime.GC()
			}
		}
	}()
}
