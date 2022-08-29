package digitvalidator

import (
	"fmt"
	"strconv"
	"strings"
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

func validateTwoDifferentDigits(digitString string) error {
	if containsSameDigits(digitString) == true {
		return fmt.Errorf("Digit must have at least two different numbers")
	}

	return nil
}

func containsSameDigits(digitString string) bool {
	splitString := strings.Split(digitString, "")

	for i := 0; i < len(splitString); i++ {
		if splitString[i] != splitString[0] {
			return false
		}
	}
	return true
}

func Validate(digitString string) error {
	validateFunctions := []func(string) error{validateCharacterLength, validateDigitIsInt, validateTwoDifferentDigits}

	for _, fn := range validateFunctions {
		err := fn(digitString)

		if err != nil {
			return err
		}
	}

	return nil
}
