package main

import (
	"errors"
	"fmt"
)

func main() {
	//-------------------------------------------------------------------------------------------------
	//--------------------------------------------Functions--------------------------------------------
	//-------------------------------------------------------------------------------------------------

	const a = 5
	const b = 4

	c := add(a, b)
	fmt.Println(c)

	//In go, you pass values, not references

	//You can ignore variables returned from functions like this:
	firstName, _ := getNames()
	//You can use the underscore _ to ignore a certain variable
	//The compiler removes it from our code
	fmt.Println(firstName)

	xV1, yV1 := getCoords_v1()
	fmt.Println(xV1, yV1)

	xV2, yV2 := getCoords_v2()
	fmt.Println(xV2, yV2)

	xV3, yV3 := getCoords_v3()
	fmt.Println(xV3, yV3)

	result, errorResult := divide(5, 0)
	fmt.Println(result, errorResult)

}

func sub(x int, y int) int {
	return x - y
}

// You can also add the types like this if the parameters are of the same type:
func add(x, y int) int {
	return x + y
}

// Functions can return more than 1 value
func getNames() (string, string) {
	return "John", "Doe"
}

// Functions can have named return types:
func getCoords_v1() (x, y int) {
	//It returns the default value of x and y
	//You would use a function like this when you want to document what x and y are for more context
	return //This is called a naked return, do not use this, use return x, y instead
}

// Overriding the named return types
func getCoords_v2() (x, y int) {
	return 5, 6 //Overrides x and y with these 2 hardcoded values
}

// Explicit return types
func getCoords_v3() (x, y int) {
	return x, y //Returns x and y
}

// Early returns/Guard clauses
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by 0")
	}
	return dividend / divisor, nil
}
