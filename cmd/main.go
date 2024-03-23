package main

import (
	"log"
	"net/http"

	"github.com/simonhammes/getting-started-with-go/pkg/handlers"
)

func main() {
	port := ":8080"

	http.HandleFunc("GET /users", handlers.GetUser)
	http.HandleFunc("POST /users", handlers.SetUser)

	log.Fatal(http.ListenAndServe(port, nil))
}
