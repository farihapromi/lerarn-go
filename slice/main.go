package main

import "fmt"

func changeSlice(p []int) []int {
	p[0] = 10
	p = append(p, 11)
	return p
}

func main() {
	// arr := [6]string{"This", "is", "go", "Interview", "Question"}
	// fmt.Println(arr)
	// s := arr[1:4]
	// fmt.Println(s)
	// s1 := s[1:2]
	// fmt.Println(s1)
	// fmt.Println(len(s1))
	// fmt.Println(cap(s))
	// s := make([]int, 3, 5)
	// s[0] = 5
	// s[3] = 20
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	//empty slice
	// var sl []int      //[] len 0 cap 0
	// s = append(sl, 1) //[1] len 1 cap 1
	// fmt.Println(sl)
	// var x []int
	// x = append(x, 1)
	// x = append(x, 2)
	// x = append(x, 3)
	// y := x
	// x = append(x, 4)
	// y = append(y, 5)
	// x[0] = 10
	// fmt.Println(x)
	// fmt.Println(y)
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)
	a := x[4:] // 4 theke shru kore last
	y := changeSlice(a)
	fmt.Println(x)
	fmt.Println(y)

}
