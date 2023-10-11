package cmd

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/glaslos/ssdeep"
	"github.com/spf13/cobra"
)

var forceHash bool
var minScore uint

type ssdeepRes struct {
	file  string
	hash  string
	score int
}

// ssdeepCmd represents the ssdeep command
var ssdeepCmd = &cobra.Command{
	Use:   "ssdeep",
	Short: "Compares files with fuzzy hashing.",
	Long:  `Compares files with fuzzy hashing. The first argument is the file of reference, others are compared against this reference.`,
	Args:  cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		ssdeep.Force = forceHash

		refHash, err := ssdeep.FuzzyFilename(args[0])
		if err != nil && !ssdeep.Force {
			return err
		}
		var res []ssdeepRes = make([]ssdeepRes, 0, len(args))

		fmt.Println(args[0], refHash)
		fmt.Println()
		for i := 1; i < len(args); i++ {
			currentFile := args[i]
			var h string
			h, err = ssdeep.FuzzyFilename(currentFile)
			if err != nil && !ssdeep.Force {
				fmt.Println("error computing hash", currentFile)
				continue
			}

			var score int
			score, err = ssdeep.Distance(refHash, h)
			if err != nil {
				fmt.Println("error computing distance", currentFile)
				continue
			}
			res = append(res, ssdeepRes{
				file:  currentFile,
				hash:  h,
				score: score,
			})
		}

		sort.Slice(res, func(i, j int) bool {
			return res[i].score < res[j].score
		})

		if cmd.Flags().Lookup("min").Changed {
			res = Filter(res, func(r ssdeepRes) bool {
				return r.score >= int(minScore)
			})
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)
		for _, r := range res {
			fmt.Fprintf(w, "%s\t%s\t%d\t\n", r.file, r.hash, r.score)
		}
		w.Flush()
		return nil
	},
}

func Filter[T any](vs []T, f func(T) bool) []T {
	filtered := make([]T, 0)
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func init() {
	rootCmd.AddCommand(ssdeepCmd)
	ssdeepCmd.Flags().BoolVarP(&forceHash, "force", "f", false, "Force hash on error or invalid input length")
	ssdeepCmd.Flags().UintVar(&minScore, "min", 0, "Minimum score between two files to display the result")
}
