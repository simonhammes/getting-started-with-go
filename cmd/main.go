package main

import (
	"log"
	"net/http"

	"github.com/simonhammes/getting-started-with-go/pkg/handlers"
)

func main() {
	port := ":8080"
	http.HandleFunc("/users", handlers.UsersHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
