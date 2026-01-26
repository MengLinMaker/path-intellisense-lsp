package handlers

import (
	"fmt"
	"log/slog"
	"os"
	"path-intellisense-lsp/src/glsp"
	protocol "path-intellisense-lsp/src/protocol_3_16"
	"path/filepath"
	"strings"
)

func TextDocumentDocumentLink(ctx *glsp.Context, params *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) {
	slog.Debug(fmt.Sprintf("TextDocumentDocumentLink for file: %s", params.TextDocument.URI))

	documentLinks := []protocol.DocumentLink{}
	for i, line := range textLines(currentFiles[params.TextDocument.URI].Text) {
		for _, match := range findPathMatches(line) {
			for _, absolutePath := range matchPath(match.Text, params.TextDocument.URI, "") {
				target := "file://" + absolutePath
				absoluteDir, _ := filepath.Split(params.TextDocument.URI[7:])

				tooltip := "📄 File: "
				fileInfo, err := os.Stat(absolutePath)
				if err == nil && fileInfo.IsDir() {
					tooltip = "📂 Folder: "
				}
				tooltip += strings.Replace(absolutePath, absoluteDir, "./", 1)

				documentLinks = append(documentLinks, protocol.DocumentLink{
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
					Target:  &target,
					Tooltip: &tooltip,
				})
			}
		}
	}
	return documentLinks, nil
}
