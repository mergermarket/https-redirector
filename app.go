package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.URL.Path
    if len(r.URL.RawQuery) > 0 {
        target += "?" + r.URL.RawQuery
    }
	http.Redirect(w, r, target, http.StatusFound)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}

