package main

import "fmt"

// single output return

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

// multiple output return
func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul

}

//string
func sayHello(name string) {
	fmt.Println("Hello everyone,", name)
}
func main() {
	// a := 10
	// var a int = 10
	// fmt.Println(a)
	// a := 10
	// b := 20
	// sum := add(a, b)
	// fmt.Println((sum))
	// p, q := getNumbers(a, b)
	// fmt.Println((p))
	// fmt.Println((q))
	// sayHello("Promi")
	// get user name as input
	// var name string
	// fmt.Println("Enter your name -")
	// fmt.Scanln(&name)
	// fmt.Println("----", name)
	sub(4, 7)

}
