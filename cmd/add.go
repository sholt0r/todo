package cmd

import (
	"strings"

	"github.com/sholt0r/todo/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the todo list",
	RunE: func(cmd *cobra.Command, args []string) error {
		todoList, err := internal.GetActiveList()
		if err != nil {
			return err
		}

		todoList.AddTask(strings.Join(args, " "))

		return nil
	},
}
