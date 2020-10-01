// +build tools

// This package is a placeholder to import tool related code that is not to be
// included as part of the build. It's a way to allow creating a soft
// dependency, as otherwise go mod will not import it automatically and a `go
// mod tidy` will remove it. See
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
// and
// https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md
// for more details.
package tools

import (
	_ "github.com/go-bindata/go-bindata/go-bindata"
)
