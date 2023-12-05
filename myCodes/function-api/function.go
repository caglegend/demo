// function_api/main.go

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Log endpoint for function API")
	// Log handling logic goes here
}

func main() {
	r := chi.NewRouter()

	// Basic Auth middleware
	r.Use(basicAuthentification)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Function API")
	})

	// Log endpoint'i
	r.Get("/log", logHandler)

	http.ListenAndServe(":8081", r)
}

func basicAuthentification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		expectedUsername := os.Getenv("USERNAME")
		expectedPassword := os.Getenv("PASSWORD")

		if !ok || !(username == expectedUsername && password == expectedPassword) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
