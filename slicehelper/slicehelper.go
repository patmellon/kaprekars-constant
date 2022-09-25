package slicehelper

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type DigitSlice []string

func (d DigitSlice) sortSlice(ascending bool) DigitSlice {
	if ascending {
		// Sorts ascending
		sort.Strings(d)
	} else {
		// Sorts descending
		sort.Slice(d, func(i, j int) bool {
			return d[i] > d[j]
		})
	}

	return d
}

func (d DigitSlice) convertSliceToInt() int {
	digit, err := strconv.Atoi(strings.Join(d, ""))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return digit
}

func (d DigitSlice) sortAndConvert(ascending bool) int {
	return d.sortSlice(ascending).convertSliceToInt()
}

func (d DigitSlice) SortDescending() int {
	return d.sortAndConvert(false)
}

func (d DigitSlice) SortAscending() int {
	return d.sortAndConvert(true)
}
