English | [简体中文](README_zh.md)

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)
[![GoDoc](https://godoc.org/github.com/hargeek/argocd-ctx-prompt?status.svg)](https://pkg.go.dev/github.com/hargeek/argocd-ctx-prompt)
[![Contributors](https://img.shields.io/github/contributors/hargeek/argocd-ctx-prompt)](https://github.com/hargeek/argocd-ctx-prompt/graphs/contributors)
[![License](https://img.shields.io/github/license/hargeek/argocd-ctx-prompt)](./LICENSE)

# argocd-ctx-prompt

A simple command-line tool that displays the current ArgoCD context in your shell prompt.

Inspired by the [kube-ps1](https://github.com/jonmosco/kube-ps1) project

## Features

- Reads ArgoCD configuration file (~/.config/argocd/config)
- Parses and displays the current active ArgoCD context
- Provides formatted output that can be integrated into your shell prompt
- Supports toggle functionality to enable/disable the prompt display at any time

## Installation

There are two installation methods available:

### Method 1: Download Pre-compiled Binary (Recommended)

Download the pre-compiled binary for your system from the [GitHub Releases](https://github.com/hargeek/argocd-ctx-prompt/releases) page. Choose the appropriate version for your operating system and architecture (e.g., Linux/macOS, amd64/arm64, etc.).

After downloading, extract and move the binary to a directory in your PATH:

```bash
# Extract the downloaded file
tar -xzf argocd-ctx-prompt_*.tar.gz

# Move to a directory in your PATH
sudo mv argocd-ctx-prompt /usr/local/bin/

# Add execute permissions
chmod +x /usr/local/bin/argocd-ctx-prompt
```

### Method 2: Install with Go (Requires Go Environment)

If you have Go installed, use the Go install command to install the tool to your system's executable directory:

```bash
go install github.com/hargeek/argocd-ctx-prompt@latest
```

This command will install the `argocd-ctx-prompt` binary to your Go bin directory (typically `$GOPATH/bin` or `$HOME/go/bin`). Make sure this directory is added to your system's PATH environment variable so you can execute the command directly.

If your Go bin directory is not in your PATH, you can add it using:

```bash
# For Bash, add to ~/.bashrc
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc

# For Zsh, add to ~/.zshrc
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.zshrc
source ~/.zshrc
```

## Usage

### Basic Usage

```bash
# Output formatted prompt
argocd-ctx-prompt prompt

# Get current context
argocd-ctx-prompt get

# Enable prompt
argocd-ctx-prompt on

# Disable prompt
argocd-ctx-prompt off

# Toggle prompt state
argocd-ctx-prompt toggle
```

### Shell Integration

Add the following to your shell configuration file (`~/.bashrc`, `~/.bash_profile`, or `~/.zshrc`):

```bash
# ArgoCD context prompt
argocd_ctx_prompt() {
  if command -v argocd-ctx-prompt &> /dev/null; then
    echo "$(argocd-ctx-prompt prompt)"
  fi
}

# With color (green)
argocd_ctx_prompt() {
  if command -v argocd-ctx-prompt &> /dev/null; then
    echo "%{$fg[green]%}$(argocd-ctx-prompt prompt)%{$reset_color%}"
  fi
}

# Add ArgoCD context to prompt
PROMPT='$(argocd_ctx_prompt) '$PROMPT  # For Zsh
PS1='$(argocd_ctx_prompt) '$PS1        # For Bash

# Add aliases for convenience
alias argocdon="argocd-ctx-prompt on"
alias argocdoff="argocd-ctx-prompt off"
```

## License

[MIT](LICENSE)
