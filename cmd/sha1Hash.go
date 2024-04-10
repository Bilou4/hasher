package cmd

import (
	"crypto/sha1"

	"github.com/spf13/cobra"
)

// sha1HashCmd represents the sha1Hash command
var sha1HashCmd = &cobra.Command{
	Use:   "sha1 [FILE]",
	Short: "Display SHA-1 checksums (160 bits).",
	Long: `Display SHA-1 checksums (160 bits).

Without FILE or when FILE is '-', read the standard input.
If the list of FILE contains a directory, it will be proceed recursively.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filesToCheck, err := getFilesToCompute(args)
		if err != nil {
			return err
		}

		numJobs := len(filesToCheck)
		jobs := make(chan JobsParam, numJobs)
		results := make(chan HashResult, numJobs)

		initWorkers(jobs, results)

		for _, filePath := range filesToCheck {
			h := sha1.New()
			jobs <- JobsParam{filePath, h}
		}
		close(jobs)

		return waitForResult(numJobs, results)
	},
}

func init() {
	hashCmd.AddCommand(sha1HashCmd)
}
