package main

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	openvmv1Connect "github.com/openvm-http/openvm-api/gen/openvm/v1/v1connect"
	"github.com/openvm-http/openvm-api/internal/interceptor"
	openvmServer "github.com/openvm-http/openvm-api/internal/service/openvm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var gitTag string
var dateTime string

func main() {
	log.Printf("OpenVM-API %s %s", gitTag, dateTime)
	if token := os.Getenv("ACCESS_TOKEN"); token != "" {
		interceptor.AccessToken = token
	} else {
		log.Printf("Security Warning: ACCESS_TOKEN has not been set!\n")
	}
	addr := "localhost:8080"
	if addrEnv := os.Getenv("ADDR"); addrEnv != "" {
		addr = addrEnv
	}

	api := http.NewServeMux()
	api.Handle(openvmv1Connect.NewApiServiceHandler(
		&openvmServer.ApiServer{},
		connect.WithInterceptors(interceptor.NewAuthInterceptor()),
	))
	apiSrv := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(api, &http2.Server{}),
	}
	log.Printf("apiSrv server listening on %s\n", addr)
	go func() {
		if err := apiSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("apiSrv listen and serve: %v\n", err)
		}
	}()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := apiSrv.Shutdown(ctx); err != nil {
		log.Fatalf("apiSrv shutdown: %v\n", err)
	}
}
