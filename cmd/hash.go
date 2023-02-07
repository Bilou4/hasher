package cmd

import (
	"fmt"
	"hash"

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

func getFilesToCompute(args []string) []string {
	if len(args) == 0 || (len(args) == 1 && args[0] == "-") {
		return []string{"-"}
	}
	return args
}

func computeFiles(filesToCheck []string, h hash.Hash) (string, error) {
	var res string
	for _, filePath := range filesToCheck {
		hashedValue, err := computeHash(filePath, h)
		if err != nil {
			return "", err
		}
		res += fmt.Sprintf("%x %s\n", hashedValue, filePath)
		h.Reset()
	}
	return res, nil
}
