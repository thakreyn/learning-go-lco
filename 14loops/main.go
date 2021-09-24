/*
	Loops:

	There is only type of loop in Go. That is the for loop.
	But this 'for' loop has different forms for various situations.

	Also, some other keywords that are useful are:
		1. break - breaks the loop and moves control to after the loop
		2. continue - Goes on to the next iteration of the loop
		3. goto - Takes the control of the program to a given label

*/
package main

import "fmt"

func main() {
	fmt.Println("Loops: ")

	days := []string{"mon", "tue", "wed", "thur", "fri", "sat"}
	fmt.Println(days)

	// Loop Type 1
	// Normal Traditional for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// Loop type 2
	// When you have to loop through all elements
	// i returns index
	for i := range days {
		fmt.Println(days[i])
	}

	// Loop type 3
	// When you have to loop thorugh entire
	// We get both index and value
	for index, day := range days {
		fmt.Println(index, day)
	}

	randomValue := 13

	// 'For' used like a 'while'
	// break, continue and goto
	for randomValue < 20 {

		if randomValue == 16 {
			break
		}

		if randomValue == 14 {
			// syntax : goto <label_name>
			goto randomLine
		}

		fmt.Println(randomValue)
		randomValue++
	}

	// This is a goto label
	// The control goes here and code starts executing from hereafter
	// It is like a checkpoint
randomLine:
	fmt.Println("Going to a random line")

	fmt.Println("At end")
}
