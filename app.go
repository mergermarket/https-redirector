package main

import (
	"log"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	target := *req.URL
	target.Scheme = "https"
	target.Host = req.Host
	http.Redirect(res, req, target.String(), http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
