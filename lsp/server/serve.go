package server

import (
	"io"
	"log/slog"
)

// See: https://github.com/sourcegraph/go-langserver/blob/master/main.go#L179

func (self *Server) ServeStream(stream io.ReadWriteCloser) {
	slog.Debug("new stream connection")
	<-self.newStreamConnection(stream).DisconnectNotify()
	slog.Debug("stream connection closed")
}
