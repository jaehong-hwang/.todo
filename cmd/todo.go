package cmd

import (
	"fmt"
	"strings"
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
			status, err := c.Flags().GetString("status")
			if err != nil {
				return err
			}

			if status != "" {
				collection.Filter.Status = []string{status}
			}

			withDone, err := c.Flags().GetBool("with-done")
			if err != nil {
				return err
			}

			collection.Filter.WithDone = withDone

			todos := collection.GetList()

			if len(todos) == 0 {
				return errors.New("todo_empty")
			}

			appResponse = &response.ListResponse{Collection: collection}
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

			setTodoFlagAttr(c, &todo)

			collection.Add(todo)

			return save()
		},
	}

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "update todo message",
		RunE: func(c *cobra.Command, args []string) error {
			id, err := c.Flags().GetInt("id")
			if err != nil {
				return err
			}

			todo, err := collection.GetTodo(id)
			if err != nil {
				return err
			}

			if len(args) > 0 && strings.TrimSpace(args[0]) != "" {
				todo.Content = args[0]
			}

			setTodoFlagAttr(c, todo)

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

func setTodoFlagAttr(c *cobra.Command, todo *t.Todo) error {
	level, err := c.Flags().GetInt("level")
	if err != nil {
		return err
	} else if level > 0 {
		todo.Level = level
	}

	status, err := c.Flags().GetString("status")
	if err != nil {
		return err
	} else if err = t.IsValidStatus(status); err != nil {
		return err
	} else {
		todo.Status = status
	}

	dueDate, err := c.Flags().GetString("due-date")
	if err != nil {
		return err
	} else if dueDate != "" {
		layout := "2006-01-02T15:04:05.000Z"
		todoTime, err := time.Parse(layout, dueDate)
		if err != nil {
			return err
		}

		todo.DueDate = todoTime
	}

	return nil
}
