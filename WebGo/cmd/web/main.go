package main

import (
	"fmt"
	"net/http"

	handler "github.com/lavakush94/lets-go/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Printf("Starting Application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
