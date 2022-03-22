package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	printBanner()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request received at / with method [%s] and host [%s] and URL [%s] and cookies %s", req.Method, req.Host, req.URL, req.Cookies())
	})

	testRedis()

	log.Println("Started at", time.Now().Format("02/01/2006 15:04:05"), "!", "Running...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func testRedis() {
	log.Println("Testing REDIS...")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("name", "Elliot", 0).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
	}

	saved, err := client.Get("name").Result()
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
