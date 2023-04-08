package cmd

import (
	"crypto/sha512"

	"github.com/spf13/cobra"
)

// sha384CheckCmd represents the sha384Check command
var sha384CheckCmd = &cobra.Command{
	Use:   "sha384",
	Short: "Check SHA-384 checksums (384 bits).",
	Long: `Check SHA-384 checksums (384 bits).

Write on the standard input, one line at a time.
Each line includes the hash to compare and the path of the file to calculate.
`,
	Example: "$ hasher check sha384\n81d2859a3505dd3d22522777b4ab0c893e279c10ca1f1a95fe53a0d7019537b40d479e5b2a18e49b294f8ffce3593d53 debian-11.6.0-amd64-DVD-1.iso",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := sha512.New384()
		return checkHash(h)
	},
}

func init() {
	checkCmd.AddCommand(sha384CheckCmd)
}
