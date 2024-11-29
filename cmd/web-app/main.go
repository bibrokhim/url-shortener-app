package main

import (
	"github.com/bibrokhim/url-shortener-app/internal/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.ShowIndex)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
