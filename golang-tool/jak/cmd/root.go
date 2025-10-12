package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jak",
	Short: "HTTP client tool",
	Long: `A HTTP client tool for making single and batch requests.

Examples:
  jak req GET https://example.com
  jak bat config.toml
  jak chain config.toml`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newReqSimpleCmd())
	rootCmd.AddCommand(newReqBatCmd())
	rootCmd.AddCommand(newReqChainCmd())
}
