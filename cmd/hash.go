package cmd

import (
	"hash"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var recursiveFlag bool

type HashResult struct {
	res  []byte
	path string
	err  error
}

type JobsParam struct {
	path string
	h    hash.Hash
}

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "The hash command computes the hash of a given FILE.",
	Long: `The hash command computes the hash of a given FILE.

Without FILE or when FILE is '-', read the standard input.
If the list of FILE contains a directory, it can be proceed recursively.`,
	SilenceErrors: true,
}

func init() {
	rootCmd.AddCommand(hashCmd)
	hashCmd.PersistentFlags().BoolVar(&recursiveFlag, "recursive", false, "When this flag is set, every directories in your input will be proceed recursively.")
}

func computeHash(path string, h hash.Hash) ([]byte, error) {
	var r io.Reader
	if path == "-" {
		r = os.Stdin
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		r = file
	}
	_, err := io.Copy(h, r)
	if err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
