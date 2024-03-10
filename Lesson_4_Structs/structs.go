package main

import "fmt"

func main() {
	//-------------------------------------------------------------------------------------------------
	//---------------------------------------------Structs---------------------------------------------
	//-------------------------------------------------------------------------------------------------

	testMessageToSend(messageToSend{message: "Sending a message", phoneNumber: "03153234"})

	messageToSendVariable := messageToSend{message: "This is our variable's message"}
	messageToSendVariable.phoneNumber = "123"
	fmt.Println(messageToSendVariable.phoneNumber)

	//Anonymous structs
	myCar := struct {
		Make  string
		Model string
	}{
		Make:  "Toyota",
		Model: "Corolla",
	}

	fmt.Println(myCar.Make)

	//Nested anonymous structs
	myOtherCar := car{
		Make:   "Mercedes",
		Model:  "E350",
		Height: 3,
		Width:  2,
		Wheel: struct {
			Radius   int
			Material string
		}{
			Radius:   5,
			Material: "Wheel Material",
		},
	}
	myOtherCar.Wheel.Radius = 7

	fmt.Println(myOtherCar.Wheel.Material)

	//Embedded structs
	//You let this truct directly access the properties of the vehicle as if it inherited those properties from it
	myTruck := truck{
		vehicle: vehicle{ //You use the same name when embedding a struct, instead of writing vehicle: struct, you write vehicle: vehicle
			Model: "Our Model",
		},
		bedSize: 2,
	}
	fmt.Println(myTruck.Model)

	//Methods on structs
	myRect := Rect{
		width:  5,
		height: 2,
	}
	fmt.Println(area(myRect))
	fmt.Println(myRect.area_v2())

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
	Wheel  struct { //This wheel is only available inside car
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
