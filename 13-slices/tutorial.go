package main

import "fmt"

func main() {
	// var x [5]int = [5]int{1, 2, 3, 4, 5}
	// var s []int = x[1:3]
	// // fmt.Println(cap(s))
	// fmt.Println(s[:cap(s)])

	// var a []int = []int{5, 6, 7, 8, 9}
	// fmt.Println(cap(a[:3]))

	// b := append(a, 10)
	// fmt.Println(b)
	// a = append(a, 10)
	// fmt.Println(a)
	a := make([]int, 5)
	fmt.Println(a)
}
