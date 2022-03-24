package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pauloalexandreanjos/simple-loadbalancer/models"
)

func startServerApi() {
	mux := http.NewServeMux()

	log.Println("Starting API...")
	mux.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Received ping / sending pong")
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("/register", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(405)
			w.Write([]byte(`{"error":"Method not Allowed"}`))
			return
		}

		var register models.Register

		decoder := json.NewDecoder(req.Body)

		err := decoder.Decode(&register)
		if err != nil {
			log.Println(err)
			return
		}

		service, err := server.GetService(register.ServiceToken)
		if err != nil {
			log.Println(err)
			log.Printf("Can't register task to service %s\n", register.ServiceToken)
			return
		}

		service.Tasks = append(service.Tasks, &models.Task{})
	})

	mux.HandleFunc("/status", func(w http.ResponseWriter, req *http.Request) {

	})

	mux.HandleFunc("/nodes", func(w http.ResponseWriter, req *http.Request) {

	})

	http.ListenAndServe(":4545", mux)
}
