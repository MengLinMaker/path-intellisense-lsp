package handlers

import (
	"fmt"
	"log/slog"
	"path-intellisense-lsp/src/glsp"
	protocol "path-intellisense-lsp/src/protocol_3_16"
	"regexp"
)

type CurrentFile struct {
	Text       string
	Version    int32
	LanguageID string
	Path       string
}

// Print CurrentFile for debugging
func (s CurrentFile) Println() {
	slog.Info("\n" + s.Path + "\n-----\n" + s.Text + "\n-----\n")
}

var currentFiles = map[string]*CurrentFile{}

func TextDocumentDidOpen(ctx *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	slog.Debug(fmt.Sprintf("Caching openned file: %s", params.TextDocument.URI))
	currentFiles[params.TextDocument.URI] = &CurrentFile{
		Text:       params.TextDocument.Text,
		Version:    params.TextDocument.Version,
		LanguageID: params.TextDocument.LanguageID,
		Path:       params.TextDocument.URI,
	}
	textDocumentPublishDiagnostics(ctx, &textDocumentPublishDiagnosticsParams{
		URI:     params.TextDocument.URI,
		Version: params.TextDocument.Version,
		Text:    params.TextDocument.Text,
	})
	return nil
}

func TextDocumentDidSave(ctx *glsp.Context, params *protocol.DidSaveTextDocumentParams) error {
	slog.Debug(fmt.Sprintf("Caching saved file: %s", params.TextDocument.URI))
	if params.Text == nil {
		return nil
	}
	currentFiles[params.TextDocument.URI].Text = *params.Text
	textDocumentPublishDiagnostics(ctx, &textDocumentPublishDiagnosticsParams{
		URI:     params.TextDocument.URI,
		Version: 0,
		Text:    *params.Text,
	})
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
	fileLines := regexp.MustCompile("\r?\n").Split(currentFile.Text, -1)

	deleteLineIds := map[uint32]bool{}
	concatLineIds := map[uint32]bool{}

	for _, contentChange := range params.ContentChanges {
		switch v := contentChange.(type) {

		case protocol.TextDocumentContentChangeEventWhole:
			currentFile.Text = v.Text

		case protocol.TextDocumentContentChangeEvent:
			// Adding to text
			if v.Range.Start == v.Range.End {
				tmp := fileLines[v.Range.End.Line]
				fileLines[v.Range.End.Line] = tmp[:v.Range.End.Character] + v.Text + tmp[v.Range.End.Character:]

				// Removing from same line
			} else if v.Range.Start.Line == v.Range.End.Line {
				tmp := fileLines[v.Range.End.Line]
				fileLines[v.Range.End.Line] = tmp[:v.Range.Start.Character] + tmp[v.Range.End.Character:]

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
	for i, line := range fileLines {
		if !deleteLineIds[uint32(i)] {
			// Explicitly concat or is last line
			if concatLineIds[uint32(i)] || i == len(fileLines)-1 {
				text += line
			} else {
				text += line + "\n"
			}
		}
	}

	currentFiles[params.TextDocument.URI].Text = text
	currentFiles[params.TextDocument.URI].Version = params.TextDocument.Version
	textDocumentPublishDiagnostics(ctx, &textDocumentPublishDiagnosticsParams{
		URI:     params.TextDocument.URI,
		Version: params.TextDocument.Version,
		Text:    text,
	})
	return nil
}
