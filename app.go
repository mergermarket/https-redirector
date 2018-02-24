package main

import "net/http"
import "encoding/json"
import "time"

type HealthCheckStatus struct {
	Status string
	Date time.Time
}

func handler(res http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	http.Redirect(res, req, target, http.StatusMovedPermanently)
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
	http.ListenAndServe(":80", nil)
}
