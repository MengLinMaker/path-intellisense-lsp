package protocol

import (
	"encoding/json"
	"errors"
	"sync"

	"path-intellisense-lsp/glsp"

	protocol316 "path-intellisense-lsp/protocol_3_16"
)

type Handler struct {
	protocol316.Handler

	Initialize             InitializeFunc
	TextDocumentDiagnostic TextDocumentDiagnosticFunc

	initialized bool
	lock        sync.Mutex
}

func (s *Handler) Handle(context *glsp.Context) (r any, validMethod bool, validParams bool, err error) {
	if !s.IsInitialized() && (context.Method != protocol316.MethodInitialize) {
		return nil, true, true, errors.New("server not initialized")
	}

	switch context.Method {
	case protocol316.MethodCancelRequest:
		if s.CancelRequest != nil {
			validMethod = true
			var params protocol316.CancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.CancelRequest(context, &params)
			}
		}

	case protocol316.MethodProgress:
		if s.Progress != nil {
			validMethod = true
			var params protocol316.ProgressParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.Progress(context, &params)
			}
		}

	// General Messages

	case MethodInitialize:
		if s.Initialize != nil {
			validMethod = true
			var params InitializeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				if r, err = s.Initialize(context, &params); err == nil {
					s.SetInitialized(true)
				}
			}
		}

	case protocol316.MethodInitialized:
		if s.Initialized != nil {
			validMethod = true
			var params protocol316.InitializedParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.Initialized(context, &params)
			}
		}

	case protocol316.MethodShutdown:
		s.SetInitialized(false)
		if s.Shutdown != nil {
			validMethod = true
			validParams = true
			err = s.Shutdown(context)
		}

	case protocol316.MethodExit:
		// Note that the server will close the connection after we handle it here
		if s.Exit != nil {
			validMethod = true
			validParams = true
			err = s.Exit(context)
		}

	case protocol316.MethodLogTrace:
		if s.LogTrace != nil {
			validMethod = true
			var params protocol316.LogTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.LogTrace(context, &params)
			}
		}

	case protocol316.MethodSetTrace:
		if s.SetTrace != nil {
			validMethod = true
			var params protocol316.SetTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.SetTrace(context, &params)
			}
		}

	// Window

	case protocol316.MethodWindowWorkDoneProgressCancel:
		if s.WindowWorkDoneProgressCancel != nil {
			validMethod = true
			var params protocol316.WorkDoneProgressCancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WindowWorkDoneProgressCancel(context, &params)
			}
		}

	// Workspace

	case protocol316.MethodWorkspaceDidChangeWorkspaceFolders:
		if s.WorkspaceDidChangeWorkspaceFolders != nil {
			validMethod = true
			var params protocol316.DidChangeWorkspaceFoldersParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidChangeWorkspaceFolders(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidChangeConfiguration:
		if s.WorkspaceDidChangeConfiguration != nil {
			validMethod = true
			var params protocol316.DidChangeConfigurationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidChangeConfiguration(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidChangeWatchedFiles:
		if s.WorkspaceDidChangeWatchedFiles != nil {
			validMethod = true
			var params protocol316.DidChangeWatchedFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidChangeWatchedFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceSymbol:
		if s.WorkspaceSymbol != nil {
			validMethod = true
			var params protocol316.WorkspaceSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.WorkspaceSymbol(context, &params)
			}
		}

	case protocol316.MethodWorkspaceExecuteCommand:
		if s.WorkspaceExecuteCommand != nil {
			validMethod = true
			var params protocol316.ExecuteCommandParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.WorkspaceExecuteCommand(context, &params)
			}
		}

	case protocol316.MethodWorkspaceWillCreateFiles:
		if s.WorkspaceWillCreateFiles != nil {
			validMethod = true
			var params protocol316.CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.WorkspaceWillCreateFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidCreateFiles:
		if s.WorkspaceDidCreateFiles != nil {
			validMethod = true
			var params protocol316.CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidCreateFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceWillRenameFiles:
		if s.WorkspaceWillRenameFiles != nil {
			validMethod = true
			var params protocol316.RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.WorkspaceWillRenameFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidRenameFiles:
		if s.WorkspaceDidRenameFiles != nil {
			validMethod = true
			var params protocol316.RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidRenameFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceWillDeleteFiles:
		if s.WorkspaceWillDeleteFiles != nil {
			validMethod = true
			var params protocol316.DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.WorkspaceWillDeleteFiles(context, &params)
			}
		}

	case protocol316.MethodWorkspaceDidDeleteFiles:
		if s.WorkspaceDidDeleteFiles != nil {
			validMethod = true
			var params protocol316.DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidDeleteFiles(context, &params)
			}
		}

	// Text Document Synchronization

	case protocol316.MethodTextDocumentDidOpen:
		if s.TextDocumentDidOpen != nil {
			validMethod = true
			var params protocol316.DidOpenTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.TextDocumentDidOpen(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDidChange:
		if s.TextDocumentDidChange != nil {
			validMethod = true
			var params protocol316.DidChangeTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.TextDocumentDidChange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentWillSave:
		if s.TextDocumentWillSave != nil {
			validMethod = true
			var params protocol316.WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.TextDocumentWillSave(context, &params)
			}
		}

	case protocol316.MethodTextDocumentWillSaveWaitUntil:
		if s.TextDocumentWillSaveWaitUntil != nil {
			validMethod = true
			var params protocol316.WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentWillSaveWaitUntil(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDidSave:
		if s.TextDocumentDidSave != nil {
			validMethod = true
			var params protocol316.DidSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.TextDocumentDidSave(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDidClose:
		if s.TextDocumentDidClose != nil {
			validMethod = true
			var params protocol316.DidCloseTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.TextDocumentDidClose(context, &params)
			}
		}

	// Language Features

	case protocol316.MethodTextDocumentCompletion:
		if s.TextDocumentCompletion != nil {
			validMethod = true
			var params protocol316.CompletionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentCompletion(context, &params)
			}
		}

	case protocol316.MethodCompletionItemResolve:
		if s.CompletionItemResolve != nil {
			validMethod = true
			var params protocol316.CompletionItem
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.CompletionItemResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentHover:
		if s.TextDocumentHover != nil {
			validMethod = true
			var params protocol316.HoverParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentHover(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSignatureHelp:
		if s.TextDocumentSignatureHelp != nil {
			validMethod = true
			var params protocol316.SignatureHelpParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentSignatureHelp(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDeclaration:
		if s.TextDocumentDeclaration != nil {
			validMethod = true
			var params protocol316.DeclarationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentDeclaration(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDefinition:
		if s.TextDocumentDefinition != nil {
			validMethod = true
			var params protocol316.DefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentDefinition(context, &params)
			}
		}

	case protocol316.MethodTextDocumentTypeDefinition:
		if s.TextDocumentTypeDefinition != nil {
			validMethod = true
			var params protocol316.TypeDefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentTypeDefinition(context, &params)
			}
		}

	case protocol316.MethodTextDocumentImplementation:
		if s.TextDocumentImplementation != nil {
			validMethod = true
			var params protocol316.ImplementationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentImplementation(context, &params)
			}
		}

	case protocol316.MethodTextDocumentReferences:
		if s.TextDocumentReferences != nil {
			validMethod = true
			var params protocol316.ReferenceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentReferences(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDocumentHighlight:
		if s.TextDocumentDocumentHighlight != nil {
			validMethod = true
			var params protocol316.DocumentHighlightParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentDocumentHighlight(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDocumentSymbol:
		if s.TextDocumentDocumentSymbol != nil {
			validMethod = true
			var params protocol316.DocumentSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentDocumentSymbol(context, &params)
			}
		}

	case protocol316.MethodTextDocumentCodeAction:
		if s.TextDocumentCodeAction != nil {
			validMethod = true
			var params protocol316.CodeActionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentCodeAction(context, &params)
			}
		}

	case protocol316.MethodCodeActionResolve:
		if s.CodeActionResolve != nil {
			validMethod = true
			var params protocol316.CodeAction
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.CodeActionResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentCodeLens:
		if s.TextDocumentCodeLens != nil {
			validMethod = true
			var params protocol316.CodeLensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentCodeLens(context, &params)
			}
		}

	case protocol316.MethodCodeLensResolve:
		if s.TextDocumentDidClose != nil {
			validMethod = true
			var params protocol316.CodeLens
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.CodeLensResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentDocumentLink:
		if s.TextDocumentDocumentLink != nil {
			validMethod = true
			var params protocol316.DocumentLinkParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentDocumentLink(context, &params)
			}
		}

	case protocol316.MethodDocumentLinkResolve:
		if s.DocumentLinkResolve != nil {
			validMethod = true
			var params protocol316.DocumentLink
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.DocumentLinkResolve(context, &params)
			}
		}

	case protocol316.MethodTextDocumentColor:
		if s.TextDocumentColor != nil {
			validMethod = true
			var params protocol316.DocumentColorParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentColor(context, &params)
			}
		}

	case protocol316.MethodTextDocumentColorPresentation:
		if s.TextDocumentColorPresentation != nil {
			validMethod = true
			var params protocol316.ColorPresentationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentColorPresentation(context, &params)
			}
		}

	case protocol316.MethodTextDocumentFormatting:
		if s.TextDocumentFormatting != nil {
			validMethod = true
			var params protocol316.DocumentFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentFormatting(context, &params)
			}
		}

	case protocol316.MethodTextDocumentRangeFormatting:
		if s.TextDocumentRangeFormatting != nil {
			validMethod = true
			var params protocol316.DocumentRangeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentRangeFormatting(context, &params)
			}
		}

	case protocol316.MethodTextDocumentOnTypeFormatting:
		if s.TextDocumentOnTypeFormatting != nil {
			validMethod = true
			var params protocol316.DocumentOnTypeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentOnTypeFormatting(context, &params)
			}
		}

	case protocol316.MethodTextDocumentRename:
		if s.TextDocumentRename != nil {
			validMethod = true
			var params protocol316.RenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentRename(context, &params)
			}
		}

	case protocol316.MethodTextDocumentPrepareRename:
		if s.TextDocumentPrepareRename != nil {
			validMethod = true
			var params protocol316.PrepareRenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentPrepareRename(context, &params)
			}
		}

	case protocol316.MethodTextDocumentFoldingRange:
		if s.TextDocumentFoldingRange != nil {
			validMethod = true
			var params protocol316.FoldingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentFoldingRange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSelectionRange:
		if s.TextDocumentSelectionRange != nil {
			validMethod = true
			var params protocol316.SelectionRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentSelectionRange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentPrepareCallHierarchy:
		if s.TextDocumentPrepareCallHierarchy != nil {
			validMethod = true
			var params protocol316.CallHierarchyPrepareParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentPrepareCallHierarchy(context, &params)
			}
		}

	case protocol316.MethodCallHierarchyIncomingCalls:
		if s.CallHierarchyIncomingCalls != nil {
			validMethod = true
			var params protocol316.CallHierarchyIncomingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.CallHierarchyIncomingCalls(context, &params)
			}
		}

	case protocol316.MethodCallHierarchyOutgoingCalls:
		if s.CallHierarchyOutgoingCalls != nil {
			validMethod = true
			var params protocol316.CallHierarchyOutgoingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.CallHierarchyOutgoingCalls(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSemanticTokensFull:
		if s.TextDocumentSemanticTokensFull != nil {
			validMethod = true
			var params protocol316.SemanticTokensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentSemanticTokensFull(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSemanticTokensFullDelta:
		if s.TextDocumentSemanticTokensFullDelta != nil {
			validMethod = true
			var params protocol316.SemanticTokensDeltaParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentSemanticTokensFullDelta(context, &params)
			}
		}

	case protocol316.MethodTextDocumentSemanticTokensRange:
		if s.TextDocumentSemanticTokensRange != nil {
			validMethod = true
			var params protocol316.SemanticTokensRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentSemanticTokensRange(context, &params)
			}
		}

	case protocol316.MethodWorkspaceSemanticTokensRefresh:
		if s.WorkspaceSemanticTokensRefresh != nil {
			validMethod = true
			validParams = true
			err = s.WorkspaceSemanticTokensRefresh(context)
		}

	case protocol316.MethodTextDocumentLinkedEditingRange:
		if s.TextDocumentLinkedEditingRange != nil {
			validMethod = true
			var params protocol316.LinkedEditingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentLinkedEditingRange(context, &params)
			}
		}

	case protocol316.MethodTextDocumentMoniker:
		if s.TextDocumentMoniker != nil {
			validMethod = true
			var params protocol316.MonikerParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentMoniker(context, &params)
			}
		}
	case MethodTextDocumentDiagnostic:
		if s.TextDocumentDiagnostic != nil {
			validMethod = true
			var params DocumentDiagnosticParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentDiagnostic(context, &params)
			}
		}

	default:
		if s.CustomRequest != nil {
			if handler, ok := s.CustomRequest[context.Method]; ok && (handler.Func != nil) {
				validMethod = true
				if err = json.Unmarshal(context.Params, &handler.Params); err == nil {
					validParams = true
					r, err = handler.Func(context, handler.Params)
				}
			}
		}
	}

	return

}

func (s *Handler) IsInitialized() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.initialized
}

func (s *Handler) SetInitialized(initialized bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.initialized = initialized
}

func (s *Handler) CreateServerCapabilities() ServerCapabilities {
	var capabilities ServerCapabilities

	if (s.TextDocumentDidOpen != nil) || (s.TextDocumentDidClose != nil) {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).OpenClose = &protocol316.True
	}

	if s.TextDocumentDidChange != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		// This can be overriden to TextDocumentSyncKindFull
		value := protocol316.TextDocumentSyncKindIncremental
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).Change = &value
	}

	if s.TextDocumentWillSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).WillSave = &protocol316.True
	}

	if s.TextDocumentWillSaveWaitUntil != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).WillSaveWaitUntil = &protocol316.True
	}

	if s.TextDocumentDidSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &protocol316.TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*protocol316.TextDocumentSyncOptions).Save = &protocol316.True
	}

	if s.TextDocumentCompletion != nil {
		capabilities.CompletionProvider = &protocol316.CompletionOptions{}
	}

	if s.TextDocumentHover != nil {
		capabilities.HoverProvider = true
	}

	if s.TextDocumentSignatureHelp != nil {
		capabilities.SignatureHelpProvider = &protocol316.SignatureHelpOptions{}
	}

	if s.TextDocumentDeclaration != nil {
		capabilities.DeclarationProvider = true
	}

	if s.TextDocumentDefinition != nil {
		capabilities.DefinitionProvider = true
	}

	if s.TextDocumentTypeDefinition != nil {
		capabilities.TypeDefinitionProvider = true
	}

	if s.TextDocumentImplementation != nil {
		capabilities.ImplementationProvider = true
	}

	if s.TextDocumentReferences != nil {
		capabilities.ReferencesProvider = true
	}

	if s.TextDocumentDocumentHighlight != nil {
		capabilities.DocumentHighlightProvider = true
	}

	if s.TextDocumentDocumentSymbol != nil {
		capabilities.DocumentSymbolProvider = true
	}

	if s.TextDocumentCodeAction != nil {
		capabilities.CodeActionProvider = true
	}

	if s.TextDocumentCodeLens != nil {
		capabilities.CodeLensProvider = &protocol316.CodeLensOptions{}
	}

	if s.TextDocumentDocumentLink != nil {
		capabilities.DocumentLinkProvider = &protocol316.DocumentLinkOptions{}
	}

	if s.TextDocumentColor != nil {
		capabilities.ColorProvider = true
	}

	if s.TextDocumentFormatting != nil {
		capabilities.DocumentFormattingProvider = true
	}

	if s.TextDocumentRangeFormatting != nil {
		capabilities.DocumentRangeFormattingProvider = true
	}

	if s.TextDocumentOnTypeFormatting != nil {
		capabilities.DocumentOnTypeFormattingProvider = &protocol316.DocumentOnTypeFormattingOptions{}
	}

	if s.TextDocumentRename != nil {
		capabilities.RenameProvider = true
	}

	if s.TextDocumentFoldingRange != nil {
		capabilities.FoldingRangeProvider = true
	}

	if s.WorkspaceExecuteCommand != nil {
		capabilities.ExecuteCommandProvider = &protocol316.ExecuteCommandOptions{}
	}

	if s.TextDocumentSelectionRange != nil {
		capabilities.SelectionRangeProvider = true
	}

	if s.TextDocumentLinkedEditingRange != nil {
		capabilities.LinkedEditingRangeProvider = true
	}

	if s.TextDocumentPrepareCallHierarchy != nil {
		capabilities.CallHierarchyProvider = true
	}

	if s.TextDocumentSemanticTokensFull != nil {
		if _, ok := capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &protocol316.SemanticTokensOptions{}
		}
		if s.TextDocumentSemanticTokensFullDelta != nil {
			capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Full = &protocol316.SemanticDelta{}
			capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Full.(*protocol316.SemanticDelta).Delta = &protocol316.True
		} else {
			capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Full = true
		}
	}

	if s.TextDocumentSemanticTokensRange != nil {
		if _, ok := capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &protocol316.SemanticTokensOptions{}
		}
		capabilities.SemanticTokensProvider.(*protocol316.SemanticTokensOptions).Range = true
	}

	// TODO: s.TextDocumentSemanticTokensRefresh?

	if s.TextDocumentMoniker != nil {
		capabilities.MonikerProvider = true
	}

	if s.WorkspaceSymbol != nil {
		capabilities.WorkspaceSymbolProvider = true
	}

	if s.WorkspaceDidCreateFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidCreate = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if s.WorkspaceWillCreateFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillCreate = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if s.WorkspaceDidRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidRename = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if s.WorkspaceWillRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillRename = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if s.WorkspaceDidDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidDelete = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if s.WorkspaceWillDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &protocol316.ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &protocol316.ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillDelete = &protocol316.FileOperationRegistrationOptions{
			Filters: []protocol316.FileOperationFilter{},
		}
	}

	if s.TextDocumentDiagnostic != nil {
		capabilities.DiagnosticProvider = DiagnosticOptions{
			InterFileDependencies: true,
			WorkspaceDiagnostics:  false,
		}
	}

	return capabilities
}
