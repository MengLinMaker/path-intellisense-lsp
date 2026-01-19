package main

import (
	"log/slog"
	"os"
	"path-intellisense-lsp/handlers"
	"strings"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

var (
	lspName        = "Path intellisense lsp"
	version string = "0.0.1"
	handler protocol.Handler
)

func main() {
	setEnvLogLevel()
	slog.Debug("Booting lsp")

	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentCompletion: handlers.TextDocumentCompletion,
	}

	server := server.NewServer(&handler, lspName, true)
	slog.Debug("Created lsp instance")

	err := server.RunStdio()
	if err == nil {
		panic("Couldn't run server")
	}
}

func setEnvLogLevel() {
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
	slog.Info("Initializing server...")

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
	slog.Info("Initialized server")
	return nil
}

func shutdown(context *glsp.Context) error {
	slog.Info("Shutdown server")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}
