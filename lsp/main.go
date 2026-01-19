package main

import (
	"log/slog"
	"path-intellisense-lsp/handlers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

const lspName = "Path intellisense language server"

var (
	version string = "0.0.1"
	handler protocol.Handler
)

func main() {
	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentCompletion: handlers.TextDocumentCompletion,
	}

	server := server.NewServer(&handler, lspName, true)

	err := server.RunStdio()
	if err == nil {
		panic("Couldn't run server")
	}
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	slog.Info("Initializing server...")

	capabilities := handler.CreateServerCapabilities()

	capabilities.CompletionProvider = &protocol.CompletionOptions{}

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
