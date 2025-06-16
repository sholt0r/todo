package cmd

import (
	"fmt"

	"github.com/sholt0r/todo/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lsCmd)
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List tasks in the todo list",
	RunE: func(cmd *cobra.Command, args []string) error {
		todoList, err := internal.GetActiveList()
		if err != nil {
			return err
		}

		fmt.Printf("%-40s %-10s\n", "Task", "Created")
		for i := range todoList.Tasks {
			task := todoList.Tasks[i]
			fmt.Printf("%-40s %-10s\n", task.Name, task.Name)
		}

		return nil
	},
}
