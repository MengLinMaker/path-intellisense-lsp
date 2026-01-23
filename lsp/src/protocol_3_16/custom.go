package protocol

import (
	"encoding/json"

	"path-intellisense-lsp/src/glsp"
)

type CustomRequestHandler struct {
	Func CustomRequestFunc
	// This field should be private however it is used in both versions of the protocol
	Params json.RawMessage
}

type CustomRequestHandlers map[string]CustomRequestHandler

type CustomRequestFunc func(ctx *glsp.Context, params json.RawMessage) (any, error)
