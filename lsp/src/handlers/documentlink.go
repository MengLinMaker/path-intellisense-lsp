package handlers

import (
	"fmt"
	"log/slog"
	"path-intellisense-lsp/src/glsp"
	protocol "path-intellisense-lsp/src/protocol_3_16"
	"path/filepath"
	"strings"
)

func TextDocumentDocumentLink(ctx *glsp.Context, params *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) {
	slog.Debug(fmt.Sprintf("TextDocumentDocumentLink for file: %s", params.TextDocument.URI))
	re := mustCompileLazyRegex(triggerCharacter + optionalPathPrefix + fmt.Sprintf("(/[^%s]+)+", illegalCharacters))

	documentLinks := []protocol.DocumentLink{}
	for i, line := range textLines(currentFiles[params.TextDocument.URI].Text) {
		matches := re.FindAllStringIndex("\n"+line, -1)

		for _, loc := range matches {
			// By spec len(loc) == 2
			start := loc[0]
			end := loc[1] - 1
			if start < 0 || end <= start || end > len(line) {
				slog.Error(fmt.Sprintf("Failed to extract path from line:\n%s\nstart(%d), end(%d), len(%d)", line, start, end, len(line)))
				continue
			}

			searchPath := line[start:end]
			for _, absolutePath := range matchPath(searchPath, params.TextDocument.URI, "") {
				target := "file://" + absolutePath
				absoluteDir, _ := filepath.Split(params.TextDocument.URI[7:])
				tooltip := strings.Replace(absolutePath, absoluteDir, "./", 1)
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
					Target:  &target,
					Tooltip: &tooltip,
				})
			}
		}
	}

	return documentLinks, nil
}
