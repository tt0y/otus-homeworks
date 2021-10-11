package hw02unpackstring

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var (
	lastRune        rune
	lastRuneIsDigit bool
	result          strings.Builder
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	for i, currentRune := range packedString {
		if unicode.IsDigit(currentRune) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(currentRune) && lastRuneIsDigit {
			return "", ErrInvalidString
		}

		if unicode.IsLetter(currentRune) {
			lastRuneIsDigit = false
			result.WriteRune(currentRune)
			lastRune = currentRune
		}

		if unicode.IsDigit(currentRune) {
			lastRuneIsDigit = true
			runesCount, err := strconv.Atoi(string(currentRune))
			if err != nil {
				log.Fatal(err)
			}

			if runesCount < 1 {
				str := result.String()
				if len(str) > 0 {
					str = str[:len(str)-1]
					result.Reset()
					result.WriteString(str)
				}
			}

			for j := 0; j < runesCount-1; j++ {
				result.WriteRune(lastRune)
			}
		}
	}

	defer result.Reset()

	return result.String(), nil
}
