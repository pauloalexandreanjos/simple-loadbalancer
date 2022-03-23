package main

import (
	"net"
	"net/http"
	"time"
)

func getHttpClient() *http.Client {

	dialer := &net.Dialer{
		Timeout:   time.Second * 5,
		KeepAlive: time.Second * 30,
	}

	transport := &http.Transport{
		Dial:                dialer.Dial,
		TLSHandshakeTimeout: time.Second * 5,
		DisableCompression:  true,
	}

	client := &http.Client{
		Timeout:   time.Second * 10, // Request timeout 10 seconds
		Transport: transport,
	}
	return client
}
