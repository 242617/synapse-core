package server

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/242617/synapse-core/api"
	"github.com/242617/synapse-core/config"
	"github.com/242617/synapse-core/log"
	"github.com/242617/synapse-core/protocol/system"
	"github.com/242617/synapse-core/protocol/tasks"
)

var logger log.Logger

func Init(base log.Logger) error {
	logger = base

	listener, err := net.Listen("tcp", config.Cfg.Server.Address)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("cannot listen")
		return err
	}

	server := grpc.NewServer()
	reflection.Register(server)
	api.RegisterSystemServer(server, &system.System{})
	api.RegisterTasksServer(server, &tasks.Tasks{})
	err = server.Serve(listener)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("cannot serve")
		return err
	}

	return nil
}
