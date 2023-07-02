package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "dehwyy",
		Short: "CLI by dehwyy",
		Long: `CLI made by dehwyy for some unknown reason`,
	}
)

func Execute() {
	rootCmd.Execute()
}
