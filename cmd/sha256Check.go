package cmd

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// sha256CheckCmd represents the sha256Check command
var sha256CheckCmd = &cobra.Command{
	Use:   "sha256",
	Short: "Check SHA256 checksums (256 bits).",
	Long: `Check SHA256 checksums (256 bits).

Write on the standard input, one line at a time.
Each line includes the hash to compare and the path of the file to calculate.
	`,
	Example: "$ hasher check sha256\n55f6f49b32d3797621297a9481a6cc3e21b3142f57d8e1279412ff5a267868d8 debian-11.6.0-amd64-DVD-1.iso",
	RunE: func(cmd *cobra.Command, args []string) error {
		h := sha256.New()
		reader := bufio.NewReader(os.Stdin)
		for {
			input, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				return err
			}
			input = strings.ReplaceAll(input, "\n", "")
			inputs := strings.Split(input, " ")
			if len(inputs) != 2 {
				return fmt.Errorf("malformed entry: HASH FilePath")
			}
			hashToCheck := strings.Trim(inputs[0], " ")
			path := strings.Trim(inputs[1], " ")
			if len(hashToCheck) != 64 {
				return fmt.Errorf("malformed hash. Expected 64 characters, got %d", len(hashToCheck))
			}
			hash, err := computeHash(path, h)
			if err != nil {
				return err
			}
			if hashToCheck != fmt.Sprintf("%x", hash) {
				return fmt.Errorf("%s FAILED", path)
			}
			fmt.Printf("%s OK\n", path)
			h.Reset()
		}
		return nil
	},
}

func init() {
	checkCmd.AddCommand(sha256CheckCmd)
}
