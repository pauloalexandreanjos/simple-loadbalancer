package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
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

	for _, service := range server.Services {

		formattedPath := service.Path

		client := &http.Client{}

		log.Printf("Adding service %s at path %s", service.Name, formattedPath)
		http.HandleFunc(formattedPath, func(w http.ResponseWriter, reqSrc *http.Request) {

			var task models.Task

			for _, t := range service.Tasks {
				task = t
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
			req.Header.Del("Accept-Encoding")

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

func normalizePath(path string) string {

	if path[len(path)-1] != '/' {
		log.Println("Formatou!")
		return fmt.Sprintf("%s%s", path, "/")
	}

	return path
}

func testRedis() {
	log.Println("Testing REDIS...")

	redisUrl := os.Getenv("REDIS_URL")
	redisPass := os.Getenv("REDIS_PASS")

	client := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPass,
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	err = client.Set(ctx, "name", "Elliot", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	saved, err := client.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Returned:", saved)
	log.Println("Test finished!")
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
