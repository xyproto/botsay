//go:build !windows
// +build !windows

package rainbow

import "os"

var stdout = os.Stdout
