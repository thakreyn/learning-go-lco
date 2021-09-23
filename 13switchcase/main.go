/*
	Switch Case:

	Switch case is also a way of controlling the flow of logic
	It is a vey simple way when we have to check only one variable
	for different conditions.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case: ")

	// Generating Random number using Math
	rand.Seed(time.Now().UnixNano())

	// For 1 to 6
	diceNumber := rand.Intn(6) + 1

	fmt.Println(diceNumber)

	// Clasic Switch Case syntax
	// switch <var-name>
	// For cases, unlike other languages, it only executes the current case
	// If we want to execute all the below, we can use fallthorigh keyword
	switch diceNumber {
	case 1:
		fmt.Println("Dice number is 1")
	case 2:
		fmt.Println("Dice number is 2")
		fallthrough
	case 3:
		fmt.Println("Dice number is 3")
	case 4:
		fmt.Println("Dice number is 4")
	case 5:
		fmt.Println("Dice number is 5")
	default:
		fmt.Println("Random!")

	}
}
