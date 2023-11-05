# Reverse Proxy

This git repository houses an origin server along with a basic reverse-proxy that offers global in-flight rate limiting. This ensures the origin server isn't overloaded with requests



### Installation and Usage
---

##### Part 1: Cloning github repo
setting up project
```bash
$ git clone https://github.com/taha5322/ProxyServer.git
```

##### Part 2: Set environment variables
Set reverse-proxy port, and origin-server endpoint that you will access.
**Note**: this is done explicitly in the script to the commented values for visibility
```bash
$ export REV_PROXY_PORT=YOUR_CHOSEN_PORT # 8080
$ export ORIGIN_URL=YOUR_ORIGIN_SERVER_URL # http://127.0.0.1:8081
```

##### Part 3: Start origin server

This can be any server you made or an existing service for which this reverse-proxy acts as middleware for. This is the server with the endpoing which you set to `ORIGIN_URL`

**(Optional)** For demo purposes, a mock origin server is provided and can be run with:
```bash
$ cd /helper/
$ go run origin.go
```
##### Part 4: Start reverse-proxy server
**note:** ensure go.mod file exists in project root folder. If not, run `go mod tidy` in /src
```bash
$ cd /src
$ go run .
```
