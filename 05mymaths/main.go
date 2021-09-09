/*
	In this part, we cover the ways of generating random numbers in Go.
	Also, the beauty of the Go documentation and package description.

	In order to generate random numbers in Go, we can use 2 methods:
		1. Math/rand
		2. crypto/rand

	The methods and documentation for these can be seen of pkg.go.dev. On the webpage for
	math/rand, it clearly states that it generates pseudo random numbers and should not be used for
	security sensitive uses.

	Whereas, crypto/rand is made specifically for that purpose. The documentation gives clear and concise
	info on how to use the packages.

	Although, doing small things seems very cumbersome in Go, It's beauty is in the fact that everything
	has a specific purpose and is very well documented. So we can use any package without much hassle.


*/

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	fmt.Println("Welcome message")

	var n1 int = 1
	var n2 float64 = 3

	fmt.Println(float64(n1) + n2)
	fmt.Printf("%T\n", n1)

	// Random Numbers
	// 1. Math using custom seed
	// 2. Random using cryptography

	// Math/rand says that it generates pseduo random number.
	// Not to be used with security sensitive data
	// The pseudo random number will be same for a givn seed, thus we give it a dynamic seed

	// rand.Seed(901004)

	// var rn = rand.Intn(25) + 1
	// fmt.Println(rn)
	// fmt.Printf("%T\n", rn)

	var properRand, _ = rand.Int(rand.Reader, big.NewInt(25))
	fmt.Println(properRand)

}
