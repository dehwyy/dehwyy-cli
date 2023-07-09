package cmd

import (
	"fmt"
	"log"
	"testing"
)

func testProvidedFlags(t *testing.T, expected bool, flags []string, n int) {
	actual := ProvidedLessThanTwoFlags(flags...)

	if expected != actual {
		log.Fatalf("Expected %t, got %t", expected, actual)
	}

	fmt.Printf("ProvidedLessThanTwoFlags: passed %d\n", n)
}

type FlagsTest struct {
	a string
	b string
	c string
	expected bool
}

func TestProvidedLessThanTwoFlagsOneFlag(t *testing.T) {
	testCases := []FlagsTest{
		// 1 flag
		{"a", "", "", true},
		{"", "b", "", true},
		{"", "", "c", true},
		// 2 flags
		{"a", "b", "", false},
		{"", "b", "c", false},
		{"a", "", "c", false},
		// 3 flags
		{"a", "b", "c", false},
	}
	for i, test := range testCases {
		flags := []string{test.a, test.b, test.c}
		testProvidedFlags(t, test.expected, flags, i + 1)
	}
	fmt.Println("----------")
}
