package ternary

import "testing"

func assert_ternary[T comparable](condition bool, ifTrue, isFalse, expected T, t *testing.T) {
	real := Use(condition, ifTrue, isFalse)
	if real != expected {
		t.Errorf("Expected %v, got %v", expected, real)
	}
}

func TestTernaryOperator(t *testing.T) {
	assert_ternary(1 < 200, 1, 200, 1, t)
	assert_ternary(1 > 200, 1, 200, 200, t)
}
