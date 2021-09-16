/*
	Hello World

	In this we just see the basics of a Go program.
	Just like languages like C/C++, Go programs have an entry point called 'main'

	Thus we use the main function for the entry to a program. Also, just like everything in Java is an object,
	similarly in Go, everything is a Type.

	To run a Go file, we use 'go run <filename>' to run it.

*/

// At the top of the code, we mention the package to which the file belongs
package main

// Import statements import packages to the code
import "fmt"

// Single line comments are like this

/*
	These are multiline comments
*/

// For functions, the curly braces need to start on the first line itself.
// func keyword is used.
// Semi-colons are suggested but not mandatory
func main() {

	// The fmt package is a package that has methods like
	// printf, println, print for output
	fmt.Println("Hello World!")
}
