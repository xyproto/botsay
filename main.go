package main

import (
	"fmt"
	"github.com/mattes/go-asciibot"
	"os"
	"strings"
	"unicode"
)

const (
	boxContentWidth = 42
	versionString   = "botsay 1.0.1"
)

type GFX struct {
	ascii string
	x     int
	y     int
}

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
	max_width := 0
	max_height := 0
	line_counter := 0
	for _, line := range strings.Split(s, "\n") {
		if len(line) > max_width {
			max_width = len(line)
		}
		line_counter++
	}
	if line_counter > max_height {
		max_height = line_counter
	}
	return max_width, max_height
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
			} else {
				return " "
			}
		}
	}
	return " "
}

// Like a blit function, but for ASCII graphics. Uses " " as the "transparent pixel".
func combine(a, b string, xoffset, yoffset int) string {
	a_w, a_h := size(a)
	b_w, b_h := size(b)
	max_w := max(a_w, b_w+xoffset)
	max_h := max(a_h, b_h+yoffset)
	var sb strings.Builder
	for y := 0; y < max_h; y++ {
		for x := 0; x < max_w; x++ {
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

// Generate ASCII graphics of a randomly generated bot with a speech bubble
func botsay(msg string) string {
	var layers []*GFX
	msgwidth := boxContentWidth
	layers = append(layers, New(asciibot.Random(), 1, 1))
	sl := splitWidthWords(msg, msgwidth)
	boxX := 18
	boxY := 1
	layers = append(layers, New(bubble(min(msgwidth, len(msg))+7, len(sl)+2), boxX, boxY))
	for i, s := range sl {
		layers = append(layers, New(s, boxX+5, boxY+1+i))
	}
	return strings.TrimRightFunc(render(layers), unicode.IsSpace) + "\n"
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "--" {
		args = args[1:]
	}
	if len(args) > 0 {
		if args[0] == "--help" {
			fmt.Println("usage: botsay [TEXT]")
			return
		} else if args[0] == "--version" {
			fmt.Println(versionString)
			return
		}
	}
	fmt.Println(botsay(strings.Join(args, " ")))
}
