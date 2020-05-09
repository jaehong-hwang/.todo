package cli

import (
	"os"

	"github.com/jaehong-hwang/todo/file"
	"github.com/jaehong-hwang/todo/response"
	"github.com/jaehong-hwang/todo/errors"
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
				return errors.New("todo_already_exists")
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
		Flags: []cli.Flag{ withDoneFlag, statusFlag },
		Aliases: []string{"l"},
		Usage:   "Print todos to the list",
		Action: func(c *cli.Context) error {
			var todos t.Todos

			status := c.String("status")
			
			if c.Bool("with-done") {
				todos = collection.Todos
			} else if status != "" {
				todos = collection.GetTodosByStatus([]string{ status })
			} else {
				todos = collection.GetTodosByStatus([]string{ t.StatusWaiting, t.StatusWorking })
			}

			appResponse = &response.ListResponse{Todos: todos}
			return nil
		},
	}

	addCommand = &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add todo",
		Action: func(c *cli.Context) error {
			if todoFile == nil {
				return errors.New("todo_doesnt_exists")
			}

			if c.NArg() == 0 {
				return errors.New("message_required")
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

	updateCommand = &cli.Command{
		Name:    "update",
		Flags: []cli.Flag{ idFlag },
		Aliases: []string{"u"},
		Usage:   "update todo message",
		Action: func(c *cli.Context) error {
			if todoFile == nil {
				return errors.New("todo_doesnt_exists")
			}

			if c.NArg() == 0 {
				return errors.New("message_required")
			}

			id := c.Int("id")
			todo := &collection.Todos[id]
			todo.Content = c.Args().Get(0)

			content, err := collection.GetTodosByJSONString()
			if err != nil {
				return err
			}

			return todoFile.FillContent(content)
		},
	}
)
