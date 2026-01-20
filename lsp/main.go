package main

import (
	"log"
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
	// setEnvLogLevel()
	// slog.Info("Booting lsp")

	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentCompletion: handlers.TextDocumentCompletion,
	}

	server := server.NewServer(&handler)
	// slog.Info("Created lsp instance")

	err := server.RunStdio()
	if err == nil {
		panic("Couldn't run server")
	}
}

func setEnvLogLevel() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	defaultLevel := slog.LevelInfo
	envLevel := os.Getenv("LOG_LEVEL")
	levelMap := map[string]slog.Level{
		"DEBUG": slog.LevelDebug,
		"INFO":  slog.LevelInfo,
		"WARN":  slog.LevelWarn,
		"ERROR": slog.LevelError,
	}
	level, ok := levelMap[strings.ToUpper(envLevel)]
	if !ok {
		level = defaultLevel
	}
	opts := slog.HandlerOptions{
		Level: level,
	}
	handler := slog.NewTextHandler(os.Stdout, &opts)
	slog.SetDefault(slog.New(handler))
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	// slog.Info("Initializing server...")

	capabilities := handler.CreateServerCapabilities()
	// capabilities.CompletionProvider = &protocol.CompletionOptions{}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lspName,
			Version: &version,
		},
	}, nil
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	// slog.Info("Initialized server")
	return nil
}

func shutdown(context *glsp.Context) error {
	// slog.Info("Shutdown server")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}
