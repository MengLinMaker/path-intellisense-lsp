use zed_extension_api::{self as zed, Command, LanguageServerId, Result, Worktree};

struct PathIntellisenseLspExtension {}

// impl PathIntellisenseLspExtension {
//     const LSP_BINARY_NAME: &'static str = "path-intellisense-lsp";
// }

impl zed::Extension for PathIntellisenseLspExtension {
    fn new() -> Self {
        Self {}
    }

    fn language_server_command(
        &mut self,
        _language_server_id: &LanguageServerId,
        worktree: &Worktree,
    ) -> Result<Command> {
        // let path = worktree
        //     .which(Self::LSP_BINARY_NAME)
        //     .ok_or_else(|| format!("Could not find {} binary", Self::LSP_BINARY_NAME))?;

        Ok(zed::Command {
            command: "/Users/menglinmaker/Documents/software-projects/open-source/path-intellisense-lsp/lsp/out/path-intellisense-lsp".to_string(),
            // command: path,
            args: vec![],
            env: worktree.shell_env(),
        })
    }
}

zed::register_extension!(PathIntellisenseLspExtension);
