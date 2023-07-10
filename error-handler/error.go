package e

import (
	"log"
)

func WithFatal[T any](result T, err error) func(s string) T {
	return func(s string) T {
		if err != nil {
			log.Fatalf("%s: %v", s, err)
		}
		return result
	}
}

func WithFatalString(err error, s string) {
		if err != nil {
			log.Fatalf("%s: %v", s, err)
		}
}
