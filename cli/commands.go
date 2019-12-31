package cli

import (
	"errors"
	"os"

	"github.com/jaehong-hwang/todo/file"
	"github.com/jaehong-hwang/todo/response"
	t "github.com/jaehong-hwang/todo/todo"
	"github.com/urfave/cli/v2"
)

// TodoCommands are collection commands
type TodoCommands []*cli.Command

var (
	initCommand = &cli.Command{
		Name:  "init",
		Usage: "set up todo for current directory",
		Action: func(c *cli.Context) error {
			if todoFile != nil {
				return errors.New("todo collection already exists")
			}

			dir, err := os.Getwd()
			if err != nil {
				return err
			}

			err = file.CreateTodoFile(dir)
			if err != nil {
				return err
			}

			appResponse = &response.MessageResponse{Message: "todo init complete"}
			return nil
		},
	}

	listCommand = &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "Print todos to the list",
		Action: func(c *cli.Context) error {
			appResponse = &response.ListResponse{Todos: collection.Todos}
			return nil
		},
	}

	addCommand = &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add todo",
		Action: func(c *cli.Context) error {
			if todoFile == nil {
				return errors.New("todo dosen't exists, you should run todo init")
			}

			todo := collection.NewTodo()
			todo.Content = c.Args().Get(0)
			todo.Status = t.StatusWaiting

			collection.Add(todo)

			content, err := collection.GetTodosByJSONString()
			if err != nil {
				return err
			}

			return todoFile.FillContent(content)
		},
	}
)
