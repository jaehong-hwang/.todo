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
	todoFile       = file.FindTodoFile()
	todoSystemFile = file.FindTodoSystemFile()
	collection     = todo.NewTodoCollection(todoFile)
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
}

func requireArgs(c *cobra.Command, args []string) error {
	if len(args) < 1 || strings.TrimSpace(args[0]) == "" {
		return errors.New("requires a state argument")
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().Bool("get-json", false, "if you want response type json")
	rootCmd.PersistentFlags().Int("id", 0, "todo item id")
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
