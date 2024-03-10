package main

import "fmt"

func main() {
	//-------------------------------------------------------------------------------------------------
	//----------------------------------------------Slices----------------------------------------------
	//-------------------------------------------------------------------------------------------------

	primes := [6]int{2, 3, 5, 7, 11, 13}

	subPrimes := primes[1:4] //1 is inclusive, 4 is exclusive
	fmt.Println(subPrimes)   //Prints out: [3 5 7]

	allPrimes := primes[:] //Prints out the whole list
	fmt.Println(allPrimes) //Prints out: [2 3 5 7 11 13]

	primesStartingFrom := primes[1:] //1 is inclusive
	fmt.Println(primesStartingFrom)  //Prints out: [3 5 7 11 13]

	primesEndingAt := primes[:4] //4 is exclusive
	fmt.Println(primesEndingAt)  //Prints out: [2 3 5 7]

	fmt.Println(len(primesEndingAt))
	fmt.Println(cap(primesEndingAt))

	//Primitve values like strings, integers... are passed by value to functions
	//If their value changes inside that function, it won't change outside the scope of that function
	//But with slices, they are passed by reference, and any modification inside the function will affect it outside that function

	mySlice := make([]int, 5, 10) //Make a slice with 5 integers having a capacity of 10, exceeding this capacity will re-assign the array in memory
	fmt.Println(mySlice)

	myStaticSlice := make([]int, 5) //Creates a slice of 5 zeros
	fmt.Println(myStaticSlice)

	//Variadic functions contain println, sprintf, printf...

	sumOfNumbers := sum(1, 2, 3)
	fmt.Println(sumOfNumbers)

	numbersToSum := []int{1, 2, 3, 4}
	fmt.Println(sum(numbersToSum...)) //Spread operator

	numbersToSum = append(numbersToSum, 5) //Outputs: [1 2 3 4 5]
	fmt.Println(numbersToSum)

	numbersToSum = append(numbersToSum, 6, 7) //Outputs: [1 2 3 4 5 6 7]
	fmt.Println(numbersToSum)

	numbersToSum = append(numbersToSum, numbersToSum...) //Outputs: [1 2 3 4 5 6 7 1 2 3 4 5 6 7]
	fmt.Println(numbersToSum)

	allCosts := []cost{{1, 5.0}, {2, 2.0}, {3, 3.0}, {5, 1.0}}
	costsByDay := getCostsByDay(allCosts)
	fmt.Println(costsByDay)

	//forEachIndexed equivalent:
	fruits := []string{"Apple", "Orange", "Banana"}
	for index, element := range fruits {
		fmt.Println(index, element)
	}
	//Outputs:
	//0 Apple
	//1 Orange
	//2 Banana
}

//-------------------------------------------------------------------------------------------------
//----------------------------------------------Slices----------------------------------------------
//-------------------------------------------------------------------------------------------------

func sum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	days := []int{}

	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		if contains(days, cost.day) {
			costsByDay[len(costsByDay)-1] += cost.value
		} else {
			days = append(days, cost.day)
			costsByDay = append(costsByDay, cost.value)
		}
	}

	return costsByDay
}

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
