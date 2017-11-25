package main

import (
	"log"
	"net/http"
	"time"

	"github.com/otknoy/av_search_be/cache"
)

var port = "8080"

func main() {
	handler := &Handler{}
	handler.Cache = cache.NewSimpleCacheRepository(24*60*time.Minute, 3*24*60*time.Minute)

	http.HandleFunc("/search", handler.Search)

	log.Print("start server: port=" + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
