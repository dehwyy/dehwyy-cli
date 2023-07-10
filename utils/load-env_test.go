package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	LoadEnv()

	envValue := os.Getenv("YANDEX_TRANSLATE_API_KEY")
	if len(envValue) == 0 {
		t.Errorf("Expected not nil value', got %s", envValue)
	} else {
		fmt.Println("LoadEnv passed")
	}
	fmt.Println("--------")
}
