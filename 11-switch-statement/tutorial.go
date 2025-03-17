package main

import "fmt"

func main() {
	ans := 10

	// switch ans {
	// case 1, -1:
	// 	fmt.Println("1. One")
	// 	fmt.Println("2. One")
	// case 2:
	// 	fmt.Println("Two")
	// default:
	// 	fmt.Println("not a case")
	// }

	switch {
	case ans > 0:
		fmt.Println("Greater than zero")
	case ans < 0:
		fmt.Println("Less than zero")
	default:
		fmt.Println("Zero")
	}
}
