package server

import (
	"fmt"
	"net/http"

	"github.com/242617/synapse-core/config"
	"github.com/242617/synapse-core/log"
)

var logger log.Logger

func Init(base log.Logger) error {
	logger = base

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Debug().
			Str("uri", r.RequestURI).
			Msg("request")
		fmt.Fprintf(w, "ok: %s", r.RequestURI)
	})

	err := http.ListenAndServe(config.Cfg.Server.Address, nil)
	return err
}
