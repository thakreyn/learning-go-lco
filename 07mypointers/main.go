/*
	Pointers

	In Go, just like in C/C++, pointers are used to store addresses
	of variables. Thus instead of passing around variables or objects themselves, we rather
	pass pointers that point to those variables. This prevents uneccesary copying of
	variables.

	In Go, pointers behave just like in C/C++. They store addresses.
	Also we can use '&' (reference) and '*' (dereference).

		& -> gives the address/reference of a variable.
		* -> gives the value stored at a particular address. (Used on pointers)

	Thus pointers allow us to pass arounf the memory location or addresses of variables,
	so that we can make changes directly to that memory.

	Following is a good resource that covers almost everything related to pointers and how they
	can be used.
		https://youtu.be/sTFJtxJXkaY
*/

package main

import "fmt"

func main() {

	// Declare a pointer without init
	// It can be read as:
	// Declaring a variable called 'ptr' that is a pointer to an integer type.
	var ptr *int
	var num int = 256

	// We assign a value to a pointer by giving the address of the variable
	ptr = &num

	// Prints address of num or &num
	fmt.Println("ptr = ", ptr)
	fmt.Println("Value in num/*ptr : ", *ptr)

	num2 := 24.6

	// Here, we see implicit declaration of a pointer.
	// When we pass address to a variable. It automatically becomes a pointer
	// In this case, ptr2 becomes of type *float64
	var ptr2 = &num2

	fmt.Println("ptr2 = ", ptr2, "\n*ptr2 + 2 or num2 + 2 = ", *ptr2+2.0)
	fmt.Printf("Type of ptr2 is : %T\n", ptr2)

	// A pointer can store address (point to) a pointer
	// Thus we can make changes to a variable my nesting pointers
	var ptr3 = &ptr

	fmt.Println("ptr3 = ", ptr3)
	fmt.Printf("Type of ptr3 = %T\n", ptr3)

	// So we can change the contents of num using ptr3 in the following way
	fmt.Println(*(*ptr3) * 2)

	// Here, strPtr implicitly becomes of type *string. ie: Pointer to a string
	var testString = "Hello"
	var strPtr = &testString
	fmt.Printf("Type of strptr is : %T", strPtr)
}
