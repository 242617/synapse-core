package server

import (
	"net"

	"google.golang.org/grpc"

	"github.com/242617/synapse-core/api"
	"github.com/242617/synapse-core/config"
	"github.com/242617/synapse-core/log"
	"github.com/242617/synapse-core/protocol/info"
	"github.com/242617/synapse-core/protocol/list"
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
	api.RegisterSystemServer(server, &info.Info{})
	api.RegisterListServer(server, &list.List{})
	err = server.Serve(listener)
	if err != nil {
		logger.Error().
			Err(err).
			Msg("cannot serve")
		return err
	}

	return nil
}
