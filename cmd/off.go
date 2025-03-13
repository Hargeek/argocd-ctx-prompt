package cmd

import (
	"fmt"
	"github.com/hargeek/argocd-ctx-prompt/pkg/argocd_config"
	"github.com/spf13/cobra"
)

// offCmd 表示 off 命令
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Disable prompt",
	Long:  `Disable the ArgoCD context prompt in your shell.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := argocd_config.DisablePrompt(); err != nil {
			fmt.Printf("Error disabling prompt: %v\n", err)
			return
		}
		fmt.Println("ArgoCD context prompt disabled")
	},
}

func init() {
	rootCmd.AddCommand(offCmd)
}
