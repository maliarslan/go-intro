package main

import "fmt"

// a struct is just a data type
type Person struct {
	Name string
	Age  int
}

// this will get a copy of passed argument and change Name property but won't change on the original one
// func changeName(person Person, newName string)  {
// 	person.Name = newName
// }

// this will add (compose) this function to Person struct but still won't change the Name of the original object
// func (person Person) changeName(newName string) {
// 	person.Name = newName
// }

// now this with `*` will change property on the original object
// because it will directly operate on the allocated memory of the original object
// func (person *Person) changeName(newName string) {
// 	person.Name = newName
// }

// TIP: use &person to see the adress of allocated memory

func main() {
	// go doesn't complain if you don't pass values of all the fields in the given struct
	// because of `zero values`
	// zeroPerson := Person{}
	// fmt.Println(zeroPerson)
	// output => { 0}

	// person := Person{
	// 	Name: "Mehmet",
	// 	Age:  33,
	// }
	// fmt.Printf("I am %+v\n", person)

	// person.changeName("Mehmet Ali")
	// fmt.Printf("I am %+v", person)

	// a := 1
	// b := a
	// fmt.Println(b) => 1

	// a := 1
	// b := &a
	// fmt.Println(b) => prints allocated memory adress of a

	// a := 1
	// b := &a
	// *b = 3
	// fmt.Println(b) => prints allocated memory adress of a
	// fmt.Println(*b) => prints value of b which is 3
	// fmt.Println(a)  => prints value of a which is 3 because line 55 change the value of allocated memory which is for a

	mySlice := []int{
		1, 2, 3,
	}

	for index, _ := range mySlice {
		mySlice[index]++
	}

	fmt.Println(mySlice)
}
