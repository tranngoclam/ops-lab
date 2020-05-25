package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	res "github.com/tranngoclam/ops-lab/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"os"
)

var resourceClient res.ResourceServiceClient

func ResourceHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	request := &res.ResourceID{Value: id}

	response, err := resourceClient.GetResource(context.Background(), request)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			sendJSON(w, http.StatusNotFound, map[string]interface{}{})
			return
		}

		log.Printf("get resource failed %+v", err)
		sendJSON(w, http.StatusInternalServerError, map[string]interface{}{})
		return
	}

	resource := map[string]interface{}{"id": response.Id, "name": response.Name}
	sendJSON(w, http.StatusOK, resource)
}

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	sendJSON(w, http.StatusOK, map[string]interface{}{})
}

func sendJSON(w http.ResponseWriter, status int, data map[string]interface{}) {
	body := map[string]interface{}{
		"message": http.StatusText(status),
		"data":    data,
	}
	bytes, _ := json.Marshal(body)
	w.WriteHeader(status)
	_, _ = w.Write(bytes)
}

func main() {
	port := os.Getenv("SERVER_PORT")
	address := ":" + port

	cc, err := grpc.Dial(os.Getenv("RESOURCE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func() {
		err := cc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	resourceClient = res.NewResourceServiceClient(cc)

	router := mux.NewRouter()
	router.HandleFunc("/health", HealthHandler).Methods(http.MethodGet)
	router.HandleFunc("/resource", ResourceHandler).Methods(http.MethodGet)

	log.Printf("[info] server is listening on %s", address)

	err = http.ListenAndServe(address, router)
	if err != nil {
		panic(err)
	}
}
