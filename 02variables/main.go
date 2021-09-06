/*
	Variables, Types and Constants in Go.

	Variables in Go can be declared in three way:
		1. Proper Representation
		2. Implicit Types
		3. Shorthand (only inside methods not globally/script form)

	Everything in Go is a type. Numeric, string, array, list etc etc.
	Basic types for variables:
		1. int (various),
		2. float32, float64
		3. string
		4. bool

	Constants in Go can be initialised at any place where var can be used.
	It can be done using 'const' in place of 'var'

	In Go, if we initialie a variable and don't use it, it raises an error.

	https://gobyexample.com/ -> is a good resource to read short go tutorials.

*/

package main

import "fmt"

const globalVar = 320

func main() {

	// Variables

	// 1st Way
	// In Go, if you init a variable and dont use it, then it raises error

	var username string = "thakreyn"

	fmt.Println(username)
	fmt.Printf("Variable type : %T\n", username)

	var isGood bool = false

	fmt.Println(isGood)
	fmt.Printf("Variable type : %T\n", isGood)

	// We can initialise multiple variables using the following syntax

	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// Numbers specs can be seen in Golang docs under lang specs in Numeric Types

	var smallVal int = 24
	fmt.Println(smallVal)
	fmt.Printf("Variable type : %T\n", smallVal)

	// Float32 prints 2 decimal places while float64 has higher precision

	var smallFloat float32 = 242424.249123123
	fmt.Println(smallFloat)
	fmt.Printf("Variable type : %T\n", smallFloat)

	// Default values and Aliases
	// Instead of storing garbage values, it defaults to certain values
	// int, float -> 0
	// string -> ""
	// bool -> false

	var randomVariable bool
	fmt.Println(randomVariable)
	fmt.Printf("Variable type : %T\n", randomVariable)

	// 2nd Way
	// Implicit Types
	// var <var-name> = <value>
	// Thus, Go interprets the datatype from assignment
	// float -> float64
	// Once Go has decided/interpretted a type, it cannot be changed

	var siteName = "thakreyn.in"
	fmt.Println(siteName)
	fmt.Printf("Variable type : %T\n", siteName)

	// 3rd way - shorthand
	// No var style (using := operator)
	// Only works when inside a method, not globally as a script

	userName := "crysys"
	fmt.Println(userName)
	fmt.Printf("Variable type : %T\n", userName)

	// Using a global variable

	fmt.Println(globalVar)

}
