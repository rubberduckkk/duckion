package luhn

import (
	"fmt"
)

const (
	LENGTH_OF_VISA_MASTERCARD = 16
	LENGTH_OF_AMEX            = 15
)

func IsValidLuhn(digits []int) (bool, error) {
	n := len(digits)
	if n != LENGTH_OF_VISA_MASTERCARD && n != LENGTH_OF_AMEX {
		return false, fmt.Errorf("invalid number of digits: len=%v", n)
	}

	var sum int
	for i := len(digits) - 2; i >= 0; i-- {
		if i%2 == 0 {
			sum += digits[i] * 2
			if digits[i] > 4 {
				sum -= 9
			}
		} else {
			sum += digits[i]
		}
	}
	return 10-sum%10 == digits[len(digits)-1], nil
}
