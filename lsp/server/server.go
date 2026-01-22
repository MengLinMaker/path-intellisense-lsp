package server

import (
	"time"

	"path-intellisense-lsp/glsp"
)

var DefaultTimeout = time.Minute

type Server struct {
	Handler       glsp.Handler
	Timeout       time.Duration
	StreamTimeout time.Duration
}

func NewServer(handler glsp.Handler) *Server {
	return &Server{
		Handler:       handler,
		Timeout:       DefaultTimeout,
		StreamTimeout: DefaultTimeout,
	}
}
