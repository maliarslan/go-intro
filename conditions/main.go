package main

import "fmt"

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	age := 24

	// if is simple
	if age < 24 {
		fmt.Printf("%v is less than 24\n", age)
	} else if age > 24 {
		fmt.Printf("%v is greater than 24\n", age)
	} else {
		fmt.Printf("%v is equal to 24\n", age)
	}

	// https://go.dev/wiki/Switch
	// you dont necessarily need to break; but you can if you need to.
	// otherwise go breaks; like an if condition
	// there are fallthrough keyword to go through other cases
	// https://go.dev/wiki/Switch#fall-through

	day := Tuesday
	switch day {
	case Monday:
		fmt.Println("First day")
	case Tuesday, Wednesday, Thursday:
		fmt.Println("Week day")
	case Friday:
		fmt.Println("TGIF")

	default:
		fmt.Println("Weekend")
	}
}
