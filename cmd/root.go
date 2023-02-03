package cmd

import (
	"bufio"
	"fmt"
	"hash"
	"io"
	"os"

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

func computeHash(path string, h hash.Hash) ([]byte, error) {
	if path == "-" {
		reader := bufio.NewReader(os.Stdin)
		buffer := make([]byte, 1024)
		for {
			n, err := reader.Read(buffer)
			h.Write(buffer[:n])
			if err == io.EOF {
				break
			}
		}
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(h, file); err != nil {
			return nil, err
		}
	}
	return h.Sum(nil), nil
}
