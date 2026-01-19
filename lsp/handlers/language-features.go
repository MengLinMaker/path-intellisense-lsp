package handlers

import (
	"path-intellisense-lsp/mappers"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completionItems []protocol.CompletionItem

	for word, emoji := range mappers.EmojiMapper {
		emojiCopy := emoji // Create a copy of emoji
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      word,
			Detail:     &emojiCopy,
			InsertText: &emojiCopy,
		})
	}

	return completionItems, nil
}
