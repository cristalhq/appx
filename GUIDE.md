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
