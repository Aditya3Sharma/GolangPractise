package main

import (
	"fmt"
	"log"
	"net/http"
	"practise/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening to port 4000...")
}
