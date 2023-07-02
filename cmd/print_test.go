package cmd

import (
	"fmt"
	"testing"
)

// StairsCase
func TestStairsCase(t *testing.T) {
	word := "Stairs"
	expected := "StAiRs"
	actual := StairsCase(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("StairsCase: Passed 1")
	}
}

func TestStairsCaseWithSpace(t *testing.T) {
	word := "Stairs test"
	expected := "StAiRs TeSt"
	actual := StairsCase(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("StairsCase: Passed 2")
	}
}

func TestStairCaseIgnoreCase(t *testing.T) {
	word := "STAIRS"
	expected := "StAiRs"
	actual := StairsCase(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("StairsCase: Passed 3")
	}
}

func TestStairsCaseSignleChar(t *testing.T) {
	word := "c"
	expected := "C"
	actual := StairsCase(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("StairsCase: Passed 4")
	}
}

func TestStairsCaseEmpty(t *testing.T) {
	word := ""
	expected := ""
	actual := StairsCase(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("StairsCase: Passed 5")
	}
	fmt.Println("----------")
}

// Capitalize

func TestCapitalize(t *testing.T) {
	word := "hello"
	expected := "Hello"
	actual := Capitalize(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("Capitalize: Passed 1")
	}
}

func TestCapitalizeWithManyWords(t *testing.T) {
	word := "hello world"
	expected := "Hello world"
	actual := Capitalize(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("Capitalize: Passed 2")
	}
}

func TestCapitalizeSingleChar(t *testing.T) {
	word := "c"
	expected := "C"
	actual := Capitalize(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("Capitalize: Passed 3")
	}
}

func TestCapitalizeEmpty(t *testing.T) {
	word := ""
	expected := ""
	actual := Capitalize(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("Capitalize: Passed 4")
	}
	fmt.Println("----------")
}
