package rainbow

import (
	"fmt"
)

var w = &Writer{Output: stdout, ColorMode: ColorMode256, Spread: 1.0, Freq: 0.5}

func Println(a ...any) (n int, err error) {
	return fmt.Fprintln(w, a...)
}

func Printf(format string, a ...any) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
}
