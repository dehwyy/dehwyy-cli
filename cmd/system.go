package cmd

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var (
	systemCmd = &cobra.Command{
		Use: "system [os | hostname | user | homedir]",
		Short: "Fast access to system data",
		Long: "Fast access to information about operation system",
		Args: cobra.RangeArgs(0, 1),
		Run: runCmdSystem,
	}
)

func init() {
	systemCmd.SetHelpTemplate(getHelperString())
	rootCmd.AddCommand(systemCmd)
}

func runCmdSystem(cmd *cobra.Command, args[]string) {
	// if no arg was provided it would reply with helpFunc
	if len(args) == 0 {
		fmt.Println(getHelperString())
		return
	}

	var output string

	switch strings.ToLower(args[0]) {
		case "os":
			output = fmt.Sprintf("Operating system: %s", runtime.GOOS)

		// About Host
		case "hostname":
			hostname, _ := os.Hostname()
			//
			output = fmt.Sprintf("Current hostname: %s", hostname)

		// About User
		case "user":
			user, _ := user.Current()
			//
			output = fmt.Sprintf("Current user: %s", user.Username)

		// CurrentUser Homedir
		case "homedir":
			homedir, _ := os.UserHomeDir()
			//
			output = fmt.Sprintf("Home directory: %s", homedir)

		// if such arg does not exists at previous cases => helpFunc with warning message
		default:
			fmt.Printf("Unknown argument: %s\n\n", args[0])
			output = getHelperString()
	}

	fmt.Println(output)
}


func getHelperString() string {
	/*
		Output is like:
		1. usage
		2. description
	*/

	usageString := "Usage: dehwyy-cli system [os | user | homedir]"


	type cmd struct {
		cmd string
		meaning string
	}

	commands := []cmd{
		{"os", "Returns current OS"},
		{"user", "Returns current user"},
		{"homedir", "Returns home directory"},
	}

	var commandUsageString string
	for _, cmd := range commands {
		commandUsageString += fmt.Sprintf("  %s \t\t%s\n", cmd.cmd, cmd.meaning)
	}


	return fmt.Sprintf("%s\n\n%s", usageString, commandUsageString)
}
