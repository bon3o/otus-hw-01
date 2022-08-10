package internalhttp

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/bon3o/otus-hw-01/hw12_13_14_15_calendar/internal/app"
)

type Server struct {
	host   string
	port   string
	logger app.Logger
	server *http.Server
}

type Application interface { // TODO
}

func NewServer(logger app.Logger, app Application, host, port string) *Server {
	server := &Server{
		host:   host,
		port:   port,
		logger: logger,
		server: nil,
	}

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: loggingMiddleware(http.HandlerFunc(server.handleHTTP), logger),
	}

	server.server = httpServer

	return server
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info(fmt.Sprintf("HTTP server listen and serve %s:%s", s.host, s.port))
	if err := s.server.ListenAndServe(); err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) handleHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
