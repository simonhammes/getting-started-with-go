package main

import (
	"github.com/simonhammes/getting-started-with-go/pkg"
	"log"
	"net/http"
)

/* type server struct {
	// db     *someDatabase
	router *mux.Router
	// cache
}*/

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	s := pkg.NewServer()

	err := http.ListenAndServe(":8080", s)

	return err
}
