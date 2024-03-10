package main

import (
	"errors"
	"fmt"
)

func main() {
	//-------------------------------------------------------------------------------------------------
	//----------------------------------------------Errors----------------------------------------------
	//-------------------------------------------------------------------------------------------------

	result2, errorResult2 := divideV2(10.0, 0.0)
	if errorResult2 != nil {
		formattedErrorResult := fmt.Errorf("operation failed: %w", errorResult2)
		fmt.Println(formattedErrorResult)
	} else {
		fmt.Println(result2)
	}

	result3, errorResult3 := divideV3(10.0, 0.0)
	var divErr *DivisionError
	if errorResult3 != nil && errors.As(errorResult3, &divErr) {
		fmt.Println(errorResult3)
	} else {
		fmt.Println(result3)
	}

}

//-------------------------------------------------------------------------------------------------
//----------------------------------------------Errors----------------------------------------------
//-------------------------------------------------------------------------------------------------

func divideV2(a, b float64) (float64, error) {
	if b == 0.0 {
		//You can use either syntax below
		//return 0, errors.New("division by zero")
		return 0, fmt.Errorf("cannot divide %v by zero", a)
	}
	return a / b, nil
}

// You can define custom error types by implementing the error interface.
// This is useful for creating error types that can carry additional context or metadata about the error
type DivisionError struct {
	Dividend float64
	Divisor  float64
	Message  string
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("%s: %v / %v", e.Message, e.Dividend, e.Divisor)
}

func divideV3(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0, &DivisionError{a, b, "division by zero"}
	}
	return a / b, nil
}
