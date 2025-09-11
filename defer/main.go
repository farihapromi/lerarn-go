package main

import "fmt"

// func a() {
// 	i := 0
// 	fmt.Println("first ", i)

// 	defer fmt.Println("second ", i)
// 	i = i + 1
// 	fmt.Println("Third", i)
// 	return
// }
func calculate() (result int) {
	fmt.Println("first", result)
	show := func() {
		result = result + 10
		fmt.Println("defer", result)

	}
	defer show()
	result = 5
	fmt.Println("second", result)
	return
}
func calc() int {
	result := 0
	fmt.Println("first", result)
	show := func() {
		result = result + 10
		fmt.Println("defer", result)

	}
	defer show()
	result = 5
	fmt.Println("second", result)
	return result
}
func main() {
	// a()
	a := calculate()
	fmt.Println("main first", a)
	b := calc()
	fmt.Println("main second", b)

}
