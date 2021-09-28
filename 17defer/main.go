/*
	Defer:

	It is a keyword that, when used in front of a function call, halts the execution
	of the statement or function until the end of the wrapping function.

	The order of multiple defer statements is LIFO (like a stack)

	Must Read! https://qvault.io/golang/defer-golang/
	It shows the proper explanation and use case for defer

*/
package main

import "fmt"

func trialFunction(toPrint bool) {

	if toPrint {
		fmt.Println("Does")
	}
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	// Output will be : 4,3,2,1

}

func main() {

	fmt.Println("Defer: ")
	defer trialFunction(false)

	trialFunction(true)
}
