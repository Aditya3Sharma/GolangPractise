package main

import (
	"fmt"
	"log"
	"net/http"
	"practise/router"
	// "practise/token"
)

func main() {
	// token.CreateToken("Aditya")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening to port 4000...")
}
