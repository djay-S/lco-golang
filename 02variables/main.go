package main

import "fmt"

// globalVar := "globalVar" // doesn't work walrus only allowed inside of a method
// var globalVar = "globalVar" // works
const LoginToken string = "qsrtghio;]75" // since the first letter is CAPS this is public

func main() {
	var userName string = "name"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	var loggedIn bool = false
	fmt.Println(loggedIn)
	fmt.Printf("Variable is of type: %T \n", loggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat32 float32 = 123.01234567890123456789
	fmt.Println(smallFloat32)
	fmt.Printf("Variable is of type: %T \n", smallFloat32)

	var smallFloat64 float64 = 123.01234567890123456789
	fmt.Println(smallFloat64)
	fmt.Printf("Variable is of type: %T \n", smallFloat64)

	//default values and aliases
	var anotherVariable int //not int8 or any thing else
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "www.github.com"
	fmt.Println(website)
	// website = 69 // -> doesn't work
	fmt.Printf("Variable is of type: %T \n", website)

	//no var style
	numberOfUsers := 255 //-> := walrus operator
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
