package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const myHost = "http://localhost:8000"

func main() {
	fmt.Println("Welcome to WebRequest in Go")

	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = myHost + "/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	var resStr strings.Builder

	bytecount, _ := resStr.Write(content)

	fmt.Println("Bytecount is:", bytecount)
	fmt.Println(string(content))
	//another way
	fmt.Println(resStr.String())
}
