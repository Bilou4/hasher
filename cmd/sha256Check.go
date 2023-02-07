package cmd

import (
	"crypto/sha256"

	"github.com/spf13/cobra"
)

// sha256CheckCmd represents the sha256Check command
var sha256CheckCmd = &cobra.Command{
	Use:   "sha256",
	Short: "Check SHA256 checksums (256 bits).",
	Long: `Check SHA256 checksums (256 bits).

Write on the standard input, one line at a time.
Each line includes the hash to compare and the path of the file to calculate.
	`,
	Example: "$ hasher check sha256\n55f6f49b32d3797621297a9481a6cc3e21b3142f57d8e1279412ff5a267868d8 debian-11.6.0-amd64-DVD-1.iso",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := sha256.New()

		return checkHash(h)
	},
}

func init() {
	checkCmd.AddCommand(sha256CheckCmd)
}
