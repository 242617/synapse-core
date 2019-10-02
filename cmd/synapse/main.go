package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"

	"github.com/242617/synapse/config"
	"github.com/242617/synapse/log"
	"github.com/242617/synapse/server"
)

var (
	configFile = flag.String("config", "config.yaml", "Application config path")
)

func main() {
	flag.Parse()

	err := config.Init(*configFile)
	if err != nil {
		fmt.Println(errors.Wrap(err, "cannot init config"))
		os.Exit(1)
	}

	err = sentry.Init(sentry.ClientOptions{Dsn: config.Cfg.Services.Sentry.DSN})
	if err != nil {
		fmt.Println(errors.Wrap(err, "cannot init sentry"))
		os.Exit(1)
	}

	base, err := log.Create()
	if err != nil {
		sentry.CaptureException(err)
		sentry.Flush(5 * time.Second)
		fmt.Println(errors.Wrap(err, "cannot create logger"))
		os.Exit(1)
	}

	err = server.Init(base.With().Str("unit", "server").Logger())
	if err != nil {
		sentry.CaptureException(err)
		sentry.Flush(5 * time.Second)
		fmt.Println(errors.Wrap(err, "cannot init server"))
		os.Exit(1)
	}
}
