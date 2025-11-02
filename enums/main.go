package main

import "fmt"

type Weekday int

// const (
// 	Monday    Weekday = 1
// 	Tuesday   Weekday = 2
// 	Wednesday Weekday = 3
// 	Thursday  Weekday = 4
// 	Friday    Weekday = 5
// 	Saturday  Weekday = 6
// 	Sunday    Weekday = 7
// )
// can be written with iota

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type Timezone int

const (
	// iota: 0, EST: -5
	EST Timezone = -(5 + iota)
	// _ is the blank identifier
	// iota: 1
	_
	// iota: 2, MST: -7
	MST
	// iota: 3, MST: -8
	PST
)

const (
	// Active = 0, Running = 100
	Active, Working = iota, iota + 100
	// Passive = 1, Stopped = 101
	Passive, Stopped
	// You can't declare like this.
	// The last expression will be repeated
	// CantDeclare => you need to pass two values
	// But, you can reset
	// the last expression
	Reset = iota
	// You can use any other
	// expression even without iota
	AnyOther = 10
)

const (
	// iota: 0, One: 1 (type: int64)
	One int64 = iota + 1
	// iota: 1, Two: 2 (type: int64)
	// Two will be declared as if:
	// Two int64 = iota + 1
	Two
	// iota: 2, Four: 4 (type: int32)
	Four int32 = iota + 2
	// iota: 3, Five: 5 (type: int32)
	// Five will be declared as if:
	// Five int32 = iota + 2
	Five
	// (type: int)
	Six = 6
	// (type: int)
	// Seven will be declared as if:
	// Seven = 6
	//	The last used expression and type will be repeated.
	Seven
)

type Activity int

// Use “iota + 1” to be sure that the enum type is initialized.
const (
	Sleeping = iota + 1
	Walking
)

func main() {
	fmt.Println(Monday)   // prints 1
	fmt.Println(Saturday) // prints 6
	fmt.Println(EST)      // prints -5
	fmt.Println(MST)      // prints -7

	var activity Activity
	fmt.Println(activity)
	// activity will be zero,
	// so it's not initialized
	activity = Sleeping
	// now you know that it's been
	// initialized
	fmt.Println(activity)
}
