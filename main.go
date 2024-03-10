package main

import (
	"errors"
	"fmt"
	"math"
)

//Lessons: Variables and Types, Conditions, Functions, Structs, Interfaces, Errors, Loops, Arrays, Slices

func main() {

	////-------------------------------------------------------------------------------------------------
	////---------------------------------------Variables and Types---------------------------------------
	////-------------------------------------------------------------------------------------------------
	//
	//var smsSendingLimit int
	//var costPerSms float64
	//var hasPermission bool
	//var username string
	//
	//fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSms, hasPermission, username)
	//
	//numCars := 0
	//temparature := 0.0
	//isFunny := false
	//empty := ""
	//
	////%T shows you the type instead of the value
	////%v is value %s is string %d is integer %f is float
	//fmt.Printf("%v %f %v %T\n", numCars, temparature, isFunny, empty)
	//
	//averageRate, displayMessage := 5.0, false
	//fmt.Printf("%v %v\n", averageRate, displayMessage)
	//
	//temparatureInt := 88
	//temperatureFloat := float64(temparatureInt)
	//
	//fmt.Printf("%f \n", temperatureFloat)
	//
	////You need to declare constants like this, no short-handed syntax
	//const premiumPlanName = "Premium plan"
	////Use Sprintf when you want to return the formatted string and assign it to a variable
	//planString := fmt.Sprintf("Your plan is %s", premiumPlanName)
	//fmt.Println(planString)
	//
	////-------------------------------------------------------------------------------------------------
	////--------------------------------------------Conditions--------------------------------------------
	////-------------------------------------------------------------------------------------------------
	//
	//height := 5
	//if height > 4 {
	//	fmt.Println("You are tall enough")
	//}
	//
	//if weight := 20; weight > 10 {
	//	fmt.Println("You are overweight")
	//}
	//
	////-------------------------------------------------------------------------------------------------
	////--------------------------------------------Functions--------------------------------------------
	////-------------------------------------------------------------------------------------------------
	//
	//const a = 5
	//const b = 4
	//
	//c := add(a, b)
	//fmt.Println(c)
	//
	////In go, you pass values, not references
	//
	////You can ignore variables returned from functions like this:
	//firstName, _ := getNames()
	////You can use the underscore _ to ignore a certain variable
	////The compiler removes it from our code
	//fmt.Println(firstName)
	//
	//xV1, yV1 := getCoords_v1()
	//fmt.Println(xV1, yV1)
	//
	//xV2, yV2 := getCoords_v2()
	//fmt.Println(xV2, yV2)
	//
	//xV3, yV3 := getCoords_v3()
	//fmt.Println(xV3, yV3)
	//
	//result, errorResult := divide(5, 0)
	//fmt.Println(result, errorResult)
	//
	////-------------------------------------------------------------------------------------------------
	////---------------------------------------------Structs---------------------------------------------
	////-------------------------------------------------------------------------------------------------
	//
	//testMessageToSend(messageToSend{message: "Sending a message", phoneNumber: "03153234"})
	//
	//messageToSendVariable := messageToSend{message: "This is our variable's message"}
	//messageToSendVariable.phoneNumber = "123"
	//fmt.Println(messageToSendVariable.phoneNumber)
	//
	////Anonymous structs
	//myCar := struct {
	//	Make  string
	//	Model string
	//}{
	//	Make:  "Toyota",
	//	Model: "Corolla",
	//}
	//
	//fmt.Println(myCar.Make)
	//
	////Nested anonymous structs
	//myOtherCar := car{
	//	Make:   "Mercedes",
	//	Model:  "E350",
	//	Height: 3,
	//	Width:  2,
	//	Wheel: struct {
	//		Radius   int
	//		Material string
	//	}{
	//		Radius:   5,
	//		Material: "Wheel Material",
	//	},
	//}
	//myOtherCar.Wheel.Radius = 7
	//
	//fmt.Println(myOtherCar.Wheel.Material)
	//
	////Embedded structs
	////You let this truct directly access the properties of the vehicle as if it inherited those properties from it
	//myTruck := truck{
	//	vehicle: vehicle{ //You use the same name when embedding a struct, instead of writing vehicle: struct, you write vehicle: vehicle
	//		Model: "Our Model",
	//	},
	//	bedSize: 2,
	//}
	//fmt.Println(myTruck.Model)
	//
	////Methods on structs
	//myRect := Rect{
	//	width:  5,
	//	height: 2,
	//}
	//fmt.Println(area(myRect))
	//fmt.Println(myRect.area_v2())
	//
	////-------------------------------------------------------------------------------------------------
	////--------------------------------------------Interfaces--------------------------------------------
	////-------------------------------------------------------------------------------------------------
	//
	////The circle doesn't explicitly have to implement the shape interface, but since it contains an area() and perimeter() functions with the same signature
	////It is automatically considered a shape
	//myCircle := circle{radius: 5}
	//print(getShapeArea(myCircle))
	////A type can implement multiple interfaces, it only needs to have all the methods of these interfaces
	//
	////Type assertions:
	//checkIfCircle(myCircle)
	//checkIfAnyWithSwitch(myCircle)
	//checkIfAnyWithSwitchV2(myCircle)
	//
	////-------------------------------------------------------------------------------------------------
	////----------------------------------------------Errors----------------------------------------------
	////-------------------------------------------------------------------------------------------------
	//
	//result2, errorResult2 := divideV2(10.0, 0.0)
	//if errorResult2 != nil {
	//	formattedErrorResult := fmt.Errorf("operation failed: %w", errorResult2)
	//	fmt.Println(formattedErrorResult)
	//} else {
	//	fmt.Println(result2)
	//}
	//
	//result3, errorResult3 := divideV3(10.0, 0.0)
	//var divErr *DivisionError
	//if errorResult3 != nil && errors.As(errorResult3, &divErr) {
	//	fmt.Println(errorResult3)
	//} else {
	//	fmt.Println(result3)
	//}
	//
	////-------------------------------------------------------------------------------------------------
	////----------------------------------------------Loops----------------------------------------------
	////-------------------------------------------------------------------------------------------------
	//
	////Same as loops in other languages, but without the parenthesis
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//finalCost := bulkSend(10)
	//fmt.Println(finalCost)
	//
	//attempts := maxMessages(12.0)
	//fmt.Println(attempts)
	//
	//getPlanHeight()

	////-------------------------------------------------------------------------------------------------
	////----------------------------------------------Arrays----------------------------------------------
	////-------------------------------------------------------------------------------------------------
	//
	//var myInts [10]int
	//fmt.Println(myInts) //Prints out: [0 0 0 0 0 0 0 0 0 0]
	//
	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(primes) //Prints out: [2 3 5 7 11 13]
	//
	//messageWithRetries := getMessageWithRetries()
	//fmt.Println(messageWithRetries) //Prints out: [Click here to sign up Pretty please click here We beg you to sign up]
	//
	////-------------------------------------------------------------------------------------------------
	////----------------------------------------------Slices----------------------------------------------
	////-------------------------------------------------------------------------------------------------
	//
	//subPrimes := primes[1:4] //1 is inclusive, 4 is exclusive
	//fmt.Println(subPrimes)   //Prints out: [3 5 7]
	//
	//allPrimes := primes[:] //Prints out the whole list
	//fmt.Println(allPrimes) //Prints out: [2 3 5 7 11 13]
	//
	//primesStartingFrom := primes[1:] //1 is inclusive
	//fmt.Println(primesStartingFrom)  //Prints out: [3 5 7 11 13]
	//
	//primesEndingAt := primes[:4] //4 is exclusive
	//fmt.Println(primesEndingAt)  //Prints out: [2 3 5 7]
	//
	//fmt.Println(len(primesEndingAt))
	//fmt.Println(cap(primesEndingAt))
	//
	////Primitve values like strings, integers... are passed by value to functions
	////If their value changes inside that function, it won't change outside the scope of that function
	////But with slices, they are passed by reference, and any modification inside the function will affect it outside that function
	//
	//mySlice := make([]int, 5, 10) //Make a slice with 5 integers having a capacity of 10, exceeding this capacity will re-assign the array in memory
	//fmt.Println(mySlice)
	//
	//myStaticSlice := make([]int, 5) //Creates a slice of 5 zeros
	//fmt.Println(myStaticSlice)
	//
	////Variadic functions contain println, sprintf, printf...
	//
	//sumOfNumbers := sum(1, 2, 3)
	//fmt.Println(sumOfNumbers)
	//
	//numbersToSum := []int{1, 2, 3, 4}
	//fmt.Println(sum(numbersToSum...)) //Spread operator
	//
	//numbersToSum = append(numbersToSum, 5) //Outputs: [1 2 3 4 5]
	//fmt.Println(numbersToSum)
	//
	//numbersToSum = append(numbersToSum, 6, 7) //Outputs: [1 2 3 4 5 6 7]
	//fmt.Println(numbersToSum)
	//
	//numbersToSum = append(numbersToSum, numbersToSum...) //Outputs: [1 2 3 4 5 6 7 1 2 3 4 5 6 7]
	//fmt.Println(numbersToSum)
	//
	//allCosts := []cost{{1, 5.0}, {2, 2.0}, {3, 3.0}, {5, 1.0}}
	//costsByDay := getCostsByDay(allCosts)
	//fmt.Println(costsByDay)
	//
	////forEachIndexed equivalent:
	//fruits := []string{"Apple", "Orange", "Banana"}
	//for index, element := range fruits {
	//	fmt.Println(index, element)
	//}
	////Outputs:
	////0 Apple
	////1 Orange
	////2 Banana

	//-------------------------------------------------------------------------------------------------
	//-----------------------------------------------Maps-----------------------------------------------
	//-------------------------------------------------------------------------------------------------

	ages := make(map[string]int)
	ages["Thaer"] = 26
	fmt.Println(ages)

	ages_v2 := map[string]int{
		"Thaer": 27,
	}
	fmt.Println(ages_v2)

	names := []string{"Thaer", "Thaer2", "Thaer3"}
	phoneNumbers := []string{"+96171700996", "+96171700997", "+96171700998"}
	usersMap, _ := getUserMap(names, phoneNumbers)
	fmt.Println(usersMap["Thaer"])

	//Inserting an element into a map:
	usersMap["Lara"] = user{
		name:        "Lara",
		phoneNumber: "+96171700995",
	}
	fmt.Println(usersMap["Lara"])

	//Getting an element
	userLara := usersMap["Lara"]
	fmt.Println(userLara)

	//Deleting an element
	delete(usersMap, "Lara")
	fmt.Println(usersMap)

	_, ok := usersMap["Thaer"]
	fmt.Println(ok)

	//Shorthanded syntax for the above check:
	if _, ok := usersMap["Lara"]; !ok {
		fmt.Println("Lara not found")
	}

	//The key has to be a comparabletype, you can't use slices, maps, and functions since they point to addresses in memory
	//And you can't compare addresses, you need to compare values
	//structs are comparable

	//-------------------------------------------------------------------------------------------------
	//--------------------------------------Higher Order Functions--------------------------------------
	//-------------------------------------------------------------------------------------------------

}

