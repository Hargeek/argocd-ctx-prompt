package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd 表示在没有子命令时的基本命令
var rootCmd = &cobra.Command{
	Use:   "argocd-ctx-prompt",
	Short: "Display current ArgoCD context in shell prompt",
	Long:  `argocd-ctx-prompt is a command-line tool that displays the current ArgoCD context in your shell prompt.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// 禁用 completion 命令输出
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// 禁用 help 命令输出，只保留 --help 标志
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
}
