package rainbow

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

const (
	DefaultSpread = 5.0
	DefaultFreq   = 0.1
	seed          = 0
)

const (
	ColorModeTrueColor = iota
	ColorMode256
	ColorMode0
)

// Writer writes colorful output
type Writer struct {
	Output    io.Writer
	ColorMode int
	lineIdx   int
	Origin    int
	Spread    float64
	Freq      float64
}

var noColor = os.Getenv("TERM") == "dumb" || (!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))

func (w *Writer) SetSpread(spread float64) {
	w.Spread = spread
}

func (w *Writer) SetFreq(freq float64) {
	w.Freq = freq
}

func (w *Writer) SetOrigin(origin int) {
	w.Origin = origin
}

// writeRaw will write a lol'd s to the underlying writer.  It does no line
// detection.
func (w *Writer) writeRaw(s string) (int, error) {
	c, err := w.getColorer()
	if err != nil {
		return -1, err
	}
	nWritten := 0
	for _, r := range s {
		c.rainbow(w.Freq, float64(w.Origin)+float64(w.lineIdx)/w.Spread)
		_, err := w.Output.Write(c.format())
		if err != nil {
			return nWritten, err
		}
		n, err := w.Output.Write([]byte(string(r)))
		if err != nil {
			return nWritten, err
		}
		_, err = w.Output.Write(c.reset())
		if err != nil {
			return nWritten, err
		}
		nWritten += n
		w.lineIdx++
	}
	return nWritten, nil
}

// getColorer will attempt to map the defined color mode, to a colorer{}
func (w *Writer) getColorer() (colorer, error) {
	switch w.ColorMode {
	case ColorModeTrueColor:
		return newTruecolorColorer(), nil
	case ColorMode256:
		return New256Colorer(), nil
	case ColorMode0:
		return New0Colorer(), nil
	default:
		return nil, fmt.Errorf("Invalid colorer: [%d]", w.ColorMode)
	}
}

// Write will write a byte slice to the Writer
func (w *Writer) Write(p []byte) (int, error) {
	nWritten := 0
	ss := strings.Split(string(p), "\n")
	for i, s := range ss {
		// TODO: strip out pre-existing ANSI codes and expand tabs. Would be
		// great to expand tabs in a context aware way (line linux expand
		// command).

		n, err := w.writeRaw(s)
		if err != nil {
			return nWritten, err
		}
		nWritten += n

		// Increment the origin (line count) for each newline.  There is a
		// newline for every item in this array except the last one.
		if i != len(ss)-1 {
			n, err := w.Output.Write([]byte("\n"))
			if err != nil {
				return nWritten, err
			}
			nWritten += n
			w.Origin++
			w.lineIdx = 0
		}
	}
	return nWritten, nil
}

// NewWriter will return a new io.Writer with a default ColorMode of 256
func NewWriter(spread, freq float64, origin int) io.Writer {
	colorMode := ColorMode256
	if noColor {
		colorMode = ColorMode0
	}
	return &Writer{
		Output:    stdout,
		ColorMode: colorMode,
		Origin:    origin,
		Spread:    spread,
		Freq:      freq,
	}
}

// NewTruecolorWriter will return a new io.Writer with a default ColorMode of truecolor
func NewTruecolorWriter(spread, freq float64, origin int) io.Writer {
	colorMode := ColorModeTrueColor
	if noColor {
		colorMode = ColorMode0
	}
	return &Writer{
		Output:    stdout,
		ColorMode: colorMode,
		Origin:    origin,
		Spread:    spread,
		Freq:      freq,
	}
}
