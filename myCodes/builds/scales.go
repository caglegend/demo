package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/go-chi/chi"
	"golang.org/x/sync/semaphore"
)

var (
	requestCount int
	mutex        sync.Mutex
	sem          *semaphore.Weighted
)

func main() {
	r := createRouter()
	port := getPort()

	// Aynı anda kaç işlemin çalışabileceğini belirten semafor
	sem = semaphore.NewWeighted(5)

	log.Printf("Server is starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func createRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		increaseRequestCount()
		w.Write([]byte("Hello, World!"))
	})

	r.Get("/manualScale", func(w http.ResponseWriter, r *http.Request) {
		// Bu endpoint, kullanıcının ölçeklendirme taleplerini karşılar.
		// Kullanıcılar buradan ölçeklendirme parametrelerini iletebilir.
		scaleFunctionManually(w, r)
	})

	return r
}

func increaseRequestCount() {
	mutex.Lock()
	defer mutex.Unlock()
	requestCount++
	fmt.Printf("Request Count: %d\n", requestCount)
}

func scaleFunctionManually(w http.ResponseWriter, r *http.Request) {
	// Aynı anda kaç işlemin çalışabileceğini kontrol et
	if err := sem.Acquire(r.Context(), 1); err != nil {
		http.Error(w, "Too many concurrent requests", http.StatusTooManyRequests)
		return
	}
	defer sem.Release(1)

	query := r.URL.Query()
	desiredCountStr := query.Get("count")

	desiredCount, err := strconv.Atoi(desiredCountStr)
	if err != nil {
		http.Error(w, "Invalid count parameter", http.StatusBadRequest)
		return
	}

	// Burada, isteğe bağlı olarak kullanıcı tarafından belirtilen sayıda ölçeklendirme yapabilirsiniz.
	// Örneğin, bir konteyner orkestrasyon aracı (Docker Swarm, Kubernetes) kullanarak.

	fmt.Printf("Manually Scaling to %d instances...\n", desiredCount)

	w.Write([]byte(fmt.Sprintf("Manually Scaling to %d instances...\n", desiredCount)))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
