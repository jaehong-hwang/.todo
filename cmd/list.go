package cmd

import (
	"time"

	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/file"
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

		orderBy, err := c.Flags().GetString("order-by")
		if err != nil {
			return err
		}

		isAll, err := c.Flags().GetBool("all")
		if err != nil {
			return err
		} else if isAll {
			collection = &t.Collection{}
			for _, dir := range system.Directories {
				tf := file.FindTodoFileWithDirectory(dir, true)
				c := t.NewTodoCollection(tf)
				collection.Todos = append(collection.Todos, c.Todos...)
			}
		}

		col := filter.Run(collection)
		err = col.Sort(orderBy)
		if err != nil {
			return err
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
	listCmd.PersistentFlags().String("order-by", "regist-date", "order by some field")
	listCmd.PersistentFlags().Bool("all", false, "showing list for all collections")

	rootCmd.AddCommand(listCmd)
}
