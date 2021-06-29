package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, black hat gophers!")

	collection := []int{2, 4, 8, 16}

	// 'range()' will yield the index and the value of the item.
	// This way, we can either read the value, or modify it using the
	// index and the collection selection functionality.
	for idx, val := range collection {
		fmt.Println(idx, val)
	}

	for idx, _ := range collection {
		fmt.Println(idx)
	}

	for _, val := range collection {
		fmt.Println(val)
	}
}
