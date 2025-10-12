package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ymatsukawa/ged/internal/crypto"
	"github.com/ymatsukawa/ged/internal/fileio"
)

var decryptCmd = &cobra.Command{
	Use:   "dec",
	Short: "decrypt file encoded by AES256-GCM with key",
	Long:  "decrypt file encoded by AES256-GCM with key",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("specify a file")
		}
		path := args[0]

		fmt.Println("decrypting...")
		text, err := fileio.ReadFile(fileio.File{
			TargetPath: path,
		})
		if err != nil {
			return fmt.Errorf("failed to read file: %w", err)
		}

		b, err := crypto.DecryptAES256GCM(text, []byte(Key))
		if err != nil {
			return fmt.Errorf("failed to decrypt: %w", err)
		}

		err = fileio.WriteFile(fileio.File{
			Data:       b,
			TargetPath: path,
			CryptoAs:   fileio.Decrypt,
		})
		if err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}

		fmt.Println("Succcess to decrypt.")
		return nil
	},
}
