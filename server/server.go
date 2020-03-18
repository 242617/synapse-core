package server

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/242617/synapse-crawler/protocol"

	"github.com/242617/synapse-core/log"
	"github.com/242617/synapse-core/version"
)

var logger log.Logger

func Init(base log.Logger) error {
	logger = base

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()

		var request protocol.Request
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			logger.Error().
				Err(err).
				Msg("cannot decode request")
			return
		}

		switch request.Name {

		case "ping":

			response := protocol.Response{Name: "ping"}

			data := struct {
				Application string `json:"application"`
				Environment string `json:"environment"`
				Version     string `json:"version"`
			}{version.Application, version.Environment, version.Version}

			barr, err := json.Marshal(data)
			if err != nil {
				logger.Error().
					Err(err).
					Msg("cannot marshal ping data")
				return
			}

			response.Data = json.RawMessage(barr)

			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				logger.Error().
					Err(err).
					Msg("cannot encode response")
				return
			}

		case "tasks":
			fmt.Fprintln(w, `{"name":"tasks","data":[{"type":"pull","resource":"__debug__"}]}`)

		case "__debug__":
			rand.Seed(time.Now().UnixNano())
			fmt.Fprintf(w, `{"name":"__debug__","data":%d}`, rand.Intn(1000))

		default:
			http.Error(w, "not implemented", http.StatusNotImplemented)

		}

	})

	err := http.ListenAndServeTLS(":443", "synapse.core.crt", "synapse.core.key", nil)
	// err := http.ListenAndServe(config.Cfg.Server.Address, nil)
	return err
}
