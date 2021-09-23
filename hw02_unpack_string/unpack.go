package hw02unpackstring

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var lastRune rune
var result strings.Builder

var ErrInvalidString = errors.New("invalid string")

func Unpack(PackedString string) (string, error) {

	for i, currentRune := range PackedString {
		if unicode.IsDigit(currentRune) && i == 0 {
			return "", ErrInvalidString
		}

		if unicode.IsLetter(currentRune) {
			result.WriteRune(currentRune)
			lastRune = currentRune
		}

		if unicode.IsDigit(currentRune) {
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

	return result.String(), nil
}
