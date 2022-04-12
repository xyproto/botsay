//go:build windows
// +build windows

package rainbow

import colorable "github.com/mattn/go-colorable"

var stdout = colorable.NewColorableStdout()
