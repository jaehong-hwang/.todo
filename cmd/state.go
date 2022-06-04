package cmd

import (
	"github.com/jaehong-hwang/todo/todo"
	"github.com/spf13/cobra"
)

var (
	stateCmd = &cobra.Command{
		Use:   "state",
		Short: "update state",
	}

	waitCmd = &cobra.Command{
		Use:   "wait",
		Short: "todo set waiting state",
		RunE:  getUpdatingStateAction(todo.STATUS_WAITING),
	}

	workCmd = &cobra.Command{
		Use:   "work",
		Short: "todo set working state",
		RunE:  getUpdatingStateAction(todo.STATUS_WORKING),
	}

	doneCmd = &cobra.Command{
		Use:   "done",
		Short: "todo set done state",
		RunE:  getUpdatingStateAction(todo.STATUS_DONE),
	}
)

func init() {
	rootCmd.AddCommand(stateCmd)
	stateCmd.AddCommand(waitCmd)
	stateCmd.AddCommand(workCmd)
	stateCmd.AddCommand(doneCmd)
}

func getUpdatingStateAction(state string) func(c *cobra.Command, args []string) error {
	return func(c *cobra.Command, args []string) error {
		id, err := c.Flags().GetString("id")
		if err != nil {
			return err
		}

		return updateState(id, state)
	}
}

func updateState(id string, status string) error {
	todo, err := collection.GetTodo(id)
	if err != nil {
		return err
	}

	todo.Status = status

	return save()
}
