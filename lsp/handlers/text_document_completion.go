package handlers

import (
	"fmt"
	"log/slog"
	"path-intellisense-lsp/mappers"

	"path-intellisense-lsp/glsp"
	protocol "path-intellisense-lsp/protocol_3_16"
)

func TextDocumentCompletion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completionItems []protocol.CompletionItem

	slog.Info(fmt.Sprintf("File path: %s", params.TextDocument.URI))

	for word, emoji := range mappers.EmojiMapper {
		emojiCopy := emoji
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      word,
			Detail:     &emojiCopy,
			InsertText: &emojiCopy,
		})
	}

	return completionItems, nil
}
