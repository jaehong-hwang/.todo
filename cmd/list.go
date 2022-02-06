package cmd

import (
	"time"

	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/response"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print todos to the list",
	RunE: func(c *cobra.Command, args []string) error {
		status, err := c.Flags().GetString("status")
		if err != nil {
			return err
		} else if status != "" {
			collection.Filter.Status = []string{status}
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

			collection.Filter.DueDateStart = dueDateStartTime
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

			collection.Filter.DueDateEnd = dueDateEndTime
		}

		collection.Filter.WithDone = withDone
		collection.Filter.Author = author

		todos := collection.GetList()

		if len(todos) == 0 {
			return errors.New("todo_empty")
		}

		appResponse = &response.ListResponse{Collection: collection}
		return nil
	},
}

func init() {
	listCmd.PersistentFlags().String("status", "", "search status")
	listCmd.PersistentFlags().Bool("with-done", false, "showing list with done status todo")
	listCmd.PersistentFlags().String("author", "", "search author name or email")
	listCmd.PersistentFlags().String("due-date-start", "", "search due-date start time")
	listCmd.PersistentFlags().String("due-date-end", "", "search due-date start end")

	rootCmd.AddCommand(listCmd)
}
