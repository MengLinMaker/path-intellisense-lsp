package handlers

import (
	"encoding/json"
	"log/slog"
	"path-intellisense-lsp/src/glsp"
	protocol "path-intellisense-lsp/src/protocol_3_16"
)

func Initialized(ctx *glsp.Context, params *protocol.InitializedParams) error {
	slog.Debug("Initialized server")
	return nil
}

func SetTrace(ctx *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}

func LogTrace(ctx *glsp.Context, params *protocol.LogTraceParams) error {
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

func Shutdown(ctx *glsp.Context) error {
	slog.Warn("Shutdown server")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func Exit(ctx *glsp.Context) error {
	slog.Warn("Exit server")
	return nil
}
