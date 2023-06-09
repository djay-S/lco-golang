package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var waitGroup sync.WaitGroup // these are pointers
var mutex sync.Mutex         // these are pointers
// Mutex provide Locks to the memory associated with a variable

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://google.com",
		"https://focalpost.com",
		"https://connectMe.com",
		"https://github.com",
		"https://fb.com",
	}

	fmt.Println("Function execution without goo routine:")
	for _, site := range websiteList {
		getStatusCodeWithoutGoRoutine(site)
	}

	fmt.Println("\nFunction execution with Goroutine")
	for _, site := range websiteList {
		go getStatusCodeWithGoRoutine(site)

		waitGroup.Add(1)
	}

	waitGroup.Wait()

	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCodeWithGoRoutine(endpoint string) {

	defer waitGroup.Done()

	result, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	mutex.Lock()
	signals = append(signals, endpoint)
	mutex.Unlock()

	fmt.Printf("%d status code for the endpoint: %s\n", result.StatusCode, endpoint)
}

func getStatusCodeWithoutGoRoutine(endpoint string) {
	result, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d status code for the endpoint: %s\n", result.StatusCode, endpoint)
}
