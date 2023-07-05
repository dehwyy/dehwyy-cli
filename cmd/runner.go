package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/dehwyy/dehwyy-cli/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var (
	logger = log.New(os.Stdout, "", 0)
	runnerCmd = &cobra.Command{
		Use: "runner [command]",
		Short: "Use is a shortcut to run several commands by using just one",
		Long: "Use is a shortcut to run several commands by using just one",
		Run: runCmdRunner,
	}
	// flags
	add string // adding new template
	exc string // executing template
	del string // deleting template
)

func runCmdRunner(cmd *cobra.Command, args[]string) {
	if !ProvidedLessThanTwoFlags(add, exc, del) {
		fmt.Println("Only one flag should be provided")
		return
	}

	// runner.db instance
	db := new(database.RunnerDB)

	// init and close right before the end of the function
	db.Init()
	defer db.Close()

	// Making table 'commands' if there is no
	db.CreateTableIfNotExists()

	switch {
		// _____________________
		// Adding new templatee
		case len(add) > 0:
			for _, cmd := range args {
				db.AddCommandByKey(cmd, add)
			}

			if len(args) == 0 {
				fmt.Println("At least one command should be provided")
			} else {
				fmt.Printf("Added new template '%s'\n", add)
			}

		// _____________________
		// Querying for template and executing its commands
		case len(exc) > 0:
			rows := db.GetCommandsByKey(exc)

			// if i == 0 => no rows.Next() was ever called => no rows were found
			var i int

			for rows.Next() {
				var command string
				rows.Scan(&command)

				executeCommand(command)

				i++
			}

			if i == 0 {
				fmt.Printf("Template '%s' wasn't found\n", exc)
			} else {
				fmt.Printf("Successfully executed template '%s'\n", exc)
			}

		// _________________
		// deleting template
		case len(del) > 0:
			rowsAffected := db.DeleteTemplateByKey(del)

			if rowsAffected == 0 {
				fmt.Printf("Template '%s' wasn't found\n", del)
			} else {
				fmt.Printf("Successfully deleted template '%s'\n", del)
			}

		// works as a helper
		default:
			rows := db.GetAvailableCommands()

			templates := []string{}
			var template string

			for rows.Next() {
				rows.Scan(&template)

				// appending template with TAB at the beginning
				templates = append(templates, fmt.Sprintf("\t%s", template))
			}
			if len(templates) == 0 {
				fmt.Println("No template yet! Create your first by typing 'dehwyy-cli runner -a [template-name] [...commands] '")
			} else {
				fmt.Printf("Available templates: \n%s\n", strings.Join(templates, "\n"))
			}
	}
}

func executeCommand(command string) {
	// args[0] is main command like "code" or "ls", args[1:] are parameters
	args := strings.Split(command, " ")

	e := exec.Command(args[0],args[1:]...).Start()
	if e != nil {
		logger.Fatalf("Error occured: %v\n", e)
	}
}

func ProvidedLessThanTwoFlags(flags ...string) bool {
	var truthyStringCounter int

	for _, flag := range flags {
		if flag != "" {
			truthyStringCounter++
		}
		if truthyStringCounter > 1 {
			return false
		}
	}
	return true
}

func init() {
	runnerCmd.Flags().StringVarP(&add, "add", "a", "", "Add a template")
	runnerCmd.Flags().StringVarP(&exc, "exec", "e", "", "Execute predefined command by its name")
	runnerCmd.Flags().StringVarP(&del, "del", "d", "", "Delete predefined command by its name")
	rootCmd.AddCommand(runnerCmd)
}
