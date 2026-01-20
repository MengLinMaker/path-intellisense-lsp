package server

import (
	contextpkg "context"
	"io"

	"github.com/sourcegraph/jsonrpc2"
)

func (s *Server) newStreamConnection(stream io.ReadWriteCloser) *jsonrpc2.Conn {
	handler := s.newHandler()

	context, cancel := contextpkg.WithTimeout(contextpkg.Background(), s.StreamTimeout)
	defer cancel()

	return jsonrpc2.NewConn(context, jsonrpc2.NewBufferedStream(stream, jsonrpc2.VSCodeObjectCodec{}), handler, nil)
}
