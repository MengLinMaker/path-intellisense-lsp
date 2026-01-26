package handlers

import (
	"fmt"
	"log/slog"
	"path-intellisense-lsp/src/glsp"
	protocol "path-intellisense-lsp/src/protocol_3_16"
)

func TextDocumentDocumentLink(ctx *glsp.Context, params *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) {
	slog.Debug(fmt.Sprintf("TextDocumentDocumentLink for file: %s", params.TextDocument.URI))

	documentLinks := []protocol.DocumentLink{}
	lines := textLines(currentFiles[params.TextDocument.URI].Text)
	re := mustCompileLazyRegex(triggerCharacter + optionalPathPrefix + fmt.Sprintf("(/[^%s]+)*", illegalCharacters))

	for i, line := range lines {
		search := "\n" + line
		matches := re.FindAllStringIndex(search, -1)
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

			matchedPaths := matchPath(pathText, string(params.TextDocument.URI), "")
			if len(matchedPaths) == 0 {
				continue
			}

			target := protocol.DocumentUri("file://" + matchedPaths[0])
			documentLinks = append(documentLinks, protocol.DocumentLink{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      uint32(i),
						Character: uint32(start),
					},
					End: protocol.Position{
						Line:      uint32(i),
						Character: uint32(end),
					},
				},
				Target: &target,
			})
		}
	}

	return documentLinks, nil
}
