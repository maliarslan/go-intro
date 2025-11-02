package main

import "fmt"

// https://go.dev/blog/maps

func main() {
	capitals := map[string]string{
		"Germany": "Berlin",
		"USA":     "Washington",
		"Turkey":  "Ankara",
	}

	fmt.Println(capitals["Germany"]) // => Berlin
	fmt.Println(capitals["UK"])      // => "" because its not defined so fall backs to zero-value

	capital, exists := capitals["UK"]

	if exists {
		fmt.Println(capital)
	} else {
		fmt.Println("Does not exist")
	}

	myMap := make(map[string]string)

	myMap["key"] = "value"

	for key, value := range myMap {
		fmt.Printf("this is key %s and this is value %s\n", key, value)
	}

	myNestedMap := map[string]map[string]string{}

	myNestedMap["firstMap"] = map[string]string{
		"nestedMap": "myFirstNestedMap",
	}

	fmt.Println(myNestedMap)
}
