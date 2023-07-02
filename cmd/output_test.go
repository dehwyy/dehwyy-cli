package cmd

import (
	"fmt"
	"testing"
)


func TestStairsCaseCorrect(t *testing.T) {
	word := "Stairs"
	expected := "StAiRs"
	actual := StairsCase(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("Passed 1")
	}
}

func TestStairsCaseCorrectWithSpace(t *testing.T) {
	word := "Stairs test"
	expected := "StAiRs TeSt"
	actual := StairsCase(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("Passed 2")
	}
}

func TestStairsCaseSignleChar(t *testing.T) {
	word := "c"
	expected := "C"
	actual := StairsCase(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("Passed 3")
	}
}

func TestStairsCaseEmpty(t *testing.T) {
	word := ""
	expected := ""
	actual := StairsCase(word)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	} else {
		fmt.Println("Passed 3")
	}
}
