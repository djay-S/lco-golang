package main

import "fmt"

func main() {
	fmt.Println("Welcome to my pointers")

	var ptr *int

	fmt.Println("Value of the pointer is:", ptr)

	myNumber := 24

	ptr = &myNumber

	fmt.Println("Value of the pointer is:", ptr)
	fmt.Println("Stored value of the pointer is:", *ptr)
}
