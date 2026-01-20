package server

import (
	contextpkg "context"
	"io"

	"github.com/sourcegraph/jsonrpc2"
)

func (self *Server) newStreamConnection(stream io.ReadWriteCloser) *jsonrpc2.Conn {
	handler := self.newHandler()

	context, cancel := contextpkg.WithTimeout(contextpkg.Background(), self.StreamTimeout)
	defer cancel()

	return jsonrpc2.NewConn(context, jsonrpc2.NewBufferedStream(stream, jsonrpc2.VSCodeObjectCodec{}), handler, nil)
}
