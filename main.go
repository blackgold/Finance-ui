package main

import (
	"config"
	"fmt"
	"gopkg.in/throttled/throttled.v2"
	"gopkg.in/throttled/throttled.v2/store/memstore"
	"handle"
	"log"
	"net/http"
)

func main() {
	conf := config.LoadConfig("/etc/finance-ui/config.json")
	r := http.NewServeMux()
	r.HandleFunc("/api/v1/info", handle.Api)
	r.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/Users/devendra/personal/finance-ui/static"))))
	store, err := memstore.New(65536)
	if err != nil {
		log.Fatal(err)
	}
	// Maximum burst of 5 which refills at 20 tokens per minute.
	quota := throttled.RateQuota{throttled.PerMin(1000), 100}
	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
	if err != nil {
		log.Fatal(err)
	}
	httpRateLimiter := throttled.HTTPRateLimiter{
		RateLimiter: rateLimiter,
	}
	http.Handle("/", httpRateLimiter.RateLimit(r))
	listenAddr := fmt.Sprintf(":%d", conf.HTTPPort)
	http.ListenAndServe(listenAddr, nil)
}