package server

import (
	"io"
	"log/slog"
)

// See: https://github.com/sourcegraph/go-langserver/blob/master/main.go#L179

func (s *Server) ServeStream(stream io.ReadWriteCloser) {
	slog.Debug("new stream connection")
	<-s.newStreamConnection(stream).DisconnectNotify()
	slog.Debug("stream connection closed")
}
