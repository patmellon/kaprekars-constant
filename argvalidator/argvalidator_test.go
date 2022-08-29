package argvalidator

import (
	"errors"
	"testing"
)

func TestSuccessfulValidate(t *testing.T) {

	args := []string{"Go compiler", "1234"}
	result := Validate(args)

	if result != nil {
		t.Errorf("Result was incorrect, got: %d, want: nil", result)
	}

}

func TestValidateWithNoArgs(t *testing.T) {

	args := []string{"Go compiler"}
	result := Validate(args)
	expectedError := errors.New("No argument provided")

	if result.Error() != expectedError.Error() {
		t.Errorf("got %q, expected %q", result, expectedError)
	}
}

func TestValidateWithTooManyArgs(t *testing.T) {

	args := []string{"Go compiler", "two", "three"}
	result := Validate(args)
	expectedError := errors.New("Only one argument accepted")

	if result.Error() != expectedError.Error() {
		t.Errorf("got %q, expected %q", result, expectedError)
	}
}
