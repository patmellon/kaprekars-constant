package digitvalidator

import (
	"fmt"
	"unicode/utf8"
)

func ValidateCharacterLength(digitString string) error {
	characterCount := utf8.RuneCountInString(digitString)

	if characterCount != 4 {
		return fmt.Errorf("Digit must have a length of 4")
	}

	return nil
}
