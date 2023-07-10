package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	LoadEnv()

	envValue := os.Getenv("API_KEY")
	if envValue != "WOLF" {
		t.Errorf("Expected 'WOLF', got %s", envValue)
	} else {
		fmt.Println("LoadEnv passed")
	}
	fmt.Println("--------")
}
