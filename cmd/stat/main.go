package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	data := struct {
		Crawlers uint `json:"crawlers"`
	}{
		Crawlers: 10,
	}
	err := json.NewEncoder(os.Stdout).Encode(&data)
	if err != nil {
		log.Fatal(err)
	}
}
