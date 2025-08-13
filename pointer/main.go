package main

import "fmt"

// func print(numbers [3]int) {
// 	fmt.Println(numbers)

// }
//pointer
func print(numbers *[3]int) {
	fmt.Println(numbers)

}

type User struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	obj := User{
		Name:   "Promi",
		Age:    29,
		Salary: 40000,
	}
	fmt.Println((obj))
}

// func main() {
// 	// x := 20
// 	// p := &x
// 	// fmt.Println("The value of x is", *p)
// 	arr := [3]int{1, 2, 3}
// 	// print(arr)
// 	print(&arr)
// }
