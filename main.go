package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/pauloalexandreanjos/simple-loadbalancer/models"
)

var ctx = context.Background()

const dateFormat = "02/01/2006 15:04:05"

var server *models.Server

func main() {

	godotenv.Load()

	printBanner()

	server = models.NewServer("My Simple Loadbalancer")
	server.MockServer()

	go startServerApi()

	for _, service := range server.Services {

		formattedPath := service.Path

		client := getHttpClient()

		log.Printf("Adding service %s at path %s", service.Name, formattedPath)
		http.HandleFunc(formattedPath, func(w http.ResponseWriter, reqSrc *http.Request) {

			task, err := getTask(service.Token)
			if err != nil {
				log.Println(err)
			}

			start := time.Now()
			req, err := http.NewRequest(reqSrc.Method, fmt.Sprintf("%s%s", task.Address, formattedPath), reqSrc.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header = reqSrc.Header
			req.Host = reqSrc.Host
			req.Body = reqSrc.Body

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Request took", time.Now().Sub(start))
			defer resp.Body.Close()
			// TODO Requesting is tooking so long, 1ms without balancer vs 300ms with balancer
			// Maybe trying with standart TCP connection ou customizing http client

			_, err = io.Copy(w, resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Request received at %s with method [%s] and host [%s] and URL [%s] and cookies %s", req.RequestURI, req.Method, req.Host, req.URL, req.Cookies())
		})
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	log.Printf("Request received at / with method [%s] and host [%s] and URL [%s] and cookies %s", req.Method, req.Host, req.URL, req.Cookies())
	// })

	// http.HandleFunc("/teste/", func(w http.ResponseWriter, req *http.Request) {
	// 	log.Printf("Request received at /teste with method [%s] and host [%s] and URL [%s] and cookies %s", req.Method, req.Host, req.URL, req.Cookies())
	// })

	testRedis()

	log.Println("Started at", time.Now().Format(dateFormat), "!", "Running...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
