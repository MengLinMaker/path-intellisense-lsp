import path from 'node:path'
import { type ExtensionContext, window } from 'vscode'

import {
	LanguageClient,
	TransportKind,
	type LanguageClientOptions,
	type ServerOptions,
} from 'vscode-languageclient/node'

let client: LanguageClient

const LSP_BINARY = 'path-intellisense-lsp'
const LSP_NAME = 'Path intellisense lsp'

export const activate = async (ctx: ExtensionContext) => {
	const lspBinaryPath = ctx.asAbsolutePath(path.join('dist', LSP_BINARY))
	const serverOptions: ServerOptions = {
		run: {
			command: lspBinaryPath,
			transport: TransportKind.stdio,
			options: {
				env: {
					LOG_LEVEL: process.env['LOG_LEVEL'] ?? 'INFO',
				},
			},
		},
		debug: {
			command: lspBinaryPath,
			transport: TransportKind.stdio,
			options: {
				env: {
					LOG_LEVEL: 'DEBUG',
				},
			},
		},
	}
	const clientOptions: LanguageClientOptions = {
		documentSelector: [{ scheme: 'file' }],
		traceOutputChannel: window.createOutputChannel(`${LSP_NAME} trace`),
	}
	client = new LanguageClient(
		LSP_BINARY,
		LSP_NAME,
		serverOptions,
		clientOptions,
	)
	client.start()
}

export const deactivate = () => (client ? client.stop() : undefined)
