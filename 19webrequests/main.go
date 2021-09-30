/*
	Handling Web Requests:

	Web requests in Go are handled using the 'http' library
	It contains loads of functions and built-in structs that can
	be used to make requests, get responses and evaluate them

	It is very important to note that the requests have to be closed
	manually. EVERYTIME! (Thus, it is best rpactice to use defer keyword)

	*> Must Read : http docs on pkg.go.dev
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// URL's are generally declared as constants
const url = "https://lco.dev"

func main() {
	fmt.Println("Web Requests")

	// Making a simple GET request
	response, err := http.Get(url)

	checkNilError(err)

	// Repsonse is a pointer to the struct called http.Response
	// Thus we can use the pointer to change data
	fmt.Printf("%+v", response)
	fmt.Printf("%T", response)

	// This is the most imp statment because http doesn't close the
	// connections and thus has to be done manually
	defer response.Body.Close()

	// Reading response body using ioutil
	data, err := ioutil.ReadAll(response.Body)
	checkNilError(err)

	fmt.Println(string(data))

}

// Common practice to make a function to check nil error
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
