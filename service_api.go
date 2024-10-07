package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/pauloalexandreanjos/simple-loadbalancer/models"
)

type Register struct {
	ServiceToken string `json:"serviceToken"`
	ServicePort  string `json:"servicePort"`
	Schema       string `json:"schema"`
	HealthUrl    string `json:"healthUrl"`
}

const ADDR = ":4545"

func startServerApi(server *models.Server) {
	mux := http.NewServeMux()

	log.Printf("Starting API on address %s...\n", ADDR)
	mux.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Received ping / sending pong")
		w.Write([]byte("pong"))
	})

	mux.HandleFunc("/register", handleRegister)
	mux.HandleFunc("/status", handleStatus)
	mux.HandleFunc("/nodes", handleNodes)

	http.ListenAndServe(ADDR, mux)
}

func handleRegister(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(405)
		w.Write([]byte(`{"error":"Method not Allowed"}`))
		return
	}

	var register Register

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

	service.AddTask(&models.Task{
		ServiceToken: register.ServiceToken,
		Id:           strconv.Itoa(rand.Int()),
		Address:      formatAddress(req.RemoteAddr, register.ServicePort, register.Schema),
	})
}

func handleStatus(w http.ResponseWriter, req *http.Request) {

}

func handleNodes(w http.ResponseWriter, req *http.Request) {

	result, err := json.Marshal(server.Services)
	if err != nil {
		fmt.Fprintf(w, "Erro, não é possivel mostrar os serviços")
		return
	}

	fmt.Fprint(w, string(result))
}

func formatAddress(remoteAddress string, port string, schema string) string {
	address := strings.Split(remoteAddress, ":")[0]
	if port != "" {
		return fmt.Sprintf("%s://%s:%s", schema, address, port)
	}

	return fmt.Sprintf("%s://%s", schema, address)
}
