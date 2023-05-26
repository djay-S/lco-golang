package main

import "fmt"

func main() {
	fmt.Println("Methods in Go")
	// no inheritence in Golang; No super or parent

	adminUser := User{"Admin", "admin@mail.com", false, 69}

	fmt.Println(adminUser)

	fmt.Printf("Admin user details: %+v\n", adminUser)
	fmt.Printf("Admin user Name: %v, email:%v\n", adminUser.Name, adminUser.Email)

	adminUser.GetStatus()
	adminUser.NewEmail()
	fmt.Printf("Admin user Name: %v, email:%v\n", adminUser.Name, adminUser.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is updated to:", u.Email)
}
