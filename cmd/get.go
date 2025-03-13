package cmd

import (
	"fmt"
	"github.com/hargeek/argocd-ctx-prompt/pkg/argocd_config"
	"github.com/spf13/cobra"
)

// getCmd 表示 get 命令
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get current ArgoCD context",
	Long:  `Get the current ArgoCD context from the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(argocd_config.GetCurrentContext())
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
