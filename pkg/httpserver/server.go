package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":8080"
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	server          *http.Server
	shutdownTimeout time.Duration
	notify          chan error
}

func New(handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}
	server := &Server{
		server:          httpServer,
		shutdownTimeout: _defaultShutdownTimeout,
		notify:          make(chan error),
	}

	// todo: apply options

	return server
}

func (s *Server) Start() <-chan error {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
	return s.notify
}

func (s *Server) ShutDown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
