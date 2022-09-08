package appx

import (
	"context"
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
