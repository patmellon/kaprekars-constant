package main

import (
	"fmt"
	"kaprekars_constant/argvalidator"
	"kaprekars_constant/digitvalidator"
	"os"
)

func main() {

	argErr := argvalidator.Validate(os.Args)

	if argErr != nil {
		fmt.Println(argErr)
		os.Exit(1)
	}

	digitString := os.Args[1]

	charLengthErr := digitvalidator.ValidateCharacterLength(digitString)

	if charLengthErr != nil {
		fmt.Println(charLengthErr)
		os.Exit(1)
	}

	fmt.Println(digitString)
}
