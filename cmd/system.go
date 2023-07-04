package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var (
	systemCmd = &cobra.Command{
		Use: "system [command]",
		Short: "Fast access to system data",
		Long: "Fast access to information about operation system",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args[]string) {
			var output string

			switch strings.ToLower(args[0]) {

				case "os":
					output = fmt.Sprintf("Operating system: %s", runtime.GOOS)

				case "user":
					user, err := user.Current()

					if err != nil {
						log.Fatalf("Error occured: %v", err)
						return
					}

					output = fmt.Sprintf("Current user: %s", user.Username)

				case "homedir":
					homedir, err := os.UserHomeDir()

					if err != nil {
						log.Fatalf("Error occured: %v", err)
						return
					}

					output = fmt.Sprintf("Home directory: %s", homedir)
			}


			fmt.Println(output)
		},
	}
)

func init() {
	rootCmd.AddCommand(systemCmd)
}
