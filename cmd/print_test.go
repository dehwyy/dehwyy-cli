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

func TestStairsCase(t *testing.T) {
	testCases := []struct{
		word string
		expected string
	}{
		{"Stairs", "StAiRs"},
		{"Stairs test", "StAiRs TeSt"},
		{"STAIRS", "StAiRs"},
		{"c", "C"},
		{"", ""},
	}

	for i, test := range testCases {
		stairsCaseSuitTest(test.word, test.expected, i + 1, t)
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

func TestCapitalize(t *testing.T) {
	testCases := []struct {
		word string
		expected string
	}{
		{"hello", "Hello"},
		{"hello world", "Hello world"},
		{"c", "C"},
		{"", ""},
	}

	for i, test := range testCases {
		capitalizeSuitTest(test.word, test.expected, i + 1, t)
	}

	fmt.Println("__________")
}

// ReverseCapitalize

func reverseCapitalizeSuitTest(w, expected string, i int, t *testing.T) {
	actual := toReverseCapitalize(w)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	fmt.Printf("ReverseCapitalize: passed %d\n", i)
}

func TestReverseCapitalize(t *testing.T) {
	testCases := []struct {
		word string
		expected string
	}{
		{"hello", "hELLO"},
		{"hello world", "hELLO WORLD"},
		{"c", "c"},
		{"", ""},
	}

	for i, test := range testCases {
		reverseCapitalizeSuitTest(test.word, test.expected, i + 1, t)
	}

	fmt.Println("__________")
}
