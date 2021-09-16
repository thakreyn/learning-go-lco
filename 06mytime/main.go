/*
	Time package

	This package is used for functionality of time in Go.
	It provides us with a time object, ability to do operations on time and store time related information.

	We can check the documentation for additional information about the package.

	In order to print the time object in the required format, we need to use the .Format function. Inside this,
	we have to give the layout string and this string has to contain a certain date.

	The string is : 01-02-2006 15:04:05 Monday

	We can alternate the format to represent time.

	Building Executables:

	We can use the Go 'env' command to view all the environment variables for our go environment.
	The main ones to focus while building are:
		1. GOOS : target OS
		2. GOARCH : target Architecture

	So, when we run `go build`, it automatically searches for the main.go file and builds for the target OS.

	If we want to generate files for other OS's we can change the go env variables. GOOS = linux or GOOS = darwin (for Mac)

	`go help build` shows us all flags and options for the build function.
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	// Initializes a time object
	presentTime := time.Now()

	// Print the time in different formats
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // mm - dd - yyyy time day
	fmt.Println(presentTime.Format("01/02/2006"))                 // mm / dd / yyyy

	// Sleep function is one of many in time package
	time.Sleep(10 * time.Second)

	fmt.Println(presentTime)
	// Time package has its own data types for seconds, weeks, days etc
}
