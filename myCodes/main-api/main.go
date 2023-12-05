// main_gateway/main.go

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	// Basic Auth middleware
	r.Use(basicAuthentification)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Main Gateway API")
	})

	http.ListenAndServe(":8080", r)
}

// Basic authentification middleware via "net/http"
func basicAuthentification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		// The environment variables are in docker-compose.yml file
		expectedUsername := os.Getenv("USERNAME")
		expectedPassword := os.Getenv("PASSWORD")

		// If the given username and password do not match with the environment variables
		if !ok || !(username == expectedUsername && password == expectedPassword) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the given username and password do match with the environment variables
		next.ServeHTTP(w, r)
	})
}
