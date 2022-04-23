package pkg

import (
	"github.com/simonhammes/getting-started-with-go/pkg/cache"
	"io/ioutil"
	"net/http"
	"time"
)

func (s *server) handleGetCacheItem() http.HandlerFunc {
	// thing := prepareThing()
	return func(w http.ResponseWriter, r *http.Request) {
		connection := cache.GetRedisConnection()
		result, err := connection.Get(r.Context(), "user").Result()
		if err != nil {
			http.Error(w, "Error!", 404)
			return
		}
		w.Write([]byte(result))
	}
}

func (s server) handleSetCacheItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return
		}
		defer r.Body.Close()

		connection := cache.GetRedisConnection()
		_, err = connection.SetNX(r.Context(), "user", body, 300*time.Second).Result()
		if err != nil {
			http.Error(w, "Internal server error", 500)
			return

		}
		w.Write(body)
	}
}
