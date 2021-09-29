/*
	File Handling:

	File handling is an important feature for any programming language.
	In Go, there are various builtin features to make this process easy

	We use the 'os' package to create, modify, delete files
	'io' to Write to the files and
	'ioutil' to read from files

	Also defer is used a lot when it comes to opening and closing connections to a file.
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files: ")

	// Multiline strings in Go can be made using ` <string here> `
	// As it is raw, escape characters like /n, /t etc don't work on it
	// The data is not interpreted, it remains as it is
	content := `Raw string literals are character sequences
				between back quotes, as in foo.
			Within the quotes, any character may appear 
		except back quote.`

	// Comma-ok syntax is used to create a file
	file, err := os.Create("./mytextfile.txt")

	// It's best to deal with errors just after a command
	checkNilErr(err)

	// Successful writing to string returns the length of the file
	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length is: ", length)

	// It is a good practice to use defer when closing files or connections etc.
	// This is so that the connection closes after everything has executed
	// Also in case of error, it ensures that the resource is closed
	defer file.Close()

	readFile("./mytextfile.txt")
}

func readFile(fileName string) {

	// Files are never read in string format, always in bytes (rather slice of bytes)
	dataInBytes, err := ioutil.ReadFile(fileName)

	checkNilErr(err)

	// To convert bytes to string, we wrap it around with string()
	fmt.Println("Text data is : \n", string(dataInBytes))

}

// It is a common practice to make a function called 'checkNilErr' to handle errors
func checkNilErr(err error) {
	if err != nil {
		// Panic shuts down execution and shows the error
		panic(err)
	}
}
