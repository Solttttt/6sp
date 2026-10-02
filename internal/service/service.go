package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseTranslater(inStr string) (string, error) {
	var outStr string
	inStr = strings.TrimSpace(inStr)

	if inStr == "" {
		return "", errors.New("empty document")
	}

	checkStr := strings.ReplaceAll(inStr, "-", "")
	checkStr = strings.ReplaceAll(checkStr, ".", "")
	checkStr = strings.ReplaceAll(checkStr, " ", "")

	if len(checkStr) > 0 {
		outStr = morse.ToMorse(strings.ToUpper(inStr))
	} else {
		outStr = morse.ToText(inStr)
	}

	return outStr, nil
}
