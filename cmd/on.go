package cmd

import (
	"fmt"
	"github.com/hargeek/argocd-ctx-prompt/pkg/argocd_config"
	"github.com/spf13/cobra"
)

// onCmd 表示 on 命令
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Enable prompt",
	Long:  `Enable the ArgoCD context prompt in your shell.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := argocd_config.EnablePrompt(); err != nil {
			fmt.Printf("Error enabling prompt: %v\n", err)
			return
		}
		fmt.Println("ArgoCD context prompt enabled")
	},
}

func init() {
	rootCmd.AddCommand(onCmd)
}
