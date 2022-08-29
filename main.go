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

	splitString := strings.Split(digitString, "")

	sortAscend := copyDigitSlice(splitString)
	sortDescend := copyDigitSlice(splitString)

	// Sorts ascending
	sort.Strings(sortAscend)

	// Sorts descending
	sort.Slice(sortDescend, func(i, j int) bool {
		return sortDescend[i] > sortDescend[j]
	})

	largerInt := convertToInt(sortDescend)
	smallerInt := convertToInt(sortAscend)

	fmt.Println(splitString)
	fmt.Println(largerInt)
	fmt.Println(smallerInt)
}

func convertToInt(digit []string) int {
	intVar, err := strconv.Atoi(strings.Join(digit, ""))

	if err != nil {
		fmt.Println(nil)
		os.Exit(1)
	}

	return intVar
}

func copyDigitSlice(splitString []string) []string {
	return append([]string{}, splitString...)
}
