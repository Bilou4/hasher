package cmd

import (
	"crypto/sha256"
	"fmt"

	"github.com/spf13/cobra"
)

// sha224HashCmd represents the sha224Hash command
var sha224HashCmd = &cobra.Command{
	Use:   "sha224 [FILE]",
	Short: "Display SHA-224 checksums (224 bits).",
	Long: `Display SHA-224 checksums (224 bits).

Without FILE or when FILE is '-', read the standard input.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filesToCheck := getFilesToCompute(args)
		h := sha256.New224()
		res, err := computeFiles(filesToCheck, h)
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

func init() {
	hashCmd.AddCommand(sha224HashCmd)
}
