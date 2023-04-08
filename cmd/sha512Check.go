package cmd

import (
	"crypto/sha512"

	"github.com/spf13/cobra"
)

// sha512CheckCmd represents the sha512Check command
var sha512CheckCmd = &cobra.Command{
	Use:   "sha512",
	Short: "Check SHA-512 checksums (512 bits).",
	Long: `Check SHA-512 checksums (512 bits).

Write on the standard input, one line at a time.
Each line includes the hash to compare and the path of the file to calculate.
`,
	Example: "$ hasher check sha512\n41735b046219d74832e033205130bce4dbbc2aa72ae81d8143aea278618a638599e1d4b7a0d9ba04f3ff44431208845be3868e313f0c258d9b232423c3f52438 debian-11.6.0-amd64-DVD-1.iso",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := sha512.New()
		return checkHash(h)
	},
}

func init() {
	checkCmd.AddCommand(sha512CheckCmd)
}
