package helper

import (
	"connectrpc.com/connect"
	"google.golang.org/protobuf/proto"
	"log"
)

func WrapConnectErrorDetail(connectErr *connect.Error, msg proto.Message) *connect.Error {
	errDetail, err := connect.NewErrorDetail(msg)
	if err != nil {
		log.Fatalf("connect.NewErrorDetail: %v\n", err)
	}
	connectErr.AddDetail(errDetail)
	return connectErr
}
