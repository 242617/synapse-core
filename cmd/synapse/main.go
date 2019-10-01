package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/242617/synapse/config"
)

var (
	configFile = flag.String("config", "config.yaml", "Application config path")
)

func init() { log.SetFlags(log.Lshortfile) }

func main() {
	flag.Parse()

	if err := config.Init(*configFile); err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %s\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request %s\n", r.RequestURI)
		fmt.Fprintf(w, "ok: %s", r.RequestURI)
	})
	fmt.Printf("server start on %s\n", config.Cfg.Server.Address)
	log.Fatal(http.ListenAndServe(config.Cfg.Server.Address, nil))
}
