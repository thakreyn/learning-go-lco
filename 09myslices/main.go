/*
	Slices:

	They are a type of data structure in Go. It is an advanced version of arrays that
	adds a layer of abstraction and a lot of functionality/features to the standard arrays.

	They are almost always used over arrays.

	Sort package:

	The 'sort' package provides a lot of useful methods for sorting different types and also
	for searching items.

	We can always refer to the documentation if we need to look into the finer details of any topic.

	** MUST READ **
	1. https://go.dev/blog/slices-intro
	2. https://gobyexample.com/slices

*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices: ")

	// 1st way of declaring slices
	// We explicitly declare the slice
	// variable called testSlice of type slice of integers
	// So basically, a slice is just an array without entering the size
	// This method inits an empty slice but does not init values
	var testSlice []int
	fmt.Println(testSlice)

	// Type of a slice is []type
	fmt.Printf("Type of slice is : %T\n", testSlice)
	fmt.Println(len(testSlice))

	// 2nd way of declaring
	// In this way, we use implicit declaration
	// Here the slice is declared as well
	var slice = []string{"a", "b", "c"}
	fmt.Println(slice)
	fmt.Printf("Type of slice is : %T\n", slice)
	fmt.Println(len(slice))

	// append is a very important method in slices
	// It allows us to append values to a slice or combine entire slices together
	slice = append(slice, "d", "e", "f")
	fmt.Println(slice)

	// Example of appending slices
	// We use a unique '...' syntax after the name of a slice
	var slice2 = []string{"x", "y", "z"}
	slice = append(slice, slice2...)
	fmt.Println(slice)

	// Just like python, we can slice the slice
	// Last range is always Non-inclusive
	fmt.Println(slice[1:3])

	// 3rd way of declaring slice
	// We can use make()
	// Make can be used to init or allocate memory for slices, maps and channels only.
	// With make, we tell the initial size of the slice (This inits the slice with default values for the types)
	highScore := make([]int, 4)
	highScore[0] = 123
	highScore[1] = 624
	highScore[2] = 125
	highScore[3] = 26

	fmt.Println(highScore)
	fmt.Printf("Type of highScore is : %T\n", highScore)

	// On appending, if the initial memory/size is exceeded, it is automatically adjusted
	highScore = append(highScore, 1023, 123123)
	fmt.Println(highScore)

	// Sort package provides methods for sorting and searching different types
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)

	fmt.Println(sort.StringsAreSorted(slice))

	// Additionally,
	// The expression []int{} -> in itself is a slice and returns a reference to this slice which can be given to a variable
	fmt.Println([]int{1, 2, 3})

	// Similarly, we can use also use the following syntax for creating a slice
	x := []int{1, 2, 3}
	fmt.Println(x)

	// What the above syntax does is that it creates an array init with the given data
	// and takes a slice equal to the entire array

	// We can use cap() to find out capacity of  a datatype
	fmt.Println(cap(x))

	// Removing values from slices based on index
	// For removing, we use the append method
	cars := []string{"merc", "beemer", "audi", "porsche", "mclaren"}
	fmt.Println(cars)

	var index int = 2

	// we use slicing to take first part and then append the remaining (use of 3 dots)
	cars = append(cars[:index], cars[index+1:]...)
	fmt.Println(cars)

}
