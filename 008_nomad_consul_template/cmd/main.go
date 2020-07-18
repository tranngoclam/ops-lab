package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type FooBar struct {
	Foo  string `json:"foo"`
	Barz string `json:"barz"`
	Env  struct {
		Fb string `json:"fb"`
	} `json:"env"`
}

func main() {
	s, ok := os.LookupEnv("FOO_BAR")
	if !ok {
		panic(errors.New("missing env `FOO_BAR`"))
	}

	log.Println(s)
	log.Println(os.Getenv("HOSTNAME"))

	fb := FooBar{}
	err := json.Unmarshal([]byte(s), &fb)
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":5000", nil))
}
