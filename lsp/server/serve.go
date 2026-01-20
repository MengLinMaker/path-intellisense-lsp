package server

import (
	"io"
	"log/slog"
)

// See: https://github.com/sourcegraph/go-langserver/blob/master/main.go#L179

func (self *Server) ServeStream(stream io.ReadWriteCloser) {
	slog.Info("new stream connection")
	<-self.newStreamConnection(stream).DisconnectNotify()
	slog.Info("stream connection closed")
}
