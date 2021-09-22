/*
	If - Else

	If - else are used to control the flow of the code and our vital components
	of any programming language.

	In go, the If - else syntax is very rigid and there's only 1 defined way to use
	them so that there's no ambiguity of formatting.

*/
package main

import (
	"fmt"
)

func main() {

	fmt.Println("If - Else")

	// Init 2 vars together
	var (
		loginCount int = 12
		result     string
	)

	fmt.Println(loginCount, result)

	// Normal if-else statement
	// Condition is written first and then the curly braces start the block
	// Else/else if has to be on the same line as the ending of prev if
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount >= 10 {
		result = "Suspicion"
	} else {
		result = "Never Reached"
	}

	fmt.Println("The user type is :", result)

	// Another Example of two conditions
	if 9%2 == 0 && 8%3 != 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// Web Request Handling example
	// Sometimes, we can directly assign a variable using := operator
	// and then check the condition
	if num := 3; num%3 == 0 {
		fmt.Print("Accepted")
	} else {
		fmt.Println("Rejected")
	}

	// Another use is Error Handling
	var some_error *string = nil

	if some_error != nil {
		fmt.Println("Some Error Here")
	}
}
