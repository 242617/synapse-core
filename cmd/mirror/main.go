package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	ServerAddress = ":8080"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request %s\n", r.RequestURI)
		fmt.Fprintf(w, "ok: %s", r.RequestURI)
	})
	fmt.Printf("server start on %s\n", ServerAddress)
	log.Fatal(http.ListenAndServe(ServerAddress, nil))
}
