package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to My Slices")

	var fruitSlice = []string{"apple", "mango", "peach"}
	fmt.Printf("The type of fruitSlice is: %T.\n", fruitSlice)
	fmt.Println("The fruitSlice:", fruitSlice)

	fruitSlice = append(fruitSlice, "melon", "banana")

	fmt.Println("The fruitSlice:", fruitSlice)

	fruitSlice = append(fruitSlice[1:])
	fmt.Println("The fruitSlice:", fruitSlice)

	fruitSlice = append(fruitSlice[1:3])
	fmt.Println("The fruitSlice:", fruitSlice)

	fruitSlice = append(fruitSlice[1:3])
	fmt.Println("The fruitSlice:", fruitSlice)

	fruitSlice = append(fruitSlice[:3])
	fmt.Println("The fruitSlice:", fruitSlice)

	// fruitSlice = append(fruitSlice[1:10])
	// fmt.Println("The fruitSlice:", fruitSlice)

	scores := make([]int, 4)

	scores[0] = 54
	scores[1] = 420
	scores[2] = 69
	scores[3] = 254
	// scores[4] = 000

	scores = append(scores, 111, 222, 333)

	fmt.Println(scores)

	sort.Ints(scores)

	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	fmt.Println(sort.Reverse(sort.IntSlice(scores)))
	fmt.Println(sort.IntsAreSorted(scores))

	var courses = []string{"reactjs", "java", "python", "go"}
	fmt.Println(courses)

	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
