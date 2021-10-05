/*
	Performing Requests and analyzing/parsing them in Go lang.

	The following Script needs a server to run. For this program , I have
	made a Flask API/Server that sends responses and is hosted on http://127.0.0.1:5000

	1. GET / => Returns a string
	2. GET /get => Returns a JSON
	3. POST /post => Echoes the same JSON that was sent
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

	const url string = "http://127.0.0.1:5000/post"

	// PerformGetRequest(url)
	PerformPostJsonRequest(url)
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

// Perform a POST request and send a JSON payload
func PerformPostJsonRequest(myurl string) {

	// Creating a JSON (in string format) and storing it in Reader format
	requestBody := strings.NewReader(`
	{
		"Name" : "CRY",
		"Age" : 20,
		"Height" : "6 feet"
	}
	`)

	// Making a post request
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	// ALWAYS REMEMBER TO CLOSE CONNECTION MANUALLY
	defer response.Body.Close()

	// Reading the response body in byte format
	byteData, _ := ioutil.ReadAll(response.Body)

	// Print the response body in string format
	fmt.Println(string(byteData))

}
