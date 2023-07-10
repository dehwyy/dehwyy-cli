package e

import (
	"errors"
	"testing"
)

func mock_function(n int) (string, error) {
	if n > 50 {
		return "fail", errors.New("greater than 50")
	} else {
		return "success", nil
	}
}

func assert_error(n int) {
	WithFatal(mock_function(n))("Assertion")
}

func TestErrorHandlerSuccess(t *testing.T) {
	assert_error(0)
	assert_error(-50)
	assert_error(50)
}

func assert_withFatalString() {

}
