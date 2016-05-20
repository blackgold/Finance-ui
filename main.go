package main

import (
	"config"
	"fmt"
	"gopkg.in/throttled/throttled.v2"
	"gopkg.in/throttled/throttled.v2/store/memstore"
	"handle"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	conf := config.LoadConfig("/etc/finance-ui/config.json")
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/{key}", handle.Api).Methods("GET")
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(conf.StaticFilePath)))).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(handle.Page404)

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
