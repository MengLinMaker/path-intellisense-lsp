package handlers

import (
	"fmt"
	"log/slog"
	"os"
	"os/user"
	"path/filepath"
	"regexp"

	"path-intellisense-lsp/src/glsp"
	protocol "path-intellisense-lsp/src/protocol_3_16"
)

func absolutePathSuggestions(absolutePath string) []string {
	searchPath := filepath.Join(absolutePath, "*")
	suggestedAbsolutePaths, err := filepath.Glob(searchPath)
	if err != nil {
		return []string{}
	}
	return suggestedAbsolutePaths
}

func relativePathSuggestions(path string, currentAbsoluteFilePath string) []string {
	currentAbsoluteDirPath, _ := filepath.Split(currentAbsoluteFilePath)
	absolutePath := filepath.Join(currentAbsoluteDirPath, path)
	return absolutePathSuggestions(absolutePath)
}

func homePathSuggestions(path string) []string {
	currentUser, err := user.Current()
	if err != nil {
		return []string{}
	}
	absolutePath := filepath.Join(currentUser.HomeDir, path[2:])
	return absolutePathSuggestions(absolutePath)
}

// Get last match of valid file path
func extractPathRegex(text string) (string, error) {
	re := regexp.MustCompile("([.]{1,2}|~)?" + "(/([*]|[^\\/:?\"<>|\r\n])+)*" + "/")
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return "", nil
	}
	path := matches[len(matches)-1]
	if len(path) == 0 {
		return "", nil
	}
	return path, nil
}

func TextDocumentCompletion(ctx *glsp.Context, params *protocol.CompletionParams) (any, error) {
	slog.Debug(fmt.Sprintf("TextDocumentCompletion: %s", params.TextDocument.URI))
	var completionItems []protocol.CompletionItem

	// Validate file path syntax
	text := currentFiles[params.TextDocument.URI].Text
	lines := regexp.MustCompile("\r?\n").Split(text, -1)
	line := lines[params.Position.Line]
	path, err := extractPathRegex(line[:params.Position.Character])
	if err != nil {
		return completionItems, nil
	}

	// Suggest path
	var suggestedAbsolutePaths = []string{}
	switch string(path[0]) {
	case "/":
		suggestedAbsolutePaths = absolutePathSuggestions(path)
	case "~":
		suggestedAbsolutePaths = homePathSuggestions(path)
	case ".":
		absoluteCurrentFilePath := params.TextDocument.URI[7:]
		suggestedAbsolutePaths = relativePathSuggestions(path, absoluteCurrentFilePath)
	}

	// Format suggested paths
	for _, suggestedAbsolutePath := range suggestedAbsolutePaths {
		_, suggestion := filepath.Split(suggestedAbsolutePath)

		doc := fmt.Sprintf(`
**Input path:**

[*%s*](file://%s)

**Absolute path:**

[*%s*](file://%s)`,
			path+suggestion, suggestedAbsolutePath, suggestedAbsolutePath, suggestedAbsolutePath)

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
