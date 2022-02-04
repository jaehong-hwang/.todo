package cmd

import (
	"fmt"
	"os"

	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/response"
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
			if todoFile.IsExist() {
				return errors.New("todo_already_exists")
			}

			dir, err := os.Getwd()
			if err != nil {
				return err
			}

			err = todoFile.CreateIfNotExist()
			if err != nil {
				return err
			}

			system.AddDirectory(dir)

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

			return todoFile.Remove()
		},
	}
)

func init() {
	rootCmd.AddCommand(collectionCmd)
	collectionCmd.AddCommand(collectionInitCmd)
	collectionCmd.AddCommand(collectionListCmd)
	collectionCmd.AddCommand(collectionRemoveCmd)
}
