package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidString = errors.New("invalid string")
	ErrSpecChar      = errors.New("found special char")
	ErrUppercaseChar = errors.New("found uppercase char")
)

func Unpack(s string) (string, error) {
	var b strings.Builder
	var lastChar rune

	sc := func(r rune) bool {
		return r < 'A' || r > 'z'
	}

	for i, val := range s {
		if !(unicode.IsDigit(val)) {
			b.WriteString(string(val))
		}

		if unicode.IsDigit(val) {
			if i == 0 || unicode.IsDigit(lastChar) {
				return "", ErrInvalidString
			}

			digit, err := strconv.Atoi(string(val))
			if err != nil {
				fmt.Println("syntax error")
			}

			if digit == 0 {
				removeLastChar(&b)
			} else {
				repeatChar(&b, string(lastChar), digit)
			}
		}

		if unicode.IsUpper(val) {
			return "", ErrUppercaseChar
		}

		if strings.IndexFunc(string(val), sc) != -1 && !(unicode.IsDigit(val)) {
			return "", ErrSpecChar
		}

		lastChar = val
	}

	output := b.String()
	return output, nil
}

func removeLastChar(b *strings.Builder) {
	l := b.String()
	l = l[:b.Len()-1]
	b.Reset()
	b.WriteString(l)
}

func repeatChar(b *strings.Builder, lastChar string, count int) {
	r := strings.Repeat(lastChar, count-1)
	b.WriteString(r)
}
