package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

type responseJson struct {
	Hello string `json:"hello"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		body, _ := json.Marshal(responseJson{Hello: fmt.Sprintf("Hello,%q", html.EscapeString(r.URL.Path))})
		w.Write(body)
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
