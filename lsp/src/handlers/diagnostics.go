package handlers

import (
	"path-intellisense-lsp/src/glsp"
	protocol "path-intellisense-lsp/src/protocol_3_16"
)

type textDocumentPublishDiagnosticsParams struct {
	URI     string
	Version int32
	Text    string
}

func textDocumentPublishDiagnostics(ctx *glsp.Context, params *textDocumentPublishDiagnosticsParams) {
	diagnostics := []protocol.Diagnostic{}

	version := uint32(params.Version)
	ctx.Notify(protocol.ServerTextDocumentPublishDiagnostics, &protocol.PublishDiagnosticsParams{
		URI:         params.URI,
		Version:     &version,
		Diagnostics: diagnostics,
	})
}
