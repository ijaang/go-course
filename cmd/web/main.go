package main

import (
	"firstWebApp/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("starting web on port ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
