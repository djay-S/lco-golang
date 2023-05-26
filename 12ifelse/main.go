package main

import "fmt"

func main() {
	fmt.Println("if else in Go")

	loginCount := 69
	var result string

	if loginCount < 69 {
		result = "< 69"
	} else if loginCount == 69 {
		result = "= 69"
	} else {
		result = "> 69"
	}

	fmt.Println(result)

	if num := 9; 9%2 == 0 {
		fmt.Printf("%v is even\n", num)
	} else {
		fmt.Printf("%v is od.\n", num)
	}
}
