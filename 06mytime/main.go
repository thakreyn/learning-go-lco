/*
	Time package

	This package is used to use the functionality of time in Go.
	It provides us with a time object, ability to do operations on time and store time related information.

	We can check the documentation for additional information about the package.

	In order to print the time object in the required format, we need to use the .Format function. Inside this,
	we have to give the layout string and this string has to contain a certain date.

	The string is : 01-02-2006 15:04:05 Monday

	We can alternate the format to represent time.
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // mm - dd - yyyy time day
	fmt.Println(presentTime.Format("01/02/2006"))                 // mm / dd / yyyy

}
