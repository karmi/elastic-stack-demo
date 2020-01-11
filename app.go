package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	listenAddr = "0.0.0.0:8000"
)

func main() {
	log.SetFlags(0)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello\n")
	})

	log.Printf("Server starting at %s...", listenAddr)
	if err := http.ListenAndServe(listenAddr, nil); err != nil && err != http.ErrServerClosed {
		log.Fatal("Unable to start server")
		os.Exit(1)
	}
}