package main

import "fmt"

func main() {
	x := 5
	y := 5
	z := &x //Here, you store the address of x as a value for z. When working with pointers, you use the & sign

	x = 7

	fmt.Println(x, y, *z) //Outputs 7 5 7

	*z = 3                //The asterisk * is how you dereference a variable to access its value
	fmt.Println(x, y, *z) //Outputs 3 5 3

	c := car{
		color: "White",
	}

	c.setColor("Blue")
	fmt.Println(c)

}

//A pointer is a variable that stores a memory address

//Pointers are typically used with method arguments to make sure that the outer reference is updated

type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}
