package main

import (
	"go-product-categories/config"
	"go-product-categories/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//homepage
	http.HandleFunc("/", homecontroller.Welcome)

	log.Println("Server Running on Port 8080")
	http.ListenAndServe(":8080", nil)
}
