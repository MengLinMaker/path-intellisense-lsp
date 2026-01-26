# Path intellisense lsp
Providing suggestions for relative, home and absolute paths.

### Capabilities
- Completion: suggest folders and files on "/" trigger.
- Diagnostics: detect file paths that are invalid.
- Documentlink: provide navigation links to matching paths.

And LSP server allows easier integration into different IDEs.

### Development
- Language server is located in `./lsp`
- IDE exrtension clients are located in `./extensions`
  - VSCode - `./extensions/vscode`
  - Zed - `./extensions/zed`

### Limitations
**File paths cannot use spaces**
This tradeoff is reasonable as spaces in filepaths will interfere with cli scripts.
Additionally, file path detection requires *trigger characters* like `"`, `'`, ```, `" "`. Using a space as trigger character enables more linting capabilities at the cost of being incompatible with filepaths with spaces.
