package main

import (
	"errors"
	"fmt"
)

// Custom Error
type DivideError struct {
	firstValue   int
	secondValue  int
	errorMessage string
}

// Implementing the error Interface
func (a *DivideError) Error() string {
	return a.errorMessage
}

func divide(firstValue int, secondValue int) (int, error) {
	if secondValue == 0 {
		divisionByZeroError := DivideError{
			firstValue:   firstValue,
			secondValue:  secondValue,
			errorMessage: "Can't divide with zero",
		}
		return -1, &divisionByZeroError
	}
	return firstValue / secondValue, nil
}
func main() {
	firstValue := 20
	secondValue := 0
	
	result, err := divide(firstValue, secondValue)
	if err != nil {
		var pointerToACustomErrorDivideError *DivideError;
		if errors.As((err), &pointerToACustomErrorDivideError) {
			fmt.Printf("Divide by 0 error")
		} else {
			fmt.Printf("Unexpected error");
		}
	} else {
		fmt.Printf("Result is : %d", result);
	}
	return;
}
