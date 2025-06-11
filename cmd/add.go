package cmd

import (
	"github.com/maslovpi/go-todo-list/servicelayer"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Long:  `Add new task to the ToDo list`,
	Run: func(cmd *cobra.Command, args []string) {
		toDoList := servicelayer.NewToDoList()
		defer toDoList.Finalize()

		toDoList.Add("New Task")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
