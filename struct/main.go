package main

import "fmt"

type User struct {
	name string
	Age  int
}

// func printUserDetails(usr User) {
// 	fmt.Println("Name", usr.name)
// 	fmt.Println("Age", usr.Age)
// }
// reciever function
func (usr User) printDetails() {
	fmt.Println("Name", usr.name)
	fmt.Println("Age", usr.Age)
}

func main() {
	var user1 User
	user1 = User{
		name: "promi",
		Age:  23,
	}
	user1.printDetails()
	// printUserDetails(user1)
	// fmt.Println("Name", user1.name)
	// fmt.Println("Age", user1.Age)
	user2 := User{
		name: "Seher",
		Age:  29,
	}
	// printUserDetails((user2))
	user2.printDetails() //reciver function
}
