package utils

import (
	"errors"
	"unicode"
)

func IsEnglishWord(word string) (bool, error) {
	var isEng bool
	for i, char := range word {
		if i == 0 {
			isEng = unicode.Is(unicode.Latin, char)
		} else {
			if isEng != unicode.Is(unicode.Latin, char) {
				return false, errors.New("Word should contain symbol from one language!")
			}
		}
	}
	return isEng, nil
}
