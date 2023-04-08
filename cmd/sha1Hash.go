package cmd

import (
	"crypto/sha1"
	"fmt"

	"github.com/spf13/cobra"
)

// sha1HashCmd represents the sha1Hash command
var sha1HashCmd = &cobra.Command{
	Use:   "sha1 [FILE]",
	Short: "Display SHA-1 checksums (160 bits).",
	Long: `Display SHA-1 checksums (160 bits).

Without FILE or when FILE is '-', read the standard input.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filesToCheck := getFilesToCompute(args)
		h := sha1.New()
		res, err := computeFiles(filesToCheck, h)
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

func init() {
	hashCmd.AddCommand(sha1HashCmd)
}
