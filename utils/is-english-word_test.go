package utils

import (
	"fmt"
	"testing"
)

func assert_is_english_word(word string, expected, hasError bool, t *testing.T, i int) {
	IsEng, err := IsEnglishWord(word)

	// Error should be when hasError is True, otherwise it should be not
	if hasError && err == nil {
		t.Errorf("Expected error, got nil : %v ", err)
	} else if !hasError && err != nil {
		t.Errorf("Expected nil, got error : %v ", err)
	}


	if IsEng != expected {
		t.Errorf("Expected %v, got %v", expected, IsEng)
	} else {
		fmt.Printf("IsEnglishWord: Passed %d\n", i)
	}
}

func TestIsEnglishWord(t *testing.T) {
	type T struct {
		word string
		expected bool
		hasError bool
	}

	testCases := []T{
		{"english", true, false},
		{"русский", false, false},
		{"russкий", false, true},
	}

	for i, test := range testCases {
		assert_is_english_word(test.word, test.expected, test.hasError, t, i + 1)
	}
	fmt.Println("----------")
}
