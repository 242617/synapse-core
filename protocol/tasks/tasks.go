package tasks

import (
	"context"
	"fmt"

	"github.com/242617/synapse-core/api"
)

type Tasks struct{}

func (*Tasks) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	fmt.Printf("Received: %v", in.Name)
	return &api.HelloReply{Message: "Hello " + in.Name}, nil
}
