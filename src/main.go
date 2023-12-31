package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/time/rate"
)

// hard coded for clarity
// avg 1 request per second; max 3 requests per burst
var limiter = rate.NewLimiter(1, 3)

func main() {

	// explicitly setting env vars (omit if already done)
	os.Setenv("REV_PROXY_PORT", "8080")
	os.Setenv("ORIGIN_URL", "http://127.0.0.1:8081")

	// retrieve origin server URL
	originServerURL, err := url.Parse(os.Getenv("ORIGIN_URL"))
	if err != nil {
		log.Fatal("Error, invalid URL")
	}

	// reverse proxy handler object
	reverseProxy := http.HandlerFunc(func(response_writer http.ResponseWriter, request *http.Request) {

		// ensuring concurrent calls are within threshold
		if !limiter.Allow() {
			http.Error(response_writer, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			response_writer.WriteHeader(http.StatusTooManyRequests)

		} else {

			// logging request
			fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())

			// setting request to point to origin server
			request.Host = originServerURL.Host
			request.URL.Host = originServerURL.Host
			request.URL.Scheme = originServerURL.Scheme
			request.RequestURI = ""

			// send request to the origin server and save
			originServerResponse, err := http.DefaultClient.Do(request)

			// catch error with response and logging
			if err != nil {
				response_writer.WriteHeader(http.StatusInternalServerError)
				_, _ = fmt.Fprint(response_writer, err)
				return
			}

			// return response to the client
			response_writer.WriteHeader(http.StatusOK)
			io.Copy(response_writer, originServerResponse.Body)

		}

	})

	// binding reverse-proxy to desired port
	log.Fatal(http.ListenAndServe(":"+os.Getenv("REV_PROXY_PORT"), reverseProxy))
}
