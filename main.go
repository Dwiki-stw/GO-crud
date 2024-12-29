package main

import (
	"go-product-categories/config"
	"go-product-categories/controllers/categoriescontroller"
	"go-product-categories/controllers/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// Home Page
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories Page
	http.HandleFunc("/categories", categoriescontroller.Index)

	log.Println("Server Running on Port 8080")
	http.ListenAndServe(":8080", nil)
}
