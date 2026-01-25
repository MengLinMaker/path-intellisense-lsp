# agent.md guidelines recommended by steipete - Github Gist https://gist.github.com/steipete/d3b9db3fa8eb1d1a692b7656217d8655
- Delete unused or obsolete files when your changes make them irrelevant (refactors, feature removals, etc.), and revert files only when the change is yours or explicitly requested. If a git operation leaves you unsure about other agents' in-flight work, stop and coordinate instead of deleting.
- **Before attempting to delete a file to resolve a local type/lint failure, stop and ask the user.** Other agents are often editing adjacent files; deleting their work to silence an error is never acceptable without explicit approval.
- NEVER edit `.env` or any environment variable files—only the user may change them.
- Coordinate with other agents before removing their in-progress edits—don't revert or delete work you didn't author unless everyone agrees.
- Moving/renaming and restoring files is allowed.
- ABSOLUTELY NEVER run destructive git operations (e.g., `git reset --hard`, `rm`, `git checkout`/`git restore` to an older commit) unless the user gives an explicit, written instruction in this conversation. Treat these commands as catastrophic; if you are even slightly unsure, stop and ask before touching them. *(When working within Cursor or Codex Web, these git limitations do not apply; use the tooling's capabilities as needed.)*
- Never use `git restore` (or similar commands) to revert files you didn't author—coordinate with other agents instead so their in-progress work stays intact.
- Always double-check git status before any commit
- Keep commits atomic: commit only the files you touched and list each path explicitly. For tracked files run `git commit -m "<scoped message>" -- path/to/file1 path/to/file2`. For brand-new files, use the one-liner `git restore --staged :/ && git add "path/to/file1" "path/to/file2" && git commit -m "<scoped message>" -- path/to/file1 path/to/file2`.
- Quote any git paths containing brackets or parentheses (e.g., `src/app/[candidate]/**`) when staging or committing so the shell does not treat them as globs or subshells.
- When running `git rebase`, avoid opening editors—export `GIT_EDITOR=:` and `GIT_SEQUENCE_EDITOR=:` (or pass `--no-edit`) so the default messages are used automatically.
- Never amend commits unless you have explicit written approval in the task thread.

# Repository Guidelines

## Project Structure & Module Organization
- `lsp/`: Go implementation of the Language Server (entry: `lsp/main.go`). Protocol definitions live under `lsp/protocol_3_16/` and `lsp/protocol_3_17/`; server wiring in `lsp/server/` and handlers in `lsp/handlers/`.
- `extensions/vscode/`: VS Code extension (TypeScript, bundled with esbuild). Output goes to `extensions/vscode/dist/`.
- `extensions/zed/`: Zed extension (Rust, `extension.toml`, sources in `extensions/zed/src/`).
- Root config: `Taskfile.yml`, `biome.json`, `tsconfig.json`.

## Build, Test, and Development Commands
- `task lsp:setup` — run `go mod tidy` for the LSP module.
- `task lsp:fmt` / `task lsp:lint` — `gofmt` and `golangci-lint` for Go.
- `task lsp:build` — build the LSP binary to `lsp/dist/path-intellisense-lsp`.
- `task -d extensions/vscode setup` — install VS Code extension deps with `pnpm`.
- `task -d extensions/vscode build` — bundle extension and produce `dist/extension.vsix`.
- `task -d extensions/zed lint` — run the Zed extension lint step (`cargo c`).
- `pnpm --dir extensions/vscode test:type` — TypeScript typecheck only.

## Coding Style & Naming Conventions
- Go: rely on `gofmt` and `golangci-lint` for formatting and linting.
- VS Code extension: `biome` formats and lints (`pnpm biome format --write`, `pnpm biome lint --write`).
- Keep file and type names descriptive and aligned with existing patterns (e.g., `handlers/*.go`, `protocol_3_17/*.go`).

## Testing Guidelines
- There are no automated unit tests in the repo yet. Use `pnpm --dir extensions/vscode test:type` for TS type safety and run `task lsp:build` to validate Go builds.

## Commit & Pull Request Guidelines
- Commit messages in history are short, imperative, and lowercase (e.g., “categorise handlers”). Follow that style.
- PRs should describe the change, include relevant commands run, and link any related issues. For UI/UX changes in extensions, add a screenshot or short clip.

## Configuration Tips
- The VS Code extension build expects the LSP binary at `lsp/dist/path-intellisense-lsp`; run `task lsp:build` before `task -d extensions/vscode build`.
