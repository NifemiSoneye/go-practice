package main

import (
	"errors"
)

type divideError struct {
	dividend float64
}

// ?

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.0 , errors.New("No dividing by 0")
	}
	return dividend / divisor, nil
}

