/*
	Maps:

	Maps are also called hash-tables or key-value pairs
	They are pretty simple in Go.


*/
package main

import "fmt"

func main() {

	fmt.Println("Works")

	// Declaration
	// We use the make command to init a map
	// map[<key_type>]<value_type>
	var languages = make(map[string]string)

	// Direct assignment
	languages["JS"] = "Javascript"
	languages["py"] = "python"
	languages["cpp"] = "C plus plus"

	fmt.Printf("%T\n", languages)
	fmt.Println(languages)

	fmt.Println(languages["cpp"])

	// Iterating (loops) in maps

	for key, value := range languages {
		fmt.Println(key, value)
	}

	// OR
	// If we can ignore either the key or value

	for key, _ := range languages {
		fmt.Println(key)
	}

}
