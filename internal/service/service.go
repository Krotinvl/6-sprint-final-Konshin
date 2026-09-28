package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ParseString(before string) (string, error) {
	if before == "" {
		return before, errors.New("Empty string")
	}

	isNotMorse := func(r rune) bool {
		return unicode.Is(unicode.Cyrillic, r)
	}

	if strings.IndexFunc(before, isNotMorse) != -1 {
		return morse.ToMorse(before), nil
	}

	return morse.ToText(before), nil
}
