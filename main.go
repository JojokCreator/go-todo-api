package main

import (
	"fmt"
	"go-todo/router"
	"log"
	"net/http"
	"os"
)

func main() {
	r := router.Router()
	fmt.Println("starting the server on port 9000....")
	port := os.Getenv("PORT")

	if port == "" {
		port = "9000"
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}
