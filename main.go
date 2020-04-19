package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	u, err := url.Parse("http://127.0.0.1:7396")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(http.ListenAndServe(":7891", httputil.NewSingleHostReverseProxy(u)))
}