//-------------------------------------------------------------------------------------------------
//--------------------------------------------Functions--------------------------------------------
//-------------------------------------------------------------------------------------------------

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

//-------------------------------------------------------------------------------------------------
//---------------------------------------------Structs---------------------------------------------
//-------------------------------------------------------------------------------------------------

type messageToSend struct {
	message     string
	phoneNumber string
}

func testMessageToSend(send messageToSend) {
	fmt.Println(send.message)
	fmt.Println(send.phoneNumber)
}

// You can nest anonymous structs like this:
type car struct {
	Make   string
	Model  string
	Height int
	Width  int
	Wheel  struct {
		Radius   int
		Material string
	}
}

// Embedded structs
type vehicle struct {
	Make  string
	Model string
}

type truck struct {
	vehicle
	bedSize int
}

// Defining methods on structs:
type Rect struct {
	width  int
	height int
}

func area(r Rect) int {
	return r.width * r.height
}

// You can define it as an extension like this:
func (r Rect) area_v2() int {
	return r.width * r.height
}

//-------------------------------------------------------------------------------------------------
//--------------------------------------------Interfaces--------------------------------------------
//-------------------------------------------------------------------------------------------------

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func getShapeArea(sh shape) float64 {
	return sh.area()
}

// You can name the parameters of your interface's functions:
type Copier interface {
	Copy(sourceFile string, destinationFile string) (bytesCoped int)
}

