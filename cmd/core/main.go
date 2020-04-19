package main

import (
	"flag"
	l "log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"

	"github.com/242617/synapse-core/berth"
	"github.com/242617/synapse-core/config"
	"github.com/242617/synapse-core/graphql"
	"github.com/242617/synapse-core/log"
	"github.com/242617/synapse-core/secret"
	"github.com/242617/synapse-core/server"
	"github.com/242617/synapse-core/types"
	"github.com/242617/synapse-core/version"
)

const (
	StopTimeout  = 10 * time.Second
	StartTimeout = 10 * time.Second
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

	resolver := graphql.NewResolver()
	srv, err := server.NewServer(cfg.Server, logger.With().Str("unit", "server").Logger(), resolver)
	if err != nil {
		logger.Error().Err(err).Msg("cannot create server")
		return
	}

	brth, err := berth.NewServer(cfg.Berth, logger.With().Str("unit", "berth").Logger())
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(5 * time.Second)
		base.Error().
			Err(err).
			Msg("cannot init server")
		os.Exit(1)
	}

	exitCode := 1
	defer func() { os.Exit(exitCode) }()

	errCh := start(srv, brth)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		l.Println("err", err)
		logger.Error().Err(err).Msg("cannot start servers")
		return
	case <-signals:
		close(signals)
	}

	errCh = stop(srv, brth)
	for err := range errCh {
		logger.Error().Err(err).Msg("cannot stop servers")
		return
	}

	exitCode = 0
}

func start(services ...types.Lifecycle) chan error {
	ch := make(chan error)
	var wg sync.WaitGroup
	wg.Add(len(services))
	go func() {
		wg.Wait()
		close(ch)
	}()
	for _, service := range services {
		service := service
		go func() {
			defer wg.Done()
			if err := service.Start(); err != nil {
				ch <- err
				return
			}
		}()
	}
	return ch
}

func stop(services ...types.Lifecycle) chan error {
	ch := make(chan error)
	var wg sync.WaitGroup
	wg.Add(len(services))
	go func() {
		wg.Wait()
		close(ch)
	}()
	for _, service := range services {
		service := service
		go func() {
			defer wg.Done()
			if err := service.Stop(); err != nil {
				ch <- err
				return
			}
		}()
	}
	return ch
}
