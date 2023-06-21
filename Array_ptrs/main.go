package main

import "fmt"

func main() {
	// Creating an array of integers
	arr := [3]int{1, 2, 3}

	// Creating a pointer to the array
	ptr := &arr

	// Accessing array elements using the pointer
	fmt.Println((*ptr)[0]) // Output: 1
	fmt.Println((*ptr)[1]) // Output: 2
	fmt.Println((*ptr)[2]) // Output: 3

	// Modifying array elements using the pointer
	(*ptr)[0] = 4
	fmt.Println(arr) // Output: [4 2 3]

	// Creating a pointer to a specific array element
	elemPtr := &arr[1]

	// Accessing and modifying the array element using the pointer
	fmt.Println(*elemPtr) // Output: 2
	*elemPtr = 5
	fmt.Println(arr) // Output: [4 5 3]
}
