package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition in Golang")

	waitGroup := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	var score = []int{0}

	waitGroup.Add(3)
	go func(waitGroup *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("One Goroutine")
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
		defer waitGroup.Done()
	}(waitGroup, mutex)

	go func(waitGroup *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Two Goroutine")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		defer waitGroup.Done()
	}(waitGroup, mutex)

	go func(waitGroup *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Three Goroutine")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
		defer waitGroup.Done()
	}(waitGroup, mutex)

	waitGroup.Wait()

	fmt.Println(score)
}
