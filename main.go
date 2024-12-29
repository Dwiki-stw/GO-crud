package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server Running on Port 8080")
	http.ListenAndServe(":8080", nil)
}
