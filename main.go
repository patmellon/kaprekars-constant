package main

import (
	"fmt"
	"kaprekars_constant/argvalidator"
	"kaprekars_constant/digitvalidator"
	sh "kaprekars_constant/slicehelper"
	"os"
	"strings"
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

	var digit sh.DigitSlice = strings.Split(digitString, "")

	fmt.Println(digit.SortAndConvert(true))
}
