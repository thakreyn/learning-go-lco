/*
	Arrays

	In Golang, unlike other languages, arrays aren't used much.
	And this is by design. We use other advanced data structures.

	But still, arrays are important and can have certain use cases.

	We can either create a array and not initialise (in this case, types are initialised to
	their default values).
	Or we can also initialise by providing values.

	Arrays don't offer much functionality like sorting etc.
	For more functionality, we use a data structure called 'slice'.

	Slice is used in most cases over arrays.
*/
package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	// Declaration of arrays
	// We just put [x] before the data type
	// Where x is len of array and data type corresponds to the content
	var arr [10]int

	arr[0] = 12
	arr[6] = 101
	fmt.Printf("Type of arrays is : %T\n", arr)

	// We can directly print out the array
	fmt.Println(arr)

	// The len(Type) can be used to find length
	// For array, the length is constant and same as initialised
	fmt.Println(len(arr))

	// Doing same for string
	var string_list [5]string

	string_list[0] = "apple"
	string_list[3] = "ball"

	fmt.Println(string_list)

	// Initialising string on declaration
	// Here although we are initialising with 3 values,
	// but as the length declared is 5. The final length will be 5
	var second_list = [5]string{"1", "2", "3"}
	fmt.Println(second_list)
	fmt.Println(len(second_list))

	// Array of pointers
	// By default it is initialised with <nil>

	n1 := 12
	n2 := 24
	n3 := 36

	var pointerArray [5]*int
	pointerArray[0] = &n1
	pointerArray[1] = &n2
	pointerArray[2] = &n3

	fmt.Println(pointerArray)

	*pointerArray[0] = 2
	fmt.Println(pointerArray, *pointerArray[0])

}
