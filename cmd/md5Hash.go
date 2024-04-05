package cmd

import (
	"crypto/md5"

	"github.com/spf13/cobra"
)

// md5HashCmd represents the md5Hash command
var md5HashCmd = &cobra.Command{
	Use:   "md5 [FILE]...",
	Short: "Display MD5 checksums (128 bits).",
	Long: `Display MD5 checksums (128 bits).

Without FILE or when FILE is '-', read the standard input.
If the list of FILE contains a directory, it will be proceed recursively.
If the list of FILE contains './...' it will proceed directories recursively from the current directory.`,
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
			h := md5.New()
			jobs <- JobsParam{filePath, h}
		}
		close(jobs)

		return waitForResult(numJobs, results)
	},
}

func init() {
	hashCmd.AddCommand(md5HashCmd)
}
