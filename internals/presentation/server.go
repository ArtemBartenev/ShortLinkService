package presentation

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) LaunchServer(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":8000",
		Handler:      handler,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) GracefulShutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func AddRoutes(handler http.Handler) {
	http.Handle("/v1/short/url", handler)
}
