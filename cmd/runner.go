package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/dehwyy/dehwyy-cli/database"
	sql_database "github.com/dehwyy/dehwyy-cli/database/sql"
	e "github.com/dehwyy/dehwyy-cli/error-handler"
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
	flagAdd string // adding new template
	flagExecute string // executing template
	flagDelete string // deleting template
)

func init() {
	runnerCmd.Flags().StringVarP(&flagAdd, "add", "a", "", "Add a template")
	runnerCmd.Flags().StringVarP(&flagExecute, "exec", "e", "", "Execute predefined command by its name")
	runnerCmd.Flags().StringVarP(&flagDelete, "delete", "d", "", "Delete predefined command by its name")
	rootCmd.AddCommand(runnerCmd)
}


func runCmdRunner(cmd *cobra.Command, args[]string) {
	// clarifying that either zero or one flag was provided
	if !providedLessThanTwoFlags(flagAdd, flagExecute, flagDelete) {
		fmt.Println("Only one flag should be provided")
		return
	}

	// db instance
	sql_db := new(sql_database.Sqlite)

	// runner queries instance
	db := database.New(sql_db)
	//  close right before the end of the function
	defer db.Close()

	// Making table 'commands' if there is no
	db.CreateTableIfNotExists()

	switch {
		// _____________________
		// Adding new templatee
		case len(flagAdd) > 0:
			for _, cmd := range args {
				db.AddCommandByKey(cmd, flagAdd)
			}

			if len(args) == 0 {
				fmt.Println("At least one command should be provided")
			} else {
				fmt.Printf("Added new template '%s'\n", flagAdd)
			}

		// _____________________
		// Querying for template and executing its commands
		case len(flagExecute) > 0:
			rows := db.GetCommandsByKey(flagExecute)

			// if i == 0 => no rows.Next() was ever called => no rows were found
			var i int

			for rows.Next() {
				var command string
				rows.Scan(&command)

				executeCommand(command)

				i++
			}

			if i == 0 {
				fmt.Printf("Template '%s' wasn't found\n", flagExecute)
			} else {
				fmt.Printf("Successfully executed template '%s'\n", flagExecute)
			}

		// _________________
		// deleting template
		case len(flagDelete) > 0:
			rowsAffected := db.DeleteTemplateByKey(flagDelete)

			if rowsAffected == 0 {
				fmt.Printf("Template '%s' wasn't found\n", flagDelete)
			} else {
				fmt.Printf("Successfully deleted template '%s'\n", flagDelete)
			}

		// works as a helper (-- help)
		default:
			rows := db.GetAvailableCommands()

			templates := []string{}
			var template string

			for rows.Next() {
				rows.Scan(&template)

				// appending template with TAB at the beginning
				templates = append(templates, fmt.Sprintf("\t%s", template))
			}

			// if there is no any templates => printing HelpingCommand
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

	cmd := exec.Command(args[0],args[1:]...)
	e.WithFatalString(cmd.Start(), "Error occured")
}

func providedLessThanTwoFlags(flags ...string) bool {
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
