package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)


var (
	outputCmd = &cobra.Command{
		Use: "output",
		Short: "Print strings with different formatting",
		Long: "Print strings with different formatting in console",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args[]string) {
			s := strings.Join(args, " ")
			switch  {
			case stairs:
				s = StairsCase(s)
			}
			fmt.Printf("%s\n", s)
		},
	}

	stairs bool
)

func init() {
	outputCmd.Flags().BoolVarP(&stairs, "stairs", "s", false, "Print in stairsCase, e: sTaIrS")
	rootCmd.AddCommand(outputCmd)
}

func StairsCase(s string) string {
	var result string
	c := 0
	for _, char := range s {
		if c%2 == 1 {
			result += string(char)
		} else {
			result += strings.ToUpper(string(char))
		}
		// " " (Escape) in string matches 32 in rune type,
		// if it's not 32 then we should increase "c"
		if char != 32 {
			c++
		}
	}
	return result
}
