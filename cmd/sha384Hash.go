package cmd

import (
	"crypto/sha512"

	"github.com/spf13/cobra"
)

// sha384HashCmd represents the sha384Hash command
var sha384HashCmd = &cobra.Command{
	Use:   "sha384 [FILE]",
	Short: "Display SHA-384 checksums (384 bits).",
	Long: `Display SHA-384 checksums (384 bits).

Without FILE or when FILE is '-', read the standard input.
If the list of FILE contains a directory, it can be proceed recursively.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filesToCheck, err := getFilesToCompute(args, recursiveFlag)
		if err != nil {
			return err
		}

		numJobs := len(filesToCheck)
		jobs := make(chan JobsParam, numJobs)
		results := make(chan HashResult, numJobs)

		initWorkers(jobs, results)

		for _, filePath := range filesToCheck {
			h := sha512.New384()
			jobs <- JobsParam{filePath, h}
		}
		close(jobs)

		return waitForResult(numJobs, results)
	},
}

func init() {
	hashCmd.AddCommand(sha384HashCmd)
}
