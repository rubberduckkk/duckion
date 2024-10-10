package util

import (
	"fmt"
	"strconv"
)

func ConvertStringToDigitArray(s string) ([]int, error) {
	arr := make([]int, 0, len(s))
	for i := range s {
		digit, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return nil, fmt.Errorf("convert string %v to digit array failed: %w", s, err)
		}
		arr = append(arr, digit)
	}
	return arr, nil
}
