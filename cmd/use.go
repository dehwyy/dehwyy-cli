package cmd

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	logger = log.New(os.Stdout, "", 0)
	useCmd = &cobra.Command{
		Use: "use [command]",
		Short: "Use is a shortcut to run several commands by using just one",
		Long: "Use is a shortcut to run several commands by using just one",
		Args: cobra.MinimumNArgs(1),
		Run: runCmdUse,
	}
)

func runCmdUse(cmd *cobra.Command, args[]string) {
	for _, s := range args {
		args:= strings.Split(s, " ")
		e := exec.Command(args[0],args[1:]...).Start()
		if e != nil {
			logger.Fatalf("Error occured: %v", e)
		}
	}
}

func init() {
	rootCmd.AddCommand(useCmd)
}
