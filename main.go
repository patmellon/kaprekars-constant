package main

import (
	"fmt"
	"kaprekars_constant/argvalidator"
	c "kaprekars_constant/computevalues"
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

	digitErr := digitvalidator.Validate(digitString)

	if digitErr != nil {
		fmt.Println(digitErr)
		os.Exit(1)
	}

	c.ComputeDigit(digitString, 1)
}
