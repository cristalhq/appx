//go:build !debug

package appx

// IsDebugBuild reports whether compilation was with `-tags=debug` flag.
const IsDebugBuild = false
