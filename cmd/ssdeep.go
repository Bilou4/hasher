package cmd

import (
	"fmt"

	"github.com/glaslos/ssdeep"
	"github.com/spf13/cobra"
)

var forceHash bool

// ssdeepCmd represents the ssdeep command
var ssdeepCmd = &cobra.Command{
	Use:   "ssdeep",
	Short: "",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 2),

	RunE: func(cmd *cobra.Command, args []string) error {
		ssdeep.Force = forceHash

		h1, err := ssdeep.FuzzyFilename(args[0])
		if err != nil && !ssdeep.Force {
			return err
		}

		if len(args) == 2 {
			var h2 string
			h2, err = ssdeep.FuzzyFilename(args[1])
			if err != nil && !ssdeep.Force {
				return err
			}

			var score int
			score, err = ssdeep.Distance(h1, h2)
			if score != 0 {
				fmt.Printf("%s matches %s (%d)\n", args[0], args[1], score)
			} else if err != nil {
				fmt.Println("distance error", err)
			} else {
				fmt.Println("The files doesn't match")
			}
		} else {
			fmt.Printf("%s,\"%s\"\n", h1, args[0])
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(ssdeepCmd)
	ssdeepCmd.Flags().BoolVarP(&forceHash, "force", "f", false, "Force hash on error or invalid input length")
}
