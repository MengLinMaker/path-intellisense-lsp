package main

import (
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
	setEnvLogLevel()

	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentCompletion: handlers.TextDocumentCompletion,
	}

	server := server.NewServer(&handler)

	err := server.RunStdio()
	if err == nil {
		slog.Error("Couldn't run server")
	}
}

func setEnvLogLevel() {
	envLevel := os.Getenv("LOG_LEVEL")

	switch strings.ToUpper(envLevel) {
	case "DEBUG":
		slog.SetLogLoggerLevel(slog.LevelDebug)
	case "WARN":
		slog.SetLogLoggerLevel(slog.LevelWarn)
	case "ERROR":
		slog.SetLogLoggerLevel(slog.LevelError)
	}
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
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

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	slog.Debug("Initialized server")
	return nil
}

func shutdown(context *glsp.Context) error {
	slog.Warn("Shutdown server")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}
