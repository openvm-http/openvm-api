package main

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	openvmv1Connect "github.com/openvm-http/openvm-api/gen/openvm/v1/v1connect"
	"github.com/openvm-http/openvm-api/internal/interceptor"
	openvmServer "github.com/openvm-http/openvm-api/internal/service/openvm"
	"github.com/rs/cors"
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

func disableCORS() *cors.Cors {
	// To let web developers play with the demo service from browsers, we need a
	// very permissive CORS setup.
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(_ /* origin */ string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
}

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

	interceptors := connect.WithInterceptors(interceptor.NewAuthInterceptor())
	api := http.NewServeMux()
	api.Handle(openvmv1Connect.NewApiServiceHandler(
		&openvmServer.ApiServer{},
		interceptors,
	))
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", api))
	var httpServerMux http.Handler
	if disableCors := os.Getenv("DISABLE_CORS"); disableCors == "YES_I_KNOWN_NOT_SAFE" {
		log.Printf("Security Warning: DISABLE_CORS set!\n")
		httpServerMux = disableCORS().Handler(mux)
	} else {
		httpServerMux = mux
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(httpServerMux, &http2.Server{}),
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
