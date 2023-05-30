package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to LCO WebRequest")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() //caller's responsibility to close the connection

	fmt.Printf("The type of the response is: %T\n", response)

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)

	fmt.Println("Website content:", content)
}
