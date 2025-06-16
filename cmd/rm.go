package cmd

import (
	"github.com/sholt0r/todo/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a task from the todo list",
	RunE: func(cmd *cobra.Command, args []string) error {
		todoList, err := internal.GetActiveList()
		if err != nil {
			return err
		}

		//todoList.RemoveTask(i)

		return nil
	},
}
