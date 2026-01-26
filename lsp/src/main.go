package main

import (
	"log/slog"
	"os"
	"path-intellisense-lsp/src/glsp"
	"path-intellisense-lsp/src/handlers"
	"path-intellisense-lsp/src/server"
	"strings"

	protocol "path-intellisense-lsp/src/protocol_3_16"
)

var (
	lspName        = "Path intellisense lsp"
	version string = "0.0.1"
	handler protocol.Handler
)

func main() {
	switch strings.ToUpper(os.Getenv("LOG_LEVEL")) {
	case "DEBUG":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "WARN":
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case "ERROR":
		slog.SetLogLoggerLevel(slog.LevelError)
	}

	handler = protocol.Handler{
		// Lifecycle
		Initialize:  initialize,
		Initialized: handlers.Initialized,
		SetTrace:    handlers.SetTrace,
		LogTrace:    handlers.LogTrace,
		Shutdown:    handlers.Shutdown,
		Exit:        handlers.Exit,
		// Handlers for basic
		CancelRequest: handlers.CancelRequest,
		// Handlers for file syncing
		TextDocumentDidOpen:   handlers.TextDocumentDidOpen,
		TextDocumentDidSave:   handlers.TextDocumentDidSave,
		TextDocumentDidClose:  handlers.TextDocumentDidClose,
		TextDocumentDidChange: handlers.TextDocumentDidChange,
		// Handlers for code completion
		TextDocumentCompletion: handlers.TextDocumentCompletion,
		// Handlers for navigation
		TextDocumentDocumentLink: handlers.TextDocumentDocumentLink,
	}

	server := server.NewServer(&handler)

	err := server.RunStdio()
	if err == nil {
		slog.Error("Couldn't run server")
	}
}

func initialize(ctx *glsp.Context, params *protocol.InitializeParams) (any, error) {
	slog.Debug("Initializing server...")

	options := protocol.ServerCapabilitiesOptions{
		CompletionOptions: &protocol.CompletionOptions{
			TriggerCharacters: []string{
				"/",
			},
		},
	}
	capabilities := handler.CreateServerCapabilities(&options)
	initializeResult := protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lspName,
			Version: &version,
		},
	}
	return initializeResult, nil
}
