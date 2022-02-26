package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":51001"

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	log.Println("[info]", "server is listening on", addr)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(http.StatusText(http.StatusOK)))
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[info]", "server received request")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(hostname))
	})

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
