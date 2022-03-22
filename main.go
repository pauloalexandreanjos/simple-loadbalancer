package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	printBanner()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Recebeu request no endpoint / method %s, host %s, cookies %s", req.Method, req.Host, req.Cookies())
	})

	log.Println("Started at", time.Now().Format("02/01/2006 15:04:05"), "!", "Running...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func printBanner() {
	log.Println(`Starting...`)
	log.Println(` __ _                 _          __    ___`)
	log.Println(`/ _(_)_ __ ___  _ __ | | ___    / /   / __\`)
	log.Println(`\ \| | '_ ' _ \| '_ \| |/ _ \  / /   /__\//`)
	log.Println(`_\ \ | | | | | | |_) | |  __/ / /___/ \/  \`)
	log.Println(`\__/_|_| |_| |_| .__/|_|\___| \____/\_____/`)
	log.Println(`               |_|                        `)
}
