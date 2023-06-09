package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"url"`
	Password string   `json:"-"`              // -> "-" skips the encoding of the attribute
	Tags     []string `json:"tags,omitempty"` // -> omitempty skips the json attribute if the value is nil
}

func main() {
	fmt.Println("Welcome to Create JSON in Go")

	EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {

	myCourses := []course{
		{"Reactjs Bootcamp", 69, "MeraWebsite", "~1igM@-u-", []string{"webdev", "js"}},
		{"MERN Bootcamp", 420, "MeraWebsite", "^d33z`", []string{"fullstack", "js"}},
		{"Java Bootcamp", 69420, "MeraWebsite", "NuT$!!oo", nil},
	}

	//package this data in JSON format

	finalJSON, err := json.Marshal(myCourses)

	if err != nil {
		panic(err)
	}

	fmt.Println("Encoded JSON:", finalJSON) // ->gives raw data
	fmt.Println("\nEncoded JSON by Println:", string(finalJSON))
	fmt.Printf("\nEncoded JSON by Printf%s\n:", finalJSON)

	jSONIndented, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%s\n", jSONIndented)
}

func DecodeJSON() {
	fmt.Println("Decode JSON")

	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var responseCourse course

	isValid := json.Valid(jsonDataFromWeb)

	if isValid {
		fmt.Println("Json was valid")

		json.Unmarshal(jsonDataFromWeb, &responseCourse)

		fmt.Printf("%#v\n", responseCourse)

	} else {
		fmt.Println("Json was not valid")
	}

	// some case where we want to add data to key value
	var myData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myData)

	fmt.Printf("%#v \n", myData)

	for k, v := range myData {
		fmt.Printf("Key is: %v, and Value is: %v and the Type is %T\n", k, v, v)
	}
}
