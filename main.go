package main

import (
	"fmt"
	"kaprekars_constant/argvalidator"
	"kaprekars_constant/digitvalidator"
	"os"
	"sort"
	"strconv"
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

	fmt.Println(digitString)
}
