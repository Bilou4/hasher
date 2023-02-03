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
		var filesToCheck []string
		if len(args) == 0 || (len(args) == 1 && args[0] == "-") {
			filesToCheck = append(filesToCheck, "-")
		} else {
			filesToCheck = args
		}
		h := sha256.New()
		for _, filePath := range filesToCheck {
			hashedValue, err := computeHash(filePath, h)
			if err != nil {
				return err
			}
			fmt.Printf("%x %s\n", hashedValue, filePath)
			h.Reset()
		}
		return nil
	},
}

func init() {
	hashCmd.AddCommand(sha256HashCmd)
}
