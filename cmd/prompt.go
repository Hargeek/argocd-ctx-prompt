package cmd

import (
	"fmt"
	"github.com/hargeek/argocd-ctx-prompt/pkg/argocd_config"
	"github.com/spf13/cobra"
)

// promptCmd 表示 prompt 命令
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Output formatted prompt string",
	Long:  `Output a formatted prompt string containing the current ArgoCD context.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(argocd_config.FormatPrompt())
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
}
