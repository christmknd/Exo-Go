package main

import (
	"fmt"
	"net/http"
	"time"
)

func ApiHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t := time.Now()
		hour := t.Hour()
		minute := t.Minute()
		fmt.Fprintf(w, "%v h %v", hour, minute)
	}
}

func main() {
	http.HandleFunc("/", ApiHandler)
	http.ListenAndServe(":4567", nil)
}
