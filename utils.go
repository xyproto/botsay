package main

import (
	"strings"
	"unicode/utf8"
)

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

// Returns unicode string runes length, not bytes
func uLen(s string) int {
	return utf8.RuneCountInString(s)
}

// Return the width and height of a given ASCII art string
func size(s string) (int, int) {
	maxWidth := 0
	maxHeight := 0
	lineCounter := 0
	for _, line := range strings.Split(s, "\n") {
		l := uLen(line)
		if l > maxWidth {
			maxWidth = l
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
func get(s []rune, x, y, w, h int) rune {
	if x < 0 || y < 0 {
		return ' '
	}
	if x >= w || y >= h {
		return ' '
	}
	// +1 to account for the trailing newlines
	pos := y*w + x
	if pos >= len(s) {
		return ' '
	}
	r := s[pos]
	switch r {
	case '\n', '\t', '\r', '\v':
		return ' '
	default:
		return r
	}
}

// Convert from a multiline-string to an indexed slice of runes (y*w+x style)
func toMap(s string, w int) []rune {
	rs := make([]rune, 0)
	for _, line := range strings.Split(s, "\n") {
		rs = append(rs, []rune(line)...)
		linelen := uLen(line)
		if linelen < w {
			// Fill out the rest of the line with spaces
			rs = append(rs, []rune(strings.Repeat(" ", w-linelen))...)
		}
	}
	return rs
}

// Like a blit function, but for ASCII graphics. Uses " " as the "transparent pixel".
func combine(a, b string, xoffset, yoffset int) string {
	aW, aH := size(a)
	bW, bH := size(b)
	maxW := max(aW, bW+xoffset)
	maxH := max(aH, bH+yoffset)
	aMap := toMap(a, aW)
	bMap := toMap(b, bW)
	var sb strings.Builder
	for y := 0; y < maxH; y++ {
		for x := 0; x < maxW; x++ {
			if get(bMap, x-xoffset, y-yoffset, bW, bH) == ' ' {
				sb.WriteRune(get(aMap, x, y, aW, aH))
			} else {
				sb.WriteRune(get(bMap, x-xoffset, y-yoffset, bW, bH))
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

// splitWords can split a string into words, keeping punctuation and trailing spaces
func splitWords(s string) []string {
	var (
		splitpoint bool
		words      []string
		letters    strings.Builder
		tmp        string
	)
	lenS := uLen(s)
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
			if uLen(tmp) > 0 {
				words = append(words, tmp)
			}
			letters.Reset()
		} else {
			letters.WriteRune(r)
		}
	}
	tmp = strings.TrimSpace(letters.String())
	if uLen(tmp) > 0 {
		words = append(words, tmp)
	}
	return words
}

// splitWithWords can split a string by words, then combine to form lines maximum w long
func splitWidthWords(s string, w int) []string {
	var (
		sl   []string
		line string
	)
	for _, word := range splitWords(s) {
		if uLen(line)+uLen(word) < w {
			line += word
		} else {
			trimmedLine := strings.TrimSpace(line)
			if strings.HasSuffix(trimmedLine, "--") {
				// Move the double dash to the beginning of the next line
				trimmedLine = trimmedLine[:uLen(trimmedLine)-2]
				sl = append(sl, trimmedLine)
				line = "-- " + word
			} else {
				sl = append(sl, trimmedLine)
				line = word
			}
		}
	}
	if uLen(line) == 0 {
		return sl
	}
	return append(sl, strings.TrimSpace(line))
}
