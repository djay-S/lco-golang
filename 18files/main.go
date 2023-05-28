package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Go")
	content := "This is a file text."

	fileName := "./myLocGoFile.txt"

	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("The file length is:", length)

	readFile(fileName)
}

func readFile(fileName string) {
	byteData, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println("Raw data in the file is: ", byteData)
	fmt.Println("Text data in the file is: ", string(byteData))
}
