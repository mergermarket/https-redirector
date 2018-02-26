package main

import (
	"log"
  "encoding/json"
  "time"
	"net/http"
)

type HealthCheckStatus struct {
	Status string
	Date time.Time
}

func handler(res http.ResponseWriter, req *http.Request) {
	target := *req.URL
	target.Scheme = "https"
	target.Host = req.Host
	http.Redirect(res, req, target.String(), http.StatusMovedPermanently)
}

func healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	healthCheckStatus := HealthCheckStatus{}
	healthCheckStatus.Status = "OK"
	healthCheckStatus.Date = time.Now().Local()

	healthCheckStatusJson, err := json.Marshal(healthCheckStatus)
	if err != nil{
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(healthCheckStatusJson)
}

func main() {
	http.HandleFunc("/health-check", healthCheckHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
