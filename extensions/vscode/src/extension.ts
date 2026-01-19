import path from 'node:path'
import { type ExtensionContext, window } from 'vscode'

import {
	LanguageClient,
	type LanguageClientOptions,
	type ServerOptions,
} from 'vscode-languageclient/node'

let client: LanguageClient

const LSP_BINARY = 'path-intellisense-lsp'
const LSP_NAME = 'Path intellisense lsp'

export const activate = async (ctx: ExtensionContext) => {
	const lspModule = ctx.asAbsolutePath(path.join('dist', LSP_BINARY))
	const serverOptions: ServerOptions = {
		run: { module: lspModule },
		debug: { module: lspModule },
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
	// activateInlayHints(ctx);
	client.start()
}

export const deactivate = () => (client ? client.stop() : undefined)

// export function activateInlayHints(ctx: ExtensionContext) {
//   const maybeUpdater = {
//     hintsProvider: null as Disposable | null,
//     updateHintsEventEmitter: new EventEmitter<void>(),

//     async onConfigChange() {
//       this.dispose()

//       // const event = this.updateHintsEventEmitter.event;
//       // this.hintsProvider = languages.registerInlayHintsProvider(
//       //   { scheme: "file", language: "nrs" },
//       //   new (class implements InlayHintsProvider {
//       //     onDidChangeInlayHints = event;
//       //     resolveInlayHint(hint: InlayHint, token: CancellationToken): ProviderResult<InlayHint> {
//       //       const ret = {
//       //         label: hint.label,
//       //         ...hint,
//       //       };
//       //       return ret;
//       //     }
//       //     async provideInlayHints(
//       //       document: TextDocument,
//       //       range: Range,
//       //       token: CancellationToken
//       //     ): Promise<InlayHint[]> {
//       //       const hints = (await client
//       //         .sendRequest("custom/inlay_hint", { path: document.uri.toString() })
//       //         .catch(err => null)) as [number, number, string][];
//       //       if (hints == null) {
//       //         return [];
//       //       } else {
//       //         return hints.map(item => {
//       //           const [start, end, label] = item;
//       //           let startPosition = document.positionAt(start);
//       //           let endPosition = document.positionAt(end);
//       //           return {
//       //             position: endPosition,
//       //             paddingLeft: true,
//       //             label: [
//       //               {
//       //                 value: `${label}`,
//       //                 location: {
//       //                   uri: document.uri,
//       //                   range: new Range(1, 0, 1, 0)
//       //                 }
//       //               command: {
//       //                   title: "hello world",
//       //                   command: "helloworld.helloWorld",
//       //                   arguments: [document.uri],
//       //                 },
//       //               },
//       //             ],
//       //           };
//       //         });
//       //       }
//       //     }
//       //   })()
//       // );
//     },

//     onDidChangeTextDocument() {
//       // debugger
//       // this.updateHintsEventEmitter.fire();
//     },

//     dispose() {
//       this.hintsProvider?.dispose()
//       this.hintsProvider = null
//       this.updateHintsEventEmitter.dispose()
//     },
//   }

//   workspace.onDidChangeConfiguration(
//     maybeUpdater.onConfigChange,
//     maybeUpdater,
//     ctx.subscriptions,
//   )
//   workspace.onDidChangeTextDocument(
//     maybeUpdater.onDidChangeTextDocument,
//     maybeUpdater,
//     ctx.subscriptions,
//   )

//   maybeUpdater.onConfigChange().catch(console.error)
// }
