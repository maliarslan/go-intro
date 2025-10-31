package main

import (
	"fmt"
	"slices"
)

func main() {
	// arrays are one type and have definite index
	// animals := [2]string{}
	// animals[0] = "dog"
	// animals[1] = "cat"
	// OR
	// animals := [2]string{
	// 	"dog",
	// 	"cat",
	// }

	// slice = flexible, dynamic size arrays
	animals := []string{
		"dog", // added to try slices.Delete
		"dog",
		"cat",
	}

	animals = append(animals, "fish")      // equivalent of [].push()
	animals = slices.Delete(animals, 0, 1) // deletes element from a slice between given range

	// **append actually gets the copy of the passed slice**
	// it passes a copy of animals, so here animals value won't be changed
	// not like in javascript “pass by reference” when dealing with objects (including arrays and functions)
	// although pass by reference seems possible as well in go
	var newAnimals = append(animals, "lion")

	fmt.Println(newAnimals, animals)

	for i := 0; i < len(animals); i++ {
		fmt.Printf("this is my animal %v\n", animals[i])
	}

	for index, value := range animals {
		fmt.Printf("this is my index %v and animal %v\n", index, value)
	}

	// loop through [0, n]
	// for value := range 10 {
	// 	fmt.Println(value)
	// }

	// equivalent of while
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
