package cmd

import (
	"crypto/sha256"

	"github.com/spf13/cobra"
)

// sha224CheckCmd represents the sha224Check command
var sha224CheckCmd = &cobra.Command{
	Use:   "sha224",
	Short: "Check SHA-224 checksums (224 bits).",
	Long: `Check SHA-224 checksums (224 bits).

Write on the standard input, one line at a time.
Each line includes the hash to compare and the path of the file to calculate.
`,
	Example: "$ hasher check sha224\nf9eb84fe0914b69d9dc39f4c831b60dbc82a0ed7af674282b81e9c7c debian-11.6.0-amd64-DVD-1.iso",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := sha256.New224()
		return checkHash(h)
	},
}

func init() {
	checkCmd.AddCommand(sha224CheckCmd)
}
