package handlers

import (
	"bufio"
	"fmt"
	"log/slog"
	"path-intellisense-lsp/glsp"
	protocol "path-intellisense-lsp/protocol_3_16"
	"strings"
)

type CurrentFile struct {
	Text       string
	Version    int32
	LanguageID string
}

var currentFiles = map[string]CurrentFile{}

func TextDocumentDidOpen(ctx *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	slog.Debug(fmt.Sprintf("Caching file: %s", params.TextDocument.URI))
	currentFiles[params.TextDocument.URI] = CurrentFile{
		Text:       params.TextDocument.Text,
		Version:    params.TextDocument.Version,
		LanguageID: params.TextDocument.LanguageID,
	}
	return nil
}

func TextDocumentDidClose(ctx *glsp.Context, params *protocol.DidCloseTextDocumentParams) error {
	slog.Debug(fmt.Sprintf("Deleting file cache: %s", params.TextDocument.URI))
	delete(currentFiles, params.TextDocument.URI)
	return nil
}

func TextDocumentDidChange(ctx *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	slog.Debug(fmt.Sprintf("Updating file cache: %s", params.TextDocument.URI))
	currentFile := currentFiles[params.TextDocument.URI]

	fileLines := []string{}
	{
		scanner := bufio.NewScanner(strings.NewReader(currentFile.Text))
		for scanner.Scan() {
			fileLines = append(fileLines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			slog.Error(err.Error())
		}
	}

	deleteLineIds := map[uint32]bool{}
	concatLineIds := map[uint32]bool{}

	for _, contentChange := range params.ContentChanges {
		switch v := contentChange.(type) {

		case protocol.TextDocumentContentChangeEventWhole:
			currentFile.Text = v.Text

		case protocol.TextDocumentContentChangeEvent:
			// Adding to text
			if v.Range.Start == v.Range.End {
				tmpLine := fileLines[v.Range.End.Line]
				fileLines[v.Range.End.Line] = tmpLine[:v.Range.End.Character] + v.Text + tmpLine[v.Range.End.Character:]

				// Removing from same line
			} else if v.Range.Start.Line == v.Range.End.Line {
				tmpLine := fileLines[v.Range.End.Line]
				fileLines[v.Range.End.Line] = tmpLine[:v.Range.Start.Character] + tmpLine[v.Range.End.Character:]

				// Removing from multiple lines
			} else {
				fileLines[v.Range.Start.Line] = fileLines[v.Range.Start.Line][:v.Range.Start.Character]
				fileLines[v.Range.End.Line] = fileLines[v.Range.End.Line][v.Range.End.Character:]
				concatLineIds[v.Range.Start.Line] = true
				for i := v.Range.Start.Line + 1; i < v.Range.End.Line; i++ {
					deleteLineIds[i] = true
				}
			}

		default:
			slog.Error("Unkown contentChange type")
		}
	}

	text := ""
	for i := 0; i < len(fileLines); i++ {
		if !deleteLineIds[uint32(i)] {
			if concatLineIds[uint32(i)] {
				text += fileLines[i]
			} else {
				text += fileLines[i] + "\n"
			}
		}
	}

	currentFiles[params.TextDocument.URI] = CurrentFile{
		Text:       text,
		Version:    params.TextDocument.Version,
		LanguageID: currentFiles[params.TextDocument.URI].LanguageID,
	}
	return nil
}
