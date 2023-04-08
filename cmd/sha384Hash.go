package cmd

import (
	"crypto/sha512"
	"fmt"

	"github.com/spf13/cobra"
)

// sha384HashCmd represents the sha384Hash command
var sha384HashCmd = &cobra.Command{
	Use:   "sha384 [FILE]",
	Short: "Display SHA-384 checksums (384 bits).",
	Long: `Display SHA-384 checksums (384 bits).

Without FILE or when FILE is '-', read the standard input.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filesToCheck := getFilesToCompute(args)
		h := sha512.New384()
		res, err := computeFiles(filesToCheck, h)
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

func init() {
	hashCmd.AddCommand(sha384HashCmd)
}
