package main

import (
	"flag"
	l "log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"

	"github.com/242617/synapse-core/config"
	"github.com/242617/synapse-core/log"
	"github.com/242617/synapse-core/secret"
	"github.com/242617/synapse-core/server"
	"github.com/242617/synapse-core/version"
)

var (
	configFile = flag.String("config", "config.yaml", "Application config path")
)

func main() {
	flag.Parse()

	err := config.Init(*configFile)
	if err != nil {
		l.Println(errors.Wrap(err, "cannot init config"))
		os.Exit(1)
	}

	err = secret.Init()
	if err != nil {
		l.Println(errors.Wrap(err, "cannot init secret"))
		os.Exit(1)
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:         secret.SentryDSN,
		Environment: version.Environment,
	})
	if err != nil {
		l.Println(errors.Wrap(err, "cannot init sentry"))
		os.Exit(1)
	}

	base, err := log.Create()
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(5 * time.Second)
		l.Println(errors.Wrap(err, "cannot create logger"))
		os.Exit(1)
	}
	base = base.With().Str("application", version.Application).Logger()

	base.Info().
		Str("environment", version.Environment).
		Str("version", version.Version).
		Str("build", version.Build).
		Msg("start")

	err = server.Init(base.With().Str("unit", "server").Logger())
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(5 * time.Second)
		base.Error().
			Err(err).
			Msg("cannot init server")
		os.Exit(1)
	}
}
