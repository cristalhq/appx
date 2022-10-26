# Guide for appx

## Basic case

```go
package main

import (
	"context"
	"os"

	"github.com/cristalhq/appx"
)

func main() {
	ctx := appx.Context()

	if err := run(ctx, os.Args[1:]); err != nil {
		panic(err)
	}
}

func run(ctx context.Context, args []string) error {
	// do good things
	return nil
}
```

## Manually trigger garbage collection

Might be useful when something bad happened and you don't want to restart the binary/instance/container/etc.

```go
func triggerGC(ctx context.Context) {
	appx.DoOnSignal(ctx, os.SIGUSR1, func(ctx context.Context) {
		runtime.GC()
	})
}
```

## Reload config on SIGHUP

Very popular pattern to add zero-restart configs to your application.

```go
type Config struct{
	// some fields
}

func rereadConfig(ctx context.Context) (*Config, error) {
	var cfg Config
	var err error

	appx.DoOnSignal(ctx, os.SIGHUP, func(ctx context.Context) {
		err = json.Unmarshal(&cfg)
		if err != nil{
			println("oh, error :(")
		}
	})

	return &cfg, err
}
```

Took from archived [cristalhq/sigreload](https://github.com/cristalhq/sigreload)
