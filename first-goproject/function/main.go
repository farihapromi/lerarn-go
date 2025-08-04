package main

import "fmt"

// func main() {
// 	add := func(x int, y int) {
// 		c := x + y
// 		fmt.Println((c))
// 	}
// 	add(2, 3)
// }

//Higher order fucntion
func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}

//First Order function
func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}
func main() {
	processOperation(3, 4, add)
}
