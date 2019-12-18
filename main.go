package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request to /")
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, hostname)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// log.Println("Got request to /health")
		fmt.Fprintf(w, "ok")
	})
	http.ListenAndServe(":9000", nil)
}
