package httpserver

import (
	"context"
	"net/http"
)

// A Server defines the wrapper for the http server.
type Server struct {
	server *http.Server
	notify chan error
}

// New returns the wrapper of the http server.
func New(s *http.Server) Server {
	return Server{server: s, notify: make(chan error)}
}

// Start starts listening on the TCP network port.
func (s Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify returns channel with server runtime error.
// Will block until a non-nil error is written to the pipe.
func (s Server) Notify() <-chan error {
	return s.notify
}

// Shutdown gracefully shuts down the server without interrupting
// active connections.
func (s Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
