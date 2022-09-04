package computevalues

import (
	"fmt"
	sh "kaprekars_constant/slicehelper"
	"strconv"
	"strings"
)

const KaprekarsConstant = 6174

func ComputeDigit(d interface{}, iterations int) {

	digit := determineType(d)

	// Descending
	descendingInt := digit.SortAndConvert(false)

	// Ascending
	ascendingInt := digit.SortAndConvert(true)

	result := descendingInt - ascendingInt

	fmt.Println(result)

	if result != KaprekarsConstant {
		ComputeDigit(result, iterations+1)
	} else {
		fmt.Printf("Kaprekar's Constant found in %v iterations", iterations)
	}
}

func convertIntToDigitSlice(digit int) sh.DigitSlice {
	digitString := strconv.Itoa(digit)
	var newDigit sh.DigitSlice = strings.Split(digitString, "")

	return newDigit
}

func determineType(digit interface{}) sh.DigitSlice {
	switch digitType := digit.(type) {
	case sh.DigitSlice:
		return digitType
	case string:
		var d sh.DigitSlice = strings.Split(digitType, "")
		return d
	case int:
		var i sh.DigitSlice = convertIntToDigitSlice(digitType)
		return i
	default:
		return sh.DigitSlice{}
	}
}
