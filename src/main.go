package main

import (
	"log"
	"net/url"
	"os"
)

func main() {

	// explicitly setting env vars (omit if already done)
	os.Setenv("REV_PROXY_PORT", "8080")
	os.Setenv("ORIGIN_URL", "http://127.0.0.1:8081")

	// retrieve origin server URL
	originServerURL, err := url.Parse(os.Getenv("ORIGIN_URL"))
	if err != nil {
		log.Fatal("Error, invalid URL")
	}

}
