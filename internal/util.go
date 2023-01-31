package internal

import (
	"strconv"
	"strings"
	"unicode"
)

func atoi(x string) (int, error) {
	y, err := strconv.ParseInt(x, 10, 32)
	return int(y), err
}

func stripSpaces(x string) string {
	var b strings.Builder
	b.Grow(len(x))
	for _, ch := range x {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
