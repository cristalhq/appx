# appx

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![version-img]][version-url]

Go library for building applications. Dramatically simple. For a CLI tool see [cristalhq/acmd](https://github.com/cristalhq/acmd).

## Features

* Simple API.
* Dependency-free.
* Dramatically simple.

See [these docs][pkg-url] or [GUIDE.md](https://github.com/cristalhq/appx/blob/main/GUIDE.md) for more details.

## Install

Go version 1.18+

```
go get github.com/cristalhq/appx
```

## Example

```go
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

See examples: [example_test.go](https://github.com/cristalhq/appx/blob/main/example_test.go).

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/appx/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/appx/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/appx
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/appx
[version-img]: https://img.shields.io/github/v/release/cristalhq/appx
[version-url]: https://github.com/cristalhq/appx/releases
