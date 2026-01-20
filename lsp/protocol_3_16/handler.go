package protocol

import (
	"encoding/json"
	"errors"
	"sync"

	"path-intellisense-lsp/glsp"
)

type Handler struct {
	// Base Protocol
	CancelRequest CancelRequestFunc
	Progress      ProgressFunc

	// General Messages
	Initialize  InitializeFunc
	Initialized InitializedFunc
	Shutdown    ShutdownFunc
	Exit        ExitFunc
	LogTrace    LogTraceFunc
	SetTrace    SetTraceFunc

	// Window
	WindowWorkDoneProgressCancel WindowWorkDoneProgressCancelFunc

	// Workspace
	WorkspaceDidChangeWorkspaceFolders WorkspaceDidChangeWorkspaceFoldersFunc
	WorkspaceDidChangeConfiguration    WorkspaceDidChangeConfigurationFunc
	WorkspaceDidChangeWatchedFiles     WorkspaceDidChangeWatchedFilesFunc
	WorkspaceSymbol                    WorkspaceSymbolFunc
	WorkspaceExecuteCommand            WorkspaceExecuteCommandFunc
	WorkspaceWillCreateFiles           WorkspaceWillCreateFilesFunc
	WorkspaceDidCreateFiles            WorkspaceDidCreateFilesFunc
	WorkspaceWillRenameFiles           WorkspaceWillRenameFilesFunc
	WorkspaceDidRenameFiles            WorkspaceDidRenameFilesFunc
	WorkspaceWillDeleteFiles           WorkspaceWillDeleteFilesFunc
	WorkspaceDidDeleteFiles            WorkspaceDidDeleteFilesFunc
	WorkspaceSemanticTokensRefresh     WorkspaceSemanticTokensRefreshFunc

	// Text Document Synchronization
	TextDocumentDidOpen           TextDocumentDidOpenFunc
	TextDocumentDidChange         TextDocumentDidChangeFunc
	TextDocumentWillSave          TextDocumentWillSaveFunc
	TextDocumentWillSaveWaitUntil TextDocumentWillSaveWaitUntilFunc
	TextDocumentDidSave           TextDocumentDidSaveFunc
	TextDocumentDidClose          TextDocumentDidCloseFunc

	// Language Features
	TextDocumentCompletion              TextDocumentCompletionFunc
	CompletionItemResolve               CompletionItemResolveFunc
	TextDocumentHover                   TextDocumentHoverFunc
	TextDocumentSignatureHelp           TextDocumentSignatureHelpFunc
	TextDocumentDeclaration             TextDocumentDeclarationFunc
	TextDocumentDefinition              TextDocumentDefinitionFunc
	TextDocumentTypeDefinition          TextDocumentTypeDefinitionFunc
	TextDocumentImplementation          TextDocumentImplementationFunc
	TextDocumentReferences              TextDocumentReferencesFunc
	TextDocumentDocumentHighlight       TextDocumentDocumentHighlightFunc
	TextDocumentDocumentSymbol          TextDocumentDocumentSymbolFunc
	TextDocumentCodeAction              TextDocumentCodeActionFunc
	CodeActionResolve                   CodeActionResolveFunc
	TextDocumentCodeLens                TextDocumentCodeLensFunc
	CodeLensResolve                     CodeLensResolveFunc
	TextDocumentDocumentLink            TextDocumentDocumentLinkFunc
	DocumentLinkResolve                 DocumentLinkResolveFunc
	TextDocumentColor                   TextDocumentColorFunc
	TextDocumentColorPresentation       TextDocumentColorPresentationFunc
	TextDocumentFormatting              TextDocumentFormattingFunc
	TextDocumentRangeFormatting         TextDocumentRangeFormattingFunc
	TextDocumentOnTypeFormatting        TextDocumentOnTypeFormattingFunc
	TextDocumentRename                  TextDocumentRenameFunc
	TextDocumentPrepareRename           TextDocumentPrepareRenameFunc
	TextDocumentFoldingRange            TextDocumentFoldingRangeFunc
	TextDocumentSelectionRange          TextDocumentSelectionRangeFunc
	TextDocumentPrepareCallHierarchy    TextDocumentPrepareCallHierarchyFunc
	CallHierarchyIncomingCalls          CallHierarchyIncomingCallsFunc
	CallHierarchyOutgoingCalls          CallHierarchyOutgoingCallsFunc
	TextDocumentSemanticTokensFull      TextDocumentSemanticTokensFullFunc
	TextDocumentSemanticTokensFullDelta TextDocumentSemanticTokensFullDeltaFunc
	TextDocumentSemanticTokensRange     TextDocumentSemanticTokensRangeFunc
	TextDocumentLinkedEditingRange      TextDocumentLinkedEditingRangeFunc
	TextDocumentMoniker                 TextDocumentMonikerFunc

	// Custom Request/Notification
	CustomRequest map[string]CustomRequestHandler

	initialized bool
	lock        sync.Mutex
}

