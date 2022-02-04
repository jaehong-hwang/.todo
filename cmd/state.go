package cmd

import (
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
		RunE:  getUpdatingStateAction("wait"),
	}

	workCmd = &cobra.Command{
		Use:   "work",
		Short: "todo set working state",
		RunE:  getUpdatingStateAction("work"),
	}

	doneCmd = &cobra.Command{
		Use:   "done",
		Short: "todo set done state",
		RunE:  getUpdatingStateAction("done"),
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
		id, err := c.Flags().GetInt("id")
		if err != nil {
			return err
		}

		return updateState(id, state)
	}
}

func updateState(id int, status string) error {
	todo, err := collection.GetTodo(id)
	if err != nil {
		return err
	}

	todo.Status = status
	content, err := collection.GetTodosJSONString()
	if err != nil {
		return err
	}

	todoFile.FillContent(content)

	return nil
}
