package main

import (
	"fmt"
	"math"
)

func main() {
	//-------------------------------------------------------------------------------------------------
	//--------------------------------------------Interfaces--------------------------------------------
	//-------------------------------------------------------------------------------------------------

	//The circle doesn't explicitly have to implement the shape interface, but since it contains an area() and perimeter() functions with the same signature
	//It is automatically considered a shape
	myCircle := circle{radius: 5}
	println(getShapeArea(myCircle))
	//A type can implement multiple interfaces, it only needs to have all the methods of these interfaces

	//Type assertions:
	checkIfCircle(myCircle)
	checkIfAnyWithSwitch(myCircle)
	checkIfAnyWithSwitchV2(myCircle)

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
