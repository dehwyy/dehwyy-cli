package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)


var (
	printCmd = &cobra.Command{
		Use: "print [...strings]",
		Short: "Print strings with different formatting",
		Long: "Print strings with different formatting in console",
		Args: cobra.MinimumNArgs(1),
		Run: runCmdPrint,
	}
	// flags
	stairs bool
	capitalize bool
)

func runCmdPrint(cmd *cobra.Command, args[]string) {
	s := strings.Join(args, " ")
	switch  {
	case stairs:
		s = StairsCase(s)
	}
	fmt.Printf("%s\n", s)
}

func init() {
	printCmd.Flags().BoolVarP(&stairs, "stairs", "s", false, "Print in stairsCase, e: stairs => StAiRs")
	printCmd.Flags().BoolVarP(&capitalize, "capital", "c", false, "Print capitalized, e: cap => Cap")
	rootCmd.AddCommand(printCmd)
}

func StairsCase(s string) string {
	var result string
	c := 0
	for _, char := range strings.ToLower(s) {
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

func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(s)
	// from LowerCase to UpperCase
	runes[0] = runes[0] - 32
	return string(runes)
}
