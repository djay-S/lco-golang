package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/djay-S/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is starting.........")
	router := router.Router()
	log.Fatal(http.ListenAndServe(":4000", router))

	fmt.Println("Listening on port: 4000")
}
