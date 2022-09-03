package slicehelper

import (
	"reflect"
	"testing"
)

func getSmallerDigit() DigitSlice {
	return DigitSlice{"1", "2", "3", "4"}
}

func getLargerDigit() DigitSlice {
	return DigitSlice{"4", "3", "2", "1"}
}

func TestSortSliceAscending(t *testing.T) {
	sorted := getLargerDigit().sortSlice(true)
	expectedResult := getSmallerDigit()

	if reflect.DeepEqual(sorted, expectedResult) == false {
		t.Errorf("got %q, expected %q", sorted, expectedResult)
	}
}

func TestSortSliceDescending(t *testing.T) {
	sorted := getSmallerDigit().sortSlice(false)
	expectedResult := getLargerDigit()

	if reflect.DeepEqual(sorted, expectedResult) == false {
		t.Errorf("got %q, expected %q", sorted, expectedResult)
	}
}

// Need a test for convertSliceToInt returning an error
func TestConvertSliceToInt(t *testing.T) {
	digitInt := getLargerDigit().convertSliceToInt()
	expectedResult := 4321

	if digitInt != expectedResult {
		t.Errorf("got %q, expected %q", digitInt, expectedResult)
	}
}

func TestSortAndConvertDescending(t *testing.T) {
	sortAndConverted := getSmallerDigit().SortAndConvert(false)
	expectedResult := 4321

	if sortAndConverted != expectedResult {
		t.Errorf("got %q, expected %q", sortAndConverted, expectedResult)
	}
}

func TestSortAndConvertAscending(t *testing.T) {
	sortAndConverted := getLargerDigit().SortAndConvert(true)
	expectedResult := 1234

	if sortAndConverted != expectedResult {
		t.Errorf("got %q, expected %q", sortAndConverted, expectedResult)
	}
}
