package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		// Oddiy metrikaning demo varianti
		now := time.Now().Unix()
		fmt.Fprintf(w, "app_up 1\napp_time %d\n", now)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from ", r.Host)
	})

	port := ":8080"
	fmt.Println("Server running on", port)
	http.ListenAndServe(port, nil)
}
