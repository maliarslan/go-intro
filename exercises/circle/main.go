package main

import "fmt"

const PI = 3.14

type Circle struct {
	radius float64
}

func (circle *Circle) circumference() {
	circumference := 2 * PI * circle.radius
	fmt.Printf("circumference %v\n", circumference)
}

func (circle Circle) area() {
	area := PI * (circle.radius * circle.radius)
	fmt.Printf("area %v\n", area)
}

func main() {
	circle := Circle{
		radius: 5,
	}

	circle.circumference()
	circle.area()
}
