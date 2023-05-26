package main

import "fmt"

func main() {
	fmt.Println("Welcome to my function")

	helloFunc()

	addRes, erMsg := myAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Printf("Sumation value:%v, My message: %v.\n", addRes, erMsg)
}

func helloFunc() {
	fmt.Println("Hello Function")
}

func myAdder(values ...int) (int, string) {
	total := 0

	for _, num := range values {
		total += num
	}
	return total, "Message"
}
