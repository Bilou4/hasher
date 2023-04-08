package cmd

import (
	"crypto/sha1"

	"github.com/spf13/cobra"
)

// sha1CheckCmd represents the sha1Check command
var sha1CheckCmd = &cobra.Command{
	Use:   "sha1",
	Short: "Check SHA-1 checksums (160 bits).",
	Long: `Check SHA-1 checksums (160 bits).

Write on the standard input, one line at a time.
Each line includes the hash to compare and the path of the file to calculate.
`,
	Example: "$ hasher check sha1\na080ff1d814cbcede3bb8938c8c1953452ce9b27 debian-11.6.0-amd64-DVD-1.iso",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := sha1.New()
		return checkHash(h)
	},
}

func init() {
	checkCmd.AddCommand(sha1CheckCmd)
}
