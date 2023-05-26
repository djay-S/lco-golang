package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	langs := make(map[string]string)

	langs["JS"] = "Javascript"
	langs["RB"] = "Ruby"
	langs["PY"] = "Python"

	fmt.Println("Map of all languages: ", langs)
	fmt.Println("Value of JS: ", langs["JS"])

	delete(langs, "JS")

	fmt.Println("Map of all languages: ", langs)


	for key, value := range(langs) {
		fmt.Printf("For key %v, value is %v.\n", key, value)
	}
}
