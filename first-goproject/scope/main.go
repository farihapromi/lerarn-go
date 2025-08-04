package main

import "fmt"

var (
	a = 10
	b = 30
)

func print(num int) {
	fmt.Println((num))

}

func add(x int, y int) {
	z := x + y
	print(z)
}
func main() {
	// p := 12
	// q := 15
	// add(p, q)
	// add(a, b)
	// add(a, p)
	// add(a, z)
	add(5, 6)
}
