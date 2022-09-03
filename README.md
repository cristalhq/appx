# appx

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Go library for building applications. Dramatically simple. For a CLI tool see [acmd](https://github.com/cristalhq/acmd).

## Features

* Simple API.
* Dependency-free.
* Dramatically simple.

See [GUIDE.md](https://github.com/cristalhq/appx/blob/main/GUIDE.md) for more details

## Install

Go version 1.17+

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

Also see examples: [examples_test.go](https://github.com/cristalhq/appx/blob/main/example_test.go).

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/appx/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/appx/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/appx
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/appx
[reportcard-img]: https://goreportcard.com/badge/cristalhq/appx
[reportcard-url]: https://goreportcard.com/report/cristalhq/appx
[coverage-img]: https://codecov.io/gh/cristalhq/appx/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/appx
[version-img]: https://img.shields.io/github/v/release/cristalhq/appx
[version-url]: https://github.com/cristalhq/appx/releases
