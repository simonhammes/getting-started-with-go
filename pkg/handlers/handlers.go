package handlers

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/simonhammes/getting-started-with-go/pkg/cache"
)

func UsersHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		connection := cache.GetRedisConnection()
		result, err := connection.Get(request.Context(), "user").Result()
		if err != nil {
			http.Error(writer, "Error!", 404)
			return
		}
		writer.Write([]byte(result))
	case "POST":
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Internal server error", 500)
			return
		}
		defer request.Body.Close()

		connection := cache.GetRedisConnection()
		_, err = connection.SetNX(request.Context(), "user", body, 5*time.Second).Result()
		if err != nil {
			http.Error(writer, "Internal server error", 500)
			return

		}
		writer.Write(body)

	}

}
