package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	address := ":51001"

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	log.Println("[info]", "server is listening on", address)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[info]", "server received request")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(hostname))
	})

	err = http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
