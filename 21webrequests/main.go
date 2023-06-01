package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const myHost = "http://localhost:8000"

func main() {
	fmt.Println("Welcome to WebRequest in Go")

	PerformGetRequest()
	PerformCorrectPostJsonRequest()
	PerformInCorrectPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	fmt.Println("Handling get requests")

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

func PerformCorrectPostJsonRequest() {
	fmt.Println("Handling Post requests")

	const myUrl = myHost + "/post"

	correctRequestBody := strings.NewReader(`
		{
			"courseName": "Go Tutorial",
			"price": 69420,
			"platform": "localhost:6942"
		}
	`)

	response, err := http.Post(myUrl, "application/json", correctRequestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformInCorrectPostJsonRequest() {
	fmt.Println("Handling Incorrect Post requests")

	const myUrl = myHost + "/post"

	correctRequestBody := strings.NewReader(`
		{
			"courseName": "Go Tutorial",
			"price": 69420
			"platform": "localhost:6942"
		}
	`)

	response, err := http.Post(myUrl, "application/json", correctRequestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	fmt.Println("Handling Post Form requests")

	const myUrl = myHost + "/postform"

	//formdata
	data := url.Values{}

	data.Add("firstName", "D")
	data.Add("lastName", "jay")
	// data.Add("phone", 69420)	// -> cannot use 69420 (untyped int constant) as string value in argument to data.Add
	data.Add("phone", "69420")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Post Form Response:", string(content))
}
