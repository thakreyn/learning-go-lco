/*
	Performing Requests and analyzing/parsing them in Go lang.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web Request Verbs")

	PerformGetRequest("http://127.0.0.1:5000/")
}

// Function that performs a GET request to the given URL
func PerformGetRequest(myurl string) {

	// Store http/response from Get
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	// Close the connection
	defer response.Body.Close()

	// Printing status code for the response
	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content Length : ", response.ContentLength)

	// Using strings Builder to print the string format for response body
	var responseString strings.Builder

	// Reading the response body
	// ReadAll returns a byte array
	byteData, _ := ioutil.ReadAll(response.Body)

	// We write the byte array to the string builder
	byteForm, _ := responseString.Write(byteData)

	fmt.Println(byteForm)

	// string builder to string
	fmt.Println(responseString.String())

	// Or
	fmt.Println(string(byteData))

}
