package main

import "fmt"

func main() {
	// var arr [5]int

	// arr[0] = 100
	// arr[4] = 900

	// fmt.Println(arr[0])

	arr := [3]int{4, 5, 6}
	fmt.Println(len(arr))

	sum := 0

	arr2D := [2][3]int{{1, 2, 3}, {3, 3, 4}}
	fmt.Println(arr2D[0][2])

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)
}
