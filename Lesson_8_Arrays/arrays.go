package main

import "fmt"

func main() {
	//-------------------------------------------------------------------------------------------------
	//----------------------------------------------Arrays----------------------------------------------
	//-------------------------------------------------------------------------------------------------

	var myInts [10]int
	fmt.Println(myInts) //Prints out: [0 0 0 0 0 0 0 0 0 0]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) //Prints out: [2 3 5 7 11 13]

	messageWithRetries := getMessageWithRetries()
	fmt.Println(messageWithRetries) //Prints out: [Click here to sign up Pretty please click here We beg you to sign up]

}

//-------------------------------------------------------------------------------------------------
//----------------------------------------------Arrays----------------------------------------------
//-------------------------------------------------------------------------------------------------

func getMessageWithRetries() [3]string {
	return [3]string{
		"Click here to sign up",
		"Pretty please click here",
		"We beg you to sign up",
	}
}
