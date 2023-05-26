package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	var myArray [3]int

	myArray[0] = 10
	myArray[2] = 12

	fmt.Println("MyArray is:", myArray)
	fmt.Println("MyArray is:", len(myArray))

	var vegArr = [3]string{"potato", "carrot", "chilly"}
	fmt.Println("My vegArr is: ", vegArr)
}
