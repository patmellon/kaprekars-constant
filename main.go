package main

import (
	"fmt"
	"kaprekars_constant/argvalidator"
	"os"
	"unicode/utf8"
)

func main() {

	argErr := argvalidator.Validate(os.Args)

	if argErr != nil {
		fmt.Println(argErr)
		os.Exit(1)
	}

	digitString := os.Args[1]

	charLengthErr := checkCharacterLength(digitString)

	if charLengthErr != nil {
		fmt.Println(charLengthErr)
		os.Exit(1)
	}

	fmt.Println(digitString)
}

func checkCharacterLength(digitString string) error {
	characterCount := utf8.RuneCountInString(digitString)

	if characterCount != 4 {
		return fmt.Errorf("Digit must have a length of 4")
	}

	return nil
}
