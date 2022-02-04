package cmd

import (
	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/todo"
	"github.com/spf13/cobra"
)

var (
	stateCmd = &cobra.Command{
		Use:   "state",
		Short: "update state",
		Args: func(c *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a state argument")
			}
			if todo.IsValidStatus(args[0]) {
				return nil
			}
			return errors.NewWithParam("unexpected_state", map[string]string{
				"state": args[0],
			})
		},
		RunE: func(c *cobra.Command, args []string) error {
			id, err := c.Flags().GetInt("id")
			if err != nil {
				return err
			}

			status := args[0]

			return updateState(id, status)
		},
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
	rootCmd.AddCommand(waitCmd)
	rootCmd.AddCommand(workCmd)
	rootCmd.AddCommand(doneCmd)
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
