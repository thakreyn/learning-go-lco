/*
	URL's (Creating and handling/parsing urls)

	The 2 main functions when it comes to URL's are:
		a. Reading and Parsing a URL
		b. Creating a URL

	The net/url package provides us with a URL struct that has predefined functions
	and data members that make the proces of dealing with URL's a lot simpler
*/
package main

import (
	"fmt"
	"net/url"
)

// URL's are generally declared as constants to prevent tampering
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&id=123123"

func main() {
	fmt.Println("Url's")
	fmt.Println(myurl)

	// Parsing the URL
	// The url library parses the url string and returns
	// a pointer to 'url.URL'
	result, err := url.Parse(myurl)
	checkNilError(err)

	// url.URL is a struct that has all the details regarding a URL
	fmt.Printf("%T <-> %v\n", result, result)

	// Printing some URL members
	fmt.Println(result.Host)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// In order to get the quries, we user .Query()
	// This returns a map of queries and their values
	queryParams := result.Query()

	fmt.Printf("%T\n", queryParams)
	fmt.Println(queryParams)

	// Iterating over the map
	for key, val := range queryParams {
		fmt.Println(key, val)
	}

	// Building a URL

	// Although there are libraries that allow building of URL's
	// If we want to do it manually, we simply create a reference to url.URL
	// and init it with the required values
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=CRY",
	}

	// Converting the given url to string
	// string(<url.URL>) can also be used
	urlString := partsOfUrl.String()
	fmt.Println(urlString)
}

// General function to check nil error
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