// ([glsp.Handler] interface)
func (s *Handler) Handle(context *glsp.Context) (r any, validMethod bool, validParams bool, err error) {
	if !s.IsInitialized() && (context.Method != MethodInitialize) {
		return nil, true, true, errors.New("server not initialized")
	}

	switch context.Method {
	// Base Protocol

	case MethodCancelRequest:
		if s.CancelRequest != nil {
			validMethod = true
			var params CancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.CancelRequest(context, &params)
			}
		}

	case MethodProgress:
		if s.Progress != nil {
			validMethod = true
			var params ProgressParams
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

	case MethodInitialized:
		if s.Initialized != nil {
			validMethod = true
			var params InitializedParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.Initialized(context, &params)
			}
		}

	case MethodShutdown:
		s.SetInitialized(false)
		if s.Shutdown != nil {
			validMethod = true
			validParams = true
			err = s.Shutdown(context)
		}

	case MethodExit:
		// Note that the server will close the connection after we handle it here
		if s.Exit != nil {
			validMethod = true
			validParams = true
			err = s.Exit(context)
		}

	case MethodLogTrace:
		if s.LogTrace != nil {
			validMethod = true
			var params LogTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.LogTrace(context, &params)
			}
		}

	case MethodSetTrace:
		if s.SetTrace != nil {
			validMethod = true
			var params SetTraceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.SetTrace(context, &params)
			}
		}

	// Window

	case MethodWindowWorkDoneProgressCancel:
		if s.WindowWorkDoneProgressCancel != nil {
			validMethod = true
			var params WorkDoneProgressCancelParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WindowWorkDoneProgressCancel(context, &params)
			}
		}

	// Workspace

	case MethodWorkspaceDidChangeWorkspaceFolders:
		if s.WorkspaceDidChangeWorkspaceFolders != nil {
			validMethod = true
			var params DidChangeWorkspaceFoldersParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidChangeWorkspaceFolders(context, &params)
			}
		}

	case MethodWorkspaceDidChangeConfiguration:
		if s.WorkspaceDidChangeConfiguration != nil {
			validMethod = true
			var params DidChangeConfigurationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidChangeConfiguration(context, &params)
			}
		}

	case MethodWorkspaceDidChangeWatchedFiles:
		if s.WorkspaceDidChangeWatchedFiles != nil {
			validMethod = true
			var params DidChangeWatchedFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidChangeWatchedFiles(context, &params)
			}
		}

	case MethodWorkspaceSymbol:
		if s.WorkspaceSymbol != nil {
			validMethod = true
			var params WorkspaceSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.WorkspaceSymbol(context, &params)
			}
		}

	case MethodWorkspaceExecuteCommand:
		if s.WorkspaceExecuteCommand != nil {
			validMethod = true
			var params ExecuteCommandParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.WorkspaceExecuteCommand(context, &params)
			}
		}

	case MethodWorkspaceWillCreateFiles:
		if s.WorkspaceWillCreateFiles != nil {
			validMethod = true
			var params CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.WorkspaceWillCreateFiles(context, &params)
			}
		}

	case MethodWorkspaceDidCreateFiles:
		if s.WorkspaceDidCreateFiles != nil {
			validMethod = true
			var params CreateFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidCreateFiles(context, &params)
			}
		}

	case MethodWorkspaceWillRenameFiles:
		if s.WorkspaceWillRenameFiles != nil {
			validMethod = true
			var params RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.WorkspaceWillRenameFiles(context, &params)
			}
		}

	case MethodWorkspaceDidRenameFiles:
		if s.WorkspaceDidRenameFiles != nil {
			validMethod = true
			var params RenameFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidRenameFiles(context, &params)
			}
		}

	case MethodWorkspaceWillDeleteFiles:
		if s.WorkspaceWillDeleteFiles != nil {
			validMethod = true
			var params DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.WorkspaceWillDeleteFiles(context, &params)
			}
		}

	case MethodWorkspaceDidDeleteFiles:
		if s.WorkspaceDidDeleteFiles != nil {
			validMethod = true
			var params DeleteFilesParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.WorkspaceDidDeleteFiles(context, &params)
			}
		}

	// Text Document Synchronization

	case MethodTextDocumentDidOpen:
		if s.TextDocumentDidOpen != nil {
			validMethod = true
			var params DidOpenTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.TextDocumentDidOpen(context, &params)
			}
		}

	case MethodTextDocumentDidChange:
		if s.TextDocumentDidChange != nil {
			validMethod = true
			var params DidChangeTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.TextDocumentDidChange(context, &params)
			}
		}

	case MethodTextDocumentWillSave:
		if s.TextDocumentWillSave != nil {
			validMethod = true
			var params WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.TextDocumentWillSave(context, &params)
			}
		}

	case MethodTextDocumentWillSaveWaitUntil:
		if s.TextDocumentWillSaveWaitUntil != nil {
			validMethod = true
			var params WillSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentWillSaveWaitUntil(context, &params)
			}
		}

	case MethodTextDocumentDidSave:
		if s.TextDocumentDidSave != nil {
			validMethod = true
			var params DidSaveTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.TextDocumentDidSave(context, &params)
			}
		}

	case MethodTextDocumentDidClose:
		if s.TextDocumentDidClose != nil {
			validMethod = true
			var params DidCloseTextDocumentParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				err = s.TextDocumentDidClose(context, &params)
			}
		}

	// Language Features

	case MethodTextDocumentCompletion:
		if s.TextDocumentCompletion != nil {
			validMethod = true
			var params CompletionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentCompletion(context, &params)
			}
		}

	case MethodCompletionItemResolve:
		if s.CompletionItemResolve != nil {
			validMethod = true
			var params CompletionItem
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.CompletionItemResolve(context, &params)
			}
		}

	case MethodTextDocumentHover:
		if s.TextDocumentHover != nil {
			validMethod = true
			var params HoverParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentHover(context, &params)
			}
		}

	case MethodTextDocumentSignatureHelp:
		if s.TextDocumentSignatureHelp != nil {
			validMethod = true
			var params SignatureHelpParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentSignatureHelp(context, &params)
			}
		}

	case MethodTextDocumentDeclaration:
		if s.TextDocumentDeclaration != nil {
			validMethod = true
			var params DeclarationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentDeclaration(context, &params)
			}
		}

	case MethodTextDocumentDefinition:
		if s.TextDocumentDefinition != nil {
			validMethod = true
			var params DefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentDefinition(context, &params)
			}
		}

	case MethodTextDocumentTypeDefinition:
		if s.TextDocumentTypeDefinition != nil {
			validMethod = true
			var params TypeDefinitionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentTypeDefinition(context, &params)
			}
		}

	case MethodTextDocumentImplementation:
		if s.TextDocumentImplementation != nil {
			validMethod = true
			var params ImplementationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentImplementation(context, &params)
			}
		}

	case MethodTextDocumentReferences:
		if s.TextDocumentReferences != nil {
			validMethod = true
			var params ReferenceParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentReferences(context, &params)
			}
		}

	case MethodTextDocumentDocumentHighlight:
		if s.TextDocumentDocumentHighlight != nil {
			validMethod = true
			var params DocumentHighlightParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentDocumentHighlight(context, &params)
			}
		}

	case MethodTextDocumentDocumentSymbol:
		if s.TextDocumentDocumentSymbol != nil {
			validMethod = true
			var params DocumentSymbolParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentDocumentSymbol(context, &params)
			}
		}

	case MethodTextDocumentCodeAction:
		if s.TextDocumentCodeAction != nil {
			validMethod = true
			var params CodeActionParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentCodeAction(context, &params)
			}
		}

	case MethodCodeActionResolve:
		if s.CodeActionResolve != nil {
			validMethod = true
			var params CodeAction
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.CodeActionResolve(context, &params)
			}
		}

	case MethodTextDocumentCodeLens:
		if s.TextDocumentCodeLens != nil {
			validMethod = true
			var params CodeLensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentCodeLens(context, &params)
			}
		}

	case MethodCodeLensResolve:
		if s.TextDocumentDidClose != nil {
			validMethod = true
			var params CodeLens
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.CodeLensResolve(context, &params)
			}
		}

	case MethodTextDocumentDocumentLink:
		if s.TextDocumentDocumentLink != nil {
			validMethod = true
			var params DocumentLinkParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentDocumentLink(context, &params)
			}
		}

	case MethodDocumentLinkResolve:
		if s.DocumentLinkResolve != nil {
			validMethod = true
			var params DocumentLink
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.DocumentLinkResolve(context, &params)
			}
		}

	case MethodTextDocumentColor:
		if s.TextDocumentColor != nil {
			validMethod = true
			var params DocumentColorParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentColor(context, &params)
			}
		}

	case MethodTextDocumentColorPresentation:
		if s.TextDocumentColorPresentation != nil {
			validMethod = true
			var params ColorPresentationParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentColorPresentation(context, &params)
			}
		}

	case MethodTextDocumentFormatting:
		if s.TextDocumentFormatting != nil {
			validMethod = true
			var params DocumentFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentFormatting(context, &params)
			}
		}

	case MethodTextDocumentRangeFormatting:
		if s.TextDocumentRangeFormatting != nil {
			validMethod = true
			var params DocumentRangeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentRangeFormatting(context, &params)
			}
		}

	case MethodTextDocumentOnTypeFormatting:
		if s.TextDocumentOnTypeFormatting != nil {
			validMethod = true
			var params DocumentOnTypeFormattingParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentOnTypeFormatting(context, &params)
			}
		}

	case MethodTextDocumentRename:
		if s.TextDocumentRename != nil {
			validMethod = true
			var params RenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentRename(context, &params)
			}
		}

	case MethodTextDocumentPrepareRename:
		if s.TextDocumentPrepareRename != nil {
			validMethod = true
			var params PrepareRenameParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentPrepareRename(context, &params)
			}
		}

	case MethodTextDocumentFoldingRange:
		if s.TextDocumentFoldingRange != nil {
			validMethod = true
			var params FoldingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentFoldingRange(context, &params)
			}
		}

	case MethodTextDocumentSelectionRange:
		if s.TextDocumentSelectionRange != nil {
			validMethod = true
			var params SelectionRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentSelectionRange(context, &params)
			}
		}

	case MethodTextDocumentPrepareCallHierarchy:
		if s.TextDocumentPrepareCallHierarchy != nil {
			validMethod = true
			var params CallHierarchyPrepareParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentPrepareCallHierarchy(context, &params)
			}
		}

	case MethodCallHierarchyIncomingCalls:
		if s.CallHierarchyIncomingCalls != nil {
			validMethod = true
			var params CallHierarchyIncomingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.CallHierarchyIncomingCalls(context, &params)
			}
		}

	case MethodCallHierarchyOutgoingCalls:
		if s.CallHierarchyOutgoingCalls != nil {
			validMethod = true
			var params CallHierarchyOutgoingCallsParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.CallHierarchyOutgoingCalls(context, &params)
			}
		}

	case MethodTextDocumentSemanticTokensFull:
		if s.TextDocumentSemanticTokensFull != nil {
			validMethod = true
			var params SemanticTokensParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentSemanticTokensFull(context, &params)
			}
		}

	case MethodTextDocumentSemanticTokensFullDelta:
		if s.TextDocumentSemanticTokensFullDelta != nil {
			validMethod = true
			var params SemanticTokensDeltaParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentSemanticTokensFullDelta(context, &params)
			}
		}

	case MethodTextDocumentSemanticTokensRange:
		if s.TextDocumentSemanticTokensRange != nil {
			validMethod = true
			var params SemanticTokensRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentSemanticTokensRange(context, &params)
			}
		}

	case MethodWorkspaceSemanticTokensRefresh:
		if s.WorkspaceSemanticTokensRefresh != nil {
			validMethod = true
			validParams = true
			err = s.WorkspaceSemanticTokensRefresh(context)
		}

	case MethodTextDocumentLinkedEditingRange:
		if s.TextDocumentLinkedEditingRange != nil {
			validMethod = true
			var params LinkedEditingRangeParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentLinkedEditingRange(context, &params)
			}
		}

	case MethodTextDocumentMoniker:
		if s.TextDocumentMoniker != nil {
			validMethod = true
			var params MonikerParams
			if err = json.Unmarshal(context.Params, &params); err == nil {
				validParams = true
				r, err = s.TextDocumentMoniker(context, &params)
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
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).OpenClose = &True
	}

	if s.TextDocumentDidChange != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		// This can be overriden to TextDocumentSyncKindFull
		value := TextDocumentSyncKindIncremental
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).Change = &value
	}

	if s.TextDocumentWillSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).WillSave = &True
	}

	if s.TextDocumentWillSaveWaitUntil != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).WillSaveWaitUntil = &True
	}

	if s.TextDocumentDidSave != nil {
		if _, ok := capabilities.TextDocumentSync.(*TextDocumentSyncOptions); !ok {
			capabilities.TextDocumentSync = &TextDocumentSyncOptions{}
		}
		capabilities.TextDocumentSync.(*TextDocumentSyncOptions).Save = &True
	}

	if s.TextDocumentCompletion != nil {
		capabilities.CompletionProvider = &CompletionOptions{}
	}

	if s.TextDocumentHover != nil {
		capabilities.HoverProvider = true
	}

	if s.TextDocumentSignatureHelp != nil {
		capabilities.SignatureHelpProvider = &SignatureHelpOptions{}
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
		capabilities.CodeLensProvider = &CodeLensOptions{}
	}

	if s.TextDocumentDocumentLink != nil {
		capabilities.DocumentLinkProvider = &DocumentLinkOptions{}
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
		capabilities.DocumentOnTypeFormattingProvider = &DocumentOnTypeFormattingOptions{}
	}

	if s.TextDocumentRename != nil {
		capabilities.RenameProvider = true
	}

	if s.TextDocumentFoldingRange != nil {
		capabilities.FoldingRangeProvider = true
	}

	if s.WorkspaceExecuteCommand != nil {
		capabilities.ExecuteCommandProvider = &ExecuteCommandOptions{}
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
		if _, ok := capabilities.SemanticTokensProvider.(*SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &SemanticTokensOptions{}
		}
		if s.TextDocumentSemanticTokensFullDelta != nil {
			capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Full = &SemanticDelta{}
			capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Full.(*SemanticDelta).Delta = &True
		} else {
			capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Full = true
		}
	}

	if s.TextDocumentSemanticTokensRange != nil {
		if _, ok := capabilities.SemanticTokensProvider.(*SemanticTokensOptions); !ok {
			capabilities.SemanticTokensProvider = &SemanticTokensOptions{}
		}
		capabilities.SemanticTokensProvider.(*SemanticTokensOptions).Range = true
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
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidCreate = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if s.WorkspaceWillCreateFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillCreate = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if s.WorkspaceDidRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidRename = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if s.WorkspaceWillRenameFiles != nil {
		capabilities.RenameProvider = true
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillRename = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if s.WorkspaceDidDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.DidDelete = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	if s.WorkspaceWillDeleteFiles != nil {
		if capabilities.Workspace == nil {
			capabilities.Workspace = &ServerCapabilitiesWorkspace{}
		}
		if capabilities.Workspace.FileOperations == nil {
			capabilities.Workspace.FileOperations = &ServerCapabilitiesWorkspaceFileOperations{}
		}
		capabilities.Workspace.FileOperations.WillDelete = &FileOperationRegistrationOptions{
			Filters: []FileOperationFilter{},
		}
	}

	return capabilities
}
