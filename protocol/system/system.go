package system

import (
	"context"
	"time"

	"github.com/242617/synapse-core/api"
	"github.com/242617/synapse-core/version"
)

var start = time.Now()

type System struct{}

func (*System) Info(context.Context, *api.Void) (*api.InfoResponse, error) {
	return &api.InfoResponse{
		Application: version.Application,
		Environment: version.Environment,
		Version:     version.Version,
	}, nil
}

func (*System) Uptime(context.Context, *api.Void) (*api.UptimeResponse, error) {
	return &api.UptimeResponse{
		Duration: time.Since(start).String(),
	}, nil
}
