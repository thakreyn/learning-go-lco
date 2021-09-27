/*
	Functions:

	Functions are an important part of programming.
	In Go, we use 'func' keyword to define functions

	Main Features:
		1. Can return multiple types
		2. Take in multiple args
*/

package main

import "fmt"

func main() {
	fmt.Println("Functions : ")

	// Normal Function Call
	greeter()

	// Function call with a return value and parameters
	result := addition(1, 3)

	// Function returning multiple values and taking indefinite values
	proResult, message := proAdder("Test String", 1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(result)
	fmt.Println(proResult, message)
}

// Function without a return type
func greeter() {
	fmt.Println("Inside Greeter")
}

// Syntax:
// Function signature: func <func_name> (parameters dataType) <return_type>
func addition(num1 int, num2 int) int {
	return num1 + num2
}

// Taking multiple indefinite values as parameter
// ... -> takes arguments and returns a slice of all of them
// For returning multiple values, we wrap them in ()
func proAdder(message string, values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}

	return total, "Status: Addition confirmed!"
}
