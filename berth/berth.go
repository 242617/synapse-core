package berth

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/242617/synapse-core/api"
	"github.com/242617/synapse-core/config"
	"github.com/242617/synapse-core/log"
	"github.com/242617/synapse-core/protocol/system"
	"github.com/242617/synapse-core/protocol/tasks"
	"github.com/242617/synapse-core/types"
)

type server struct {
	cfg config.BerthConfig
	log log.Logger
	srv *grpc.Server
}

func NewServer(config config.BerthConfig, logger log.Logger) (types.Lifecycle, error) {
	s := grpc.NewServer()
	reflection.Register(s)
	api.RegisterSystemServer(s, &system.System{})
	api.RegisterTasksServer(s, &tasks.Tasks{})

	return &server{
		log: logger,
		srv: s,
	}, nil
}

func (s *server) Start() error {
	listener, err := net.Listen("tcp", s.cfg.Address)
	if err != nil {
		s.log.Error().Err(err).Msg("cannot listen")
		return err
	}

	s.log.Info().Msg("starting server...")
	if err := s.srv.Serve(listener); err != nil {
		s.log.Error().Err(err).Msg("cannot start serving")
		return err
	}

	return nil
}

func (s *server) Stop() error {
	s.log.Info().Msg("stopping server...")
	s.srv.GracefulStop()
	return nil
}
