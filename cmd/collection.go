package cmd

import (
	"fmt"

	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/file"
	"github.com/jaehong-hwang/todo/response"
	"github.com/jaehong-hwang/todo/todo"
	"github.com/spf13/cobra"
)

var (
	collectionCmd = &cobra.Command{
		Use:   "collection",
		Short: "managing collections",
	}

	collectionInitCmd = &cobra.Command{
		Use:   "init",
		Short: "set up todo for current directory",
		RunE: func(c *cobra.Command, args []string) error {
			dir := file.GetCurrentDirectory()
			newTodoFile := file.FindTodoWorkspace(dir, false)

			if newTodoFile.IsExist() {
				return errors.New("todo_already_exists")
			}

			err := newTodoFile.CreateIfNotExist()
			if err != nil {
				return err
			}

			name, err := c.Flags().GetString("name")
			if err != nil {
				return err
			}

			system.AddDirectory(todo.Directory{
				Name: name,
				Path: dir,
			})

			appResponse = &response.MessageResponse{Message: "todo init complete"}
			return nil
		},
	}

	collectionListCmd = &cobra.Command{
		Use:   "list of collection directories",
		Short: "Print directories of todo collection",
		Run: func(c *cobra.Command, args []string) {
			appResponse = &response.DirectoryResponse{Directories: system.Directories}
		},
	}

	collectionRemoveCmd = &cobra.Command{
		Use:   "remove",
		Short: "remove current todo collection",
		RunE: func(c *cobra.Command, args []string) error {
			yn := "y"
			fmt.Print("Do you want remove current todo collection? (y, n): ")
			fmt.Scanln(&yn)
			if yn != "y" && yn != "Y" {
				return nil
			}

			system.RemoveDirectory(todoFile.GetDirectory())

			return todoFile.Remove()
		},
	}
)

func init() {
	rootCmd.AddCommand(collectionCmd)
	rootCmd.AddCommand(collectionInitCmd)
	collectionCmd.AddCommand(collectionInitCmd)
	collectionCmd.AddCommand(collectionListCmd)
	collectionCmd.AddCommand(collectionRemoveCmd)

	collectionInitCmd.PersistentFlags().String("name", "", "alias directory name")
}
