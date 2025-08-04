package main

import "fmt"

const a = 30

var p = 100

func outer() func() {
	money := 100
	age := 27
	fmt.Println("AGE:", age)
	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}
func call() {
	incr1 := outer()
	incr1()
	incr1()
	incr2 := outer()
	incr2()
	incr2()

}
func init() {
	fmt.Println("=======BANK==========")
}
