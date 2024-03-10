package main

import "fmt"

func main() {

	//-------------------------------------------------------------------------------------------------
	//---------------------------------------Variables and Types---------------------------------------
	//-------------------------------------------------------------------------------------------------

	var smsSendingLimit int
	var costPerSms float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSms, hasPermission, username)

	numCars := 0
	temparature := 0.0
	isFunny := false
	empty := ""

	//%T shows you the type instead of the value
	//%v is value %s is string %d is integer %f is float
	fmt.Printf("%v %f %v %T\n", numCars, temparature, isFunny, empty)

	averageRate, displayMessage := 5.0, false
	fmt.Printf("%v %v\n", averageRate, displayMessage)

	temparatureInt := 88
	temperatureFloat := float64(temparatureInt)

	fmt.Printf("%f \n", temperatureFloat)

	//You need to declare constants like this, no short-handed syntax
	const premiumPlanName = "Premium plan"
	//Use Sprintf when you want to return the formatted string and assign it to a variable
	planString := fmt.Sprintf("Your plan is %s", premiumPlanName)
	fmt.Println(planString)

}
