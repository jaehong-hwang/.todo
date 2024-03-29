package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/jaehong-hwang/todo/errors"
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

			if err := handleTodoFlags(c, &todo); err != nil {
				return err
			}

			collection.Add(todo)

			todoFile.AddLog(todo.ID, "regist", "")

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

			if err = handleTodoFlags(c, todo); err != nil {
				return err
			}

			todoFile.AddLog(todo.ID, "updated", "")

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

			todoFile.AddLog(todo.ID, "removed", "")

			return save()
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(removeCmd)

	for _, statuses := range t.RepeatStatuses {
		addCmd.PersistentFlags().String(statuses, "", "set repeat "+statuses+" information")
		updateCmd.PersistentFlags().String(statuses, "", "set repeat "+statuses+" information")
	}
}

func handleTodoFlags(c *cobra.Command, todo *t.Todo) error {
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

	for _, statuses := range t.RepeatStatuses {
		repeat, err := c.Flags().GetString(statuses)
		if err != nil {
			return err
		}

		if repeat == "" {
			continue
		}

		if todo.Repeat != nil {
			return errors.New("repeat_set_only_one")
		}

		todo.Repeat = &t.Repeat{
			Types: statuses,
			Data:  strings.Split(repeat, ","),
		}
	}

	return nil
}
