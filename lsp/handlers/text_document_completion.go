package handlers

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"path-intellisense-lsp/glsp"
	protocol "path-intellisense-lsp/protocol_3_16"
)

func TextDocumentCompletion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	slog.Debug(fmt.Sprintf("TextDocumentCompletion: %s", params.TextDocument.URI))

	var completionItems []protocol.CompletionItem

	text := currentFiles[params.TextDocument.URI].Text
	line, err := extractFileLine(text, params.Position.Line)
	if err != nil {
		return completionItems, err
	}

	// Proceed if input is relative file path
	rePathSyntax := regexp.MustCompile(`[.]+(\/[^\\/:*?"<>|\r\n]+)*\/`)
	matches := rePathSyntax.FindAllString(line[:params.Position.Character], -1)
	if len(matches) == 0 {
		return completionItems, nil
	}

	path := matches[len(matches)-1]
	absoluteDir, _ := filepath.Split(params.TextDocument.URI[7:])
	absolutePath := filepath.Join(absoluteDir, path, "*")

	suggestedAbsolutePaths, err := filepath.Glob(absolutePath)
	if err != nil {
		return completionItems, err
	}

	for _, suggestedAbsolutePath := range suggestedAbsolutePaths {
		_, suggestion := filepath.Split(suggestedAbsolutePath)

		detail := ""
		label := ""

		fileInfo, err := os.Stat(suggestedAbsolutePath)
		if err == nil && fileInfo.IsDir() {
			detail = "📂 Folder"
			label = "📂 " + suggestion
		} else {
			detail = "📄 File"
			label = "📄 " + suggestion
		}
		detail += "\n" + path + suggestion

		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      label,
			Detail:     &detail,
			InsertText: &suggestion,
		})
	}
	return completionItems, nil
}

func extractFileLine(text string, linePosition uint32) (string, error) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanLine := uint32(0)
	for scanner.Scan() {
		line := scanner.Text()
		if scanLine == linePosition {
			return line, nil
		}
		scanLine += 1
	}

	slog.Error(fmt.Sprintf("Cannot find line %d in file", linePosition))
	return "", errors.New("cannot find line in file")
}
