package main

import "fmt"

type rect struct {
	width  float64
	height float64
}

type cir struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c cir) area() float64 {
	return 2 * 3.14 * c.radius
}

func (c cir) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	r := rect{
		width:  10,
		height: 50,
	}
	rectArea := r.area()
	fmt.Println("Area: ", rectArea)

	rectPerimeter := r.perimeter()
	fmt.Println("Perimeter: ", rectPerimeter)

	c := cir{radius: 4}

	cirArea := c.area()
	fmt.Println("Circle Area: ", cirArea)

	cirPerimeter := c.perimeter()
	fmt.Println("Circle Perimeter: ", cirPerimeter)
}
