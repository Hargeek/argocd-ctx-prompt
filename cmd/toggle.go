package cmd

import (
	"fmt"
	"github.com/hargeek/argocd-ctx-prompt/pkg/argocd_config"
	"github.com/spf13/cobra"
)

// toggleCmd 表示 toggle 命令
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggle prompt state",
	Long:  `Toggle the prompt between enabled and disabled states.`,
	Run: func(cmd *cobra.Command, args []string) {
		if argocd_config.IsPromptEnabled() {
			if err := argocd_config.DisablePrompt(); err != nil {
				fmt.Printf("Error disabling prompt: %v\n", err)
				return
			}
			fmt.Println("ArgoCD context prompt disabled")
		} else {
			if err := argocd_config.EnablePrompt(); err != nil {
				fmt.Printf("Error enabling prompt: %v\n", err)
				return
			}
			fmt.Println("ArgoCD context prompt enabled")
		}
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
