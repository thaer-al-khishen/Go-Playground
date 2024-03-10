package main

import "fmt"

func main() {
	fmt.Println(aggregate(2, 3, 4, add))      //Outputs 9
	fmt.Println(aggregate(2, 3, 4, multiply)) //Outputs 24

	//Currying functions
	add2 := addV2(2)    // This gives us a function that adds 2 to its argument
	result := add2(3)   // Now, we use that function with 3 as the argument
	fmt.Println(result) // Output: 5

	//Defer
	//The defer statement in Go (Golang) is a powerful feature for managing resources cleanly and efficiently,
	//ensuring that cleanup actions are performed later, usually for purposes of releasing resources or doing necessary clean-up tasks.
	//When you use defer, the function call it precedes is delayed until the surrounding function completes, either because it reached its end,
	//returned, or encountered a panic. This behavior is particularly useful for closing files, releasing locks, or any other cleanup tasks
	//that you need to guarantee happen regardless of where the function exits due to errors or return statements.
	defer fmt.Println("world")
	fmt.Println("hello")
	//Outputs Hello World

	//Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
	//Outputs:
	//start
	//end
	//middle

	//Closures
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("Khishen")
	fmt.Println(harryPotterAggregator("24"))
	//Outputs: Mr. Khishen 24

}

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

// Currying
// Currying is a functional programming concept where a function that takes multiple arguments is transformed into a sequence of functions,
// each taking a single argument. In essence, it's about breaking down a function that takes multiple arguments into a series of functions
// that take one argument each. This technique can lead to more modular and reusable code.
// First function takes an int and returns another function
func addV2(a int) func(int) int {
	// The returned function takes an int and returns the result
	return func(b int) int {
		return a + b
	}
}

// Closures
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}
