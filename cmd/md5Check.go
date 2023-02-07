package cmd

import (
	"crypto/md5"

	"github.com/spf13/cobra"
)

// md5CheckCmd represents the md5Check command
var md5CheckCmd = &cobra.Command{
	Use:   "md5",
	Short: "Check MD5 checksums (128 bits).",
	Long: `Check MD5 checksums (128 bits).

Write on the standard input, one line at a time.
Each line includes the hash to compare and the path of the file to calculate.
	`,
	Example: "$ hasher check md5\n52016f1168b24120ad9135475f65dd2e debian-11.6.0-amd64-DVD-1.iso",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := md5.New()

		return checkHash(h)
	},
}

func init() {
	checkCmd.AddCommand(md5CheckCmd)
}
