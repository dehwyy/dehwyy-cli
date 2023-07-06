package cmd

import (
	"fmt"
	"testing"
)

// StairsCase

func stairsCaseSuitTest(w, expected string, i int, t *testing.T) {
	actual := toStairsCase(w)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	fmt.Printf("StairsCase: passed %d\n", i)
}

type StairsCaseTest struct {
	word string
	expected string
}

func TestStairsCase(t *testing.T) {
	testCases := []StairsCaseTest{
		{"Stairs", "StAiRs"},
		{"Stairs test", "StAiRs TeSt"},
		{"STAIRS", "StAiRs"},
		{"c", "C"},
		{"", ""},
	}

	for i, test := range testCases {
		stairsCaseSuitTest(test.word, test.expected, i, t)
	}

	fmt.Println("__________")
}


// Capitalize

func capitalizeSuitTest(w, expected string, i int, t *testing.T) {
	actual := toCapitalize(w)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	fmt.Printf("Capitalize: passed %d\n", i)
}

type CapitalizeTest struct {
	word string
	expected string
}

func TestCapitalize(t *testing.T) {
	testCases := []CapitalizeTest{
		{"hello", "Hello"},
		{"hello world", "Hello world"},
		{"c", "C"},
		{"", ""},
	}

	for i, test := range testCases {
		capitalizeSuitTest(test.word, test.expected, i, t)
	}

	fmt.Println("__________")
}
