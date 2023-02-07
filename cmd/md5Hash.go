package cmd

import (
	"crypto/md5"
	"fmt"

	"github.com/spf13/cobra"
)

// md5HashCmd represents the md5Hash command
var md5HashCmd = &cobra.Command{
	Use:   "md5 [FILE]...",
	Short: "Display MD5 checksums (128 bits).",
	Long: `Display MD5 checksums (128 bits).

Without FILE or when FILE is '-', read the standard input.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filesToCheck := getFilesToCompute(args)
		h := md5.New()
		res, err := computeFiles(filesToCheck, h)
		if err != nil {
			return err
		}
		fmt.Print(res)
		return nil
	},
}

func init() {
	hashCmd.AddCommand(md5HashCmd)
}