func checkIfCircle(sh shape) {
	//This is how you check if something is of a certain type
	//You use interface variable.(type): sh.(circle)
	//ok represents whether or not this cast was successful
	myCircle, ok := sh.(circle)
	if ok {
		fmt.Println(myCircle.radius, ok)
	}
}

func checkIfAnyWithSwitch(sh shape) {
	switch sh.(type) {
	case circle:
		fmt.Println("This is a circle")
	case rect:
		fmt.Println("This is a rect")
	default:
		fmt.Println("This is none")
	}
}

func checkIfAnyWithSwitchV2(sh shape) {
	switch typ := sh.(type) {
	case circle:
		fmt.Printf("This is a circle with a radius of %.1f", typ.radius)
	case rect:
		fmt.Printf("This is a rectangle with an area of %f", typ.area())
	default:
		fmt.Println("This is none")
	}
}

// Interfaces extending other interfaces:
type transportationMeans interface {
	getModel() string
	getMake() string
}

type firetruck interface {
	transportationMeans
	HoseLength() int
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

//-------------------------------------------------------------------------------------------------
//-----------------------------------------------Maps-----------------------------------------------
//-------------------------------------------------------------------------------------------------

func getUserMap(names []string, phoneNumbers []string) (map[string]user, error) {
	userMap := map[string]user{}
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid sizes")
	}
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber string
}
