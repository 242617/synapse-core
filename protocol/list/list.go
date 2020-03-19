package list

import (
	"context"
	"fmt"

	"github.com/242617/synapse-core/api"
)

type List struct{}

func (l *List) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	fmt.Printf("Received: %v", in.Name)
	return &api.HelloReply{Message: "Hello " + in.Name}, nil
}
