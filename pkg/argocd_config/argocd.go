package argocd_config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// ArgoCDConfig ArgoCD 配置文件的结构
type ArgoCDConfig struct {
	CurrentContext string               `yaml:"current-context"`
	Contexts       []ArgoCDContextEntry `yaml:"contexts"`
	Servers        []ArgoCDServerEntry  `yaml:"servers"`
}

// ArgoCDContextEntry ArgoCD 上下文条目
type ArgoCDContextEntry struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
	User   string `yaml:"user"`
}

// ArgoCDServerEntry ArgoCD Server条目
type ArgoCDServerEntry struct {
	Name   string `yaml:"name"`
	Server string `yaml:"server"`
}

// DefaultSymbol 默认的 ArgoCD 图标 (Unicode \ue734，需要 Nerd Font 支持)
const DefaultSymbol = "\ue734"

// getStateFilePath 返回状态文件的路径
func getStateFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".config", "argocd-ctx-prompt", "state")
}

// ensureStateDir 确保状态目录存在
func ensureStateDir() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	stateDir := filepath.Join(homeDir, ".config", "argocd-ctx-prompt")
	return os.MkdirAll(stateDir, 0755)
}

// IsPromptEnabled 检查提示符是否启用
func IsPromptEnabled() bool {
	stateFile := getStateFilePath()
	if stateFile == "" {
		return true // 默认启用
	}

	data, err := os.ReadFile(stateFile)
	if err != nil {
		if os.IsNotExist(err) {
			return true // 文件不存在时默认启用
		}
		return true // 读取错误时默认启用
	}

	return string(data) != "off"
}

// EnablePrompt 启用提示符
func EnablePrompt() error {
	if err := ensureStateDir(); err != nil {
		return err
	}
	stateFile := getStateFilePath()
	return os.WriteFile(stateFile, []byte("on"), 0644)
}

// DisablePrompt 禁用提示符
func DisablePrompt() error {
	if err := ensureStateDir(); err != nil {
		return err
	}
	stateFile := getStateFilePath()
	return os.WriteFile(stateFile, []byte("off"), 0644)
}

// GetArgoCDConfigPath 返回 ArgoCD 配置文件的路径
func GetArgoCDConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".config", "argocd", "config")
}

// getCurrentContextImpl 从 ArgoCD 配置文件获取当前上下文的实现
func getCurrentContextImpl() string {
	configPath := GetArgoCDConfigPath()
	if configPath == "" {
		return ""
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return ""
	}

	var config ArgoCDConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return ""
	}

	return config.CurrentContext
}

// GetCurrentContext 是一个可以在测试中被替换的变量
var GetCurrentContext = getCurrentContextImpl

// FormatPrompt 格式化提示符输出
func FormatPrompt() string {
	// 如果提示符被禁用，返回空字符串
	if !IsPromptEnabled() {
		return ""
	}

	ctx := GetCurrentContext()
	if ctx == "" {
		return ""
	}

	// 使用图标和括号包裹上下文
	return fmt.Sprintf("(%s)", ctx)
}
