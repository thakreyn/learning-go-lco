/*
	Type Conversion (Str -> int):

		In order to convert the string to int, we need to use 'strconv' package.
		Also, we need to remove the trailing '\n'. For that we use 'strings' package.

		The 'strings' package is used to do actions on strings like slicing, cleaning, editing etc

		'strconv': Used to convert string to different types

	Basic error handling rule:
		The first thing we do when we encounter an error prone place or assignment is
		that we handle it immediatley after that.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome!")
	fmt.Println("Enter a number : ")

	// Init the reader
	reader := bufio.NewReader(os.Stdin)

	// Take input from reader
	// In this case, we are discarding or ignoring the error by using '_', but this is bad practice
	user_input, _ := reader.ReadString('\n')

	fmt.Println("Read : ", user_input)

	// We convert the string to int. but we also clean the string before so that only digits are left.
	// Also we store the error in 'err' to evaluate
	num_input, err := strconv.ParseInt(strings.TrimSpace("12.0"), 10, 64)

	// The following shows how we should check for errors
	// If there is no error -> error will be nil and we can carry on
	// Otherwise we have to handle the particular error
	if err != nil {
		fmt.Printf("%T\n", err)
		fmt.Println(err)
	} else {
		fmt.Printf("Type of num_input is : %T\n", num_input)
		fmt.Println("num_input square : ", num_input*num_input)
	}

	// Errors is a wide subject and will be discussed later
	// We can make custom errors as well
	// Error in itself is an interface
	// Also their are keywords like log.fatal or panic related to errors
}
