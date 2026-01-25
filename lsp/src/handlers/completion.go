package handlers

import (
	"fmt"
	"log/slog"
	"os"
	"os/user"
	"path/filepath"

	"path-intellisense-lsp/src/glsp"
	protocol "path-intellisense-lsp/src/protocol_3_16"
)

func TextDocumentCompletion(ctx *glsp.Context, params *protocol.CompletionParams) (any, error) {
	slog.Debug(fmt.Sprintf("TextDocumentCompletion: %s", params.TextDocument.URI))
	var completionItems []protocol.CompletionItem

	// Validate file path syntax
	text := currentFiles[params.TextDocument.URI].Text
	line := textLines(text)[params.Position.Line]
	paths, err := extractPathsRegex(line[:params.Position.Character])
	if err != nil {
		return completionItems, nil
	}
	path := paths[len(paths)-1]

	// Suggest path
	var suggestedAbsolutePaths = []string{}
	switch string(path[0]) {
	case "/":
		suggestedAbsolutePaths = absolutePathSuggestions(path)
	case "~":
		suggestedAbsolutePaths = homePathSuggestions(path)
	case ".":
		suggestedAbsolutePaths = relativePathSuggestions(path, params.TextDocument.URI)
	}

	// Format suggested paths
	for _, suggestedAbsolutePath := range suggestedAbsolutePaths {
		_, suggestion := filepath.Split(suggestedAbsolutePath)
		doc := documentPathMarkdown(path+suggestion, suggestedAbsolutePath)

		fileInfo, err := os.Stat(suggestedAbsolutePath)
		if err == nil && fileInfo.IsDir() {
			detail := "📂 Folder"
			kind := protocol.CompletionItemKindFolder
			completionItems = append(completionItems, protocol.CompletionItem{
				Label:  suggestion,
				Kind:   &kind,
				Detail: &detail,
				Documentation: protocol.MarkupContent{
					Kind:  protocol.MarkupKindMarkdown,
					Value: "**📂 Folder**\n" + doc,
				},
				InsertText: &suggestion,
			})
		} else {
			detail := "📄 File"
			kind := protocol.CompletionItemKindFile
			completionItems = append(completionItems, protocol.CompletionItem{
				Label:  suggestion,
				Kind:   &kind,
				Detail: &detail,
				Documentation: protocol.MarkupContent{
					Kind:  protocol.MarkupKindMarkdown,
					Value: "**📄 File**\n" + doc,
				},
				InsertText: &suggestion,
			})
		}
	}
	return completionItems, nil
}

func documentPathMarkdown(inputPath, absolutePath string) string {
	return fmt.Sprintf(`
**Input path:**

*%s*

**Absolute path:**

[*%s*](file://%s)`,
		inputPath, absolutePath, absolutePath)
}

func absolutePathSuggestions(absolutePath string) []string {
	searchPath := filepath.Join(absolutePath, "*")
	suggestedAbsolutePaths, err := filepath.Glob(searchPath)
	if err != nil {
		return []string{}
	}
	return suggestedAbsolutePaths
}

func homePathSuggestions(path string) []string {
	currentUser, err := user.Current()
	if err != nil {
		return []string{}
	}
	absolutePath := filepath.Join(currentUser.HomeDir, path[2:])
	return absolutePathSuggestions(absolutePath)
}

func relativePathSuggestions(path string, fileUri string) []string {
	currentAbsoluteDirPath, _ := filepath.Split(fileUri[7:])
	absolutePath := filepath.Join(currentAbsoluteDirPath, path)
	return absolutePathSuggestions(absolutePath)
}
