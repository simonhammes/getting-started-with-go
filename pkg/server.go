package pkg

import (
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	// db     *someDatabase
	router *mux.Router
	// cache
}

func NewServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}

	// Setup routing
	s.routes()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
