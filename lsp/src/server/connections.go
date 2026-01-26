package server

import (
	"context"
	"io"

	"github.com/sourcegraph/jsonrpc2"
)

func (s *Server) newStreamConnection(stream io.ReadWriteCloser) *jsonrpc2.Conn {
	handler := s.newHandler()

	ctx, cancel := context.WithTimeout(context.Background(), s.StreamTimeout)
	defer cancel()

	return jsonrpc2.NewConn(ctx, jsonrpc2.NewBufferedStream(stream, jsonrpc2.VSCodeObjectCodec{}), handler, nil)
}
