package handlers

import (
	"fmt"
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

type pathMatch struct {
	Text  string
	Start int
	End   int
}

func findPathMatches(line string) []pathMatch {
	search := "\n" + line
	re := mustCompileLazyRegex(triggerCharacter + optionalPathPrefix + fmt.Sprintf("(/[^%s]+)*", illegalCharacters))

	matches := re.FindAllStringIndex(search, -1)
	if len(matches) == 0 {
		return nil
	}

	results := make([]pathMatch, 0, len(matches))
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		start := match[0]
		end := match[1] - 1
		if start < 0 || end <= start || end > len(line) {
			continue
		}
		pathText := search[start:match[1]][1:]
		if pathText == "" {
			continue
		}
		results = append(results, pathMatch{
			Start: start,
			End:   end,
			Text:  pathText,
		})
	}
	return results
}
