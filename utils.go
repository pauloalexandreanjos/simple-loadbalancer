package main

import (
	"fmt"
	"log"
)

func normalizePath(path string) string {

	if path[len(path)-1] != '/' {
		log.Println("Formatou!")
		return fmt.Sprintf("%s%s", path, "/")
	}

	return path
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
