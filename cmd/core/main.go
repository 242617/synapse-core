package main

import (
	"flag"
	l "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"

	"github.com/242617/synapse-core/config"
	"github.com/242617/synapse-core/graphql"
	"github.com/242617/synapse-core/log"
	"github.com/242617/synapse-core/secret"
	"github.com/242617/synapse-core/server"
	"github.com/242617/synapse-core/types"
	"github.com/242617/synapse-core/version"
)

var (
	configFile = flag.String("config", "config.yaml", "Application config path")
)

func main() {
	flag.Parse()

	cfg, err := config.Init(*configFile)
	if err != nil {
		l.Println(errors.Wrap(err, "cannot init config"))
		os.Exit(1)
	}

	l.Println("cfg")
	l.Println(cfg)

	scrt, err := secret.Init(cfg.Services.Vault)
	if err != nil {
		l.Println(errors.Wrap(err, "cannot init secret"))
		os.Exit(1)
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:         scrt.SentryDSN,
		Environment: version.Environment,
	})
	if err != nil {
		l.Println(errors.Wrap(err, "cannot init sentry"))
		os.Exit(1)
	}

	base, err := log.Create(cfg.Logger)
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
	logger := base

	exitCode := 1
	defer func() { os.Exit(exitCode) }()

	srv, err := server.NewServer(cfg.Server, logger.With().Str("unit", "server").Logger(), &graphql.Resolver{})
	if err != nil {
		logger.Error().Err(err).Msg("cannot create server")
		return
	}

	err = run(srv)
	if err != nil {
		logger.Error().Err(err).Msg("cannot run server")
		return
	}

	exitCode = 0

	// err = berth.Init(base.With().Str("unit", "server").Logger())
	// if err != nil {
	// 	sentry.CaptureException(err)
	// 	defer sentry.Flush(5 * time.Second)
	// 	base.Error().
	// 		Err(err).
	// 		Msg("cannot init server")
	// 	os.Exit(1)
	// }
}

func run(srv types.Lifecycle) error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	if err := srv.Start(); err != nil {
		return errors.Wrap(err, "connot start server")
	}

	<-signals
	close(signals)

	if err := srv.Stop(); err != nil {
		return errors.Wrap(err, "cannot stop server")
	}

	return nil
}
