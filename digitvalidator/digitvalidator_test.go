package digitvalidator

import (
	"errors"
	"testing"
)

func TestValidateWithNoError(t *testing.T) {

	digitString := "1234"
	result := validateCharacterLength(digitString)

	if result != nil {
		t.Errorf("Result was incorrect, got: %d, want: nil", result)
	}

}

func TestValidateLengthError(t *testing.T) {

	tests := []string{"123", "45678"}

	expectedError := errors.New("Digit must have a length of 4")

	for _, s := range tests {
		result := validateCharacterLength(s)

		if result.Error() != expectedError.Error() {
			t.Errorf("Result was incorrect, got: %d, want: nil", result)
		}
	}

}

func TestValidateIntError(t *testing.T) {

	digitString := "word"
	expectedError := errors.New("Argument is not a number")

	result := validateDigitIsInt(digitString)

	if result.Error() != expectedError.Error() {
		t.Errorf("got %q, expected %q", result, expectedError)
	}

}

func TestValidateWithNoErrors(t *testing.T) {
	digitString := "1234"

	result := Validate(digitString)

	if result != nil {
		t.Errorf("Result was incorrect, got: %d, want: nil", result)
	}
}

func TestValidateWithCharErr(t *testing.T) {
	digitString := "123"
	expectedError := errors.New("Digit must have a length of 4")

	result := Validate(digitString)

	if result.Error() != expectedError.Error() {
		t.Errorf("Result was incorrect, got: %d, want: nil", result)
	}
}

func TestValidateWithIntErr(t *testing.T) {
	digitString := "word"
	expectedError := errors.New("Argument is not a number")

	result := Validate(digitString)

	if result.Error() != expectedError.Error() {
		t.Errorf("Result was incorrect, got: %d, want: nil", result)
	}
}
