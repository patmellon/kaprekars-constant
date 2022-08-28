package main

import (
	"fmt"
	"kaprekars_constant/argchecks"
	"os"
	"unicode/utf8"
)

func main() {

	argChecksErr := argchecks.CheckArgs(os.Args)

	if argChecksErr != nil {
		fmt.Println(argChecksErr)
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
