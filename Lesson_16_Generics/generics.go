package main

import "fmt"

func main() {
	fmt.Println(Swap(1, 2))             // works with int
	fmt.Println(Swap("hello", "world")) // works with string

	//Constraints
	fmt.Println(Sum(Int(1), Int(2))) // Output: 3

	//Type Sets and Interfaces
	fmt.Println(Add(1, 2))     // int
	fmt.Println(Add(1.5, 2.3)) // float64
	//fmt.Println(Add("hello", "world")) // This line would cause a compile-time error

}

// Swap swaps the values of a and b.
func Swap[T any](a, b T) (T, T) {
	return b, a
}

//Constraints

//In this example, Adder is an interface that defines a constraint.
//Only types that satisfy the Adder interface (i.e., types that have an Add method) can be used with the Sum function.

type Adder[T any] interface {
	Add(a, b T) T
}

func Sum[T Adder[T]](a, b T) T {
	return a.Add(a, b)
}

type Int int

func (i Int) Add(a, b Int) Int {
	return a + b
}

//Type Sets and Interfaces

type Number interface {
	~int | ~float64 // Type set syntax: only types identical to int or float64
}

func Add[T Number](a, b T) T {
	return a + b
}
