package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	port, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		port = "8000"
	}
	address := ":" + port

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())

	log.Println("[info]", "server is listening on", address)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[info]", "server received request")

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(hostname))
	})

	err = http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
