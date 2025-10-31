package imports

import "fmt"

// go makes a field/function public by checking the first char
// capital = public, small = private
// ticket.ID = public, ticket.event = private

type Ticket struct {
	ID    int
	Event string
}

func (ticket Ticket) PrintEvent() {
	fmt.Println(ticket.Event)
}
