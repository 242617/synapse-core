package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"

	"github.com/242617/synapse-core/config"
	"github.com/242617/synapse-core/graphql/generated"
	"github.com/242617/synapse-core/log"
	"github.com/242617/synapse-core/types"
)

const (
	ReadTimeout    = 20 * time.Second
	WriteTimeout   = 20 * time.Second
	MaxHeaderBytes = 1 << 20
)

type server struct {
	cfg config.ServerConfig
	log log.Logger
	srv *http.Server
	gql *handler.Server
}

func NewServer(config config.ServerConfig, logger log.Logger, resolver generated.ResolverRoot) (types.Lifecycle, error) {
	es := generated.NewExecutableSchema(generated.Config{Resolvers: resolver})
	gql := handler.New(es)
	gql.AddTransport(transport.Options{})
	gql.AddTransport(transport.GET{})
	gql.AddTransport(transport.POST{})
	gql.Use(extension.Introspection{})

	srv := &http.Server{
		Handler:        gql,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: MaxHeaderBytes,
	}

	return &server{
		cfg: config,
		log: logger,
		srv: srv,
		gql: gql,
	}, nil
}

func (s *server) Start() error {
	address, err := net.ResolveTCPAddr("tcp4", s.cfg.Address)
	if err != nil {
		return fmt.Errorf("cannot resolve address %s", s.cfg.Address)
	}

	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		return fmt.Errorf("cannot listen %s", address)
	}

	go func() {
		err := s.srv.Serve(listener)
		if err != nil {
			s.log.Error().Err(err).Msg("cannot start serving")
		}
	}()
	s.log.Info().Msg("starting...")

	return nil
}

func (s *server) Stop() error {
	s.log.Info().Msg("stopping...")
	err := s.srv.Shutdown(context.TODO())
	return err
}
