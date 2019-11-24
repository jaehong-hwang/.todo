package main

import (
	"errors"
	"os"

	"github.com/urfave/cli/v2"
)

// TodoCommands are collection commands
type TodoCommands []*cli.Command

// GetCommands are making todo collection
func (a *App) GetCommands() TodoCommands {
	return []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "print todos to the list",
			Action: func(c *cli.Context) error {
				ResponseChan <- &ListResponse{todos: a.collection.Todos}
				return nil
			},
		},
		{
			Name:  "init",
			Usage: "set up todo for current directory",
			Action: func(c *cli.Context) error {
				if a.file.IsExists() {
					return errors.New("todo collection already exists")
				}

				dir, err := os.Getwd()
				if err != nil {
					return err
				}

				err = a.file.CreateFile(dir)
				if err != nil {
					return err
				}

				ResponseChan <- &MessageResponse{message: "todo init complete"}
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add todo",
			Action: func(c *cli.Context) error {
				if !a.file.IsExists() {
					return errors.New("todo dosen't exists, you should run todo init")
				}

				todo := a.collection.NewTodo()
				todo.Content = c.Args().Get(0)

				a.collection.Add(todo)

				content, err := a.collection.GetTodosByJSONString()
				if err != nil {
					return err
				}

				a.file.FillContent(content)

				return nil
			},
		},
	}
}
