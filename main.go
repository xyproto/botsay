package main

import (
	"fmt"
	"github.com/mattes/go-asciibot"
	"github.com/xyproto/rainbow"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

const (
	boxContentWidth = 42
	versionString   = "botsay 1.2.0"
)

// GFX is ASCII graphics as a string, and where to place it on the canvas
type GFX struct {
	ascii string
	x     int
	y     int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// New creates a new GFX struct, with an ASCII art string and a position
func New(ascii string, x, y int) *GFX {
	return &GFX{ascii, x, y}
}

// Draw an ASCII bubble
func bubble(w, h int) string {
	var sb strings.Builder
	sb.WriteString("   .")
	sb.WriteString(strings.Repeat("-", w-5))
	sb.WriteString(".\n")
	for i := 0; i < (h - 2); i++ {
		if i == 1 {
			sb.WriteString("--<|")
		} else {
			sb.WriteString("   |")
		}
		sb.WriteString(strings.Repeat(" ", w-5))
		sb.WriteString("|\n")
	}
	sb.WriteString("   '")
	sb.WriteString(strings.Repeat("-", w-5))
	sb.WriteString("'\n")
	return sb.String()
}

// Return the width and height of a given ASCII art string
func size(s string) (int, int) {
	maxWidth := 0
	maxHeight := 0
	lineCounter := 0
	for _, line := range strings.Split(s, "\n") {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
		lineCounter++
	}
	if lineCounter > maxHeight {
		maxHeight = lineCounter
	}
	return maxWidth, maxHeight
}

// Return a character at (x,y) in a multiline string.
// If anythings go wrong, or if (x,y) is out of bounds, return a space.
func get(s string, x, y int) string {
	if x < 0 || y < 0 {
		return " "
	}
	w, h := size(s)
	if x >= w || y >= h {
		return " "
	}
	for i, line := range strings.Split(s, "\n") {
		if i == y {
			if x < len(line) {
				return string(line[x])
			}
			return " "
		}
	}
	return " "
}

// Like a blit function, but for ASCII graphics. Uses " " as the "transparent pixel".
func combine(a, b string, xoffset, yoffset int) string {
	aW, aH := size(a)
	bW, bH := size(b)
	maxW := max(aW, bW+xoffset)
	maxH := max(aH, bH+yoffset)
	var sb strings.Builder
	for y := 0; y < maxH; y++ {
		for x := 0; x < maxW; x++ {
			if get(b, x-xoffset, y-yoffset) == " " {
				sb.WriteString(get(a, x, y))
			} else {
				sb.WriteString(get(b, x-xoffset, y-yoffset))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// Combine several ASCII graphics layers (with a position each) into one layer
func render(layers []*GFX) string {
	var canvas string
	for _, gfx := range layers {
		canvas = combine(canvas, gfx.ascii, gfx.x, gfx.y)
	}
	return canvas
}

// Split a string into words, keepin punctuation and trailing spaces
func splitWords(s string) []string {
	var (
		splitpoint bool
		words      []string
		letters    strings.Builder
		tmp        string
	)
	lenS := len(s)
	for i, r := range s {
		splitpoint = false
		switch r {
		case '.', '!', ',', ':', '-', ' ', '?', ';', '\n':
			// Check if the next character is not an end quote
			if i+1 < lenS && s[i+1] != '"' && s[i+1] != '\'' {
				splitpoint = true
			}
		}
		// Combine repeated dashes
		if r == '-' && i+1 < lenS && s[i+1] == '-' {
			splitpoint = false
		}
		// Combine repeated dots
		if r == '.' && i+1 < lenS && s[i+1] == '.' {
			splitpoint = false
		}
		if splitpoint || i == lenS {
			letters.WriteRune(r)
			tmp = letters.String()
			if len(tmp) > 0 {
				words = append(words, tmp)
			}
			letters.Reset()
		} else {
			letters.WriteRune(r)
		}
	}
	tmp = strings.TrimSpace(letters.String())
	if len(tmp) > 0 {
		words = append(words, tmp)
	}
	return words
}

// Split a string by words, then combine to form lines maximum w long
func splitWidthWords(s string, w int) []string {
	var sl []string
	var line string
	for _, word := range splitWords(s) {
		if len(line)+len(word) < w {
			line += word
		} else {
			trimmedLine := strings.TrimSpace(line)
			if strings.HasSuffix(trimmedLine, "--") {
				// Move the double dash to the beginning of the next line
				trimmedLine = trimmedLine[:len(trimmedLine)-2]
				sl = append(sl, trimmedLine)
				line = "-- " + word
			} else {
				sl = append(sl, trimmedLine)
				line = word
			}
		}
	}
	if len(line) > 0 {
		sl = append(sl, strings.TrimSpace(line))
	}
	return sl
}

// Generate ASCII graphics of a randomly generated bot with a speech bubble
func botsay(msg string) string {
	var layers []*GFX
	trimmed := strings.TrimSpace(msg)
	msgwidth := boxContentWidth
	lineCount := strings.Count(trimmed, "\n") + 1
	layers = append(layers, New(asciibot.Random(), 1, 1))
	sl := splitWidthWords(msg, msgwidth)
	boxX := 18
	boxY := 1
	if len(trimmed) > 0 {
		layers = append(layers, New(bubble(min(msgwidth, len(msg))+7, len(sl)+lineCount+1), boxX, boxY))
		for i, s := range sl {
			layers = append(layers, New(s, boxX+5, boxY+1+i))
		}
	}
	return strings.TrimRightFunc(render(layers), unicode.IsSpace) + "\n"
}

func main() {
	rainbowMode := false
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	if len(args) > 0 {
		if args[0] == "--help" {
			fmt.Println("usage: botsay [-c] [TEXT or \"-\"]")
			return
		} else if args[0] == "--version" {
			fmt.Println(versionString)
			return
		} else if args[0] == "-c" {
			rainbowMode = true
			if len(args) > 1 {
				args = args[1:]
			} else {
				args = []string{}
			}
		}
	}
	// Join all arguments to a single string
	msg := strings.Join(args, " ")
	// Read from /dev/stdin if "-" is given
	if msg == "-" {
		data, err := ioutil.ReadFile("/dev/stdin")
		if err != nil {
			panic(err)
		}
		msg = string(data)
	}
	if rainbowMode {
		rw := rainbow.NewTruecolorWriter(3, 0.4, 10)
		rw.Write([]byte(botsay(msg) + "\n"))
	} else {
		fmt.Println(botsay(msg))
	}
}
