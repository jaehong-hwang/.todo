package cmd

import (
	t "github.com/jaehong-hwang/todo/todo"
	"github.com/spf13/cobra"
)

var (
	labelAddCmd = &cobra.Command{
		Use:   "add-label",
		Short: "add label to todo",
		Args:  requireArgs,
		RunE: func(c *cobra.Command, args []string) error {
			id, err := c.Flags().GetInt("id")
			if err != nil {
				return err
			}

			todo, err := collection.GetTodo(id)
			if err != nil {
				return err
			}

			labelText := args[0]
			label := t.Label{
				Text: labelText,
			}

			err = todo.AddLabel(&label)
			if err != nil {
				return err
			}

			return save()
		},
	}

	labelRemoveCmd = &cobra.Command{
		Use:   "remove-label",
		Short: "remove label from todo",
		Args:  requireArgs,
		RunE: func(c *cobra.Command, args []string) error {
			id, err := c.Flags().GetInt("id")
			if err != nil {
				return err
			}

			todo, err := collection.GetTodo(id)
			if err != nil {
				return err
			}

			labelText := args[0]
			err = todo.RemoveLabel(labelText)
			if err != nil {
				return err
			}

			return save()
		},
	}
)

func init() {
	rootCmd.AddCommand(labelAddCmd)
	rootCmd.AddCommand(labelRemoveCmd)
}
