/*
	Methods :

	As Go doesn't have Classes, thus to add some functionality to structs,
	methods are used.

	They are very similar to functions, but have a slightly different signature

	signature -> func (<struct_name)> <struct_type>) <func_name> () {}

	We give the method a struct to associate with so that that function becomes a
	part of the struct.

	Also if we want to change the values, we have to pass a pointer to the struct
*/
package main

import "fmt"

func main() {
	fmt.Println("Methods : ")

	// init the struct instance
	user1 := User{"CRY", "asda@kl.cas", []string{"acc1", "acc2"}, 12}
	fmt.Println(user1)

	// Using a method that takes reference
	user1.ChangeAge()

	// Using normal function
	AddAge(user1)
	fmt.Println(user1)

	// Using a method with pointer to struct
	user1.addAccount()
	fmt.Println(user1)

}

// Using the user struct
type User struct {
	Name     string
	Email    string
	Accounts []string
	Age      int
}

// This is a method for struct User
// It changes the age to 13
// But as only a reference is passed, it creates a copy of the struct
func (u User) ChangeAge() {
	u.Age = 13
	fmt.Println("Age is : ", u.Age)
}

// The following method adds an account to the User.Accounts
// It takes a pointer to a suer struct, so the changes last
func (u *User) addAccount() {
	u.Accounts = append(u.Accounts, "newacc")
	fmt.Println(u.Accounts)
}

// A normal function to add age
func AddAge(u User) {
	u.Age++
	fmt.Println(u)
}
