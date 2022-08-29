package digitvalidator

import (
	"errors"
	"testing"
)

func TestValidateWithNoError(t *testing.T) {

	digitString := "1234"
	result := ValidateCharacterLength(digitString)

	if result != nil {
		t.Errorf("Result was incorrect, got: %d, want: nil", result)
	}

}

func TestValidateWithError(t *testing.T) {

	tests := []string{"123", "45678"}

	expectedError := errors.New("Digit must have a length of 4")

	for _, e := range tests {
		result := ValidateCharacterLength(e)

		if result.Error() != expectedError.Error() {
			t.Errorf("Result was incorrect, got: %d, want: nil", result)
		}
	}

}
