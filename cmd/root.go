package cmd

import (
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "hasher",
	Short:         "",
	Long:          ``,
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func getFilesToCompute(args []string) ([]string, error) {
	if len(args) == 0 || (len(args) == 1 && args[0] == "-") {
		return []string{"-"}, nil
	}

	// check for dirs and all files exist
	for idx, elem := range args {
		f, err := os.Open(elem)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		fInfo, err := f.Stat()
		if err != nil {
			return nil, err
		}
		if fInfo.IsDir() {
			args = remove(args, idx)
			args, err = computeDirRecursively(args, elem)
			if err != nil {
				return nil, err
			}
		}
	}
	fmt.Printf("Computing hash for: %d files.\n", len(args))

	return args, nil
}

func computeDirRecursively(l []string, rootFolder string) ([]string, error) {
	err := filepath.WalkDir(rootFolder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			l = append(l, path)
		}
		return nil
	})
	return l, err
}

func initWorkers(jobs chan JobsParam, results chan HashResult) {
	for i := 0; i <= runtime.NumCPU(); i++ {
		go worker(jobs, results)
	}
}

func waitForResult(nbJobs int, results chan HashResult) error {
	for a := 1; a <= nbJobs; a++ {
		res := <-results
		if res.err != nil {
			return res.err
		}
		fmt.Printf("%s %s\n", hex.EncodeToString(res.res), res.path)
	}
	return nil
}

func worker(jobs <-chan JobsParam, results chan<- HashResult) {
	for j := range jobs {
		hashedValue, err := computeHash(j.path, j.h)
		results <- HashResult{hashedValue, j.path, err}
	}
}
