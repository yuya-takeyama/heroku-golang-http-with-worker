package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var count uint64

func countUpdater() {
	for {
		atomic.AddUint64(&count, 1)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$$PORT must be set")
	}

	go countUpdater()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Count: %d\n", count)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
