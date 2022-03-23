package main

import (
	"log"
	"net/http"
)

func startServerApi() {
	mux := http.NewServeMux()

	log.Println("Starting API...")
	mux.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Received ping / sending pong")
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("/register", func(w http.ResponseWriter, req *http.Request) {

	})

	mux.HandleFunc("/status", func(w http.ResponseWriter, req *http.Request) {

	})

	mux.HandleFunc("/nodes", func(w http.ResponseWriter, req *http.Request) {

	})

	http.ListenAndServe(":4545", mux)
}
