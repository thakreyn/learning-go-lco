/*
	Structs:

	Golang is not a Object-Oriented (atleast completely)
	Thus there is no concept of 'Classes', 'Inheritance' etc.
	This is a design choice because inheritance hides a lot of
	things, whereas the philosophy of Go is that the working should
	be clear by seeing the code.

	So, for this, Go uses STRUCTS.

	Struct is like a custom type that is defined by the user.

*/
package main

import "fmt"

// Declaration of a struct
// We just enlist the data members
type User struct {
	Name     string
	Email    string
	Accounts []string
	Age      int
}

func main() {

	fmt.Println("Structs:")

	// There is no inheritance in Go-lang
	// So no concepts like parent, child, etc
	// This is because inheritance hides a lot of things

	// We have to allot the values while init a type struct variable
	crysys := User{"cry", "xyz@abc.co", []string{"acc1", "acc2"}, 20}

	// We can directly print our struct. In this case, only values are printed
	fmt.Println(crysys)

	// We can also access specific members
	fmt.Println(crysys.Accounts)

	// %v is used to print the value of a variable in printf
	fmt.Printf("Age is %v\n", crysys.Age)

	// %+v is used to print a struct with the name of the variable as well
	fmt.Printf("Name : %+v\nAll details : %+v\n", crysys.Name, crysys)

	// Also is not necessary to init all members of a point,
	// In those cases, the members are set to default type valus
	var test User = User{Name: "Skrinar"}
	fmt.Println(test)

	fmt.Printf("Type of test is : %T", test)

}
