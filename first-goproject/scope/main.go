package main

import "fmt"

var (
	a = 10
	b = 30
)

func add(x int, y int) {
	z := x + y
	fmt.Println(z)
}
func main() {
	p := 12
	q := 15
	add(p, q)
	add(a, b)
	add(a, p)
	add(a, z)
}
