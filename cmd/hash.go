package cmd

import (
	"github.com/spf13/cobra"
)

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:           "hash",
	Short:         "",
	Long:          ``,
	SilenceErrors: true,
}

func init() {
	rootCmd.AddCommand(hashCmd)
}
