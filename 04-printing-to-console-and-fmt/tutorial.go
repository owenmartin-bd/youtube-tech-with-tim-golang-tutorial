package main

import "fmt"

func main() {
	fmt.Printf("Hello %T %v", 10, 10)
	var x string = fmt.Sprintf("Hello %T %v", 10, 10)
	fmt.Println(x)
}
