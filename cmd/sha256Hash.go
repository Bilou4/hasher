package cmd

import (
	"crypto/sha256"
	"fmt"

	"github.com/spf13/cobra"
)

// sha256HashCmd represents the sha256Hash command
var sha256HashCmd = &cobra.Command{
	Use:   "sha256 [FILE]...",
	Short: "Display SHA256 checksums (256 bits).",
	Long: `Display SHA256 checksums (256 bits).

Without FILE or when FILE is '-', read the standard input.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filesToCheck := getFilesToCompute(args)
		h := sha256.New()
		res, err := computeFiles(filesToCheck, h)
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

func init() {
	hashCmd.AddCommand(sha256HashCmd)
}
