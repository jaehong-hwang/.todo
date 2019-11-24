package cli

import (
	"errors"
	"os"

	"github.com/jaehong-hwang/todo/response"
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
			Action:  a.list,
		},
		{
			Name:   "init",
			Usage:  "set up todo for current directory",
			Action: a.init,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add todo",
			Action:  a.add,
		},
	}
}

func (a *App) list(c *cli.Context) error {
	a.response = &response.ListResponse{Todos: a.collection.Todos}
	return nil
}

func (a *App) init(c *cli.Context) error {
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

	a.response = &response.MessageResponse{Message: "todo init complete"}
	return nil
}

func (a *App) add(c *cli.Context) error {
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
}
