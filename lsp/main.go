package main

import (
	"encoding/json"
	"log/slog"
	"os"
	"path-intellisense-lsp/glsp"
	"path-intellisense-lsp/handlers"
	"path-intellisense-lsp/server"
	"strings"

	protocol "path-intellisense-lsp/protocol_3_16"
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
		Initialized: initialized,
		SetTrace:    setTrace,
		LogTrace:    logTrace,
		Shutdown:    shutdown,
		Exit:        exit,
		// Handlers
		TextDocumentDidOpen:    handlers.TextDocumentDidOpen,
		TextDocumentDidClose:   handlers.TextDocumentDidClose,
		TextDocumentDidChange:  handlers.TextDocumentDidChange,
		TextDocumentCompletion: handlers.TextDocumentCompletion,
	}

	server := server.NewServer(&handler)

	err := server.RunStdio()
	if err == nil {
		slog.Error("Couldn't run server")
	}
}

func initialize(ctx *glsp.Context, params *protocol.InitializeParams) (any, error) {
	slog.Debug("Initializing server...")

	capabilities := handler.CreateServerCapabilities()

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lspName,
			Version: &version,
		},
	}, nil
}

func initialized(ctx *glsp.Context, params *protocol.InitializedParams) error {
	slog.Debug("Initialized server")
	return nil
}

func setTrace(ctx *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}

func logTrace(ctx *glsp.Context, params *protocol.LogTraceParams) error {
	traceValue := protocol.GetTraceValue()

	switch traceValue {
	case protocol.TraceValueMessage:
		slog.Info(params.Message)

	case protocol.TraceValueVerbose:
		jsonData, err := json.MarshalIndent(params, "", "  ")
		if err != nil {
			return err
		}
		slog.Debug(string(jsonData))
	}

	return nil
}

func shutdown(ctx *glsp.Context) error {
	slog.Warn("Shutdown server")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func exit(ctx *glsp.Context) error {
	slog.Warn("Exit server")
	return nil
}
