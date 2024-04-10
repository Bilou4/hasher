package cmd

import (
	"crypto/sha256"

	"github.com/spf13/cobra"
)

// sha256HashCmd represents the sha256Hash command
var sha256HashCmd = &cobra.Command{
	Use:   "sha256 [FILE]...",
	Short: "Display SHA256 checksums (256 bits).",
	Long: `Display SHA256 checksums (256 bits).

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
			h := sha256.New()
			jobs <- JobsParam{filePath, h}
		}
		close(jobs)

		return waitForResult(numJobs, results)
	},
}

func init() {
	hashCmd.AddCommand(sha256HashCmd)
}
