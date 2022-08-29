package digitvalidator

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func validateCharacterLength(digitString string) error {
	characterCount := utf8.RuneCountInString(digitString)

	if characterCount != 4 {
		return fmt.Errorf("Digit must have a length of 4")
	}

	return nil
}

func validateDigitIsInt(digitString string) error {
	_, err := strconv.Atoi(digitString)

	if err != nil {
		return fmt.Errorf("Argument is not a number")
	}

	return nil
}

func Validate(digitString string) error {
	charLengthErr := validateCharacterLength(digitString)
	intErr := validateDigitIsInt(digitString)

	if charLengthErr != nil {
		return charLengthErr
	} else if intErr != nil {
		return intErr
	}

	return nil
}
