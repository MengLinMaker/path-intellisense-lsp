package handlers

import (
	"fmt"
	"log/slog"
	"path-intellisense-lsp/src/glsp"
	protocol "path-intellisense-lsp/src/protocol_3_16"
)

type textDocumentPublishDiagnosticsParams struct {
	URI     string
	Version int32
	Text    string
}

func textDocumentPublishDiagnostics(ctx *glsp.Context, params *textDocumentPublishDiagnosticsParams) {
	slog.Debug(fmt.Sprintf("TextDocumentPublishDiagnostics for file: %s", params.URI))

	diagnostics := []protocol.Diagnostic{}
	for i, line := range textLines(params.Text) {
		for _, match := range findPathMatches(line) {
			if len(matchPath(match.Text, params.URI, "")) > 0 {
				continue
			}
			severity := protocol.DiagnosticSeverityError
			source := "path-intellisense-lsp"
			diagnostics = append(diagnostics, protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      uint32(i),
						Character: uint32(match.Start),
					},
					End: protocol.Position{
						Line:      uint32(i),
						Character: uint32(match.End),
					},
				},
				Severity: &severity,
				Source:   &source,
				Message:  fmt.Sprintf("Path not found: %s", match.Text),
			})
		}
	}

	version := uint32(params.Version)
	ctx.Notify(protocol.ServerTextDocumentPublishDiagnostics, &protocol.PublishDiagnosticsParams{
		URI:         params.URI,
		Version:     &version,
		Diagnostics: diagnostics,
	})
}
