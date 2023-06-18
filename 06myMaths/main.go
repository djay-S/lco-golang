package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to Maths in Golang")

	//random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// random from crypto
	myRandomNumber, err := rand.Int(rand.Reader, big.NewInt(5))

	if err != nil {
		panic(err)
	}

	fmt.Println(myRandomNumber)
}
