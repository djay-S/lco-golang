package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Welcome to User input"
	fmt.Println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number:")

	//comma ok, error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the input, ", input)
	fmt.Printf("Type of input is %T\n", input)

}
