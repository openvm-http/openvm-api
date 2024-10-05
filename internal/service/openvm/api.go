package openvm

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	openvmv1 "github.com/openvm-http/openvm-api/gen/openvm/v1"
	"log"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[openvmv1.GreetRequest],
) (*connect.Response[openvmv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&openvmv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
