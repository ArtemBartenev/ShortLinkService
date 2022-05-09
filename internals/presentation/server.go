package presentation

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) LaunchServer(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
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
