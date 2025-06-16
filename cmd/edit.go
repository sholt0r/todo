package cmd

import (
	"github.com/sholt0r/todo/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a task in the todo list",
	RunE: func(cmd *cobra.Command, args []string) error {
		todoList, err := internal.GetActiveList()
		if err != nil {
			return err
		}

		//todoList.Tasks[i].Name = args[3]

		return nil
	},
}
