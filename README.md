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


### Design Decisions
---

##### Forward Proxy

The forward proxy was simple and primarily a placeholder for this process so only had the few following considerations

1. Returned a verbose response upon call 
2. logging the simple timestamp of each request

The obvious limitation of this design is the lack of state-based responses from the origin server. Though it does the job, it's not similar to an origin server in a production setting 

##### Reverse Proxy

The reverse proxy implements the **global in-flight request limit** functionality, preventing the origin to be pinged while it is serving a certain number of concurrent requests. Here is the basic functionality of the reverse proxy:

1. In normal case, it forwards the origin's message to the user
2. If origin is serving more than the allowed concurrent users, a `http.StatusTooManyRequests` status code is returned
3. The timestamp of each request is logged

**Reverse-proxy Design:**

The rate limiting logic was set up through the `x/time/rate` package which uses a [token bucket](https://en.wikipedia.org/wiki/Token_bucket) limiting algorithm; a standard algorithm in telecomm networks. The `limiter.Allow()` function, which ensures concurrent requests are below the threshold, is protected by a **mutex** which is why it's reliable during concurrent requests. This is helpful as it abstracts away the reliability of our reverse-proxy.

Additionally, the rate-limiting logic was set up in an if-else statement rather than being put into a helper package in another file. This was done for two reasons:
1. Improving code readability by preventing the need to traverse various files for relatively simple functionality   
2. Adding additional features to this workflow, such as the 'sharded rate-limiting' or 'retries', could be added by chaining that code together together in additional else-if blocks, easing the time taken to integrate newer functions while keeping the process flow easy to understand


### Scaling the Server
----

To provide effective scalability, a mixture of battle-tested techniques for managing large-scale workloads and through workflows that allow engineers to pragmaticlly diagnose flaws ahead of time.

##### Load Balancing
As concurrent requests may need to be exceedingly high for some servers, a load balancing mechanism can be introduced to split workloads over multiple origin servers. This is to prevent overload on one server by evenly allocating requests throughout origin server infrastructure, to optimize for performance and reliability. In addition, failiure to respond for one origin server can simply mean requests are dynamically re-forwarded to origin servers that are still working

##### Caching

Depending on the server's usecase, several responses can come from similar or identical requests. Therefore, a caching mechanism could help our service perform reliably at scale by ensuring identical workloads aren't repeatedly processed

##### Monitoring

Setting up visibility services like Grafana and disaster-monitoring software like PagerDuty can help our infra teams stay ahead of the curve, while allowing regular engineers to easily monitor the state of the production system. This will allow them to pick up anomolous patterns in performance and fix them before a large-scale break


### Improving Server Security
---
Improving security can be done through following best-practices common in the space. One such way is through introducing encryption.

As large-scale services can expect to be targetted by attacks, the reverse-proxy can be in-charge of encrypting and decrypting messages to and from the origin server. This adds a layer of security to ensure our origin is protected and can reliably perform despite the threats
