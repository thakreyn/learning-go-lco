/*
	User Input and packages in Golang.

	For information on packages, one can visit: pkg.go.dev
	It contains all the needed information about the packages, simliar to flutter's packages

	There are no try catch in Go, by nature the language expects you to handle
	or name errors and treat them individually.
		This is done with the help of comma ok form of syntax.
	In this form, an object that can give an error, returns the error along with the desired data.
	It's the users choice to use this error or ignore using '_'

	Here we have used bufio and Reader, but the reader returns a string. Thus to use other values,
	we need to convert them later.

*/

package main

// Multiple imports can be clubbed in a single import statement
import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	message := "HELO"
	fmt.Println(message)

	// Go automatically imports or uses the os module
	// Also, we can observe that package names have lower case starting whereas,
	// Their public funstions have upper case
	path, error := os.UserHomeDir()
	fmt.Println(path, error)
	fmt.Printf("%T", path)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter some text : ")

	// Comma - ok syntax or error ok syntax
	// So some methods in Go return two values
	// In this case, one is the input read
	// And the other is error
	// We can either use the error or simply use '_' to ignore it
	// We cannot use _ as value. It is kinda reserved, only as intermediate

	input, _ := reader.ReadString('\n')
	fmt.Println(input)

}
