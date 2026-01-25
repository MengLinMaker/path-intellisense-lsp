package handlers

import (
	"encoding/json"
	"errors"
	"path-intellisense-lsp/src/glsp"
	protocol "path-intellisense-lsp/src/protocol_3_16"
)

func CancelRequest(ctx *glsp.Context, params *protocol.CancelParams) error {
	message, _ := json.Marshal(params)
	return errors.New(string(message))
}
