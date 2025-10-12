package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Key string
)

var rootCmd = &cobra.Command{
	Use:   "ged",
	Short: "minimal enc/dec tool using AES256-GCM",
	Long: `minimal enc/dec tool using AES256-GCM.

Examples:
  ged enc ./path/to/myfile.txt -k mykey     # out ./path/to/myfile.txt.enc
	ged dec ./path/to/myfile.txt.enc -k mykey # out ./path/to/myfile.txt.dec`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
	encryptCmd.Flags().StringVarP(&Key, "key", "k", "", "key to encrypt")
	encryptCmd.MarkFlagRequired("key")

	decryptCmd.Flags().StringVarP(&Key, "key", "k", "", "key to decrypt")
	decryptCmd.MarkFlagRequired("key")
}
