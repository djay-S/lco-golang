package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	// no inheritence in Golang; No super or parent

	adminUser := User{"Admin", "admin@mail.com", false, 69}

	fmt.Println(adminUser)

	fmt.Printf("Admin user details: %+v\n", adminUser)
	fmt.Printf("Admin user Name: %v, email:%v\n", adminUser.Name, adminUser.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
