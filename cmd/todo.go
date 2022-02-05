package cmd

import (
	"fmt"
	"time"

	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/response"
	t "github.com/jaehong-hwang/todo/todo"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Print todos to the list",
		RunE: func(c *cobra.Command, args []string) error {
			var col t.Collection

			status, err := c.Flags().GetString("status")
			if err != nil {
				return err
			}

			withDone, err := c.Flags().GetBool("with-done")
			if err != nil {
				return err
			}

			if withDone == true {
				col = *collection
			} else if status != "" {
				col = collection.SearchByStatus([]string{status})
			} else {
				col = collection.SearchByStatus([]string{t.StatusWaiting, t.StatusWorking})
			}

			if len(col.Todos) == 0 {
				return errors.New("todo_empty")
			}

			appResponse = &response.ListResponse{Collection: col}
			return nil
		},
	}

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add todo",
		Args:  requireArgs,
		RunE: func(c *cobra.Command, args []string) error {
			todo := collection.NewTodo()
			todo.Content = args[0]
			todo.Status = t.StatusWaiting
			todo.Author = system.Author.Name
			todo.AuthorEmail = system.Author.Email
			todo.RegistDate = time.Now()

			collection.Add(todo)

			return save()
		},
	}

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "update todo message",
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

			todo.Content = args[0]

			return save()
		},
	}

	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "remove todo message",
		RunE: func(c *cobra.Command, args []string) error {
			id, err := c.Flags().GetInt("id")
			if err != nil {
				return err
			}

			todo, err := collection.GetTodo(id)
			if err != nil {
				return err
			}

			yn := "y"
			fmt.Print("Do you want remove this todo?\nContent: ", todo.Content, " (y, n): ")
			fmt.Scanln(&yn)
			if yn != "y" && yn != "Y" {
				return nil
			}

			collection.Remove(id)

			return save()
		},
	}
)

func init() {
	listCmd.PersistentFlags().String("status", "", "search status")
	listCmd.PersistentFlags().Bool("with-done", false, "showing list with done status todo")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(removeCmd)
}
