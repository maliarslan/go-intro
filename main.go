package main

import (
	"femBasics/imports"
	"fmt"
)

func main() {
	fmt.Println("hello")

	newTicket := imports.Ticket{
		ID:    1,
		Event: "Go",
	}

	fmt.Printf("newTicket: %+v\n", newTicket)
	newTicket.PrintEvent()
}
