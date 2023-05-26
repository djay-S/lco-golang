package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time file")

	presentTime := time.Now()

	fmt.Println("Current time is: ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdTimeStamp := time.Date(2000, time.April, 1, 4, 20, 6, 9, time.Now().Location())

	fmt.Println("Created TimeStamp:", createdTimeStamp)
	fmt.Println(createdTimeStamp.Format("01-02-2006 Monday 15:04:05"))
}
