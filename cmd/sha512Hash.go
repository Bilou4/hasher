package cmd

import (
	"crypto/sha512"
	"fmt"

	"github.com/spf13/cobra"
)

// sha512HashCmd represents the sha512Hash command
var sha512HashCmd = &cobra.Command{
	Use:   "sha512 [FILE]",
	Short: "Display SHA-512 checksums (512 bits).",
	Long: `Display SHA-512 checksums (512 bits).

Without FILE or when FILE is '-', read the standard input.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filesToCheck := getFilesToCompute(args)
		h := sha512.New()
		res, err := computeFiles(filesToCheck, h)
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

func init() {
	hashCmd.AddCommand(sha512HashCmd)
}
