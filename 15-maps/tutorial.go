package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	// mp := make(map[string]int)
	// fmt.Println(mp["apples"])

	// mp["Owen"] = 900 // Add a new key-value pair

	// delete(mp, "apple") // Delete a key-value pair

	val, ok := mp["apple"]

	fmt.Println(val, ok)

	fmt.Println(mp)
}
