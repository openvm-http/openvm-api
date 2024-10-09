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
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
		log.Printf("Security Warning: Non-local network!\n")
	}

	interceptors := connect.WithInterceptors(interceptor.NewAuthInterceptor())
	api := http.NewServeMux()
	api.Handle(openvmv1Connect.NewApiServiceHandler(
		&openvmServer.ApiServer{},
		interceptors,
	))
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", api))

	srv := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
	log.Printf("HTTP server listening on %s\n", addr)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP listen and serve: %v\n", err)
		}
	}()

	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown: %v\n", err)
	}
}
