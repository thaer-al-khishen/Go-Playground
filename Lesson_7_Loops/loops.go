package main

import "fmt"

func main() {
	//-------------------------------------------------------------------------------------------------
	//----------------------------------------------Loops----------------------------------------------
	//-------------------------------------------------------------------------------------------------

	//Same as loops in other languages, but without the parenthesis
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	finalCost := bulkSend(10)
	fmt.Println(finalCost)

	attempts := maxMessages(12.0)
	fmt.Println(attempts)

	getPlanHeight()
}

//-------------------------------------------------------------------------------------------------
//----------------------------------------------Loops----------------------------------------------
//-------------------------------------------------------------------------------------------------

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

func maxMessages(thresh float64) int {
	totalCost := 0.0
	//INITIAL, CONDITION, AFTER
	//You can emit any of these 3, if you emit the condition, it will run forever
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}

// Go doesn't have a while loop
// You just use a for loop with just a condition
func getPlanHeight() {
	planHeight := 1
	for planHeight < 5 {
		fmt.Println("Still growing! ", planHeight)
		planHeight++
	}
	fmt.Println("Plan has grown to: ", planHeight)
}

// Continue
// Used to stop the current iteration of a loop and continue to the next
func printOddNumbers() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

// Break
// Used to stop the current iteration of a loop and continue to the next
func stopAt5() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}
