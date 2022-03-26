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
	server.MockService()
	server.MockTask()
	server.MockTask()
	server.MockTask()
	server.MockTask()
	server.MockTask()
	server.MockTask()
	server.MockTask()

	go startServerApi(server)

	for _, service := range server.Services {

		formattedPath := service.Path

		client := getHttpClient()

		log.Printf("Adding service %s at path %s", service.Name, formattedPath)
		http.HandleFunc(formattedPath, func(w http.ResponseWriter, reqSrc *http.Request) {

			task, err := service.NextTask()
			if err != nil {
				w.WriteHeader(404)
				fmt.Println(err)
				return
			}

			start := time.Now()
			req, err := http.NewRequest(reqSrc.Method, fmt.Sprintf("%s%s%s", task.Address, task.TaskPath, formattedPath), reqSrc.Body)
			if err != nil {
				w.WriteHeader(500)
				fmt.Println(err)
				return
			}
			req.Header = reqSrc.Header
			req.Host = reqSrc.Host
			req.Body = reqSrc.Body

			resp, err := client.Do(req)
			if err != nil {
				w.WriteHeader(500)
				fmt.Println(err)
				return
			}

			fmt.Println("Request took", time.Now().Sub(start))
			defer resp.Body.Close()
			// TODO Request is tooking too long(only in windows) to complete, 1ms without balancer vs 300ms with balancer for python static file server
			// this behavior isn't seen in a simple http server written in go. Must check for how python deal with this requests, maybe trying to run from a netcat
			// with http request send as a file through pipeline like [cat http_request | nc localhost 8000]
			// Another approach may be trying with standart TCP connection

			w.WriteHeader(resp.StatusCode)
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				w.WriteHeader(500)
				fmt.Println(err)
				return
			}

			log.Printf("Request received at %s with method [%s] and host [%s] and URL [%s] and cookies %s", req.RequestURI, req.Method, req.Host, req.URL, req.Cookies())
		})
	}

	testRedis()

	log.Println("Started at", time.Now().Format(dateFormat), "!", "Running...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
