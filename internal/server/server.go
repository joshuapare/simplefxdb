package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/fx"
)

// NewHTTPServer builds an HTTP server that will begin serving requests
// when the Fx application starts.
func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: ":4422", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			fmt.Println("Starting HTTP Server")
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Shutting down HTTP Server")
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

// NewServeMux builds a ServeMux that will route requests
// to the given EchoHandler.
func NewServeMux(query *QueryHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/query", query)
	return mux
}
