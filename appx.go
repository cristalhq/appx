package appx

import (
	"context"
	"os"
	"os/signal"
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

// DoOnSignal runs fn on every signal.
// Function is async, context is used to close underlying goroutine.
func DoOnSignal(ctx context.Context, signal os.Signal, fn func(ctx context.Context)) {
	ch := newChan(signal)

	go func() {
		for {
			select {
			case <-ch:
				fn(ctx)
			case <-ctx.Done():
				return
			}
		}
	}()
}

// newChan returns a channel triggered on every sig.
func newChan(sig os.Signal) <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, sig)
	return ch
}
