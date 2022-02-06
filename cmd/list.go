package cmd

import (
	"sort"
	"time"

	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/response"
	t "github.com/jaehong-hwang/todo/todo"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print todos to the list",
	RunE: func(c *cobra.Command, args []string) error {
		status, err := c.Flags().GetString("status")
		filter := t.Filters{}
		if err != nil {
			return err
		} else if status != "" {
			filter.Status = []string{status}
		}

		withDone, err := c.Flags().GetBool("with-done")
		if err != nil {
			return err
		}

		author, err := c.Flags().GetString("author")
		if err != nil {
			return err
		}

		dueDateStart, err := c.Flags().GetString("due-date-start")
		if err != nil {
			return err
		} else if dueDateStart != "" {
			layout := "2006-01-02"
			dueDateStartTime, err := time.Parse(layout, dueDateStart)
			if err != nil {
				return err
			}

			filter.DueDateStart = dueDateStartTime
		}

		dueDateEnd, err := c.Flags().GetString("due-date-end")
		if err != nil {
			return err
		} else if dueDateEnd != "" {
			layout := "2006-01-02"
			dueDateEndTime, err := time.Parse(layout, dueDateEnd)
			if err != nil {
				return err
			}

			filter.DueDateEnd = dueDateEndTime
		}

		filter.WithDone = withDone
		filter.Author = author

		col := filter.Run(collection)

		orderBy, err := c.Flags().GetString("order-by")
		if err != nil {
			return err
		} else {
			sort.Slice(col.Todos, func(i, j int) bool {
				switch orderBy {
				case "level":
					return col.Todos[i].Level > col.Todos[j].Level
				case "due-date":
					return (col.Todos[i].DueDate.Unix() > 0 && col.Todos[j].DueDate.Unix() < 0) || col.Todos[i].DueDate.Unix() < col.Todos[j].DueDate.Unix()
				default:
					return col.Todos[i].RegistDate.Unix() < col.Todos[j].RegistDate.Unix()
				}
			})
		}

		if len(col.Todos) == 0 {
			return errors.New("todo_empty")
		}

		appResponse = &response.ListResponse{Collection: col}
		return nil
	},
}

func init() {
	listCmd.PersistentFlags().String("status", "", "search status")
	listCmd.PersistentFlags().Bool("with-done", false, "showing list with done status todo")
	listCmd.PersistentFlags().String("author", "", "search author name or email")
	listCmd.PersistentFlags().String("due-date-start", "", "search due-date start time")
	listCmd.PersistentFlags().String("due-date-end", "", "search due-date start end")
	listCmd.PersistentFlags().String("order-by", "", "order by some field")

	rootCmd.AddCommand(listCmd)
}
