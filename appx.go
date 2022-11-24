package appx

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
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

var (
	env     string
	envOnce sync.Once
)

// Env returns application environment set via SetEnv func.
func Env() string {
	return env
}

// SetEnv for the application. Can be called once, all next calls panics.
// This function should be called at the start of the application.
func SetEnv(v string) {
	if env != "" {
		panic("appx: SetEnv cannot be called twice")
	}

	envOnce.Do(func() { env = v })
}

// OnSignal run fn.
// Function is async, context is used to close underlying goroutine.
func OnSignal(ctx context.Context, signal os.Signal, fn func(ctx context.Context)) {
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

// BuildInfo of the app (if -buildvcs flag provided).
func BuildInfo() (revision string, at time.Time, isModified, ok bool) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", time.Time{}, false, false
	}

	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			revision = setting.Value
		case "vcs.time":
			at, _ = time.Parse(time.RFC3339, setting.Value)
		case "vcs.modified":
			isModified = setting.Value == "true"
		}
	}
	return revision, at, isModified, true
}

// SendInterrupt signal to the current process.
func SendInterrupt() error {
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		return fmt.Errorf("unable to find own pid: %w", err)
	}

	if err := proc.Signal(os.Interrupt); err != nil {
		return fmt.Errorf("signal: %w", err)
	}
	return nil
}
