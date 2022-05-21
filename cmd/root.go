package cmd

import (
	"strings"

	"github.com/jaehong-hwang/todo/errors"
	"github.com/jaehong-hwang/todo/file"
	"github.com/jaehong-hwang/todo/response"
	"github.com/jaehong-hwang/todo/todo"
	"github.com/spf13/cobra"
)

var (
	todoFile       *file.File
	collection     *todo.Collection
	todoSystemFile = file.FindTodoSystemFile()
	system         = todo.NewSystem(todoSystemFile)
)

var appResponse response.Response

var rootCmd = &cobra.Command{
	Use:     "todo",
	Version: "0.0.3",
	Short:   "todo is a directory-based cli todo management tool.",
	Long: `.todo was created with a motif from git.
				It is a tool that creates a .todo file that manages to-dos by directory, shows list, and manages.
				For detailed explanation, see https://jaehong-hwang.github.com/todo`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		dir, err := cmd.Flags().GetString("directory")
		if err != nil {
			return err
		} else if dir != "" {
			todoFile = file.FindTodoFileWithDirectory(dir, false)
			collection = todo.NewTodoCollection(todoFile)
		} else {
			todoFile = file.FindTodoFile(true)
			collection = todo.NewTodoCollection(todoFile)
		}

		return nil
	},
}

func requireArgs(c *cobra.Command, args []string) error {
	if len(args) < 1 || strings.TrimSpace(args[0]) == "" {
		return errors.New("requires a state argument")
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().String("directory", "", "running directory")
	rootCmd.PersistentFlags().Bool("get-json", false, "if you want response type json")
	rootCmd.PersistentFlags().Int("id", 0, "todo item id")
	rootCmd.PersistentFlags().String("status", "", "todo item status")
	rootCmd.PersistentFlags().Int("level", 0, "todo item level")
	rootCmd.PersistentFlags().String("due-date", "", "todo item's due-date")
}

func Execute() (response.Response, bool) {
	if err := rootCmd.Execute(); err != nil {
		appResponse = &response.ErrorResponse{Err: err}
	}

	isJson, err := rootCmd.Flags().GetBool("get-json")
	if err != nil {
		appResponse = &response.ErrorResponse{Err: err}
	}

	return appResponse, isJson
}
