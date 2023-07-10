package ternary


func Use[T any](condition bool, ifTrue, isFalse T) T {
	if condition {
		return ifTrue
	}
	return isFalse
}
