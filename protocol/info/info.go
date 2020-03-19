package info

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/242617/synapse-core/api"
	"github.com/242617/synapse-core/version"
)

type Info struct{}

func (i *Info) Info(context.Context, *empty.Empty) (*api.InfoResponse, error) {

	// struct {
	// 	Application string `json:"application"`
	// 	Environment string `json:"environment"`
	// 	Version     string `json:"version"`
	// }{version.Application, version.Environment, version.Version}

	return &api.InfoResponse{
		Version: version.Version,
	}, nil
}
