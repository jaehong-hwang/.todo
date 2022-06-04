package cmd

import (
	"fmt"
	"strings"
	"time"

	t "github.com/jaehong-hwang/todo/todo"
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add todo",
		Args:  requireArgs,
		RunE: func(c *cobra.Command, args []string) error {
			todo := collection.NewTodo()
			todo.Content = args[0]
			todo.Status = t.STATUS_WAITING
			todo.Author = system.Author.Name
			todo.AuthorEmail = system.Author.Email
			todo.RegistDate = time.Now()

			handleRepeatFlag(c, &todo)
			setTodoFlagAttr(c, &todo)

			collection.Add(todo)

			return save()
		},
	}

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "update todo message",
		RunE: func(c *cobra.Command, args []string) error {
			id, err := c.Flags().GetString("id")
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

			if err = handleRepeatFlag(c, todo); err != nil {
				return err
			}

			if err = setTodoFlagAttr(c, todo); err != nil {
				return err
			}

			return save()
		},
	}

	removeCmd = &cobra.Command{
		Use:   "remove",
		Short: "remove todo message",
		RunE: func(c *cobra.Command, args []string) error {
			id, err := c.Flags().GetString("id")
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
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(removeCmd)

	addCmd.PersistentFlags().String("repeat", "", "set repeat information")
	updateCmd.PersistentFlags().String("repeat", "", "set repeat information")
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
	} else if status != "" {
		if err = t.IsValidStatus(status); err != nil {
			return err
		} else {
			todo.Status = status
		}
	}

	dueDate, err := c.Flags().GetString("due-date")
	if err != nil {
		return err
	} else if dueDate != "" {
		layout := "2006-01-02"
		todoTime, err := time.Parse(layout, dueDate)
		if err != nil {
			return err
		}

		todo.DueDate = todoTime
	}

	return nil
}

func handleRepeatFlag(c *cobra.Command, todo *t.Todo) error {
	repeat, err := c.Flags().GetString("repeat")
	if err != nil {
		return err
	}

	todo.Repeat = t.Repeat{
		Types: repeat,
		Data:  nil,
	}

	return nil
}
