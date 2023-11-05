package main

import (
	"os"
)

func main() {

	// explicitly setting env vars (omit if already done)
	os.Setenv("REV_PROXY_PORT", "8080")
	os.Setenv("ORIGIN_URL", "http://127.0.0.1:8081")

}
