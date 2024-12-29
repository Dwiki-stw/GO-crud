package main

import (
	"go-product-categories/config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	log.Println("Server Running on Port 8080")
	http.ListenAndServe(":8080", nil)
}
