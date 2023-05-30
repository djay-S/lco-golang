package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=edfgb5vdghnb"

func main() {
	fmt.Println("Welcome to Urls")
	fmt.Println("url", myUrl)

	//parsing the url
	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of the qparams is %T\n", qparams)

	fmt.Println("Course name: ", qparams["coursename"])
	fmt.Println("Payment Info: ", qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "myWebsite",
		Path:     "/myroute",
		RawQuery: "user=admin",
	}

	createdUrl := partsOfUrl.String()

	fmt.Println("The custom created Url is: ", createdUrl)
}
