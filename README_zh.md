# argocd-ctx-prompt

一个简单的命令行工具，用于在命令行提示符中显示当前的 ArgoCD 上下文。灵感来源于 [kube-ps1](https://github.com/jonmosco/kube-ps1) 项目。

[English](README.md) | 简体中文

## 功能

- 读取 ArgoCD 配置文件 (~/.config/argocd/config)
- 解析并显示当前活动的 ArgoCD 上下文
- 提供格式化的输出，可以集成到命令行提示符中
- 支持开关功能，可以随时启用或禁用提示符显示

## 安装

有两种安装方式可供选择：

### 方式 1：下载预编译的二进制文件（推荐）

从 [GitHub Releases](https://github.com/hargeek/argocd-ctx-prompt/releases) 页面下载适合系统的预编译二进制文件。根据操作系统和架构（如 Linux/macOS、amd64/arm64 等）选择对应的版本。

下载后，解压并将二进制文件移动到系统 PATH 目录中：

```bash
# 解压下载的文件
tar -xzf argocd-ctx-prompt_*.tar.gz

# 移动到系统 PATH 目录
sudo mv argocd-ctx-prompt /usr/local/bin/

# 添加执行权限
chmod +x /usr/local/bin/argocd-ctx-prompt
```

### 方式 2：使用 Go 安装（需要 Go 环境）

如果已安装 Go 环境，可以使用 Go 安装命令将工具安装到系统的可执行目录中：

```bash
go install github.com/hargeek/argocd-ctx-prompt@latest
```

这个命令会将 `argocd-ctx-prompt` 二进制文件安装到 Go 的 bin 目录（通常是 `$GOPATH/bin` 或 `$HOME/go/bin`）。确保这个目录已添加到系统的 PATH 环境变量中，以便可以直接执行该命令。

如果 Go 的 bin 目录不在 PATH 中，可以通过以下方式添加：

```bash
# 对于 Bash，添加到 ~/.bashrc
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc

# 对于 Zsh，添加到 ~/.zshrc
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.zshrc
source ~/.zshrc
```

## 使用方法

### 基本用法

```bash
# 输出格式化的提示符
argocd-ctx-prompt prompt

# 获取当前上下文
argocd-ctx-prompt get

# 启用提示符
argocd-ctx-prompt on

# 禁用提示符
argocd-ctx-prompt off

# 切换提示符状态
argocd-ctx-prompt toggle
```

### Shell 集成

将以下内容添加到您的 Shell 配置文件中（`~/.bashrc`、`~/.bash_profile` 或 `~/.zshrc`）：

```bash
# ArgoCD 上下文提示符
argocd_ctx_prompt() {
  if command -v argocd-ctx-prompt &> /dev/null; then
    echo "$(argocd-ctx-prompt prompt)"
  fi
}

# 带颜色版本（绿色）
argocd_ctx_prompt() {
  if command -v argocd-ctx-prompt &> /dev/null; then
    echo "%{$fg[green]%}$(argocd-ctx-prompt prompt)%{$reset_color%}"
  fi
}

# 将 ArgoCD 上下文添加到提示符
PROMPT='$(argocd_ctx_prompt) '$PROMPT  # Zsh 用户
PS1='$(argocd_ctx_prompt) '$PS1        # Bash 用户

# 添加别名，方便开关提示符
alias argocdon="argocd-ctx-prompt on"
alias argocdoff="argocd-ctx-prompt off"
```

## 许可证

MIT 