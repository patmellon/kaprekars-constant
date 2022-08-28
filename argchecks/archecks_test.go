package argchecks

import (
	"errors"
	"testing"
)

func TestSuccessfulCheckArgs(t *testing.T) {

	args := []string{"Go compiler", "1234"}
	result := CheckArgs(args)

	if result != nil {
		t.Errorf("Result was incorrect, got: %d, want: nil", result)
	}

}

func TestUnsuccessfulCheckArgs(t *testing.T) {

	args := []string{"Go compiler"}
	result := CheckArgs(args)
	expectedError := errors.New("No argument provided")

	if result.Error() != expectedError.Error() {
		t.Errorf("got %q, expected %q", result, expectedError)
	}
}
