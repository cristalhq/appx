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
	appEnv     string
	appEnvOnce sync.Once
)

// Env where application is running. Can be set via SetEnv func.
func Env() string { return appEnv }

// SetEnv for the application from environment variable or a default.
// Can be called once, all next calls panics.
// Example:
//
//	appx.SetEnv("ENV", "dev") // set from ENV or set default "dev"
//	appx.SetEnv("ENV", "") // set from ENV or leave unset
//	appx.SetEnv("", "dev") // set "dev" directly
//
// This function should be called at the start of the application.
func SetEnv(env, def string) {
	if appEnv != "" {
		panic("appx: SetEnv called twice")
	}

	appEnvOnce.Do(func() {
		v := os.Getenv(env)
		if v == "" {
			v = def
		}
		appEnv = v
	})
}

// OnSignal run fn.
// Function is async, context is used to close underlying goroutine.
func OnSignal(ctx context.Context, sig os.Signal, fn func(ctx context.Context)) {
	ch := newChan(sig)

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

// SelfSignal sends signal to the current process.
func SelfSignal(sig os.Signal) error {
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		return fmt.Errorf("unable to find own pid: %w", err)
	}

	if err := proc.Signal(sig); err != nil {
		return fmt.Errorf("signal: %w", err)
	}
	return nil
}
