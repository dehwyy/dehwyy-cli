package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "prettify",
		Short: "CLI tool for print text",
		Long: `CLI tool for print text in different styles`,
	}
)

func Execute() {
	rootCmd.Execute()
}
